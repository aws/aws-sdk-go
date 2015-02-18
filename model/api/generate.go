package api

import (
	"go/format"
	"strings"
)

func (s *Shape) Code(a *API) (code string) {
	code = "type " + s.ShapeName + " "
	switch s.Type {
	case "structure":
		code += "struct {\n"
		for n, m := range s.MemberRefs {
			code += n + " " + m.GoType(a) + " " + m.GoTags(a) + "\n"
		}
		metaStruct := strings.ToLower(s.ShapeName[0:1]) + s.ShapeName[1:] + "Metadata"
		ref := &ShapeRef{ShapeName: s.ShapeName}
		code += "\n" + metaStruct + "\n"
		code += "}\n\n"
		code += "type " + metaStruct + " struct {\n"
		code += "SDKShapeTraits bool " + ref.GoTags(a)
		code += "}"
	case "map":
		code += "map[" + s.KeyRef.ShapeName + "]" + s.ValueRef.ShapeName
	}

	formatted, _ := format.Source([]byte(code))
	code = string(formatted)
	return
}

func (ref *ShapeRef) GoType(a *API) (code string) {
	shape := a.Shapes[ref.ShapeName]
	switch shape.Type {
	case "structure":
		code = "*" + shape.ShapeName
	case "map":
		code = "*map[" + stripStar(shape.KeyRef.GoType(a)) + "]" + shape.ValueRef.GoType(a)
	case "list":
		code = "[]" + shape.MemberRef.GoType(a)
	case "boolean":
		code = "*bool"
	case "string":
		code = "*string"
	case "blob":
		code = "[]byte"
	case "integer":
		code = "*int"
	case "long":
		code = "*int64"
	case "float":
		code = "*float32"
	case "double":
		code = "*float64"
	default:
		panic("Unsupported API type: " + shape.Type)
	}
	return
}

func stripStar(name string) string {
	if name[0:1] == "*" {
		return name[1:]
	}
	return name
}

func (ref *ShapeRef) GoTags(a *API) (code string) {
	code = "`"
	if ref.Location != "" {
		code += `location:"` + ref.Location + `" `
	}
	if ref.LocationName != "" {
		code += `locationName:"` + ref.LocationName + `" `
	}
	shape := a.Shapes[ref.ShapeName]
	code += `type:"` + shape.Type + `" `
	if shape.Payload != "" {
		code += `payload:"` + shape.Payload + `" `
	}
	if len(shape.Required) > 0 {
		code += `required:"` + strings.Join(shape.Required, ",") + `" `
	}
	if strings.Contains(a.Metadata.Protocol, "json") {
		code += `json:",omitempty"`
	}

	code = strings.TrimSpace(code)
	code += "`"
	return
}
