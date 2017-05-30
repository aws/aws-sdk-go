// +build codegen

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go/private/util"
)

type Examples map[string][]Example

// ExamplesDefinition is the structural representation of the examples-1.json file
type ExamplesDefinition struct {
	*API     `json:"-"`
	Examples Examples `json:"examples"`
}

// Example is a single entry within the examples-1.json file.
type Example struct {
	API           *API                   `json:"-"`
	Operation     *Operation             `json:"-"`
	OperationName string                 `json:"-"`
	Index         string                 `json:"-"`
	VisitedErrors map[string]struct{}    `json:"-"`
	Title         string                 `json:"title"`
	Description   string                 `json:"description"`
	ID            string                 `json:"id"`
	Comments      Comments               `json:"comments"`
	Input         map[string]interface{} `json:"input"`
	Output        map[string]interface{} `json:"output"`
}

type Comments struct {
	Input  map[string]interface{} `json:"input"`
	Output map[string]interface{} `json:"output"`
}

var exampleFuncMap = template.FuncMap{
	"commentify":           commentify,
	"wrap":                 wrap,
	"generateExampleInput": generateExampleInput,
	"generateTypes":        generateTypes,
}

var exampleCustomizations = map[string]template.FuncMap{}

var exampleTmpls = template.Must(template.New("example").Funcs(exampleFuncMap).Parse(`
{{ generateTypes . }}
{{ commentify (wrap .Title 80 false) }}
//
{{ commentify (wrap .Description 80 false) }}
func Example{{ .API.StructName }}_{{ .MethodName }}() {
	svc := {{ .API.PackageName }}.New(session.New())
	input := &{{ .API.PackageName }}.{{ .Operation.InputRef.Shape.ShapeName }} {
		{{ generateExampleInput . -}}
	}

	result, err := svc.{{ .OperationName }}(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
				{{ range $_, $ref := .Operation.ErrorRefs -}}
					{{ if not ($.HasVisitedError $ref) -}}
			case {{ .API.PackageName }}.{{ $ref.Shape.ErrorCodeName }}:
				fmt.Println({{ .API.PackageName }}.{{ $ref.Shape.ErrorCodeName }}, aerr.Error())
					{{ end -}}
				{{ end -}}
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
`))

// Names will return the name of the example. This will also be the name of the operation
// that is to be tested.
func (exs Examples) Names() []string {
	names := make([]string, 0, len(exs))
	for k := range exs {
		names = append(names, k)
	}

	sort.Strings(names)
	return names
}

