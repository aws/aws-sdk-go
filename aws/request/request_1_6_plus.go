// +build go1.6

package request

import (
	"net"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func shouldRetryCancel(r *Request) bool {
	awsErr, ok := r.Error.(awserr.Error)
	timeoutErr := false
	if ok {
		netErr, netOK := awsErr.OrigErr().(net.Error)
		timeoutErr = netOK && netErr.Timeout()
	}

	// There can be two types of canceled errors here.
	// The first being a net.Error and the other being an error.
	// If the request was timed out, we want to continue the retry
	// process. Otherwise, return the canceled error.
	return timeoutErr || !strings.Contains(r.Error.Error(), "net/http: request canceled")
}
