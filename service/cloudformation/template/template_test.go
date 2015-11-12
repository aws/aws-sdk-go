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

func TestResourceAttributeDependsOn(t *testing.T) {
	resource := Resource{
		Type: "AWS::EC2::Route",
		Properties: map[string]interface{}{
			"DestinationCidrBlock": "0.0.0.0/0",
			"InstanceId":           Ref("NATInstance"),
			"RouteTableId":         Ref("PrivateRouteTable"),
		},
		DependsOn: "NATInstance",
	}
	expected := `{
                "Type": "AWS::EC2::Route",
                "Properties": {
                    "InstanceId": { "Ref": "NATInstance" },
                    "DestinationCidrBlock": "0.0.0.0/0",
                    "RouteTableId": { "Ref": "PrivateRouteTable" }
                },
                "DependsOn": "NATInstance"
            }`

	assertMarshalsTo(t, resource, expected)

	resource.DependsOn = []string{"one", "two"}
	expected = `{
                "Type": "AWS::EC2::Route",
                "Properties": {
                    "InstanceId": { "Ref": "NATInstance" },
                    "DestinationCidrBlock": "0.0.0.0/0",
                    "RouteTableId": { "Ref": "PrivateRouteTable" }
                },
                "DependsOn": [ "one", "two" ]
            }`

	assertMarshalsTo(t, resource, expected)
}

func TestResourceAttributeCreationAndUpdatePolicy(t *testing.T) {
	resource := Resource{
		Type: "AWS::AutoScaling::AutoScalingGroup",
		Properties: map[string]interface{}{
			"AvailabilityZones":       FnGetAZs(""),
			"LaunchConfigurationName": Ref("LaunchConfig"),
			"DesiredCapacity":         "3",
			"MinSize":                 "1",
			"MaxSize":                 "4",
		},
		CreationPolicy: map[string]interface{}{
			"ResourceSignal": map[string]string{
				"Count":   "3",
				"Timeout": "PT15M",
			},
		},
		UpdatePolicy: map[string]interface{}{
			"AutoScalingScheduledAction": map[string]string{
				"IgnoreUnmodifiedGroupSizeProperties": "true",
			},
			"AutoScalingRollingUpdate": map[string]string{
				"MinInstancesInService": "1",
				"MaxBatchSize":          "2",
				"PauseTime":             "PT1M",
				"WaitOnResourceSignals": "true",
			},
		},
	}
	expected := `
{
  "Type": "AWS::AutoScaling::AutoScalingGroup",
  "Properties": {
    "AvailabilityZones": { "Fn::GetAZs": "" },
    "LaunchConfigurationName": { "Ref": "LaunchConfig" },
    "DesiredCapacity": "3",
    "MinSize": "1",
    "MaxSize": "4"
  },
  "CreationPolicy": {
    "ResourceSignal": {
      "Count": "3",
      "Timeout": "PT15M"
    }
  },
  "UpdatePolicy" : {
    "AutoScalingScheduledAction" : {
      "IgnoreUnmodifiedGroupSizeProperties" : "true"
    },
    "AutoScalingRollingUpdate" : {
      "MinInstancesInService" : "1",
      "MaxBatchSize" : "2",
      "PauseTime" : "PT1M",
      "WaitOnResourceSignals" : "true"
    }
  }
}`

	assertMarshalsTo(t, resource, expected)
}

func TestResourceAttributeDeletionPolicy(t *testing.T) {
	resource := Resource{
		Type:           "AWS::S3::Bucket",
		DeletionPolicy: "Retain",
		Properties: map[string]interface{}{
			"BucketName": "some-bucket",
		},
	}
	expected := ` {
  "Type" : "AWS::S3::Bucket",
  "DeletionPolicy" : "Retain",
  "Properties": { "BucketName": "some-bucket" }
}`

	assertMarshalsTo(t, resource, expected)
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
