package eventstreamapi

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
)

// Marshaler provides a marshaling interface for event types to event stream
// messages.
type Marshaler interface {
	MarshalEvent(protocol.PayloadMarshaler) (eventstream.Message, error)
}

// EventWriter provides a wrapper around the underlying event stream encoder
// for an io.Writer.
type EventWriter struct {
	writer  io.Writer
	encoder *eventstream.Encoder
	signer  *MessageSigner

	payloadMarshaler protocol.PayloadMarshaler
}

// NewEventWriter returns a new event stream writer, that will write to the
// writer provided. Use the WriteStream method to write an event to the stream.
func NewEventWriter(writer io.Writer,
	payloadMarshaler protocol.PayloadMarshaler,
	signer *MessageSigner,
) *EventWriter {
	return &EventWriter{
		writer:           writer,
		encoder:          eventstream.NewEncoder(writer),
		payloadMarshaler: payloadMarshaler,
		signer:           signer,
	}
}

// UseLogger instructs the EventWriter to use the logger and log level
// specified.
func (w *EventWriter) UseLogger(logger aws.Logger, logLevel aws.LogLevelType) {
	if logger != nil && logLevel.Matches(aws.LogDebugWithEventStreamBody) {
		w.encoder.UseLogger(logger)
	}
}

func (w *EventWriter) signMessage(msg *eventstream.Message) error {
	if w.signer == nil {
		return nil
	}

	return w.signer.SignMessage(msg, timeNow())
}

// WriteEvent writes an event to the stream. Returns an error if the event
// fails to marshal into a message, or writing to the underlying writer fails.
func (w *EventWriter) WriteEvent(event Marshaler) error {
	msg, err := event.MarshalEvent(w.payloadMarshaler)
	if err != nil {
		return err
	}

	if err = w.signMessage(&msg); err != nil {
		return err
	}

	return w.encoder.Encode(msg)
}
