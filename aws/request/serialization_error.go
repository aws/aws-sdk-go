// +build !appengine

package request

import (
	"net"
	"os"
	"syscall"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func isSerializationErrorRetryable(err error) bool {
	if err == nil {
		return false
	}

	if aerr, ok := err.(awserr.Error); ok {
		return isCodeRetryable(aerr.Code())
	}

	if opErr, ok := err.(*net.OpError); ok {
		if sysErr, ok := opErr.Err.(*os.SyscallError); ok {
			return sysErr.Err == syscall.ECONNRESET
		}
	}

	return false
}
