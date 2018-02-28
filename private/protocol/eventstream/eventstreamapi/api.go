package eventstreamapi

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
)

// EventUnmarshaler provides the interface for unmarshaling a EventStream
// message into a SDK type.
type EventUnmarshaler interface {
	UnmarshalEvent(protocol.PayloadUnmarshaler, eventstream.Message) error
}

// EventTypeHeader is the EventStream header used by APIs to identify the event
// message type.
const EventTypeHeader = `:event-type`

// EventReader provides reading from the EventStream of an reader.
type EventReader struct {
	reader  io.ReadCloser
	decoder *eventstream.Decoder

	unmarshalerForEventType func(string) (EventUnmarshaler, error)
	payloadUnmarshaler      protocol.PayloadUnmarshaler

	payloadBuf []byte
}

// NewEventReader returns a EventReader built from the reader and unmarshaler
// provided.  Use ReadStream method to start reading from the EventStream.
func NewEventReader(
	reader io.ReadCloser,
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	unmarshalerForEventType func(string) (EventUnmarshaler, error),
) *EventReader {
	return &EventReader{
		reader:                  reader,
		decoder:                 eventstream.NewDecoder(reader),
		payloadUnmarshaler:      payloadUnmarshaler,
		unmarshalerForEventType: unmarshalerForEventType,
		payloadBuf:              make([]byte, 10*1024),
	}
}

// UseLogger instructs the EventReader to use the logger and log level
// specified.
func (s *EventReader) UseLogger(logger aws.Logger, logLevel aws.LogLevelType) {
	if logger != nil && logLevel.Matches(aws.LogDebugWithEventStreamBody) {
		s.decoder.UseLogger(logger)
	}
}

// ReadEvent attempts to read a message from the EventStream and return the
// unmarshaled event value that the message is for.
//
// EventUnmarshalers called with EventStream messages must take copies of the
// message's Payload. The payload will is reused between events read.
func (s *EventReader) ReadEvent() (event interface{}, err error) {
	//	defer func() {
	//		// Ignore errors that occur as a result of the user closing the
	//		// EventReader.
	//		// TODO this is not complete, should check for EOF
	//		if closed := atomic.LoadInt32(&s.closed); closed != 1 {
	//			err = nil
	//		}
	//	}()

	msg, err := s.decoder.Decode(s.payloadBuf)
	if err != nil {
		return nil, err
	}

	eventType, err := GetEventType(msg)
	if err != nil {
		return nil, err
	}

	ev, err := s.unmarshalerForEventType(eventType)
	if err != nil {
		return nil, err
	}

	err = ev.UnmarshalEvent(s.payloadUnmarshaler, msg)
	if err != nil {
		return nil, err
	}
	s.payloadBuf = msg.Payload[0:0]

	return ev, nil
}

// Close closes the EventReader's EventStream reader.
func (s *EventReader) Close() error {
	return s.reader.Close()
}

// GetEventType returns the type of the event message. Returns error if the
// event type was not found.
func GetEventType(msg eventstream.Message) (string, error) {
	eventType := msg.Headers.Get(EventTypeHeader)
	if eventType == nil {
		return "", fmt.Errorf("event type header %s not present", EventTypeHeader)
	}

	eventName, ok := eventType.Get().(string)
	if !ok {
		return "", fmt.Errorf("event type not string value, %T", eventType)
	}

	return eventName, nil
}
