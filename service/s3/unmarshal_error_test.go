package s3_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/s3"
)

type testErrorCase struct {
	RespFn    func() *http.Response
	ReqID     string
	Code, Msg string
}

var testUnmarshalCases = []testErrorCase{
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode:    301,
				Header:        http.Header{"X-Amz-Request-Id": []string{"abc123"}},
				Body:          ioutil.NopCloser(nil),
				ContentLength: -1,
			}
		},
		ReqID: "abc123",
		Code:  "BucketRegionError", Msg: "incorrect region, the bucket is not in 'mock-region' region",
	},
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode:    403,
				Header:        http.Header{"X-Amz-Request-Id": []string{"abc123"}},
				Body:          ioutil.NopCloser(nil),
				ContentLength: 0,
			}
		},
		ReqID: "abc123",
		Code:  "Forbidden", Msg: "Forbidden",
	},
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode:    400,
				Header:        http.Header{"X-Amz-Request-Id": []string{"abc123"}},
				Body:          ioutil.NopCloser(nil),
				ContentLength: 0,
			}
		},
		ReqID: "abc123",
		Code:  "BadRequest", Msg: "Bad Request",
	},
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode:    404,
				Header:        http.Header{"X-Amz-Request-Id": []string{"abc123"}},
				Body:          ioutil.NopCloser(nil),
				ContentLength: 0,
			}
		},
		ReqID: "abc123",
		Code:  "NotFound", Msg: "Not Found",
	},
	{
		RespFn: func() *http.Response {
			body := `<Error><Code>SomeException</Code><Message>Exception message</Message></Error>`
			return &http.Response{
				StatusCode:    500,
				Header:        http.Header{"X-Amz-Request-Id": []string{"abc123"}},
				Body:          ioutil.NopCloser(strings.NewReader(body)),
				ContentLength: int64(len(body)),
			}
		},
		ReqID: "abc123",
		Code:  "SomeException", Msg: "Exception message",
	},
}

func TestUnmarshalError(t *testing.T) {
	for _, c := range testUnmarshalCases {
		s := s3.New(unit.Session)
		s.Handlers.Send.Clear()
		s.Handlers.Send.PushBack(func(r *request.Request) {
			r.HTTPResponse = c.RespFn()
			r.HTTPResponse.Status = http.StatusText(r.HTTPResponse.StatusCode)
		})
		_, err := s.PutBucketAcl(&s3.PutBucketAclInput{
			Bucket: aws.String("bucket"), ACL: aws.String("public-read"),
		})

		fmt.Printf("%#v\n", err)

		assert.Error(t, err)
		assert.Equal(t, c.Code, err.(awserr.Error).Code())
		assert.Equal(t, c.Msg, err.(awserr.Error).Message())
		assert.Equal(t, c.ReqID, err.(awserr.RequestFailure).RequestID())
	}
}
