package api

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/awslabs/aws-sdk-go/internal/util"
)

type Operation struct {
	API           *API `json: "-"`
	ExportedName  string
	Name          string
	Documentation string
	HTTP          HTTPInfo
	InputRef      ShapeRef `json:"input"`
	OutputRef     ShapeRef `json:"output"`
}

type HTTPInfo struct {
	Method       string
	RequestURI   string
	ResponseCode uint
}

func (o *Operation) HasInput() bool {
	return o.InputRef.ShapeName != ""
}

func (o *Operation) HasOutput() bool {
	return o.OutputRef.ShapeName != ""
}

func (o *Operation) Docstring() string {
	if o.Documentation != "" {
		return docstring(o.Documentation)
	}
	return ""
}

var tplOperation = template.Must(template.New("operation").Parse(`
// {{ .ExportedName }}Request generates a request for the {{ .ExportedName }} operation.
func (c *{{ .API.StructName }}) {{ .ExportedName }}Request(` +
	`input {{ .InputRef.GoType }}) (req *aws.Request, output {{ .OutputRef.GoType }}) {
	oprw.Lock()
	defer oprw.Unlock()

	if op{{ .ExportedName }} == nil {
		op{{ .ExportedName }} = &aws.Operation{
			Name:       "{{ .Name }}",
			{{ if ne .HTTP.Method "" }}HTTPMethod: "{{ .HTTP.Method }}",
			{{ end }}{{ if ne .HTTP.RequestURI "" }}HTTPPath:   "{{ .HTTP.RequestURI }}",
			{{ end }}
		}
	}

	if input == nil {
		input = &{{ .InputRef.GoTypeElem }}{}
	}

	req = c.newRequest(op{{ .ExportedName }}, input, output)
	output = &{{ .OutputRef.GoTypeElem }}{}
	req.Data = output
	return
}

{{ .Docstring }}func (c *{{ .API.StructName }}) {{ .ExportedName }}(` +
	`input {{ .InputRef.GoType }}) (output {{ .OutputRef.GoType }}, err error) {
	req, out := c.{{ .ExportedName }}Request(input)
	output = out
	err = req.Send()
	return
}

var op{{ .ExportedName }} *aws.Operation
`))

func (o *Operation) GoCode() string {
	var buf bytes.Buffer
	err := tplOperation.Execute(&buf, o)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(util.GoFmt(buf.String()))
}

var tplExample = template.Must(template.New("operationExample").Parse(`
func Example{{ .API.StructName }}_{{ .ExportedName }}() {
	svc := {{ .API.PackageName }}.New(nil)

	{{ .ExampleInput }}
	resp, err := svc.{{ .ExportedName }}(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
`))

func (o *Operation) Example() string {
	var buf bytes.Buffer
	err := tplExample.Execute(&buf, o)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(util.GoFmt(buf.String()))
}

func (o *Operation) ExampleInput() string {
	if len(o.InputRef.Shape.MemberRefs) == 0 {
		return fmt.Sprintf("var params *%s.%s",
			o.API.PackageName(), o.InputRef.GoTypeElem())
	}
	e := example{o, map[string]int{}}
	return "params := " + e.traverseAny(o.InputRef.Shape, false, false)
}

type example struct {
	*Operation
	visited map[string]int
}

func (e *example) traverseAny(s *Shape, required, payload bool) string {
	str := ""
	e.visited[s.ShapeName]++

	switch s.Type {
	case "structure":
		str = e.traverseStruct(s, required, payload)
	case "list":
		str = e.traverseList(s, required, payload)
	case "map":
		str = e.traverseMap(s, required, payload)
	default:
		str = e.traverseScalar(s, required, payload)
	}

	e.visited[s.ShapeName]--

	return str
}

var reType = regexp.MustCompile(`\b([A-Z])`)

func (e *example) traverseStruct(s *Shape, required, payload bool) string {
	var buf bytes.Buffer
	buf.WriteString("&" + s.API.PackageName() + "." + s.GoTypeElem() + "{")
	if required {
		buf.WriteString(" // Required")
	}
	buf.WriteString("\n")

	req := make([]string, len(s.Required))
	copy(req, s.Required)
	sort.Strings(req)

	if e.visited[s.ShapeName] < 2 {
		for _, n := range req {
			m := s.MemberRefs[n].Shape
			p := n == s.Payload && (s.MemberRefs[n].Streaming || m.Streaming)
			buf.WriteString(n + ": " + e.traverseAny(m, true, p) + ",")
			if m.Type != "list" && m.Type != "structure" && m.Type != "map" {
				buf.WriteString(" // Required")
			}
			buf.WriteString("\n")
		}

		for _, n := range s.MemberNames() {
			if s.IsRequired(n) {
				continue
			}
			m := s.MemberRefs[n].Shape
			p := n == s.Payload && (s.MemberRefs[n].Streaming || m.Streaming)
			buf.WriteString(n + ": " + e.traverseAny(m, false, p) + ",\n")
		}
	} else {
		buf.WriteString("// Recursive values...\n")
	}

	buf.WriteString("}")
	return buf.String()
}

func (e *example) traverseMap(s *Shape, required, payload bool) string {
	var buf bytes.Buffer
	t := reType.ReplaceAllString(s.GoTypeElem(), s.API.PackageName()+".$1")
	buf.WriteString("&" + t + "{")
	if required {
		buf.WriteString(" // Required")
	}
	buf.WriteString("\n")

	if e.visited[s.ShapeName] < 2 {
		m := s.ValueRef.Shape
		buf.WriteString("\"Key\": " + e.traverseAny(m, true, false) + ",")
		if m.Type != "list" && m.Type != "structure" && m.Type != "map" {
			buf.WriteString(" // Required")
		}
		buf.WriteString("\n// More values...\n")
	} else {
		buf.WriteString("// Recursive values...\n")
	}
	buf.WriteString("}")

	return buf.String()
}

func (e *example) traverseList(s *Shape, required, payload bool) string {
	var buf bytes.Buffer
	t := reType.ReplaceAllString(s.GoTypeElem(), s.API.PackageName()+".$1")
	buf.WriteString(t + "{")
	if required {
		buf.WriteString(" // Required")
	}
	buf.WriteString("\n")

	if e.visited[s.ShapeName] < 2 {
		m := s.MemberRef.Shape
		buf.WriteString(e.traverseAny(m, true, false) + ",")
		if m.Type != "list" && m.Type != "structure" && m.Type != "map" {
			buf.WriteString(" // Required")
		}
		buf.WriteString("\n// More values...\n")
	} else {
		buf.WriteString("// Recursive values...\n")
	}
	buf.WriteString("}")

	return buf.String()
}

func (e *example) traverseScalar(s *Shape, required, payload bool) string {
	str := ""
	switch s.Type {
	case "integer", "long":
		str = `aws.Long(1)`
	case "float", "double":
		str = `aws.Double(1.0)`
	case "string", "character":
		str = `aws.String("` + s.ShapeName + `")`
	case "blob":
		if payload {
			str = `bytes.NewReader([]byte("PAYLOAD"))`
		} else {
			str = `[]byte("PAYLOAD")`
		}
	case "boolean":
		str = `aws.Boolean(true)`
	case "timestamp":
		str = `aws.Time(time.Now())`
	default:
		panic("unsupported shape " + s.Type)
	}

	return str
}
