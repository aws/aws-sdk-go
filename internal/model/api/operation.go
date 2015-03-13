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

var tplOperation = template.Must(template.New("operation").Parse(`
// {{ .ExportedName }}Request generates a request for the {{ .ExportedName }} operation.
func (c *{{ .API.StructName }}) {{ .ExportedName }}Request(` +
	`{{ if .HasInput }}input {{ .InputRef.GoType }}{{ end }}) ` +
	`(req *aws.Request{{ if .HasOutput }}, output {{ .OutputRef.GoType }}{{ end }}) {
	if op{{ .ExportedName }} == nil {
		op{{ .ExportedName }} = &aws.Operation{
			Name:       "{{ .Name }}",
			{{ if ne .HTTP.Method "" }}HTTPMethod: "{{ .HTTP.Method }}",
			{{ end }}{{ if ne .HTTP.RequestURI "" }}HTTPPath:   "{{ .HTTP.RequestURI }}",
			{{ end }}{{ if ne .OutputRef.ResultWrapper "" }}ResultWrapper: "{{ .OutputRef.ResultWrapper }}",
			{{ end }}
		}
	}

	req = aws.NewRequest(c.Service, op{{ .ExportedName }}, ` +
	`{{ if .HasInput }}input{{ else }}nil{{ end }}, {{ if .HasOutput }}output{{ else }}nil{{ end }})
	{{ if .HasOutput }}output = &{{ .OutputRef.GoTypeElem }}{}
	req.Data = output{{ end }}
	return
}

func (c *{{ .API.StructName }}) {{ .ExportedName }}(` +
	`{{ if .HasInput }}input {{ .InputRef.GoType }}{{ end }}) ` +
	`({{ if .HasOutput }}output {{ .OutputRef.GoType }},{{ end }} err error) {
	req{{ if .HasOutput }}, out{{ end }} := c.{{ .ExportedName }}Request({{ if .HasInput }}input{{ end }})
	{{ if .HasOutput }}output = out
	{{ end }}err = req.Send()
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
