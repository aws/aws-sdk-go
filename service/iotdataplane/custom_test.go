// +build go1.8

package iotdataplane_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

func Test_DescribeEndpoint(t *testing.T) {
	sess := unit.Session
	if v := aws.StringValue(sess.Config.Region); len(v) == 0 {
		sess.Config.Region = aws.String("us-east-1")
	}

	ctrlSvc := iot.New(sess)
	req, _ := ctrlSvc.DescribeEndpointRequest(&iot.DescribeEndpointInput{})
	if e, a := "iot.mock-region.amazonaws.com", req.HTTPRequest.URL.Host; e!=a {
		t.Fatalf("expected %v, got %v", e, a)
	}

	dataSvc := iotdataplane.New(sess, &aws.Config{
		Endpoint: aws.String("mockEndpoint"),
	})
	req, _ = dataSvc.GetThingShadowRequest(&iotdataplane.GetThingShadowInput{
		ThingName: aws.String("fake-thing"),
	})
	if e, a := "mockEndpoint", req.HTTPRequest.URL.Host; e!=a {
		t.Fatalf("expected %v, got %v", e, a)
	}
	if e, a:= "/things/{thingName}/shadow",req.HTTPRequest.URL.RawPath; e!=a {
		t.Fatalf("expected %v, got %v", e, a)
	}
}
