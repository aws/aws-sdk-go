package api

import (
	"sort"
	"strings"
)

type ShapeRef struct {
	API           *API   `json: "-"`
	Shape         *Shape `json: "-"`
	Documentation string
	ShapeName     string `json:"shape"`
	Location      string
	LocationName  string
	XMLNamespace  XMLInfo
}

type XMLInfo struct {
	Prefix string
	URI    string
}

type Shape struct {
	API           *API `json: "-"`
	ShapeName     string
	Documentation string
	MemberRefs    map[string]*ShapeRef `json:"members"`
	MemberRef     ShapeRef             `json:"member"`
	KeyRef        ShapeRef             `json:"key"`
	ValueRef      ShapeRef             `json:"value"`
	Required      []string
	Payload       string
	Type          string
	Exception     bool
	Enum          []string

	refs    []*ShapeRef
	visited bool
}

func (s *Shape) Rename(newName string) {
	for _, r := range s.refs {
		r.ShapeName = newName
	}

	delete(s.API.Shapes, s.ShapeName)
	s.API.Shapes[newName] = s
	s.ShapeName = newName
}

func (s *Shape) MemberNames() []string {
	i, names := 0, make([]string, len(s.MemberRefs))
	for n, _ := range s.MemberRefs {
		names[i] = n
		i++
	}
	sort.Strings(names)
	return names
}

func (ref *ShapeRef) GoTypeElem() string {
	switch ref.Shape.Type {
	case "structure":
		return ref.Shape.ShapeName
	case "map":
		return "map[" + ref.Shape.KeyRef.GoTypeElem() + "]" + ref.Shape.ValueRef.GoType()
	case "list":
		return "[]" + ref.Shape.MemberRef.GoType()
	case "boolean":
		return "bool"
	case "string":
		return "string"
	case "blob":
		return "[]byte"
	case "integer":
		return "int"
	case "long":
		return "int64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "timestamp":
		ref.API.imports["time"] = true
		return "time.Time"
	default:
		panic("Unsupported shape type: " + ref.Shape.Type)
	}
}

func (ref *ShapeRef) GoType() string {
	return "*" + ref.GoTypeElem()
}

func (ref *ShapeRef) GoTags() string {
	code := "`"
	if ref.Location != "" {
		code += `location:"` + ref.Location + `" `
	}
	if ref.LocationName != "" {
		code += `locationName:"` + ref.LocationName + `" `
	}
	code += `type:"` + ref.Shape.Type + `" `

	// embed the timestamp type for easier lookups
	if ref.Shape.Type == "timestamp" {
		code += `timestampFormat:"`
		if ref.Location == "header" {
			code += "rfc822"
		} else {
			switch ref.API.Metadata.Protocol {
			case "json", "rest-json":
				code += "unix"
			case "rest-xml", "ec2", "query":
				code += "iso8601"
			}
		}
		code += `" `
	}

	if ref.Shape.Payload != "" {
		code += `payload:"` + ref.Shape.Payload + `" `
	}
	if len(ref.Shape.Required) > 0 {
		code += `required:"` + strings.Join(ref.Shape.Required, ",") + `" `
	}
	if strings.Contains(ref.API.Metadata.Protocol, "json") {
		code += `json:",omitempty"`
	}

	return strings.TrimSpace(code) + "`"
}

func (s *Shape) GoCode() string {
	code := "type " + s.ShapeName + " "
	switch s.Type {
	case "structure":
		code += "struct {\n"
		for _, n := range s.MemberNames() {
			m := s.MemberRefs[n]

			// TODO figure out why this is needed
			m.Shape = s.API.Shapes[m.ShapeName]
			m.API = s.API

			code += n + " " + m.GoType() + " " + m.GoTags() + "\n"
		}
		metaStruct := strings.ToLower(s.ShapeName[0:1]) + s.ShapeName[1:] + "Metadata"
		ref := &ShapeRef{ShapeName: s.ShapeName, API: s.API, Shape: s}
		code += "\n" + metaStruct + "\n"
		code += "}\n\n"
		code += "type " + metaStruct + " struct {\n"
		code += "SDKShapeTraits bool " + ref.GoTags()
		code += "}"
	default:
		panic("Cannot generate toplevel shape for " + s.Type)
	}

	return gofmt(code)
}
