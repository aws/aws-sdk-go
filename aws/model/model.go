package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
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
	InputRef      *ShapeRef `json:"Input"`
	OutputRef     *ShapeRef `json:"Output"`
}

func (o Operation) Input() *Shape {
	return o.InputRef.Shape()
}

func (o Operation) Output() *Shape {
	return o.OutputRef.Shape()
}

type Error struct {
	Code           string
	HTTPStatusCode int
	SenderFault    bool
}

type ShapeRef struct {
	ShapeName     string `json:"Shape"`
	Documentation string
	Wrapper       bool
	ResultWrapper string
}

func (ref *ShapeRef) Shape() *Shape {
	if ref == nil {
		return nil
	}
	return s.Shapes[ref.ShapeName]
}

type Member struct {
	Name     string
	Shape    *Shape
	Required bool
}

func (m Member) JSONTag() string {
	if !m.Required {
		return fmt.Sprintf("`json:\"%s,omitempty\"`", m.Name)
	}
	return fmt.Sprintf("`json:\"%s\"`", m.Name)
}

func (m Member) Type() string {
	return m.Shape.Type()
}

type Shape struct {
	Name          string
	ShapeType     string `json:"Type"`
	Required      []string
	MemberRefs    map[string]ShapeRef `json:"Members"`
	MemberRef     *ShapeRef           `json:"Member"`
	KeyRef        *ShapeRef           `json:"Key"`
	ValueRef      *ShapeRef           `json:"Value"`
	Error         Error
	Exception     bool
	Documentation string
	Min           int
	Max           int
	Pattern       string
	Sensitive     bool
	Wrapper       bool
}

func (s Shape) Key() *Shape {
	return s.KeyRef.Shape()
}

func (s Shape) Member() *Shape {
	return s.MemberRef.Shape()
}

func (s Shape) Members() map[string]Member {
	required := func(v string) bool {
		for _, s := range s.Required {
			if s == v {
				return true
			}
		}
		return false
	}

	members := map[string]Member{}
	for name, ref := range s.MemberRefs {
		members[name] = Member{
			Name:     name,
			Shape:    ref.Shape(),
			Required: required(name),
		}
	}
	return members
}

func (s Shape) Value() *Shape {
	return s.ValueRef.Shape()
}

func (s Shape) Type() string {
	switch s.ShapeType {
	case "structure":
		return exportable(s.Name)
	case "integer", "long":
		return "int"
	case "double":
		return "float64"
	case "string":
		return "string"
	case "map":
		return "map[" + s.Key().Type() + "]" + s.Value().Type()
	case "list":
		return "[]" + s.Member().Type()
	case "boolean":
		return "bool"
	case "blob":
		return "[]byte"
	case "timestamp":
		return "time.Time"
	}

	panic(fmt.Errorf("type %q (%q) not found", s.Name, s.ShapeType))
}

type Service struct {
	Name          string
	FullName      string
	PackageName   string
	Metadata      Metadata
	Documentation string
	Operations    map[string]Operation
	Shapes        map[string]*Shape
}

var s Service

func Load(name string, r io.Reader) error {
	s = Service{}
	if err := json.NewDecoder(r).Decode(&s); err != nil {
		return err
	}

	for name, shape := range s.Shapes {
		shape.Name = name
	}

	s.FullName = s.Metadata.ServiceFullName
	s.PackageName = strings.ToLower(name)
	s.Name = name

	return nil
}
