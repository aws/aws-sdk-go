// +build codegen

package api

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

func renderEventStreamAPI(w io.Writer, op *Operation) error {
	// Imports needed by the EventStream APIs.
	op.API.AddImport("fmt")
	op.API.AddImport("bytes")
	op.API.AddImport("io")
	op.API.AddImport("sync")
	op.API.AddImport("sync/atomic")
	op.API.AddSDKImport("aws")
	op.API.AddSDKImport("aws/awserr")
	op.API.AddSDKImport("aws/request")
	op.API.AddSDKImport("private/protocol/eventstream")
	op.API.AddSDKImport("private/protocol/eventstream/eventstreamapi")

	return eventStreamAPITmpl.Execute(w, op)
}

// Template for an EventStream API Shape that will provide read/writing events
// across the EventStream. This is a special shape that's only public members
// are the Events channel and a Close and Err method.
//
// Executed in the context of a Shape.
var eventStreamAPITmpl = template.Must(
	template.New("eventStreamAPITmplDef").
		Funcs(template.FuncMap{
			"unexported": func(v string) string {
				return strings.ToLower(string(v[0])) + v[1:]
			},
		}).
		Parse(eventStreamAPITmplDef),
)

const eventStreamAPITmplDef = `
{{- $esName := $.EventStreamAPI.Name }}
{{- $writerName := printf "%sWriter" $.EventStreamAPI.Name }}
{{- $readerName := printf "%sReader" $.EventStreamAPI.Name }}
{{- $outputStream := $.EventStreamAPI.OutputStream }}
{{- $inputStream := $.EventStreamAPI.InputStream }}

// {{ $esName }} provides the event stream handling for the {{ $.ExportedName }}.
type {{ $esName }} struct {
	{{- if $inputStream }}

		// Writer is the EventStream writer for the {{ $inputStream.Name }}
		// events. This value is automatically set by the SDK when the API call is made
		// Use this member when unit testing your code with the SDK to mock out the
		// EventStream Writer.
		//
		// Must not be nil.
		Writer {{ $inputStream.Name }}Writer

		inputWriter io.WriteCloser

		{{- if eq .API.Metadata.Protocol "json" }}

			input {{ $.InputRef.GoType }}
		{{- end }}
	{{- end }}

	{{- if $outputStream }}

		// Reader is the EventStream reader for the {{ $outputStream.Name }}
		// events. This value is automatically set by the SDK when the API call is made
		// Use this member when unit testing your code with the SDK to mock out the
		// EventStream Reader.
		//
		// Must not be nil.
		Reader {{ $outputStream.Name }}Reader

		{{- if eq .API.Metadata.Protocol "json" }}

			output {{ $.OutputRef.GoType }}
		{{- end }}
	{{- end }}

	// StreamCloser is the io.Closer for the EventStream connection. For HTTP
	// EventStream this is the response Body. The stream will be closed when
	// the Close method of the EventStream is called.
	StreamCloser io.Closer
}

func (es *{{ $esName }}) setStreamCloser(r *request.Request) {
	var closers aws.MultiCloser
	if es.StreamCloser != nil {
		closers = append(closers, es.StreamCloser)
	}

	{{- if $inputStream }}
		closers = append(closers, es.inputWriter)
	{{- end }}

	{{- if $outputStream }}
		closers = append(closers, r.HTTPResponse.Body)
	{{- end }}

	es.StreamCloser =  closers
}

{{- if $inputStream }}

	// Send writes the event to the stream blocking until the event is written.
	// Returns an error if the event was not written.
	//
	// These events are:
	// {{ range $_, $event := $inputStream.Events }}
	//     * {{ $event.Shape.ShapeName }}
	{{- end }}
	func (es *{{ $esName }}) Send(event {{ $inputStream.Name }}Event) {
		es.Writer.Send(event)
	}

	func (es *{{ $esName }}) runInputStream(r *request.Request) {
		var signer *eventstreamapi.MessageSigner
		{{- if and false $.ShouldSignRequestBody }}
			{{- $.API.AddSDKImport "aws/signer/v4" }}
			sigSeed, err := v4.GetSignedRequestSignature(r.HTTPRequest)
			if err != nil {
				r.Error = awserr.New("", "unable to get initial request's signature", err)
				return
			}
			signer = &eventstreamapi.MessageSigner{
					Signer: v4.NewStreamSigner(
						r.ClientInfo.SigningRegion, r.ClientInfo.SigningName,
						sigSeed, r.Config.Credentials),
			}
		{{- end }}

		writer := newWrite{{ $inputStream.Name }}(
			es.inputWriter, 
			r.Handlers.BuildStream,
			signer,
			r.Config.Logger,
			r.Config.LogLevel.Value(),
		)
		es.Writer = writer
		go writer.writeEventStream()
	}

	{{- if eq .API.Metadata.Protocol "json" }}

		func (es *{{ $esName }}) sendInitialRequestEvent(r *request.Request) {
			if err := es.Send(es.input); err != nil {
				r.Error = err
			}
		}
	{{- end }}
{{- end }}

{{- if $outputStream }}

	// Events returns a channel to read events from.
	//
	// These events are:
	// {{ range $_, $event := $outputStream.Events }}
	//     * {{ $event.Shape.ShapeName }}
	{{- end }}
	func (es *{{ $esName }}) Events() <-chan {{ $outputStream.Name }}Event {
		return es.Reader.Events()
	}

	func (es *{{ $esName }}) runOutputStream(r *request.Request) {
		reader := newRead{{ $outputStream.Name }}(
			r.HTTPResponse.Body,
			r.Handlers.UnmarshalStream,
			r.Config.Logger,
			r.Config.LogLevel.Value(),
		)
		es.Reader = reader
		go reader.readEventStream()
	}

	{{- if eq .API.Metadata.Protocol "json" }}

		func (es *{{ $esName }}) recvInitialResponseEvent(r *request.Request) {
			// Wait for the initial response event, which must be the first
			// event to be received from the API.
			select {
			case event, ok := <- es.Events():
				if !ok {
					return
				}

				v, ok := event.({{ $.OutputRef.GoType }})
				if !ok || v == nil {
					r.Error = awserr.New(
						request.ErrCodeSerialization,
						fmt.Sprintf("invalid event, %T, expect %T, %v",
							event, ({{ $.OutputRef.GoType }})(nil), v),
						nil,
					)
					return
				}

				*es.output = *v
				es.output.{{ $.EventStreamAPI.OutputMemberName  }} = es
			}
		}
	{{- end }}
{{- end }}

// Close closes the EventStream. This will also cause the Events channel to be
// closed. You can use the closing of the Events channel to terminate your
// application's read from the API's EventStream.
{{- if $inputStream }}
//
// Will close the underlying EventStream writer. 
{{- end }}
{{- if $outputStream }}
//
// Will close the underlying EventStream reader.
{{- end }}
//
// For EventStream over HTTP connection this will also close the HTTP connection.
//
// Close must be called when done using the EventStream API. Not calling Close
// may result in resource leaks.
func (es *{{ $esName }}) Close() (err error) {
	{{- if $outputStream }}
		es.Reader.Close()
	{{- end }}
	{{- if $inputStream }}
		es.Writer.Close()
	{{- end }}
	es.StreamCloser.Close()

	return es.Err()
}

// Err returns any error that occurred while reading or writing EventStream
// Events from the service API's response. Returns nil if there were no errors.
func (es *{{ $esName }}) Err() error {
	{{- if $inputStream }}
		if err := es.Writer.Err(); err != nil {
			return err
		}
	{{- end }}
	{{- if $outputStream }}
		if err := es.Reader.Err(); err != nil {
			return err
		}
	{{- end }}

	return nil
}
`

