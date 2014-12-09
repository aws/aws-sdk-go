package model

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"text/template"
)

func Generate(w io.Writer) error {
	t := template.New("root").Funcs(template.FuncMap{
		"godoc":      godoc,
		"exportable": exportable,
	})
	template.Must(common(t))
	template.Must(jsonClient(t))
	template.Must(queryClient(t))
	template.Must(restXMLClient(t))
	template.Must(restJSONClient(t))

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

func common(t *template.Template) (*template.Template, error) {
	return t.Parse(`
{{ define "header" }}
// Package {{ .PackageName }} provides a client for {{ .FullName }}.
package {{ .PackageName }}

import (
  "encoding/xml"
  "net/http"
  "time"

  "github.com/stripe/aws-go/aws"
  "github.com/stripe/aws-go/aws/gen/endpoints"
)

{{ end }}

{{ define "footer" }}
// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
{{ end }}

`)
}

func jsonClient(t *template.Template) (*template.Template, error) {
	return t.Parse(`
{{ define "json" }}
{{ template "header" $ }}

// {{ .Name }} is a client for {{ .FullName }}.
type {{ .Name }} struct {
  client *aws.JSONClient
}

// New returns a new {{ .Name }} client.
func New(key, secret, region string, client *http.Client) *{{ .Name }} {
  if client == nil {
     client = http.DefaultClient
  }

  service := "{{ .Metadata.EndpointPrefix }}"
  endpoint, service, region := endpoints.Lookup("{{ .Metadata.EndpointPrefix }}", region)

  return &{{ .Name }}{
    client: &aws.JSONClient{
      Signer: &aws.V4Signer{
        Key: key,
        Secret: secret,
        Service: service,
        Region: region,
        IncludeXAmzContentSha256: true,
      },
      Client: client,
      Endpoint: endpoint,
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

`)
}

func queryClient(t *template.Template) (*template.Template, error) {
	return t.Parse(`
{{ define "query" }}
{{ template "header" $ }}

// {{ .Name }} is a client for {{ .FullName }}.
type {{ .Name }} struct {
  client *aws.QueryClient
}

// New returns a new {{ .Name }} client.
func New(key, secret, region string, client *http.Client) *{{ .Name }} {
  if client == nil {
     client = http.DefaultClient
  }

  service := "{{ .Metadata.EndpointPrefix }}"
  endpoint, service, region := endpoints.Lookup("{{ .Metadata.EndpointPrefix }}", region)

  return &{{ .Name }}{
    client: &aws.QueryClient{
      Signer: &aws.V4Signer{
        Key: key,
        Secret: secret,
        Service: service,
        Region: region,
        IncludeXAmzContentSha256: true,
      },
      Client: client,
      Endpoint: endpoint,
      APIVersion: "{{ .Metadata.APIVersion }}",
    },
  }
}

{{ range $name, $op := .Operations }}

{{ godoc $name $op.Documentation }} func (c *{{ $.Name }}) {{ exportable $name }}({{ if $op.InputRef }}req {{ exportable $op.InputRef.WrappedType }}{{ end }}) ({{ if $op.OutputRef }}resp *{{ exportable $op.OutputRef.WrappedType }},{{ end }} err error) {
  {{ if $op.Output }}resp = &{{ exportable $op.OutputRef.WrappedType }}{}{{ else }}// NRE{{ end }}
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
{{ exportable $name }} {{ $m.Type }} {{ $m.XMLTag $s.ResultWrapper }}  {{ end }}
}

{{ end }}
{{ end }}
{{ end }}

{{ range $wname, $s := .Wrappers }}

// {{ exportable $wname }} is a wrapper for {{ $s.Name }}.
type {{ exportable $wname }} struct {
    XMLName xml.Name {{ $s.MessageTag }}
{{ range $name, $m := $s.Members }}
{{ exportable $name }} {{ $m.Type }} {{ $m.XMLTag $wname }}  {{ end }}
}

{{ end }}

{{ template "footer" }}
{{ end }}

`)
}

func restXMLClient(t *template.Template) (*template.Template, error) {
	return t.Parse(`
{{ define "rest-xml" }}
{{ template "header" $ }}

import (
  "bytes"
  "fmt"
  "io"
  "io/ioutil"
  "net/url"
  "strconv"
  "strings"
)

// {{ .Name }} is a client for {{ .FullName }}.
type {{ .Name }} struct {
  client *aws.RestClient
}

// New returns a new {{ .Name }} client.
func New(key, secret, region string, client *http.Client) *{{ .Name }} {
  if client == nil {
     client = http.DefaultClient
  }

  service := "{{ .Metadata.EndpointPrefix }}"
  endpoint, service, region := endpoints.Lookup("{{ .Metadata.EndpointPrefix }}", region)

  return &{{ .Name }}{
    client: &aws.RestClient{
      Signer: &aws.V4Signer{
        Key: key,
        Secret: secret,
        Service: service,
        Region: region,
        IncludeXAmzContentSha256: true,
      },
      Client: client,
      Endpoint: endpoint,
      APIVersion: "{{ .Metadata.APIVersion }}",
    },
  }
}

{{ range $name, $op := .Operations }}

{{ godoc $name $op.Documentation }} func (c *{{ $.Name }}) {{ exportable $name }}({{ if $op.Input }}req {{ exportable $op.Input.Type }}{{ end }}) ({{ if $op.Output }}resp *{{ exportable $op.Output.Type }},{{ end }} err error) {
  {{ if $op.Output }}resp = &{{ $op.Output.Type }}{}{{ else }}// NRE{{ end }}

  var body io.Reader
  var contentType string
  {{ if $op.Input }}
  {{ if $op.Input.Payload }}
  {{ with $m := index $op.Input.Members $op.Input.Payload }}
  {{ if $m.Streaming }}
  body = req.{{ exportable $m.Name  }}
  {{ else }}
  contentType = "application/xml"
  b, err := xml.Marshal(req.{{ exportable $m.Name  }})
  if err != nil {
    return
  }
  body = bytes.NewReader(b)
  {{ end }}
  {{ end }}
  {{ end }}
  {{ end }}


  uri := c.client.Endpoint + "{{ $op.HTTP.RequestURI }}"
  {{ if $op.Input }}
  {{ range $name, $m := $op.Input.Members }}
  {{ if eq $m.Location "uri" }}

  uri = strings.Replace(uri, "{"+"{{ $m.LocationName }}"+"}", req.{{ exportable $name }}, -1)
  uri = strings.Replace(uri, "{"+"{{ $m.LocationName }}+"+"}", req.{{ exportable $name }}, -1)

  {{ end }}
  {{ end }}
  {{ end }}

  q := url.Values{}

  {{ if $op.Input }}
  {{ range $name, $m := $op.Input.Members }}
  {{ if eq $m.Location "querystring" }}


  {{ if eq $m.Shape.ShapeType "string" }}
  if s := req.{{ exportable $name }}; s != "" {
  {{ else if eq $m.Shape.ShapeType "timestamp" }}
  if s := req.{{ exportable $name }}.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {
  {{ else if eq $m.Shape.ShapeType "integer" }}
  if s := strconv.Itoa(req.{{ exportable $name }}); req.{{ exportable $name }} != 0 {
  {{ else }}
  if s := fmt.Sprintf("%v", req.{{ exportable $name }}); s != "" {
  {{ end }}
    q.Set("{{ $m.LocationName }}", s)
  }

  {{ end }}
  {{ end }}
  {{ end }}

  if len(q) > 0 {
    uri += "?" + q.Encode()
  }

  httpReq, err := http.NewRequest("{{ $op.HTTP.Method }}", uri, body)
  if err != nil {
    return
  }

  if contentType != "" {
    httpReq.Header.Set("Content-Type", contentType)
  }

  {{ if $op.Input }}
  {{ range $name, $m := $op.Input.Members }}
  {{ if eq $m.Location "header" }}

  {{ if eq $m.Shape.ShapeType "string" }}
  if s := req.{{ exportable $name }}; s != "" {
  {{ else if eq $m.Shape.ShapeType "timestamp" }}
  if s := req.{{ exportable $name }}.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {
  {{ else if eq $m.Shape.ShapeType "integer" }}
  if s := strconv.Itoa(req.{{ exportable $name }}); req.{{ exportable $name }} != 0 {
  {{ else }}
  if s := fmt.Sprintf("%v", req.{{ exportable $name }}); s != "" {
  {{ end }}
    httpReq.Header.Set("{{ $m.LocationName }}", s)
  }

  {{ end }}
  {{ end }}
  {{ end }}

  httpResp, err := c.client.Do(httpReq)
  if err != nil {
    return
  }

  {{ if $op.Output }}
    {{ with $name := "Body" }}
    {{ with $m := index $op.Output.Members $name }}
    {{ if $m }}

      {{ if $m.Streaming }}
  resp.Body = httpResp.Body
      {{ else }}
  defer httpResp.Body.Close()
  err = xml.NewDecoder(httpResp.Body).Decode(resp)
      {{ end }}


    {{ else }}
  defer httpResp.Body.Close()
    {{ end }}
    {{ end }}

  {{ range $name, $m := $op.Output.Members }}
    {{ if ne $name "Body" }}
      {{ if eq $m.Location "header" }}
        if s := httpResp.Header.Get("{{ $m.LocationName }}"); s != "" {
         {{ if eq $m.Shape.ShapeType "string" }}
          resp.{{ exportable $name }} = s
         {{ else if eq $m.Shape.ShapeType "timestamp" }}
          if t, e := time.Parse(s, time.RFC822); e != nil {
           err = e
           return
          } else {
           resp.{{ exportable $name }} = t
          }
         {{ else if eq $m.Shape.ShapeType "integer" }}
          if n, e := strconv.Atoi(s); e != nil {
           err = e
           return
          } else {
           resp.{{ exportable $name }} = n
          }
         {{ else if eq $m.Shape.ShapeType "boolean" }}
         if v, e := strconv.ParseBool(s); e != nil {
           err = e
           return
          } else {
           resp.{{ exportable $name }} = v
          }
         {{ else }}
         // TODO: add support for {{ $m.Shape.ShapeType }} headers
         {{ end }}
        }
      {{ else if eq $m.Location "headers" }}
      resp.{{ exportable $name }} = {{ $m.Shape.Type }}{}
      for name := range httpResp.Header {
        if strings.HasPrefix(name, "{{ $m.Location  }}") {
          resp.{{ exportable $name }}[name] = httpResp.Header.Get(name)
        }
      }
      {{ else if eq $m.Location "statusCode" }}
        resp.{{ exportable $name }} = httpResp.StatusCode
      {{ else if ne $m.Location "" }}
      // TODO: add support for extracting output members from {{ $m.Location }} to support {{ exportable $name }}
      {{ end }}

    {{ end }}
  {{ end }}
  {{ end }}
  {{ else }}
  defer httpResp.Body.Close()
  {{ end }}


  return
}

{{ end }}

{{ range $name, $s := .Shapes }}
{{ if eq $s.ShapeType "structure" }}
{{ if not $s.Exception }}

// {{ exportable $name }} is undocumented.
type {{ exportable $name }} struct {
{{ range $name, $m := $s.Members }}
{{ exportable $name }} {{ $m.Type }} {{ $m.XMLTag $s.ResultWrapper }}  {{ end }}
}

{{ end }}
{{ end }}
{{ end }}

{{ template "footer" }}
var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
{{ end }}

`)
}

func restJSONClient(t *template.Template) (*template.Template, error) {
	return t.Parse(`
{{ define "rest-json" }}
{{ template "header" $ }}

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "net/url"
  "strconv"
  "strings"
)

// {{ .Name }} is a client for {{ .FullName }}.
type {{ .Name }} struct {
  client *aws.RestClient
}

// New returns a new {{ .Name }} client.
func New(key, secret, region string, client *http.Client) *{{ .Name }} {
  if client == nil {
     client = http.DefaultClient
  }

  service := "{{ .Metadata.EndpointPrefix }}"
  endpoint, service, region := endpoints.Lookup("{{ .Metadata.EndpointPrefix }}", region)

  return &{{ .Name }}{
    client: &aws.RestClient{
      Signer: &aws.V4Signer{
        Key: key,
        Secret: secret,
        Service: service,
        Region: region,
        IncludeXAmzContentSha256: true,
      },
      Client: client,
      Endpoint: endpoint,
      APIVersion: "{{ .Metadata.APIVersion }}",
    },
  }
}

{{ range $name, $op := .Operations }}

{{ godoc $name $op.Documentation }} func (c *{{ $.Name }}) {{ exportable $name }}({{ if $op.Input }}req {{ exportable $op.Input.Type }}{{ end }}) ({{ if $op.Output }}resp *{{ exportable $op.Output.Type }},{{ end }} err error) {
  {{ if $op.Output }}resp = &{{ $op.Output.Type }}{}{{ else }}// NRE{{ end }}

  var body io.Reader
  var contentType string
  {{ if $op.Input }}
  {{ if $op.Input.Payload }}
  {{ with $m := index $op.Input.Members $op.Input.Payload }}
  {{ if $m.Streaming }}
  body = req.{{ exportable $m.Name  }}
  {{ else }}
  contentType = "application/json"
  b, err := json.Marshal(req.{{ exportable $m.Name  }})
  if err != nil {
    return
  }
  body = bytes.NewReader(b)
  {{ end }}
  {{ end }}
  {{ end }}
  {{ end }}


  uri := c.client.Endpoint + "{{ $op.HTTP.RequestURI }}"
  {{ if $op.Input }}
  {{ range $name, $m := $op.Input.Members }}
  {{ if eq $m.Location "uri" }}

  uri = strings.Replace(uri, "{"+"{{ $m.LocationName }}"+"}", req.{{ exportable $name }}, -1)
  uri = strings.Replace(uri, "{"+"{{ $m.LocationName }}+"+"}", req.{{ exportable $name }}, -1)

  {{ end }}
  {{ end }}
  {{ end }}

  q := url.Values{}

  {{ if $op.Input }}
  {{ range $name, $m := $op.Input.Members }}
  {{ if eq $m.Location "querystring" }}


  {{ if eq $m.Shape.ShapeType "string" }}
  if s := req.{{ exportable $name }}; s != "" {
  {{ else if eq $m.Shape.ShapeType "timestamp" }}
  if s := req.{{ exportable $name }}.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {
  {{ else if eq $m.Shape.ShapeType "integer" }}
  if s := strconv.Itoa(req.{{ exportable $name }}); req.{{ exportable $name }} != 0 {
  {{ else }}
  if s := fmt.Sprintf("%v", req.{{ exportable $name }}); s != "" {
  {{ end }}
    q.Set("{{ $m.LocationName }}", s)
  }

  {{ end }}
  {{ end }}
  {{ end }}

  if len(q) > 0 {
    uri += "?" + q.Encode()
  }

  httpReq, err := http.NewRequest("{{ $op.HTTP.Method }}", uri, body)
  if err != nil {
    return
  }

  if contentType != "" {
    httpReq.Header.Set("Content-Type", contentType)
  }

  {{ if $op.Input }}
  {{ range $name, $m := $op.Input.Members }}
  {{ if eq $m.Location "header" }}

  {{ if eq $m.Shape.ShapeType "string" }}
  if s := req.{{ exportable $name }}; s != "" {
  {{ else if eq $m.Shape.ShapeType "timestamp" }}
  if s := req.{{ exportable $name }}.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {
  {{ else if eq $m.Shape.ShapeType "integer" }}
  if s := strconv.Itoa(req.{{ exportable $name }}); req.{{ exportable $name }} != 0 {
  {{ else }}
  if s := fmt.Sprintf("%v", req.{{ exportable $name }}); s != "" {
  {{ end }}
    httpReq.Header.Set("{{ $m.LocationName }}", s)
  }

  {{ end }}
  {{ end }}
  {{ end }}

  httpResp, err := c.client.Do(httpReq)
  if err != nil {
    return
  }

  {{ if $op.Output }}
    {{ with $name := "Body" }}
    {{ with $m := index $op.Output.Members $name }}
    {{ if $m }}

      {{ if $m.Streaming }}
  resp.Body = httpResp.Body
      {{ else }}
  defer httpResp.Body.Close()
  err = xml.NewDecoder(httpResp.Body).Decode(resp)
      {{ end }}


    {{ else }}
  defer httpResp.Body.Close()
    {{ end }}
    {{ end }}

  {{ range $name, $m := $op.Output.Members }}
    {{ if ne $name "Body" }}
      {{ if eq $m.Location "header" }}
        if s := httpResp.Header.Get("{{ $m.LocationName }}"); s != "" {
         {{ if eq $m.Shape.ShapeType "string" }}
          resp.{{ exportable $name }} = s
         {{ else if eq $m.Shape.ShapeType "timestamp" }}
          if t, e := time.Parse(s, time.RFC822); e != nil {
           err = e
           return
          } else {
           resp.{{ exportable $name }} = t
          }
         {{ else if eq $m.Shape.ShapeType "integer" }}
          if n, e := strconv.Atoi(s); e != nil {
           err = e
           return
          } else {
           resp.{{ exportable $name }} = n
          }
         {{ else if eq $m.Shape.ShapeType "boolean" }}
         if v, e := strconv.ParseBool(s); e != nil {
           err = e
           return
          } else {
           resp.{{ exportable $name }} = v
          }
         {{ else }}
         // TODO: add support for {{ $m.Shape.ShapeType }} headers
         {{ end }}
        }
      {{ else if eq $m.Location "headers" }}
      resp.{{ exportable $name }} = {{ $m.Shape.Type }}{}
      for name := range httpResp.Header {
        if strings.HasPrefix(name, "{{ $m.Location  }}") {
          resp.{{ exportable $name }}[name] = httpResp.Header.Get(name)
        }
      }
      {{ else if eq $m.Location "statusCode" }}
        resp.{{ exportable $name }} = httpResp.StatusCode
      {{ else if ne $m.Location "" }}
      // TODO: add support for extracting output members from {{ $m.Location }} to support {{ exportable $name }}
      {{ end }}

    {{ end }}
  {{ end }}
  {{ end }}
  {{ else }}
  defer httpResp.Body.Close()
  {{ end }}


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
var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
var _ json.RawMessage
{{ end }}
`)
}
