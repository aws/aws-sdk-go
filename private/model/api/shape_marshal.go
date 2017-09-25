// +build codegen

package api

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// MarshalShapeGoCode renders the shape's MarshalFields method with marshalers
// for each field within the shape. A string is returned of the rendered Go code.
//
// Will panic if error.
func MarshalShapeGoCode(s *Shape) string {
	w := &bytes.Buffer{}
	if err := marshalShapeTmpl.Execute(w, s); err != nil {
		panic(fmt.Sprintf("failed to render shape's fields marshaler, %v", err))
	}

	return w.String()
}

// MarshalShapeRefGoCode renders protocol encode for the shape ref's type.
//
// Will panic if error.
func MarshalShapeRefGoCode(refName string, ref *ShapeRef, context *Shape) string {
	if ref.XMLAttribute {
		return "// Skipping " + refName + " XML Attribute."
	}
	if context.IsRefPayloadReader(refName, ref) {
		if strings.HasSuffix(context.ShapeName, "Output") {
			return "// Skipping " + refName + " Output type's body not valid."
		}
	}

	mRef := marshalShapeRef{
		Name:    refName,
		Ref:     ref,
		Context: context,
	}

	switch mRef.Location() {
	case "StatusCode":
		return "// ignoring invalid encode state, StatusCode. " + refName
	}

	w := &bytes.Buffer{}
	if err := marshalShapeRefTmpl.ExecuteTemplate(w, "encode field", mRef); err != nil {
		panic(fmt.Sprintf("failed to marshal shape ref, %s, %v", ref.Shape.Type, err))
	}

	return w.String()
}

var marshalShapeTmpl = template.Must(template.New("marshalShapeTmpl").Funcs(
	map[string]interface{}{
		"MarshalShapeRefGoCode": MarshalShapeRefGoCode,
	},
).Parse(`
{{ $shapeName := $.ShapeName -}}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s *{{ $shapeName }}) MarshalFields(e protocol.FieldEncoder) error {
	{{ range $name, $ref := $.MemberRefs -}}
		{{ MarshalShapeRefGoCode $name $ref $ }}
	{{ end }}
	return nil
}

{{ if $.UsedInList -}}
func encode{{ $shapeName }}List(vs []*{{ $shapeName }}) func(protocol.ListEncoder) {
	return func(le protocol.ListEncoder) {
		for _, v := range vs {
			le.ListAddFields(v)
		}
	}
}
{{- end }}

{{ if $.UsedInMap -}}
func encode{{ $shapeName }}Map(vs map[string]*{{ $shapeName }}) func(protocol.MapEncoder) {
	return func(me protocol.MapEncoder) {
		for k, v := range vs {
			me.MapSetFields(k, v)
		}
	}
}
{{- end }}
`))

