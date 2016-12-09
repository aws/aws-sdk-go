// +build go1.8

package request

import (
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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
	r.safeBody = newOffsetReader(r.Body, r.BodyStart)

	seekable := false
	// Determine if the seeker is actually seekable. ReaderSeekerCloser
	// hides the fact that a io.Readers might not actually be seekable.
	switch v := r.Body.(type) {
	case aws.ReaderSeekerCloser:
		seekable = v.IsSeeker()
	case *aws.ReaderSeekerCloser:
		seekable = v.IsSeeker()
	default:
		seekable = true
	}

	if seekable {
		curOffset, err := r.Body.Seek(0, io.SeekCurrent)
		if err != nil {
			r.Error = awserr.New("SerializationError", "failed to seek request body", err)
			return
		}
		endOffset, err := r.Body.Seek(0, io.SeekEnd)
		if err != nil {
			r.Error = awserr.New("SerializationError", "failed to seek request body", err)
			return
		}
		_, err = r.Body.Seek(r.BodyStart, io.SeekStart)
		if err != nil {
			r.Error = awserr.New("SerializationError", "failed to seek request body", err)
			return
		}

		if endOffset-curOffset == 0 {
			r.HTTPRequest.Body = http.NoBody
		} else {
			r.HTTPRequest.Body = r.safeBody
		}
	} else {
		// Hack to prevent sending bodies for methods where the body
		// should be ignored by the server. Sending bodies on these
		// methods without an associated ContentLength will cause the
		// request to socket timeout because the server does not handle
		// Transfer-Encoding: chunked bodies for these methods.
		switch r.Operation.HTTPMethod {
		case "GET", "HEAD", "DELETE":
			r.HTTPRequest.Body = http.NoBody
		default:
			r.HTTPRequest.Body = r.safeBody
		}
	}
}
