package query_test

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/query"
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService1ProtocolTest{service}
}

// OutputService1TestCaseOperation1Request generates a request for the OutputService1TestCaseOperation1 operation.
func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1Request(input *OutputService1TestShapeOutputService1TestCaseOperation1Input) (req *aws.Request, output *OutputService1TestShapeOutputService1TestCaseOperation1Output) {
	if opOutputService1TestCaseOperation1 == nil {
		opOutputService1TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService1TestCaseOperation1, input, output)
	output = &OutputService1TestShapeOutputService1TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1(input *OutputService1TestShapeOutputService1TestCaseOperation1Input) (output *OutputService1TestShapeOutputService1TestCaseOperation1Output, err error) {
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

type OutputService1TestShapeOutputService1TestCaseOperation1Output struct {
	Char      *string    `type:"character"`
	Double    *float64   `type:"double"`
	FalseBool *bool      `type:"boolean"`
	Float     *float64   `type:"float"`
	Long      *int64     `type:"long"`
	Num       *int64     `locationName:"FooNum" type:"integer"`
	Str       *string    `type:"string"`
	Timestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`
	TrueBool  *bool      `type:"boolean"`

	metadataOutputService1TestShapeOutputService1TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService1TestShapeOutputService1TestCaseOperation1Output struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService2ProtocolTest{service}
}

// OutputService2TestCaseOperation1Request generates a request for the OutputService2TestCaseOperation1 operation.
func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1Request(input *OutputService2TestShapeOutputService2TestCaseOperation1Input) (req *aws.Request, output *OutputService2TestShapeOutputService2TestCaseOperation1Output) {
	if opOutputService2TestCaseOperation1 == nil {
		opOutputService2TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService2TestCaseOperation1, input, output)
	output = &OutputService2TestShapeOutputService2TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1(input *OutputService2TestShapeOutputService2TestCaseOperation1Input) (output *OutputService2TestShapeOutputService2TestCaseOperation1Output, err error) {
	req, out := c.OutputService2TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService2TestCaseOperation1 *aws.Operation

type OutputService2TestShapeOutputService2TestCaseOperation1Input struct {
	metadataOutputService2TestShapeOutputService2TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService2TestShapeOutputService2TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService2TestShapeOutputService2TestCaseOperation1Output struct {
	Num *int64  `type:"integer"`
	Str *string `type:"string"`

	metadataOutputService2TestShapeOutputService2TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService2TestShapeOutputService2TestCaseOperation1Output struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService3ProtocolTest{service}
}

// OutputService3TestCaseOperation1Request generates a request for the OutputService3TestCaseOperation1 operation.
func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1Request(input *OutputService3TestShapeOutputService3TestCaseOperation1Input) (req *aws.Request, output *OutputService3TestShapeOutputService3TestCaseOperation1Output) {
	if opOutputService3TestCaseOperation1 == nil {
		opOutputService3TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService3TestCaseOperation1, input, output)
	output = &OutputService3TestShapeOutputService3TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1(input *OutputService3TestShapeOutputService3TestCaseOperation1Input) (output *OutputService3TestShapeOutputService3TestCaseOperation1Output, err error) {
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

type OutputService3TestShapeOutputService3TestCaseOperation1Output struct {
	Blob []byte `type:"blob"`

	metadataOutputService3TestShapeOutputService3TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService3TestShapeOutputService3TestCaseOperation1Output struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService4ProtocolTest{service}
}

// OutputService4TestCaseOperation1Request generates a request for the OutputService4TestCaseOperation1 operation.
func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1Request(input *OutputService4TestShapeOutputService4TestCaseOperation1Input) (req *aws.Request, output *OutputService4TestShapeOutputService4TestCaseOperation1Output) {
	if opOutputService4TestCaseOperation1 == nil {
		opOutputService4TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService4TestCaseOperation1, input, output)
	output = &OutputService4TestShapeOutputService4TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1(input *OutputService4TestShapeOutputService4TestCaseOperation1Input) (output *OutputService4TestShapeOutputService4TestCaseOperation1Output, err error) {
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

type OutputService4TestShapeOutputService4TestCaseOperation1Output struct {
	ListMember []*string `type:"list"`

	metadataOutputService4TestShapeOutputService4TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService4TestShapeOutputService4TestCaseOperation1Output struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService5ProtocolTest{service}
}

// OutputService5TestCaseOperation1Request generates a request for the OutputService5TestCaseOperation1 operation.
func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1Request(input *OutputService5TestShapeOutputService5TestCaseOperation1Input) (req *aws.Request, output *OutputService5TestShapeOutputService5TestCaseOperation1Output) {
	if opOutputService5TestCaseOperation1 == nil {
		opOutputService5TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService5TestCaseOperation1, input, output)
	output = &OutputService5TestShapeOutputService5TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1(input *OutputService5TestShapeOutputService5TestCaseOperation1Input) (output *OutputService5TestShapeOutputService5TestCaseOperation1Output, err error) {
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

type OutputService5TestShapeOutputService5TestCaseOperation1Output struct {
	ListMember []*string `locationNameList:"item" type:"list"`

	metadataOutputService5TestShapeOutputService5TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService5TestShapeOutputService5TestCaseOperation1Output struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService6ProtocolTest{service}
}

// OutputService6TestCaseOperation1Request generates a request for the OutputService6TestCaseOperation1 operation.
func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1Request(input *OutputService6TestShapeOutputService6TestCaseOperation1Input) (req *aws.Request, output *OutputService6TestShapeOutputService6TestCaseOperation1Output) {
	if opOutputService6TestCaseOperation1 == nil {
		opOutputService6TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService6TestCaseOperation1, input, output)
	output = &OutputService6TestShapeOutputService6TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1(input *OutputService6TestShapeOutputService6TestCaseOperation1Input) (output *OutputService6TestShapeOutputService6TestCaseOperation1Output, err error) {
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

type OutputService6TestShapeOutputService6TestCaseOperation1Output struct {
	ListMember []*string `type:"list" flattened:"true"`

	metadataOutputService6TestShapeOutputService6TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService6TestShapeOutputService6TestCaseOperation1Output struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService7ProtocolTest{service}
}

// OutputService7TestCaseOperation1Request generates a request for the OutputService7TestCaseOperation1 operation.
func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1Request(input *OutputService7TestShapeOutputService7TestCaseOperation1Input) (req *aws.Request, output *OutputService7TestShapeOutputService7TestCaseOperation1Output) {
	if opOutputService7TestCaseOperation1 == nil {
		opOutputService7TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService7TestCaseOperation1, input, output)
	output = &OutputService7TestShapeOutputService7TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1(input *OutputService7TestShapeOutputService7TestCaseOperation1Input) (output *OutputService7TestShapeOutputService7TestCaseOperation1Output, err error) {
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

type OutputService7TestShapeOutputService7TestCaseOperation1Output struct {
	ListMember []*string `type:"list" flattened:"true"`

	metadataOutputService7TestShapeOutputService7TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService7TestShapeOutputService7TestCaseOperation1Output struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService8ProtocolTest{service}
}

// OutputService8TestCaseOperation1Request generates a request for the OutputService8TestCaseOperation1 operation.
func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1Request(input *OutputService8TestShapeOutputService8TestCaseOperation1Input) (req *aws.Request, output *OutputService8TestShapeOutputService8TestCaseOperation1Output) {
	if opOutputService8TestCaseOperation1 == nil {
		opOutputService8TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService8TestCaseOperation1, input, output)
	output = &OutputService8TestShapeOutputService8TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1(input *OutputService8TestShapeOutputService8TestCaseOperation1Input) (output *OutputService8TestShapeOutputService8TestCaseOperation1Output, err error) {
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

type OutputService8TestShapeOutputService8TestCaseOperation1Output struct {
	List []*OutputService8TestShapeStructureShape `type:"list"`

	metadataOutputService8TestShapeOutputService8TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService8TestShapeOutputService8TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService8TestShapeStructureShape struct {
	Bar *string `type:"string"`
	Baz *string `type:"string"`
	Foo *string `type:"string"`

	metadataOutputService8TestShapeStructureShape `json:"-", xml:"-"`
}

type metadataOutputService8TestShapeStructureShape struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService9ProtocolTest{service}
}

// OutputService9TestCaseOperation1Request generates a request for the OutputService9TestCaseOperation1 operation.
func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1Request(input *OutputService9TestShapeOutputService9TestCaseOperation1Input) (req *aws.Request, output *OutputService9TestShapeOutputService9TestCaseOperation1Output) {
	if opOutputService9TestCaseOperation1 == nil {
		opOutputService9TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService9TestCaseOperation1, input, output)
	output = &OutputService9TestShapeOutputService9TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1(input *OutputService9TestShapeOutputService9TestCaseOperation1Input) (output *OutputService9TestShapeOutputService9TestCaseOperation1Output, err error) {
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

type OutputService9TestShapeOutputService9TestCaseOperation1Output struct {
	List []*OutputService9TestShapeStructureShape `type:"list" flattened:"true"`

	metadataOutputService9TestShapeOutputService9TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService9TestShapeOutputService9TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService9TestShapeStructureShape struct {
	Bar *string `type:"string"`
	Baz *string `type:"string"`
	Foo *string `type:"string"`

	metadataOutputService9TestShapeStructureShape `json:"-", xml:"-"`
}

type metadataOutputService9TestShapeStructureShape struct {
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService10ProtocolTest{service}
}

// OutputService10TestCaseOperation1Request generates a request for the OutputService10TestCaseOperation1 operation.
func (c *OutputService10ProtocolTest) OutputService10TestCaseOperation1Request(input *OutputService10TestShapeOutputService10TestCaseOperation1Input) (req *aws.Request, output *OutputService10TestShapeOutputService10TestCaseOperation1Output) {
	if opOutputService10TestCaseOperation1 == nil {
		opOutputService10TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService10TestCaseOperation1, input, output)
	output = &OutputService10TestShapeOutputService10TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService10ProtocolTest) OutputService10TestCaseOperation1(input *OutputService10TestShapeOutputService10TestCaseOperation1Input) (output *OutputService10TestShapeOutputService10TestCaseOperation1Output, err error) {
	req, out := c.OutputService10TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService10TestCaseOperation1 *aws.Operation

type OutputService10TestShapeOutputService10TestCaseOperation1Input struct {
	metadataOutputService10TestShapeOutputService10TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService10TestShapeOutputService10TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService10TestShapeOutputService10TestCaseOperation1Output struct {
	List []*string `locationNameList:"NamedList" type:"list" flattened:"true"`

	metadataOutputService10TestShapeOutputService10TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService10TestShapeOutputService10TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
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
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService11ProtocolTest{service}
}

// OutputService11TestCaseOperation1Request generates a request for the OutputService11TestCaseOperation1 operation.
func (c *OutputService11ProtocolTest) OutputService11TestCaseOperation1Request(input *OutputService11TestShapeOutputService11TestCaseOperation1Input) (req *aws.Request, output *OutputService11TestShapeOutputService11TestCaseOperation1Output) {
	if opOutputService11TestCaseOperation1 == nil {
		opOutputService11TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService11TestCaseOperation1, input, output)
	output = &OutputService11TestShapeOutputService11TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService11ProtocolTest) OutputService11TestCaseOperation1(input *OutputService11TestShapeOutputService11TestCaseOperation1Input) (output *OutputService11TestShapeOutputService11TestCaseOperation1Output, err error) {
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

type OutputService11TestShapeOutputService11TestCaseOperation1Output struct {
	Map *map[string]*OutputService11TestShapeStructType `type:"map"`

	metadataOutputService11TestShapeOutputService11TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService11TestShapeOutputService11TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService11TestShapeStructType struct {
	Foo *string `locationName:"foo" type:"string"`

	metadataOutputService11TestShapeStructType `json:"-", xml:"-"`
}

type metadataOutputService11TestShapeStructType struct {
	SDKShapeTraits bool `type:"structure"`
}

// OutputService12ProtocolTest is a client for OutputService12ProtocolTest.
type OutputService12ProtocolTest struct {
	*aws.Service
}

type OutputService12ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService12ProtocolTest client.
func NewOutputService12ProtocolTest(config *OutputService12ProtocolTestConfig) *OutputService12ProtocolTest {
	if config == nil {
		config = &OutputService12ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice12protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService12ProtocolTest{service}
}

// OutputService12TestCaseOperation1Request generates a request for the OutputService12TestCaseOperation1 operation.
func (c *OutputService12ProtocolTest) OutputService12TestCaseOperation1Request(input *OutputService12TestShapeOutputService12TestCaseOperation1Input) (req *aws.Request, output *OutputService12TestShapeOutputService12TestCaseOperation1Output) {
	if opOutputService12TestCaseOperation1 == nil {
		opOutputService12TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService12TestCaseOperation1, input, output)
	output = &OutputService12TestShapeOutputService12TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService12ProtocolTest) OutputService12TestCaseOperation1(input *OutputService12TestShapeOutputService12TestCaseOperation1Input) (output *OutputService12TestShapeOutputService12TestCaseOperation1Output, err error) {
	req, out := c.OutputService12TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService12TestCaseOperation1 *aws.Operation

type OutputService12TestShapeOutputService12TestCaseOperation1Input struct {
	metadataOutputService12TestShapeOutputService12TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService12TestShapeOutputService12TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService12TestShapeOutputService12TestCaseOperation1Output struct {
	Map *map[string]*string `type:"map" flattened:"true"`

	metadataOutputService12TestShapeOutputService12TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService12TestShapeOutputService12TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

// OutputService13ProtocolTest is a client for OutputService13ProtocolTest.
type OutputService13ProtocolTest struct {
	*aws.Service
}

type OutputService13ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService13ProtocolTest client.
func NewOutputService13ProtocolTest(config *OutputService13ProtocolTestConfig) *OutputService13ProtocolTest {
	if config == nil {
		config = &OutputService13ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice13protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService13ProtocolTest{service}
}

// OutputService13TestCaseOperation1Request generates a request for the OutputService13TestCaseOperation1 operation.
func (c *OutputService13ProtocolTest) OutputService13TestCaseOperation1Request(input *OutputService13TestShapeOutputService13TestCaseOperation1Input) (req *aws.Request, output *OutputService13TestShapeOutputService13TestCaseOperation1Output) {
	if opOutputService13TestCaseOperation1 == nil {
		opOutputService13TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService13TestCaseOperation1, input, output)
	output = &OutputService13TestShapeOutputService13TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService13ProtocolTest) OutputService13TestCaseOperation1(input *OutputService13TestShapeOutputService13TestCaseOperation1Input) (output *OutputService13TestShapeOutputService13TestCaseOperation1Output, err error) {
	req, out := c.OutputService13TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService13TestCaseOperation1 *aws.Operation

type OutputService13TestShapeOutputService13TestCaseOperation1Input struct {
	metadataOutputService13TestShapeOutputService13TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService13TestShapeOutputService13TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService13TestShapeOutputService13TestCaseOperation1Output struct {
	Map *map[string]*string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	metadataOutputService13TestShapeOutputService13TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService13TestShapeOutputService13TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

// OutputService14ProtocolTest is a client for OutputService14ProtocolTest.
type OutputService14ProtocolTest struct {
	*aws.Service
}

type OutputService14ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new OutputService14ProtocolTest client.
func NewOutputService14ProtocolTest(config *OutputService14ProtocolTestConfig) *OutputService14ProtocolTest {
	if config == nil {
		config = &OutputService14ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "outputservice14protocoltest",
		APIVersion:  "",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &OutputService14ProtocolTest{service}
}

// OutputService14TestCaseOperation1Request generates a request for the OutputService14TestCaseOperation1 operation.
func (c *OutputService14ProtocolTest) OutputService14TestCaseOperation1Request(input *OutputService14TestShapeOutputService14TestCaseOperation1Input) (req *aws.Request, output *OutputService14TestShapeOutputService14TestCaseOperation1Output) {
	if opOutputService14TestCaseOperation1 == nil {
		opOutputService14TestCaseOperation1 = &aws.Operation{
			Name: "OperationName",
		}
	}

	req = aws.NewRequest(c.Service, opOutputService14TestCaseOperation1, input, output)
	output = &OutputService14TestShapeOutputService14TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService14ProtocolTest) OutputService14TestCaseOperation1(input *OutputService14TestShapeOutputService14TestCaseOperation1Input) (output *OutputService14TestShapeOutputService14TestCaseOperation1Output, err error) {
	req, out := c.OutputService14TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opOutputService14TestCaseOperation1 *aws.Operation

type OutputService14TestShapeOutputService14TestCaseOperation1Input struct {
	metadataOutputService14TestShapeOutputService14TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataOutputService14TestShapeOutputService14TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type OutputService14TestShapeOutputService14TestCaseOperation1Output struct {
	Map *map[string]*string `locationNameKey:"foo" locationNameValue:"bar" type:"map" flattened:"true"`

	metadataOutputService14TestShapeOutputService14TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataOutputService14TestShapeOutputService14TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//
// Tests begin here
//

func TestOutputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	svc := NewOutputService1ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Str>myname</Str><FooNum>123</FooNum><FalseBool>false</FalseBool><TrueBool>true</TrueBool><Float>1.2</Float><Double>1.3</Double><Long>200</Long><Char>a</Char><Timestamp>2015-01-25T08:00:00Z</Timestamp></OperationNameResult><ResponseMetadata><RequestId>request-id</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService1TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "a", *out.Char)
	assert.Equal(t, 1.3, *out.Double)
	assert.Equal(t, false, *out.FalseBool)
	assert.Equal(t, 1.2, *out.Float)
	assert.Equal(t, 200, *out.Long)
	assert.Equal(t, 123, *out.Num)
	assert.Equal(t, "myname", *out.Str)
	assert.Equal(t, time.Unix(1.4221728e+09, 0).UTC().String(), out.Timestamp.String())
	assert.Equal(t, true, *out.TrueBool)

}

func TestOutputService2ProtocolTestNotAllMembersInResponseCase1(t *testing.T) {
	svc := NewOutputService2ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Str>myname</Str></OperationNameResult><ResponseMetadata><RequestId>request-id</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService2TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "myname", *out.Str)

}

func TestOutputService3ProtocolTestBlobCase1(t *testing.T) {
	svc := NewOutputService3ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Blob>dmFsdWU=</Blob></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService3TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "value", string(out.Blob))

}

func TestOutputService4ProtocolTestListsCase1(t *testing.T) {
	svc := NewOutputService4ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><ListMember><member>abc</member><member>123</member></ListMember></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService4TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", *out.ListMember[0])
	assert.Equal(t, "123", *out.ListMember[1])

}

func TestOutputService5ProtocolTestListWithCustomMemberNameCase1(t *testing.T) {
	svc := NewOutputService5ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><ListMember><item>abc</item><item>123</item></ListMember></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService5TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", *out.ListMember[0])
	assert.Equal(t, "123", *out.ListMember[1])

}

func TestOutputService6ProtocolTestFlattenedListCase1(t *testing.T) {
	svc := NewOutputService6ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><ListMember>abc</ListMember><ListMember>123</ListMember></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService6TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", *out.ListMember[0])
	assert.Equal(t, "123", *out.ListMember[1])

}

func TestOutputService7ProtocolTestFlattenedSingleElementListCase1(t *testing.T) {
	svc := NewOutputService7ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><ListMember>abc</ListMember></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService7TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", *out.ListMember[0])

}

func TestOutputService8ProtocolTestListOfStructuresCase1(t *testing.T) {
	svc := NewOutputService8ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse xmlns=\"https://service.amazonaws.com/doc/2010-05-08/\"><OperationNameResult><List><member><Foo>firstfoo</Foo><Bar>firstbar</Bar><Baz>firstbaz</Baz></member><member><Foo>secondfoo</Foo><Bar>secondbar</Bar><Baz>secondbaz</Baz></member></List></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService8TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "firstbar", *out.List[0].Bar)
	assert.Equal(t, "firstbaz", *out.List[0].Baz)
	assert.Equal(t, "firstfoo", *out.List[0].Foo)
	assert.Equal(t, "secondbar", *out.List[1].Bar)
	assert.Equal(t, "secondbaz", *out.List[1].Baz)
	assert.Equal(t, "secondfoo", *out.List[1].Foo)

}

func TestOutputService9ProtocolTestFlattenedListOfStructuresCase1(t *testing.T) {
	svc := NewOutputService9ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse xmlns=\"https://service.amazonaws.com/doc/2010-05-08/\"><OperationNameResult><List><Foo>firstfoo</Foo><Bar>firstbar</Bar><Baz>firstbaz</Baz></List><List><Foo>secondfoo</Foo><Bar>secondbar</Bar><Baz>secondbaz</Baz></List></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService9TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "firstbar", *out.List[0].Bar)
	assert.Equal(t, "firstbaz", *out.List[0].Baz)
	assert.Equal(t, "firstfoo", *out.List[0].Foo)
	assert.Equal(t, "secondbar", *out.List[1].Bar)
	assert.Equal(t, "secondbaz", *out.List[1].Baz)
	assert.Equal(t, "secondfoo", *out.List[1].Foo)

}

func TestOutputService10ProtocolTestFlattenedListWithLocationNameCase1(t *testing.T) {
	svc := NewOutputService10ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse xmlns=\"https://service.amazonaws.com/doc/2010-05-08/\"><OperationNameResult><NamedList>a</NamedList><NamedList>b</NamedList></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService10TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "a", *out.List[0])
	assert.Equal(t, "b", *out.List[1])

}

func TestOutputService11ProtocolTestNormalMapCase1(t *testing.T) {
	svc := NewOutputService11ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse xmlns=\"https://service.amazonaws.com/doc/2010-05-08\"><OperationNameResult><Map><entry><key>qux</key><value><foo>bar</foo></value></entry><entry><key>baz</key><value><foo>bam</foo></value></entry></Map></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService11TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "bam", *(*out.Map)["baz"].Foo)
	assert.Equal(t, "bar", *(*out.Map)["qux"].Foo)

}

func TestOutputService12ProtocolTestFlattenedMapCase1(t *testing.T) {
	svc := NewOutputService12ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Map><key>qux</key><value>bar</value></Map><Map><key>baz</key><value>bam</value></Map></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService12TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "bam", *(*out.Map)["baz"])
	assert.Equal(t, "bar", *(*out.Map)["qux"])

}

func TestOutputService13ProtocolTestFlattenedMapInShapeDefinitionCase1(t *testing.T) {
	svc := NewOutputService13ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Attribute><Name>qux</Name><Value>bar</Value></Attribute></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService13TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "bar", *(*out.Map)["qux"])

}

func TestOutputService14ProtocolTestNamedMapCase1(t *testing.T) {
	svc := NewOutputService14ProtocolTest(nil)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Map><foo>qux</foo><bar>bar</bar></Map><Map><foo>baz</foo><bar>bam</bar></Map></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req, out := svc.OutputService14TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req)
	query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "bam", *(*out.Map)["baz"])
	assert.Equal(t, "bar", *(*out.Map)["qux"])

}