func renderEventStreamShape(w io.Writer, s *Shape) error {
	// Imports needed by the EventStream APIs.
	s.API.AddImport("fmt")
	s.API.AddImport("bytes")
	s.API.AddImport("io")
	s.API.AddImport("sync")
	s.API.AddImport("sync/atomic")
	s.API.AddSDKImport("aws")
	s.API.AddSDKImport("aws/awserr")
	s.API.AddSDKImport("private/protocol/eventstream")
	s.API.AddSDKImport("private/protocol/eventstream/eventstreamapi")

	return eventStreamShapeTmpl.Execute(w, s)
}

var eventStreamShapeTmpl = func() *template.Template {
	t := template.Must(
		template.New("eventStreamShapeTmplDef").
			Parse(eventStreamShapeTmplDef),
	)
	template.Must(
		t.AddParseTree(
			"eventStreamAPIWriterTmpl", eventStreamAPIWriterTmpl.Tree),
	)
	template.Must(
		t.AddParseTree(
			"eventStreamAPIReaderTmpl", eventStreamAPIReaderTmpl.Tree),
	)

	return t
}()

const eventStreamShapeTmplDef = `
{{- if $.IsInputEventStream }}
	{{- template "eventStreamAPIWriterTmpl" $ }}
{{- end }}

{{- if $.IsOutputEventStream }}
	{{- template "eventStreamAPIReaderTmpl" $ }}
{{- end }}
`

