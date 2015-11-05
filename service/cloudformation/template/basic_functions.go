package template

import "fmt"

// FnBase64 marshals to the CloudFormation intrinsic function Fn::Base64
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-base64.html
func FnBase64(valueToEncode interface{}) interface{} {
	return struct {
		Encode interface{} `json:"Fn::Base64"`
	}{valueToEncode}
}

// FnFindInMap marshals to the CloudFormation intrinsic function Fn::FindInMap
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-findinmap.html
func FnFindInMap(mapName string, topLevelKey interface{}, secondLevelKey interface{}) interface{} {
	return struct {
		Keys []interface{} `json:"Fn::FindInMap"`
	}{[]interface{}{mapName, topLevelKey, secondLevelKey}}
}

// FnGetAtt marshals to the CloudFormation intrinsic function Fn::GetAtt
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html
func FnGetAtt(logicalNameOfResource string, attributeName interface{}) interface{} {
	return struct {
		Att []interface{} `json:"Fn::GetAtt"`
	}{[]interface{}{logicalNameOfResource, attributeName}}
}

// FnGetAZs marshals to the CloudFormation intrinsic function Fn::GetAZs
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getavailabilityzones.html
func FnGetAZs(region interface{}) interface{} {
	if region == nil {
		region = ""
	}
	return struct {
		AZs interface{} `json:"Fn::GetAZs"`
	}{region}
}

// FnJoin marshals to the CloudFormation intrinsic function Fn::Join
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-join.html
func FnJoin(delimeter interface{}, listOfValues ...interface{}) interface{} {
	return struct {
		Join []interface{} `json:"Fn::Join"`
	}{[]interface{}{delimeter, listOfValues}}
}

// FnSelect marshals to the CloudFormation intrinsic function Fn::Select
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-select.html
func FnSelect(index interface{}, listOfObjects interface{}) interface{} {
	switch intIndex := index.(type) {
	case int:
		index = fmt.Sprintf("%d", intIndex)
	}
	return struct {
		Select []interface{} `json:"Fn::Select"`
	}{[]interface{}{index, listOfObjects}}
}

// Ref marshals to the CloudFormation intrinsic function Ref
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html
func Ref(logicalName string) interface{} {
	return struct{ Ref string }{logicalName}
}
