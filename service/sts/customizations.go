package sts

import "github.com/aws/aws-sdk-go/aws/request"

func init() {
	initRequest = customizeRequest
}

func customizeRequest(r *request.Request) {
	r.RetryCodes = append(r.RetryCodes, ErrCodeIDPCommunicationErrorException)
}
