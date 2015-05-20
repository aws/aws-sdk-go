package cloudhsm_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cloudhsm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudHSM_CreateHAPG() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.CreateHAPGInput{
		Label: aws.String("Label"), // Required
	}
	resp, err := svc.CreateHAPG(params)

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

func ExampleCloudHSM_CreateHSM() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.CreateHSMInput{
		IAMRoleARN:       aws.String("IamRoleArn"),       // Required
		SSHKey:           aws.String("SshKey"),           // Required
		SubnetID:         aws.String("SubnetId"),         // Required
		SubscriptionType: aws.String("SubscriptionType"), // Required
		ClientToken:      aws.String("ClientToken"),
		ENIIP:            aws.String("IpAddress"),
		ExternalID:       aws.String("ExternalId"),
		SyslogIP:         aws.String("IpAddress"),
	}
	resp, err := svc.CreateHSM(params)

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

func ExampleCloudHSM_CreateLunaClient() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.CreateLunaClientInput{
		Certificate: aws.String("Certificate"), // Required
		Label:       aws.String("ClientLabel"),
	}
	resp, err := svc.CreateLunaClient(params)

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

func ExampleCloudHSM_DeleteHAPG() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DeleteHAPGInput{
		HAPGARN: aws.String("HapgArn"), // Required
	}
	resp, err := svc.DeleteHAPG(params)

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

func ExampleCloudHSM_DeleteHSM() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DeleteHSMInput{
		HSMARN: aws.String("HsmArn"), // Required
	}
	resp, err := svc.DeleteHSM(params)

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

func ExampleCloudHSM_DeleteLunaClient() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DeleteLunaClientInput{
		ClientARN: aws.String("ClientArn"), // Required
	}
	resp, err := svc.DeleteLunaClient(params)

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

func ExampleCloudHSM_DescribeHAPG() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DescribeHAPGInput{
		HAPGARN: aws.String("HapgArn"), // Required
	}
	resp, err := svc.DescribeHAPG(params)

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

func ExampleCloudHSM_DescribeHSM() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DescribeHSMInput{
		HSMARN:          aws.String("HsmArn"),
		HSMSerialNumber: aws.String("HsmSerialNumber"),
	}
	resp, err := svc.DescribeHSM(params)

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

func ExampleCloudHSM_DescribeLunaClient() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DescribeLunaClientInput{
		CertificateFingerprint: aws.String("CertificateFingerprint"),
		ClientARN:              aws.String("ClientArn"),
	}
	resp, err := svc.DescribeLunaClient(params)

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

func ExampleCloudHSM_GetConfig() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.GetConfigInput{
		ClientARN:     aws.String("ClientArn"),     // Required
		ClientVersion: aws.String("ClientVersion"), // Required
		HAPGList: []*string{ // Required
			aws.String("HapgArn"), // Required
			// More values...
		},
	}
	resp, err := svc.GetConfig(params)

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

func ExampleCloudHSM_ListAvailableZones() {
	svc := cloudhsm.New(nil)

	var params *cloudhsm.ListAvailableZonesInput
	resp, err := svc.ListAvailableZones(params)

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

func ExampleCloudHSM_ListHSMs() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ListHSMsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListHSMs(params)

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

func ExampleCloudHSM_ListHapgs() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ListHapgsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListHapgs(params)

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

func ExampleCloudHSM_ListLunaClients() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ListLunaClientsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListLunaClients(params)

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

func ExampleCloudHSM_ModifyHAPG() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ModifyHAPGInput{
		HAPGARN: aws.String("HapgArn"), // Required
		Label:   aws.String("Label"),
		PartitionSerialList: []*string{
			aws.String("PartitionSerial"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyHAPG(params)

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

func ExampleCloudHSM_ModifyHSM() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ModifyHSMInput{
		HSMARN:     aws.String("HsmArn"), // Required
		ENIIP:      aws.String("IpAddress"),
		ExternalID: aws.String("ExternalId"),
		IAMRoleARN: aws.String("IamRoleArn"),
		SubnetID:   aws.String("SubnetId"),
		SyslogIP:   aws.String("IpAddress"),
	}
	resp, err := svc.ModifyHSM(params)

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

func ExampleCloudHSM_ModifyLunaClient() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ModifyLunaClientInput{
		Certificate: aws.String("Certificate"), // Required
		ClientARN:   aws.String("ClientArn"),   // Required
	}
	resp, err := svc.ModifyLunaClient(params)

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