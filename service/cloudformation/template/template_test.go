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

	actual, err := json.Marshal(param)
	assert.Nil(t, err)

	expected := `{ "Type": "some-parameter", "NoEcho": "true", "MinLength": "2", "MaxLength": "22" }`

	assertEquivalentJSON(t, expected, string(actual))
}

func TestRefs(t *testing.T) {
	ref := Ref("AWS::StackId")

	actual, err := json.Marshal(ref)
	assert.Nil(t, err)

	assertEquivalentJSON(t, `{"Ref":"AWS::StackId" }`, string(actual))
}

func TestFnBase64(t *testing.T) {
	base64 := FnBase64("some string to encode")

	actual, err := json.Marshal(base64)
	assert.Nil(t, err)

	assertEquivalentJSON(t, `{ "Fn::Base64" : "some string to encode" }`, string(actual))
}

func TestFnBase64WithObjects(t *testing.T) {
	base64 := FnBase64(Ref("some-logical-name"))

	actual, err := json.Marshal(base64)
	assert.Nil(t, err)

	assertEquivalentJSON(t, `{ "Fn::Base64" : { "Ref": "some-logical-name" } }`, string(actual))
}

func TestFnFindInMap(t *testing.T) {
	mapped := FnFindInMap("RegionMap", Ref("AWS::Region"), "32")

	actual, err := json.Marshal(mapped)
	assert.Nil(t, err)

	assertEquivalentJSON(t, `{ "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "32"]}`, string(actual))
}

func TestFnJoin(t *testing.T) {
	joined := FnJoin("\n", "some", "list", Ref("AWS::StackId"))

	actual, err := json.Marshal(joined)
	assert.Nil(t, err)

	assertEquivalentJSON(t, `{ "Fn::Join" : ["\n", [ "some", "list", { "Ref" : "AWS::StackId" } ] ] }`, string(actual))
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

	actual := template.String()

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
	assertEquivalentJSON(t, expected, actual)
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
