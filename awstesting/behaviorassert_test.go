package awstesting_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/request"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Used for custom client initialization logic
var initClient func(*client.Client)

// Error types are copied from awserr as they are unexported
type awsError awserr.Error

type RequestError struct {
	awsError
	statusCode int
	requestID  string
	bytes      []byte
}

type MockService struct {
	*client.Client
}

// MockInput1 is doesn't have payload defined and is used
// as input for all assertion unit tests except "requestBodyEqualsBytes"
type MockInput1 struct {
	_ struct{} `type:"structure"`

	//BodyStream io.ReadSeeker `type:"blob"`

	Data *SimpleStruct `locationName:"DataNode" type:"structure" xmlURI:"http://xml/ns"`

	// HeaderBlob is automatically base64 encoded/decoded by the SDK.
	HeaderBlob []byte `location:"header" locationName:"Header-Binary" type:"blob"`

	HeaderBoolean *bool `location:"header" locationName:"Header-Boolean" type:"boolean"`

	HeaderDouble *float64 `location:"header" locationName:"Header-Double" type:"double"`

	HeaderJsonValue aws.JSONValue `location:"header" locationName:"Header-Json-Value" type:"jsonvalue"`

	HeaderString *string `location:"header" locationName:"Header-String" type:"string"`

	// QueryBlob is automatically base64 encoded/decoded by the SDK.
	QueryBlob []byte `location:"querystring" locationName:"binary-value" type:"blob"`

	QueryString *string `location:"querystring" locationName:"string" type:"string"`

	String_ *string `type:"string"`

	// UriPath is a required field
	UriPath *string `location:"uri" locationName:"second" type:"string" required:"true"`

	// UriPathSegment is a required field
	UriPathSegment *string `location:"uri" locationName:"first" type:"string" required:"true"`
}

// MockInput2 has payload defined as 'BodyStream' and is used as
// as input for "requestBodyEqualsBytes" assertion unit test
type MockInput2 struct {
	_ struct{} `type:"structure" payload:"BodyStream"`

	BodyStream io.ReadSeeker `type:"blob"`

	// UriPath is a required field
	UriPath *string `location:"uri" locationName:"second" type:"string" required:"true"`

	// UriPathSegment is a required field
	UriPathSegment *string `location:"uri" locationName:"first" type:"string" required:"true"`
}

type SimpleStruct struct {
	_ struct{} `type:"structure"`

	Value *string `type:"string"`
}

type MockOutput struct {
	_ struct{} `type:"structure"`

	// HeaderBlob is automatically base64 encoded/decoded by the SDK.
	HeaderBlob []byte `location:"header" locationName:"Header-Binary" type:"blob"`

	HeaderBoolean *bool `location:"header" locationName:"Header-Boolean" type:"boolean"`

	HeaderDouble *float64 `location:"header" locationName:"Header-Double" type:"double"`

	HttpStatusCode *int64 `location:"statusCode" type:"integer"`
}

// newRequest creates a new request for a MockService operation and runs any
// custom request initialization.
func (c *MockService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)
	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}
	return req
}

