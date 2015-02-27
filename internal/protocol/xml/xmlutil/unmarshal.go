package xmlutil

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func UnmarshalXML(v interface{}, d *xml.Decoder) error {
	n, _ := xmlToStruct(d, nil)
	if n.children != nil {
		for _, root := range n.children {
			for _, c := range root {
				err := parse(reflect.ValueOf(v), c, "")
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	return fmt.Errorf("Missing root XML node")
}

func parse(r reflect.Value, node *xmlNode, tag reflect.StructTag) error {
	t := r.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // check kind of actual element type
	}

	switch t.Kind() {
	case reflect.Struct:
		if field, ok := t.FieldByName("SDKShapeTraits"); ok {
			tag = field.Tag
		}
		return parseStruct(r, node, tag)
	case reflect.Slice:
		if tag.Get("type") == "blob" { // this is a scalar slice, not a list
			return parseScalar(r, node, tag)
		} else {
			return parseList(r, node, tag)
		}
	case reflect.Map:
		return parseMap(r, node, tag)
	default:
		return parseScalar(r, node, tag)
	}
}

func parseStruct(r reflect.Value, node *xmlNode, tag reflect.StructTag) error {
	t := r.Type()
	if r.Kind() == reflect.Ptr {
		if r.IsNil() { // create the structure if it's nil
			s := reflect.New(r.Type().Elem())
			r.Set(s)
			r = s
		}

		r = r.Elem()
		t = t.Elem()
	}

	// unwrap any wrappers
	if wrapper := tag.Get("resultWrapper"); wrapper != "" {
		if children, ok := node.children[wrapper]; ok {
			for _, c := range children {
				err := parseStruct(r, c, "")
				if err != nil {
					return err
				}
			}
			return nil
		}
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if c := field.Name[0:1]; strings.ToLower(c) == c {
			continue // ignore unexported fields
		}

		// figure out what this field is called
		name := field.Name
		if locName := field.Tag.Get("locationName"); locName != "" {
			name = locName
		}

		// try to find the field by name in elements
		elems := node.children[name]

		if elems == nil { // try to find the field in attributes
			for _, a := range node.attributes {
				if name == a.Name.Local {
					// turn this into a text node for de-serializing
					elems = []*xmlNode{&xmlNode{text: a.Value}}
				}
			}
		}

		member := r.FieldByName(field.Name)
		for _, elem := range elems {
			err := parse(member, elem, field.Tag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func parseList(r reflect.Value, node *xmlNode, tag reflect.StructTag) error {
	t := r.Type()

	if tag.Get("flattened") == "" { // look at all item entries
		mname := "member"
		if name := tag.Get("locationNameList"); name != "" {
			mname = name
		}

		if children, ok := node.children[mname]; ok {
			if r.IsNil() {
				r.Set(reflect.MakeSlice(t, len(children), len(children)))
			}

			for i, c := range children {
				err := parse(r.Index(i), c, "")
				if err != nil {
					return err
				}
			}
		}
	} else { // flattened list means this is a single element
		if r.IsNil() {
			r.Set(reflect.MakeSlice(t, 0, 0))
		}

		childR := reflect.Zero(t.Elem())
		r.Set(reflect.Append(r, childR))
		err := parse(r.Index(r.Len()-1), node, "")
		if err != nil {
			return err
		}
	}

	return nil
}

func parseMap(r reflect.Value, node *xmlNode, tag reflect.StructTag) error {
	t := r.Type()
	if r.Kind() == reflect.Ptr {
		t = t.Elem()
		if r.IsNil() {
			r.Set(reflect.New(t))
			r.Elem().Set(reflect.MakeMap(t))
		}

		r = r.Elem()
	}

	if tag.Get("flattened") == "" { // look at all child entries
		for _, entry := range node.children["entry"] {
			parseMapEntry(r, entry, tag)
		}
	} else { // this element is itself an entry
		parseMapEntry(r, node, tag)
	}

	return nil
}

func parseMapEntry(r reflect.Value, node *xmlNode, tag reflect.StructTag) error {
	kname, vname := "key", "value"
	if n := tag.Get("locationNameKey"); n != "" {
		kname = n
	}
	if n := tag.Get("locationNameValue"); n != "" {
		vname = n
	}

	keys, ok := node.children[kname]
	values := node.children[vname]
	if ok {
		for i, key := range keys {
			keyR := reflect.ValueOf(key.text)
			value := values[i]
			valueR := reflect.New(r.Type().Elem()).Elem()

			parse(valueR, value, "")
			r.SetMapIndex(keyR, valueR)
		}
	}
	return nil
}

func parseScalar(r reflect.Value, node *xmlNode, tag reflect.StructTag) error {
	t := r.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.String:
		r.Set(reflect.ValueOf(&node.text))
		return nil
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 { // blob type
			b, err := base64.StdEncoding.DecodeString(node.text)
			if err != nil {
				return err
			}
			r.Set(reflect.ValueOf(b))
			return nil
		}
	case reflect.Bool:
		v, err := strconv.ParseBool(node.text)
		if err != nil {
			return err
		}
		r.Set(reflect.ValueOf(&v))
		return nil
	case reflect.Int64:
		v, err := strconv.ParseInt(node.text, 10, t.Bits())
		if err != nil {
			return err
		}
		r.Set(reflect.ValueOf(&v))
		return nil
	case reflect.Int:
		v, err := strconv.ParseInt(node.text, 10, t.Bits())
		if err != nil {
			return err
		}
		i := int(v)
		r.Set(reflect.ValueOf(&i))
		return nil
	case reflect.Float64:
		v, err := strconv.ParseFloat(node.text, t.Bits())
		if err != nil {
			return err
		}
		r.Set(reflect.ValueOf(&v))
		return nil
	case reflect.Float32:
		v, err := strconv.ParseFloat(node.text, t.Bits())
		if err != nil {
			return err
		}
		f := float32(v)
		r.Set(reflect.ValueOf(&f))
		return nil
		// case reflect.Struct:
		// 	// const ISO8601UTC = "2006-01-02T15:04:05Z"
		// 	// v.Set(name, value.UTC().Format(ISO8601UTC))
	}
	return fmt.Errorf("Unsupported value: %v (%s)", r.Interface(), t.Name())
}
