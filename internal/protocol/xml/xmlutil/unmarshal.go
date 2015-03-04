package xmlutil

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func UnmarshalXML(v interface{}, d *xml.Decoder) error {
	n, _ := XMLToStruct(d, nil)
	if n.Children != nil {
		for _, root := range n.Children {
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

func parse(r reflect.Value, node *XMLNode, tag reflect.StructTag) error {
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

func parseStruct(r reflect.Value, node *XMLNode, tag reflect.StructTag) error {
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
		if Children, ok := node.Children[wrapper]; ok {
			for _, c := range Children {
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
		elems := node.Children[name]

		if elems == nil { // try to find the field in attributes
			for _, a := range node.Attr {
				if name == a.Name.Local {
					// turn this into a text node for de-serializing
					elems = []*XMLNode{&XMLNode{Text: a.Value}}
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

func parseList(r reflect.Value, node *XMLNode, tag reflect.StructTag) error {
	t := r.Type()

	if tag.Get("flattened") == "" { // look at all item entries
		mname := "member"
		if name := tag.Get("locationNameList"); name != "" {
			mname = name
		}

		if Children, ok := node.Children[mname]; ok {
			if r.IsNil() {
				r.Set(reflect.MakeSlice(t, len(Children), len(Children)))
			}

			for i, c := range Children {
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

func parseMap(r reflect.Value, node *XMLNode, tag reflect.StructTag) error {
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
		for _, entry := range node.Children["entry"] {
			parseMapEntry(r, entry, tag)
		}
	} else { // this element is itself an entry
		parseMapEntry(r, node, tag)
	}

	return nil
}

func parseMapEntry(r reflect.Value, node *XMLNode, tag reflect.StructTag) error {
	kname, vname := "key", "value"
	if n := tag.Get("locationNameKey"); n != "" {
		kname = n
	}
	if n := tag.Get("locationNameValue"); n != "" {
		vname = n
	}

	keys, ok := node.Children[kname]
	values := node.Children[vname]
	if ok {
		for i, key := range keys {
			keyR := reflect.ValueOf(key.Text)
			value := values[i]
			valueR := reflect.New(r.Type().Elem()).Elem()

			parse(valueR, value, "")
			r.SetMapIndex(keyR, valueR)
		}
	}
	return nil
}

func parseScalar(r reflect.Value, node *XMLNode, tag reflect.StructTag) error {
	switch r.Interface().(type) {
	case *string:
		r.Set(reflect.ValueOf(&node.Text))
		return nil
	case []byte:
		b, err := base64.StdEncoding.DecodeString(node.Text)
		if err != nil {
			return err
		}
		r.Set(reflect.ValueOf(b))
	case *bool:
		v, err := strconv.ParseBool(node.Text)
		if err != nil {
			return err
		}
		r.Set(reflect.ValueOf(&v))
	case *int64:
		v, err := strconv.ParseInt(node.Text, 10, 64)
		if err != nil {
			return err
		}
		r.Set(reflect.ValueOf(&v))
	case *int:
		v, err := strconv.ParseInt(node.Text, 10, 32)
		if err != nil {
			return err
		}
		i := int(v)
		r.Set(reflect.ValueOf(&i))
	case *float64:
		v, err := strconv.ParseFloat(node.Text, 64)
		if err != nil {
			return err
		}
		r.Set(reflect.ValueOf(&v))
	case *float32:
		v, err := strconv.ParseFloat(node.Text, 32)
		if err != nil {
			return err
		}
		f := float32(v)
		r.Set(reflect.ValueOf(&f))
	case *time.Time:
		const ISO8601UTC = "2006-01-02T15:04:05Z"
		t, err := time.Parse(ISO8601UTC, node.Text)
		if err != nil {
			return err
		} else {
			r.Set(reflect.ValueOf(&t))
		}
	default:
		return fmt.Errorf("Unsupported value: %v (%s)", r.Interface(), r.Type())
	}
	return nil
}
