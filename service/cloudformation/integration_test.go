// +build integration

package cloudformation_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/util/utilassert"
	"github.com/awslabs/aws-sdk-go/service/cloudformation"
	"github.com/stretchr/testify/assert"
)

var (
	_ = assert.Equal
	_ = utilassert.Match
)

func TestMakingABasicRequest(t *testing.T) {
	client := cloudformation.New(nil)
	resp, e := client.ListStacks(&cloudformation.ListStacksInput{})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NoError(t, nil, err)

}

func TestErrorHandling(t *testing.T) {
	client := cloudformation.New(nil)
	resp, e := client.CreateStack(&cloudformation.CreateStackInput{
		StackName:   aws.String("fakestack"),
		TemplateURL: aws.String("http://s3.amazonaws.com/foo/bar"),
	})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "ValidationError", err.Code)
	utilassert.Match(t, "TemplateURL must reference a valid S3 object to which you have access.", err.Message)

}
