// +build codegen

package api

import (
	"bytes"
	"fmt"
	"io"
	"text/template"
)

// EventStreamAPI provides details about the event stream async API and
// associated EventStream shapes.
type EventStreamAPI struct {
	Name      string
	Operation *Operation
	Shape     *Shape
	Inbound   *EventStream
	Outbound  *EventStream
}

// EventStream represents a single eventstream group (input/output) and the
// modeled events that are known for the stream.
type EventStream struct {
	Name   string
	Shape  *Shape
	Events []*Event
}

// Event is a single EventStream event that can be sent or received in an
// EventStream.
type Event struct {
	Name  string
	Shape *Shape
	For   *EventStream
}

// ShapeDoc returns the docstring for the EventStream API.
func (esAPI *EventStreamAPI) ShapeDoc() string {
	tmpl := template.Must(template.New("eventStreamShapeDoc").Parse(`
{{- $.Name }} provides handling of EventStreams for
the {{ $.Operation.ExportedName }} API. 
{{- if $.Inbound }}

Use this type to receive {{ $.Inbound.Name }} events. The events
can be read from the Events channel member.

The events that can be received are:
{{ range $_, $event := $.Inbound.Events }}
    * {{ $event.Shape.ShapeName }}
{{- end }}

{{- end }}

{{- if $.Outbound }}

Use this type to send {{ $.Outbound.Name }} events. The events 
can be sent with the Send method.

The events that can be sent are:
{{ range $_, $event := $.Outbound.Events -}}
    * {{ $event.Shape.ShapeName }}
{{- end }}

{{- end }}`))

	var w bytes.Buffer
	if err := tmpl.Execute(&w, esAPI); err != nil {
		panic(fmt.Sprintf("failed to generate eventstream shape template for %v, %v", esAPI.Name, err))
	}

	return commentify(w.String())
}

func eventStreamAPIShapeRefDoc(refName string) string {
	return commentify(fmt.Sprintf("Use %s to use the API's stream.", refName))
}

func (a *API) setupEventStreams() {
	const eventStreamMemberName = "EventStream"

	for _, op := range a.Operations {
		outbound := setupEventStream(op.InputRef.Shape)
		inbound := setupEventStream(op.OutputRef.Shape)

		if outbound == nil && inbound == nil {
			continue
		}

		if outbound != nil {
			panic(fmt.Sprintf("Outbound stream support not implemented, %s, %s",
				outbound.Name, outbound.Shape.ShapeName))
		}

		eventStreamAPI := &EventStreamAPI{
			Name:      op.ExportedName + eventStreamMemberName,
			Operation: op,
			Outbound:  outbound,
			Inbound:   inbound,
		}

		streamShape := &Shape{
			API:            a,
			ShapeName:      eventStreamAPI.Name,
			Documentation:  eventStreamAPI.ShapeDoc(),
			Type:           "structure",
			EventStreamAPI: eventStreamAPI,
		}
		streamShapeRef := &ShapeRef{
			API:           a,
			ShapeName:     streamShape.ShapeName,
			Shape:         streamShape,
			Documentation: eventStreamAPIShapeRefDoc(eventStreamMemberName),
		}
		streamShape.refs = []*ShapeRef{streamShapeRef}
		eventStreamAPI.Shape = streamShape

		if _, ok := op.OutputRef.Shape.MemberRefs[eventStreamMemberName]; ok {
			panic(fmt.Sprintf("shape ref already exists, %s.%s",
				op.OutputRef.Shape.ShapeName, eventStreamMemberName))
		}
		op.OutputRef.Shape.MemberRefs[eventStreamMemberName] = streamShapeRef
		op.OutputRef.Shape.EventStreamsMemberName = eventStreamMemberName
		if _, ok := a.Shapes[streamShape.ShapeName]; ok {
			panic("shape already exists, " + streamShape.ShapeName)
		}
		a.Shapes[streamShape.ShapeName] = streamShape

		a.HasEventStream = true
	}
}

