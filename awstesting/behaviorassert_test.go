package awstesting_test

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
	"io/ioutil"
	"net/http"
	"testing"
)
type sampleStruct struct{
	A int
	B string
	C float64
	D map[string]string
	E map[string]float64
	F map[string]interface{}
}

func TestStringEqual(t *testing.T) {
	cases := map[string]struct {
		expectString string
		actualString string
	}{
		"Test1": {
			expectString: "hello world",
			actualString: "hello world",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.StringEqual(t, c.expectString, c.actualString){
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

type MockInput struct{
	_ struct{} `type:"structure"`

	Header *string `location:"header" locationName:"Some-Header" type:"string"`

	String_ *string `type:"string"`

	Struct *SimpleStruct `type:"structure"`
}

type SimpleStruct struct {
	_ struct{} `type:"structure"`

	Value *string `type:"string"`
}

type MockOutput struct {
	_ struct{} `type:"structure"`
}

func CreateRequest(input *MockInput) *request.Request {
	c := awstesting.NewClient(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("akid", "secret", ""),
	})
	c.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	c.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	c.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	c.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	c.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	op := &request.Operation{
		Name:       "",
		HTTPMethod: "PUT",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MockInput{}
	}

	output := &MockOutput{}
	req := c.NewRequest(op, input, output)
	/*
		for key, val := range actualHeader {
			req.HTTPRequest.Header.Set(key, reflect.ValueOf(val).String())
		}
	*/
	req.Handlers.Unmarshal.Swap(restxml.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)

	MockHTTPResponseHandler := request.NamedHandler{Name: "core.SendHandler", Fn: func(r *request.Request) {
		r.HTTPResponse = &http.Response{StatusCode: 200,
			Header: http.Header{},
			Body:   ioutil.NopCloser(&bytes.Buffer{}),
		}
	}}
	req.Handlers.Send.Swap(corehandlers.SendHandler.Name, MockHTTPResponseHandler)

	err := req.Send()
	if err != nil {
		panic(err)
	}

	return req
}

func TestAssertRequestHeadersMatch(t *testing.T) {
	cases := map[string]struct {
		req *request.Request
		expectHeader map[string]interface{}
	}{
		"Test1": {
			req: CreateRequest(&MockInput{
				Header:  aws.String("value 1"),
				String_: aws.String("value 2"),
				Struct: &SimpleStruct{
					Value: aws.String("value 3"),
				},
			}),
			expectHeader: map[string]interface{}{
				"String_": "value 2",
				"Header": "value 1",
			},

		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			mockT := &testing.T{}
			if !awstesting.AssertRequestHeadersMatch(mockT, c.expectHeader, c.req){
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertResponseErrorIsKindOf(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		err error
	}{
		"Test1": {
			expectVal: CreateRequest(&MockInput{
				Header:  aws.String("value 1"),
				String_: aws.String("value 2"),
				Struct: &SimpleStruct{
					Value: aws.String("value 3"),
				},
			}),
			err: errors.New()

		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			mockT := &testing.T{}
			if !awstesting.AssertResponseErrorIsKindOf(mockT, c.expectVal, c.err){
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}
func TestAssertResponseDataEquals(t *testing.T) {
	cases := map[string]struct {
		expectStruct sampleStruct
		actualStruct sampleStruct
	}{
		"Test1": {
			expectStruct: sampleStruct{
				A: 1,
				B: "hey",
				C: 32,
				D: map[string]string{"hello": "world"},
			},
			actualStruct: sampleStruct{
				A: 1,
				B: "hey",
				C: 3.000,
				D: map[string]string{"hello": "world"},
			},

		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertResponseDataEquals(t, c.expectStruct, c.actualStruct){
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}