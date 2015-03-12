// +build integration

package elastictranscoder_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/util/utilassert"
	"github.com/awslabs/aws-sdk-go/service/elastictranscoder"
	"github.com/stretchr/testify/assert"
)

var (
	_ = assert.Equal
	_ = utilassert.Match
)

func TestMakingABasicRequest(t *testing.T) {
	client := elastictranscoder.New(nil)
	resp, e := client.ListPresets(&elastictranscoder.ListPresetsInput{})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NoError(t, nil, err)

}

func TestErrorHandling(t *testing.T) {
	client := elastictranscoder.New(nil)
	resp, e := client.ReadJob(&elastictranscoder.ReadJobInput{
		ID: aws.String("fake_job"),
	})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "ValidationException", err.Code)
	utilassert.Match(t, "Value 'fake_job' at 'id' failed to satisfy constraint", err.Message)

}
