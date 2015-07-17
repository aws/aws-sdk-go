package s3_test

import (
	"crypto/md5"
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsconv"
	"github.com/aws/aws-sdk-go/internal/test/unit"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

var _ = unit.Imported

func assertMD5(t *testing.T, req *aws.Request) {
	err := req.Build()
	assert.NoError(t, err)

	b, _ := ioutil.ReadAll(req.HTTPRequest.Body)
	out := md5.Sum(b)
	assert.NotEmpty(t, b)
	assert.Equal(t, base64.StdEncoding.EncodeToString(out[:]), req.HTTPRequest.Header.Get("Content-MD5"))
}

func TestMD5InPutBucketCORS(t *testing.T) {
	svc := s3.New(nil)
	req, _ := svc.PutBucketCORSRequest(&s3.PutBucketCORSInput{
		Bucket: awsconv.String("bucketname"),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{
				{AllowedMethods: []*string{awsconv.String("GET")}},
			},
		},
	})
	assertMD5(t, req)
}

func TestMD5InPutBucketLifecycle(t *testing.T) {
	svc := s3.New(nil)
	req, _ := svc.PutBucketLifecycleRequest(&s3.PutBucketLifecycleInput{
		Bucket: awsconv.String("bucketname"),
		LifecycleConfiguration: &s3.LifecycleConfiguration{
			Rules: []*s3.LifecycleRule{
				{
					ID:     awsconv.String("ID"),
					Prefix: awsconv.String("Prefix"),
					Status: awsconv.String("Enabled"),
				},
			},
		},
	})
	assertMD5(t, req)
}

func TestMD5InPutBucketPolicy(t *testing.T) {
	svc := s3.New(nil)
	req, _ := svc.PutBucketPolicyRequest(&s3.PutBucketPolicyInput{
		Bucket: awsconv.String("bucketname"),
		Policy: awsconv.String("{}"),
	})
	assertMD5(t, req)
}

func TestMD5InPutBucketTagging(t *testing.T) {
	svc := s3.New(nil)
	req, _ := svc.PutBucketTaggingRequest(&s3.PutBucketTaggingInput{
		Bucket: awsconv.String("bucketname"),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{Key: awsconv.String("KEY"), Value: awsconv.String("VALUE")},
			},
		},
	})
	assertMD5(t, req)
}

func TestMD5InDeleteObjects(t *testing.T) {
	svc := s3.New(nil)
	req, _ := svc.DeleteObjectsRequest(&s3.DeleteObjectsInput{
		Bucket: awsconv.String("bucketname"),
		Delete: &s3.Delete{
			Objects: []*s3.ObjectIdentifier{
				{Key: awsconv.String("key")},
			},
		},
	})
	assertMD5(t, req)
}