var marshalShapeRefTmpl = template.Must(template.New("marshalShapeRefTmpl").Parse(`
{{ define "encode field" -}}
	{{ if $.IsIdempotencyToken -}}
		{{ template "idempotency token" $ }}
	{
		v := {{ $.Name }}
	{{ else -}}
	if {{ template "is ref set" $ }} {
		v := {{ template "ref value" $ }}
	{{- end }}
		{{ if $.HasAttributes -}}
			{{ template "attributes" $ }}
		{{- end }}
		e.Set{{ $.MarshalerType }}(protocol.{{ $.Location }}Target, "{{ $.LocationName }}", {{ template "marshaler" $ }}, {{ template "metadata" $ }})
	}
{{- end }}

{{ define "marshaler" -}}
	{{- if $.IsShapeType "list" -}}
		{{- $helperName := $.EncodeHelperName "list" -}}
		{{- if $helperName -}}
			{{ $helperName }}
		{{- else -}}
			func(le protocol.ListEncoder) {
				{{ $memberRef := $.ListMemberRef -}}
				for _, item := range v {
					v := {{ if $memberRef.Ref.UseIndirection }}*{{ end }}item
					le.ListAdd{{ $memberRef.MarshalerType }}({{ template "marshaler" $memberRef }})
				}
			}
		{{- end -}}
	{{- else if $.IsShapeType "map" -}}
		{{- $helperName := $.EncodeHelperName "map" -}}
		{{- if $helperName -}}
			{{ $helperName }}
		{{- else -}}
			func(me protocol.MapEncoder) {
				{{ $valueRef := $.MapValueRef -}}
				for k, item := range v {
					v := {{ if $valueRef.Ref.UseIndirection }}*{{ end }}item
					me.MapSet{{ $valueRef.MarshalerType }}(k, {{ template "marshaler" $valueRef }})
				}
			}
		{{- end -}}
	{{- else if $.IsShapeType "structure" -}}
		v
	{{- else if $.IsShapeType "timestamp" -}}
		protocol.TimeValue{V: v, Format: {{ $.TimeFormat }} }
	{{- else if $.IsShapeType "jsonvalue" -}}
		protocol.JSONValue{V: v {{ if eq $.Location "Header" }}, Base64: true{{ end }} }
	{{- else if $.IsPayloadStream -}}
		protocol.{{ $.GoType }}{{ $.MarshalerType }}{V:v}
	{{- else -}}
		protocol.{{ $.GoType }}{{ $.MarshalerType }}(v)
	{{- end -}}
{{- end }}

{{ define "metadata" -}}
	protocol.Metadata{
		{{- if $.IsFlattened -}}
			Flatten: true,
		{{- end -}}

		{{- if $.HasAttributes -}}
			Attributes: attrs,
		{{- end -}}

		{{- if $.XMLNamespacePrefix -}}
			XMLNamespacePrefix: "{{ $.XMLNamespacePrefix }}",
		{{- end -}}

		{{- if $.XMLNamespaceURI -}}
			XMLNamespaceURI: "{{ $.XMLNamespaceURI }}",
		{{- end -}}

		{{- if $.ListLocationName -}}
			ListLocationName: "{{ $.ListLocationName }}",
		{{- end -}}

		{{- if $.MapLocationNameKey -}}
			MapLocationNameKey: "{{ $.MapLocationNameKey }}",
		{{- end -}}

		{{- if $.MapLocationNameValue -}}
			MapLocationNameValue: "{{ $.MapLocationNameValue }}",
		{{- end -}}
	}
{{- end }}

{{ define "ref value" -}}
	{{ if $.Ref.UseIndirection }}*{{ end }}s.{{ $.Name }}
{{- end}}

{{ define "is ref set" -}}
	{{ $isList := $.IsShapeType "list" -}}
	{{ $isMap := $.IsShapeType "map" -}}
	{{- if or $isList $isMap -}}
		len(s.{{ $.Name }}) > 0
	{{- else -}}
		s.{{ $.Name }} != nil
	{{- end -}}
{{- end }}

{{ define "attributes" -}}
	attrs := make([]protocol.Attribute, 0, {{ $.NumAttributes }})

	{{ range $name, $child := $.ChildrenRefs -}}
		{{ if $child.Ref.XMLAttribute -}}
			if s.{{ $.Name }}.{{ $name }} != nil {
				v := {{ if $child.Ref.UseIndirection }}*{{ end }}s.{{ $.Name }}.{{ $name }}
				attrs = append(attrs, protocol.Attribute{Name: "{{ $child.LocationName }}", Value: {{ template "marshaler" $child }}, Meta: {{ template "metadata" $child }}})
			}
		{{- end }}
	{{- end }}
{{- end }}

{{ define "idempotency token" -}}
    var {{ $.Name }} string
	if {{ template "is ref set" $ }} {
		{{ $.Name }} = {{ template "ref value" $ }}
	} else {
		{{ $.Name }} = protocol.GetIdempotencyToken()
	}
{{- end }}
`))

type marshalShapeRef struct {
	Name    string
	Ref     *ShapeRef
	Context *Shape
}

