package query

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

func Build(r *aws.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.Service.APIVersion},
	}
	if err := parseValue(body, r.Params, "", ""); err != nil {
		r.Error = err
		return
	}

	r.HTTPRequest.Method = "POST"
	r.HTTPRequest.Body = ioutil.NopCloser(bytes.NewReader([]byte(body.Encode())))
}

func elemOf(value reflect.Value) reflect.Value {
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}

func parseValue(v url.Values, i interface{}, prefix string, tag reflect.StructTag) error {
	value := elemOf(reflect.ValueOf(i))

	// no need to handle zero values
	if !value.IsValid() {
		return nil
	}

	switch value.Kind() {
	case reflect.Struct:
		return parseStruct(v, value, prefix)
	case reflect.Slice:
		if tag.Get("type") == "blob" { // this is a scalar slice, not a list
			return parseScalar(v, value, prefix, tag)
		} else {
			return parseList(v, value, prefix, tag)
		}
	case reflect.Map:
		return parseMap(v, value, prefix, tag)
	default:
		return parseScalar(v, value, prefix, tag)
	}
}

func parseStruct(v url.Values, value reflect.Value, prefix string) error {
	if !value.IsValid() {
		return nil
	}

	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		if c := t.Field(i).Name[0:1]; strings.ToLower(c) == c {
			continue // ignore unexported fields
		}

		value := elemOf(value.Field(i))
		field := t.Field(i)
		name := field.Tag.Get("locationName")
		if name == "" {
			name = field.Name
		}
		if prefix != "" {
			name = prefix + "." + name
		}

		if err := parseValue(v, value.Interface(), name, field.Tag); err != nil {
			return err
		}
	}
	return nil
}

func parseList(v url.Values, value reflect.Value, prefix string, tag reflect.StructTag) error {
	// check for unflattened list member
	if tag.Get("flattened") == "" {
		prefix += ".member"
	}

	for i := 0; i < value.Len(); i++ {
		slicePrefix := prefix
		if slicePrefix == "" {
			slicePrefix = strconv.Itoa(i + 1)
		} else {
			slicePrefix = slicePrefix + "." + strconv.Itoa(i+1)
		}
		if err := parseValue(v, value.Index(i).Interface(), slicePrefix, ""); err != nil {
			return err
		}
	}
	return nil
}

func parseMap(v url.Values, value reflect.Value, prefix string, tag reflect.StructTag) error {
	// check for unflattened list member
	if tag.Get("flattened") == "" {
		prefix += ".entry"
	}

	for i, mapKey := range value.MapKeys() {
		mapValue := value.MapIndex(mapKey)

		// serialize key
		var keyName string
		if prefix == "" {
			keyName = strconv.Itoa(i+1) + ".key"
		} else {
			keyName = prefix + "." + strconv.Itoa(i+1) + ".key"
		}

		if err := parseValue(v, mapKey.Interface(), keyName, ""); err != nil {
			return err
		}

		// serialize value
		var valueName string
		if prefix == "" {
			valueName = strconv.Itoa(i+1) + ".value"
		} else {
			valueName = prefix + "." + strconv.Itoa(i+1) + ".value"
		}

		if err := parseValue(v, mapValue.Interface(), valueName, ""); err != nil {
			return err
		}
	}

	return nil
}

func parseScalar(v url.Values, r reflect.Value, name string, tag reflect.StructTag) error {
	switch value := r.Interface().(type) {
	case string:
		v.Set(name, value)
	case []byte:
		v.Set(name, base64.StdEncoding.EncodeToString(value))
	case bool:
		v.Set(name, strconv.FormatBool(value))
	case int64:
		v.Set(name, strconv.FormatInt(value, 10))
	case int:
		v.Set(name, strconv.Itoa(value))
	case float64:
		v.Set(name, strconv.FormatFloat(value, 'f', -1, 64))
	case float32:
		v.Set(name, strconv.FormatFloat(float64(value), 'f', -1, 32))
	case time.Time:
		const ISO8601UTC = "2006-01-02T15:04:05Z"
		v.Set(name, value.UTC().Format(ISO8601UTC))
	default:
		return fmt.Errorf("Unsupported value for param %s: %v (%s)", name, r.Interface(), r.Type().Name())
	}
	return nil
}
