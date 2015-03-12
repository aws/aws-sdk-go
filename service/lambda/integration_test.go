// +build integration

package lambda_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/util/utilassert"
	"github.com/awslabs/aws-sdk-go/service/lambda"
	"github.com/stretchr/testify/assert"
)

var (
	_ = assert.Equal
	_ = utilassert.Match
)

func TestMakingABasicRequest(t *testing.T) {
	client := lambda.New(nil)
	resp, e := client.ListEventSources(&lambda.ListEventSourcesInput{})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NoError(t, nil, err)

}

func TestErrorHandling(t *testing.T) {
	client := lambda.New(nil)
	resp, e := client.GetEventSource(&lambda.GetEventSourceInput{
		UUID: aws.String("fake-uuid"),
	})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "ResourceNotFoundException", err.Code)
	utilassert.Match(t, "does not exist", err.Message)

}
