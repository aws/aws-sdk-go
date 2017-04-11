// +build appengine

package request

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func isSerializationErrorRetryable(err error) bool {
	if err == nil {
		return false
	}

	if aerr, ok := err.(awserr.Error); ok {
		return isCodeRetryable(aerr.Code())
	}

	return strings.Contain(err.Error(), "connection reset")
}
