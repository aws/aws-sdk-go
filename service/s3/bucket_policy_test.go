package s3_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
)

const getBucketPolicyRespJSON = `{"Version":"2012-10-17","Id":"Policy1234567890123","Statement":[{"Sid":"Stmt1234567890123","Effect":"Allow","Principal":{"AWS":"arn:aws:iam::123456789012:user/UserName"},"Action":"s3:*","Resource":"arn:aws:s3:::bucket-name"}]}`

func TestUnmarshalGetBucketPolicy(t *testing.T) {
	svc := s3.New(nil)
	req, output := svc.GetBucketPolicyRequest(&s3.GetBucketPolicyInput{Bucket: aws.String("bucket")})

	req.Handlers.Send.Clear()
	req.Handlers.Send.PushBack(func(r *aws.Request) {
		req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(getBucketPolicyRespJSON)))}
	})

	req.Build()
	req.Send()

	assert.NoError(t, req.Error, "Expect no error")
	assert.Equal(t, getBucketPolicyRespJSON, *output.Policy, "Expected policy JSON to match")
}
