package configservice_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/configservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleConfigService_DeleteDeliveryChannel() {
	svc := configservice.New(nil)

	params := &configservice.DeleteDeliveryChannelInput{
		DeliveryChannelName: aws.String("ChannelName"), // Required
	}
	resp, err := svc.DeleteDeliveryChannel(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_DeliverConfigSnapshot() {
	svc := configservice.New(nil)

	params := &configservice.DeliverConfigSnapshotInput{
		DeliveryChannelName: aws.String("ChannelName"), // Required
	}
	resp, err := svc.DeliverConfigSnapshot(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_DescribeConfigurationRecorderStatus() {
	svc := configservice.New(nil)

	params := &configservice.DescribeConfigurationRecorderStatusInput{
		ConfigurationRecorderNames: []*string{
			aws.String("RecorderName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeConfigurationRecorderStatus(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_DescribeConfigurationRecorders() {
	svc := configservice.New(nil)

	params := &configservice.DescribeConfigurationRecordersInput{
		ConfigurationRecorderNames: []*string{
			aws.String("RecorderName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeConfigurationRecorders(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_DescribeDeliveryChannelStatus() {
	svc := configservice.New(nil)

	params := &configservice.DescribeDeliveryChannelStatusInput{
		DeliveryChannelNames: []*string{
			aws.String("ChannelName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeDeliveryChannelStatus(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_DescribeDeliveryChannels() {
	svc := configservice.New(nil)

	params := &configservice.DescribeDeliveryChannelsInput{
		DeliveryChannelNames: []*string{
			aws.String("ChannelName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeDeliveryChannels(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_GetResourceConfigHistory() {
	svc := configservice.New(nil)

	params := &configservice.GetResourceConfigHistoryInput{
		ResourceID:         aws.String("ResourceId"),   // Required
		ResourceType:       aws.String("ResourceType"), // Required
		ChronologicalOrder: aws.String("ChronologicalOrder"),
		EarlierTime:        aws.Time(time.Now()),
		LaterTime:          aws.Time(time.Now()),
		Limit:              aws.Long(1),
		NextToken:          aws.String("NextToken"),
	}
	resp, err := svc.GetResourceConfigHistory(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_PutConfigurationRecorder() {
	svc := configservice.New(nil)

	params := &configservice.PutConfigurationRecorderInput{
		ConfigurationRecorder: &configservice.ConfigurationRecorder{ // Required
			Name:    aws.String("RecorderName"),
			RoleARN: aws.String("String"),
		},
	}
	resp, err := svc.PutConfigurationRecorder(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_PutDeliveryChannel() {
	svc := configservice.New(nil)

	params := &configservice.PutDeliveryChannelInput{
		DeliveryChannel: &configservice.DeliveryChannel{ // Required
			Name:         aws.String("ChannelName"),
			S3BucketName: aws.String("String"),
			S3KeyPrefix:  aws.String("String"),
			SNSTopicARN:  aws.String("String"),
		},
	}
	resp, err := svc.PutDeliveryChannel(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_StartConfigurationRecorder() {
	svc := configservice.New(nil)

	params := &configservice.StartConfigurationRecorderInput{
		ConfigurationRecorderName: aws.String("RecorderName"), // Required
	}
	resp, err := svc.StartConfigurationRecorder(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleConfigService_StopConfigurationRecorder() {
	svc := configservice.New(nil)

	params := &configservice.StopConfigurationRecorderInput{
		ConfigurationRecorderName: aws.String("RecorderName"), // Required
	}
	resp, err := svc.StopConfigurationRecorder(params)

	if reqerr, ok := err.(awserr.RequestFailure); ok {
		// A service error occurred
		fmt.Println(reqerr.Code(), reqerr.Message(), reqerr.StatusCode(), reqerr.RequestID())
	} else {
		// A non-service error occurred.
		fmt.Println(err.Code(), reqerr.Message(), err.OrigErr())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}