package query

import (
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

func Build(r *aws.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.Service.APIVersion},
	}
	if err := loadValues(body, r.Params, ""); err != nil {
		r.Error = err
		return
	}

	r.HTTPRequest.Method = "GET"
	r.HTTPRequest.URL.RawQuery = body.Encode()
}

func loadValues(v url.Values, i interface{}, prefix string) error {
	value := reflect.ValueOf(i)

	// follow any pointers
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// no need to handle zero values
	if !value.IsValid() {
		return nil
	}

	switch value.Kind() {
	case reflect.Struct:
		return loadStruct(v, value, prefix)
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			slicePrefix := prefix
			if slicePrefix == "" {
				slicePrefix = strconv.Itoa(i + 1)
			} else {
				slicePrefix = slicePrefix + "." + strconv.Itoa(i+1)
			}
			if err := loadValues(v, value.Index(i).Interface(), slicePrefix); err != nil {
				return err
			}
		}
		return nil
	case reflect.Map:
		sortedKeys := []string{}
		keysByString := map[string]reflect.Value{}
		for _, k := range value.MapKeys() {
			s := fmt.Sprintf("%v", k.Interface())
			sortedKeys = append(sortedKeys, s)
			keysByString[s] = k
		}
		sort.Strings(sortedKeys)

		for i, sortKey := range sortedKeys {
			mapKey := keysByString[sortKey]

			var keyName string
			if prefix == "" {
				keyName = strconv.Itoa(i+1) + ".Name"
			} else {
				keyName = prefix + "." + strconv.Itoa(i+1) + ".Name"
			}

			if err := loadValue(v, mapKey, keyName); err != nil {
				return err
			}

			mapValue := value.MapIndex(mapKey)

			var valueName string
			if prefix == "" {
				valueName = strconv.Itoa(i+1) + ".Value"
			} else {
				valueName = prefix + "." + strconv.Itoa(i+1) + ".Value"
			}

			if err := loadValue(v, mapValue, valueName); err != nil {
				return err
			}
		}

		return nil
	default:
		panic("unknown request member type: " + value.String())
	}
}

func loadStruct(v url.Values, value reflect.Value, prefix string) error {
	if !value.IsValid() {
		return nil
	}

	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		value := value.Field(i)
		name := t.Field(i).Tag.Get("name")
		if name == "" {
			name = t.Field(i).Name
		}
		if prefix != "" {
			name = prefix + "." + name
		}
		if err := loadValue(v, value, name); err != nil {
			return err
		}
	}
	return nil
}

func loadValue(v url.Values, value reflect.Value, name string) error {
	switch casted := value.Interface().(type) {
	case string:
		if casted != "" {
			v.Set(name, casted)
		}
	case *string:
		if casted != nil {
			v.Set(name, *casted)
		}
	case *bool:
		if casted != nil {
			v.Set(name, strconv.FormatBool(*casted))
		}
	case *int64:
		if casted != nil {
			v.Set(name, strconv.FormatInt(*casted, 10))
		}
	case *int:
		if casted != nil {
			v.Set(name, strconv.Itoa(*casted))
		}
	case *float64:
		if casted != nil {
			v.Set(name, strconv.FormatFloat(*casted, 'f', -1, 64))
		}
	case *float32:
		if casted != nil {
			v.Set(name, strconv.FormatFloat(float64(*casted), 'f', -1, 32))
		}
	case time.Time:
		if !casted.IsZero() {
			const ISO8601UTC = "2006-01-02T15:04:05Z"
			v.Set(name, casted.UTC().Format(ISO8601UTC))
		}
	case []string:
		if len(casted) != 0 {
			for i, val := range casted {
				v.Set(fmt.Sprintf("%s.%d", name, i+1), val)
			}
		}
	default:
		if err := loadValues(v, value.Interface(), name); err != nil {
			return err
		}
	}
	return nil
}
