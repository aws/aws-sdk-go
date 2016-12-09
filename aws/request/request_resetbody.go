// +build go1.8

package request

import (
	"io"
	"net/http"
)

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
