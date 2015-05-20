package cloudsearchdomain_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cloudsearchdomain"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudSearchDomain_Search() {
	svc := cloudsearchdomain.New(nil)

	params := &cloudsearchdomain.SearchInput{
		Query:        aws.String("Query"), // Required
		Cursor:       aws.String("Cursor"),
		Expr:         aws.String("Expr"),
		Facet:        aws.String("Facet"),
		FilterQuery:  aws.String("FilterQuery"),
		Highlight:    aws.String("Highlight"),
		Partial:      aws.Boolean(true),
		QueryOptions: aws.String("QueryOptions"),
		QueryParser:  aws.String("QueryParser"),
		Return:       aws.String("Return"),
		Size:         aws.Long(1),
		Sort:         aws.String("Sort"),
		Start:        aws.Long(1),
	}
	resp, err := svc.Search(params)

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

func ExampleCloudSearchDomain_Suggest() {
	svc := cloudsearchdomain.New(nil)

	params := &cloudsearchdomain.SuggestInput{
		Query:     aws.String("Query"),     // Required
		Suggester: aws.String("Suggester"), // Required
		Size:      aws.Long(1),
	}
	resp, err := svc.Suggest(params)

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

func ExampleCloudSearchDomain_UploadDocuments() {
	svc := cloudsearchdomain.New(nil)

	params := &cloudsearchdomain.UploadDocumentsInput{
		ContentType: aws.String("ContentType"),          // Required
		Documents:   bytes.NewReader([]byte("PAYLOAD")), // Required
	}
	resp, err := svc.UploadDocuments(params)

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