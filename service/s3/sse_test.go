package s3_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/awscfg"
	"github.com/aws/aws-sdk-go/aws/awsconv"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/internal/test/unit"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

var _ = unit.Imported

func TestSSECustomerKeyOverHTTPError(t *testing.T) {
	s := s3.New(&awscfg.Config{DisableSSL: awsconv.Bool(true)})
	req, _ := s.CopyObjectRequest(&s3.CopyObjectInput{
		Bucket:         awsconv.String("bucket"),
		CopySource:     awsconv.String("bucket/source"),
		Key:            awsconv.String("dest"),
		SSECustomerKey: awsconv.String("key"),
	})
	err := req.Build()

	assert.Error(t, err)
	assert.Equal(t, "ConfigError", err.(awserr.Error).Code())
	assert.Contains(t, err.(awserr.Error).Message(), "cannot send SSE keys over HTTP")
}

func TestCopySourceSSECustomerKeyOverHTTPError(t *testing.T) {
	s := s3.New(&awscfg.Config{DisableSSL: awsconv.Bool(true)})
	req, _ := s.CopyObjectRequest(&s3.CopyObjectInput{
		Bucket:     awsconv.String("bucket"),
		CopySource: awsconv.String("bucket/source"),
		Key:        awsconv.String("dest"),
		CopySourceSSECustomerKey: awsconv.String("key"),
	})
	err := req.Build()

	assert.Error(t, err)
	assert.Equal(t, "ConfigError", err.(awserr.Error).Code())
	assert.Contains(t, err.(awserr.Error).Message(), "cannot send SSE keys over HTTP")
}

func TestComputeSSEKeys(t *testing.T) {
	s := s3.New(nil)
	req, _ := s.CopyObjectRequest(&s3.CopyObjectInput{
		Bucket:                   awsconv.String("bucket"),
		CopySource:               awsconv.String("bucket/source"),
		Key:                      awsconv.String("dest"),
		SSECustomerKey:           awsconv.String("key"),
		CopySourceSSECustomerKey: awsconv.String("key"),
	})
	err := req.Build()

	assert.NoError(t, err)
	assert.Equal(t, "a2V5", req.HTTPRequest.Header.Get("x-amz-server-side-encryption-customer-key"))
	assert.Equal(t, "a2V5", req.HTTPRequest.Header.Get("x-amz-copy-source-server-side-encryption-customer-key"))
	assert.Equal(t, "PG4LipwVIkqCKLmpjKFTHQ==", req.HTTPRequest.Header.Get("x-amz-server-side-encryption-customer-key-md5"))
	assert.Equal(t, "PG4LipwVIkqCKLmpjKFTHQ==", req.HTTPRequest.Header.Get("x-amz-copy-source-server-side-encryption-customer-key-md5"))
}

func TestComputeSSEKeysShortcircuit(t *testing.T) {
	s := s3.New(nil)
	req, _ := s.CopyObjectRequest(&s3.CopyObjectInput{
		Bucket:                      awsconv.String("bucket"),
		CopySource:                  awsconv.String("bucket/source"),
		Key:                         awsconv.String("dest"),
		SSECustomerKey:              awsconv.String("key"),
		CopySourceSSECustomerKey:    awsconv.String("key"),
		SSECustomerKeyMD5:           awsconv.String("MD5"),
		CopySourceSSECustomerKeyMD5: awsconv.String("MD5"),
	})
	err := req.Build()

	assert.NoError(t, err)
	assert.Equal(t, "a2V5", req.HTTPRequest.Header.Get("x-amz-server-side-encryption-customer-key"))
	assert.Equal(t, "a2V5", req.HTTPRequest.Header.Get("x-amz-copy-source-server-side-encryption-customer-key"))
	assert.Equal(t, "MD5", req.HTTPRequest.Header.Get("x-amz-server-side-encryption-customer-key-md5"))
	assert.Equal(t, "MD5", req.HTTPRequest.Header.Get("x-amz-copy-source-server-side-encryption-customer-key-md5"))
}
