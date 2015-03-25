package api

import (
	"bytes"
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
	if op{{ .ExportedName }} == nil {
		op{{ .ExportedName }} = &aws.Operation{
			Name:       "{{ .Name }}",
			{{ if ne .HTTP.Method "" }}HTTPMethod: "{{ .HTTP.Method }}",
			{{ end }}{{ if ne .HTTP.RequestURI "" }}HTTPPath:   "{{ .HTTP.RequestURI }}",
			{{ end }}
		}
	}

	req = aws.NewRequest(c.Service, op{{ .ExportedName }}, input, output)
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
