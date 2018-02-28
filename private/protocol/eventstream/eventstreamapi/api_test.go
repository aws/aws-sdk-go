package eventstreamapi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

func TestEventReader(t *testing.T) {
	stream := createStream(
		eventstream.Message{
			Headers: eventstream.Headers{
				eventstream.Header{
					Name:  EventTypeHeader,
					Value: eventstream.StringValue("eventABC"),
				},
			},
		},
		eventstream.Message{
			Headers: eventstream.Headers{
				eventstream.Header{
					Name:  EventTypeHeader,
					Value: eventstream.StringValue("eventEFG"),
				},
			},
		},
	)

	var unmarshalers request.HandlerList
	unmarshalers.PushBackNamed(restjson.UnmarshalHandler)

	eventReader := NewEventReader(stream,
		protocol.HandlerPayloadUnmarshal{
			Unmarshalers: unmarshalers,
		},
		unmarshalerForEventType,
	)

	event, err := eventReader.ReadEvent()
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if event == nil {
		t.Fatalf("expect event got none")
	}

	event, err = eventReader.ReadEvent()
	if err == nil {
		t.Fatalf("expect error for unknown event, got none")
	}

	if event != nil {
		t.Fatalf("expect no event, got %T, %v", event, event)
	}

	fmt.Println(event)
}

func BenchmarkEventReader(b *testing.B) {
	var buf bytes.Buffer
	encoder := eventstream.NewEncoder(&buf)
	msg := eventstream.Message{
		Headers: eventstream.Headers{
			eventstream.Header{
				Name:  EventTypeHeader,
				Value: eventstream.StringValue("eventABC"),
			},
		},
	}
	if err := encoder.Encode(msg); err != nil {
		b.Fatalf("failed to encode message, %v", err)
	}
	stream := bytes.NewReader(buf.Bytes())

	var unmarshalers request.HandlerList
	unmarshalers.PushBackNamed(restjson.UnmarshalHandler)

	eventReader := NewEventReader(ioutil.NopCloser(stream),
		protocol.HandlerPayloadUnmarshal{
			Unmarshalers: unmarshalers,
		},
		unmarshalerForEventType,
	)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stream.Seek(0, 0)

		event, err := eventReader.ReadEvent()
		if err != nil {
			b.Fatalf("expect no error, got %v", err)
		}
		if event == nil {
			b.Fatalf("expect event got none")
		}
	}
}

func unmarshalerForEventType(eventType string) (EventUnmarshaler, error) {
	switch eventType {
	case "eventABC":
		return &eventABC{}, nil
	default:
		return nil, fmt.Errorf("unknown event type, %v", eventType)
	}
}

type eventABC struct {
	_ struct{}

	HeaderField string
	Payload     []byte
}

func (e *eventABC) UnmarshalEvent(
	unmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	return nil
}

func createStream(msgs ...eventstream.Message) io.ReadCloser {
	w := bytes.NewBuffer(nil)

	encoder := eventstream.NewEncoder(w)

	for _, msg := range msgs {
		if err := encoder.Encode(msg); err != nil {
			panic("createStream failed, " + err.Error())
		}
	}

	return ioutil.NopCloser(w)
}
