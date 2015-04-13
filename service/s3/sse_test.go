package s3_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestSSECustomerKeyOverHTTPError(t *testing.T) {
	s := s3.New(baseConfig.Merge(&aws.Config{DisableSSL: true}))
	req, _ := s.CopyObjectRequest(&s3.CopyObjectInput{
		Bucket:         aws.String("bucket"),
		CopySource:     aws.String("bucket/source"),
		Key:            aws.String("dest"),
		SSECustomerKey: aws.String("key"),
	})
	err := req.Build()

	assert.Error(t, err)
	aerr := aws.Error(err)
	assert.Equal(t, "ConfigError", aerr.Code)
	assert.Contains(t, aerr.Message, "cannot send SSE keys over HTTP")
}

func TestCopySourceSSECustomerKeyOverHTTPError(t *testing.T) {
	s := s3.New(baseConfig.Merge(&aws.Config{DisableSSL: true}))
	req, _ := s.CopyObjectRequest(&s3.CopyObjectInput{
		Bucket:     aws.String("bucket"),
		CopySource: aws.String("bucket/source"),
		Key:        aws.String("dest"),
		CopySourceSSECustomerKey: aws.String("key"),
	})
	err := req.Build()

	assert.Error(t, err)
	aerr := aws.Error(err)
	assert.Equal(t, "ConfigError", aerr.Code)
	assert.Contains(t, aerr.Message, "cannot send SSE keys over HTTP")
}
