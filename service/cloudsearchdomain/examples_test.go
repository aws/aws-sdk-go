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

func ExampleCloudSearchDomain_Suggest() {
	svc := cloudsearchdomain.New(nil)

	params := &cloudsearchdomain.SuggestInput{
		Query:     aws.String("Query"),     // Required
		Suggester: aws.String("Suggester"), // Required
		Size:      aws.Long(1),
	}
	resp, err := svc.Suggest(params)

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

func ExampleCloudSearchDomain_UploadDocuments() {
	svc := cloudsearchdomain.New(nil)

	params := &cloudsearchdomain.UploadDocumentsInput{
		ContentType: aws.String("ContentType"),          // Required
		Documents:   bytes.NewReader([]byte("PAYLOAD")), // Required
	}
	resp, err := svc.UploadDocuments(params)

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