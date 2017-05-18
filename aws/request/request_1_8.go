// +build go1.8

package request

import (
	"net/http"
)

// NoBody is a http.NoBody reader instructing Go HTTP client to not include
// and body in the HTTP request.
var NoBody = http.NoBody

// ResetBody rewinds the request body back to its starting position, and
// set's the HTTP Request body reference. When the body is read prior
// to being sent in the HTTP request it will need to be rewound.
//
// Will also set the Go 1.8's http.Request.GetBody member to allow retrying
// PUT/POST redirects.
func (r *Request) ResetBody() {
	body, err := r.getNextRequestBody()
	if err != nil {
		r.Error = err
		return
	}

	r.HTTPRequest.Body = body
	r.HTTPRequest.GetBody = r.getNextRequestBody
}
