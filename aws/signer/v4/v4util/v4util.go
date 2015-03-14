package v4util

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
)

// Add a Authorization header to requests using the v4 signature format.
//
// It is required that the request has previously been signed using the query
// string. This is useful for non-production services like DynamoDB Local which
// do not properly support signing using the query string.
func SignWithHeader(req *aws.Request) {
	headers := map[string]string{
		"X-Amz-Algorithm":     "",
		"X-Amz-Credential":    "",
		"X-Amz-SignedHeaders": "",
		"X-Amz-Signature":     "",
	}
	query := req.HTTPRequest.URL.Query()
	for header, _ := range headers {
		v := query.Get(header)
		if v == "" {
			req.Error = fmt.Errorf("'%s' was not found in the query string", header)
			return
		}
		headers[header] = v
	}
	authorization := fmt.Sprintf("%s Credential=%s, SignedHeaders=%s, Signature=%s",
		headers["X-Amz-Algorithm"], headers["X-Amz-Credential"],
		headers["X-Amz-SignedHeaders"], headers["X-Amz-Signature"])
	req.HTTPRequest.Header.Set("Authorization", authorization)
}
