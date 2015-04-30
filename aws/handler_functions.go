package aws

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

var sleepDelay = func(delay time.Duration) {
	time.Sleep(delay)
}

type lener interface {
	Len() int
}

func BuildContentLength(r *Request) {
	if slength := r.HTTPRequest.Header.Get("Content-Length"); slength != "" {
		length, _ := strconv.ParseInt(slength, 10, 64)
		r.HTTPRequest.ContentLength = length
		return
	}

	var length int64
	switch body := r.Body.(type) {
	case nil:
		length = 0
	case lener:
		length = int64(body.Len())
	case io.Seeker:
		r.bodyStart, _ = body.Seek(0, 1)
		end, _ := body.Seek(0, 2)
		body.Seek(r.bodyStart, 0) // make sure to seek back to original location
		length = end - r.bodyStart
	default:
		panic("Cannot get length of body, must provide `ContentLength`")
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
		// this may be replaced by an UnmarshalError handler
		r.Error = &APIError{
			StatusCode: r.HTTPResponse.StatusCode,
			Code:       "UnknownError",
			Message:    "unknown error",
		}
	}
}

func RetryHandler(r *Request) {
	// If one of the other handlers already set the retry state
	// we don't want to override it based on the service's state
	if !r.Retryable.IsSet() {
		r.Retryable.Set(r.Service.ShouldRetry(r))
	}

	r.RetryDelay = r.Service.RetryRules(r)
}

func AfterRetryHandler(r *Request) {
	if r.WillRetry() {
		sleepDelay(r.RetryDelay)

		r.RetryCount++
		r.Error = nil
	}
}

var (
	ErrMissingRegion   = fmt.Errorf("could not find region configuration.")
	ErrMissingEndpoint = fmt.Errorf("`Endpoint' configuration is required for this service.")
)

func ValidateEndpointHandler(r *Request) {
	if r.Service.SigningRegion == "" && r.Service.Config.Region == "" {
		r.Error = ErrMissingRegion
	} else if r.Service.Endpoint == "" {
		r.Error = ErrMissingEndpoint
	}
}
