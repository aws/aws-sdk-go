// +build go1.7

package v2

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
)

func requestContext(r *http.Request) aws.Context {
	return r.Context()
}
