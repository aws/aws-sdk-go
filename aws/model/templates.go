package model

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"text/template"
)

// TODO: support client-side validation
// TODO: support enum values
// TODO: support exceptions
// TODO: support query clients
// TODO: support rest-xml clients
// TODO: support ec2 clients

func Generate(w io.Writer) error {
	t := template.New("root").Funcs(template.FuncMap{
		"godoc":      godoc,
		"exportable": exportable,
	})
	t, err := t.Parse(awsTmpl)
	if err != nil {
		return err
	}

	out := bytes.NewBuffer(nil)
	if err := t.ExecuteTemplate(out, service.Metadata.Protocol, service); err != nil {
		return err
	}

	b, err := format.Source(out.Bytes())
	if err != nil {
		fmt.Fprint(os.Stderr, out.String())
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(b))
	return err
}

const (
	awsTmpl = `
{{ define "json" }}
{{ template "header" $ }}

// New returns a new {{ .Name }} client.
func New(key, secret, region string, client *http.Client) *{{ .Name }} {
  if client == nil {
     client = http.DefaultClient
  }

  return &{{ .Name }}{
    client: &aws.JSONClient{
      Client: client,
      Region: region,
      Endpoint: fmt.Sprintf("https://{{ .Metadata.EndpointPrefix }}.%s.amazonaws.com", region),
      Prefix: "{{ .Metadata.EndpointPrefix }}",
      Key: key,
      Secret: secret,
      JSONVersion: "{{ .Metadata.JSONVersion }}",
      TargetPrefix: "{{ .Metadata.TargetPrefix }}",
    },
  }
}

{{ range $name, $op := .Operations }}

{{ godoc $name $op.Documentation }} func (c *{{ $.Name }}) {{ exportable $name }}({{ if $op.Input }}req {{ exportable $op.Input.Type }}{{ end }}) ({{ if $op.Output }}resp *{{ exportable $op.Output.Type }},{{ end }} err error) {
  {{ if $op.Output }}resp = &{{ $op.Output.Type }}{}{{ else }}// NRE{{ end }}
  err = c.client.Do("{{ $name }}", "{{ $op.HTTP.Method }}", "{{ $op.HTTP.RequestURI }}", {{ if $op.Input }} req {{ else }} nil {{ end }}, {{ if $op.Output }} resp {{ else }} nil {{ end }})
  return
}

{{ end }}

{{ range $name, $s := .Shapes }}
{{ if eq $s.ShapeType "structure" }}
{{ if not $s.Exception }}

// {{ exportable $name }} is undocumented.
type {{ exportable $name }} struct {
{{ range $name, $m := $s.Members }}
{{ exportable $name }} {{ $m.Type }} {{ $m.JSONTag }}  {{ end }}
}

{{ end }}
{{ end }}
{{ end }}

{{ template "footer" }}
{{ end }}

{{ define "header" }}
// Package {{ .PackageName }} provides a client for {{ .FullName }}.
package {{ .PackageName }}

import (
  "fmt"
  "net/http"
  "time"

  "github.com/stripe/aws-go/aws"
)

// {{ .Name }} is a client for {{ .FullName }}.
type {{ .Name }} struct {
  client aws.Client
}
{{ end }}

{{ define "footer" }}
// avoid errors if the packages aren't referenced
var _ time.Time
{{ end }}
`
)
