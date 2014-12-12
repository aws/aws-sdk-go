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
	ChecksumFormat      string
	GlobalEndpoint      string
	TimestampFormat     string
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
	Location      string
	LocationName  string
	Wrapper       bool
	ResultWrapper string
	Streaming     bool
}

func (ref *ShapeRef) WrappedType() string {
	if ref.ResultWrapper != "" {
		return "*" + exportable(ref.ResultWrapper)
	}
	return ref.Shape().Type()
}

func (ref *ShapeRef) WrappedLiteral() string {
	if ref.ResultWrapper != "" {
		return "&" + exportable(ref.ResultWrapper) + "{}"
	}
	return ref.Shape().Literal()
}

func (ref *ShapeRef) Shape() *Shape {
	if ref == nil {
		return nil
	}
	return service.Shapes[ref.ShapeName]
}

type Member struct {
	ShapeRef
	Name     string
	Required bool
}

func (m Member) JSONTag() string {
	if !m.Required {
		return fmt.Sprintf("`json:\"%s,omitempty\"`", m.Name)
	}
	return fmt.Sprintf("`json:\"%s\"`", m.Name)
}

func (m Member) XMLTag(wrapper string) string {
	var path []string
	if wrapper != "" {
		path = append(path, wrapper)
	}

	if m.LocationName != "" {
		path = append(path, m.LocationName)
	} else {
		path = append(path, m.Name)
	}

	if m.Shape().ShapeType == "list" && service.Metadata.Protocol != "rest-xml" {
		loc := m.Shape().MemberRef.LocationName
		if loc == "" {
			loc = "member"
		}
		path = append(path, loc)
	}

	return fmt.Sprintf("`xml:\"%s\"`", strings.Join(path, ">"))
}

func (m Member) EC2Tag() string {
	var path []string
	if m.LocationName != "" {
		path = append(path, m.LocationName)
	} else {
		path = append(path, m.Name)
	}

	if m.Shape().ShapeType == "list" {
		loc := m.Shape().MemberRef.LocationName
		if loc == "" {
			loc = "member"
		}
		path = append(path, loc)
	}

	// Literally no idea how to distinguish between a location name that's
	// required (e.g. DescribeImagesRequest#Filters) and one that's weirdly
	// misleading (e.g. ModifyInstanceAttributeRequest#InstanceId) besides this.

	// Use the locationName unless it's missing or unless it starts with a
	// lowercase letter. Not even making this up.
	var name = m.LocationName
	if name == "" || strings.ToLower(name[0:1]) == name[0:1] {
		name = m.Name
	}

	return fmt.Sprintf("`ec2:%q xml:%q`", name, strings.Join(path, ">"))
}

func (m Member) Shape() *Shape {
	return m.ShapeRef.Shape()
}

func (m Member) Type() string {
	if m.Streaming {
		return "io.ReadCloser"
	}
	return m.Shape().Type()
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
	Payload       string
}

func (s *Shape) Message() bool {
	return strings.HasSuffix(s.Name, "Message") && s.ResultWrapper() != ""
}

func (s *Shape) MessageTag() string {
	tag := strings.TrimSuffix(s.ResultWrapper(), "Result") + "Response"
	return fmt.Sprintf("`xml:\"%s\"`", tag)
}

func (s *Shape) Key() *Shape {
	return s.KeyRef.Shape()
}

func (s *Shape) Member() *Shape {
	return s.MemberRef.Shape()
}

func (s *Shape) Members() map[string]Member {
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
			Required: required(name),
			ShapeRef: ref,
		}
	}
	return members
}

func (s *Shape) ResultWrapper() string {
	var wrappers []string

	for _, op := range service.Operations {
		if op.OutputRef != nil && op.OutputRef.ShapeName == s.Name {
			wrappers = append(wrappers, op.OutputRef.ResultWrapper)
		}
	}

	if len(wrappers) == 1 {
		return wrappers[0]
	}

	return ""
}

func (s *Shape) Value() *Shape {
	return s.ValueRef.Shape()
}

func (s *Shape) Literal() string {
	if s.ShapeType == "structure" {
		return "&" + s.Type()[1:] + "{}"
	}
	panic("trying to make a literal non-structure for " + s.Name)
}

func (s *Shape) ElementType() string {
	switch s.ShapeType {
	case "structure":
		return exportable(s.Name)
	case "integer":
		return "int"
	case "long":
		return "int64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "string":
		return "string"
	case "map":
		return "map[" + s.Key().ElementType() + "]" + s.Value().ElementType()
	case "list":
		return "[]" + s.Member().ElementType()
	case "boolean":
		return "bool"
	case "blob":
		return "[]byte"
	case "timestamp":
		return "time.Time"
	}

	panic(fmt.Errorf("type %q (%q) not found", s.Name, s.ShapeType))
}

func (s *Shape) Type() string {
	switch s.ShapeType {
	case "structure":
		return "*" + exportable(s.Name)
	case "integer":
		return "aws.IntegerValue"
	case "long":
		return "aws.LongValue"
	case "float":
		return "aws.FloatValue"
	case "double":
		return "aws.DoubleValue"
	case "string":
		return "aws.StringValue"
	case "map":
		return "map[" + s.Key().ElementType() + "]" + s.Value().ElementType()
	case "list":
		return "[]" + s.Member().ElementType()
	case "boolean":
		return "aws.BooleanValue"
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

func (s Service) Wrappers() map[string]*Shape {
	wrappers := map[string]*Shape{}

	// collect all wrapper types
	for _, op := range s.Operations {
		if op.InputRef != nil && op.InputRef.ResultWrapper != "" {
			wrappers[op.InputRef.ResultWrapper] = op.Input()
		}

		if op.OutputRef != nil && op.OutputRef.ResultWrapper != "" {
			wrappers[op.OutputRef.ResultWrapper] = op.Output()
		}
	}

	// remove all existing types?
	for name := range wrappers {
		if _, ok := s.Shapes[name]; ok {
			delete(wrappers, name)
		}
	}

	return wrappers
}

var service Service

func Load(name string, r io.Reader) error {
	service = Service{}
	if err := json.NewDecoder(r).Decode(&service); err != nil {
		return err
	}

	for name, shape := range service.Shapes {
		shape.Name = name
	}

	service.FullName = service.Metadata.ServiceFullName
	service.PackageName = strings.ToLower(name)
	service.Name = name

	return nil
}
