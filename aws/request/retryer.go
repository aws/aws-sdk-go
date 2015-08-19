package request

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

type Retryer interface {
	RetryRules(*Request) time.Duration
	ShouldRetry(*Request) bool
	MaxRetries() uint
}

// retryableCodes is a collection of service response codes which are retry-able
// without any further action.
var retryableCodes = map[string]struct{}{
	"RequestError":                           {},
	"ProvisionedThroughputExceededException": {},
	"Throttling":                             {},
	"ThrottlingException":                    {},
	"RequestLimitExceeded":                   {},
	"RequestThrottled":                       {},
}

// credsExpiredCodes is a collection of error codes which signify the credentials
// need to be refreshed. Expired tokens require refreshing of credentials, and
// resigning before the request can be retried.
var credsExpiredCodes = map[string]struct{}{
	"ExpiredToken":          {},
	"ExpiredTokenException": {},
	"RequestExpired":        {}, // EC2 Only
}

func isCodeRetryable(code string) bool {
	if _, ok := retryableCodes[code]; ok {
		return true
	}

	return isCodeExpiredCreds(code)
}

func isCodeExpiredCreds(code string) bool {
	_, ok := credsExpiredCodes[code]
	return ok
}

func (r *Request) IsErrorRetryable() bool {
	if r.Error != nil {
		if err, ok := r.Error.(awserr.Error); ok {
			return isCodeRetryable(err.Code())
		}
	}
	return false
}

func (r *Request) IsErrorExpired() bool {
	if r.Error != nil {
		if err, ok := r.Error.(awserr.Error); ok {
			return isCodeExpiredCreds(err.Code())
		}
	}
	return false
}
