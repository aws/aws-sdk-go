package aws

import (
	"encoding/xml"
	"reflect"
	"strings"
)

// MarshalXML is a weird and stunted version of xml.Marshal which is used by the
// REST-XML request types to get around a bug in encoding/xml which doesn't
// allow us to marshal pointers to zero values:
//
// https://github.com/golang/go/issues/5452
func MarshalXML(v interface{}, e *xml.Encoder, start xml.StartElement) error {
	value := reflect.ValueOf(v)
	t := value.Type()
	switch value.Kind() {
	case reflect.Ptr:
		if !value.IsNil() {
			return MarshalXML(value.Elem().Interface(), e, start)
		}
	case reflect.Struct:
		var rootInfo xmlFieldInfo

		// detect xml.Name, if any
		for i := 0; i < value.NumField(); i++ {
			f := t.Field(i)
			if f.Type == xmlName {
				rootInfo = parseXMLTag(f.Tag.Get("xml"))
			}
		}

		if err := e.EncodeToken(rootInfo.start(t.Name())); err != nil {
			return err
		}

		for i := 0; i < value.NumField(); i++ {
			ft := value.Type().Field(i)

			if ft.Type == xmlName {
				continue
			}

			fv := value.Field(i)
			fi := parseXMLTag(ft.Tag.Get("xml"))

			if fi.omit {
				switch fv.Kind() {
				case reflect.Ptr:
					if fv.IsNil() {
						continue
					}
				default:
					if !fv.IsValid() {
						continue
					}
				}
			}

			start := fi.start(ft.Name)
			if err := e.EncodeElement(fv.Interface(), start); err != nil {
				return err
			}
		}

		if err := e.EncodeToken(rootInfo.end(t.Name())); err != nil {
			return err
		}
	default:
		return e.Encode(v)
	}
	return nil
}

var xmlName = reflect.TypeOf(xml.Name{})

type xmlFieldInfo struct {
	name string
	ns   string
	omit bool
}

func (fi xmlFieldInfo) start(name string) xml.StartElement {
	if fi.name != "" {
		name = fi.name
	}
	return xml.StartElement{
		Name: xml.Name{
			Local: name,
			Space: fi.ns,
		},
	}
}

func (fi xmlFieldInfo) end(name string) xml.EndElement {
	if fi.name != "" {
		name = fi.name
	}
	return xml.EndElement{
		Name: xml.Name{
			Local: name,
			Space: fi.ns,
		},
	}
}

func parseXMLTag(t string) xmlFieldInfo {
	parts := strings.Split(t, ",")

	var omit bool
	for _, p := range parts {
		omit = omit || p == "omitempty"
	}

	var name, ns string
	if len(parts) > 0 {
		nameParts := strings.Split(parts[0], " ")
		if len(nameParts) == 2 {
			name = nameParts[1]
			ns = nameParts[0]
		} else if len(nameParts) == 1 {
			name = nameParts[0]
		}

	}

	return xmlFieldInfo{
		name: name,
		ns:   ns,
		omit: omit,
	}
}
