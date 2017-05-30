// +build codegen

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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
func buildShape(ref *ShapeRef, shapes map[string]interface{}, isMap bool) string {
	order := make([]string, len(shapes))
	for k := range shapes {
		order = append(order, k)
	}
	sort.Strings(order)

	ret := ""
	for _, name := range order {
		shape := shapes[name]

		// If the shape isn't a map, we want to export the value, since every field
		// defined in our shapes are exported.
		if len(name) > 0 && !isMap && strings.ToLower(name[0:1]) == name[0:1] {
			name = strings.Title(name)
		}

		memName := name
		if isMap {
			memName = fmt.Sprintf("%q", memName)
		} else if ref != nil {
		}

		switch v := shape.(type) {
		case map[string]interface{}:
			ret += buildComplex(name, memName, ref, v)
		default:
			ret += buildScalar(name, memName, ref, v)
		}
	}
	return ret
}

func buildScalar(name, memName string, ref *ShapeRef, shape interface{}) string {
	switch v := shape.(type) {
	case bool:
		return fmt.Sprintf("%s: aws.Bool(%t),\n", memName, v)
	case int:
		if ref.Shape.MemberRefs[name].Shape.Type == "timestamp" {
			return parseTimeString(ref, memName, fmt.Sprintf("%d", v))
		} else {
			return fmt.Sprintf("%s: aws.Int64(%d),\n", memName, v)
		}
	case string:
		if ref != nil && ref.Shape.MemberRefs[name] != nil && ref.Shape.MemberRefs[name].Shape.Type == "timestamp" {
			return parseTimeString(ref, memName, fmt.Sprintf("%s", v))
		} else if ref != nil && ref.Shape.MemberRefs[name] != nil && ref.Shape.MemberRefs[name].Shape.Type == "blob" {
			if (ref.Shape.MemberRefs[name].Streaming || ref.Shape.MemberRefs[name].Shape.Streaming) && ref.Shape.Payload == name {
				return fmt.Sprintf("%s: aws.ReadSeekCloser(bytes.NewBuffer([]byte(%q))),\n", memName, v)
			} else {
				return fmt.Sprintf("%s: []byte(%q),\n", memName, v)
			}
		} else {
			return fmt.Sprintf("%s: aws.String(%q),\n", memName, v)
		}
	}
	return ""
}

func buildComplex(name, memName string, ref *ShapeRef, v map[string]interface{}) string {
	shapeName, t := "", ""
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
		return fmt.Sprintf(`%s: &%s.%s{
				%s
			},
			`, memName, ref.API.PackageName(), shapeName, buildShape(ref.Shape.MemberRefs[name], v, false))
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
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/awserr"
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
	case "int":
		return "*int64"
	case "bool":
		return "*bool"
	case "list":
		return fmt.Sprintf("[]%s", getNestedType(shape.MemberRef.Shape))
	case "structure":
		return fmt.Sprintf("*%s.%s", shape.API.PackageName(), shape.ShapeName)
	default:
		panic("Unsupported shape " + shape.ValueRef.Shape.Type)
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