func setupEventStream(topShape *Shape) *EventStream {
	var eventStream *EventStream
	for refName, ref := range topShape.MemberRefs {
		if !ref.Shape.IsEventStream {
			continue
		}
		if eventStream != nil {
			panic(fmt.Sprintf("multiple shape ref eventstreams, %s, prev: %s",
				refName, eventStream.Name))
		}

		eventStream = &EventStream{
			Name:  ref.Shape.ShapeName,
			Shape: ref.Shape,
		}
		for _, eventRefName := range ref.Shape.MemberNames() {
			eventRef := ref.Shape.MemberRefs[eventRefName]
			if !eventRef.Shape.IsEvent {
				panic(fmt.Sprintf("unexpected non-event member reference %s.%s",
					ref.Shape.ShapeName, eventRefName))
			}

			updateEventPayloadRef(eventRef.Shape)

			eventRef.Shape.EventFor = append(eventRef.Shape.EventFor, eventStream)
			eventStream.Events = append(eventStream.Events, &Event{
				Name:  eventRefName,
				Shape: eventRef.Shape,
				For:   eventStream,
			})
		}

		// Remove the eventstream references as they will be added elsewhere.
		ref.Shape.removeRef(ref)
		delete(topShape.MemberRefs, refName)
		delete(topShape.API.Shapes, ref.Shape.ShapeName)
	}

	return eventStream
}

func updateEventPayloadRef(parent *Shape) {
	if parent.API.Metadata.Protocol != "rest-xml" {
		return
	}

	refName := parent.PayloadRefName()
	if len(refName) == 0 {
		return
	}

	payloadRef := parent.MemberRefs[refName]

	if payloadRef.Shape.Type == "blob" {
		return
	}

	if len(payloadRef.LocationName) != 0 {
		return
	}

	payloadRef.LocationName = refName
}

func renderEventStreamAPIShape(w io.Writer, s *Shape) error {
	// Imports needed by the EventStream APIs.
	s.API.imports["bytes"] = true
	s.API.imports["io"] = true
	s.API.imports["sync"] = true
	s.API.imports["sync/atomic"] = true
	s.API.imports["github.com/aws/aws-sdk-go/private/protocol/eventstream"] = true
	s.API.imports["github.com/aws/aws-sdk-go/private/protocol/eventstream/eventstreamapi"] = true

	return eventStreamAPIShapeTmpl.Execute(w, s)
}

// Template for an EventStream API Shape that will provide read/writing events
// across the EventStream. This is a special shape that's only public members
// are the Events channel and a Close and Err method.
//
// Executed in the context of a Shape.
var eventStreamAPIShapeTmpl = template.Must(template.New("eventStreamAPIShapeTmpl").Parse(`
{{ if $.EventStreamAPI.Inbound }}
	// {{ $.EventStreamAPI.Inbound.Name }}Event groups together all EventStream
	// events read from the {{ $.EventStreamAPI.Operation.ExportedName }} API.
	//
	// These events are:
	// {{ range $_, $event := $.EventStreamAPI.Inbound.Events }}
	//     * {{ $event.Shape.ShapeName }}
	{{- end }}
	type {{ $.EventStreamAPI.Inbound.Name }}Event interface {
		event{{ $.EventStreamAPI.Name }}()
	}
{{ end }}

{{ $.Documentation }}
type {{ $.ShapeName }} struct {
	{{ if $.EventStreamAPI.Inbound }}
		eventReader *eventstreamapi.EventReader
		inboundStream chan {{ $.EventStreamAPI.Inbound.Name }}Event
		Events <-chan {{ $.EventStreamAPI.Inbound.Name }}Event
	{{ end }}

	{{ if $.EventStreamAPI.Outbound }}
		// TODO Outbound EventStream
	{{ end }}

	done chan struct{}
	closeOnce sync.Once
	errVal atomic.Value
}

func new{{ $.ShapeName }}(
	reader io.ReadCloser,
	unmarshalers request.HandlerList,
	logger aws.Logger,
	logLevel aws.LogLevelType,
) *{{ $.ShapeName }} {
	es := {{ $.ShapeName }}{
		done: make(chan struct{}),
	}

	{{ if $.EventStreamAPI.Inbound }}
		payloadUnmarshaler := protocol.HandlerPayloadUnmarshal{
			Unmarshalers: unmarshalers,
		}
		es.eventReader = eventstreamapi.NewEventReader(
			reader, payloadUnmarshaler, es.unmarshalerForEventType,
		)
		es.eventReader.UseLogger(logger, logLevel)

		es.inboundStream = make(chan {{ $.EventStreamAPI.Inbound.Name }}Event)
		es.Events = es.inboundStream
	{{ end }}

	{{ if $.EventStreamAPI.Outbound }}
		// TODO Outbound EventStream
	{{ end }}

	return &es
}

// Close closes the EventStream. This will also cause the Events channel to be
// closed. You can use the closing of the Events channel to terminate your
// application's read from the API's EventStream.
func (es *{{ $.ShapeName }}) Close() (err error) { 
	es.closeOnce.Do(func() {
		close(es.done)
		err = es.eventReader.Close()
	})

	return err
}

func (es *{{ $.ShapeName }}) Err() error { 
	if v := es.errVal.Load(); v != nil {
		return v.(error)
	}

	return nil
}

{{ if $.EventStreamAPI.Inbound }}
	{{ template "inbound eventstream" $ }}
{{ end }}

{{ define "inbound eventstream" }}
func (es *{{ $.ShapeName }}) readEventStream() {
	defer close(es.inboundStream)
	for {
		event, err := es.eventReader.ReadEvent()
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-es.done:
				// If closed already ignore the error
				return
			default:
			}
			es.errVal.Store(err)
			return
		}

		select {
		case es.inboundStream <- event.({{ $.EventStreamAPI.Inbound.Name }}Event):
		case <-es.done:
			return
		}
	}
}

func (es *{{ $.ShapeName }}) unmarshalerForEventType(
	eventType string,
) (eventstreamapi.Unmarshaler, error) {
	switch eventType {
		{{ range $_, $event := $.EventStreamAPI.Inbound.Events }}
			case {{ printf "%q" $event.Name }}:
				return &{{ $event.Shape.ShapeName }}{}, nil
		{{ end }}
	default:
		return nil, fmt.Errorf(
			"unknown event type name, %s, for {{ $.ShapeName }}", eventType)
	}
}
{{ end }}
`))

