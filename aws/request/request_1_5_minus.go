// +build !go1.6

package request

import (
	"strings"
)

func shouldRetryCancel(r *Request) bool {
	return strings.Contains(r.Error.Error(), "Timeout") &&
		strings.Contains(r.Error.Error(), "net/http: request canceled")
}
