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
	"commentify":       commentify,
	"wrap":             wrap,
	"generateFunction": generateExample,
	"generateTypes":    generateTypes,
}

var exampleCustomizations = map[string]template.FuncMap{}

var exampleTmpls = template.Must(template.New("example").Funcs(exampleFuncMap).Parse(`
{{ generateTypes . }}
{{ commentify (wrap .Title 80 false) }}
//
{{ commentify (wrap .Description 80 false) }}
func Example{{ .API.StructName }}_{{ .OperationName }}_{{.Index}}() {
	svc := {{ .API.PackageName }}.New(session.New())

	{{ generateFunction . }}
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
	for _, examples := range exs {
		for i, ex := range examples {
			ex.Index = fmt.Sprintf("%c", rune(i+int('a')))
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

func generateExample(ex Example) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("input := &%s.%s {\n", ex.API.PackageName(), ex.Operation.InputRef.Shape.ShapeName))
	if ex.Operation.HasInput() {
		buf.WriteString(buildShape(&ex.Operation.InputRef, ex.Input, false))
	}
	buf.WriteString("}\n")
	buf.WriteString(fmt.Sprintf("\nresult, err := svc.%s(input)\n", ex.OperationName))
	return buf.String()
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
		t := ""
		if isMap {
			memName = fmt.Sprintf("%q", memName)
		} else if ref != nil {
		}

		switch v := shape.(type) {
		case bool:
			ret += fmt.Sprintf("%s: aws.Bool(%t),\n", memName, v)
		case int:
			if ref.Shape.MemberRefs[name].Shape.Type == "timestamp" {
				ret += parseTimeString(ref, memName, fmt.Sprintf("%d", v))
			} else {
				ret += fmt.Sprintf("%s: aws.Int64(%d),\n", memName, v)
			}
		case string:
			if ref != nil && ref.Shape.MemberRefs[name] != nil && ref.Shape.MemberRefs[name].Shape.Type == "timestamp" {
				ret += parseTimeString(ref, memName, fmt.Sprintf("%s", v))
			} else if ref != nil && ref.Shape.MemberRefs[name] != nil && ref.Shape.MemberRefs[name].Shape.Type == "blob" {
				if (ref.Shape.MemberRefs[name].Streaming || ref.Shape.MemberRefs[name].Shape.Streaming) && ref.Shape.Payload == name {
					ret += fmt.Sprintf("%s: aws.ReadSeekCloser(bytes.NewBuffer([]byte(%q))),\n", memName, v)
				} else {
					ret += fmt.Sprintf("%s: []byte(%q),\n", memName, v)
				}
			} else {
				ret += fmt.Sprintf("%s: aws.String(%q),\n", memName, v)
			}
		case map[string]interface{}:
			if ref == nil {
				ret += fmt.Sprintf(`%s: {
				%s
			},
			`, memName, buildShape(nil, v, isMap))
			} else {
				if isMap {
					t = ref.Shape.Type
				} else {
					t = ref.Shape.MemberRefs[name].Shape.Type
				}
				switch t {
				case "structure":
					shapeName := ""
					if ref.Shape.MemberRefs[name] != nil {
						shapeName = ref.Shape.MemberRefs[name].Shape.ShapeName
					} else {
						shapeName = ref.Shape.ShapeName
					}
					ret += fmt.Sprintf(`%s: &%s.%s{
				%s
			},
			`, memName, ref.API.PackageName(), shapeName, buildShape(ref.Shape.MemberRefs[name], v, false))
				case "map":
					shapeType := getNestedType(ref.Shape.MemberRefs[name].Shape.ValueRef.Shape)
					ret += fmt.Sprintf(`%s: map[string]%s{
				%s
			},
			`, name, shapeType, buildShape(&ref.Shape.MemberRefs[name].Shape.ValueRef, v, true))
				}
			}
		}
	}
	return ret
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
			e.VisitedErrors = map[string]struct{}{}
			op := p.API.Operations[e.OperationName]
			e.OperationName = p.ExportableName(e.OperationName)
			e.Operation = op
			p.Examples[n][i] = e
		}
	}

	p.API.Examples = p.Examples
}

// ExamplesGoCode will return a code representation of the entry within the
// examples.json file.
func (a *API) ExamplesGoCode() string {
	var buf bytes.Buffer
	imports := []string{
		"fmt",
		"bytes",
		"time",
		"github.com/aws/aws-sdk-go/aws",
		"github.com/aws/aws-sdk-go/aws/session",
		"github.com/aws/aws-sdk-go/aws/awserr",
		"github.com/aws/aws-sdk-go/service/" + a.PackageName(),
	}
	buf.WriteString("import (\n")
	for _, importValue := range imports {
		buf.WriteString(fmt.Sprintf("%q\n", importValue))
	}
	buf.WriteString(")\n")
	buf.WriteString("\nvar _ time.Duration\n")
	buf.WriteString("var _ bytes.Buffer\n")
	buf.WriteString("var _ aws.Config\n\n")
	buf.WriteString(`func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}`)
	buf.WriteString("\n")

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
