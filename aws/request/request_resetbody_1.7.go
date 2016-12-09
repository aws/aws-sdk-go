// +build !go1.8

package request

// Pre Go 1.8 funcitonality to always set Request.Body to a value regardless
// if it will actually be used. Go 1.8 removed undocumented feature the SDK
// expected where the request body would only be sent if the length was > 0.
func resetBody(r *Request) {
	r.safeBody = newOffsetReader(r.Body, r.BodyStart)
	r.HTTPRequest.Body = r.safeBody
}
