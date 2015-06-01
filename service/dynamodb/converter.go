package dynamodb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

// ConvertTo accepts a map or struct converts it to the map[string]*AttributeValue
// type used to interact with the DynamoDB Item APIs.
//
// If in is a struct, we first JSON encode/decode it to get the data as a map.
// This can/should be optimized later.
func ConvertTo(in interface{}) (item map[string]*AttributeValue, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(runtime.Error); ok {
				err = e
			} else if s, ok := r.(string); ok {
				err = errors.New(s)
			} else {
				err = r.(error)
			}
			item = nil
		}
	}()

	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Struct:
		item = convertToStruct(in)
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			return nil, errors.New("item must be a map[string]interface{} or struct (or a non-nil pointer to one), got " + v.Type().String())
		}
		item = convertToMap(in)
	case reflect.Ptr:
		if v.IsNil() {
			return nil, errors.New("item must not be nil")
		}
		return ConvertTo(v.Elem().Interface())
	default:
		return nil, errors.New("item must be a map[string]interface{} or struct (or a non-nil pointer to one), got " + v.Type().String())
	}
	return item, nil
}

// ConvertFrom accepts the map[string]*AttributeValue type returned by the
// DynamoDB Item APIs and converts it to a map or struct.
//
// If v points to a struct, we first convert it to a basic map, then JSON
// encode/decode it to convert to a struct. This can/should be optimized later.
func ConvertFrom(item map[string]*AttributeValue, v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(runtime.Error); ok {
				err = e
			} else if s, ok := r.(string); ok {
				err = errors.New(s)
			} else {
				err = r.(error)
			}
			item = nil
		}
	}()

	m := make(map[string]interface{})
	for k, v := range item {
		m[k] = convertFrom(v)
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("v must be a non-nil pointer to a map[string]interface{} or struct, got " + rv.Type().String())
	}

	switch rv.Elem().Kind() {
	case reflect.Struct:
		// TODO: We convert basic maps into structs by JSON encoding/decoding.
		// This can be made more efficient by simplying recursing over the
		// struct and populating it from the map.
		b, err := json.Marshal(m)
		if err != nil {
			return err
		}
		return json.Unmarshal(b, v)
	case reflect.Map:
		if rv.Elem().Type().Key().Kind() != reflect.String {
			return errors.New("v must be a non-nil pointer to a map[string]interface{} or struct, got " + rv.Type().String())
		}
		rv.Elem().Set(reflect.ValueOf(m))
	default:
		return errors.New("v must be a non-nil pointer to a map[string]interface{} or struct, got " + rv.Type().String())
	}

	return nil
}

func convertToStruct(in interface{}) map[string]*AttributeValue {
	// TODO: We convert structs into basic maps by JSON encoding/decoding. This
	// can be made more efficient by recursing over the struct directly.
	b, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m)
	if err != nil {
		panic(err)
	}

	return convertToMap(m)
}

func convertToMap(in interface{}) map[string]*AttributeValue {
	item := make(map[string]*AttributeValue)
	m := in.(map[string]interface{})
	for k, v := range m {
		item[k] = convertTo(v)
	}
	return item
}

func convertTo(in interface{}) *AttributeValue {
	a := &AttributeValue{}

	if in == nil {
		a.NULL = new(bool)
		*a.NULL = true
		return a
	}

	if m, ok := in.(map[string]interface{}); ok {
		mp := make(map[string]*AttributeValue)
		for k, v := range m {
			mp[k] = convertTo(v)
		}
		a.M = &mp
		return a
	}

	if l, ok := in.([]interface{}); ok {
		a.L = make([]*AttributeValue, len(l))
		for index, v := range l {
			a.L[index] = convertTo(v)
		}
		return a
	}

	// Only primitive types should remain.
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Bool:
		a.BOOL = new(bool)
		*a.BOOL = v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a.N = new(string)
		*a.N = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		a.N = new(string)
		*a.N = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		a.N = new(string)
		*a.N = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.String:
		if n, ok := in.(json.Number); ok {
			a.N = new(string)
			*a.N = n.String()
		} else {
			a.S = new(string)
			*a.S = v.String()
		}
	default:
		panic(fmt.Sprintf(`the type %s is not supported`, v.Type().String()))
	}

	return a
}

func convertFrom(a *AttributeValue) interface{} {
	if a.S != nil {
		return *a.S
	}

	if a.N != nil {
		// Number is tricky b/c we don't know which numeric type to use. Here we
		// simply try the different types from most to least restrictive.
		if n, err := strconv.ParseInt(*a.N, 10, 64); err == nil {
			return int(n)
		}
		if n, err := strconv.ParseUint(*a.N, 10, 64); err == nil {
			return uint(n)
		}
		n, err := strconv.ParseFloat(*a.N, 64)
		if err != nil {
			panic(err)
		}
		return n
	}

	if a.BOOL != nil {
		return *a.BOOL
	}

	if a.NULL != nil {
		return nil
	}

	if a.M != nil {
		m := make(map[string]interface{})
		for k, v := range *a.M {
			m[k] = convertFrom(v)
		}
		return m
	}

	if a.L != nil {
		l := make([]interface{}, len(a.L))
		for index, v := range a.L {
			l[index] = convertFrom(v)
		}
		return l
	}

	panic(fmt.Sprintf("unsupported dynamo attribute %#v", a))
}
