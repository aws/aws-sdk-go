// +build codegen

package api

import "text/template"

var eventStreamAPIWriterTmpl = template.Must(template.New("eventStreamAPIWriterTmpl").
	Funcs(template.FuncMap{}).
	Parse(`
{{- $eventStream := $.EventStream }}
{{- $eventStreamEventGroup := printf "%sEvent" $eventStream.Name }}
{{- $esWriterAPI := printf "%sWriter" $eventStream.Name }}
{{- $esWriterImpl := printf "write%s" $eventStream.Name }}

// {{ $esWriterAPI }} provides the interface for writing events to the stream.
// The default implementation for this interface will be {{ $.ShapeName }}.
//
// The writer's Close method must allow multiple concurrent calls.
//
// These events are:
// {{ range $_, $event := $eventStream.Events }}
//     * {{ $event.Shape.ShapeName }}
{{- end }}
type {{ $esWriterAPI }} interface {
	// Sends writes events to the stream blocking until the event has been
	// written. An error is returned if the write fails.
	Send({{ $eventStreamEventGroup }}) error

	// Close will stop the writer writing to the event stream.
	Close() error

	// Returns any error that has occurred while writing to the event stream.
	Err() error
}

{{- $eventGroupWrapperName := printf "outbound%s" $eventStreamEventGroup }}
type {{ $eventGroupWrapperName }} struct {
	Event {{ $eventStreamEventGroup }}
	Result chan <- error
}

type {{ $esWriterImpl }} struct {
	eventWriter *eventstreamapi.EventWriter
	stream chan {{ $eventGroupWrapperName }}
	errVal atomic.Value

	done chan struct{}
	closeOnce sync.Once
}

func newWrite{{ $eventStream.Name }}(
	writer io.Writer,
	buildHandlers request.HandlerList,
	signer *eventstreamapi.MessageSigner,
	logger aws.Logger, logLevel aws.LogLevelType,
) *{{ $esWriterImpl }} {
	w := & {{ $esWriterImpl }}{
		stream: make(chan {{ $eventGroupWrapperName }}),
		done: make(chan struct{}),
	}
	w.eventWriter = eventstreamapi.NewEventWriter(writer,
		protocol.HandlerPayloadMarshal{
			Marshalers: buildHandlers,
		},
		 signer,
	)
	w.eventWriter.UseLogger(logger, logLevel)

	return w
}

// Close will stop the writer writing to the event stream.
func (w *{{ $esWriterImpl }}) Close() error {
	w.closeOnce.Do(w.safeClose)
	return w.Err()
}

func (w *{{ $esWriterImpl }}) safeClose() {
	close(w.done)
}

// Err returns any error occurred writing events..
func (w *{{ $esWriterImpl }}) Err() error {
	if v := w.errVal.Load(); v != nil {
		return v.(error)
	}

	return nil
}

func (w *{{ $esWriterImpl }}) Send(event {{ $eventStreamEventGroup }}) error {
	if err := w.Err(); err != nil {
		return err
	}
	resultCh := make(chan error)
	wrapped := {{ $eventGroupWrapperName }}{
		Event: event,
		Result: resultCh,
	}

	select {
	case w.stream <- wrapped:
	case <- w.done:
		return fmt.Errorf("stream closed, unable to send event")
	}

	select {
	case err := <- resultCh:
		return err
	case <- w.done:
		return fmt.Errorf("stream closed, unable to send event")
	}
}

func (w *{{ $esWriterImpl }}) writeEventStream() {
	defer close(w.stream)

	for {
		select { 
		case wrapper := <- w.stream:
			err := w.eventWriter.WriteEvent(wrapper.Event)
			select {
			case wrapper.Result <- err:
			case <-w.done:
				return
			}

		case <- w.done:
			return
		}
	}
}
`))
