package awstesting_test

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockInput struct {
	_ struct{} `type:"structure"`

	HeaderBoolean *bool `location:"header" locationName:"Header-Boolean" type:"boolean"`

	HeaderDouble *float64 `location:"header" locationName:"Header-Double" type:"double"`

	HeaderString *string `location:"header" locationName:"Header-String" type:"string"`

	QueryString *string `location:"querystring" locationName:"string" type:"string"`

	String_ *string `type:"string"`

	// UriPath is a required field
	UriPath *string `location:"uri" locationName:"second" type:"string" required:"true"`

	// UriPathSegment is a required field
	UriPathSegment *string `location:"uri" locationName:"first" type:"string" required:"true"`
}

type MockOutput struct {
	_ struct{} `type:"structure"`

	HeaderBoolean *bool `location:"header" locationName:"Header-Boolean" type:"boolean"`

	HeaderDouble *float64 `location:"header" locationName:"Header-Double" type:"double"`

	HttpStatusCode *int64 `location:"statusCode" type:"integer"`
}

type MockClient struct {
	*client.Client
}

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Used for custom client initialization logic
var initClient func(*client.Client)

// newRequest creates a new request for a SampleResetXmlProtocolService operation and runs any
// custom request initialization.
func (c *MockClient) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

func New(p client.ConfigProvider, cfgs ...*aws.Config) *MockClient {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// Service information constants
const (
	ServiceName = "Rest Xml Protocol Service" // Name of service.
	EndpointsID = "rest-xml-svc"              // ID to lookup a service endpoint with.
	ServiceID   = "Rest Xml Protocol Service" // ServiceID is a unique identifer of a specific service.
)

func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *MockClient {
	svc := &MockClient{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
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

func (c MockClient) MockRequest(input *MockInput) (req *request.Request, output *MockOutput) {
	op := &request.Operation{
		Name:       "",
		HTTPMethod: "PUT",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MockInput{}
	}

	output = &MockOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restxml.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

func CreateRequest(input *MockInput) *request.Request {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("akid", "secret", "")}))
	svc := New(sess)
	req, resp := svc.MockRequest(input)
	_ = resp

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

type awsError awserr.Error

type RequestError struct {
	awsError
	statusCode int
	requestID  string
	bytes      []byte
}

func NewRequestError(err awsError, statusCode int, requestID string) *RequestError {
	return &RequestError{
		awsError:   err,
		statusCode: statusCode,
		requestID:  requestID,
	}
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
			if !awstesting.StringEqual(t, c.expectString, c.actualString) {
				t.Errorf("input and output time don't match for %s case", name)
			}
		})
	}
}

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
			req: CreateRequest(&MockInput{
				QueryString: aws.String("string-value"),
			}),
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
			req: CreateRequest(&MockInput{
				HeaderBoolean: aws.Bool(true),
				HeaderDouble:  aws.Float64(123.456),
			}),
			expectHeader: map[string]interface{}{
				"Header-Boolean": "true",
				"Header-Double":  "123.456",
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
