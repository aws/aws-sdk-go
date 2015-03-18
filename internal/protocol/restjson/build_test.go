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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService1ProtocolTest{service}
}

// InputService1TestCaseOperation1Request generates a request for the InputService1TestCaseOperation1 operation.
func (c *InputService1ProtocolTest) InputService1TestCaseOperation1Request(input *InputService1TestShapeInputService1TestCaseOperation1Input) (req *aws.Request, output *InputService1TestShapeInputService1TestCaseOperation1Output) {
	if opInputService1TestCaseOperation1 == nil {
		opInputService1TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
		}
	}

	req = aws.NewRequest(c.Service, opInputService1TestCaseOperation1, input, output)
	output = &InputService1TestShapeInputService1TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService1ProtocolTest) InputService1TestCaseOperation1(input *InputService1TestShapeInputService1TestCaseOperation1Input) (output *InputService1TestShapeInputService1TestCaseOperation1Output, err error) {
	req, out := c.InputService1TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService1TestCaseOperation1 *aws.Operation

type InputService1TestShapeInputService1TestCaseOperation1Input struct {
	PipelineId *string `location:"uri" type:"string" json:"-" xml:"-"`

	metadataInputService1TestShapeInputService1TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataInputService1TestShapeInputService1TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService1TestShapeInputService1TestCaseOperation1Output struct {
	metadataInputService1TestShapeInputService1TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService1TestShapeInputService1TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService2ProtocolTest{service}
}

// InputService2TestCaseOperation1Request generates a request for the InputService2TestCaseOperation1 operation.
func (c *InputService2ProtocolTest) InputService2TestCaseOperation1Request(input *InputService2TestShapeInputService2TestCaseOperation1Input) (req *aws.Request, output *InputService2TestShapeInputService2TestCaseOperation1Output) {
	if opInputService2TestCaseOperation1 == nil {
		opInputService2TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
		}
	}

	req = aws.NewRequest(c.Service, opInputService2TestCaseOperation1, input, output)
	output = &InputService2TestShapeInputService2TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService2ProtocolTest) InputService2TestCaseOperation1(input *InputService2TestShapeInputService2TestCaseOperation1Input) (output *InputService2TestShapeInputService2TestCaseOperation1Output, err error) {
	req, out := c.InputService2TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService2TestCaseOperation1 *aws.Operation

type InputService2TestShapeInputService2TestCaseOperation1Input struct {
	Foo *string `location:"uri" locationName:"PipelineId" type:"string" json:"-" xml:"-"`

	metadataInputService2TestShapeInputService2TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataInputService2TestShapeInputService2TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService2TestShapeInputService2TestCaseOperation1Output struct {
	metadataInputService2TestShapeInputService2TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService2TestShapeInputService2TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService3ProtocolTest{service}
}

// InputService3TestCaseOperation1Request generates a request for the InputService3TestCaseOperation1 operation.
func (c *InputService3ProtocolTest) InputService3TestCaseOperation1Request(input *InputService3TestShapeInputService3TestCaseOperation1Input) (req *aws.Request, output *InputService3TestShapeInputService3TestCaseOperation1Output) {
	if opInputService3TestCaseOperation1 == nil {
		opInputService3TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
		}
	}

	req = aws.NewRequest(c.Service, opInputService3TestCaseOperation1, input, output)
	output = &InputService3TestShapeInputService3TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService3ProtocolTest) InputService3TestCaseOperation1(input *InputService3TestShapeInputService3TestCaseOperation1Input) (output *InputService3TestShapeInputService3TestCaseOperation1Output, err error) {
	req, out := c.InputService3TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService3TestCaseOperation1 *aws.Operation

type InputService3TestShapeInputService3TestCaseOperation1Input struct {
	Ascending  *string `location:"querystring" locationName:"Ascending" type:"string" json:"-" xml:"-"`
	PageToken  *string `location:"querystring" locationName:"PageToken" type:"string" json:"-" xml:"-"`
	PipelineId *string `location:"uri" locationName:"PipelineId" type:"string" json:"-" xml:"-"`

	metadataInputService3TestShapeInputService3TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataInputService3TestShapeInputService3TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService3TestShapeInputService3TestCaseOperation1Output struct {
	metadataInputService3TestShapeInputService3TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService3TestShapeInputService3TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService4ProtocolTest{service}
}

// InputService4TestCaseOperation1Request generates a request for the InputService4TestCaseOperation1 operation.
func (c *InputService4ProtocolTest) InputService4TestCaseOperation1Request(input *InputService4TestShapeInputService4TestCaseOperation1Input) (req *aws.Request, output *InputService4TestShapeInputService4TestCaseOperation1Output) {
	if opInputService4TestCaseOperation1 == nil {
		opInputService4TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
		}
	}

	req = aws.NewRequest(c.Service, opInputService4TestCaseOperation1, input, output)
	output = &InputService4TestShapeInputService4TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService4ProtocolTest) InputService4TestCaseOperation1(input *InputService4TestShapeInputService4TestCaseOperation1Input) (output *InputService4TestShapeInputService4TestCaseOperation1Output, err error) {
	req, out := c.InputService4TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService4TestCaseOperation1 *aws.Operation

type InputService4TestShapeInputService4TestCaseOperation1Input struct {
	Ascending  *string                           `location:"querystring" locationName:"Ascending" type:"string" json:"-" xml:"-"`
	Config     *InputService4TestShapeStructType `type:"structure" json:",omitempty"`
	PageToken  *string                           `location:"querystring" locationName:"PageToken" type:"string" json:"-" xml:"-"`
	PipelineId *string                           `location:"uri" locationName:"PipelineId" type:"string" json:"-" xml:"-"`

	metadataInputService4TestShapeInputService4TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataInputService4TestShapeInputService4TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService4TestShapeInputService4TestCaseOperation1Output struct {
	metadataInputService4TestShapeInputService4TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService4TestShapeInputService4TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService4TestShapeStructType struct {
	A *string `type:"string" json:",omitempty"`
	B *string `type:"string" json:",omitempty"`

	metadataInputService4TestShapeStructType `json:"-", xml:"-"`
}

type metadataInputService4TestShapeStructType struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService5ProtocolTest{service}
}

// InputService5TestCaseOperation1Request generates a request for the InputService5TestCaseOperation1 operation.
func (c *InputService5ProtocolTest) InputService5TestCaseOperation1Request(input *InputService5TestShapeInputService5TestCaseOperation1Input) (req *aws.Request, output *InputService5TestShapeInputService5TestCaseOperation1Output) {
	if opInputService5TestCaseOperation1 == nil {
		opInputService5TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
		}
	}

	req = aws.NewRequest(c.Service, opInputService5TestCaseOperation1, input, output)
	output = &InputService5TestShapeInputService5TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService5ProtocolTest) InputService5TestCaseOperation1(input *InputService5TestShapeInputService5TestCaseOperation1Input) (output *InputService5TestShapeInputService5TestCaseOperation1Output, err error) {
	req, out := c.InputService5TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService5TestCaseOperation1 *aws.Operation

type InputService5TestShapeInputService5TestCaseOperation1Input struct {
	Ascending  *string                           `location:"querystring" locationName:"Ascending" type:"string" json:"-" xml:"-"`
	Checksum   *string                           `location:"header" locationName:"x-amz-checksum" type:"string" json:"-" xml:"-"`
	Config     *InputService5TestShapeStructType `type:"structure" json:",omitempty"`
	PageToken  *string                           `location:"querystring" locationName:"PageToken" type:"string" json:"-" xml:"-"`
	PipelineId *string                           `location:"uri" locationName:"PipelineId" type:"string" json:"-" xml:"-"`

	metadataInputService5TestShapeInputService5TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataInputService5TestShapeInputService5TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService5TestShapeInputService5TestCaseOperation1Output struct {
	metadataInputService5TestShapeInputService5TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService5TestShapeInputService5TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService5TestShapeStructType struct {
	A *string `type:"string" json:",omitempty"`
	B *string `type:"string" json:",omitempty"`

	metadataInputService5TestShapeStructType `json:"-", xml:"-"`
}

type metadataInputService5TestShapeStructType struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService6ProtocolTest{service}
}

// InputService6TestCaseOperation1Request generates a request for the InputService6TestCaseOperation1 operation.
func (c *InputService6ProtocolTest) InputService6TestCaseOperation1Request(input *InputService6TestShapeInputService6TestCaseOperation1Input) (req *aws.Request, output *InputService6TestShapeInputService6TestCaseOperation1Output) {
	if opInputService6TestCaseOperation1 == nil {
		opInputService6TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/vaults/{vaultName}/archives",
		}
	}

	req = aws.NewRequest(c.Service, opInputService6TestCaseOperation1, input, output)
	output = &InputService6TestShapeInputService6TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService6ProtocolTest) InputService6TestCaseOperation1(input *InputService6TestShapeInputService6TestCaseOperation1Input) (output *InputService6TestShapeInputService6TestCaseOperation1Output, err error) {
	req, out := c.InputService6TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService6TestCaseOperation1 *aws.Operation

type InputService6TestShapeInputService6TestCaseOperation1Input struct {
	Body      []byte  `locationName:"body" type:"blob" json:"body,omitempty"`
	Checksum  *string `location:"header" locationName:"x-amz-sha256-tree-hash" type:"string" json:"-" xml:"-"`
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"json:"-" xml:"-"`

	metadataInputService6TestShapeInputService6TestCaseOperation1Input `json:"-", xml:"-"`
}

type metadataInputService6TestShapeInputService6TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure" payload:"Body" json:",omitempty"`
}

type InputService6TestShapeInputService6TestCaseOperation1Output struct {
	metadataInputService6TestShapeInputService6TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService6TestShapeInputService6TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService7ProtocolTest{service}
}

// InputService7TestCaseOperation1Request generates a request for the InputService7TestCaseOperation1 operation.
func (c *InputService7ProtocolTest) InputService7TestCaseOperation1Request(input *InputService7TestShapeInputShape) (req *aws.Request, output *InputService7TestShapeInputService7TestCaseOperation1Output) {
	if opInputService7TestCaseOperation1 == nil {
		opInputService7TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService7TestCaseOperation1, input, output)
	output = &InputService7TestShapeInputService7TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService7ProtocolTest) InputService7TestCaseOperation1(input *InputService7TestShapeInputShape) (output *InputService7TestShapeInputService7TestCaseOperation1Output, err error) {
	req, out := c.InputService7TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService7TestCaseOperation1 *aws.Operation

// InputService7TestCaseOperation2Request generates a request for the InputService7TestCaseOperation2 operation.
func (c *InputService7ProtocolTest) InputService7TestCaseOperation2Request(input *InputService7TestShapeInputShape) (req *aws.Request, output *InputService7TestShapeInputService7TestCaseOperation2Output) {
	if opInputService7TestCaseOperation2 == nil {
		opInputService7TestCaseOperation2 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path?abc=mno",
		}
	}

	req = aws.NewRequest(c.Service, opInputService7TestCaseOperation2, input, output)
	output = &InputService7TestShapeInputService7TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService7ProtocolTest) InputService7TestCaseOperation2(input *InputService7TestShapeInputShape) (output *InputService7TestShapeInputService7TestCaseOperation2Output, err error) {
	req, out := c.InputService7TestCaseOperation2Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService7TestCaseOperation2 *aws.Operation

type InputService7TestShapeInputService7TestCaseOperation1Output struct {
	metadataInputService7TestShapeInputService7TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService7TestShapeInputService7TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService7TestShapeInputService7TestCaseOperation2Output struct {
	metadataInputService7TestShapeInputService7TestCaseOperation2Output `json:"-", xml:"-"`
}

type metadataInputService7TestShapeInputService7TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService7TestShapeInputShape struct {
	Foo *string `location:"querystring" locationName:"param-name" type:"string" json:"-" xml:"-"`

	metadataInputService7TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService7TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

// InputService8ProtocolTest is a client for InputService8ProtocolTest.
type InputService8ProtocolTest struct {
	*aws.Service
}

type InputService8ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService8ProtocolTest client.
func NewInputService8ProtocolTest(config *InputService8ProtocolTestConfig) *InputService8ProtocolTest {
	if config == nil {
		config = &InputService8ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice8protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService8ProtocolTest{service}
}

// InputService8TestCaseOperation1Request generates a request for the InputService8TestCaseOperation1 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation1Request(input *InputService8TestShapeInputShape) (req *aws.Request, output *InputService8TestShapeInputService8TestCaseOperation1Output) {
	if opInputService8TestCaseOperation1 == nil {
		opInputService8TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService8TestCaseOperation1, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation1(input *InputService8TestShapeInputShape) (output *InputService8TestShapeInputService8TestCaseOperation1Output, err error) {
	req, out := c.InputService8TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService8TestCaseOperation1 *aws.Operation

// InputService8TestCaseOperation2Request generates a request for the InputService8TestCaseOperation2 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation2Request(input *InputService8TestShapeInputShape) (req *aws.Request, output *InputService8TestShapeInputService8TestCaseOperation2Output) {
	if opInputService8TestCaseOperation2 == nil {
		opInputService8TestCaseOperation2 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService8TestCaseOperation2, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation2(input *InputService8TestShapeInputShape) (output *InputService8TestShapeInputService8TestCaseOperation2Output, err error) {
	req, out := c.InputService8TestCaseOperation2Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService8TestCaseOperation2 *aws.Operation

// InputService8TestCaseOperation3Request generates a request for the InputService8TestCaseOperation3 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation3Request(input *InputService8TestShapeInputShape) (req *aws.Request, output *InputService8TestShapeInputService8TestCaseOperation3Output) {
	if opInputService8TestCaseOperation3 == nil {
		opInputService8TestCaseOperation3 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService8TestCaseOperation3, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation3Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation3(input *InputService8TestShapeInputShape) (output *InputService8TestShapeInputService8TestCaseOperation3Output, err error) {
	req, out := c.InputService8TestCaseOperation3Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService8TestCaseOperation3 *aws.Operation

// InputService8TestCaseOperation4Request generates a request for the InputService8TestCaseOperation4 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation4Request(input *InputService8TestShapeInputShape) (req *aws.Request, output *InputService8TestShapeInputService8TestCaseOperation4Output) {
	if opInputService8TestCaseOperation4 == nil {
		opInputService8TestCaseOperation4 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService8TestCaseOperation4, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation4Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation4(input *InputService8TestShapeInputShape) (output *InputService8TestShapeInputService8TestCaseOperation4Output, err error) {
	req, out := c.InputService8TestCaseOperation4Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService8TestCaseOperation4 *aws.Operation

// InputService8TestCaseOperation5Request generates a request for the InputService8TestCaseOperation5 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation5Request(input *InputService8TestShapeInputShape) (req *aws.Request, output *InputService8TestShapeInputService8TestCaseOperation5Output) {
	if opInputService8TestCaseOperation5 == nil {
		opInputService8TestCaseOperation5 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService8TestCaseOperation5, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation5Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation5(input *InputService8TestShapeInputShape) (output *InputService8TestShapeInputService8TestCaseOperation5Output, err error) {
	req, out := c.InputService8TestCaseOperation5Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService8TestCaseOperation5 *aws.Operation

// InputService8TestCaseOperation6Request generates a request for the InputService8TestCaseOperation6 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation6Request(input *InputService8TestShapeInputShape) (req *aws.Request, output *InputService8TestShapeInputService8TestCaseOperation6Output) {
	if opInputService8TestCaseOperation6 == nil {
		opInputService8TestCaseOperation6 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService8TestCaseOperation6, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation6Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation6(input *InputService8TestShapeInputShape) (output *InputService8TestShapeInputService8TestCaseOperation6Output, err error) {
	req, out := c.InputService8TestCaseOperation6Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService8TestCaseOperation6 *aws.Operation

type InputService8TestShapeInputService8TestCaseOperation1Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService8TestShapeInputService8TestCaseOperation2Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation2Output `json:"-", xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService8TestShapeInputService8TestCaseOperation3Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation3Output `json:"-", xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation3Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService8TestShapeInputService8TestCaseOperation4Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation4Output `json:"-", xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation4Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService8TestShapeInputService8TestCaseOperation5Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation5Output `json:"-", xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation5Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService8TestShapeInputService8TestCaseOperation6Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation6Output `json:"-", xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation6Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService8TestShapeInputShape struct {
	RecursiveStruct *InputService8TestShapeRecursiveStructType `type:"structure" json:",omitempty"`

	metadataInputService8TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService8TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService8TestShapeRecursiveStructType struct {
	NoRecurse       *string                                                `type:"string" json:",omitempty"`
	RecursiveList   []*InputService8TestShapeRecursiveStructType           `type:"list" json:",omitempty"`
	RecursiveMap    *map[string]*InputService8TestShapeRecursiveStructType `type:"map" json:",omitempty"`
	RecursiveStruct *InputService8TestShapeRecursiveStructType             `type:"structure" json:",omitempty"`

	metadataInputService8TestShapeRecursiveStructType `json:"-", xml:"-"`
}

type metadataInputService8TestShapeRecursiveStructType struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

// InputService9ProtocolTest is a client for InputService9ProtocolTest.
type InputService9ProtocolTest struct {
	*aws.Service
}

type InputService9ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService9ProtocolTest client.
func NewInputService9ProtocolTest(config *InputService9ProtocolTestConfig) *InputService9ProtocolTest {
	if config == nil {
		config = &InputService9ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice9protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService9ProtocolTest{service}
}

// InputService9TestCaseOperation1Request generates a request for the InputService9TestCaseOperation1 operation.
func (c *InputService9ProtocolTest) InputService9TestCaseOperation1Request(input *InputService9TestShapeInputShape) (req *aws.Request, output *InputService9TestShapeInputService9TestCaseOperation1Output) {
	if opInputService9TestCaseOperation1 == nil {
		opInputService9TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService9TestCaseOperation1, input, output)
	output = &InputService9TestShapeInputService9TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService9ProtocolTest) InputService9TestCaseOperation1(input *InputService9TestShapeInputShape) (output *InputService9TestShapeInputService9TestCaseOperation1Output, err error) {
	req, out := c.InputService9TestCaseOperation1Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService9TestCaseOperation1 *aws.Operation

// InputService9TestCaseOperation2Request generates a request for the InputService9TestCaseOperation2 operation.
func (c *InputService9ProtocolTest) InputService9TestCaseOperation2Request(input *InputService9TestShapeInputShape) (req *aws.Request, output *InputService9TestShapeInputService9TestCaseOperation2Output) {
	if opInputService9TestCaseOperation2 == nil {
		opInputService9TestCaseOperation2 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService9TestCaseOperation2, input, output)
	output = &InputService9TestShapeInputService9TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService9ProtocolTest) InputService9TestCaseOperation2(input *InputService9TestShapeInputShape) (output *InputService9TestShapeInputService9TestCaseOperation2Output, err error) {
	req, out := c.InputService9TestCaseOperation2Request(input)
	output = out
	err = req.Send()
	return
}

var opInputService9TestCaseOperation2 *aws.Operation

type InputService9TestShapeInputService9TestCaseOperation1Output struct {
	metadataInputService9TestShapeInputService9TestCaseOperation1Output `json:"-", xml:"-"`
}

type metadataInputService9TestShapeInputService9TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService9TestShapeInputService9TestCaseOperation2Output struct {
	metadataInputService9TestShapeInputService9TestCaseOperation2Output `json:"-", xml:"-"`
}

type metadataInputService9TestShapeInputService9TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InputService9TestShapeInputShape struct {
	TimeArg         *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	TimeArgInHeader *time.Time `location:"header" locationName:"x-amz-timearg" type:"timestamp" timestampFormat:"rfc822" json:"-" xml:"-"`

	metadataInputService9TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService9TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

//
// Tests begin here
//

func TestInputService1ProtocolTestURIParameterOnlyWithNoLocationNameCase1(t *testing.T) {
	svc := NewInputService1ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService1TestShapeInputService1TestCaseOperation1Input{
		PipelineId: aws.String("foo"),
	}
	req, _ := svc.InputService1TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo", r.URL.String())

	// assert headers

}

func TestInputService2ProtocolTestURIParameterOnlyWithLocationNameCase1(t *testing.T) {
	svc := NewInputService2ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService2TestShapeInputService2TestCaseOperation1Input{
		Foo: aws.String("bar"),
	}
	req, _ := svc.InputService2TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/bar", r.URL.String())

	// assert headers

}

func TestInputService3ProtocolTestURIParameterAndQuerystringParamsCase1(t *testing.T) {
	svc := NewInputService3ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService3TestShapeInputService3TestCaseOperation1Input{
		Ascending:  aws.String("true"),
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req, _ := svc.InputService3TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers

}

func TestInputService4ProtocolTestURIParameterQuerystringParamsAndJSONBodyCase1(t *testing.T) {
	svc := NewInputService4ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService4TestShapeInputService4TestCaseOperation1Input{
		Ascending: aws.String("true"),
		Config: &InputService4TestShapeStructType{
			A: aws.String("one"),
			B: aws.String("two"),
		},
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req, _ := svc.InputService4TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"Config":{"A":"one","B":"two"}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers

}

func TestInputService5ProtocolTestURIParameterQuerystringParamsHeadersAndJSONBodyCase1(t *testing.T) {
	svc := NewInputService5ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService5TestShapeInputService5TestCaseOperation1Input{
		Ascending: aws.String("true"),
		Checksum:  aws.String("12345"),
		Config: &InputService5TestShapeStructType{
			A: aws.String("one"),
			B: aws.String("two"),
		},
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req, _ := svc.InputService5TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"Config":{"A":"one","B":"two"}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers
	assert.Equal(t, "12345", r.Header.Get("x-amz-checksum"))

}

func TestInputService6ProtocolTestStreamingPayloadCase1(t *testing.T) {
	svc := NewInputService6ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService6TestShapeInputService6TestCaseOperation1Input{
		Body:      []byte("contents"),
		Checksum:  aws.String("foo"),
		VaultName: aws.String("name"),
	}
	req, _ := svc.InputService6TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`contents`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/vaults/name/archives", r.URL.String())

	// assert headers
	assert.Equal(t, "foo", r.Header.Get("x-amz-sha256-tree-hash"))

}

func TestInputService7ProtocolTestOmitsNullQueryParamsButSerializesEmptyStringsCase1(t *testing.T) {
	svc := NewInputService7ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService7TestShapeInputShape{}
	req, _ := svc.InputService7TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService7ProtocolTestOmitsNullQueryParamsButSerializesEmptyStringsCase2(t *testing.T) {
	svc := NewInputService7ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService7TestShapeInputShape{
		Foo: aws.String(""),
	}
	req, _ := svc.InputService7TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/path?abc=mno&param-name=", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestRecursiveShapesCase1(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{
		RecursiveStruct: &InputService8TestShapeRecursiveStructType{
			NoRecurse: aws.String("foo"),
		},
	}
	req, _ := svc.InputService8TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"NoRecurse":"foo"}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestRecursiveShapesCase2(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{
		RecursiveStruct: &InputService8TestShapeRecursiveStructType{
			RecursiveStruct: &InputService8TestShapeRecursiveStructType{
				NoRecurse: aws.String("foo"),
			},
		},
	}
	req, _ := svc.InputService8TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveStruct":{"NoRecurse":"foo"}}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestRecursiveShapesCase3(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{
		RecursiveStruct: &InputService8TestShapeRecursiveStructType{
			RecursiveStruct: &InputService8TestShapeRecursiveStructType{
				RecursiveStruct: &InputService8TestShapeRecursiveStructType{
					RecursiveStruct: &InputService8TestShapeRecursiveStructType{
						NoRecurse: aws.String("foo"),
					},
				},
			},
		},
	}
	req, _ := svc.InputService8TestCaseOperation3Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveStruct":{"RecursiveStruct":{"RecursiveStruct":{"NoRecurse":"foo"}}}}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestRecursiveShapesCase4(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{
		RecursiveStruct: &InputService8TestShapeRecursiveStructType{
			RecursiveList: []*InputService8TestShapeRecursiveStructType{
				&InputService8TestShapeRecursiveStructType{
					NoRecurse: aws.String("foo"),
				},
				&InputService8TestShapeRecursiveStructType{
					NoRecurse: aws.String("bar"),
				},
			},
		},
	}
	req, _ := svc.InputService8TestCaseOperation4Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveList":[{"NoRecurse":"foo"},{"NoRecurse":"bar"}]}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestRecursiveShapesCase5(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{
		RecursiveStruct: &InputService8TestShapeRecursiveStructType{
			RecursiveList: []*InputService8TestShapeRecursiveStructType{
				&InputService8TestShapeRecursiveStructType{
					NoRecurse: aws.String("foo"),
				},
				&InputService8TestShapeRecursiveStructType{
					RecursiveStruct: &InputService8TestShapeRecursiveStructType{
						NoRecurse: aws.String("bar"),
					},
				},
			},
		},
	}
	req, _ := svc.InputService8TestCaseOperation5Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveList":[{"NoRecurse":"foo"},{"RecursiveStruct":{"NoRecurse":"bar"}}]}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestRecursiveShapesCase6(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{
		RecursiveStruct: &InputService8TestShapeRecursiveStructType{
			RecursiveMap: &map[string]*InputService8TestShapeRecursiveStructType{
				"bar": &InputService8TestShapeRecursiveStructType{
					NoRecurse: aws.String("bar"),
				},
				"foo": &InputService8TestShapeRecursiveStructType{
					NoRecurse: aws.String("foo"),
				},
			},
		},
	}
	req, _ := svc.InputService8TestCaseOperation6Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveMap":{"bar":{"NoRecurse":"bar"},"foo":{"NoRecurse":"foo"}}}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService9ProtocolTestTimestampValuesCase1(t *testing.T) {
	svc := NewInputService9ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService9TestShapeInputShape{
		TimeArg: aws.Time(time.Unix(1422172800, 0)),
	}
	req, _ := svc.InputService9TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"TimeArg":1422172800}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService9ProtocolTestTimestampValuesCase2(t *testing.T) {
	svc := NewInputService9ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService9TestShapeInputShape{
		TimeArgInHeader: aws.Time(time.Unix(1422172800, 0)),
	}
	req, _ := svc.InputService9TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers
	assert.Equal(t, "Sun, 25 Jan 2015 08:00:00 GMT", r.Header.Get("x-amz-timearg"))

}

