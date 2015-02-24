package ec2query_test

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/ec2query"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/awslabs/aws-sdk-go/internal/util"
	"github.com/stretchr/testify/assert"
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
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice1protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)

	return &InputService1ProtocolTest{service}
}

// InputService1TestCaseOperation1Request generates a request for the InputService1TestCaseOperation1 operation.
func (c *InputService1ProtocolTest) InputService1TestCaseOperation1Request(input *InputService1TestShapeInputShape) (req *aws.Request) {
	if opInputService1TestCaseOperation1 == nil {
		opInputService1TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
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
	Bar *string `type:"string"`
	Foo *string `type:"string"`

	metadataInputService1TestShapeInputShape
}

type metadataInputService1TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService2ProtocolTest is a client for InputService2ProtocolTest.
type InputService2ProtocolTest struct {
	*aws.Service
}

type InputService2ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService2ProtocolTest client.
func NewInputService2ProtocolTest(config *InputService2ProtocolTestConfig) *InputService2ProtocolTest {
	if config == nil {
		config = &InputService2ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice2protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)

	return &InputService2ProtocolTest{service}
}

// InputService2TestCaseOperation1Request generates a request for the InputService2TestCaseOperation1 operation.
func (c *InputService2ProtocolTest) InputService2TestCaseOperation1Request(input *InputService2TestShapeInputShape) (req *aws.Request) {
	if opInputService2TestCaseOperation1 == nil {
		opInputService2TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opInputService2TestCaseOperation1, input, nil)

	return
}

func (c *InputService2ProtocolTest) InputService2TestCaseOperation1(input *InputService2TestShapeInputShape) (err error) {
	req := c.InputService2TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService2TestCaseOperation1 *aws.Operation

type InputService2TestShapeInputShape struct {
	Bar  *string `locationName:"barLocationName" type:"string"`
	Foo  *string `type:"string"`
	Yuck *string `locationName:"yuckLocationName" queryName:"yuckQueryName" type:"string"`

	metadataInputService2TestShapeInputShape
}

type metadataInputService2TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService3ProtocolTest is a client for InputService3ProtocolTest.
type InputService3ProtocolTest struct {
	*aws.Service
}

type InputService3ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService3ProtocolTest client.
func NewInputService3ProtocolTest(config *InputService3ProtocolTestConfig) *InputService3ProtocolTest {
	if config == nil {
		config = &InputService3ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice3protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)

	return &InputService3ProtocolTest{service}
}

// InputService3TestCaseOperation1Request generates a request for the InputService3TestCaseOperation1 operation.
func (c *InputService3ProtocolTest) InputService3TestCaseOperation1Request(input *InputService3TestShapeInputShape) (req *aws.Request) {
	if opInputService3TestCaseOperation1 == nil {
		opInputService3TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opInputService3TestCaseOperation1, input, nil)

	return
}

func (c *InputService3ProtocolTest) InputService3TestCaseOperation1(input *InputService3TestShapeInputShape) (err error) {
	req := c.InputService3TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService3TestCaseOperation1 *aws.Operation

type InputService3TestShapeInputShape struct {
	StructArg *InputService3TestShapeStructType `locationName:"Struct" type:"structure"`

	metadataInputService3TestShapeInputShape
}

type metadataInputService3TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService3TestShapeStructType struct {
	ScalarArg *string `locationName:"Scalar" type:"string"`

	metadataInputService3TestShapeStructType
}

type metadataInputService3TestShapeStructType struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService4ProtocolTest is a client for InputService4ProtocolTest.
type InputService4ProtocolTest struct {
	*aws.Service
}

type InputService4ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService4ProtocolTest client.
func NewInputService4ProtocolTest(config *InputService4ProtocolTestConfig) *InputService4ProtocolTest {
	if config == nil {
		config = &InputService4ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice4protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)

	return &InputService4ProtocolTest{service}
}

// InputService4TestCaseOperation1Request generates a request for the InputService4TestCaseOperation1 operation.
func (c *InputService4ProtocolTest) InputService4TestCaseOperation1Request(input *InputService4TestShapeInputShape) (req *aws.Request) {
	if opInputService4TestCaseOperation1 == nil {
		opInputService4TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opInputService4TestCaseOperation1, input, nil)

	return
}

func (c *InputService4ProtocolTest) InputService4TestCaseOperation1(input *InputService4TestShapeInputShape) (err error) {
	req := c.InputService4TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService4TestCaseOperation1 *aws.Operation

type InputService4TestShapeInputShape struct {
	ListArg []*string `type:"list"`

	metadataInputService4TestShapeInputShape
}

type metadataInputService4TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService5ProtocolTest is a client for InputService5ProtocolTest.
type InputService5ProtocolTest struct {
	*aws.Service
}

type InputService5ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService5ProtocolTest client.
func NewInputService5ProtocolTest(config *InputService5ProtocolTestConfig) *InputService5ProtocolTest {
	if config == nil {
		config = &InputService5ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice5protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)

	return &InputService5ProtocolTest{service}
}

// InputService5TestCaseOperation1Request generates a request for the InputService5TestCaseOperation1 operation.
func (c *InputService5ProtocolTest) InputService5TestCaseOperation1Request(input *InputService5TestShapeInputShape) (req *aws.Request) {
	if opInputService5TestCaseOperation1 == nil {
		opInputService5TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opInputService5TestCaseOperation1, input, nil)

	return
}

func (c *InputService5ProtocolTest) InputService5TestCaseOperation1(input *InputService5TestShapeInputShape) (err error) {
	req := c.InputService5TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService5TestCaseOperation1 *aws.Operation

type InputService5TestShapeInputShape struct {
	ListArg []*string `locationName:"ListMemberName" type:"list"`

	metadataInputService5TestShapeInputShape
}

type metadataInputService5TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService6ProtocolTest is a client for InputService6ProtocolTest.
type InputService6ProtocolTest struct {
	*aws.Service
}

type InputService6ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService6ProtocolTest client.
func NewInputService6ProtocolTest(config *InputService6ProtocolTestConfig) *InputService6ProtocolTest {
	if config == nil {
		config = &InputService6ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice6protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)

	return &InputService6ProtocolTest{service}
}

// InputService6TestCaseOperation1Request generates a request for the InputService6TestCaseOperation1 operation.
func (c *InputService6ProtocolTest) InputService6TestCaseOperation1Request(input *InputService6TestShapeInputShape) (req *aws.Request) {
	if opInputService6TestCaseOperation1 == nil {
		opInputService6TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opInputService6TestCaseOperation1, input, nil)

	return
}

func (c *InputService6ProtocolTest) InputService6TestCaseOperation1(input *InputService6TestShapeInputShape) (err error) {
	req := c.InputService6TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService6TestCaseOperation1 *aws.Operation

type InputService6TestShapeInputShape struct {
	ListArg []*string `locationName:"ListMemberName" queryName:"ListQueryName" type:"list"`

	metadataInputService6TestShapeInputShape
}

type metadataInputService6TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService7ProtocolTest is a client for InputService7ProtocolTest.
type InputService7ProtocolTest struct {
	*aws.Service
}

type InputService7ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService7ProtocolTest client.
func NewInputService7ProtocolTest(config *InputService7ProtocolTestConfig) *InputService7ProtocolTest {
	if config == nil {
		config = &InputService7ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice7protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)

	return &InputService7ProtocolTest{service}
}

// InputService7TestCaseOperation1Request generates a request for the InputService7TestCaseOperation1 operation.
func (c *InputService7ProtocolTest) InputService7TestCaseOperation1Request(input *InputService7TestShapeInputShape) (req *aws.Request) {
	if opInputService7TestCaseOperation1 == nil {
		opInputService7TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opInputService7TestCaseOperation1, input, nil)

	return
}

func (c *InputService7ProtocolTest) InputService7TestCaseOperation1(input *InputService7TestShapeInputShape) (err error) {
	req := c.InputService7TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService7TestCaseOperation1 *aws.Operation

type InputService7TestShapeInputShape struct {
	BlobArg []byte `type:"blob"`

	metadataInputService7TestShapeInputShape
}

type metadataInputService7TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

//
// Tests begin here
//

func TestInputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	svc := NewInputService1ProtocolTest(nil)

	var input InputService1TestShapeInputShape
	json.Unmarshal([]byte("{\"Bar\":\"val2\",\"Foo\":\"val1\"}"), &input)
	req := svc.InputService1TestCaseOperation1Request(&input)
	r := req.HTTPRequest

	// build request
	ec2query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("Action=OperationName&Bar=val2&Foo=val1&Version=2014-01-01"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "/", r.URL.Path)

	// assert headers

}

func TestInputService2ProtocolTestStructureWithLocationNameAndQueryNameAppliedToMembersCase1(t *testing.T) {
	svc := NewInputService2ProtocolTest(nil)

	var input InputService2TestShapeInputShape
	json.Unmarshal([]byte("{\"Bar\":\"val2\",\"Foo\":\"val1\",\"Yuck\":\"val3\"}"), &input)
	req := svc.InputService2TestCaseOperation1Request(&input)
	r := req.HTTPRequest

	// build request
	ec2query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("Action=OperationName&BarLocationName=val2&Foo=val1&Version=2014-01-01&yuckQueryName=val3"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "/", r.URL.Path)

	// assert headers

}

func TestInputService3ProtocolTestNestedStructureMembersCase1(t *testing.T) {
	svc := NewInputService3ProtocolTest(nil)

	var input InputService3TestShapeInputShape
	json.Unmarshal([]byte("{\"StructArg\":{\"ScalarArg\":\"foo\"}}"), &input)
	req := svc.InputService3TestCaseOperation1Request(&input)
	r := req.HTTPRequest

	// build request
	ec2query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("Action=OperationName&Struct.Scalar=foo&Version=2014-01-01"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "/", r.URL.Path)

	// assert headers

}

func TestInputService4ProtocolTestListTypesCase1(t *testing.T) {
	svc := NewInputService4ProtocolTest(nil)

	var input InputService4TestShapeInputShape
	json.Unmarshal([]byte("{\"ListArg\":[\"foo\",\"bar\",\"baz\"]}"), &input)
	req := svc.InputService4TestCaseOperation1Request(&input)
	r := req.HTTPRequest

	// build request
	ec2query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("Action=OperationName&ListArg.1=foo&ListArg.2=bar&ListArg.3=baz&Version=2014-01-01"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "/", r.URL.Path)

	// assert headers

}

func TestInputService5ProtocolTestListWithLocationNameAppliedToMemberCase1(t *testing.T) {
	svc := NewInputService5ProtocolTest(nil)

	var input InputService5TestShapeInputShape
	json.Unmarshal([]byte("{\"ListArg\":[\"a\",\"b\",\"c\"]}"), &input)
	req := svc.InputService5TestCaseOperation1Request(&input)
	r := req.HTTPRequest

	// build request
	ec2query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("Action=OperationName&ListMemberName.1=a&ListMemberName.2=b&ListMemberName.3=c&Version=2014-01-01"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "/", r.URL.Path)

	// assert headers

}

func TestInputService6ProtocolTestListWithLocationNameAndQueryNameCase1(t *testing.T) {
	svc := NewInputService6ProtocolTest(nil)

	var input InputService6TestShapeInputShape
	json.Unmarshal([]byte("{\"ListArg\":[\"a\",\"b\",\"c\"]}"), &input)
	req := svc.InputService6TestCaseOperation1Request(&input)
	r := req.HTTPRequest

	// build request
	ec2query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("Action=OperationName&ListQueryName.1=a&ListQueryName.2=b&ListQueryName.3=c&Version=2014-01-01"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "/", r.URL.Path)

	// assert headers

}

func TestInputService7ProtocolTestBase64EncodedBlobsCase1(t *testing.T) {
	svc := NewInputService7ProtocolTest(nil)

	var input InputService7TestShapeInputShape
	json.Unmarshal([]byte("{\"BlobArg\":\"Zm9v\"}"), &input)
	req := svc.InputService7TestCaseOperation1Request(&input)
	r := req.HTTPRequest

	// build request
	ec2query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("Action=OperationName&BlobArg=Zm9v&Version=2014-01-01"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "/", r.URL.Path)

	// assert headers

}
