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
		cur, _ := body.Seek(0, 1)
		end, _ := body.Seek(0, 2)
		body.Seek(cur, 0) // make sure to seek back to original location
		length = end - cur
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

var (
	ErrMissingRegion      = fmt.Errorf("could not find region configuration.")
	ErrMissingCredentials = fmt.Errorf("could not find credentials configuration.")
)

func ValidateEndpointHandler(r *Request) {
	if r.Service.SigningRegion == "" && r.Service.Config.Region == "" {
		r.Error = ErrMissingRegion
	}
}

func ValidateCredentialsHandler(r *Request) {
	if r.Service.Config.Credentials == nil {
		r.Error = ErrMissingCredentials
		return
	}

	creds, err := r.Service.Config.Credentials.Credentials()
	if err != nil || creds.AccessKeyID == "" || creds.SecretAccessKey == "" {
		r.Error = ErrMissingCredentials
	}
}
