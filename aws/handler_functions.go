package aws

import (
	"fmt"
	"strconv"
	"time"
)

var sleepDelay = func(delay time.Duration) {
	time.Sleep(delay)
}

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

func UserAgentHandler(r *Request) {
	r.HTTPRequest.Header.Set("User-Agent", SDKName+"/"+SDKVersion)
}

func SendHandler(r *Request) {
	r.HTTPResponse, r.Error = r.Service.Config.HTTPClient.Do(r.HTTPRequest)
}

func ValidateResponseHandler(r *Request) {
	if r.HTTPResponse.StatusCode == 0 || r.HTTPResponse.StatusCode >= 400 {
		err := APIError{
			StatusCode: r.HTTPResponse.StatusCode,
			RetryCount: r.RetryCount,
		}
		r.Error = err
		err.Retryable = r.Service.ShouldRetry(r)
		err.RetryDelay = r.Service.RetryRules(r)
		r.Error = err
	}
}

func AfterRetryHandler(r *Request) {
	delay := 0 * time.Second
	willRetry := false

	if err := Error(r.Error); err != nil {
		delay = err.RetryDelay
		if err.Retryable && r.RetryCount < r.Service.MaxRetries() {
			r.RetryCount++
			willRetry = true
		}
	}

	if willRetry {
		r.Error = nil
		sleepDelay(delay)
	}
}
