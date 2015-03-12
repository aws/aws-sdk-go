package util

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go/format"
	"io"
	"reflect"
	"regexp"
	"strings"

	"github.com/awslabs/aws-sdk-go/internal/protocol/xml/xmlutil"
)

func GoFmt(buf string) string {
	formatted, err := format.Source([]byte(buf))
	if err != nil {
		panic(fmt.Errorf("%s\nOriginal code:\n%s", err.Error(), buf))
	}
	return string(formatted)
}

var reTrim = regexp.MustCompile(`\s{2,}`)

func Trim(s string) string {
	return strings.TrimSpace(reTrim.ReplaceAllString(s, " "))
}

func Capitalize(s string) string {
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

func SortXML(r io.Reader) string {
	var buf bytes.Buffer
	d := xml.NewDecoder(r)
	root, _ := xmlutil.XMLToStruct(d, nil)
	e := xml.NewEncoder(&buf)
	xmlutil.StructToXML(e, root, true)
	return buf.String()
}

func PrettyPrint(v interface{}) string {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Struct:
		str := fullName(value.Type()) + "{\n"
		for i := 0; i < value.NumField(); i++ {
			l := string(value.Type().Field(i).Name[0])
			if strings.ToUpper(l) == l {
				str += value.Type().Field(i).Name + ": "
				str += PrettyPrint(value.Field(i).Interface())
				str += ",\n"
			}
		}
		str += "}"
		return str
	case reflect.Map:
		str := "map[" + fullName(value.Type().Key()) + "]" + fullName(value.Type().Elem()) + "{\n"
		for _, k := range value.MapKeys() {
			str += "\"" + k.String() + "\": "
			str += PrettyPrint(value.MapIndex(k).Interface())
			str += ",\n"
		}
		str += "}"
		return str
	case reflect.Ptr:
		if e := value.Elem(); e.IsValid() {
			return "&" + PrettyPrint(e.Interface())
		} else {
			return "nil"
		}
	case reflect.Slice:
		str := "[]" + fullName(value.Type().Elem()) + "{\n"
		for i := 0; i < value.Len(); i++ {
			str += PrettyPrint(value.Index(i).Interface())
			str += ",\n"
		}
		str += "}"
		return str
	default:
		return fmt.Sprintf("%#v", v)
	}
	return ""
}

func pkgName(t reflect.Type) string {
	pkg := t.PkgPath()
	c := strings.Split(pkg, "/")
	return c[len(c)-1]
}

func fullName(t reflect.Type) string {
	if pkg := pkgName(t); pkg != "" {
		return pkg + "." + t.Name()
	}
	return t.Name()
}
