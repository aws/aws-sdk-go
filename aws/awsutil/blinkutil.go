package awsutil

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"reflect"
	"sort"
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
		"map",
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

func underlayParameter(parameter string, shape reflect.Value) string {
	shapeKind := getShapeKind(shape)

	if shapeKind == reflect.Struct || shapeKind == reflect.Map {
		structPrefixIndex := strings.IndexByte(parameter, '{')
		structSuffixIndex := strings.IndexByte(parameter, '}')

		if structPrefixIndex == 0 && structSuffixIndex == len(parameter)-1 {
			parameter = parameter[1:structSuffixIndex]
		}
	} else if shapeKind == reflect.Slice {
		slicePrefixIndex := strings.IndexByte(parameter, '[')
		sliceSuffixIndex := strings.LastIndexByte(parameter, ']')
		if slicePrefixIndex == 0 && sliceSuffixIndex == len(parameter)-1 {
			parameter = parameter[1:sliceSuffixIndex]
		}
	}

	return parameter
}

func resolveMapParameter(parent string, parameter string, shape reflect.Value) map[string]string {
	unpacked := make(map[string]string)

	var concatItems []string
	lastItemIndex := -1
	mapIndex := -1
	sliceIndex := -1
	for _, keyValue := range strings.SplitAfter(parameter, ",") {
		keyValue = strings.ReplaceAll(keyValue, ",", "")
		if sliceIndex == -1 && mapIndex == -1 && strings.Contains(keyValue, "=") {
			concatItems = append(concatItems, keyValue)
			if mapIndex == -1 && sliceIndex == -1 {
				lastItemIndex++
			}
		} else {
			if concatItems == nil || 0 > lastItemIndex || lastItemIndex >= len(concatItems) {
				continue
			}
			concatItems[lastItemIndex] += fmt.Sprintf("%s%s", ",", keyValue)
		}

		if strings.Contains(keyValue, "[") {
			sliceIndex += 1
		}
		if strings.Contains(keyValue, "{") {
			mapIndex += 1
		}

		if strings.Contains(keyValue, "}") {
			mapIndex -= 1
		}
		if strings.Contains(keyValue, "]") {
			sliceIndex -= 1
		}
	}

	for _, keyValue := range concatItems {
		itemKeyValue := strings.Split(keyValue, "=")
		key, value := itemKeyValue[0], itemKeyValue[1]
		if len(itemKeyValue) > 2 {
			value = strings.Join(itemKeyValue[1:], "=")
		}
		unpacked[key] = value
	}

	return unpacked
}

func correctShape(shape reflect.Value) reflect.Value {
	if !shape.IsValid() {
		return reflect.ValueOf(nil)
	}

	shapeKind := shape.Kind()
	if shapeKind == reflect.Invalid {
		return reflect.ValueOf(nil)
	}

	for shapeKind == reflect.Ptr || shapeKind == reflect.UnsafePointer || shapeKind == reflect.Uintptr {
		shapeKind = shape.Type().Elem().Kind()
		shape = reflect.Indirect(reflect.New(shape.Type().Elem()))
	}

	return shape
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

func getSlices(parameter string) []string {
	spaceIndexes := getAllIndexes(parameter, ' ')
	delimiterIndexes := getAllIndexes(parameter, ',')
	slicePrefixIndexes := getAllIndexes(parameter, '[')
	sliceSuffixIndexes := getAllIndexes(parameter, ']')
	mapPrefixIndexes := getAllIndexes(parameter, '{')
	mapSuffixIndexes := getAllIndexes(parameter, '}')

	realDelimiterIndexes := make([]int, 0)
	if len(slicePrefixIndexes) == 0 && len(sliceSuffixIndexes) == 0 && len(mapPrefixIndexes) == 0 && len(mapSuffixIndexes) == 0 {
		if spaceIndexes != nil {
			realDelimiterIndexes = spaceIndexes
		} else {
			realDelimiterIndexes = delimiterIndexes
		}
	} else {
		for _, index := range delimiterIndexes {
			betweenSlices := betweenPrefixAndSuffix(index, slicePrefixIndexes, sliceSuffixIndexes)
			betweenMaps := betweenPrefixAndSuffix(index, mapPrefixIndexes, mapSuffixIndexes)
			if betweenSlices {
				realDelimiterIndexes = append(realDelimiterIndexes, index)
			} else if betweenMaps {
				realDelimiterIndexes = append(realDelimiterIndexes, index)
			}
		}
	}

	slices := make([]string, 0)
	lastIndex := 0
	for _, delimiterIndex := range realDelimiterIndexes {
		delimiterIndex -= 1
		slices = append(slices, parameter[lastIndex:delimiterIndex])
		lastIndex = delimiterIndex + 1
	}
	slices = append(slices, parameter[lastIndex:])

	return slices
}

func unpackSlice(parent string, parameter string, shape reflect.Value) []interface{} {
	shapeKind := getSliceKind(shape)
	if shapeKind == reflect.Invalid {
		return nil
	}

	slicedShape := getSliceShape(shape)
	slices := getSlices(parameter)

	resolvedList := make([]interface{}, 0)
	for _, item := range slices {
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
	case reflect.Map:
		return resolveMapParameter(parent, parameter, shape)
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

	shape = correctShape(shape)
	parameter = underlayParameter(parameter, shape)
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

func GetServiceRegions(serviceName string) []string {
	awsPartition := endpoints.AwsPartition()
	operationServiceRegions := awsPartition.Regions()
	services := awsPartition.Services()

	if operationService, ok := services[strings.ToLower(serviceName)]; ok {
		operationServiceRegions = operationService.Regions()
	}

	operationRegions := make([]string, 0)
	for _, region := range operationServiceRegions {
		operationRegions = append(operationRegions, region.ID())
	}

	return operationRegions
}

func GetActionTypeFromShape(shape interface{}) string {
	shapeValue := reflect.ValueOf(shape)
	shapeValue = correctShape(shapeValue)
	shapeKind := getShapeKind(shapeValue)

	return getActionTypeFromShapeKind(shapeValue, shapeKind)
}

func getActionTypeFromShapeKind(shape reflect.Value, shapeKind reflect.Kind) string {
	switch shapeKind {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return "int"
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "uint"
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	case reflect.Map:
		return "code:map"
	case reflect.Struct:
		return "code:json"
	case reflect.Slice:
		sliceKind := getSliceKind(shape)
		slicedShape := getSliceShape(shape)
		return getActionTypeFromShapeKind(slicedShape, sliceKind)
	}

	return "string"
}

func getAllIndexes(str string, c byte) []int {
	var indexes []int

	lastIndex := 0
	for len(str) > 0 {
		index := strings.IndexByte(str, c) + 1
		if index == 0 {
			break
		}
		indexes = append(indexes, lastIndex+index)
		str = str[index:]
		lastIndex += index
	}

	sort.Ints(indexes)
	return indexes
}

func betweenPrefixAndSuffix(index int, prefixIndexes []int, suffixIndexes []int) bool {
	nearPrefix, nearSuffix := false, false

	for _, prefix := range prefixIndexes {
		if prefix == index+1 {
			nearPrefix = true
		}
	}

	for _, suffix := range suffixIndexes {
		if suffix == index-1 {
			nearSuffix = true
		}
	}

	return nearPrefix && nearSuffix
}
