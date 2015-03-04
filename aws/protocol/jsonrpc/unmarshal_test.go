package jsonrpc_test

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"

	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/awslabs/aws-sdk-go/internal/protocol/xml/xmlutil"
	"github.com/awslabs/aws-sdk-go/internal/util"
	"github.com/stretchr/testify/assert"
)

var _ bytes.Buffer // always import bytes
var _ http.Request
var _ json.Marshaler
var _ time.Time
var _ xmlutil.XMLNode
var _ xml.Attr
var _ = ioutil.Discard

// OutputService1ProtocolTest is a client for OutputService1ProtocolTest.
type OutputService1ProtocolTest struct {
	*aws.Service
}

type OutputService1ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService1ProtocolTest client.
func NewOutputService1ProtocolTest(config *OutputService1ProtocolTestConfig) *OutputService1ProtocolTest {
	if config == nil {
		config = &OutputService1ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "outputservice1protocoltest",
		APIVersion:   "",
		JSONVersion:  "",
		TargetPrefix: "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &OutputService1ProtocolTest{service}
}

// OutputService1TestCaseOperation1Request generates a request for the OutputService1TestCaseOperation1 operation.
func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1Request() (req *aws.Request, output *OutputService1TestShapeOutputShape) {
	if opOutputService1TestCaseOperation1 == nil {
		opOutputService1TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService1TestCaseOperation1, nil, output)
	output = &OutputService1TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1() (output *OutputService1TestShapeOutputShape, err error) {
	req, out := c.OutputService1TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService1TestCaseOperation1 *aws.Operation

type OutputService1TestShapeOutputShape struct {
	Char      *string  `type:"character" json:",omitempty"`
	Double    *float64 `type:"double" json:",omitempty"`
	FalseBool *bool    `type:"boolean" json:",omitempty"`
	Float     *float32 `type:"float" json:",omitempty"`
	Long      *int64   `type:"long" json:",omitempty"`
	Num       *int     `type:"integer" json:",omitempty"`
	Str       *string  `type:"string" json:",omitempty"`
	TrueBool  *bool    `type:"boolean" json:",omitempty"`

	metadataOutputService1TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService1TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

// OutputService2ProtocolTest is a client for OutputService2ProtocolTest.
type OutputService2ProtocolTest struct {
	*aws.Service
}

type OutputService2ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService2ProtocolTest client.
func NewOutputService2ProtocolTest(config *OutputService2ProtocolTestConfig) *OutputService2ProtocolTest {
	if config == nil {
		config = &OutputService2ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "outputservice2protocoltest",
		APIVersion:   "",
		JSONVersion:  "",
		TargetPrefix: "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &OutputService2ProtocolTest{service}
}

// OutputService2TestCaseOperation1Request generates a request for the OutputService2TestCaseOperation1 operation.
func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1Request() (req *aws.Request, output *OutputService2TestShapeOutputShape) {
	if opOutputService2TestCaseOperation1 == nil {
		opOutputService2TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService2TestCaseOperation1, nil, output)
	output = &OutputService2TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1() (output *OutputService2TestShapeOutputShape, err error) {
	req, out := c.OutputService2TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService2TestCaseOperation1 *aws.Operation

type OutputService2TestShapeBlobContainer struct {
	Foo []byte `locationName:"foo" type:"blob" json:",omitempty"`

	metadataOutputService2TestShapeBlobContainer `json:"-", xml:"-"`
}

type metadataOutputService2TestShapeBlobContainer struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type OutputService2TestShapeOutputShape struct {
	BlobMember   []byte                                `type:"blob" json:",omitempty"`
	StructMember *OutputService2TestShapeBlobContainer `type:"structure" json:",omitempty"`

	metadataOutputService2TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService2TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

// OutputService3ProtocolTest is a client for OutputService3ProtocolTest.
type OutputService3ProtocolTest struct {
	*aws.Service
}

type OutputService3ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService3ProtocolTest client.
func NewOutputService3ProtocolTest(config *OutputService3ProtocolTestConfig) *OutputService3ProtocolTest {
	if config == nil {
		config = &OutputService3ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "outputservice3protocoltest",
		APIVersion:   "",
		JSONVersion:  "",
		TargetPrefix: "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &OutputService3ProtocolTest{service}
}

// OutputService3TestCaseOperation1Request generates a request for the OutputService3TestCaseOperation1 operation.
func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1Request() (req *aws.Request, output *OutputService3TestShapeOutputShape) {
	if opOutputService3TestCaseOperation1 == nil {
		opOutputService3TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService3TestCaseOperation1, nil, output)
	output = &OutputService3TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1() (output *OutputService3TestShapeOutputShape, err error) {
	req, out := c.OutputService3TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService3TestCaseOperation1 *aws.Operation

type OutputService3TestShapeOutputShape struct {
	ListMember []*string `type:"list" json:",omitempty"`

	metadataOutputService3TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService3TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

// OutputService4ProtocolTest is a client for OutputService4ProtocolTest.
type OutputService4ProtocolTest struct {
	*aws.Service
}

type OutputService4ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService4ProtocolTest client.
func NewOutputService4ProtocolTest(config *OutputService4ProtocolTestConfig) *OutputService4ProtocolTest {
	if config == nil {
		config = &OutputService4ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "outputservice4protocoltest",
		APIVersion:   "",
		JSONVersion:  "",
		TargetPrefix: "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &OutputService4ProtocolTest{service}
}

// OutputService4TestCaseOperation1Request generates a request for the OutputService4TestCaseOperation1 operation.
func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1Request() (req *aws.Request, output *OutputService4TestShapeOutputShape) {
	if opOutputService4TestCaseOperation1 == nil {
		opOutputService4TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService4TestCaseOperation1, nil, output)
	output = &OutputService4TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1() (output *OutputService4TestShapeOutputShape, err error) {
	req, out := c.OutputService4TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService4TestCaseOperation1 *aws.Operation

type OutputService4TestShapeOutputShape struct {
	MapMember *map[string][]*int `type:"map" json:",omitempty"`

	metadataOutputService4TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService4TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

// OutputService5ProtocolTest is a client for OutputService5ProtocolTest.
type OutputService5ProtocolTest struct {
	*aws.Service
}

type OutputService5ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService5ProtocolTest client.
func NewOutputService5ProtocolTest(config *OutputService5ProtocolTestConfig) *OutputService5ProtocolTest {
	if config == nil {
		config = &OutputService5ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "outputservice5protocoltest",
		APIVersion:   "",
		JSONVersion:  "",
		TargetPrefix: "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &OutputService5ProtocolTest{service}
}

// OutputService5TestCaseOperation1Request generates a request for the OutputService5TestCaseOperation1 operation.
func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1Request() (req *aws.Request, output *OutputService5TestShapeOutputShape) {
	if opOutputService5TestCaseOperation1 == nil {
		opOutputService5TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService5TestCaseOperation1, nil, output)
	output = &OutputService5TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1() (output *OutputService5TestShapeOutputShape, err error) {
	req, out := c.OutputService5TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService5TestCaseOperation1 *aws.Operation

type OutputService5TestShapeOutputShape struct {
	StrType *string `type:"string" json:",omitempty"`

	metadataOutputService5TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService5TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

//
// Tests begin here
//

func TestOutputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	svc := NewOutputService1ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"Str\": \"myname\", \"Num\": 123, \"FalseBool\": false, \"TrueBool\": true, \"Float\": 1.2, \"Double\": 1.3, \"Long\": 200, \"Char\": \"a\"}"))
	req, _ := svc.OutputService1TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf)}

	// unmarshal response
	jsonrpc.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"Char\":\"a\",\"Double\":1.3,\"FalseBool\":false,\"Float\":1.2,\"Long\":200,\"Num\":123,\"Str\":\"myname\",\"TrueBool\":true}"), util.Trim(string(jBuf)))

	// assert headers

}

func TestOutputService2ProtocolTestBlobMembersCase1(t *testing.T) {
	svc := NewOutputService2ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"BlobMember\": \"aGkh\", \"StructMember\": {\"foo\": \"dGhlcmUh\"}}"))
	req, _ := svc.OutputService2TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf)}

	// unmarshal response
	jsonrpc.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"BlobMember\":\"aGkh\",\"StructMember\":{\"Foo\":\"dGhlcmUh\"}}"), util.Trim(string(jBuf)))

	// assert headers

}

func TestOutputService3ProtocolTestListsCase1(t *testing.T) {
	svc := NewOutputService3ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"ListMember\": [\"a\", \"b\"]}"))
	req, _ := svc.OutputService3TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf)}

	// unmarshal response
	jsonrpc.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"ListMember\":[\"a\",\"b\"]}"), util.Trim(string(jBuf)))

	// assert headers

}

func TestOutputService4ProtocolTestMapsCase1(t *testing.T) {
	svc := NewOutputService4ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"MapMember\": {\"a\": [1, 2], \"b\": [3, 4]}}"))
	req, _ := svc.OutputService4TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf)}

	// unmarshal response
	jsonrpc.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"MapMember\":{\"a\":[1,2],\"b\":[3,4]}}"), util.Trim(string(jBuf)))

	// assert headers

}

func TestOutputService5ProtocolTestIgnoresExtraDataCase1(t *testing.T) {
	svc := NewOutputService5ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"foo\": \"bar\"}"))
	req, _ := svc.OutputService5TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf)}

	// unmarshal response
	jsonrpc.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{}"), util.Trim(string(jBuf)))

	// assert headers

}
