package support_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/support"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSupport_AddAttachmentsToSet() {
	svc := support.New(nil)

	params := &support.AddAttachmentsToSetInput{
		Attachments: []*support.Attachment{ // Required
			&support.Attachment{ // Required
				Data:     []byte("PAYLOAD"),
				FileName: aws.String("FileName"),
			},
			// More values...
		},
		AttachmentSetID: aws.String("AttachmentSetId"),
	}
	resp, err := svc.AddAttachmentsToSet(params)

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

func ExampleSupport_AddCommunicationToCase() {
	svc := support.New(nil)

	params := &support.AddCommunicationToCaseInput{
		CommunicationBody: aws.String("CommunicationBody"), // Required
		AttachmentSetID:   aws.String("AttachmentSetId"),
		CCEmailAddresses: []*string{
			aws.String("CcEmailAddress"), // Required
			// More values...
		},
		CaseID: aws.String("CaseId"),
	}
	resp, err := svc.AddCommunicationToCase(params)

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

func ExampleSupport_CreateCase() {
	svc := support.New(nil)

	params := &support.CreateCaseInput{
		CommunicationBody: aws.String("CommunicationBody"), // Required
		Subject:           aws.String("Subject"),           // Required
		AttachmentSetID:   aws.String("AttachmentSetId"),
		CCEmailAddresses: []*string{
			aws.String("CcEmailAddress"), // Required
			// More values...
		},
		CategoryCode: aws.String("CategoryCode"),
		IssueType:    aws.String("IssueType"),
		Language:     aws.String("Language"),
		ServiceCode:  aws.String("ServiceCode"),
		SeverityCode: aws.String("SeverityCode"),
	}
	resp, err := svc.CreateCase(params)

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

func ExampleSupport_DescribeAttachment() {
	svc := support.New(nil)

	params := &support.DescribeAttachmentInput{
		AttachmentID: aws.String("AttachmentId"), // Required
	}
	resp, err := svc.DescribeAttachment(params)

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

func ExampleSupport_DescribeCases() {
	svc := support.New(nil)

	params := &support.DescribeCasesInput{
		AfterTime:  aws.String("AfterTime"),
		BeforeTime: aws.String("BeforeTime"),
		CaseIDList: []*string{
			aws.String("CaseId"), // Required
			// More values...
		},
		DisplayID:             aws.String("DisplayId"),
		IncludeCommunications: aws.Boolean(true),
		IncludeResolvedCases:  aws.Boolean(true),
		Language:              aws.String("Language"),
		MaxResults:            aws.Long(1),
		NextToken:             aws.String("NextToken"),
	}
	resp, err := svc.DescribeCases(params)

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

func ExampleSupport_DescribeCommunications() {
	svc := support.New(nil)

	params := &support.DescribeCommunicationsInput{
		CaseID:     aws.String("CaseId"), // Required
		AfterTime:  aws.String("AfterTime"),
		BeforeTime: aws.String("BeforeTime"),
		MaxResults: aws.Long(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeCommunications(params)

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

func ExampleSupport_DescribeServices() {
	svc := support.New(nil)

	params := &support.DescribeServicesInput{
		Language: aws.String("Language"),
		ServiceCodeList: []*string{
			aws.String("ServiceCode"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeServices(params)

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

func ExampleSupport_DescribeSeverityLevels() {
	svc := support.New(nil)

	params := &support.DescribeSeverityLevelsInput{
		Language: aws.String("Language"),
	}
	resp, err := svc.DescribeSeverityLevels(params)

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

func ExampleSupport_DescribeTrustedAdvisorCheckRefreshStatuses() {
	svc := support.New(nil)

	params := &support.DescribeTrustedAdvisorCheckRefreshStatusesInput{
		CheckIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrustedAdvisorCheckRefreshStatuses(params)

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

func ExampleSupport_DescribeTrustedAdvisorCheckResult() {
	svc := support.New(nil)

	params := &support.DescribeTrustedAdvisorCheckResultInput{
		CheckID:  aws.String("String"), // Required
		Language: aws.String("String"),
	}
	resp, err := svc.DescribeTrustedAdvisorCheckResult(params)

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

func ExampleSupport_DescribeTrustedAdvisorCheckSummaries() {
	svc := support.New(nil)

	params := &support.DescribeTrustedAdvisorCheckSummariesInput{
		CheckIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrustedAdvisorCheckSummaries(params)

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

func ExampleSupport_DescribeTrustedAdvisorChecks() {
	svc := support.New(nil)

	params := &support.DescribeTrustedAdvisorChecksInput{
		Language: aws.String("String"), // Required
	}
	resp, err := svc.DescribeTrustedAdvisorChecks(params)

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

func ExampleSupport_RefreshTrustedAdvisorCheck() {
	svc := support.New(nil)

	params := &support.RefreshTrustedAdvisorCheckInput{
		CheckID: aws.String("String"), // Required
	}
	resp, err := svc.RefreshTrustedAdvisorCheck(params)

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

func ExampleSupport_ResolveCase() {
	svc := support.New(nil)

	params := &support.ResolveCaseInput{
		CaseID: aws.String("CaseId"),
	}
	resp, err := svc.ResolveCase(params)

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