// Template for the EventStream API Output shape that contains the EventStream
// member.
//
// Executed in the context of a Shape.
var eventStreamAPILoopMethodTmpl = template.Must(
	template.New("eventStreamAPILoopMethodTmpl").Parse(`
func (s *{{ $.ShapeName }}) runEventStreamLoop(r *request.Request) {
	if r.Error != nil {
		return
	}

	// TODO need to be be generated based on inbound and outbound streams
	{{ $esMemberRef := index $.MemberRefs $.EventStreamsMemberName }}
	s.{{ $.EventStreamsMemberName }} = new{{ $esMemberRef.ShapeName }}(
		r.HTTPResponse.Body, r.Handlers.UnmarshalStream,
		r.Config.Logger, r.Config.LogLevel.Value(),
	)

	go s.{{ $.EventStreamsMemberName }}.readEventStream()
}
`))

// Template for an EventStream Event shape. This is a normal API shape that is
// decorated as an EventStream Event.
//
// Executed in the context of a Shape.
var eventStreamEventShapeTmpl = template.Must(template.New("eventStreamEventShapeTmpl").Parse(`
{{ range $_, $eventstream := $.EventFor }}
	// The {{ $.ShapeName }} is and event in the {{ $eventstream.Name }} group of events.
	func (s *{{ $.ShapeName }}) event{{ $eventstream.Name }}() {}
{{ end }}

// UnmarshalEvent unmarshals the EventStream Message into the {{ $.ShapeName }} value.
// This method is only used internally within the SDK's EventStream handling.
func (s *{{ $.ShapeName }}) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	{{- range $fieldIdx, $fieldName := $.MemberNames }}
		{{- $fieldRef := index $.MemberRefs $fieldName -}}
		{{ if $fieldRef.IsEventHeader }}
			// TODO handle event header, {{ $fieldName }}
		{{- else if (and ($fieldRef.IsEventPayload) (eq $fieldRef.Shape.Type "blob")) }}
			s.{{ $fieldName }} = make([]byte, len(msg.Payload))
			copy(s.{{ $fieldName }}, msg.Payload)
		{{- else }}
			if err := payloadUnmarshaler.UnmarshalPayload(
				bytes.NewReader(msg.Payload), s,
			); err != nil {
				return fmt.Errorf("failed to unmarshal payload, %v", err)
			}
		{{- end }}
	{{- end }}
	return nil
}
`))

var eventStreamTestTmpl = template.Must(template.New("eventStreamTestTmpl").Parse(`
`))
