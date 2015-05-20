package cloudtrail_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cloudtrail"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudTrail_CreateTrail() {
	svc := cloudtrail.New(nil)

	params := &cloudtrail.CreateTrailInput{
		Name:                       aws.String("String"), // Required
		S3BucketName:               aws.String("String"), // Required
		CloudWatchLogsLogGroupARN:  aws.String("String"),
		CloudWatchLogsRoleARN:      aws.String("String"),
		IncludeGlobalServiceEvents: aws.Boolean(true),
		S3KeyPrefix:                aws.String("String"),
		SNSTopicName:               aws.String("String"),
	}
	resp, err := svc.CreateTrail(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudTrail_DeleteTrail() {
	svc := cloudtrail.New(nil)

	params := &cloudtrail.DeleteTrailInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.DeleteTrail(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudTrail_DescribeTrails() {
	svc := cloudtrail.New(nil)

	params := &cloudtrail.DescribeTrailsInput{
		TrailNameList: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrails(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudTrail_GetTrailStatus() {
	svc := cloudtrail.New(nil)

	params := &cloudtrail.GetTrailStatusInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.GetTrailStatus(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudTrail_LookupEvents() {
	svc := cloudtrail.New(nil)

	params := &cloudtrail.LookupEventsInput{
		EndTime: aws.Time(time.Now()),
		LookupAttributes: []*cloudtrail.LookupAttribute{
			&cloudtrail.LookupAttribute{ // Required
				AttributeKey:   aws.String("LookupAttributeKey"), // Required
				AttributeValue: aws.String("String"),             // Required
			},
			// More values...
		},
		MaxResults: aws.Long(1),
		NextToken:  aws.String("NextToken"),
		StartTime:  aws.Time(time.Now()),
	}
	resp, err := svc.LookupEvents(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudTrail_StartLogging() {
	svc := cloudtrail.New(nil)

	params := &cloudtrail.StartLoggingInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.StartLogging(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudTrail_StopLogging() {
	svc := cloudtrail.New(nil)

	params := &cloudtrail.StopLoggingInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.StopLogging(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudTrail_UpdateTrail() {
	svc := cloudtrail.New(nil)

	params := &cloudtrail.UpdateTrailInput{
		Name: aws.String("String"), // Required
		CloudWatchLogsLogGroupARN:  aws.String("String"),
		CloudWatchLogsRoleARN:      aws.String("String"),
		IncludeGlobalServiceEvents: aws.Boolean(true),
		S3BucketName:               aws.String("String"),
		S3KeyPrefix:                aws.String("String"),
		SNSTopicName:               aws.String("String"),
	}
	resp, err := svc.UpdateTrail(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}