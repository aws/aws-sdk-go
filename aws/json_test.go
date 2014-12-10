package aws_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stripe/aws-go/aws"
)

func TestJSONRequest(t *testing.T) {
	var m sync.Mutex
	var httpReq *http.Request
	var body []byte

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			m.Lock()
			defer m.Unlock()

			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			defer r.Body.Close()

			httpReq = r
			body = b

			fmt.Fprintln(w, `{"TailWagged":true}`)
		},
	))
	defer server.Close()

	client := aws.JSONClient{
		Context: aws.Context{
			Service: "animals",
			Region:  "us-west-2",
			Credentials: aws.Creds(
				"accessKeyID",
				"secretAccessKey",
				"securityToken",
			),
			Clock: func() time.Time {
				return time.Date(
					2012, 12, 10, 11, 12, 43, 0,
					time.FixedZone("PDT", -8),
				)
			},
		},
		Client:       http.DefaultClient,
		Endpoint:     server.URL,
		TargetPrefix: "Animals",
		JSONVersion:  "1.1",
	}

	req := fakeJSONRequest{Name: "Penny"}
	var resp fakeJSONResponse
	if err := client.Do("PetTheDog", "POST", "/", req, &resp); err != nil {
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

	if v, want := httpReq.Header.Get("Content-Type"), "application/x-amz-json-1.1"; v != want {
		t.Errorf("Content-Type was %v but expected %v", v, want)
	}

	if v, want := httpReq.Header.Get("User-Agent"), "aws-go"; v != want {
		t.Errorf("User-Agent was %v but expected %v", v, want)
	}

	if v, want := httpReq.Header.Get("X-Amz-Target"), "Animals.PetTheDog"; v != want {
		t.Errorf("X-Amz-Target was %v but expected %v", v, want)
	}

	if v, want := string(body), `{"Name":"Penny"}`; v != want {
		t.Errorf("Body was %v but expected %v", v, want)
	}

	x := fakeJSONResponse{TailWagged: true}
	if v, want := resp, x; v != want {
		t.Errorf("Response was %#v but expected %#v", v, want)
	}
}

func TestJSONRequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(400)
			fmt.Fprintln(w, `{"__type":"Problem", "message":"What even"}`)
		},
	))
	defer server.Close()

	client := aws.JSONClient{
		Context: aws.Context{
			Service: "animals",
			Region:  "us-west-2",
			Credentials: aws.Creds(
				"accessKeyID",
				"secretAccessKey",
				"securityToken",
			),
			Clock: time.Now,
		},
		Client:       http.DefaultClient,
		Endpoint:     server.URL,
		TargetPrefix: "Animals",
		JSONVersion:  "1.1",
	}

	req := fakeJSONRequest{Name: "Penny"}
	var resp fakeJSONResponse
	err := client.Do("PetTheDog", "POST", "/", req, &resp)
	if err == nil {
		t.Fatal("Expected an error but none was returned")
	}

	if err, ok := err.(aws.APIError); ok {
		if v, want := err.Type, "Problem"; v != want {
			t.Errorf("Error type was %v, but expected %v", v, want)
		}

		if v, want := err.Message, "What even"; v != want {
			t.Errorf("Error message was %v, but expected %v", v, want)
		}
	} else {
		t.Errorf("Unknown error returned: %#v", err)
	}
}

type fakeJSONRequest struct {
	Name string
}

type fakeJSONResponse struct {
	TailWagged bool
}
