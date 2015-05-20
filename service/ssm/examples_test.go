package ssm_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/ssm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSSM_CreateAssociation() {
	svc := ssm.New(nil)

	params := &ssm.CreateAssociationInput{
		InstanceID: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.CreateAssociation(params)

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

func ExampleSSM_CreateAssociationBatch() {
	svc := ssm.New(nil)

	params := &ssm.CreateAssociationBatchInput{
		Entries: []*ssm.CreateAssociationBatchRequestEntry{ // Required
			&ssm.CreateAssociationBatchRequestEntry{ // Required
				InstanceID: aws.String("InstanceId"),
				Name:       aws.String("DocumentName"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateAssociationBatch(params)

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

func ExampleSSM_CreateDocument() {
	svc := ssm.New(nil)

	params := &ssm.CreateDocumentInput{
		Content: aws.String("DocumentContent"), // Required
		Name:    aws.String("DocumentName"),    // Required
	}
	resp, err := svc.CreateDocument(params)

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

func ExampleSSM_DeleteAssociation() {
	svc := ssm.New(nil)

	params := &ssm.DeleteAssociationInput{
		InstanceID: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.DeleteAssociation(params)

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

func ExampleSSM_DeleteDocument() {
	svc := ssm.New(nil)

	params := &ssm.DeleteDocumentInput{
		Name: aws.String("DocumentName"), // Required
	}
	resp, err := svc.DeleteDocument(params)

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

func ExampleSSM_DescribeAssociation() {
	svc := ssm.New(nil)

	params := &ssm.DescribeAssociationInput{
		InstanceID: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.DescribeAssociation(params)

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

func ExampleSSM_DescribeDocument() {
	svc := ssm.New(nil)

	params := &ssm.DescribeDocumentInput{
		Name: aws.String("DocumentName"), // Required
	}
	resp, err := svc.DescribeDocument(params)

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

func ExampleSSM_GetDocument() {
	svc := ssm.New(nil)

	params := &ssm.GetDocumentInput{
		Name: aws.String("DocumentName"), // Required
	}
	resp, err := svc.GetDocument(params)

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

func ExampleSSM_ListAssociations() {
	svc := ssm.New(nil)

	params := &ssm.ListAssociationsInput{
		AssociationFilterList: []*ssm.AssociationFilter{ // Required
			&ssm.AssociationFilter{ // Required
				Key:   aws.String("AssociationFilterKey"),   // Required
				Value: aws.String("AssociationFilterValue"), // Required
			},
			// More values...
		},
		MaxResults: aws.Long(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListAssociations(params)

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

func ExampleSSM_ListDocuments() {
	svc := ssm.New(nil)

	params := &ssm.ListDocumentsInput{
		DocumentFilterList: []*ssm.DocumentFilter{
			&ssm.DocumentFilter{ // Required
				Key:   aws.String("DocumentFilterKey"),   // Required
				Value: aws.String("DocumentFilterValue"), // Required
			},
			// More values...
		},
		MaxResults: aws.Long(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListDocuments(params)

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

func ExampleSSM_UpdateAssociationStatus() {
	svc := ssm.New(nil)

	params := &ssm.UpdateAssociationStatusInput{
		AssociationStatus: &ssm.AssociationStatus{ // Required
			Date:           aws.Time(time.Now()),                // Required
			Message:        aws.String("StatusMessage"),         // Required
			Name:           aws.String("AssociationStatusName"), // Required
			AdditionalInfo: aws.String("StatusAdditionalInfo"),
		},
		InstanceID: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.UpdateAssociationStatus(params)

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