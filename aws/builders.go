package aws

import (
	"fmt"
	"strconv"
)

func BuildContentLength(r *Request) {
	var length int64
	strlen := r.HTTPRequest.Header.Get("Content-Length")
	if strlen == "" && r.Body != nil {
		body, ok := r.Body.(interface {
			Len() int
		})
		if ok {
			length = int64(body.Len())
		} else {
			panic("Cannot get length of body, must provide `ContentLength`")
		}
	} else {
		length, _ = strconv.ParseInt(strlen, 10, 64)
	}
	r.HTTPRequest.ContentLength = length
	r.HTTPRequest.Header.Set("Content-Length", fmt.Sprintf("%d", length))
}
