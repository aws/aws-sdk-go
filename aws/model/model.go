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
	Input         *ShapeRef
	Output        *ShapeRef
}

type Error struct {
	Code           string
	HTTPStatusCode int
	SenderFault    bool
}

type ShapeRef struct {
	Shape         string
	Documentation string
	Wrapper       bool
	ResultWrapper string
}

func (ref ShapeRef) JSONTags(name string, required []string) string {
	omitempty := true
	for _, r := range required {
		if name == r {
			omitempty = false
			break
		}
	}

	if omitempty {
		return fmt.Sprintf("`json:\"%s,omitempty\"`", name)
	}
	return fmt.Sprintf("`json:\"%s\"`", name)
}

type Shape struct {
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
	Wrapper       bool
}

type Service struct {
	Name          string
	FullName      string
	PackageName   string
	Metadata      Metadata
	Documentation string
	Operations    map[string]Operation
	Shapes        map[string]Shape
}

func Parse(name string, r io.Reader) (*Service, error) {
	var service Service
	if err := json.NewDecoder(r).Decode(&service); err != nil {
		return nil, err
	}
	service.init(name)
	return &service, nil
}

func (s *Service) Type(name string) string {
	shape := s.Shapes[name]

	switch shape.Type {
	case "structure":
		return exportable(name)
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

	panic(fmt.Errorf("type %q (%q) not found", name, shape.Type))
}

func (s *Service) init(name string) {
	s.FullName = s.Metadata.ServiceFullName
	s.PackageName = strings.ToLower(name)
	s.Name = name
}
