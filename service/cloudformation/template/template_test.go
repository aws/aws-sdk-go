package template_test

import (
	"encoding/json"

	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/awstesting"
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
	awstesting.AssertJSON(t, expected, asString)
}

func TestResourcePropertiesNotRequired(t *testing.T) {
	resource := Resource{Type: "AWS::EC2::InternetGateway"}
	expected := `{ "Type": "AWS::EC2::InternetGateway" }`

	assertMarshalsTo(t, resource, expected)
}

func assertMarshalsTo(t *testing.T, value interface{}, expectedJSON string) {
	actual, err := json.Marshal(value)
	assert.Nil(t, err)

	awstesting.AssertJSON(t, expectedJSON, string(actual))
}
