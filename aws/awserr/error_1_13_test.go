//go:build go1.13
// +build go1.13

package awserr_test

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

type temporaryError interface {
	error
	Temporary() bool
}

func TestNew(t *testing.T) {
	err := awserr.New("RequestTimeout", "RequestTimeout", context.DeadlineExceeded)
	if err.Code() != "RequestTimeout" {
		t.Errorf("expected RequestTimeout, but received %v", err.Code())
	}
	if err.Message() != "RequestTimeout" {
		t.Errorf("expected RequestTimeout, but received %v", err.Message())
	}
	if !errors.Is(err.OrigErr(), context.DeadlineExceeded) {
		t.Error("expected err is a context.DeadlineExceeded")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Error("expected err is a context.DeadlineExceeded")
	}

	var e temporaryError
	if !errors.As(err, &e) {
		t.Errorf("expacted err can be casted to temporaryError")
	}
}

func TestNewBatchError(t *testing.T) {
	errs := []error{context.DeadlineExceeded}
	err := awserr.NewBatchError("RequestTimeout", "RequestTimeout", errs)
	if err.Code() != "RequestTimeout" {
		t.Errorf("expected RequestTimeout, but received %v", err.Code())
	}
	if !reflect.DeepEqual(err.OrigErrs(), errs) {
		t.Errorf("expected values to be equivalent but received %v and %v", err.OrigErrs(), errs)
	}
	if !errors.Is(err.OrigErr(), context.DeadlineExceeded) {
		t.Error("expected err is a context.DeadlineExceeded")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Error("expected err is a context.DeadlineExceeded")
	}

	var e temporaryError
	if !errors.As(err, &e) {
		t.Fatal("expected err can be casted to temporaryError")
	}
	if !e.Temporary() {
		t.Error("expect true")
	}
}

func TestNewRequestFailure(t *testing.T) {
	reqID := "123"
	err := awserr.NewRequestFailure(
		awserr.New("RequestTimeout", "RequestTimeout", context.DeadlineExceeded), http.StatusBadRequest, reqID)
	if err.Code() != "RequestTimeout" {
		t.Errorf("expected RequestTimeout but received %v", err.Code())
	}
	if err.Message() != "RequestTimeout" {
		t.Errorf("expected RequestTimeout but received %v", err.Message())
	}
	if err.RequestID() != reqID {
		t.Errorf("expected values to be equivalent but received %v and %v", err.RequestID(), reqID)
	}
	if err.StatusCode() != http.StatusBadRequest {
		t.Errorf("expected values to be equivalent but received %v and %v", err.StatusCode(), http.StatusBadRequest)
	}
	if !errors.Is(err.OrigErr(), context.DeadlineExceeded) {
		t.Errorf("expected values to be equivalent but received %v and %v", err.OrigErr(), context.DeadlineExceeded)
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Error("expected err is a context.DeadlineExceeded")
	}

	var e temporaryError
	if !errors.As(err, &e) {
		t.Fatal("expected err can be casted to temporaryError")
	}
	if !e.Temporary() {
		t.Error("expect true")
	}
}

func TestNewUnmarshalError(t *testing.T) {
	err := awserr.NewUnmarshalError(context.DeadlineExceeded, "unexpected EOF", []byte("unexpected EOF"))
	if err.Code() != "UnmarshalError" {
		t.Errorf("expected UnmarshalError but received %v", err.Code())
	}
	if err.Message() != "unexpected EOF" {
		t.Errorf("expected 'unexpected EOF' but received %v", err.Message())
	}
	if !errors.Is(err.OrigErr(), context.DeadlineExceeded) {
		t.Errorf("expected context.DeadlineExceeded but received %v", err.OrigErr())
	}
	if string(err.Bytes()) != "unexpected EOF" {
		t.Errorf("expected 'unexpected EOF' but received %s", err.Bytes())
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Error("expected err is a context.DeadlineExceeded")
	}

	var e temporaryError
	if !errors.As(err, &e) {
		t.Fatal("expected err can be casted to temporaryError")
	}
	if !e.Temporary() {
		t.Error("expect true")
	}
}