// NewRestXML creates a new instance of the MockService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
func NewRestXML(p client.ConfigProvider, cfgs ...*aws.Config) *MockService {
	c := p.ClientConfig("rest-xml-svc", cfgs...)
	return newRestXMLClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newRestXMLClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *MockService {
	svc := &MockService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "Rest Xml Protocol Service",
				ServiceID:     "Rest Xml Protocol Service",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// NewJSONRpc creates a new instance of the MockService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
func NewJSONRpc(p client.ConfigProvider, cfgs ...*aws.Config) *MockService {
	c := p.ClientConfig("endpoint-prefix", cfgs...)
	return newJSONRpcClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newJSONRpcClient creates, initializes and returns a new service client instance.
func newJSONRpcClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *MockService {
	svc := &MockService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "Json One One Protocol Service",
				ServiceID:     "Json One One Protocol Service",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
				JSONVersion:   "1.1",
				TargetPrefix:  "JsonProtocolService_20180101",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// CreateRequest1 generates a "aws/request.Request". The "output" return value
// will be populated with the request's response once the request completes
// successfully. The input here is MockInput1
// method is the HTTPMethod and requestUri is the HTTPPath
func (c MockService) CreateRequest1(input *MockInput1, method string, requestUri string) (req *request.Request, output *MockOutput) {
	op := &request.Operation{
		Name:       "",
		HTTPMethod: method,
		HTTPPath:   requestUri,
	}

	if input == nil {
		input = &MockInput1{}
	}

	output = &MockOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restxml.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// BuildRequest1 creates the request, stubs out a mock response by
// replacing the send handler with a custom handler and outputs
// request and response
func BuildRequest1(input *MockInput1, method string, requestUri string, clientType string) (req *request.Request, resp *MockOutput) {
	sess := unit.Session
	svc := NewRestXML(sess)
	if clientType == "json" {
		svc = NewJSONRpc(sess)
	}
	req, resp = svc.CreateRequest1(input, method, requestUri)
	_ = resp

	MockHTTPResponseHandler := request.NamedHandler{Name: "core.SendHandler", Fn: func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(&bytes.Buffer{}),
		}
	}}
	req.Handlers.Send.Swap(corehandlers.SendHandler.Name, MockHTTPResponseHandler)

	err := req.Send()
	if err != nil {
		panic(err)
	}
	return
}

// GetRequest1 returns the request by calling BuildRequest1
func GetRequest1(input *MockInput1, method string, requestUri string, clientType string) *request.Request {
	req, _ := BuildRequest1(input, method, requestUri, clientType)
	return req
}

// GetResponse1 returns the response by calling BuildRequest1
func GetResponse1(input *MockInput1, method string, requestUri string, clientType string) *MockOutput {
	_, resp := BuildRequest1(input, method, requestUri, clientType)
	return resp
}

// CreateRequest2 generates a "aws/request.Request". The "output" return value
// will be populated with the request's response once the request completes
// successfully. The input here is MockInput2
// method is the HTTPMethod and requestUri is the HTTPPath
func (c MockService) CreateRequest2(input *MockInput2, method string, requestUri string) (req *request.Request, output *MockOutput) {
	op := &request.Operation{
		Name:       "",
		HTTPMethod: method,
		HTTPPath:   requestUri,
	}

	if input == nil {
		input = &MockInput2{}
	}

	output = &MockOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restxml.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// BuildRequest2 creates the request, stubs out a mock response by
// replacing the send handler with a custom handler and outputs
// request and response
func BuildRequest2(input *MockInput2, method string, requestUri string, clientType string) (req *request.Request, resp *MockOutput) {
	sess := unit.Session
	svc := NewRestXML(sess)
	if clientType == "json" {
		svc = NewJSONRpc(sess)
	}
	req, resp = svc.CreateRequest2(input, method, requestUri)
	_ = resp

	MockHTTPResponseHandler := request.NamedHandler{Name: "core.SendHandler", Fn: func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(&bytes.Buffer{}),
		}
	}}
	req.Handlers.Send.Swap(corehandlers.SendHandler.Name, MockHTTPResponseHandler)

	err := req.Send()
	if err != nil {
		panic(err)
	}
	return
}

// GetRequest2 returns the request by calling BuildRequest2
func GetRequest2(input *MockInput2, method string, requestUri string, clientType string) *request.Request {
	req, _ := BuildRequest2(input, method, requestUri, clientType)
	return req
}

// GetResponse2 returns the response by calling BuildRequest2
func GetResponse2(input *MockInput2, method string, requestUri string, clientType string) *MockOutput {
	_, resp := BuildRequest2(input, method, requestUri, clientType)
	return resp
}

func NewRequestError(err awsError, statusCode int, requestID string) *RequestError {
	return &RequestError{
		awsError:   err,
		statusCode: statusCode,
		requestID:  requestID,
	}
}

// Unit Test for assertions start here
func TestAssertRequestURLMatches(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		actualVal string
	}{
		"Test1": {
			expectVal: "https://inside.amazon.com/",
			actualVal: "https://inside.amazon.com/",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertRequestURLMatches(t, c.expectVal, c.actualVal) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertRequestURLQueryMatches(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		req       *request.Request
	}{
		"Test1": {
			expectVal: "string=string-value",
			req: GetRequest1(&MockInput1{
				QueryString: aws.String("string-value"),
			}, "PUT", "/", "xml"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertRequestURLQueryMatches(t, c.expectVal, c.req) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertRequestHeadersMatch(t *testing.T) {
	cases := map[string]struct {
		req          *request.Request
		expectHeader map[string]interface{}
	}{
		"Test1": {
			req: GetRequest1(&MockInput1{
				HeaderBoolean: aws.Bool(true),
				HeaderDouble:  aws.Float64(123.456),
			}, "PUT", "/", "xml"),
			expectHeader: map[string]interface{}{
				"Header-Boolean": "true",
				"Header-Double":  "123.456",
			},
		},
		"Test2": {
			req: GetRequest1(&MockInput1{
				HeaderJsonValue: aws.JSONValue{"array": []interface{}{1, 2, 3, 4}, "boolFalse": false, "boolTrue": true, "null": interface{}(nil), "number": 1234.5, "object": map[string]interface{}{"key": "value"}, "string": "value"},
			}, "PUT", "/", "xml"),
			expectHeader: map[string]interface{}{
				"Header-Json-Value": "eyJzdHJpbmciOiJ2YWx1ZSIsIm51bWJlciI6MTIzNC41LCJib29sVHJ1ZSI6dHJ1ZSwiYm9vbEZhbHNlIjpmYWxzZSwiYXJyYXkiOlsxLDIsMyw0XSwib2JqZWN0Ijp7ImtleSI6InZhbHVlIn0sIm51bGwiOm51bGx9",
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertRequestHeadersMatch(t, c.expectHeader, c.req) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertRequestBodyEqualsBytes(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		req       *request.Request
	}{
		"Test1": {
			expectVal: "YmluYXJ5LXZhbHVl",
			req: GetRequest2(&MockInput2{
				BodyStream:     aws.ReadSeekCloser(strings.NewReader("YmluYXJ5LXZhbHVl")),
				UriPath:        aws.String("path"),
				UriPathSegment: aws.String("segment"),
			}, "PUT", "/{first}/{second+}", "xml"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			mockT := &testing.T{}
			if !awstesting.AssertRequestBodyEqualsBytes(mockT, c.expectVal, c.req) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertRequestBodyEqualsJSON(t *testing.T) {
	cases := map[string]struct {
		expectVal map[string]interface{}
		req       *request.Request
	}{
		"Test1": {
			expectVal: map[string]interface{}{"String_": "abc xyz"},
			req: GetRequest1(&MockInput1{
				String_: aws.String("abc xyz"),
			}, "POST", "/", "json"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertRequestBodyEqualsJSON(t, c.expectVal, c.req) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertRequestBodyMatchesXML(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		req       *request.Request
	}{
		"Test1": {
			expectVal: "<DataNode xmlns=\"http://xml/ns\"><Value>string value</Value></DataNode>",
			req: GetRequest1(&MockInput1{
				Data: &SimpleStruct{
					Value: aws.String("string value"),
				},
			}, "PUT", "/", "xml"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertRequestBodyMatchesXML(t, c.expectVal, c.req, MockInput1{}) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertRequestBodyEqualsString(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		req       *request.Request
	}{
		"Test1": {
			expectVal: "",
			req:       GetRequest1(&MockInput1{}, "PUT", "/", "xml"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertRequestBodyEqualsString(t, c.expectVal, c.req, MockInput1{}) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertResponseDataEquals(t *testing.T) {
	cases := map[string]struct {
		expectStruct MockOutput
		actualStruct MockOutput
	}{
		"Test1": {
			expectStruct: MockOutput{
				HeaderBoolean:  aws.Bool(true),
				HeaderDouble:   aws.Float64(123.456),
				HttpStatusCode: aws.Int64(200),
			},
			actualStruct: MockOutput{
				HeaderBoolean:  aws.Bool(true),
				HeaderDouble:   aws.Float64(123.456000),
				HttpStatusCode: aws.Int64(200),
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertResponseDataEquals(t, c.expectStruct, c.actualStruct) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertResponseErrorIsKindOf(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		err       error
	}{
		"Test1": {
			expectVal: "ErrorWithoutMembers",
			err:       NewRequestError(awserr.New("ErrorWithoutMembers", "", nil), 500, ""),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertResponseErrorIsKindOf(t, c.expectVal, c.err) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertResponseErrorMessageEquals(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		err       error
	}{
		"Test1": {
			expectVal: "Something went wrong",
			err:       NewRequestError(awserr.New("ErrorWithMembers", "Something went wrong", nil), 500, ""),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertResponseErrorMessageEquals(t, c.expectVal, c.err) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

func TestAssertResponseErrorRequestIDEquals(t *testing.T) {
	cases := map[string]struct {
		expectVal string
		err       error
	}{
		"Test1": {
			expectVal: "amazon-uniq-request-id",
			err:       NewRequestError(awserr.New("ErrorWithMembers", "Something went wrong", nil), 500, "amazon-uniq-request-id"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if !awstesting.AssertResponseErrorRequestIDEquals(t, c.expectVal, c.err) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}
