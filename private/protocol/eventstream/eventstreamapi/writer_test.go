package eventstreamapi

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

func TestEventWriter(t *testing.T) {
	cases := map[string]struct {
		Event  Marshaler
		Expect eventstream.Message
	}{
		"structured event": {
			Event: &eventStructured{
				String: aws.String("stringfield"),
				Number: aws.Int64(123),
				Nested: &eventStructured{
					String: aws.String("fieldstring"),
					Number: aws.Int64(321),
				},
			},
			Expect: eventstream.Message{
				Headers: eventstream.Headers{
					eventMessageTypeHeader,
					eventstream.Header{
						Name:  EventTypeHeader,
						Value: eventstream.StringValue("eventStructured"),
					},
				},
				Payload: []byte(`{"String":"stringfield","Number":123,"Nested":{"String":"fieldstring","Number":321}}`),
			},
		},
	}

	var marshalers request.HandlerList
	marshalers.PushBackNamed(restjson.BuildHandler)

	var stream bytes.Buffer
	eventWriter := NewEventWriter(&stream,
		protocol.HandlerPayloadMarshal{
			Marshalers: marshalers,
		},
	)

	decoder := eventstream.NewDecoder(&stream)

	decodeBuf := make([]byte, 1024)
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			eventWriter.UseLogger(t, aws.LogDebugWithEventStreamBody)

			if err := eventWriter.WriteEvent(c.Event); err != nil {
				t.Fatalf("expect no write error, got %v", err)
			}

			msg, err := decoder.Decode(decodeBuf)
			if err != nil {
				t.Fatalf("expect no decode error got, %v", err)
			}

			if e, a := c.Expect, msg; !reflect.DeepEqual(e, a) {
				t.Errorf("expect:%v\nactual:%v\n", e, a)
			}
		})
	}
}

func BenchmarkEventWriter(b *testing.B) {
	var marshalers request.HandlerList
	marshalers.PushBackNamed(restjson.BuildHandler)

	var stream bytes.Buffer
	eventWriter := NewEventWriter(&stream,
		protocol.HandlerPayloadMarshal{
			Marshalers: marshalers,
		},
	)

	event := &eventStructured{
		String: aws.String("stringfield"),
		Number: aws.Int64(123),
		Nested: &eventStructured{
			String: aws.String("fieldstring"),
			Number: aws.Int64(321),
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := eventWriter.WriteEvent(event); err != nil {
			b.Fatalf("expect no write error, got %v", err)
		}
	}
}

type eventStructured struct {
	_ struct{} `type:"structure"`

	String *string          `type:"string"`
	Number *int64           `type:"long"`
	Nested *eventStructured `type:"structure"`
}

func (e *eventStructured) MarshalEvent(pm protocol.PayloadMarshaler) (eventstream.Message, error) {
	var msg eventstream.Message
	msg.Headers.Set(MessageTypeHeader, eventstream.StringValue(EventMessageType))
	msg.Headers.Set(EventTypeHeader, eventstream.StringValue("eventStructured"))

	var buf bytes.Buffer
	if err := pm.MarshalPayload(&buf, e); err != nil {
		return eventstream.Message{}, err
	}

	msg.Payload = buf.Bytes()

	return msg, nil
}

func (e *eventStructured) UnmarshalEvent(pm protocol.PayloadUnmarshaler, msg eventstream.Message) error {
	return pm.UnmarshalPayload(bytes.NewReader(msg.Payload), e)
}
