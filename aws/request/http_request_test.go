// +build go1.5

package request_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting"
)

func TestRequestCancel(t *testing.T) {
	c := make(chan struct{})

	s := awstesting.NewMockClient()
	s.Handlers.Validate.Clear()
	out := &testData{}
	r := s.NewRequest(&request.Operation{Name: "Operation"}, nil, out)
	r.HTTPRequest.Cancel = c
	close(c)

	err := r.Send()
	assert.True(t, strings.Contains(err.Error(), "canceled"))
}

func TestRequestCancelFailAfterSend(t *testing.T) {
	c := make(chan struct{})

	s := awstesting.NewMockClient()
	s.Handlers.Validate.Clear()
	out := &testData{}
	r := s.NewRequest(&request.Operation{Name: "Operation"}, nil, out)
	r.HTTPRequest.Cancel = c

	err := r.Send()
	close(c)
	assert.Nil(t, err)
}

func TestRequestCancelRetry(t *testing.T) {
	c := make(chan struct{})

	reqNum := 0
	s := awstesting.NewMockClient(aws.NewConfig().WithMaxRetries(10))
	s.Handlers.Validate.Clear()
	s.Handlers.Unmarshal.PushBack(unmarshal)
	s.Handlers.UnmarshalError.PushBack(unmarshalError)
	s.Handlers.Send.PushBack(func(r *request.Request) {
		reqNum++
	})
	out := &testData{}
	r := s.NewRequest(&request.Operation{Name: "Operation"}, nil, out)
	r.HTTPRequest.Cancel = c
	close(c)

	err := r.Send()
	assert.True(t, strings.Contains(err.Error(), "canceled"))
	assert.Equal(t, 1, reqNum)
}
