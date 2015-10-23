//Package iotdataplane provides gucumber integration tests support.
package iotdataplane

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@iotdataplane", func() {
		svc := iot.New(nil)
		result, err := svc.DescribeEndpoint(&iot.DescribeEndpointInput{})
		if err != nil {
			World["error"] = err
			return
		}

		World["client"] = iotdataplane.New(aws.NewConfig().
			WithEndpoint(*result.EndpointAddress))
	})
}
