package sqs

import "github.com/aws/aws-sdk-go/aws/service"

func init() {
	initRequest = func(r *service.Request) {
		setupChecksumValidation(r)
	}
}
