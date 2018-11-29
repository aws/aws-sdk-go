package request

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

var ersrTimeout = awserr.New("foo", "bar", errors.New("net/http: request canceled Timeout"))

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

func Test_shouldRetryCancel_timeout(t *testing.T) {

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cli := http.Client{
		Timeout:   time.Nanosecond,
		Transport: tr,
	}

	resp, err := cli.Do(r(t, "https://179.179.179.179/no/such/host"))
	if resp != nil {
		resp.Body.Close()
	}
	if err == nil {
		t.Fatal("This should have failed.")
	}
	debugerr(t, err)

	if shouldRetryCancel(err) == false {
		t.Errorf("this request timed out and should be retried")
	}
}

func Test_shouldRetryCancel_cancelled(t *testing.T) {

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cli := http.Client{
		Transport: tr,
	}

	unblockc := make(chan bool)
	srvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello")
		w.(http.Flusher).Flush() // send headers and some body
		<-unblockc               // block forever
	}))
	defer srvr.Close()
	defer close(unblockc)

	r := r(t, srvr.URL)
	ch := make(chan struct{})
	r.Cancel = ch
	close(ch) // request is cancelled before anything

	resp, err := cli.Do(r)
	if resp != nil {
		resp.Body.Close()
	}
	if err == nil {
		t.Fatal("This should have failed.")
	}

	debugerr(t, err)

	if shouldRetryCancel(err) == true {
		t.Errorf("this request was cancelled and should not be retried")
	}
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
