package sts_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/sts"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSTS_AssumeRole() {
	svc := sts.New(nil)

	params := &sts.AssumeRoleInput{
		RoleARN:         aws.String("arnType"),      // Required
		RoleSessionName: aws.String("userNameType"), // Required
		DurationSeconds: aws.Long(1),
		ExternalID:      aws.String("externalIdType"),
		Policy:          aws.String("sessionPolicyDocumentType"),
		SerialNumber:    aws.String("serialNumberType"),
		TokenCode:       aws.String("tokenCodeType"),
	}
	resp, err := svc.AssumeRole(params)

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

func ExampleSTS_AssumeRoleWithSAML() {
	svc := sts.New(nil)

	params := &sts.AssumeRoleWithSAMLInput{
		PrincipalARN:    aws.String("arnType"),           // Required
		RoleARN:         aws.String("arnType"),           // Required
		SAMLAssertion:   aws.String("SAMLAssertionType"), // Required
		DurationSeconds: aws.Long(1),
		Policy:          aws.String("sessionPolicyDocumentType"),
	}
	resp, err := svc.AssumeRoleWithSAML(params)

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

func ExampleSTS_AssumeRoleWithWebIdentity() {
	svc := sts.New(nil)

	params := &sts.AssumeRoleWithWebIdentityInput{
		RoleARN:          aws.String("arnType"),         // Required
		RoleSessionName:  aws.String("userNameType"),    // Required
		WebIdentityToken: aws.String("clientTokenType"), // Required
		DurationSeconds:  aws.Long(1),
		Policy:           aws.String("sessionPolicyDocumentType"),
		ProviderID:       aws.String("urlType"),
	}
	resp, err := svc.AssumeRoleWithWebIdentity(params)

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

func ExampleSTS_DecodeAuthorizationMessage() {
	svc := sts.New(nil)

	params := &sts.DecodeAuthorizationMessageInput{
		EncodedMessage: aws.String("encodedMessageType"), // Required
	}
	resp, err := svc.DecodeAuthorizationMessage(params)

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

func ExampleSTS_GetFederationToken() {
	svc := sts.New(nil)

	params := &sts.GetFederationTokenInput{
		Name:            aws.String("userNameType"), // Required
		DurationSeconds: aws.Long(1),
		Policy:          aws.String("sessionPolicyDocumentType"),
	}
	resp, err := svc.GetFederationToken(params)

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

func ExampleSTS_GetSessionToken() {
	svc := sts.New(nil)

	params := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Long(1),
		SerialNumber:    aws.String("serialNumberType"),
		TokenCode:       aws.String("tokenCodeType"),
	}
	resp, err := svc.GetSessionToken(params)

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