package helpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/awslabs/aws-sdk-go/internal/model/api"
	"github.com/awslabs/aws-sdk-go/internal/util"
	"github.com/awslabs/aws-sdk-go/internal/util/utilsort"
)

type paramFiller struct {
	prefixPackageName bool
}

func (f paramFiller) typeName(shape *api.Shape) string {
	if f.prefixPackageName && shape.Type == "structure" {
		return "*" + shape.API.PackageName() + "." + shape.GoTypeElem()
	} else {
		return shape.GoType()
	}
}

func ParamsStructFromJSON(value interface{}, shape *api.Shape, prefixPackageName bool) string {
	f := paramFiller{prefixPackageName: prefixPackageName}
	return util.GoFmt(f.paramsStructAny(value, shape))
}

func (f paramFiller) paramsStructAny(value interface{}, shape *api.Shape) string {
	if value == nil {
		return ""
	}

	switch shape.Type {
	case "structure":
		if value != nil {
			vmap := value.(map[string]interface{})
			return f.paramsStructStruct(vmap, shape)
		}
	case "list":
		vlist := value.([]interface{})
		return f.paramsStructList(vlist, shape)
	case "map":
		vmap := value.(map[string]interface{})
		return f.paramsStructMap(vmap, shape)
	case "string", "character":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.String(%#v)", v.Interface())
		}
	case "blob":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("[]byte(%#v)", v.Interface())
		}
	case "boolean":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.Boolean(%#v)", v.Interface())
		}
	case "integer", "long":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.Long(%v)", v.Interface())
		}
	case "float", "double":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.Double(%v)", v.Interface())
		}
	case "timestamp":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.Time(time.Unix(%d, 0))", int(v.Float()))
		}
	default:
		panic("Unhandled type " + shape.Type)
	}
	return ""
}

func (f paramFiller) paramsStructStruct(value map[string]interface{}, shape *api.Shape) string {
	out := "&" + f.typeName(shape)[1:] + "{\n"
	for _, n := range shape.MemberNames() {
		ref := shape.MemberRefs[n]
		name := findMember(value, n)

		if val := f.paramsStructAny(value[name], ref.Shape); val != "" {
			out += fmt.Sprintf("%s: %s,\n", n, val)
		}
	}
	out += "}"
	return out
}

func (f paramFiller) paramsStructMap(value map[string]interface{}, shape *api.Shape) string {
	out := "&" + f.typeName(shape)[1:] + "{\n"
	keys := utilsort.SortedKeys(value)
	for _, k := range keys {
		v := value[k]
		out += fmt.Sprintf("%q: %s,\n", k, f.paramsStructAny(v, shape.ValueRef.Shape))
	}
	out += "}"
	return out
}

func (f paramFiller) paramsStructList(value []interface{}, shape *api.Shape) string {
	out := f.typeName(shape) + "{\n"
	for _, v := range value {
		out += fmt.Sprintf("%s,\n", f.paramsStructAny(v, shape.MemberRef.Shape))
	}
	out += "}"
	return out
}

func findMember(value map[string]interface{}, key string) string {
	for actualKey, _ := range value {
		if strings.ToLower(key) == strings.ToLower(actualKey) {
			return actualKey
		}
	}
	return ""
}
