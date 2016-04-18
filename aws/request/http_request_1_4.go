// +build go1.4 !go1.5

import "net/http"

func newHTTPRequest(r *http.Request) *http.Request {
	return &http.Request{
		URL:           r.URL,
		Header:        r.Header,
		Close:         r.Close,
		Form:          r.Form,
		PostForm:      r.PostForm,
		Body:          r.Body,
		MultipartForm: r.MultipartForm,
		Host:          r.Host,
		Method:        r.Method,
		Proto:         r.Proto,
		ContentLength: r.ContentLength,
	}
}