// EventStreamHeaderTypeMap provides the mapping of a EventStream Header's
// Value type to the shape reference's member type.
type EventStreamHeaderTypeMap struct {
	Header string
	Member string
}

// Returns if the event has any members which are not the event's blob payload,
// nor a header.
func eventHasNonBlobPayloadMembers(s *Shape) bool {
	num := len(s.MemberRefs)
	for _, ref := range s.MemberRefs {
		if ref.IsEventHeader || (ref.IsEventPayload && (ref.Shape.Type == "blob" || ref.Shape.Type == "string")) {
			num--
		}
	}
	return num > 0
}

func setEventHeaderValueForType(s *Shape, memVar string) string {
	switch s.Type {
	case "blob":
		return fmt.Sprintf("eventstream.BytesValue(%s)", memVar)
	case "string":
		return fmt.Sprintf("eventstream.StringValue(*%s)", memVar)
	case "boolean":
		return fmt.Sprintf("eventstream.BoolValue(*%s)", memVar)
	case "byte":
		return fmt.Sprintf("eventstream.Int8Value(int8(*%s))", memVar)
	case "short":
		return fmt.Sprintf("eventstream.Int16Value(int16(*%s))", memVar)
	case "integer":
		return fmt.Sprintf("eventstream.Int32Value(int32(*%s))", memVar)
	case "long":
		return fmt.Sprintf("eventstream.Int64Value(*%s)", memVar)
	case "float":
		return fmt.Sprintf("eventstream.Float32Value(float32(*%s))", memVar)
	case "double":
		return fmt.Sprintf("eventstream.Float64Value(*%s)", memVar)
	case "timestamp":
		return fmt.Sprintf("eventstream.TimestampValue(*%s)", memVar)
	default:
		panic(fmt.Sprintf("value type %s not supported for event headers, %s", s.Type, s.ShapeName))
	}
}

func shapeMessageType(s *Shape) string {
	if s.Exception {
		return "eventstreamapi.ExceptionMessageType"
	}
	return "eventstreamapi.EventMessageType"
}

var eventStreamEventShapeTmplFuncs = template.FuncMap{
	"EventStreamHeaderTypeMap": func(ref *ShapeRef) EventStreamHeaderTypeMap {
		switch ref.Shape.Type {
		case "boolean":
			return EventStreamHeaderTypeMap{Header: "bool", Member: "bool"}
		case "byte":
			return EventStreamHeaderTypeMap{Header: "int8", Member: "int64"}
		case "short":
			return EventStreamHeaderTypeMap{Header: "int16", Member: "int64"}
		case "integer":
			return EventStreamHeaderTypeMap{Header: "int32", Member: "int64"}
		case "long":
			return EventStreamHeaderTypeMap{Header: "int64", Member: "int64"}
		case "timestamp":
			return EventStreamHeaderTypeMap{Header: "time.Time", Member: "time.Time"}
		case "blob":
			return EventStreamHeaderTypeMap{Header: "[]byte", Member: "[]byte"}
		case "string":
			return EventStreamHeaderTypeMap{Header: "string", Member: "string"}
		case "uuid":
			return EventStreamHeaderTypeMap{Header: "[]byte", Member: "[]byte"}
		default:
			panic("unsupported EventStream header type, " + ref.Shape.Type)
		}
	},
	"EventHeaderValueForType":  setEventHeaderValueForType,
	"ShapeMessageType":         shapeMessageType,
	"HasNonBlobPayloadMembers": eventHasNonBlobPayloadMembers,
}

