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

func ExampleSTS_DecodeAuthorizationMessage() {
	svc := sts.New(nil)

	params := &sts.DecodeAuthorizationMessageInput{
		EncodedMessage: aws.String("encodedMessageType"), // Required
	}
	resp, err := svc.DecodeAuthorizationMessage(params)

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

func ExampleSTS_GetFederationToken() {
	svc := sts.New(nil)

	params := &sts.GetFederationTokenInput{
		Name:            aws.String("userNameType"), // Required
		DurationSeconds: aws.Long(1),
		Policy:          aws.String("sessionPolicyDocumentType"),
	}
	resp, err := svc.GetFederationToken(params)

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

func ExampleSTS_GetSessionToken() {
	svc := sts.New(nil)

	params := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Long(1),
		SerialNumber:    aws.String("serialNumberType"),
		TokenCode:       aws.String("tokenCodeType"),
	}
	resp, err := svc.GetSessionToken(params)

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