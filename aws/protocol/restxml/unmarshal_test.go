package restxml_test

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/restxml"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"

	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/awslabs/aws-sdk-go/internal/protocol/xml/xmlutil"
	"github.com/awslabs/aws-sdk-go/internal/util"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
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
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice1protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &OutputService1ProtocolTest{service}
}

// OutputService1TestCaseOperation1Request generates a request for the OutputService1TestCaseOperation1 operation.
func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1Request() (req *aws.Request, output *OutputService1TestShapeOutputService1TestShapeOutputShape) {
	if opOutputService1TestCaseOperation1 == nil {
		opOutputService1TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService1TestCaseOperation1, nil, output)
	output = &OutputService1TestShapeOutputService1TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1() (output *OutputService1TestShapeOutputService1TestShapeOutputShape, err error) {
	req, out := c.OutputService1TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService1TestCaseOperation1 *aws.Operation

type OutputService1TestShapeOutputService1TestShapeOutputShape struct {
	Char              *string  `type:"character"`
	Double            *float64 `type:"double"`
	FalseBool         *bool    `type:"boolean"`
	Float             *float32 `type:"float"`
	ImaHeader         *string  `location:"header" type:"string"`
	ImaHeaderLocation *string  `location:"header" locationName:"X-Foo" type:"string"`
	Long              *int64   `type:"long"`
	Num               *int     `locationName:"FooNum" type:"integer"`
	Str               *string  `type:"string"`
	TrueBool          *bool    `type:"boolean"`

	metadataOutputService1TestShapeOutputService1TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService1TestShapeOutputService1TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
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
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice2protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

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

type OutputService2TestShapeOutputShape struct {
	Blob []byte `type:"blob"`

	metadataOutputService2TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService2TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
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
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice3protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

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
	ListMember []*string `type:"list"`

	metadataOutputService3TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService3TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
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
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice4protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

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
	ListMember []*string `locationNameList:"item" type:"list"`

	metadataOutputService4TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService4TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
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
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice5protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

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
	ListMember []*string `type:"list" flattened:"true"`

	metadataOutputService5TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService5TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// OutputService6ProtocolTest is a client for OutputService6ProtocolTest.
type OutputService6ProtocolTest struct {
	*aws.Service
}

type OutputService6ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService6ProtocolTest client.
func NewOutputService6ProtocolTest(config *OutputService6ProtocolTestConfig) *OutputService6ProtocolTest {
	if config == nil {
		config = &OutputService6ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice6protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &OutputService6ProtocolTest{service}
}

// OutputService6TestCaseOperation1Request generates a request for the OutputService6TestCaseOperation1 operation.
func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1Request() (req *aws.Request, output *OutputService6TestShapeOutputShape) {
	if opOutputService6TestCaseOperation1 == nil {
		opOutputService6TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService6TestCaseOperation1, nil, output)
	output = &OutputService6TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1() (output *OutputService6TestShapeOutputShape, err error) {
	req, out := c.OutputService6TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService6TestCaseOperation1 *aws.Operation

type OutputService6TestShapeOutputShape struct {
	Map *map[string]*OutputService6TestShapeSingleStructure `type:"map"`

	metadataOutputService6TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService6TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService6TestShapeSingleStructure struct {
	Foo *string `locationName:"foo" type:"string"`

	metadataOutputService6TestShapeSingleStructure `json:"-", xml:"-"`
}

type metadataOutputService6TestShapeSingleStructure struct {
	SDKShapeTraits bool `type:"structure"`
}

// OutputService7ProtocolTest is a client for OutputService7ProtocolTest.
type OutputService7ProtocolTest struct {
	*aws.Service
}

type OutputService7ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService7ProtocolTest client.
func NewOutputService7ProtocolTest(config *OutputService7ProtocolTestConfig) *OutputService7ProtocolTest {
	if config == nil {
		config = &OutputService7ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice7protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &OutputService7ProtocolTest{service}
}

// OutputService7TestCaseOperation1Request generates a request for the OutputService7TestCaseOperation1 operation.
func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1Request() (req *aws.Request, output *OutputService7TestShapeOutputShape) {
	if opOutputService7TestCaseOperation1 == nil {
		opOutputService7TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService7TestCaseOperation1, nil, output)
	output = &OutputService7TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1() (output *OutputService7TestShapeOutputShape, err error) {
	req, out := c.OutputService7TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService7TestCaseOperation1 *aws.Operation

type OutputService7TestShapeOutputShape struct {
	Map *map[string]*string `type:"map" flattened:"true"`

	metadataOutputService7TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService7TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// OutputService8ProtocolTest is a client for OutputService8ProtocolTest.
type OutputService8ProtocolTest struct {
	*aws.Service
}

type OutputService8ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService8ProtocolTest client.
func NewOutputService8ProtocolTest(config *OutputService8ProtocolTestConfig) *OutputService8ProtocolTest {
	if config == nil {
		config = &OutputService8ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice8protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &OutputService8ProtocolTest{service}
}

// OutputService8TestCaseOperation1Request generates a request for the OutputService8TestCaseOperation1 operation.
func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1Request() (req *aws.Request, output *OutputService8TestShapeOutputShape) {
	if opOutputService8TestCaseOperation1 == nil {
		opOutputService8TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService8TestCaseOperation1, nil, output)
	output = &OutputService8TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1() (output *OutputService8TestShapeOutputShape, err error) {
	req, out := c.OutputService8TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService8TestCaseOperation1 *aws.Operation

type OutputService8TestShapeOutputShape struct {
	Map *map[string]*string `locationNameKey:"foo" locationNameValue:"bar" type:"map"`

	metadataOutputService8TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService8TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// OutputService9ProtocolTest is a client for OutputService9ProtocolTest.
type OutputService9ProtocolTest struct {
	*aws.Service
}

type OutputService9ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService9ProtocolTest client.
func NewOutputService9ProtocolTest(config *OutputService9ProtocolTestConfig) *OutputService9ProtocolTest {
	if config == nil {
		config = &OutputService9ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice9protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &OutputService9ProtocolTest{service}
}

// OutputService9TestCaseOperation1Request generates a request for the OutputService9TestCaseOperation1 operation.
func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1Request() (req *aws.Request, output *OutputService9TestShapeOutputShape) {
	if opOutputService9TestCaseOperation1 == nil {
		opOutputService9TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService9TestCaseOperation1, nil, output)
	output = &OutputService9TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1() (output *OutputService9TestShapeOutputShape, err error) {
	req, out := c.OutputService9TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService9TestCaseOperation1 *aws.Operation

type OutputService9TestShapeOutputShape struct {
	Data   *OutputService9TestShapeSingleStructure `type:"structure"`
	Header *string                                 `location:"header" locationName:"X-Foo" type:"string"`

	metadataOutputService9TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService9TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Data"`
}

type OutputService9TestShapeSingleStructure struct {
	Foo *string `type:"string"`

	metadataOutputService9TestShapeSingleStructure `json:"-", xml:"-"`
}

type metadataOutputService9TestShapeSingleStructure struct {
	SDKShapeTraits bool `type:"structure"`
}

// OutputService10ProtocolTest is a client for OutputService10ProtocolTest.
type OutputService10ProtocolTest struct {
	*aws.Service
}

type OutputService10ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService10ProtocolTest client.
func NewOutputService10ProtocolTest(config *OutputService10ProtocolTestConfig) *OutputService10ProtocolTest {
	if config == nil {
		config = &OutputService10ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice10protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &OutputService10ProtocolTest{service}
}

// OutputService10TestCaseOperation1Request generates a request for the OutputService10TestCaseOperation1 operation.
func (c *OutputService10ProtocolTest) OutputService10TestCaseOperation1Request() (req *aws.Request, output *OutputService10TestShapeOutputShape) {
	if opOutputService10TestCaseOperation1 == nil {
		opOutputService10TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService10TestCaseOperation1, nil, output)
	output = &OutputService10TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService10ProtocolTest) OutputService10TestCaseOperation1() (output *OutputService10TestShapeOutputShape, err error) {
	req, out := c.OutputService10TestCaseOperation1Request()
	output = out
	err = req.Send()
	return
}

var opOutputService10TestCaseOperation1 *aws.Operation

type OutputService10TestShapeOutputShape struct {
	Stream []byte `type:"blob"`

	metadataOutputService10TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService10TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Stream"`
}

//
// Tests begin here
//

func TestOutputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	svc := NewOutputService1ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><Str>myname</Str><FooNum>123</FooNum><FalseBool>false</FalseBool><TrueBool>true</TrueBool><Float>1.2</Float><Double>1.3</Double><Long>200</Long><Char>a</Char></OperationNameResponse>"))
	req, _ := svc.OutputService1TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("ImaHeader", "test")
	req.HTTPResponse.Header.Set("X-Foo", "abc")

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"Char\":\"a\",\"Double\":1.3,\"FalseBool\":false,\"Float\":1.2,\"ImaHeader\":\"test\",\"ImaHeaderLocation\":\"abc\",\"Long\":200,\"Num\":123,\"Str\":\"myname\",\"TrueBool\":true}"), util.Trim(string(jBuf)))
}

func TestOutputService2ProtocolTestBlobCase1(t *testing.T) {
	svc := NewOutputService2ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResult><Blob>dmFsdWU=</Blob></OperationNameResult>"))
	req, _ := svc.OutputService2TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"Blob\":\"dmFsdWU=\"}"), util.Trim(string(jBuf)))
}

func TestOutputService3ProtocolTestListsCase1(t *testing.T) {
	svc := NewOutputService3ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResult><ListMember><member>abc</member><member>123</member></ListMember></OperationNameResult>"))
	req, _ := svc.OutputService3TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"ListMember\":[\"abc\",\"123\"]}"), util.Trim(string(jBuf)))
}

func TestOutputService4ProtocolTestListWithCustomMemberNameCase1(t *testing.T) {
	svc := NewOutputService4ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResult><ListMember><item>abc</item><item>123</item></ListMember></OperationNameResult>"))
	req, _ := svc.OutputService4TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"ListMember\":[\"abc\",\"123\"]}"), util.Trim(string(jBuf)))
}

func TestOutputService5ProtocolTestFlattenedListCase1(t *testing.T) {
	svc := NewOutputService5ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResult><ListMember>abc</ListMember><ListMember>123</ListMember></OperationNameResult>"))
	req, _ := svc.OutputService5TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"ListMember\":[\"abc\",\"123\"]}"), util.Trim(string(jBuf)))
}

func TestOutputService6ProtocolTestNormalMapCase1(t *testing.T) {
	svc := NewOutputService6ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResult><Map><entry><key>qux</key><value><foo>bar</foo></value></entry><entry><key>baz</key><value><foo>bam</foo></value></entry></Map></OperationNameResult>"))
	req, _ := svc.OutputService6TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"Map\":{\"baz\":{\"Foo\":\"bam\"},\"qux\":{\"Foo\":\"bar\"}}}"), util.Trim(string(jBuf)))
}

func TestOutputService7ProtocolTestFlattenedMapCase1(t *testing.T) {
	svc := NewOutputService7ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResult><Map><key>qux</key><value>bar</value></Map><Map><key>baz</key><value>bam</value></Map></OperationNameResult>"))
	req, _ := svc.OutputService7TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"Map\":{\"baz\":\"bam\",\"qux\":\"bar\"}}"), util.Trim(string(jBuf)))
}

func TestOutputService8ProtocolTestNamedMapCase1(t *testing.T) {
	svc := NewOutputService8ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResult><Map><entry><foo>qux</foo><bar>bar</bar></entry><entry><foo>baz</foo><bar>bam</bar></entry></Map></OperationNameResult>"))
	req, _ := svc.OutputService8TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"Map\":{\"baz\":\"bam\",\"qux\":\"bar\"}}"), util.Trim(string(jBuf)))
}

func TestOutputService9ProtocolTestXMLPayloadCase1(t *testing.T) {
	svc := NewOutputService9ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><Foo>abc</Foo></OperationNameResponse>"))
	req, _ := svc.OutputService9TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("X-Foo", "baz")

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"Data\":{\"Foo\":\"abc\"},\"Header\":\"baz\"}"), util.Trim(string(jBuf)))
}

func TestOutputService10ProtocolTestStreamingPayloadCase1(t *testing.T) {
	svc := NewOutputService10ProtocolTest(nil)

	buf := bytes.NewReader([]byte("abc"))
	req, _ := svc.OutputService10TestCaseOperation1Request()
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req)
	restxml.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	jBuf, _ := json.Marshal(req.Data)
	assert.Equal(t, util.Trim("{\"Stream\":\"YWJj\"}"), util.Trim(string(jBuf)))
}

