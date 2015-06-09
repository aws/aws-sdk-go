package sqs

import "github.com/aws/aws-sdk-go/aws"

func init() {
	initRequest = func(r *aws.Request) {
		setupChecksumValidation(r)
	}
}
