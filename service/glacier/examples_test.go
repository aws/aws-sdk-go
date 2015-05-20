package glacier_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/glacier"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleGlacier_AbortMultipartUpload() {
	svc := glacier.New(nil)

	params := &glacier.AbortMultipartUploadInput{
		AccountID: aws.String("string"), // Required
		UploadID:  aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.AbortMultipartUpload(params)

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

func ExampleGlacier_CompleteMultipartUpload() {
	svc := glacier.New(nil)

	params := &glacier.CompleteMultipartUploadInput{
		AccountID:   aws.String("string"), // Required
		UploadID:    aws.String("string"), // Required
		VaultName:   aws.String("string"), // Required
		ArchiveSize: aws.String("string"),
		Checksum:    aws.String("string"),
	}
	resp, err := svc.CompleteMultipartUpload(params)

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

func ExampleGlacier_CreateVault() {
	svc := glacier.New(nil)

	params := &glacier.CreateVaultInput{
		AccountID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.CreateVault(params)

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

func ExampleGlacier_DeleteArchive() {
	svc := glacier.New(nil)

	params := &glacier.DeleteArchiveInput{
		AccountID: aws.String("string"), // Required
		ArchiveID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DeleteArchive(params)

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

func ExampleGlacier_DeleteVault() {
	svc := glacier.New(nil)

	params := &glacier.DeleteVaultInput{
		AccountID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DeleteVault(params)

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

func ExampleGlacier_DeleteVaultNotifications() {
	svc := glacier.New(nil)

	params := &glacier.DeleteVaultNotificationsInput{
		AccountID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DeleteVaultNotifications(params)

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

func ExampleGlacier_DescribeJob() {
	svc := glacier.New(nil)

	params := &glacier.DescribeJobInput{
		AccountID: aws.String("string"), // Required
		JobID:     aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DescribeJob(params)

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

func ExampleGlacier_DescribeVault() {
	svc := glacier.New(nil)

	params := &glacier.DescribeVaultInput{
		AccountID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DescribeVault(params)

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

func ExampleGlacier_GetDataRetrievalPolicy() {
	svc := glacier.New(nil)

	params := &glacier.GetDataRetrievalPolicyInput{
		AccountID: aws.String("string"), // Required
	}
	resp, err := svc.GetDataRetrievalPolicy(params)

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

func ExampleGlacier_GetJobOutput() {
	svc := glacier.New(nil)

	params := &glacier.GetJobOutputInput{
		AccountID: aws.String("string"), // Required
		JobID:     aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Range:     aws.String("string"),
	}
	resp, err := svc.GetJobOutput(params)

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

func ExampleGlacier_GetVaultNotifications() {
	svc := glacier.New(nil)

	params := &glacier.GetVaultNotificationsInput{
		AccountID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.GetVaultNotifications(params)

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

func ExampleGlacier_InitiateJob() {
	svc := glacier.New(nil)

	params := &glacier.InitiateJobInput{
		AccountID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		JobParameters: &glacier.JobParameters{
			ArchiveID:   aws.String("string"),
			Description: aws.String("string"),
			Format:      aws.String("string"),
			InventoryRetrievalParameters: &glacier.InventoryRetrievalJobInput{
				EndDate:   aws.String("string"),
				Limit:     aws.String("string"),
				Marker:    aws.String("string"),
				StartDate: aws.String("string"),
			},
			RetrievalByteRange: aws.String("string"),
			SNSTopic:           aws.String("string"),
			Type:               aws.String("string"),
		},
	}
	resp, err := svc.InitiateJob(params)

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

func ExampleGlacier_InitiateMultipartUpload() {
	svc := glacier.New(nil)

	params := &glacier.InitiateMultipartUploadInput{
		AccountID:          aws.String("string"), // Required
		VaultName:          aws.String("string"), // Required
		ArchiveDescription: aws.String("string"),
		PartSize:           aws.String("string"),
	}
	resp, err := svc.InitiateMultipartUpload(params)

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

func ExampleGlacier_ListJobs() {
	svc := glacier.New(nil)

	params := &glacier.ListJobsInput{
		AccountID:  aws.String("string"), // Required
		VaultName:  aws.String("string"), // Required
		Completed:  aws.String("string"),
		Limit:      aws.String("string"),
		Marker:     aws.String("string"),
		Statuscode: aws.String("string"),
	}
	resp, err := svc.ListJobs(params)

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

func ExampleGlacier_ListMultipartUploads() {
	svc := glacier.New(nil)

	params := &glacier.ListMultipartUploadsInput{
		AccountID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Limit:     aws.String("string"),
		Marker:    aws.String("string"),
	}
	resp, err := svc.ListMultipartUploads(params)

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

func ExampleGlacier_ListParts() {
	svc := glacier.New(nil)

	params := &glacier.ListPartsInput{
		AccountID: aws.String("string"), // Required
		UploadID:  aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Limit:     aws.String("string"),
		Marker:    aws.String("string"),
	}
	resp, err := svc.ListParts(params)

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

func ExampleGlacier_ListVaults() {
	svc := glacier.New(nil)

	params := &glacier.ListVaultsInput{
		AccountID: aws.String("string"), // Required
		Limit:     aws.String("string"),
		Marker:    aws.String("string"),
	}
	resp, err := svc.ListVaults(params)

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

func ExampleGlacier_SetDataRetrievalPolicy() {
	svc := glacier.New(nil)

	params := &glacier.SetDataRetrievalPolicyInput{
		AccountID: aws.String("string"), // Required
		Policy: &glacier.DataRetrievalPolicy{
			Rules: []*glacier.DataRetrievalRule{
				&glacier.DataRetrievalRule{ // Required
					BytesPerHour: aws.Long(1),
					Strategy:     aws.String("string"),
				},
				// More values...
			},
		},
	}
	resp, err := svc.SetDataRetrievalPolicy(params)

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

func ExampleGlacier_SetVaultNotifications() {
	svc := glacier.New(nil)

	params := &glacier.SetVaultNotificationsInput{
		AccountID: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		VaultNotificationConfig: &glacier.VaultNotificationConfig{
			Events: []*string{
				aws.String("string"), // Required
				// More values...
			},
			SNSTopic: aws.String("string"),
		},
	}
	resp, err := svc.SetVaultNotifications(params)

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

func ExampleGlacier_UploadArchive() {
	svc := glacier.New(nil)

	params := &glacier.UploadArchiveInput{
		AccountID:          aws.String("string"), // Required
		VaultName:          aws.String("string"), // Required
		ArchiveDescription: aws.String("string"),
		Body:               bytes.NewReader([]byte("PAYLOAD")),
		Checksum:           aws.String("string"),
	}
	resp, err := svc.UploadArchive(params)

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

func ExampleGlacier_UploadMultipartPart() {
	svc := glacier.New(nil)

	params := &glacier.UploadMultipartPartInput{
		AccountID: aws.String("string"), // Required
		UploadID:  aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Body:      bytes.NewReader([]byte("PAYLOAD")),
		Checksum:  aws.String("string"),
		Range:     aws.String("string"),
	}
	resp, err := svc.UploadMultipartPart(params)

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