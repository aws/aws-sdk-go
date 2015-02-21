package jsonrpc_test

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"

	"bytes"
	"encoding/json"
	"github.com/awslabs/aws-sdk-go/internal/util"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var _ bytes.Buffer // always import bytes
var _ http.Request

// InputService1ProtocolTest is a client for InputService1ProtocolTest.
type InputService1ProtocolTest struct {
	*aws.Service
}

type InputService1ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService1ProtocolTest client.
func NewInputService1ProtocolTest(config *InputService1ProtocolTestConfig) *InputService1ProtocolTest {
	if config == nil {
		config = &InputService1ProtocolTestConfig{}
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
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &InputService1ProtocolTest{service}
}

// InputService1TestCaseOperation1Request generates a request for the InputService1TestCaseOperation1 operation.
func (c *InputService1ProtocolTest) InputService1TestCaseOperation1Request(input *InputService1TestShapeInputShape) (req *aws.Request) {
	if opInputService1TestCaseOperation1 == nil {
		opInputService1TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "",
		}
	}

	req = aws.NewRequest(c.Service, opInputService1TestCaseOperation1, input, nil)

	return
}

func (c *InputService1ProtocolTest) InputService1TestCaseOperation1(input *InputService1TestShapeInputShape) (err error) {
	req := c.InputService1TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService1TestCaseOperation1 *aws.Operation

type InputService1TestShapeInputShape struct {
	Name *string `type:"string" json:",omitempty"`

	metadataInputService1TestShapeInputShape
}

type metadataInputService1TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

//
// Tests begin here
//

func TestInputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	svc := NewInputService1ProtocolTest(nil)

	var input InputService1TestShapeInputShape
	json.Unmarshal([]byte("{\"Name\":\"myname\"}"), &input)
	req := svc.InputService1TestCaseOperation1Request(&input)
	r := req.HTTPRequest

	// build request
	jsonrpc.Build(req)

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("{\"Name\": \"myname\"}"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "/", r.URL.Path)

	// assert headers
	assert.Equal(t, "application/x-amz-json-1.1", r.Header.Get("Content-Type"))
	assert.Equal(t, "com.amazonaws.foo.OperationName", r.Header.Get("X-Amz-Target"))

}

