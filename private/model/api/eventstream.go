// +build codegen

package api

import (
	"bytes"
	"fmt"
	"text/template"
)

// EventStreamAPI provides details about the event stream async API and
// associated EventStream shapes.
type EventStreamAPI struct {
	API       *API
	Operation *Operation
	//	Shape     *Shape
	Name         string
	InputStream  *EventStream
	OutputStream *EventStream

	// The eventstream generated code was generated with an older model that
	// does not scale with bi-directional models. This drives the need to
	// expose the output shape's event stream member as an exported member.
	Legacy bool
}

// EventStream represents a single eventstream group (input/output) and the
// modeled events that are known for the stream.
type EventStream struct {
	Name       string
	Shape      *Shape
	Events     []*Event
	Exceptions []*Event
}

// Event is a single EventStream event that can be sent or received in an
// EventStream.
type Event struct {
	Name    string
	Shape   *Shape
	For     *EventStream
	Private bool
}

// ShapeDoc returns the docstring for the EventStream API.
func (esAPI *EventStreamAPI) ShapeDoc() string {
	tmpl := template.Must(template.New("eventStreamShapeDoc").Parse(`
{{- $.Name }} provides handling of EventStreams for
the {{ $.Operation.ExportedName }} API.

{{- if $.OutputStream }}

Use this type to receive {{ $.OutputStream.Name }} events. The events
can be read from the stream.

The events that can be received are:
{{- range $_, $event := $.OutputStream.Events }}
    * {{ $event.Shape.ShapeName }}
{{- end }}

{{- end }}

{{- if $.InputStream }}

Use this type to send {{ $.InputStream.Name }} events. The events
can be written to the stream.

The events that can be sent are:
{{ range $_, $event := $.InputStream.Events -}}
    * {{ $event.Shape.ShapeName }}
{{- end }}

{{- end }}`))

	var w bytes.Buffer
	if err := tmpl.Execute(&w, esAPI); err != nil {
		panic(fmt.Sprintf("failed to generate eventstream shape template for %v, %v",
			esAPI.Operation.ExportedName, err))
	}

	return commentify(w.String())
}

func hasEventStream(topShape *Shape) bool {
	for _, ref := range topShape.MemberRefs {
		if ref.Shape.IsEventStream {
			return true
		}
	}

	return false
}

func eventStreamAPIShapeRefDoc(refName string) string {
	return commentify(fmt.Sprintf("Use %s to use the API's stream.", refName))
}

func (a *API) setupEventStreams() error {
	for opName, op := range a.Operations {
		_, inRef := getEventStream(op.InputRef.Shape)
		_, outRef := getEventStream(op.OutputRef.Shape)

		if inRef == nil && outRef == nil {
			continue
		}
		if inRef != nil && outRef == nil {
			return fmt.Errorf("event stream input only stream not supported for protocol %s, %s, %v",
				a.NiceName(), opName, a.Metadata.Protocol)
		}
		switch a.Metadata.Protocol {
		case `rest-json`, `rest-xml`, `json`:
		default:
			return UnsupportedAPIModelError{
				Err: fmt.Errorf("EventStream not supported for protocol %s, %s, %v",
					a.NiceName(), opName, a.Metadata.Protocol),
			}
		}

		// TODO inputStream and outputStream are generated per Operation, not
		// per instance of that event type. This will cause conflicts if an
		// eventstream decorated shape is stream between multiple operations.
		//
		// This needs to be split into two passes:
		// 1.) Gather EventStream with their associated set of events.
		// 2.) Gather operations that use EventStreams for an EventStreamAPI.

		var inputStream *EventStream
		if inRef != nil {
			inputStream = setupEventStream(inRef)
			inputStream.Shape.IsInputEventStream = true

			if op.API.Metadata.Protocol == "json" {
				op.InputRef.Shape.EventFor = append(op.InputRef.Shape.EventFor, inputStream)
			}
		}

		var outputStream *EventStream
		if outRef != nil {
			outputStream = setupEventStream(outRef)
			outputStream.Shape.IsOutputEventStream = true

			if op.API.Metadata.Protocol == "json" {
				op.OutputRef.Shape.EventFor = append(op.OutputRef.Shape.EventFor, outputStream)
			}
		}

		a.HasEventStream = true
		op.EventStreamAPI = &EventStreamAPI{
			API:          a,
			Operation:    op,
			Name:         op.ExportedName + "EventStream",
			InputStream:  inputStream,
			OutputStream: outputStream,
			Legacy:       isLegacyEventStream(op),
		}
		op.OutputRef.Shape.OutputEventStreamAPI = op.EventStreamAPI

		if s, ok := a.Shapes[op.EventStreamAPI.Name]; ok {
			newName := op.EventStreamAPI.Name + "Data"
			if _, ok := a.Shapes[newName]; ok {
				panic(fmt.Sprintf(
					"%s: attempting to rename %s to %s, but shape with that name already exists",
					a.NiceName(), op.EventStreamAPI.Name, newName))
			}
			s.Rename(newName)
		}
	}

	return nil
}

var legacyEventStream = map[string]map[string]struct{}{
	"s3": {
		"SelectObjectContent": struct{}{},
	},
	"kinesis": {
		"SubscribeToShard": struct{}{},
	},
}

func isLegacyEventStream(op *Operation) bool {
	if s, ok := legacyEventStream[op.API.PackageName()]; ok {
		if _, ok = s[op.ExportedName]; ok {
			return true
		}
	}
	return false
}

func (e EventStreamAPI) OutputMemberName() string {
	if e.Legacy {
		return "EventStream"
	}

	return "eventStream"
}

func getEventStream(topShape *Shape) (string, *ShapeRef) {
	for refName, ref := range topShape.MemberRefs {
		if !ref.Shape.IsEventStream {
			continue
		}
		return refName, ref
	}

	return "", nil
}

func setupEventStream(ref *ShapeRef) *EventStream {
	//	// Swap out the modeled shape with a copy so that references to the
	//	// events are not lost, and the modeled shape can be dropped if
	//	// unneeded.
	//	ref.Shape.removeRef(ref)
	//	clonedShape := ref.Shape.Clone(ref.Shape.ShapeName + "EventStream")
	//	ref.Shape = clonedShape
	//	clonedShape.refs = append(clonedShape.refs, ref)

	eventStream := &EventStream{
		Name:  ref.Shape.ShapeName,
		Shape: ref.Shape,
	}
	ref.Shape.EventStream = eventStream

	for _, eventRefName := range ref.Shape.MemberNames() {
		eventRef := ref.Shape.MemberRefs[eventRefName]
		if !(eventRef.Shape.IsEvent || eventRef.Shape.Exception) {
			panic(fmt.Sprintf("unexpected non-event member reference %s.%s",
				ref.Shape.ShapeName, eventRefName))
		}

		updateEventPayloadRef(eventRef.Shape)

		eventRef.Shape.EventFor = append(eventRef.Shape.EventFor, eventStream)

		// Exceptions and events are two different lists to allow the SDK
		// to easily generate code with the two handled differently.
		event := &Event{
			Name:  eventRefName,
			Shape: eventRef.Shape,
			For:   eventStream,
		}
		if eventRef.Shape.Exception {
			eventStream.Exceptions = append(eventStream.Exceptions, event)
		} else {
			eventStream.Events = append(eventStream.Events, event)
		}
	}

	return eventStream
}

func updateEventPayloadRef(parent *Shape) {
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
