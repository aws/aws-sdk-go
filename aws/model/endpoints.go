package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"
)

type Constraint []interface{}

func (c Constraint) Condition() string {
	str := func(i interface{}) string {
		if i == nil {
			return ""
		}
		return i.(string)
	}

	switch c[1] {
	case "startsWith":
		return fmt.Sprintf("strings.HasPrefix(%s, %q)", str(c[0]), str(c[2]))
	case "notStartsWith":
		return fmt.Sprintf("!strings.HasPrefix(%s, %q)", str(c[0]), str(c[2]))
	case "equals":
		return fmt.Sprintf("%s == %q", str(c[0]), str(c[2]))
	case "notEquals":
		return fmt.Sprintf("%s == %q", str(c[0]), str(c[2]))
	case "oneOf":
		var values []string
		for _, v := range c[2].([]interface{}) {
			values = append(values, str(v))
		}

		var conditions []string
		for _, v := range values {
			conditions = append(conditions, fmt.Sprintf("%s == %q", str(c[0]), v))
		}

		return strings.Join(conditions, " || ")
	default:
		panic(fmt.Sprintf("unknown operator: %v", c[1]))
	}
}

type CredentialScope struct {
	Region           string
	SignatureVersion string
}

type Properties struct {
	CredentialScope CredentialScope
}

type Endpoint struct {
	Name        string
	URI         string
	Properties  Properties
	Constraints []Constraint
}

func (e Endpoint) Conditions() string {
	var conds []string
	for _, c := range e.Constraints {
		conds = append(conds, "("+c.Condition()+")")
	}
	return strings.Join(conds, " && ")
}

type Endpoints map[string][]Endpoint

func (e *Endpoints) Parse(r io.Reader) error {
	return json.NewDecoder(r).Decode(e)
}

func (e Endpoints) Render(w io.Writer) error {
	tmpl, err := template.New("endpoints").Parse(t)
	if err != nil {
		return err
	}

	out := bytes.NewBuffer(nil)
	if err := tmpl.Execute(out, e); err != nil {
		return err
	}

	b, err := format.Source(bytes.TrimSpace(out.Bytes()))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(b))
	return err
}

const t = `
// Package endpoints provides lookups for all AWS service endpoints.
package endpoints

import (
  "strings"
)

// Lookup returns the endpoint for the given service in the given region.
func Lookup(service, region string) string {
  switch service {
    {{ range $name, $endpoints := . }}
    {{ if ne $name "_default" }}
    case "{{ $name}}" :
    {{ range $endpoints }}
      {{ if .Constraints }}if {{ .Conditions }} { {{ end }}
        return format("{{ .URI }}", service, region)
      {{ if .Constraints }} } {{ end }}
    {{ end }}
    {{ end }}
    {{ end }}
    default:
    {{ with $endpoints := index . "_default" }}
    {{ range $endpoints }}
      {{ if .Constraints }}if {{ .Conditions }} { {{ end }}
        return format("{{ .URI }}", service, region)
      {{ if .Constraints }} } {{ end }}
    {{ end }}
    {{ end }}
  }

  panic("unknown endpoint for " + service + " in " + region)
}

func format(uri, service, region string) string {
  uri = strings.Replace(uri, "{scheme}", "https", -1)
  uri = strings.Replace(uri, "{service}", service, -1)
  uri = strings.Replace(uri, "{region}", region, -1)
  return uri
}
`
