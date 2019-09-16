package request_test

import (
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/mock"
)

func TestRequestCancelRetry(t *testing.T) {

	reqNum := 0
	s := mock.NewMockClient(&aws.Config{
		MaxRetries: aws.Int(1),
	})
	s.Handlers.Validate.Clear()
	s.Handlers.Unmarshal.Clear()
	s.Handlers.UnmarshalMeta.Clear()
	s.Handlers.UnmarshalError.Clear()
	s.Handlers.Send.PushFront(func(r *request.Request) {
		reqNum++
	})
	out := &testData{}
	ctx, cancelFn := context.WithCancel(context.Background())
	r := s.NewRequest(&request.Operation{Name: "Operation"}, nil, out)
	r.SetContext(ctx)
	cancelFn() // cancelling the context associated with the request

	err := r.Send()
	if !strings.Contains(err.Error(), "canceled") {
		t.Errorf("expect canceled in error, %v", err)
	}
	if e, a := 1, reqNum; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
