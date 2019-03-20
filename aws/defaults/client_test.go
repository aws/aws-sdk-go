package defaults

import (
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"
)

func TestGetDefaultClient(t *testing.T) {
	c := getDefaultClient()

	if e, a := http.DefaultClient, c; e == a {
		t.Fatalf("expect not equal")
	}
	if e, a := http.DefaultClient.Transport, c.Transport; e == a {
		t.Errorf("expect not equal")
	}

	tr := c.Transport.(*http.Transport)
	if e, a := defaultResponseHeaderTimeout, tr.ResponseHeaderTimeout; e != a {
		t.Errorf("expect %v response header timeout, got %v", e, a)
	}
	if e, a := defaultMaxIdleConnsPerHost, tr.MaxIdleConnsPerHost; e != a {
		t.Errorf("expect %v idle conns per host, got %v", e, a)
	}
}

func TestCopyClient(t *testing.T) {
	orig := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 10,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Jar:     &cookiejar.Jar{},
		Timeout: 10 * time.Second,
	}

	c := copyClient(orig)

	if c == orig {
		t.Fatalf("expect not equal")
	}
	if e, a := orig.Transport, c.Transport; e == a {
		t.Errorf("expect transports not equal")
	}
	origTr := orig.Transport.(*http.Transport)
	tr := orig.Transport.(*http.Transport)
	if e, a := origTr.MaxIdleConns, tr.MaxIdleConns; e != a {
		t.Errorf("expect %v max idle conns, got %v", e, a)
	}
	if c.CheckRedirect == nil {
		t.Errorf("expect CheckRedirect set")
	}
	if e, a := orig.Jar, c.Jar; e != a {
		t.Errorf("expect same jar, was not")
	}
	if e, a := orig.Timeout, c.Timeout; e != a {
		t.Errorf("expect %v timeout, got %v", e, a)
	}
}

type mockRoundTripper struct {
	MaxIdleConns int
}

func (m *mockRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, nil
}

func TestCopyClient_NonTransportRoundtripper(t *testing.T) {
	orig := &http.Client{
		Transport: &mockRoundTripper{
			MaxIdleConns: 10,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Jar:     &cookiejar.Jar{},
		Timeout: 10 * time.Second,
	}

	c := copyClient(orig)

	if c == orig {
		t.Fatalf("expect not equal")
	}
	if e, a := orig.Transport, c.Transport; e != a {
		t.Errorf("expect equal")
	}
	tr := orig.Transport.(*mockRoundTripper)
	if e, a := 10, tr.MaxIdleConns; e != a {
		t.Errorf("expect equal")
	}
	if c.CheckRedirect == nil {
		t.Errorf("expect set")
	}
	if e, a := orig.Jar, c.Jar; e != a {
		t.Errorf("expect equal")
	}
	if e, a := orig.Timeout, c.Timeout; e != a {
		t.Errorf("expect equal")
	}
}