func (exs Examples) GoCode() string {
	buf := bytes.NewBuffer(nil)
	for _, opName := range exs.Names() {
		examples := exs[opName]
		for _, ex := range examples {
			buf.WriteString(util.GoFmt(ex.GoCode()))
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

// ExampleCode will generate the example code for the given Example shape.
func (ex Example) GoCode() string {
	var buf bytes.Buffer
	m := exampleFuncMap
	if fMap, ok := exampleCustomizations[ex.API.PackageName()]; ok {
		m = fMap
	}
	tmpl := exampleTmpls.Funcs(m)
	if err := tmpl.ExecuteTemplate(&buf, "example", &ex); err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

func generateExampleInput(ex Example) string {
	if ex.Operation.HasInput() {
		return buildShape(&ex.Operation.InputRef, ex.Input, false)
	}
	return ""
}

// generateTypes will generate no types for default examples, but customizations may
// require their own defined types.
func generateTypes(ex Example) string {
	return ""
}

// buildShape will recursively build the referenced shape based on the json object
// provided.
// isMap will dictate how the field name is specified. If isMap is true, we will expect
// the member name to be quotes like "Foo".
func buildShape(ref *ShapeRef, shapes map[string]interface{}, isMap bool) string {
	order := make([]string, len(shapes))
	for k := range shapes {
		order = append(order, k)
	}
	sort.Strings(order)

	ret := ""
	for _, name := range order {
		if name == "" {
			continue
		}
		shape := shapes[name]

		// If the shape isn't a map, we want to export the value, since every field
		// defined in our shapes are exported.
		if len(name) > 0 && !isMap && strings.ToLower(name[0:1]) == name[0:1] {
			name = strings.Title(name)
		}

		memName := name
		if isMap {
			memName = fmt.Sprintf("%q", memName)
		}

		switch v := shape.(type) {
		case map[string]interface{}:
			ret += buildComplex(name, memName, ref, v)
		case []interface{}:
			ret += buildList(ref, name, memName, v)
		default:
			ret += buildScalar(name, memName, ref, v)
		}
	}
	return ret
}

func buildList(ref *ShapeRef, name, memName string, v []interface{}) string {
	ret := ""

	if len(v) == 0 || ref == nil {
		return ""
	}

	t := ""
	dataType := ""
	format := ""
	isComplex := false
	passRef := ref

	if ref.Shape.MemberRefs[name] != nil {
		t = getType(ref.Shape.MemberRefs[name].Shape.MemberRef.Shape)
		dataType = ref.Shape.MemberRefs[name].Shape.MemberRef.Shape.Type
		passRef = ref.Shape.MemberRefs[name]
		if dataType == "map" {
			t = fmt.Sprintf("map[string]%s.%s", ref.API.PackageName(), ref.Shape.MemberRefs[name].Shape.MemberRef.Shape.ValueRef.Shape.ShapeName)
			passRef = &ref.Shape.MemberRefs[name].Shape.MemberRef.Shape.ValueRef
		}
	} else if ref.Shape.MemberRef.Shape != nil && ref.Shape.MemberRef.Shape.MemberRefs[name] != nil {
		t = getType(ref.Shape.MemberRef.Shape.MemberRefs[name].Shape.MemberRef.Shape)
		dataType = ref.Shape.MemberRef.Shape.MemberRefs[name].Shape.MemberRef.Shape.Type
		passRef = &ref.Shape.MemberRef.Shape.MemberRefs[name].Shape.MemberRef
	} else {
		t = getType(ref.Shape.MemberRef.Shape)
		dataType = ref.Shape.MemberRef.Shape.Type
		passRef = &ref.Shape.MemberRef
	}

	switch v[0].(type) {
	case string:
		format = "%s"
	case bool:
		format = "%t"
	case float64:
		if dataType == "integer" || dataType == "int64" {
			format = "%d"
		} else {
			format = "%f"
		}
	default:
		if ref.Shape.MemberRefs[name] != nil {
		} else {
			passRef = ref.Shape.MemberRef.Shape.MemberRefs[name]

			// if passRef is nil that means we are either in a map or within a nested array
			if passRef == nil {
				passRef = &ref.Shape.MemberRef
			}
		}
		isComplex = true
	}
	ret += fmt.Sprintf("%s: []*%s {\n", memName, t)
	for _, elem := range v {
		if isComplex {
			ret += fmt.Sprintf("{\n%s\n},\n", buildShape(passRef, elem.(map[string]interface{}), false))
		} else {
			if dataType == "integer" || dataType == "int64" || dataType == "long" {
				elem = int(elem.(float64))
			}
			ret += fmt.Sprintf("%s,\n", getValue(t, fmt.Sprintf(format, elem)))
		}
	}
	ret += "},\n"
	return ret
}

func buildScalar(name, memName string, ref *ShapeRef, shape interface{}) string {
	if ref == nil || ref.Shape == nil {
		return ""
	} else if ref.Shape.MemberRefs[name] == nil {
		if ref.Shape.MemberRef.Shape != nil && ref.Shape.MemberRef.Shape.MemberRefs[name] != nil {
			return correctType(memName, ref.Shape.MemberRef.Shape.MemberRefs[name].Shape.Type, shape)
		}
		if ref.Shape.Type != "structure" && ref.Shape.Type != "map" {
			return correctType(memName, ref.Shape.Type, shape)
		}
		return ""
	}

	switch v := shape.(type) {
	case bool:
		return convertToCorrectType(memName, ref.Shape.MemberRefs[name].Shape.Type, fmt.Sprintf("%t", v))
	case int:
		if ref.Shape.MemberRefs[name].Shape.Type == "timestamp" {
			return parseTimeString(ref, memName, fmt.Sprintf("%d", v))
		} else {
			return convertToCorrectType(memName, ref.Shape.MemberRefs[name].Shape.Type, fmt.Sprintf("%d", v))
		}
	case float64:
		return convertToCorrectType(memName, ref.Shape.MemberRefs[name].Shape.Type, fmt.Sprintf("%f", v))
	case string:
		t := ref.Shape.MemberRefs[name].Shape.Type
		switch t {
		case "timestamp":
			return parseTimeString(ref, memName, fmt.Sprintf("%s", v))
		case "blob":
			if (ref.Shape.MemberRefs[name].Streaming || ref.Shape.MemberRefs[name].Shape.Streaming) && ref.Shape.Payload == name {
				return fmt.Sprintf("%s: aws.ReadSeekCloser(bytes.NewBuffer([]byte(%q))),\n", memName, v)
			} else {
				return fmt.Sprintf("%s: []byte(%q),\n", memName, v)
			}
		default:
			return convertToCorrectType(memName, t, v)
		}
	default:
		panic(fmt.Errorf("Unsupported scalar type: %v", reflect.TypeOf(v)))
	}
	return ""
}

func correctType(memName string, t string, value interface{}) string {
	if value == nil {
		return ""
	}

	v := ""
	switch value.(type) {
	case string:
		v = value.(string)
	case int:
		v = fmt.Sprintf("%d", value.(int))
	case float64:
		if t == "integer" || t == "long" || t == "int64" {
			v = fmt.Sprintf("%d", int(value.(float64)))
		} else {
			v = fmt.Sprintf("%f", value.(float64))
		}
	case bool:
		v = fmt.Sprintf("%t", value.(bool))
	}

	return convertToCorrectType(memName, t, v)
}

func convertToCorrectType(memName, t, v string) string {
	return fmt.Sprintf("%s: %s,\n", memName, getValue(t, v))
}

func getValue(t, v string) string {
	switch t {
	case "string":
		return fmt.Sprintf("aws.String(%q)", v)
	case "integer", "long", "int64":
		return fmt.Sprintf("aws.Int64(%s)", v)
	case "float", "float64":
		return fmt.Sprintf("aws.Float64(%s)", v)
	case "boolean":
		return fmt.Sprintf("aws.Bool(%s)", v)
	default:
		panic("Unsupported type: " + t)
	}
}

func buildComplex(name, memName string, ref *ShapeRef, v map[string]interface{}) string {
	shapeName, t := "", ""
	if ref == nil {
		return buildShape(nil, v, true)
	}

	member := ref.Shape.MemberRefs[name]

	if member != nil && member.Shape != nil {
		shapeName = ref.Shape.MemberRefs[name].Shape.ShapeName
		t = ref.Shape.MemberRefs[name].Shape.Type
	} else {
		shapeName = ref.Shape.ShapeName
		t = ref.Shape.Type
	}

	switch t {
	case "structure":
		passRef := ref.Shape.MemberRefs[name]
		// passRef will be nil if the entry is a map. In that case
		// we want to pass the reference, because the previous call
		// passed the value reference.
		if passRef == nil {
			passRef = ref
		}
		return fmt.Sprintf(`%s: &%s.%s{
				%s
			},
			`, memName, ref.API.PackageName(), shapeName, buildShape(passRef, v, false))
	case "map":
		shapeType := getNestedType(ref.Shape.MemberRefs[name].Shape.ValueRef.Shape)
		return fmt.Sprintf(`%s: map[string]%s{
				%s
			},
			`, name, shapeType, buildShape(&ref.Shape.MemberRefs[name].Shape.ValueRef, v, true))
	}

	return ""
}

// AttachExamples will create a new ExamplesDefinition from the examples file
// and reference the API object.
func (a *API) AttachExamples(filename string) {
	a.Setup()
	a.customizationPasses()
	p := ExamplesDefinition{API: a}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		panic(err)
	}

	p.setup()
}

func (p *ExamplesDefinition) setup() {
	keys := p.Examples.Names()
	for _, n := range keys {
		examples := p.Examples[n]
		for i, e := range examples {
			n = p.ExportableName(n)
			e.OperationName = n
			e.API = p.API
			e.Index = fmt.Sprintf("shared%02d", i)
			e.VisitedErrors = map[string]struct{}{}
			op := p.API.Operations[e.OperationName]
			e.OperationName = p.ExportableName(e.OperationName)
			e.Operation = op
			p.Examples[n][i] = e
		}
	}

	p.API.Examples = p.Examples
}

var exampleHeader = template.Must(template.New("exampleHeader").Parse(`
import (
	"fmt"
	"bytes"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/{{ .PackageName }}"
)

var _ time.Duration
var _ bytes.Buffer
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

`))

// ExamplesGoCode will return a code representation of the entry within the
// examples.json file.
func (a *API) ExamplesGoCode() string {
	var buf bytes.Buffer
	if err := exampleHeader.ExecuteTemplate(&buf, "exampleHeader", &a); err != nil {
		panic(err)
	}

	code := a.Examples.GoCode()
	if len(code) == 0 {
		return ""
	}

	buf.WriteString(code)
	return buf.String()
}

// TODO: In the operation docuentation where we list errors, this needs to be done
// there as well.
func (ex *Example) HasVisitedError(errRef *ShapeRef) bool {
	errName := errRef.Shape.ErrorCodeName()
	_, ok := ex.VisitedErrors[errName]
	ex.VisitedErrors[errName] = struct{}{}
	return ok
}

func getNestedType(shape *Shape) string {
	switch shape.Type {
	case "string":
		return "*string"
	case "integer", "int64":
		return "*int64"
	case "bool":
		return "*bool"
	case "list":
		return fmt.Sprintf("[]%s", getNestedType(shape.MemberRef.Shape))
	case "structure":
		return fmt.Sprintf("*%s.%s", shape.API.PackageName(), shape.ShapeName)
	default:
		panic("Unsupported shape " + shape.Type)
	}
}

func getType(shape *Shape) string {
	switch shape.Type {
	case "string":
		return "string"
	case "integer", "int64":
		return "int64"
	case "bool":
		return "bool"
	case "structure":
		return fmt.Sprintf("%s.%s", shape.API.PackageName(), shape.ShapeName)
	case "map":
		return fmt.Sprintf("map[string]%s.%s", shape.API.PackageName(), shape.ShapeName)
	default:
		panic("Unsupported shape " + shape.Type)
	}
}

func parseTimeString(ref *ShapeRef, memName, v string) string {
	if ref.Location == "header" {
		return fmt.Sprintf("%s: parseTime(%q, %q),\n", memName, "Mon, 2 Jan 2006 15:04:05 GMT", v)
	} else {
		switch ref.API.Metadata.Protocol {
		case "json", "rest-json":
			return fmt.Sprintf("%s: time.Unix(int64(%s), 0).UTC()", memName, v)
		case "rest-xml", "ec2", "query":
			return fmt.Sprintf("%s: parseTime(%q, %q),\n", memName, "2006-01-02T15:04:05Z", v)
		default:
			panic("Unsupported time type: " + ref.API.Metadata.Protocol)
		}
	}
}

func (ex *Example) MethodName() string {
	return fmt.Sprintf("%s_%s", ex.OperationName, ex.Index)
}
