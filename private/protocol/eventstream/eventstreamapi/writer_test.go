// +build go1.7

package eventstreamapi

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

func TestEventWriter(t *testing.T) {
	cases := map[string]struct {
		Event    Marshaler
		Signer   *MessageSigner
		TimeFunc func() time.Time
		Expect   eventstream.Message
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
		"signed event": {
			Event: &eventStructured{
				String: aws.String("stringfield"),
				Number: aws.Int64(123),
				Nested: &eventStructured{
					String: aws.String("fieldstring"),
					Number: aws.Int64(321),
				},
			},
			Signer: &MessageSigner{Signer: &mockChunkSigner{signature: "524f1d03d1d81e94a099042736d40bd9681b867321443ff58a4568e274dbd83bff"}},
			TimeFunc: func() time.Time {
				return time.Date(2019, 1, 27, 22, 37, 54, 0, time.UTC)
			},
			Expect: eventstream.Message{
				Headers: eventstream.Headers{
					eventMessageTypeHeader,
					eventstream.Header{
						Name:  EventTypeHeader,
						Value: eventstream.StringValue("eventStructured"),
					},
					{
						Name:  DateHeader,
						Value: eventstream.TimestampValue(time.Date(2019, 1, 27, 22, 37, 54, 0, time.UTC)),
					},
					{
						Name:  ChunkSignatureHeader,
						Value: eventstream.BytesValue(mustDecodeHex(hex.DecodeString("524f1d03d1d81e94a099042736d40bd9681b867321443ff58a4568e274dbd83bff"))),
					},
				},
				Payload: []byte(`{"String":"stringfield","Number":123,"Nested":{"String":"fieldstring","Number":321}}`),
			},
		},
	}

	var marshalers request.HandlerList
	marshalers.PushBackNamed(restjson.BuildHandler)

	var stream bytes.Buffer

	decodeBuf := make([]byte, 1024)
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			defer swapTimeNow(c.TimeFunc)()

			stream.Reset()

			eventWriter := NewEventWriter(&stream,
				protocol.HandlerPayloadMarshal{
					Marshalers: marshalers,
				},
				c.Signer,
			)

			decoder := eventstream.NewDecoder(&stream)

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
		nil,
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
