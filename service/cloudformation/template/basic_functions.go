package template

import "fmt"

func FnBase64(valueToEncode interface{}) interface{} {
	return struct {
		Encode interface{} `json:"Fn::Base64"`
	}{valueToEncode}
}

func FnFindInMap(mapName string, topLevelKey interface{}, secondLevelKey interface{}) interface{} {
	return struct {
		Keys []interface{} `json:"Fn::FindInMap"`
	}{[]interface{}{mapName, topLevelKey, secondLevelKey}}
}

func FnGetAtt(logicalNameOfResource string, attributeName interface{}) interface{} {
	return struct {
		Att []interface{} `json:"Fn::GetAtt"`
	}{[]interface{}{logicalNameOfResource, attributeName}}
}

func FnGetAZs(region interface{}) interface{} {
	if region == nil {
		region = ""
	}
	return struct {
		AZs interface{} `json:"Fn::GetAZs"`
	}{region}
}

func FnSelect(index interface{}, listOfObjects interface{}) interface{} {
	switch intIndex := index.(type) {
	case int:
		index = fmt.Sprintf("%d", intIndex)
	}
	return struct {
		Select []interface{} `json:"Fn::Select"`
	}{[]interface{}{index, listOfObjects}}
}

func FnJoin(delimeter interface{}, listOfValues ...interface{}) interface{} {
	return struct {
		Join []interface{} `json:"Fn::Join"`
	}{[]interface{}{delimeter, listOfValues}}
}

func Ref(logicalName string) interface{} {
	return struct{ Ref string }{logicalName}
}
