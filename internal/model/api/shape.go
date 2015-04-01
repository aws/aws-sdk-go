package api

import (
	"fmt"
	"sort"
	"strings"

	"github.com/awslabs/aws-sdk-go/internal/util"
)

type ShapeRef struct {
	API           *API   `json: "-"`
	Shape         *Shape `json: "-"`
	Documentation string
	ShapeName     string `json:"shape"`
	Location      string
	LocationName  string
	QueryName     string
	Flattened     bool
	Streaming     bool
	XMLAttribute  bool
	XMLNamespace  XMLInfo
	Payload       string
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
	Flattened     bool
	Streaming     bool
	Location      string
	LocationName  string
	XMLNamespace  XMLInfo

	refs []*ShapeRef
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

func (s *Shape) GoType() string {
	switch s.Type {
	case "structure":
		return "*" + s.ShapeName
	case "map":
		return "*map[string]" + s.ValueRef.GoType()
	case "list":
		return "[]" + s.MemberRef.GoType()
	case "boolean":
		return "*bool"
	case "string", "character":
		return "*string"
	case "blob":
		return "[]byte"
	case "integer", "long":
		return "*int64"
	case "float", "double":
		return "*float64"
	case "timestamp":
		s.API.imports["time"] = true
		return "*time.Time"
	default:
		panic("Unsupported shape type: " + s.Type)
	}
}

func (ref *ShapeRef) GoType() string {
	if ref.Shape == nil {
		panic(fmt.Errorf("missing shape definition on reference for %#v", ref))
	}

	return ref.Shape.GoType()
}

func (s *Shape) GoTypeElem() string {
	t := s.GoType()
	if strings.HasPrefix(t, "*") {
		return t[1:]
	}
	return t
}

func (ref *ShapeRef) GoTypeElem() string {
	if ref.Shape == nil {
		panic(fmt.Errorf("missing shape definition on reference for %#v", ref))
	}

	return ref.Shape.GoTypeElem()
}

func (ref *ShapeRef) GoTags(toplevel bool, isRequired bool) string {
	code := "`"
	if ref.Location != "" {
		code += `location:"` + ref.Location + `" `
	} else if ref.Shape.Location != "" {
		code += `location:"` + ref.Shape.Location + `" `
	}
	if ref.LocationName != "" {
		code += `locationName:"` + ref.LocationName + `" `
	} else if ref.Shape.LocationName != "" {
		code += `locationName:"` + ref.Shape.LocationName + `" `
	}
	if ref.QueryName != "" {
		code += `queryName:"` + ref.QueryName + `" `
	}
	if ref.Shape.MemberRef.LocationName != "" {
		code += `locationNameList:"` + ref.Shape.MemberRef.LocationName + `" `
	}
	if ref.Shape.KeyRef.LocationName != "" {
		code += `locationNameKey:"` + ref.Shape.KeyRef.LocationName + `" `
	}
	if ref.Shape.ValueRef.LocationName != "" {
		code += `locationNameValue:"` + ref.Shape.ValueRef.LocationName + `" `
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

	if ref.Shape.Flattened || ref.Flattened {
		code += `flattened:"true" `
	}

	if ref.XMLAttribute {
		code += `xmlAttribute:"true" `
	}

	if isRequired {
		code += `required:"true"`
	}

	if toplevel {
		if ref.Shape.Payload != "" {
			code += `payload:"` + ref.Shape.Payload + `" `
		}
		if ref.XMLNamespace.Prefix != "" {
			code += `xmlPrefix:"` + ref.XMLNamespace.Prefix + `" `
		} else if ref.Shape.XMLNamespace.Prefix != "" {
			code += `xmlPrefix:"` + ref.Shape.XMLNamespace.Prefix + `" `
		}
		if ref.XMLNamespace.URI != "" {
			code += `xmlURI:"` + ref.XMLNamespace.URI + `" `
		} else if ref.Shape.XMLNamespace.URI != "" {
			code += `xmlURI:"` + ref.Shape.XMLNamespace.URI + `" `
		}

	}

	return strings.TrimSpace(code) + "`"
}

func (ref *ShapeRef) Docstring() string {
	if ref.Documentation != "" {
		return docstring(ref.Documentation)
	}
	return ref.Shape.Docstring()
}

func (s *Shape) Docstring() string {
	return docstring(s.Documentation)
}

func (s *Shape) GoCode() string {
	code := s.Docstring() + "type " + s.ShapeName + " "
	switch s.Type {
	case "structure":
		code += "struct {\n"
		for _, n := range s.MemberNames() {
			m := s.MemberRefs[n]
			code += m.Docstring()
			if (m.Streaming || m.Shape.Streaming) && s.Payload == n {
				rtype := "io.ReadSeeker"
				if len(s.refs) > 1 {
					rtype = "aws.ReaderSeekCloser"
				} else if strings.HasSuffix(s.ShapeName, "Output") {
					rtype = "io.ReadCloser"
				}

				s.API.imports["io"] = true
				code += n + " " + rtype + " " + m.GoTags(false, s.IsRequired(n)) + "\n\n"
			} else {
				code += n + " " + m.GoType() + " " + m.GoTags(false, s.IsRequired(n)) + "\n\n"
			}
		}
		metaStruct := "metadata" + s.ShapeName
		ref := &ShapeRef{ShapeName: s.ShapeName, API: s.API, Shape: s}
		code += "\n" + metaStruct + "  `json:\"-\", xml:\"-\"`\n"
		code += "}\n\n"
		code += "type " + metaStruct + " struct {\n"
		code += "SDKShapeTraits bool " + ref.GoTags(true, false)
		code += "}"
	default:
		panic("Cannot generate toplevel shape for " + s.Type)
	}

	return util.GoFmt(code)
}

func (s *Shape) IsRequired(member string) bool {
	for _, n := range s.Required {
		if n == member {
			return true
		}
	}
	return false
}
