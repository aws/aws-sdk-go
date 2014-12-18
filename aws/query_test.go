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
		PresentString:      aws.String("string"),
		PresentBoolean:     aws.True(),
		PresentInteger:     aws.Integer(1),
		PresentLong:        aws.Long(2),
		PresentDouble:      aws.Double(1.2),
		PresentFloat:       aws.Float(2.3),
		PresentSlice:       []string{"one", "two"},
		PresentStruct:      &EmbeddedStruct{Value: aws.String("v")},
		PresentStructSlice: []EmbeddedStruct{{Value: aws.String("p")}},
		PresentMap: map[string]EmbeddedStruct{
			"aa": EmbeddedStruct{Value: aws.String("AA")},
			"bb": EmbeddedStruct{Value: aws.String("BB")},
		},
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
		"Action":                            []string{"GetIP"},
		"Version":                           []string{"1.1"},
		"PresentString":                     []string{"string"},
		"PresentBoolean":                    []string{"true"},
		"PresentInteger":                    []string{"1"},
		"PresentLong":                       []string{"2"},
		"PresentDouble":                     []string{"1.2"},
		"PresentFloat":                      []string{"2.3"},
		"PresentSlice.member.1":             []string{"one"},
		"PresentSlice.member.2":             []string{"two"},
		"PresentStruct.Value":               []string{"v"},
		"PresentStructSlice.member.1.Value": []string{"p"},
		"PresentMap.1.Name":                 []string{"aa"},
		"PresentMap.1.Value.Value":          []string{"AA"},
		"PresentMap.2.Name":                 []string{"bb"},
		"PresentMap.2.Value.Value":          []string{"BB"},
	}

	if !reflect.DeepEqual(form, expectedForm) {
		t.Errorf("Post body was \n%s\n but expected \n%s", form.Encode(), expectedForm.Encode())
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

	req := fakeQueryRequest{}
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
	PresentString aws.StringValue `xml:"PresentString"`
	MissingString aws.StringValue `xml:"MissingString"`

	PresentInteger aws.IntegerValue `xml:"PresentInteger"`
	MissingInteger aws.IntegerValue `xml:"MissingInteger"`

	PresentLong aws.LongValue `xml:"PresentLong"`
	MissingLong aws.LongValue `xml:"MissingLong"`

	PresentDouble aws.DoubleValue `xml:"PresentDouble"`
	MissingDouble aws.DoubleValue `xml:"MissingDouble"`

	PresentFloat aws.FloatValue `xml:"PresentFloat"`
	MissingFloat aws.FloatValue `xml:"MissingFloat"`

	PresentBoolean aws.BooleanValue `xml:"PresentBoolean"`
	MissingBoolean aws.BooleanValue `xml:"MissingBoolean"`

	PresentSlice []string `xml:"PresentSlice"`
	MissingSlice []string `xml:"MissingSlice"`

	PresentStructSlice []EmbeddedStruct `xml:"PresentStructSlice"`
	MissingStructSlice []EmbeddedStruct `xml:"MissingStructSlice"`

	PresentMap map[string]EmbeddedStruct `xml:"PresentMap"`
	MissingMap map[string]EmbeddedStruct `xml:"MissingMap"`

	PresentStruct *EmbeddedStruct `xml:"PresentStruct"`
	MissingStruct *EmbeddedStruct `xml:"MissingStruct"`
}

type EmbeddedStruct struct {
	Value aws.StringValue
}

type fakeQueryResponse struct {
	IPAddress string `xml:"IpAddress"`
}
