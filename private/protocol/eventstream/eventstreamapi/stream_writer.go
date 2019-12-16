package eventstreamapi

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws"
)

type StreamWriter struct {
	eventWriter *EventWriter
	stream      chan eventWriteAsyncReport

	errVal    atomic.Value
	done      chan struct{}
	closeOnce sync.Once
}

func NewStreamWriter(eventWriter *EventWriter) *StreamWriter {
	writer := &StreamWriter{
		eventWriter: eventWriter,
		stream:      make(chan eventWriteAsyncReport),
		done:        make(chan struct{}),
	}
	go writer.writeStream()

	return writer
}

// Close terminates the writers ability to write new events to the stream. Any
// future call to Send will fail with an error.
func (w *StreamWriter) Close() error {
	w.closeOnce.Do(w.safeClose)
	return w.Err()
}

func (w *StreamWriter) safeClose() {
	close(w.done)
}

// Err returns any error that occurred while attempting to write an event to the
// stream.
func (w *StreamWriter) Err() error {
	if v := w.errVal.Load(); v != nil {
		return v.(error)
	}

	return nil
}

// Send writes a single event to the stream returning an error if the write
// failed.
//
// Send may be called concurrently. Events will be written to the stream
// safely.
func (w *StreamWriter) Send(ctx aws.Context, event Marshaler) error {
	if err := w.Err(); err != nil {
		return err
	}

	resultCh := make(chan error)
	wrapped := eventWriteAsyncReport{
		Event:  event,
		Result: resultCh,
	}

	select {
	case w.stream <- wrapped:
	case <-ctx.Done():
		return ctx.Err()
	case <-w.done:
		return fmt.Errorf("stream closed, unable to send event")
	}

	select {
	case err := <-resultCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-w.done:
		return fmt.Errorf("stream closed, unable to send event")
	}
}

func (w *StreamWriter) writeStream() {
	defer close(w.stream)

	for {
		select {
		case wrapper := <-w.stream:
			err := w.eventWriter.WriteEvent(wrapper.Event)
			wrapper.ReportResult(w.done, err)
			if err != nil {
				w.errVal.Store(err)
				return
			}

		case <-w.done:
			return
		}
	}
}

type eventWriteAsyncReport struct {
	Event  Marshaler
	Result chan<- error
}

func (e eventWriteAsyncReport) ReportResult(cancel <-chan struct{}, err error) bool {
	select {
	case e.Result <- err:
		return true
	case <-cancel:
		return false
	}
}
