// +build go1.8

package request

import (
	"io"
	"net/http"
)

// Go 1.8 tightened and clarified the rules code needs to use when building
// requests with the http package. Go 1.8 removed the automatic detection
// of if the Request.Body was empty, or actually had bytes in it. The SDK
// always sets the Request.Body even if it is empty and should not actually
// be sent. This is incorrect.
//
// Go 1.8 did add a http.NoBody value that the SDK can use to tell the http
// client that the request really should be sent without a body. The
// Request.Body cannot be set to nil, which is preferable, because the
// field is exported and could introduce nil pointer dereferences for users
// of the SDK if they used that field.
//
// Related golang/go#18257
func resetBody(r *Request) {
	curOffset, _ := r.Body.Seek(0, io.SeekCurrent)
	endOffset, _ := r.Body.Seek(0, io.SeekEnd)
	r.Body.Seek(r.BodyStart, io.SeekStart)

	r.safeBody = newOffsetReader(r.Body, r.BodyStart)

	if endOffset-curOffset == 0 {
		r.HTTPRequest.Body = http.NoBody
	} else {
		r.HTTPRequest.Body = r.safeBody
	}
}
