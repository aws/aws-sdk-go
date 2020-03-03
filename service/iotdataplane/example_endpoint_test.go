package iotdataplane_test

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

func ExampleIoTDataPlane_DescribeEndpoint_shared00() {
	sess := unit.Session
	if v := aws.StringValue(sess.Config.Region); len(v) == 0 {
		sess.Config.Region = aws.String("us-east-1")
	}

	ctrlSvc := iot.New(sess)
	descResp, err := ctrlSvc.DescribeEndpoint(&iot.DescribeEndpointInput{})
	if err != nil {
		log.Fatalf("failed to get dataplane endpoint, %v", err)
	}

	dataSvc := iotdataplane.New(sess, &aws.Config{
		Endpoint: descResp.EndpointAddress,
	})
	_, err = dataSvc.GetThingShadow(&iotdataplane.GetThingShadowInput{
		ThingName: aws.String("fake-thing"),
	})
	if err == nil {
		log.Fatalf("expect error")
	}

	aerr, ok := err.(awserr.Error)
	if !ok {
		log.Fatalf("expect awserr.Error, got %T, %v", err, err)
	}
	if e, a := "ResourceNotFoundException", aerr.Code(); e != a {
		log.Fatalf("expect %v error, got %v", e, aerr)
	}
}
