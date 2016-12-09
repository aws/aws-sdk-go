// +build !go1.8

package request

func resetBody(r *Request) {
	r.safeBody = newOffsetReader(r.Body, r.BodyStart)
	r.HTTPRequest.Body = r.safeBody
}
