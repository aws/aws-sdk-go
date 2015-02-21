package jsonrpc_test

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"

	"encoding/json"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var reTrim = regexp.MustCompile("\\s")

func trim(s string) string {
	return reTrim.ReplaceAllString(s, "")
}

// Service1ProtocolTest is a client for Service1ProtocolTest.
type Service1ProtocolTest struct {
	*aws.Service
}

type Service1ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new Service1ProtocolTest client.
func NewService1ProtocolTest(config *Service1ProtocolTestConfig) *Service1ProtocolTest {
	if config == nil {
		config = &Service1ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "",
		APIVersion:   "",
		JSONVersion:  "1.1",
		TargetPrefix: "com.amazonaws.foo",
	}
	service.Initialize()

	// Handlers

	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &Service1ProtocolTest{service}
}

// Service1TestCaseOperation0Request generates a request for the Service1TestCaseOperation0 operation.
func (c *Service1ProtocolTest) Service1TestCaseOperation0Request(input *Service1TestShapeInputShape) (req *aws.Request) {
	if opService1TestCaseOperation0 == nil {
		opService1TestCaseOperation0 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "",
		}
	}

	req = aws.NewRequest(c.Service, opService1TestCaseOperation0, input, nil)

	return
}

func (c *Service1ProtocolTest) Service1TestCaseOperation0(input *Service1TestShapeInputShape) (err error) {
	req := c.Service1TestCaseOperation0Request(input)
	err = req.Send()
	return
}

var opService1TestCaseOperation0 *aws.Operation

type Service1TestShapeInputShape struct {
	Name *string `type:"string" json:",omitempty"`

	metadataService1TestShapeInputShape
}

type metadataService1TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

func TestScalarMembersCase1(t *testing.T) {
	svc := NewService1ProtocolTest(nil)

	var input Service1TestShapeInputShape
	json.Unmarshal([]byte("{\"Name\":\"myname\"}"), &input)
	req := svc.Service1TestCaseOperation0Request(&input)
	req.Build()
	r := req.HTTPRequest

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, trim(string(body)), trim("{\"Name\": \"myname\"}"))

	// assert URL
	assert.Equal(t, r.URL.Path, "/")

	// assert headers
	assert.Equal(t, r.Header.Get("Content-Type"), "application/x-amz-json-1.1")
	assert.Equal(t, r.Header.Get("X-Amz-Target"), "com.amazonaws.foo.OperationName")

}
