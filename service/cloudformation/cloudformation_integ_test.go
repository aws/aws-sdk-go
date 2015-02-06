// +build integration

package cloudformation_test

import (
	"os"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/cloudformation"
)

var cloudformationClient *cloudformation.CloudFormation
var exampleTemplate string

func TestMain(m *testing.M) {

	exampleTemplate =
		`{
	"AWSTemplateFormatVersion" : "2010-09-09",

	"Description" : "This is a sample template",

	"Parameters" : {
		"TopicName" : {
		    "Type" : "String",
            "Default" : "TheTopic",
            "Description" : "A topic."
		},
        "QueueName" : {
            "Type" : "String",
            "Default" : "TheQueue",
            "Description" : "A queue."
        }
	},

	"Resources" : {

		"TheQueue" : {
		    "Type" : "AWS::SQS::Queue",
		    "Properties" : {
				"QueueName" : { "Ref" : "QueueName" }
		    }
		},

        "TheTopic" : {
            "Type" : "AWS::SNS::Topic",
            "Properties" : {
				"TopicName" : { "Ref" : "TopicName" },
                "Subscription" : [
					{"Protocol" : "sqs", "Endpoint" : {"Fn::GetAtt" : [ "TheQueue", "Arn"]}}
                ]
            }
        }
	},

	"Outputs" : {
		"TopicARN" : {
		    "Value" : { "Ref" : "TheTopic" }
		},
        "QueueURL" : {
            "Value" : { "Ref" : "TheQueue" }
        }
	}
}`

	cloudformationClient = cloudformation.New(aws.DefaultCreds(), "us-east-1", nil)
	os.Exit(m.Run())
}

func TestGetTemplateSummary(t *testing.T) {

	output, err := cloudformationClient.GetTemplateSummary(&cloudformation.GetTemplateSummaryInput{
		TemplateBody: &exampleTemplate,
	})
	if err != nil {
		t.Fatal("Failed to call GetTemplateSummaryInput: ", err)
	}

	if len(output.Parameters) != 2 {
		t.Fatal("Incorrect number of paraters")
	}

	if *(output.Parameters[0].ParameterKey) != "TopicName" {
		t.Fatal("Incorrect parameter name: ", *(output.Parameters[0].ParameterKey))
	}
	if *(output.Parameters[0].ParameterType) != "String" {
		t.Fatal("Incorrect parameter type: ", *(output.Parameters[0].ParameterType))
	}
	if *(output.Parameters[0].DefaultValue) != "TheTopic" {
		t.Fatal("Incorrect parameter default: ", *(output.Parameters[0].DefaultValue))
	}

}
