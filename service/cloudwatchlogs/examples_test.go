package cloudwatchlogs_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cloudwatchlogs"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudWatchLogs_CreateLogGroup() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.CreateLogGroup(params)

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

func ExampleCloudWatchLogs_CreateLogStream() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
	}
	resp, err := svc.CreateLogStream(params)

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

func ExampleCloudWatchLogs_DeleteLogGroup() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteLogGroupInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteLogGroup(params)

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

func ExampleCloudWatchLogs_DeleteLogStream() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteLogStreamInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
	}
	resp, err := svc.DeleteLogStream(params)

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

func ExampleCloudWatchLogs_DeleteMetricFilter() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteMetricFilterInput{
		FilterName:   aws.String("FilterName"),   // Required
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteMetricFilter(params)

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

func ExampleCloudWatchLogs_DeleteRetentionPolicy() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteRetentionPolicyInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteRetentionPolicy(params)

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

func ExampleCloudWatchLogs_DescribeLogGroups() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DescribeLogGroupsInput{
		Limit:              aws.Long(1),
		LogGroupNamePrefix: aws.String("LogGroupName"),
		NextToken:          aws.String("NextToken"),
	}
	resp, err := svc.DescribeLogGroups(params)

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

func ExampleCloudWatchLogs_DescribeLogStreams() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        aws.String("LogGroupName"), // Required
		Descending:          aws.Boolean(true),
		Limit:               aws.Long(1),
		LogStreamNamePrefix: aws.String("LogStreamName"),
		NextToken:           aws.String("NextToken"),
		OrderBy:             aws.String("OrderBy"),
	}
	resp, err := svc.DescribeLogStreams(params)

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

func ExampleCloudWatchLogs_DescribeMetricFilters() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DescribeMetricFiltersInput{
		LogGroupName:     aws.String("LogGroupName"), // Required
		FilterNamePrefix: aws.String("FilterName"),
		Limit:            aws.Long(1),
		NextToken:        aws.String("NextToken"),
	}
	resp, err := svc.DescribeMetricFilters(params)

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

func ExampleCloudWatchLogs_GetLogEvents() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
		EndTime:       aws.Long(1),
		Limit:         aws.Long(1),
		NextToken:     aws.String("NextToken"),
		StartFromHead: aws.Boolean(true),
		StartTime:     aws.Long(1),
	}
	resp, err := svc.GetLogEvents(params)

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

func ExampleCloudWatchLogs_PutLogEvents() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{ // Required
			&cloudwatchlogs.InputLogEvent{ // Required
				Message:   aws.String("EventMessage"), // Required
				Timestamp: aws.Long(1),                // Required
			},
			// More values...
		},
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
		SequenceToken: aws.String("SequenceToken"),
	}
	resp, err := svc.PutLogEvents(params)

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

func ExampleCloudWatchLogs_PutMetricFilter() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutMetricFilterInput{
		FilterName:    aws.String("FilterName"),    // Required
		FilterPattern: aws.String("FilterPattern"), // Required
		LogGroupName:  aws.String("LogGroupName"),  // Required
		MetricTransformations: []*cloudwatchlogs.MetricTransformation{ // Required
			&cloudwatchlogs.MetricTransformation{ // Required
				MetricName:      aws.String("MetricName"),      // Required
				MetricNamespace: aws.String("MetricNamespace"), // Required
				MetricValue:     aws.String("MetricValue"),     // Required
			},
			// More values...
		},
	}
	resp, err := svc.PutMetricFilter(params)

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

func ExampleCloudWatchLogs_PutRetentionPolicy() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutRetentionPolicyInput{
		LogGroupName:    aws.String("LogGroupName"), // Required
		RetentionInDays: aws.Long(1),                // Required
	}
	resp, err := svc.PutRetentionPolicy(params)

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

func ExampleCloudWatchLogs_TestMetricFilter() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.TestMetricFilterInput{
		FilterPattern: aws.String("FilterPattern"), // Required
		LogEventMessages: []*string{ // Required
			aws.String("EventMessage"), // Required
			// More values...
		},
	}
	resp, err := svc.TestMetricFilter(params)

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