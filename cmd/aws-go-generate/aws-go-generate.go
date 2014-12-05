package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/doc"
	"go/format"
	"io"
	"os"
	"sort"
	"strings"
	"text/template"

	"code.google.com/p/go.net/html"
	"github.com/aarzilli/sandblast"
)

type Metadata struct {
	APIVersion          string
	EndpointPrefix      string
	JSONVersion         string
	ServiceAbbreviation string
	ServiceFullName     string
	SignatureVersion    string
	TargetPrefix        string
	Protocol            string
}

type HTTPOptions struct {
	Method     string
	RequestURI string
}

type Operation struct {
	Name          string
	Documentation string
	HTTP          HTTPOptions
	Input         ShapeRef
	Output        *ShapeRef
}

type Operations map[string]Operation

func (l Operations) Sorted() []Operation {
	var names []string
	for name := range l {
		names = append(names, name)
	}
	sort.Strings(names)

	var ops []Operation
	for _, name := range names {
		ops = append(ops, l[name])
	}
	return ops
}

type Error struct {
	Code           string
	HTTPStatusCode int
	SenderFault    bool
}

type ShapeRef struct {
	Shape         string
	Documentation string
}

type Shapes map[string]Shape

func (l Shapes) Sorted() []Shape {
	var names []string
	for name := range l {
		names = append(names, name)
	}
	sort.Strings(names)

	var shapes []Shape
	for _, name := range names {
		shapes = append(shapes, l[name])
	}
	return shapes
}

func (l Shapes) Structures() []Shape {
	var shapes []Shape
	for _, shape := range l.Sorted() {
		if shape.Type == "structure" && !shape.Exception {
			shapes = append(shapes, shape)
		}
	}
	return shapes
}

func (l Shapes) Exceptions() []Shape {
	var shapes []Shape
	for _, shape := range l.Sorted() {
		if shape.Type == "structure" && shape.Exception {
			shapes = append(shapes, shape)
		}
	}
	return shapes
}

type Shape struct {
	Name          string
	Type          string
	Required      []string
	Members       map[string]ShapeRef
	Member        *ShapeRef
	Key           *ShapeRef
	Value         *ShapeRef
	Error         Error
	Exception     bool
	Documentation string
	Min           int
	Max           int
	Pattern       string
	Sensitive     bool
}

type Service struct {
	Name          string
	FullName      string
	PackageName   string
	Metadata      Metadata
	Documentation string
	Operations    Operations
	Shapes        Shapes
}

func (s *Service) Type(name string) string {
	shape := s.Shapes[name]

	switch shape.Type {
	case "structure":
		return shape.Name
	case "integer", "long":
		return "int"
	case "double":
		return "float64"
	case "string":
		return "string"
	case "map":
		return "map[" + s.Type(shape.Key.Shape) + "]" + s.Type(shape.Value.Shape)
	case "list":
		return "[]" + s.Type(shape.Member.Shape)
	case "boolean":
		return "bool"
	case "blob":
		return "[]byte"
	case "timestamp":
		return "time.Time"
	}

	panic(fmt.Errorf("type %s (%s) not found", shape.Name, shape.Type))
}

var replacements = map[string]string{
	"Id":  "ID",
	"Arn": "ARN",
}

func (s *Service) FixName(name string) string {
	for from, to := range replacements {
		name = strings.Replace(name, from, to, -1)
	}
	return name
}

func (s *Service) Tags(shape Shape, name string, ref ShapeRef) string {
	required := false
	for _, s := range shape.Required {
		if s == name {
			required = true
			break
		}
	}

	tag := name
	if !required {
		tag += ",omitempty"
	}

	return fmt.Sprintf("`json:%q`", tag)
}

func (*Service) Doc(name, doco string) string {
	node, err := html.Parse(strings.NewReader(doco))
	if err != nil {
		return ""
	}

	_, v, err := sandblast.Extract(node)
	if err != nil {
		return ""
	}

	v = strings.TrimSpace(v)
	if v == "" {
		return "// " + name + " is undocumented.\n"
	}

	if name != "" {
		v = name + " " + strings.ToLower(v[0:1]) + v[1:]
	}

	out := bytes.NewBuffer(nil)
	doc.ToText(out, v, "// ", "", 72)
	return out.String()
}

func (s *Service) Generate(name string, w io.Writer) error {
	// do all the setup
	for name, shape := range s.Shapes {
		shape.Name = name
		s.Shapes[name] = shape
	}

	s.FullName = s.Metadata.ServiceFullName
	s.PackageName = strings.ToLower(name)
	s.Name = name

	pkg := bytes.NewBuffer(nil)
	if err := pkgTmpl.Execute(pkg, s); err != nil {
		return err
	}

	b, err := format.Source(pkg.Bytes())
	if err != nil {
		fmt.Println(pkg.String())
		return err
	}

	if _, err := w.Write(b); err != nil {
		return err
	}

	return nil
}

var pkgTmpl = template.Must(template.New("pkg").Parse(`
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
   client *aws.JSONClient
}

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

{{ range $s := .Operations.Sorted }}
{{ $.Doc $s.Name $s.Documentation }} func (c *{{ $.Name }}) {{$s.Name}}(req {{ $.Type $s.Input.Shape }}) ({{if $s.Output }}resp *{{ $.Type $s.Output.Shape  }},{{end}} err error)  {
  {{ if $s.Output }} resp = &{{ $.Type $s.Output.Shape }}{} {{ else }} // NRE {{ end }}
  err = c.client.Do("{{ $s.Name }}", "{{ $s.HTTP.Method }}", "{{ $s.HTTP.RequestURI }}", req, {{ if $s.Output }} resp {{ else }} nil {{ end }})
  return
}
{{ end }}

{{ range $s := .Shapes.Structures }}

type {{ $s.Name }} struct { {{ range $name, $member := $s.Members }}
    {{ $.FixName $name }} {{ $.Type $member.Shape }} {{ $.Tags $s $name $member }} {{ end }}
}

{{ end }}

{{ range $s := .Shapes.Exceptions }}
type {{ $s.Name }} struct {
  // exception
}
{{ end }}

// IMPORTS

var _ time.Time
`))

func main() {
	f, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var service Service
	if err := json.NewDecoder(f).Decode(&service); err != nil {
		panic(err)
	}

	if err := service.Generate(os.Args[1], os.Stdout); err != nil {
		panic(err)
	}
}
