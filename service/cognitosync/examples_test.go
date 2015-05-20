package cognitosync_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cognitosync"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCognitoSync_BulkPublish() {
	svc := cognitosync.New(nil)

	params := &cognitosync.BulkPublishInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.BulkPublish(params)

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

func ExampleCognitoSync_DeleteDataset() {
	svc := cognitosync.New(nil)

	params := &cognitosync.DeleteDatasetInput{
		DatasetName:    aws.String("DatasetName"),    // Required
		IdentityID:     aws.String("IdentityId"),     // Required
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DeleteDataset(params)

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

func ExampleCognitoSync_DescribeDataset() {
	svc := cognitosync.New(nil)

	params := &cognitosync.DescribeDatasetInput{
		DatasetName:    aws.String("DatasetName"),    // Required
		IdentityID:     aws.String("IdentityId"),     // Required
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeDataset(params)

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

func ExampleCognitoSync_DescribeIdentityPoolUsage() {
	svc := cognitosync.New(nil)

	params := &cognitosync.DescribeIdentityPoolUsageInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeIdentityPoolUsage(params)

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

func ExampleCognitoSync_DescribeIdentityUsage() {
	svc := cognitosync.New(nil)

	params := &cognitosync.DescribeIdentityUsageInput{
		IdentityID:     aws.String("IdentityId"),     // Required
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeIdentityUsage(params)

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

func ExampleCognitoSync_GetBulkPublishDetails() {
	svc := cognitosync.New(nil)

	params := &cognitosync.GetBulkPublishDetailsInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.GetBulkPublishDetails(params)

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

func ExampleCognitoSync_GetIdentityPoolConfiguration() {
	svc := cognitosync.New(nil)

	params := &cognitosync.GetIdentityPoolConfigurationInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.GetIdentityPoolConfiguration(params)

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

func ExampleCognitoSync_ListDatasets() {
	svc := cognitosync.New(nil)

	params := &cognitosync.ListDatasetsInput{
		IdentityID:     aws.String("IdentityId"),     // Required
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
		MaxResults:     aws.Long(1),
		NextToken:      aws.String("String"),
	}
	resp, err := svc.ListDatasets(params)

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

func ExampleCognitoSync_ListIdentityPoolUsage() {
	svc := cognitosync.New(nil)

	params := &cognitosync.ListIdentityPoolUsageInput{
		MaxResults: aws.Long(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListIdentityPoolUsage(params)

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

func ExampleCognitoSync_ListRecords() {
	svc := cognitosync.New(nil)

	params := &cognitosync.ListRecordsInput{
		DatasetName:      aws.String("DatasetName"),    // Required
		IdentityID:       aws.String("IdentityId"),     // Required
		IdentityPoolID:   aws.String("IdentityPoolId"), // Required
		LastSyncCount:    aws.Long(1),
		MaxResults:       aws.Long(1),
		NextToken:        aws.String("String"),
		SyncSessionToken: aws.String("SyncSessionToken"),
	}
	resp, err := svc.ListRecords(params)

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

func ExampleCognitoSync_RegisterDevice() {
	svc := cognitosync.New(nil)

	params := &cognitosync.RegisterDeviceInput{
		IdentityID:     aws.String("IdentityId"),     // Required
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
		Platform:       aws.String("Platform"),       // Required
		Token:          aws.String("PushToken"),      // Required
	}
	resp, err := svc.RegisterDevice(params)

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

func ExampleCognitoSync_SetIdentityPoolConfiguration() {
	svc := cognitosync.New(nil)

	params := &cognitosync.SetIdentityPoolConfigurationInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
		CognitoStreams: &cognitosync.CognitoStreams{
			RoleARN:         aws.String("AssumeRoleArn"),
			StreamName:      aws.String("StreamName"),
			StreamingStatus: aws.String("StreamingStatus"),
		},
		PushSync: &cognitosync.PushSync{
			ApplicationARNs: []*string{
				aws.String("ApplicationArn"), // Required
				// More values...
			},
			RoleARN: aws.String("AssumeRoleArn"),
		},
	}
	resp, err := svc.SetIdentityPoolConfiguration(params)

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

func ExampleCognitoSync_SubscribeToDataset() {
	svc := cognitosync.New(nil)

	params := &cognitosync.SubscribeToDatasetInput{
		DatasetName:    aws.String("DatasetName"),    // Required
		DeviceID:       aws.String("DeviceId"),       // Required
		IdentityID:     aws.String("IdentityId"),     // Required
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.SubscribeToDataset(params)

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

func ExampleCognitoSync_UnsubscribeFromDataset() {
	svc := cognitosync.New(nil)

	params := &cognitosync.UnsubscribeFromDatasetInput{
		DatasetName:    aws.String("DatasetName"),    // Required
		DeviceID:       aws.String("DeviceId"),       // Required
		IdentityID:     aws.String("IdentityId"),     // Required
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.UnsubscribeFromDataset(params)

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

func ExampleCognitoSync_UpdateRecords() {
	svc := cognitosync.New(nil)

	params := &cognitosync.UpdateRecordsInput{
		DatasetName:      aws.String("DatasetName"),      // Required
		IdentityID:       aws.String("IdentityId"),       // Required
		IdentityPoolID:   aws.String("IdentityPoolId"),   // Required
		SyncSessionToken: aws.String("SyncSessionToken"), // Required
		ClientContext:    aws.String("ClientContext"),
		DeviceID:         aws.String("DeviceId"),
		RecordPatches: []*cognitosync.RecordPatch{
			&cognitosync.RecordPatch{ // Required
				Key:                    aws.String("RecordKey"), // Required
				Op:                     aws.String("Operation"), // Required
				SyncCount:              aws.Long(1),             // Required
				DeviceLastModifiedDate: aws.Time(time.Now()),
				Value: aws.String("RecordValue"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateRecords(params)

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