func (r marshalShapeRef) ListMemberRef() marshalShapeRef {
	return marshalShapeRef{
		Name:    r.Name + "ListMember",
		Ref:     &r.Ref.Shape.MemberRef,
		Context: r.Ref.Shape,
	}
}
func (r marshalShapeRef) MapValueRef() marshalShapeRef {
	return marshalShapeRef{
		Name:    r.Name + "MapValue",
		Ref:     &r.Ref.Shape.ValueRef,
		Context: r.Ref.Shape,
	}
}
func (r marshalShapeRef) ChildrenRefs() map[string]marshalShapeRef {
	children := map[string]marshalShapeRef{}

	for name, ref := range r.Ref.Shape.MemberRefs {
		children[name] = marshalShapeRef{
			Name:    name,
			Ref:     ref,
			Context: r.Ref.Shape,
		}
	}

	return children
}
func (r marshalShapeRef) IsShapeType(typ string) bool {
	return r.Ref.Shape.Type == typ
}
func (r marshalShapeRef) IsPayloadStream() bool {
	return r.Context.IsRefPayloadReader(r.Name, r.Ref)
}
func (r marshalShapeRef) MarshalerType() string {
	switch r.Ref.Shape.Type {
	case "list":
		return "List"
	case "map":
		return "Map"
	case "structure":
		return "Fields"
	default:
		// Streams have a special case
		if r.Context.IsRefPayload(r.Name) {
			return "Stream"
		}
		return "Value"
	}
}
func (r marshalShapeRef) EncodeHelperName(typ string) string {
	if r.Ref.Shape.Type != typ {
		return ""
	}

	var memberRef marshalShapeRef
	switch r.Ref.Shape.Type {
	case "map":
		memberRef = r.MapValueRef()
	case "list":
		memberRef = r.ListMemberRef()
	default:
		return ""
	}

	switch memberRef.Ref.Shape.Type {
	case "list", "map":
		return ""
	case "structure":
		shapeName := memberRef.Ref.Shape.ShapeName
		return "encode" + shapeName + strings.Title(typ) + "(v)"
	default:
		return "protocol.Encode" + memberRef.GoType() + strings.Title(typ) + "(v)"
	}
}
func (r marshalShapeRef) GoType() string {
	switch r.Ref.Shape.Type {
	case "boolean":
		return "Bool"
	case "string", "character":
		return "String"
	case "integer", "long":
		return "Int64"
	case "float", "double":
		return "Float64"
	case "timestamp":
		return "Time"
	case "jsonvalue":
		return "JSONValue"
	case "blob":
		if r.Context.IsRefPayloadReader(r.Name, r.Ref) {
			if strings.HasSuffix(r.Context.ShapeName, "Output") {
				return "ReadCloser"
			}
			return "ReadSeeker"
		}
		return "Bytes"
	default:
		panic(fmt.Sprintf("unknown marshal shape ref type, %s", r.Ref.Shape.Type))
	}
}
func (r marshalShapeRef) Location() string {
	var loc string
	if l := r.Ref.Location; len(l) > 0 {
		loc = l
	} else if l := r.Ref.Shape.Location; len(l) > 0 {
		loc = l
	}

	switch loc {
	case "querystring":
		return "Query"
	case "header":
		return "Header"
	case "headers": // headers means key is header prefix
		return "Headers"
	case "uri":
		return "Path"
	case "statusCode":
		return "StatusCode"
	default:
		if len(loc) != 0 {
			panic(fmt.Sprintf("unknown marshal shape ref location, %s", loc))
		}

		if r.Context.IsRefPayload(r.Name) {
			return "Payload"
		}

		return "Body"
	}
}
func (r marshalShapeRef) LocationName() string {
	if l := r.Ref.QueryName; len(l) > 0 {
		// Special case for EC2 query
		return l
	}

	locName := r.Name
	if l := r.Ref.LocationName; len(l) > 0 {
		locName = l
	} else if l := r.Ref.Shape.LocationName; len(l) > 0 {
		locName = l
	}

	return locName
}
func (r marshalShapeRef) IsFlattened() bool {
	return r.Ref.Flattened || r.Ref.Shape.Flattened
}
func (r marshalShapeRef) XMLNamespacePrefix() string {
	if v := r.Ref.XMLNamespace.Prefix; len(v) != 0 {
		return v
	}
	return r.Ref.Shape.XMLNamespace.Prefix
}
func (r marshalShapeRef) XMLNamespaceURI() string {
	if v := r.Ref.XMLNamespace.URI; len(v) != 0 {
		return v
	}
	return r.Ref.Shape.XMLNamespace.URI
}
func (r marshalShapeRef) ListLocationName() string {
	if v := r.Ref.Shape.MemberRef.LocationName; len(v) > 0 {
		if !(r.Ref.Shape.Flattened || r.Ref.Flattened) {
			return v
		}
	}
	return ""
}
func (r marshalShapeRef) MapLocationNameKey() string {
	return r.Ref.Shape.KeyRef.LocationName
}
func (r marshalShapeRef) MapLocationNameValue() string {
	return r.Ref.Shape.ValueRef.LocationName
}
func (r marshalShapeRef) HasAttributes() bool {
	for _, ref := range r.Ref.Shape.MemberRefs {
		if ref.XMLAttribute {
			return true
		}
	}
	return false
}
func (r marshalShapeRef) NumAttributes() (n int) {
	for _, ref := range r.Ref.Shape.MemberRefs {
		if ref.XMLAttribute {
			n++
		}
	}
	return n
}
func (r marshalShapeRef) IsIdempotencyToken() bool {
	return r.Ref.IdempotencyToken || r.Ref.Shape.IdempotencyToken
}
func (r marshalShapeRef) TimeFormat() string {
	switch r.Location() {
	case "Header", "Headers":
		return "protocol.RFC822TimeFromat"
	default:
		switch r.Context.API.Metadata.Protocol {
		case "json", "rest-json":
			return "protocol.UnixTimeFormat"
		case "rest-xml", "ec2", "query":
			return "protocol.ISO8601TimeFormat"
		default:
			panic(fmt.Sprintf("unable to determine time format for %s ref", r.Name))
		}
	}
}

