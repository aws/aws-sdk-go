// +build !go1.6

package request_test

import (
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

var errTimeout = awserr.New("foo", "bar", errors.New("net/http: request canceled Timeout"))

func r(t *testing.T, url string) *http.Request {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("can't forge request: %v", err)
	}
	return r
}

type BrokenRoundTripper struct {
	fn func(*http.Request)
	rt http.RoundTripper
}

func (r *BrokenRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	r.fn(req)
	return r.rt.RoundTrip(req)
}

func Test_shouldRetryCancel(t *testing.T) {

	///////////
	// too quick timeout
	cli := http.Client{
		Timeout:   time.Nanosecond,
		Transport: http.DefaultTransport,
	}

	// resp, err := cli.Do(r(t, "https://179.179.179.179/no/such/host"))
	// if resp != nil {
	// 	resp.Body.Close()
	// }
	// if err == nil {
	// 	t.Fatal("This should have failed.")
	// }
	// debugerr(t, err)

	// if shouldRetryCancel(err) == false {
	// 	t.Errorf("this request timed out and should be retried")
	// }

	///////////
	// request cancelled
	tr := http.DefaultTransport.(*http.Transport)
	tr.MaxIdleConnsPerHost = 1
	cli.Timeout = 0 // reset timeout
	// cli.Jar = append()

	srvr := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		tr.CancelRequest(req)
		http.Redirect(rw, req, "potato", 308)
	}))
	defer srvr.Close()
	r := r(t, srvr.URL)
	// r.GetBody = func() (io.ReadCloser, error) {
	// 	tr.CancelRequest(r)
	// 	time.Sleep(1 * time.Second)
	// 	return nil, nil
	// }
	tr.Dial = func(network, addr string) (net.Conn, error) {
		tr.CancelRequest(r)
		return (&net.Dialer{}).Dial(network, addr)
	}

	cli.Transport = &BrokenRoundTripper{
		fn: func(req *http.Request) {
			http.DefaultTransport.(*http.Transport).CancelRequest(req)
		},
		rt: http.DefaultTransport,
	}
	resp, err := cli.Do(r)
	if resp != nil {
		resp.Body.Close()
	}
	if err == nil {
		t.Fatal("This should have failed.")
	}

	debugerr(t, err)

	t.Fatal("just for the logs")
}

func debugerr(t *testing.T, err error) {
	type temporaryError interface {
		Temporary() bool
	}

	switch err := err.(type) {
	case temporaryError:
		t.Logf("%s is a temporary error: %t", err, err.Temporary())
		return
	case *url.Error:
		// we should be before 1.5
		// that's our case !
		t.Logf("err: %s", err)
		t.Logf("err: %#v", err.Err)
		if operr, ok := err.Err.(*net.OpError); ok {
			t.Logf("operr: %#v", operr)
		}
		debugerr(t, err.Err)
		return
	default:
		return
	}
}