// Template for an EventStream Event shape. This is a normal API shape that is
// decorated as an EventStream Event.
//
// Executed in the context of a Shape.
var eventStreamEventShapeTmpl = template.Must(template.New("eventStreamEventShapeTmpl").
	Funcs(eventStreamEventShapeTmplFuncs).Parse(`
{{ range $_, $eventStream := $.EventFor }}
	// The {{ $.ShapeName }} is and event in the {{ $eventStream.Name }} group of events.
	func (s *{{ $.ShapeName }}) event{{ $eventStream.Name }}() {}
{{ end }}

// UnmarshalEvent unmarshals the EventStream Message into the {{ $.ShapeName }} value.
// This method is only used internally within the SDK's EventStream handling.
func (s *{{ $.ShapeName }}) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	{{- range $memName, $memRef := $.MemberRefs }}
		{{- if $memRef.IsEventHeader }}
			if hv := msg.Headers.Get("{{ $memName }}"); hv != nil {
				{{ $types := EventStreamHeaderTypeMap $memRef -}}
				v := hv.Get().({{ $types.Header }})
				{{- if ne $types.Header $types.Member }}
					m := {{ $types.Member }}(v)
					s.{{ $memName }} = {{ if $memRef.UseIndirection }}&{{ end }}m
				{{- else }}
					s.{{ $memName }} = {{ if $memRef.UseIndirection }}&{{ end }}v
				{{- end }}
			}
		{{- else if (and ($memRef.IsEventPayload) (eq $memRef.Shape.Type "blob")) }}
			s.{{ $memName }} = make([]byte, len(msg.Payload))
			copy(s.{{ $memName }}, msg.Payload)
		{{- else if (and ($memRef.IsEventPayload) (eq $memRef.Shape.Type "string")) }}
			s.{{ $memName }} = aws.String(string(msg.Payload))
		{{- end }}
	{{- end }}
	{{- if HasNonBlobPayloadMembers $ }}
		if err := payloadUnmarshaler.UnmarshalPayload(
			bytes.NewReader(msg.Payload), s,
		); err != nil {
			return err
		}
	{{- end }}
	return nil
}

func (s *{{ $.ShapeName}}) MarshalEvent(pm protocol.PayloadMarshaler) (msg eventstream.Message, err error) {
	{{- $messageType := ShapeMessageType $ }}
	msg.Headers.Set(eventstreamapi.MessageTypeHeader,
		eventstream.StringValue({{ $messageType }}))
	msg.Headers.Set(eventstreamapi.EventTypeHeader,
		eventstream.StringValue("{{ $.OrigShapeName }}"))
	{{- range $memName, $memRef := $.MemberRefs }}
		{{- if $memRef.IsEventHeader }}
			{{ $memVar := printf "s.%s" $memName -}}
			{{ $typedMem := EventHeaderValueForType $memRef $memVar -}}
			msg.Header.Set("{{ $memName }}", {{ $typedMem }})
		{{- else if (and ($memRef.IsEventPayload) (eq $memRef.Shape.Type "blob")) }}
			msg.Payload = s.{{ $memName }}
		{{- else if (and ($memRef.IsEventPayload) (eq $memRef.Shape.Type "string")) }}
			msg.Payload = []byte(s.{{ $memName }})
		{{- end }}
	{{- end }}
	{{- if HasNonBlobPayloadMembers $ }}
		var buf bytes.Buffer
		if err = pm.MarshalPayload(&buf, s); err != nil {
			return eventstream.Message{}, err
		}
		msg.Payload = buf.Bytes()
	{{- end }}
	return msg, err
}
`))

var eventStreamExceptionEventShapeTmpl = template.Must(
	template.New("eventStreamExceptionEventShapeTmpl").Parse(`
// Code returns the exception type name.
func (s {{ $.ShapeName }}) Code() string {
	{{- if $.ErrorInfo.Code }}
		return "{{ $.ErrorInfo.Code }}"
	{{- else }}
		return "{{ $.ShapeName }}"
	{{ end -}}
}

// Message returns the exception's message.
func (s {{ $.ShapeName }}) Message() string {
	{{- if index $.MemberRefs "Message_" }}
		return *s.Message_
	{{- else }}
		return ""
	{{ end -}}
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s {{ $.ShapeName }}) OrigErr() error {
	return nil
}

func (s {{ $.ShapeName }}) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}
`))
