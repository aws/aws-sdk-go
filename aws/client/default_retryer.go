package client

import (
	"math"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/internal/sdkrand"
)

// DefaultRetryer implements basic retry logic using exponential backoff for
// most services. If you want to implement custom retry logic, implement the
// request.Retryer interface or create a structure type that composes this
// struct and override the specific methods. For example, to override only
// the MaxRetries method:
//
//   type retryer struct {
//       client.DefaultRetryer
//   }
//
//   // This implementation always has 100 max retries
//   func (d retryer) MaxRetries() int { return 100 }
type DefaultRetryer struct {
	NumMaxRetries    int
	MinRetryDelay    time.Duration
	MinThrottleDelay time.Duration
	MaxRetryDelay    time.Duration
	MaxThrottleDelay time.Duration
}

const (
	// DefaultRetryerMaxNumRetries sets maximum number of retries
	DefaultRetryerMaxNumRetries = 3

	// DefaultRetryerMinRetryDelay sets minimum retry delay
	DefaultRetryerMinRetryDelay = 30 * time.Millisecond

	// DefaultRetryerMinThrottleDelay sets minimum delay when throttled
	DefaultRetryerMinThrottleDelay = 500 * time.Millisecond

	// DefaultRetryerMaxRetryDelay sets maximum retry delay
	DefaultRetryerMaxRetryDelay = 300 * time.Second

	// DefaultRetryerMaxThrottleDelay sets maximum delay when throttled
	DefaultRetryerMaxThrottleDelay = 300 * time.Second
)

// MaxRetries returns the number of maximum returns the service will use to make
// an individual API request.
func (d DefaultRetryer) MaxRetries() int {
	return d.NumMaxRetries
}

// setRetryerDefaults sets the default values of the retryer if not set
func (d *DefaultRetryer) setRetryerDefaults() {
	if d.MinRetryDelay == 0 {
		d.MinRetryDelay = DefaultRetryerMinRetryDelay
	}
	if d.MaxRetryDelay == 0 {
		d.MaxRetryDelay = DefaultRetryerMaxRetryDelay
	}
	if d.MinThrottleDelay == 0 {
		d.MinThrottleDelay = DefaultRetryerMinThrottleDelay
	}
	if d.MaxThrottleDelay == 0 {
		d.MaxThrottleDelay = DefaultRetryerMaxThrottleDelay
	}
}

// RetryRules returns the delay duration before retrying this request again
func (d DefaultRetryer) RetryRules(r *request.Request) time.Duration {

	// if number of max retries is zero, no retries will be performed.
	if d.NumMaxRetries == 0 {
		return 0
	}

	// Sets default value for retryer members
	d.setRetryerDefaults()

	// minDelay is the minimum retryer delay
	minDelay := d.MinRetryDelay

	var initialDelay time.Duration

	isThrottle := r.IsErrorThrottle()
	if isThrottle {
		if delay, ok := getRetryAfterDelay(r); ok {
			initialDelay = delay
		}
		minDelay = d.MinThrottleDelay
	}

	retryCount := r.RetryCount

	// maxDelay the maximum retryer delay
	maxDelay := d.MaxRetryDelay

	if isThrottle {
		maxDelay = d.MaxThrottleDelay
	}

	var delay time.Duration

	// Logic to cap the retry count based on the minDelay provided
	actualRetryCount := int(math.Log2(float64(minDelay))) + 1
	if actualRetryCount < 63-retryCount {
		delay = time.Duration(1<<uint64(retryCount)) * getJitterDelay(minDelay)
		if delay > maxDelay {
			delay = getJitterDelay(maxDelay / 2)
		}
	} else {
		delay = getJitterDelay(maxDelay / 2)
	}
	return delay + initialDelay
}

// getJitterDelay returns a jittered delay for retry
func getJitterDelay(duration time.Duration) time.Duration {
	return time.Duration(sdkrand.SeededRand.Int63n(int64(duration)) + int64(duration))
}

// ShouldRetry returns true if the request should be retried.
func (d DefaultRetryer) ShouldRetry(r *request.Request) bool {
	// If one of the other handlers already set the retry state
	// we don't want to override it based on the service's state
	if r.Retryable != nil {
		return *r.Retryable
	}

	if r.HTTPResponse.StatusCode >= 500 && r.HTTPResponse.StatusCode != 501 {
		return true
	}

	return r.IsErrorRetryable() || r.IsErrorThrottle()
}

// This will look in the Retry-After header, RFC 7231, for how long
// it will wait before attempting another request
func getRetryAfterDelay(r *request.Request) (time.Duration, bool) {
	if !canUseRetryAfterHeader(r) {
		return 0, false
	}

	delayStr := r.HTTPResponse.Header.Get("Retry-After")
	if len(delayStr) == 0 {
		return 0, false
	}

	delay, err := strconv.Atoi(delayStr)
	if err != nil {
		return 0, false
	}

	return time.Duration(delay) * time.Second, true
}

// Will look at the status code to see if the retry header pertains to
// the status code.
func canUseRetryAfterHeader(r *request.Request) bool {
	switch r.HTTPResponse.StatusCode {
	case 429:
	case 503:
	default:
		return false
	}

	return true
}
