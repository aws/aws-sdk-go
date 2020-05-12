// +build codegen

package api

import "text/template"

var eventStreamShapeReaderTmpl = template.Must(template.New("eventStreamShapeReaderTmpl").
	Funcs(template.FuncMap{}).
	Parse(`
{{- $es := $.EventStream }}

// {{ $es.StreamReaderAPIName }} provides the interface for reading to the stream. The
// default implementation for this interface will be {{ $.ShapeName }}.
//
// The reader's Close method must allow multiple concurrent calls.
//
// These events are:
// {{ range $_, $event := $es.Events }}
//     * {{ $event.Shape.ShapeName }}
{{- end }}
type {{ $es.StreamReaderAPIName }} interface {
	// Returns a channel of events as they are read from the event stream.
	Events() <-chan {{ $es.EventGroupName }}

	// Close will stop the reader reading events from the stream.
	Close() error

	// Returns any error that has occurred while reading from the event stream.
	Err() error
}

type {{ $es.StreamReaderImplName }} struct {
	eventReader *eventstreamapi.EventReader
	stream      chan {{ $es.EventGroupName }}
	err         *eventstreamapi.OnceError

	done      chan struct{}
	closeOnce sync.Once
}

func {{ $es.StreamReaderImplConstructorName }}(eventReader *eventstreamapi.EventReader) *{{ $es.StreamReaderImplName }} {
	r := &{{ $es.StreamReaderImplName }}{
		eventReader: eventReader,
		stream: make(chan {{ $es.EventGroupName }}),
		done:   make(chan struct{}),
		err:    eventstreamapi.NewOnceError(),
	}
	go r.readEventStream()

	return r
}

// Close will close the underlying event stream reader.
func (r *{{ $es.StreamReaderImplName }}) Close() error {
	r.closeOnce.Do(r.safeClose)
	return r.Err()
}

func (r *{{ $es.StreamReaderImplName }}) ErrorSet() <-chan struct{} {
	return r.err.ErrorSet()
}

func (r *{{ $es.StreamReaderImplName }}) Closed() <-chan struct{} {
	return r.done
}

func (r *{{ $es.StreamReaderImplName }}) safeClose() {
	close(r.done)
}

func (r *{{ $es.StreamReaderImplName }}) Err() error {
	return r.err.Err()
}

func (r *{{ $es.StreamReaderImplName }}) Events() <-chan {{ $es.EventGroupName }} {
	return r.stream
}

func (r *{{ $es.StreamReaderImplName }}) readEventStream() {
	defer r.Close()
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
			r.err.SetError(err)
			return
		}

		select {
		case r.stream <- event.({{ $es.EventGroupName }}):
		case <-r.done:
			return
		}
	}
}

type {{ $es.StreamUnmarshalerForEventName }} struct {
	metadata protocol.ResponseMetadata
}

func (u {{ $es.StreamUnmarshalerForEventName }}) UnmarshalerForEventName(eventType string) (eventstreamapi.Unmarshaler, error) {
	switch eventType {
		{{- range $_, $event := $es.Events }}
			case {{ printf "%q" $event.Name }}:
				return &{{ $event.Shape.ShapeName }}{}, nil
		{{- end }}
		{{- range $_, $event := $es.Exceptions }}
			case {{ printf "%q" $event.Name }}:
				return newError{{ $event.Shape.ShapeName }}(u.metadata).(eventstreamapi.Unmarshaler), nil
		{{- end }}
	default:
		return nil, awserr.New(
			request.ErrCodeSerialization,
			fmt.Sprintf("unknown event type name, %s, for {{ $es.Name }}", eventType),
			nil,
		)
	}
}
`))
