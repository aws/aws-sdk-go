package aws

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	Data string
}

var delays = []time.Duration{}
var sleepDelay = func(delay time.Duration) {
	delays = append(delays, delay)
}

func body(str string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(str)))
}

func TestRequestRecoverRetry(t *testing.T) {
	reqNum := 0
	reqs := []http.Response{
		http.Response{StatusCode: 500, Body: body(`{"__type":"UnknownError","message":"An error occurred."}`)},
		http.Response{StatusCode: 500, Body: body(`{"__type":"UnknownError","message":"An error occurred."}`)},
		http.Response{StatusCode: 200, Body: body(`{"data":"valid"}`)},
	}

	s := aws.NewService(&aws.Config{MaxRetries: -1})
	s.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	s.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)
	s.Handlers.Send.Init() // mock sending
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		r.HTTPResponse = &reqs[reqNum]
		reqNum++
	})
	out := &testData{}
	r := aws.NewRequest(s, &aws.Operation{Name: "Operation"}, nil, out)
	err := r.Send()
	assert.Nil(t, err)
	assert.Equal(t, 2, int(r.RetryCount))
	assert.Equal(t, "valid", out.Data)
}

func TestRequestExhaustRetries(t *testing.T) {
	reqNum := 0
	reqs := []http.Response{
		http.Response{StatusCode: 500, Body: body(`{"__type":"UnknownError","message":"An error occurred."}`)},
		http.Response{StatusCode: 500, Body: body(`{"__type":"UnknownError","message":"An error occurred."}`)},
		http.Response{StatusCode: 500, Body: body(`{"__type":"UnknownError","message":"An error occurred."}`)},
	}

	s := aws.NewService(&aws.Config{MaxRetries: -1})
	s.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	s.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)
	s.Handlers.Send.Init() // mock sending
	s.Handlers.Send.PushBack(func(r *aws.Request) {
		r.HTTPResponse = &reqs[reqNum]
		reqNum++
	})
	r := aws.NewRequest(s, &aws.Operation{Name: "Operation"}, nil, nil)
	err := r.Send()
	apiErr := aws.Error(err)
	assert.NotNil(t, err)
	assert.NotNil(t, apiErr)
	assert.Equal(t, 500, apiErr.StatusCode)
	assert.Equal(t, "UnknownError", apiErr.Code)
	assert.Equal(t, "An error occurred.", apiErr.Message)
	assert.Equal(t, 3, int(r.RetryCount))
}
