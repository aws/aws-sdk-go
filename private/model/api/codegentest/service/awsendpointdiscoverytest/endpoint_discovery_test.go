package awsendpointdiscoverytest

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
)

func TestEndpointDiscovery(t *testing.T) {
	cfg := &aws.Config{
		Region:                  aws.String("mock-region"),
		Credentials:             credentials.NewStaticCredentials("AKID", "SECRET", "SESSION"),
		EnableEndpointDiscovery: aws.Bool(true),
	}
	sess := session.New(cfg)
	sess.Handlers = removeHandlers(sess.Handlers, true)

	sess.Handlers.Send.PushBack(func(r *request.Request) {
		out, ok := r.Data.(*DescribeEndpointsOutput)
		if !ok {
			return
		}

		out.Endpoints = []*Endpoint{
			{
				Address: aws.String("foo"),
			},
		}
		r.Data = out
	})

	svc := New(sess)
	svc.Handlers = removeHandlers(svc.Handlers, false)

	req, _ := svc.TestDiscoveryIdentifiersRequiredRequest(&TestDiscoveryIdentifiersRequiredInput{
		Sdk: aws.String("sdk"),
	})

	req.Handlers = removeHandlers(req.Handlers, false)

	req.Handlers.Send.PushBack(func(r *request.Request) {
		if e, a := "foo", r.HTTPRequest.URL.Host; e != a {
			t.Errorf("expected %q, but received %q", e, a)
		}
	})

	if err := req.Send(); err != nil {
		t.Fatal(err)
	}
}

func TestAsyncEndpointDiscovery(t *testing.T) {
	cfg := &aws.Config{
		Region:                  aws.String("mock-region"),
		Credentials:             credentials.NewStaticCredentials("AKID", "SECRET", "SESSION"),
		EnableEndpointDiscovery: aws.Bool(true),
	}
	sess := session.New(cfg)
	sess.Handlers = removeHandlers(sess.Handlers, true)

	sess.Handlers.Send.PushBack(func(r *request.Request) {
		out, ok := r.Data.(*DescribeEndpointsOutput)
		if !ok {
			return
		}

		out.Endpoints = []*Endpoint{
			{
				Address: aws.String("foo"),
			},
		}
		r.Data = out
	})

	svc := New(sess)
	svc.Handlers = removeHandlers(svc.Handlers, false)

	req, _ := svc.TestDiscoveryOptionalRequest(&TestDiscoveryOptionalInput{
		Sdk: aws.String("sdk"),
	})

	req.Handlers = removeHandlers(req.Handlers, false)

	const expectedURI = "awsendpointdiscoverytestservice.mock-region.amazonaws.com"
	req.Handlers.Send.PushBack(func(r *request.Request) {
		if e, a := expectedURI, r.HTTPRequest.URL.Host; e != a {
			t.Errorf("expected %q, but received %q", e, a)
		}
	})

	if err := req.Send(); err != nil {
		t.Fatal(err)
	}

	// prompt the go scheduler to run the goroutine to
	// discover the new endpoint
	time.Sleep(1 * time.Millisecond)

	req, _ = svc.TestDiscoveryOptionalRequest(&TestDiscoveryOptionalInput{
		Sdk: aws.String("sdk"),
	})

	req.Handlers = removeHandlers(req.Handlers, false)
	req.Handlers.Send.PushBack(func(r *request.Request) {
		if e, a := "foo", r.HTTPRequest.URL.Host; e != a {
			t.Errorf("expected %q, but received %q", e, a)
		}
	})

	if err := req.Send(); err != nil {
		t.Fatal(err)
	}
}

func removeHandlers(h request.Handlers, removeSendHandlers bool) request.Handlers {
	if removeSendHandlers {
		h.Send.Clear()
	}
	h.Unmarshal.Clear()
	h.UnmarshalStream.Clear()
	h.UnmarshalMeta.Clear()
	h.UnmarshalError.Clear()
	h.Validate.Clear()
	h.Complete.Clear()
	h.ValidateResponse.Clear()
	return h
}
