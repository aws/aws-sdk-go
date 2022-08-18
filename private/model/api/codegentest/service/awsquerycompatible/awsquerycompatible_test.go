//go:build go1.13 && codegen
// +build go1.13,codegen

package awsquerycompatible

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"
)

func TestAWSQuery(t *testing.T) {
	cases := map[string]struct {
		statusCode   int
		responseBody interface {
			io.Reader
			Len() int
		}
		headers         http.Header
		expectErrorCode string
	}{
		"legacy query code": {
			statusCode:      500,
			responseBody:    strings.NewReader(`{"__type":"com.amazonaws.awsquerycompatible#QueueDeletedRecently", "message":"Some user-visible message"}`),
			expectErrorCode: "AWS.SimpleQueueService.QueueDeletedRecently",
		},
		"new error code": {
			statusCode:      500,
			responseBody:    strings.NewReader(`{"__type":"com.amazonaws.awsquerycompatible#QueueNameExists", "message":"Some user-visible message"}`),
			expectErrorCode: "QueueNameExists",
		},
		"mapped unmodlled error code": {
			statusCode:      400,
			responseBody:    strings.NewReader(`{"__type":"com.amazonaws.awsquerycompatible#AccessDeniedException", "message":"Some user-visible message"}`),
			expectErrorCode: "AccessDenied",
		},
		"unmapped unmodelled error code": {
			statusCode:      400,
			responseBody:    strings.NewReader(`{"__type":"com.amazonaws.awsquerycompatible#SomeException", "message":"Some user-visible message"}`),
			expectErrorCode: "SomeException",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			client := New(unit.Session.Copy(), &aws.Config{
				MaxRetries: aws.Int(0),
			})

			client.Handlers.Send.Swap(corehandlers.SendHandler.Name, request.NamedHandler{
				Name: corehandlers.SendHandler.Name,
				Fn: func(r *request.Request) {
					r.HTTPResponse = &http.Response{
						StatusCode: c.statusCode,
						ContentLength: func() int64 {
							if c.responseBody == nil {
								return 0
							}
							return int64(c.responseBody.Len())
						}(),
						Header: func() http.Header {
							if c.headers == nil {
								return http.Header{}
							}
							return c.headers
						}(),
						Body: ioutil.NopCloser(c.responseBody),
					}
				},
			})

			_, err := client.CreateQueue(&CreateQueueInput{
				QueueName: aws.String("queueName"),
			})
			if err == nil {
				t.Fatalf("expect error, got none")
			}

			var awsErr awserr.RequestFailure
			if !errors.As(err, &awsErr) {
				t.Fatalf("expect RequestFailure error, got %#v", err)
			}

			if e, a := c.expectErrorCode, awsErr.Code(); e != a {
				t.Errorf("expect %v code, got %v", e, a)
			}
		})
	}
}

func Test_MappedErrorCode(t *testing.T) {
	if e, a := "AWS.SimpleQueueService.QueueDeletedRecently", ErrCodeQueueDeletedRecently; e != a {
		t.Errorf("expect %v code, got %v", e, a)
	}
}

func Test_UnmappedErrorCode(t *testing.T) {
	if e, a := "QueueNameExists", ErrCodeQueueNameExists; e != a {
		t.Errorf("expect %v code, got %v", e, a)
	}
}


