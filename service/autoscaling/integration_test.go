// +build integration

package autoscaling_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/util/utilassert"
	"github.com/awslabs/aws-sdk-go/service/autoscaling"
	"github.com/stretchr/testify/assert"
)

var (
	_ = assert.Equal
	_ = utilassert.Match
)

func TestMakingABasicRequest(t *testing.T) {
	client := autoscaling.New(nil)
	resp, e := client.DescribeScalingProcessTypes(&autoscaling.DescribeScalingProcessTypesInput{})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NoError(t, nil, err)

}

func TestErrorHandling(t *testing.T) {
	client := autoscaling.New(nil)
	resp, e := client.CreateLaunchConfiguration(&autoscaling.CreateLaunchConfigurationInput{
		ImageID:                 aws.String("ami-12345678"),
		InstanceType:            aws.String("m1.small"),
		LaunchConfigurationName: aws.String(""),
	})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "ValidationError", err.Code)
	utilassert.Match(t, "Member must have length greater than or equal to 1", err.Message)

}
