// +build codegen

package api

import "text/template"

var eventStreamShapeReaderTmpl = template.Must(template.New("eventStreamShapeReaderTmpl").
	Funcs(template.FuncMap{}).
	Parse(`
{{- $eventStream := $.EventStream }}
{{- $eventStreamEventGroup := printf "%sEvent" $eventStream.Name }}
{{- $esReaderAPI := printf "%sReader" $eventStream.Name }}
{{- $esReaderImpl := printf "read%s" $eventStream.Name }}

// {{ $esReaderAPI }} provides the interface for reading to the stream. The
// default implementation for this interface will be {{ $.ShapeName }}.
//
// The reader's Close method must allow multiple concurrent calls.
//
// These events are:
// {{ range $_, $event := $eventStream.Events }}
//     * {{ $event.Shape.ShapeName }}
{{- end }}
type {{ $esReaderAPI }} interface {
	// Returns a channel of events as they are read from the event stream.
	Events() <-chan {{ $eventStreamEventGroup }}

	// Close will stop the reader reading events from the stream.
	Close() error

	// Returns any error that has occurred while reading from the event stream.
	Err() error
}


type {{ $esReaderImpl }} struct {
	eventReader *eventstreamapi.EventReader
	stream chan {{ $eventStreamEventGroup }}
	errVal atomic.Value

	done      chan struct{}
	closeOnce sync.Once
}

func newRead{{ $eventStream.Name }}(
	reader io.Reader,
	unmarshalers request.HandlerList,
	logger aws.Logger,
	logLevel aws.LogLevelType,
	unmarshalerForEvent func(string) (eventstreamapi.Unmarshaler, error),
) *{{ $esReaderImpl }} {
	r := &{{ $esReaderImpl }}{
		stream: make(chan {{ $eventStreamEventGroup }}),
		done: make(chan struct{}),
	}

	r.eventReader = eventstreamapi.NewEventReader(
		reader,
		protocol.HandlerPayloadUnmarshal{
			Unmarshalers: unmarshalers,
		},
		unmarshalerForEvent,
	)
	r.eventReader.UseLogger(logger, logLevel)

	return r
}

// Close will close the underlying event stream reader.
func (r *{{ $esReaderImpl }}) Close() error {
	r.closeOnce.Do(r.safeClose)

	return r.Err()
}

func (r *{{ $esReaderImpl }}) safeClose() {
	close(r.done)
}

func (r *{{ $esReaderImpl }}) Err() error {
	if v := r.errVal.Load(); v != nil {
		return v.(error)
	}

	return nil
}

func (r *{{ $esReaderImpl }}) Events() <-chan {{ $eventStreamEventGroup }} {
	return r.stream
}

func (r *{{ $esReaderImpl }}) readEventStream() {
	defer close(r.stream)

	for {
		event, err := r.eventReader.ReadEvent()
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				// If closed already ignore the error
				return
			default:
			}
			r.errVal.Store(err)
			return
		}

		select {
		case r.stream <- event.({{ $eventStreamEventGroup }}):
		case <-r.done:
			return
		}
	}
}

func unmarshalerFor{{ $eventStream.Name }}Event(eventType string) (eventstreamapi.Unmarshaler, error) {
	switch eventType {
		{{- range $_, $event := $eventStream.Events }}
			case {{ printf "%q" $event.Name }}:
				return &{{ $event.Shape.ShapeName }}{}, nil
		{{ end -}}
		{{- range $_, $event := $eventStream.Exceptions }}
			case {{ printf "%q" $event.Name }}:
				return &{{ $event.Shape.ShapeName }}{}, nil
		{{ end -}}
	default:
		return nil, awserr.New(
			request.ErrCodeSerialization,
			fmt.Sprintf("unknown event type name, %s, for {{ $eventStream.Name }}", eventType),
			nil,
		)
	}
}
`))
