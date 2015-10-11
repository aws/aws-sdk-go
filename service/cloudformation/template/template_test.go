package template_test

import (
	"encoding/json"
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/aws/aws-sdk-go/service/cloudformation/template"
)

func TestParameterStringify(t *testing.T) {
	param := Parameter{
		Type:      "some-parameter",
		NoEcho:    true,
		MaxLength: 22,
		MinLength: 2,
	}
	assertMarshalsTo(t, param, `{ "Type": "some-parameter", "NoEcho": "true", "MinLength": "2", "MaxLength": "22" }`)
}

func TestFnBase64(t *testing.T) {
	base64 := FnBase64("some string to encode")
	assertMarshalsTo(t, base64, `{ "Fn::Base64" : "some string to encode" }`)
}

func TestFnBase64WithObjects(t *testing.T) {
	base64 := FnBase64(Ref("some-logical-name"))
	assertMarshalsTo(t, base64, `{ "Fn::Base64" : { "Ref": "some-logical-name" } }`)
}

func TestFnFindInMap(t *testing.T) {
	mapped := FnFindInMap("RegionMap", Ref("AWS::Region"), "32")
	assertMarshalsTo(t, mapped, `{ "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "32"]}`)
}

func TestFnGetAtt(t *testing.T) {
	att := FnGetAtt("MyLB", Ref("something"))
	assertMarshalsTo(t, att, `{ "Fn::GetAtt": [ "MyLB" , { "Ref": "something" } ] }`)
}

func TestFnGetAZs(t *testing.T) {
	azs := FnGetAZs(Ref("AWS::Region"))
	assertMarshalsTo(t, azs, `{ "Fn::GetAZs": { "Ref": "AWS::Region" } }`)
}

func TestFnGetAZsEmpty(t *testing.T) {
	azs := FnGetAZs("")
	assertMarshalsTo(t, azs, `{ "Fn::GetAZs": "" }`)
}

func TestFnGetAZsNil(t *testing.T) {
	azs := FnGetAZs(nil)
	assertMarshalsTo(t, azs, `{ "Fn::GetAZs": "" }`)
}

func TestFnJoin(t *testing.T) {
	joined := FnJoin("\n", "some", "list", Ref("AWS::StackId"))
	assertMarshalsTo(t, joined, `{ "Fn::Join" : ["\n", [ "some", "list", { "Ref" : "AWS::StackId" } ] ] }`)
}

func TestFnSelect(t *testing.T) {
	selected := FnSelect(Ref("MyIndex"), FnGetAZs(""))
	assertMarshalsTo(t, selected, `{ "Fn::Select" : [ { "Ref": "MyIndex" }, { "Fn::GetAZs": "" } ] }`)
}

func TestFnSelectWithLiteralIndex(t *testing.T) {
	selected := FnSelect(2, FnGetAZs(""))
	assertMarshalsTo(t, selected, `{ "Fn::Select" : [ "2", { "Fn::GetAZs": "" } ] }`)
}

func TestRefs(t *testing.T) {
	ref := Ref("AWS::StackId")
	assertMarshalsTo(t, ref, `{"Ref":"AWS::StackId" }`)
}

func TestTemplateCreation(t *testing.T) {
	template := Template{

		AWSTemplateFormatVersion: "2010-09-09",

		Description: "This is a test template",

		Metadata: map[string]interface{}{
			"some-key":       map[string]int{"thing": 42},
			"some-other-key": []bool{true, true, false, true},
		},

		Parameters: map[string]Parameter{
			"KeyName": {
				Type:        "AWS::EC2::KeyPair::KeyName",
				Description: "SSH KeyPair to use for instances",
			},
			"DBPassword": {
				NoEcho:                true,
				Description:           "Password for the DB",
				Type:                  "String",
				MinLength:             5,
				MaxLength:             31,
				AllowedPattern:        "[a-zA-Z0-9]*",
				ConstraintDescription: "must contain only alphanumeric characters",
			},
			"InstanceType": {
				Type:          "String",
				Default:       "t1.micro",
				AllowedValues: []string{"t1.micro", "t2.micro"},
			},
		},
	}

	expected := `
{
  "AWSTemplateFormatVersion" : "2010-09-09",

  "Description" : "This is a test template",

    "Metadata": {
        "some-key": { "thing": 42 },
        "some-other-key": [ true, true, false, true ]
    },

  "Parameters": {
    "KeyName": {
      "Description" : "SSH KeyPair to use for instances",
      "Type": "AWS::EC2::KeyPair::KeyName"
    },
    "DBPassword": {
      "NoEcho": "true",
      "Description" : "Password for the DB",
      "Type": "String",
      "MinLength": "5",
      "MaxLength": "31",
      "AllowedPattern" : "[a-zA-Z0-9]*",
      "ConstraintDescription" : "must contain only alphanumeric characters"
    },
    "InstanceType" : {
      "Type" : "String",
      "Default" : "t1.micro",
      "AllowedValues" : [ "t1.micro", "t2.micro" ]
    }
  }
}
`
	assertMarshalsTo(t, template, expected)

	asString := template.String()
	assertEquivalentJSON(t, expected, asString)
}

func assertMarshalsTo(t *testing.T, value interface{}, expectedJSON string) {
	actual, err := json.Marshal(value)
	assert.Nil(t, err)

	assertEquivalentJSON(t, expectedJSON, string(actual))
}

func assertEquivalentJSON(t *testing.T, expected, actual string) {
	var rawExpected, rawActual interface{}

	err := json.Unmarshal([]byte(expected), &rawExpected)
	if err != nil {
		t.Errorf("Unable to unmarshal expected data (%s) as JSON: %s", expected, err)
	}

	err = json.Unmarshal([]byte(actual), &rawActual)
	if err != nil {
		t.Errorf("Unable to unmarshal actual data (%s) as JSON: %s", actual, err)
	}

	assert.EqualValues(t, rawExpected, rawActual, fmt.Sprintf("JSON mismatch: Expected\n\n%s\n\n\t\tbut instead got\n\n%s", expected, actual))
}
