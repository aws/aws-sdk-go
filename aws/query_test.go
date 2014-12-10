package aws_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"sync"
	"testing"

	"github.com/stripe/aws-go/aws"
)

func TestQueryRequest(t *testing.T) {
	var m sync.Mutex
	var httpReq *http.Request
	var form url.Values

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			m.Lock()
			defer m.Unlock()

			httpReq = r

			if err := r.ParseForm(); err != nil {
				t.Fatal(err)
			}
			form = r.Form

			fmt.Fprintln(w, `<Thing><IpAddress>woo</IpAddress></Thing>`)
		},
	))
	defer server.Close()

	client := aws.QueryClient{
		Context: aws.Context{
			Service: "animals",
			Region:  "us-west-2",
			Credentials: aws.Creds(
				"accessKeyID",
				"secretAccessKey",
				"securityToken",
			),
		},
		Client:     http.DefaultClient,
		Endpoint:   server.URL,
		APIVersion: "1.1",
	}

	req := fakeQueryRequest{
		ID:       "one",
		Count:    500,
		BigCount: 600,
		OK:       true,
		Things:   []string{"woo", "yeah"},
	}
	var resp fakeQueryResponse
	if err := client.Do("GetIP", "POST", "/", &req, &resp); err != nil {
		t.Fatal(err)
	}

	m.Lock()
	defer m.Unlock()

	if v, want := httpReq.Method, "POST"; v != want {
		t.Errorf("Method was %v but expected %v", v, want)
	}

	if httpReq.Header.Get("Authorization") == "" {
		t.Error("Authorization header is missing")
	}

	if v, want := httpReq.Header.Get("Content-Type"), "application/x-www-form-urlencoded"; v != want {
		t.Errorf("Content-Type was %v but expected %v", v, want)
	}

	if v, want := httpReq.Header.Get("User-Agent"), "aws-go"; v != want {
		t.Errorf("User-Agent was %v but expected %v", v, want)
	}

	if err := httpReq.ParseForm(); err != nil {
		t.Fatal(err)
	}

	expectedForm := url.Values{
		"Action":          []string{"GetIP"},
		"Version":         []string{"1.1"},
		"Id":              []string{"one"},
		"BigCount":        []string{"600"},
		"Count":           []string{"500"},
		"OK":              []string{"true"},
		"Things.member.1": []string{"woo"},
		"Things.member.2": []string{"yeah"},
	}

	if !reflect.DeepEqual(form, expectedForm) {
		t.Errorf("Post body was %s, but expected %s", form.Encode(), expectedForm.Encode())
	}

	if want := (fakeQueryResponse{IPAddress: "woo"}); want != resp {
		t.Errorf("Response was %#v, but expected %#v", resp, want)
	}
}

func TestQueryRequestError(t *testing.T) {
	var m sync.Mutex
	var httpReq *http.Request
	var form url.Values

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			m.Lock()
			defer m.Unlock()

			httpReq = r

			if err := r.ParseForm(); err != nil {
				t.Fatal(err)
			}
			form = r.Form

			w.WriteHeader(400)
			fmt.Fprintln(w, `<ErrorResponse>
<RequestId>woo</RequestId>
<Error>
<Type>Problem</Type>
<Code>Uh Oh</Code>
<Message>You done did it</Message>
</Error>
</ErrorResponse>`)
		},
	))
	defer server.Close()

	client := aws.QueryClient{
		Context: aws.Context{
			Service: "animals",
			Region:  "us-west-2",
			Credentials: aws.Creds(
				"accessKeyID",
				"secretAccessKey",
				"securityToken",
			),
		},
		Client:     http.DefaultClient,
		Endpoint:   server.URL,
		APIVersion: "1.1",
	}

	req := fakeQueryRequest{
		ID:       "one",
		Count:    500,
		BigCount: 600,
		OK:       true,
		Things:   []string{"woo", "yeah"},
	}
	var resp fakeQueryResponse
	err := client.Do("GetIP", "POST", "/", &req, &resp)
	if err == nil {
		t.Fatal("Expected an error but none was returned")
	}

	if err, ok := err.(aws.APIError); ok {
		if v, want := err.Type, "Problem"; v != want {
			t.Errorf("Error type was %v, but expected %v", v, want)
		}

		if v, want := err.Code, "Uh Oh"; v != want {
			t.Errorf("Error type was %v, but expected %v", v, want)
		}

		if v, want := err.Message, "You done did it"; v != want {
			t.Errorf("Error message was %v, but expected %v", v, want)
		}
	} else {
		t.Errorf("Unknown error returned: %#v", err)
	}
}

type fakeQueryRequest struct {
	ID       string `xml:"Id"`
	Count    int
	BigCount int64
	OK       bool
	Things   []string
}

type fakeQueryResponse struct {
	IPAddress string `xml:"IpAddress"`
}
