package s3

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
)

var errSSERequiresSSL = aws.APIError{
	Code:    "ConfigError",
	Message: "cannot send SSE keys over HTTP.",
}

func validateSSERequiresSSL(r *aws.Request) {
	if r.HTTPRequest.URL.Scheme != "https" {
		p := awsutil.ValuesAtPath(r.Params, "SSECustomerKey||CopySourceSSECustomerKey")
		if len(p) > 0 {
			r.Error = errSSERequiresSSL
		}
	}
}
