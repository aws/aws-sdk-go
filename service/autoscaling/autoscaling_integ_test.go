// +build integration

package autoscaling_test

import (
	"os"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/autoscaling"
)

var autoscalingClient *autoscaling.AutoScaling

func TestMain(m *testing.M) {

	autoscalingClient = autoscaling.New(aws.DefaultCreds(), "us-east-1", nil)
	os.Exit(m.Run())
}

func TestDescribeAccountLimits(t *testing.T) {

	output, err := autoscalingClient.DescribeAccountLimits()
	if err != nil {
		t.Fatal("Failed to DescribeAccountLimits: ", err)
	}

	if *(output.MaxNumberOfAutoScalingGroups) == 0 {
		t.Fatal("Failed to marshall MaxNumberOfAutoScalingGroups")
	}
	if *(output.MaxNumberOfLaunchConfigurations) == 0 {
		t.Fatal("Failed to marshall MaxNumberOfAutoScalingGroups")
	}
}