// UnmarshalShapeGoCode renders the shape's UnmarshalAWS method with unmarshalers
// for each field within the shape. A string is returned of the rendered Go code.
//
// Will panic if error.
func UnmarshalShapeGoCode(s *Shape) string {
	w := &bytes.Buffer{}
	if err := unmarshalShapeTmpl.Execute(w, s); err != nil {
		panic(fmt.Sprintf("failed to render shape's fields unmarshaler, %v", err))
	}

	return w.String()
}

var unmarshalShapeTmpl = template.Must(template.New("unmarshalShapeTmpl").Funcs(
	template.FuncMap{
		"UnmarshalShapeRefGoCode": UnmarshalShapeRefGoCode,
	},
).Parse(`
{{ $shapeName := $.ShapeName -}}

// UnmarshalAWS decodes the AWS API shape using the passed in protocol decoder.
func (s *{{ $shapeName }}) UnmarshalAWS(d protocol.FieldDecoder) {
	{{ range $name, $ref := $.MemberRefs -}}
		{{ UnmarshalShapeRefGoCode $name $ref $ }}
	{{ end }}
}

{{ if $.UsedInList -}}
func decode{{ $shapeName }}List(vsp *[]*{{ $shapeName }}) func(int, protocol.ListDecoder) {
	return func(n int, ld protocol.ListDecoder) {
		vs := make([]{{ $shapeName }}, n)
		*vsp = make([]*{{ $shapeName }}, n)
		for i := 0; i < n; i++ {
			ld.ListGetUnmarshaler(&vs[i])
			(*vsp)[i] = &vs[i]
		}
	}
}
{{- end }}

{{ if $.UsedInMap -}}
func decode{{ $shapeName }}Map(vsp *map[string]*{{ $shapeName }}) func([]string, protocol.MapDecoder) {
	return func(ks []string, md protocol.MapDecoder) {
		vs := make(map[string]*{{ $shapeName }}, n)
		for _, k range ks {
			v := &{{ $shapeName }}{}
			md.MapGetUnmarshaler(k, v)
			vs[k] = v
		}
	}
}
{{- end }}
`))

// UnmarshalShapeRefGoCode generates the Go code to unmarshal an API shape.
func UnmarshalShapeRefGoCode(refName string, ref *ShapeRef, context *Shape) string {
	if ref.XMLAttribute {
		return "// Skipping " + refName + " XML Attribute."
	}

	mRef := marshalShapeRef{
		Name:    refName,
		Ref:     ref,
		Context: context,
	}

	switch mRef.Location() {
	case "Path":
		return "// ignoring invalid decode state, Path. " + refName
	case "Query":
		return "// ignoring invalid decode state, Query. " + refName
	}

	w := &bytes.Buffer{}
	if err := unmarshalShapeRefTmpl.ExecuteTemplate(w, "decode", mRef); err != nil {
		panic(fmt.Sprintf("failed to decode shape ref, %s, %v", ref.Shape.Type, err))
	}

	return w.String()
}

var unmarshalShapeRefTmpl = template.Must(template.New("unmarshalShapeRefTmpl").Parse(`
//  Decode {{ $.Name }} {{ $.GoType }} {{ $.MarshalerType }} to {{ $.Location }} at {{ $.LocationName }}
`))
