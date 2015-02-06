// +build integration

package codedeploy_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/codedeploy"
)

func TestCreateDeleteApplication(t *testing.T) {

	client := codedeploy.New(aws.DefaultCreds(), "us-east-1", nil)
	applicationName := "awsgosdk-" + strconv.FormatInt(time.Now().UnixNano(), 10)
	createOutput, err := client.CreateApplication(&codedeploy.CreateApplicationInput{
		ApplicationName: &applicationName,
	})
	if err != nil {
		t.Fatal("Failed to create application: ", err)
	}
	if *(createOutput.ApplicationID) == "" {
		t.Fatal("Failed to marshall create response")
	}

	client.DeleteApplication(&codedeploy.DeleteApplicationInput{
		ApplicationName: &applicationName,
	})
	if err != nil {
		t.Fatal("Failed to delete application: ", err)
	}
}
