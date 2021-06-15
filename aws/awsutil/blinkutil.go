package awsutil

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	scalarTypes = []string{
		"bool",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"float32",
		"float64",
		"string",
	}
	complexTypes = []string{
		"slice",
		"struct",
	}
)

func resolveParameter(parent string, parameter string, shape reflect.Value) interface{} {
	if strings.Contains(parameter, " ") {
		return resolveListMapParameter(parent, strings.Split(parameter, " "), shape)
	}

	if strings.Contains(parameter, ",") && !strings.Contains(parameter, "=") {
		return strings.Split(parameter, ",")
	}

	return parameter
}

func resolveListMapParameter(parent string, parameters []string, shape reflect.Value) []map[string]string {
	parameterListMap := make([]map[string]string, 0)
	for _, parameter := range parameters {
		parameterListMap = append(parameterListMap, resolveMapParameter(parent, parameter, shape))
	}

	return parameterListMap
}

func resolveMapParameter(parent string, parameter string, shape reflect.Value) map[string]string {
	unpacked := make(map[string]string)

	var concatItems []string
	lastItemIndex := -1
	for _, keyValue := range strings.SplitAfter(parameter, ",") {
		keyValue = strings.ReplaceAll(keyValue, ",", "")
		if strings.Contains(keyValue, "=") {
			concatItems = append(concatItems, keyValue)
			lastItemIndex++
		} else {
			if concatItems == nil || 0 > lastItemIndex || lastItemIndex >= len(concatItems) {
				continue
			}
			concatItems[lastItemIndex] += fmt.Sprintf("%s%s", ",", keyValue)
		}
	}

	for _, keyValue := range concatItems {
		itemKeyValue := strings.Split(keyValue, "=")
		key, value := itemKeyValue[0], itemKeyValue[1]
		unpacked[key] = value
	}
	return unpacked
}

func getShapeKind(shape reflect.Value) reflect.Kind {
	if !shape.IsValid() {
		return reflect.Invalid
	}

	shapeKind := shape.Kind()
	if shapeKind == reflect.Invalid {
		return reflect.Invalid
	}

	shapeType := shape.Type()
	for shapeKind == reflect.Ptr || shapeKind == reflect.UnsafePointer || shapeKind == reflect.Uintptr {
		shapeKind = shapeType.Elem().Kind()
		shapeType = shapeType.Elem()
	}

	return shapeKind
}

func getSliceKind(shape reflect.Value) reflect.Kind {
	shapeKind := shape.Kind()
	if shapeKind != reflect.Slice {
		return reflect.Invalid
	}
	shapeType := shape.Type().Elem()
	for shapeKind == reflect.Slice || shapeKind == reflect.Ptr || shapeKind == reflect.UnsafePointer || shapeKind == reflect.Uintptr {
		shapeKind = shapeType.Elem().Kind()
		shapeType = shapeType.Elem()
	}

	return shapeKind
}

func getSliceShape(shape reflect.Value) reflect.Value {
	shapeKind := shape.Kind()
	if shapeKind != reflect.Slice {
		return reflect.ValueOf(nil)
	}
	shapeType := shape.Type().Elem()
	for shapeKind == reflect.Slice || shapeKind == reflect.Ptr || shapeKind == reflect.UnsafePointer || shapeKind == reflect.Uintptr {
		shapeKind = shapeType.Elem().Kind()
		shapeType = shapeType.Elem()
	}

	return reflect.Indirect(reflect.New(shapeType))
}

func isScalar(shapeType reflect.Kind) bool {
	for _, scalarType := range scalarTypes {
		if strings.EqualFold(scalarType, shapeType.String()) {
			return true
		}
	}
	return false
}

func isComplex(shapeType reflect.Kind) bool {
	for _, complexType := range complexTypes {
		if strings.EqualFold(complexType, shapeType.String()) {
			return true
		}
	}
	return false
}

func unpackStruct(parameters map[string]string, shape reflect.Value) map[string]interface{} {
	unpacked := make(map[string]interface{})
	for key, value := range parameters {
		structField := shape.FieldByName(key)
		shapeKind := getShapeKind(structField)

		unpacked[key] = UnpackParameter(key, value, structField, shapeKind)
	}
	return unpacked
}

func unpackSlice(parent string, parameter string, shape reflect.Value) []interface{} {
	shapeKind := getSliceKind(shape)
	if shapeKind == reflect.Invalid {
		return nil
	}
	slicedShape := getSliceShape(shape)

	resolvedList := make([]interface{}, 0)
	for _, item := range strings.Split(parameter, " ") {
		unpackedParameter := UnpackParameter(parent, item, slicedShape, shapeKind)
		switch unpackedParameter.(type) {
		case []string:
			for _, unpackedItem := range unpackedParameter.([]string) {
				resolvedList = append(resolvedList, unpackedItem)
			}
		default:
			resolvedList = append(resolvedList, unpackedParameter)
		}
	}
	return resolvedList
}

func unpackScalar(parent string, parameter string, shape reflect.Value, scalarType reflect.Kind) interface{} {
	unpacked := interface{}(parameter)
	switch scalarType {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		unpacked, _ = strconv.ParseInt(parameter, 10, 64)
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		unpacked, _ = strconv.ParseUint(parameter, 10, 64)
	case reflect.Float32, reflect.Float64:
		unpacked, _ = strconv.ParseFloat(parameter, 64)
	case reflect.Bool:
		unpacked, _ = strconv.ParseBool(parameter)
	case reflect.String:
		unpacked = resolveParameter(parent, parameter, shape)
	}
	return unpacked
}

func unpackComplex(parent string, parameter string, shape reflect.Value) interface{} {
	complexKind := getShapeKind(shape)
	switch complexKind {
	case reflect.Invalid:
		return parameter
	case reflect.Struct:
		return unpackStruct(resolveMapParameter(parent, parameter, shape), shape)
	case reflect.Slice:
		return unpackSlice(parent, parameter, shape)
	}
	return nil
}

func UnpackParameter(parent string, parameter string, shape reflect.Value, shapeKind reflect.Kind) interface{} {
	if shapeKind == reflect.Invalid {
		return parameter
	}
	if isScalar(shapeKind) {
		return unpackScalar(parent, parameter, shape, shapeKind)
	}
	if isComplex(shapeKind) {
		return unpackComplex(parent, parameter, shape)
	}
	return parameter
}

func unpackParameters(parameters map[string]string, shape interface{}) map[string]interface{} {
	shapeValue := reflect.ValueOf(shape)
	shapeKind := getShapeKind(shapeValue)

	unpackedParameters := make(map[string]interface{})
	for key, value := range parameters {
		unpackedParameters[key] = value
	}
	if shapeKind == reflect.Invalid {
		return unpackedParameters
	}

	if shapeKind == reflect.Struct {
		return unpackStruct(parameters, shapeValue)
	}

	return unpackedParameters
}


func UnpackParameters(parameters map[string]interface{}, shape interface{}) map[string]interface{} {
	parametersMap := make(map[string]string)
	for key, value := range parameters {
		switch value.(type) {
		case string:
			parametersMap[key] = value.(string)
		default:
			parametersMap[key] = fmt.Sprintf("%v", value)
		}
	}

	return unpackParameters(parametersMap, shape)
}