package restjson_test

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/restjson"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"

	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/awslabs/aws-sdk-go/internal/protocol/xml/xmlutil"
	"github.com/awslabs/aws-sdk-go/internal/util"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
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
var _ = util.Trim("")
var _ = url.Values{}

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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService1ProtocolTest{service}
}

// OutputService1TestCaseOperation1Request generates a request for the OutputService1TestCaseOperation1 operation.
func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1Request(input *OutputService1TestShapeOutputService1TestCaseOperation1Input) (req *aws.Request, output *OutputService1TestShapeOutputShape) {
	if opOutputService1TestCaseOperation1 == nil {
		opOutputService1TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService1TestCaseOperation1, input, output)
	output = &OutputService1TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1(input *OutputService1TestShapeOutputService1TestCaseOperation1Input) (output *OutputService1TestShapeOutputShape, err error) {
	req, out := c.OutputService1TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService1TestCaseOperation1 *aws.Operation

type OutputService1TestShapeOutputService1TestCaseOperation1Input struct {
	metadataOutputService1TestShapeOutputService1TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService1TestShapeOutputService1TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService1TestShapeOutputShape struct {
	Char *string `type:"character"`

	Double *float64 `type:"double"`

	FalseBool *bool `type:"boolean"`

	Float *float64 `type:"float"`

	ImaHeader *string `location:"header" type:"string"`

	ImaHeaderLocation *string `location:"header" locationName:"X-Foo" type:"string"`

	Long *int64 `type:"long"`

	Num *int64 `type:"integer"`

	Status *int64 `location:"statusCode" type:"integer"`

	Str *string `type:"string"`

	TrueBool *bool `type:"boolean"`

	metadataOutputService1TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService1TestShapeOutputShape struct {
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService2ProtocolTest{service}
}

// OutputService2TestCaseOperation1Request generates a request for the OutputService2TestCaseOperation1 operation.
func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1Request(input *OutputService2TestShapeOutputService2TestCaseOperation1Input) (req *aws.Request, output *OutputService2TestShapeOutputShape) {
	if opOutputService2TestCaseOperation1 == nil {
		opOutputService2TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService2TestCaseOperation1, input, output)
	output = &OutputService2TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1(input *OutputService2TestShapeOutputService2TestCaseOperation1Input) (output *OutputService2TestShapeOutputShape, err error) {
	req, out := c.OutputService2TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService2TestCaseOperation1 *aws.Operation

type OutputService2TestShapeBlobContainer struct {
	Foo []byte `locationName:"foo" type:"blob"`

	metadataOutputService2TestShapeBlobContainer `json:"-", xml:"-"`
}

type metadataOutputService2TestShapeBlobContainer struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService2TestShapeOutputService2TestCaseOperation1Input struct {
	metadataOutputService2TestShapeOutputService2TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService2TestShapeOutputService2TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService2TestShapeOutputShape struct {
	BlobMember []byte `type:"blob"`

	StructMember *OutputService2TestShapeBlobContainer `type:"structure"`

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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService3ProtocolTest{service}
}

// OutputService3TestCaseOperation1Request generates a request for the OutputService3TestCaseOperation1 operation.
func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1Request(input *OutputService3TestShapeOutputService3TestCaseOperation1Input) (req *aws.Request, output *OutputService3TestShapeOutputShape) {
	if opOutputService3TestCaseOperation1 == nil {
		opOutputService3TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService3TestCaseOperation1, input, output)
	output = &OutputService3TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1(input *OutputService3TestShapeOutputService3TestCaseOperation1Input) (output *OutputService3TestShapeOutputShape, err error) {
	req, out := c.OutputService3TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService3TestCaseOperation1 *aws.Operation

type OutputService3TestShapeOutputService3TestCaseOperation1Input struct {
	metadataOutputService3TestShapeOutputService3TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService3TestShapeOutputService3TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService3TestShapeOutputShape struct {
	StructMember *OutputService3TestShapeTimeContainer `type:"structure"`

	TimeMember *time.Time `type:"timestamp" timestampFormat:"unix"`

	metadataOutputService3TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService3TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService3TestShapeTimeContainer struct {
	Foo *time.Time `locationName:"foo" type:"timestamp" timestampFormat:"unix"`

	metadataOutputService3TestShapeTimeContainer `json:"-", xml:"-"`
}

type metadataOutputService3TestShapeTimeContainer struct {
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService4ProtocolTest{service}
}

// OutputService4TestCaseOperation1Request generates a request for the OutputService4TestCaseOperation1 operation.
func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1Request(input *OutputService4TestShapeOutputService4TestCaseOperation1Input) (req *aws.Request, output *OutputService4TestShapeOutputShape) {
	if opOutputService4TestCaseOperation1 == nil {
		opOutputService4TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService4TestCaseOperation1, input, output)
	output = &OutputService4TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1(input *OutputService4TestShapeOutputService4TestCaseOperation1Input) (output *OutputService4TestShapeOutputShape, err error) {
	req, out := c.OutputService4TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService4TestCaseOperation1 *aws.Operation

type OutputService4TestShapeOutputService4TestCaseOperation1Input struct {
	metadataOutputService4TestShapeOutputService4TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService4TestShapeOutputService4TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService4TestShapeOutputShape struct {
	ListMember []*string `type:"list"`

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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService5ProtocolTest{service}
}

// OutputService5TestCaseOperation1Request generates a request for the OutputService5TestCaseOperation1 operation.
func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1Request(input *OutputService5TestShapeOutputService5TestCaseOperation1Input) (req *aws.Request, output *OutputService5TestShapeOutputShape) {
	if opOutputService5TestCaseOperation1 == nil {
		opOutputService5TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService5TestCaseOperation1, input, output)
	output = &OutputService5TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1(input *OutputService5TestShapeOutputService5TestCaseOperation1Input) (output *OutputService5TestShapeOutputShape, err error) {
	req, out := c.OutputService5TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService5TestCaseOperation1 *aws.Operation

type OutputService5TestShapeOutputService5TestCaseOperation1Input struct {
	metadataOutputService5TestShapeOutputService5TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService5TestShapeOutputService5TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService5TestShapeOutputShape struct {
	ListMember []*OutputService5TestShapeSingleStruct `type:"list"`

	metadataOutputService5TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService5TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService5TestShapeSingleStruct struct {
	Foo *string `type:"string"`

	metadataOutputService5TestShapeSingleStruct `json:"-", xml:"-"`
}

type metadataOutputService5TestShapeSingleStruct struct {
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService6ProtocolTest{service}
}

// OutputService6TestCaseOperation1Request generates a request for the OutputService6TestCaseOperation1 operation.
func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1Request(input *OutputService6TestShapeOutputService6TestCaseOperation1Input) (req *aws.Request, output *OutputService6TestShapeOutputShape) {
	if opOutputService6TestCaseOperation1 == nil {
		opOutputService6TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService6TestCaseOperation1, input, output)
	output = &OutputService6TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1(input *OutputService6TestShapeOutputService6TestCaseOperation1Input) (output *OutputService6TestShapeOutputShape, err error) {
	req, out := c.OutputService6TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService6TestCaseOperation1 *aws.Operation

type OutputService6TestShapeOutputService6TestCaseOperation1Input struct {
	metadataOutputService6TestShapeOutputService6TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService6TestShapeOutputService6TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService6TestShapeOutputShape struct {
	MapMember *map[string][]*int64 `type:"map"`

	metadataOutputService6TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService6TestShapeOutputShape struct {
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService7ProtocolTest{service}
}

// OutputService7TestCaseOperation1Request generates a request for the OutputService7TestCaseOperation1 operation.
func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1Request(input *OutputService7TestShapeOutputService7TestCaseOperation1Input) (req *aws.Request, output *OutputService7TestShapeOutputShape) {
	if opOutputService7TestCaseOperation1 == nil {
		opOutputService7TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService7TestCaseOperation1, input, output)
	output = &OutputService7TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1(input *OutputService7TestShapeOutputService7TestCaseOperation1Input) (output *OutputService7TestShapeOutputShape, err error) {
	req, out := c.OutputService7TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService7TestCaseOperation1 *aws.Operation

type OutputService7TestShapeOutputService7TestCaseOperation1Input struct {
	metadataOutputService7TestShapeOutputService7TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService7TestShapeOutputService7TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService7TestShapeOutputShape struct {
	MapMember *map[string]*time.Time `type:"map"`

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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService8ProtocolTest{service}
}

// OutputService8TestCaseOperation1Request generates a request for the OutputService8TestCaseOperation1 operation.
func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1Request(input *OutputService8TestShapeOutputService8TestCaseOperation1Input) (req *aws.Request, output *OutputService8TestShapeOutputShape) {
	if opOutputService8TestCaseOperation1 == nil {
		opOutputService8TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService8TestCaseOperation1, input, output)
	output = &OutputService8TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1(input *OutputService8TestShapeOutputService8TestCaseOperation1Input) (output *OutputService8TestShapeOutputShape, err error) {
	req, out := c.OutputService8TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService8TestCaseOperation1 *aws.Operation

type OutputService8TestShapeOutputService8TestCaseOperation1Input struct {
	metadataOutputService8TestShapeOutputService8TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService8TestShapeOutputService8TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService8TestShapeOutputShape struct {
	StrType *string `type:"string"`

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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService9ProtocolTest{service}
}

// OutputService9TestCaseOperation1Request generates a request for the OutputService9TestCaseOperation1 operation.
func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1Request(input *OutputService9TestShapeOutputService9TestCaseOperation1Input) (req *aws.Request, output *OutputService9TestShapeOutputShape) {
	if opOutputService9TestCaseOperation1 == nil {
		opOutputService9TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService9TestCaseOperation1, input, output)
	output = &OutputService9TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1(input *OutputService9TestShapeOutputService9TestCaseOperation1Input) (output *OutputService9TestShapeOutputShape, err error) {
	req, out := c.OutputService9TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService9TestCaseOperation1 *aws.Operation

type OutputService9TestShapeOutputService9TestCaseOperation1Input struct {
	metadataOutputService9TestShapeOutputService9TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService9TestShapeOutputService9TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService9TestShapeOutputShape struct {
	AllHeaders *map[string]*string `location:"headers" type:"map"`

	PrefixedHeaders *map[string]*string `location:"headers" locationName:"X-" type:"map"`

	metadataOutputService9TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService9TestShapeOutputShape struct {
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService10ProtocolTest{service}
}

// OutputService10TestCaseOperation1Request generates a request for the OutputService10TestCaseOperation1 operation.
func (c *OutputService10ProtocolTest) OutputService10TestCaseOperation1Request(input *OutputService10TestShapeOutputService10TestCaseOperation1Input) (req *aws.Request, output *OutputService10TestShapeOutputShape) {
	if opOutputService10TestCaseOperation1 == nil {
		opOutputService10TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService10TestCaseOperation1, input, output)
	output = &OutputService10TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService10ProtocolTest) OutputService10TestCaseOperation1(input *OutputService10TestShapeOutputService10TestCaseOperation1Input) (output *OutputService10TestShapeOutputShape, err error) {
	req, out := c.OutputService10TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService10TestCaseOperation1 *aws.Operation

type OutputService10TestShapeBodyStructure struct {
	Foo *string `type:"string"`

	metadataOutputService10TestShapeBodyStructure `json:"-", xml:"-"`
}

type metadataOutputService10TestShapeBodyStructure struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService10TestShapeOutputService10TestCaseOperation1Input struct {
	metadataOutputService10TestShapeOutputService10TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService10TestShapeOutputService10TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService10TestShapeOutputShape struct {
	Data *OutputService10TestShapeBodyStructure `type:"structure"`

	Header *string `location:"header" locationName:"X-Foo" type:"string"`

	metadataOutputService10TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService10TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Data"`
}

// OutputService11ProtocolTest is a client for OutputService11ProtocolTest.
type OutputService11ProtocolTest struct {
	*aws.Service
}

type OutputService11ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService11ProtocolTest client.
func NewOutputService11ProtocolTest(config *OutputService11ProtocolTestConfig) *OutputService11ProtocolTest {
	if config == nil {
		config = &OutputService11ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice11protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &OutputService11ProtocolTest{service}
}

// OutputService11TestCaseOperation1Request generates a request for the OutputService11TestCaseOperation1 operation.
func (c *OutputService11ProtocolTest) OutputService11TestCaseOperation1Request(input *OutputService11TestShapeOutputService11TestCaseOperation1Input) (req *aws.Request, output *OutputService11TestShapeOutputShape) {
	if opOutputService11TestCaseOperation1 == nil {
		opOutputService11TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService11TestCaseOperation1, input, output)
	output = &OutputService11TestShapeOutputShape{}
	req.Data = output
	return
}

func (c *OutputService11ProtocolTest) OutputService11TestCaseOperation1(input *OutputService11TestShapeOutputService11TestCaseOperation1Input) (output *OutputService11TestShapeOutputShape, err error) {
	req, out := c.OutputService11TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService11TestCaseOperation1 *aws.Operation

type OutputService11TestShapeOutputService11TestCaseOperation1Input struct {
	metadataOutputService11TestShapeOutputService11TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService11TestShapeOutputService11TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService11TestShapeOutputShape struct {
	Stream []byte `type:"blob"`

	metadataOutputService11TestShapeOutputShape `json:"-", xml:"-"`
}

type metadataOutputService11TestShapeOutputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Stream"`
}

//
// Tests begin here
//

func TestOutputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	svc := NewOutputService1ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"Str\": \"myname\", \"Num\": 123, \"FalseBool\": false, \"TrueBool\": true, \"Float\": 1.2, \"Double\": 1.3, \"Long\": 200, \"Char\": \"a\"}"))
	req, out := svc.OutputService1TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("ImaHeader", "test")
	req.HTTPResponse.Header.Set("X-Foo", "abc")

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "a", *out.Char)
	assert.Equal(t, 1.3, *out.Double)
	assert.Equal(t, false, *out.FalseBool)
	assert.Equal(t, 1.2, *out.Float)
	assert.Equal(t, "test", *out.ImaHeader)
	assert.Equal(t, "abc", *out.ImaHeaderLocation)
	assert.Equal(t, 200, *out.Long)
	assert.Equal(t, 123, *out.Num)
	assert.Equal(t, 200, *out.Status)
	assert.Equal(t, "myname", *out.Str)
	assert.Equal(t, true, *out.TrueBool)

}

func TestOutputService2ProtocolTestBlobMembersCase1(t *testing.T) {
	svc := NewOutputService2ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"BlobMember\": \"aGkh\", \"StructMember\": {\"foo\": \"dGhlcmUh\"}}"))
	req, out := svc.OutputService2TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "hi!", string(out.BlobMember))
	assert.Equal(t, "there!", string(out.StructMember.Foo))

}

func TestOutputService3ProtocolTestTimestampMembersCase1(t *testing.T) {
	svc := NewOutputService3ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"TimeMember\": 1398796238, \"StructMember\": {\"foo\": 1398796238}}"))
	req, out := svc.OutputService3TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, time.Unix(1.398796238e+09, 0).UTC().String(), out.StructMember.Foo.String())
	assert.Equal(t, time.Unix(1.398796238e+09, 0).UTC().String(), out.TimeMember.String())

}

func TestOutputService4ProtocolTestListsCase1(t *testing.T) {
	svc := NewOutputService4ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"ListMember\": [\"a\", \"b\"]}"))
	req, out := svc.OutputService4TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "a", *out.ListMember[0])
	assert.Equal(t, "b", *out.ListMember[1])

}

func TestOutputService5ProtocolTestListsWithStructureMemberCase1(t *testing.T) {
	svc := NewOutputService5ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"ListMember\": [{\"Foo\": \"a\"}, {\"Foo\": \"b\"}]}"))
	req, out := svc.OutputService5TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "a", *out.ListMember[0].Foo)
	assert.Equal(t, "b", *out.ListMember[1].Foo)

}

func TestOutputService6ProtocolTestMapsCase1(t *testing.T) {
	svc := NewOutputService6ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"MapMember\": {\"a\": [1, 2], \"b\": [3, 4]}}"))
	req, out := svc.OutputService6TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, 1, *(*out.MapMember)["a"][0])
	assert.Equal(t, 2, *(*out.MapMember)["a"][1])
	assert.Equal(t, 3, *(*out.MapMember)["b"][0])
	assert.Equal(t, 4, *(*out.MapMember)["b"][1])

}

func TestOutputService7ProtocolTestComplexMapValuesCase1(t *testing.T) {
	svc := NewOutputService7ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"MapMember\": {\"a\": 1398796238, \"b\": 1398796238}}"))
	req, out := svc.OutputService7TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, time.Unix(1.398796238e+09, 0).UTC().String(), (*out.MapMember)["a"].String())
	assert.Equal(t, time.Unix(1.398796238e+09, 0).UTC().String(), (*out.MapMember)["b"].String())

}

func TestOutputService8ProtocolTestIgnoresExtraDataCase1(t *testing.T) {
	svc := NewOutputService8ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"foo\": \"bar\"}"))
	req, out := svc.OutputService8TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used

}

func TestOutputService9ProtocolTestSupportsHeaderMapsCase1(t *testing.T) {
	svc := NewOutputService9ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{}"))
	req, out := svc.OutputService9TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("Content-Length", "10")
	req.HTTPResponse.Header.Set("X-bam", "boo")
	req.HTTPResponse.Header.Set("x-Foo", "bar")

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "10", *(*out.AllHeaders)["Content-Length"])
	assert.Equal(t, "boo", *(*out.AllHeaders)["X-Bam"])
	assert.Equal(t, "bar", *(*out.AllHeaders)["X-Foo"])
	assert.Equal(t, "boo", *(*out.PrefixedHeaders)["Bam"])
	assert.Equal(t, "bar", *(*out.PrefixedHeaders)["Foo"])

}

func TestOutputService10ProtocolTestJSONPayloadCase1(t *testing.T) {
	svc := NewOutputService10ProtocolTest(nil)

	buf := bytes.NewReader([]byte("{\"Foo\": \"abc\"}"))
	req, out := svc.OutputService10TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("X-Foo", "baz")

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", *out.Data.Foo)
	assert.Equal(t, "baz", *out.Header)

}

func TestOutputService11ProtocolTestStreamingPayloadCase1(t *testing.T) {
	svc := NewOutputService11ProtocolTest(nil)

	buf := bytes.NewReader([]byte("abc"))
	req, out := svc.OutputService11TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restjson.UnmarshalMeta(req)
	restjson.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", string(out.Stream))

}

