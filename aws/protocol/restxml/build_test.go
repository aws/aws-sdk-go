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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService1ProtocolTest{service}
}

// InputService1TestCaseOperation1Request generates a request for the InputService1TestCaseOperation1 operation.
func (c *InputService1ProtocolTest) InputService1TestCaseOperation1Request(input *InputService1TestShapeInputShape) (req *aws.Request) {
	if opInputService1TestCaseOperation1 == nil {
		opInputService1TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
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

// InputService1TestCaseOperation2Request generates a request for the InputService1TestCaseOperation2 operation.
func (c *InputService1ProtocolTest) InputService1TestCaseOperation2Request(input *InputService1TestShapeInputShape) (req *aws.Request) {
	if opInputService1TestCaseOperation2 == nil {
		opInputService1TestCaseOperation2 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "PUT",
			HTTPPath:   "/2014-01-01/hostedzone",
		}
	}

	req = aws.NewRequest(c.Service, opInputService1TestCaseOperation2, input, nil)

	return
}

func (c *InputService1ProtocolTest) InputService1TestCaseOperation2(input *InputService1TestShapeInputShape) (err error) {
	req := c.InputService1TestCaseOperation2Request(input)
	err = req.Send()
	return
}

var opInputService1TestCaseOperation2 *aws.Operation

type InputService1TestShapeInputShape struct {
	Description *string `type:"string"`
	Name        *string `type:"string"`

	metadataInputService1TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService1TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService2ProtocolTest{service}
}

// InputService2TestCaseOperation1Request generates a request for the InputService2TestCaseOperation1 operation.
func (c *InputService2ProtocolTest) InputService2TestCaseOperation1Request(input *InputService2TestShapeInputShape) (req *aws.Request) {
	if opInputService2TestCaseOperation1 == nil {
		opInputService2TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
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
	First  *bool    `type:"boolean"`
	Fourth *int     `type:"integer"`
	Second *bool    `type:"boolean"`
	Third  *float32 `type:"float"`

	metadataInputService2TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService2TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService3ProtocolTest{service}
}

// InputService3TestCaseOperation1Request generates a request for the InputService3TestCaseOperation1 operation.
func (c *InputService3ProtocolTest) InputService3TestCaseOperation1Request(input *InputService3TestShapeInputShape) (req *aws.Request) {
	if opInputService3TestCaseOperation1 == nil {
		opInputService3TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
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
	Description  *string                             `type:"string"`
	SubStructure *InputService3TestShapeSubStructure `type:"structure"`

	metadataInputService3TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService3TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
}

type InputService3TestShapeSubStructure struct {
	Bar *string `type:"string"`
	Foo *string `type:"string"`

	metadataInputService3TestShapeSubStructure `json:"-", xml:"-"`
}

type metadataInputService3TestShapeSubStructure struct {
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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService4ProtocolTest{service}
}

// InputService4TestCaseOperation1Request generates a request for the InputService4TestCaseOperation1 operation.
func (c *InputService4ProtocolTest) InputService4TestCaseOperation1Request(input *InputService4TestShapeInputShape) (req *aws.Request) {
	if opInputService4TestCaseOperation1 == nil {
		opInputService4TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
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
	Description  *string                             `type:"string"`
	SubStructure *InputService4TestShapeSubStructure `type:"structure"`

	metadataInputService4TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService4TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
}

type InputService4TestShapeSubStructure struct {
	Bar *string `type:"string"`
	Foo *string `type:"string"`

	metadataInputService4TestShapeSubStructure `json:"-", xml:"-"`
}

type metadataInputService4TestShapeSubStructure struct {
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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService5ProtocolTest{service}
}

// InputService5TestCaseOperation1Request generates a request for the InputService5TestCaseOperation1 operation.
func (c *InputService5ProtocolTest) InputService5TestCaseOperation1Request(input *InputService5TestShapeInputShape) (req *aws.Request) {
	if opInputService5TestCaseOperation1 == nil {
		opInputService5TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
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
	ListParam []*string `type:"list"`

	metadataInputService5TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService5TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService6ProtocolTest{service}
}

// InputService6TestCaseOperation1Request generates a request for the InputService6TestCaseOperation1 operation.
func (c *InputService6ProtocolTest) InputService6TestCaseOperation1Request(input *InputService6TestShapeInputShape) (req *aws.Request) {
	if opInputService6TestCaseOperation1 == nil {
		opInputService6TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
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
	ListParam []*string `locationName:"AlternateName" locationNameList:"NotMember" type:"list"`

	metadataInputService6TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService6TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService7ProtocolTest{service}
}

// InputService7TestCaseOperation1Request generates a request for the InputService7TestCaseOperation1 operation.
func (c *InputService7ProtocolTest) InputService7TestCaseOperation1Request(input *InputService7TestShapeInputShape) (req *aws.Request) {
	if opInputService7TestCaseOperation1 == nil {
		opInputService7TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
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
	ListParam []*string `type:"list" flattened:"true"`

	metadataInputService7TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService7TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService8ProtocolTest{service}
}

// InputService8TestCaseOperation1Request generates a request for the InputService8TestCaseOperation1 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation1Request(input *InputService8TestShapeInputShape) (req *aws.Request) {
	if opInputService8TestCaseOperation1 == nil {
		opInputService8TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
		}
	}

	req = aws.NewRequest(c.Service, opInputService8TestCaseOperation1, input, nil)

	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation1(input *InputService8TestShapeInputShape) (err error) {
	req := c.InputService8TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService8TestCaseOperation1 *aws.Operation

type InputService8TestShapeInputShape struct {
	ListParam []*string `locationName:"item" type:"list" flattened:"true"`

	metadataInputService8TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService8TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
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
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService9ProtocolTest{service}
}

// InputService9TestCaseOperation1Request generates a request for the InputService9TestCaseOperation1 operation.
func (c *InputService9ProtocolTest) InputService9TestCaseOperation1Request(input *InputService9TestShapeInputShape) (req *aws.Request) {
	if opInputService9TestCaseOperation1 == nil {
		opInputService9TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
		}
	}

	req = aws.NewRequest(c.Service, opInputService9TestCaseOperation1, input, nil)

	return
}

func (c *InputService9ProtocolTest) InputService9TestCaseOperation1(input *InputService9TestShapeInputShape) (err error) {
	req := c.InputService9TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService9TestCaseOperation1 *aws.Operation

type InputService9TestShapeInputShape struct {
	ListParam []*InputService9TestShapeSingleFieldStruct `locationName:"item" type:"list" flattened:"true"`

	metadataInputService9TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService9TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
}

type InputService9TestShapeSingleFieldStruct struct {
	Element *string `locationName:"value" type:"string"`

	metadataInputService9TestShapeSingleFieldStruct `json:"-", xml:"-"`
}

type metadataInputService9TestShapeSingleFieldStruct struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService10ProtocolTest is a client for InputService10ProtocolTest.
type InputService10ProtocolTest struct {
	*aws.Service
}

type InputService10ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService10ProtocolTest client.
func NewInputService10ProtocolTest(config *InputService10ProtocolTestConfig) *InputService10ProtocolTest {
	if config == nil {
		config = &InputService10ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice10protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService10ProtocolTest{service}
}

// InputService10TestCaseOperation1Request generates a request for the InputService10TestCaseOperation1 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation1Request(input *InputService10TestShapeInputShape) (req *aws.Request) {
	if opInputService10TestCaseOperation1 == nil {
		opInputService10TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-01-01/hostedzone",
		}
	}

	req = aws.NewRequest(c.Service, opInputService10TestCaseOperation1, input, nil)

	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation1(input *InputService10TestShapeInputShape) (err error) {
	req := c.InputService10TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService10TestCaseOperation1 *aws.Operation

type InputService10TestShapeInputShape struct {
	StructureParam *InputService10TestShapeStructureShape `type:"structure"`

	metadataInputService10TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService10TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
}

type InputService10TestShapeStructureShape struct {
	B []byte     `locationName:"b" type:"blob"`
	T *time.Time `locationName:"t" type:"timestamp" timestampFormat:"iso8601"`

	metadataInputService10TestShapeStructureShape `json:"-", xml:"-"`
}

type metadataInputService10TestShapeStructureShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService11ProtocolTest is a client for InputService11ProtocolTest.
type InputService11ProtocolTest struct {
	*aws.Service
}

type InputService11ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService11ProtocolTest client.
func NewInputService11ProtocolTest(config *InputService11ProtocolTestConfig) *InputService11ProtocolTest {
	if config == nil {
		config = &InputService11ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice11protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService11ProtocolTest{service}
}

// InputService11TestCaseOperation1Request generates a request for the InputService11TestCaseOperation1 operation.
func (c *InputService11ProtocolTest) InputService11TestCaseOperation1Request(input *InputService11TestShapeInputShape) (req *aws.Request) {
	if opInputService11TestCaseOperation1 == nil {
		opInputService11TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opInputService11TestCaseOperation1, input, nil)

	return
}

func (c *InputService11ProtocolTest) InputService11TestCaseOperation1(input *InputService11TestShapeInputShape) (err error) {
	req := c.InputService11TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService11TestCaseOperation1 *aws.Operation

type InputService11TestShapeInputShape struct {
	Foo *map[string]*string `location:"headers" locationName:"x-foo-" type:"map"`

	metadataInputService11TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService11TestShapeInputShape struct {
	SDKShapeTraits bool `locationName:"OperationRequest" type:"structure" xmlURI:"https://foo/"`
}

// InputService12ProtocolTest is a client for InputService12ProtocolTest.
type InputService12ProtocolTest struct {
	*aws.Service
}

type InputService12ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService12ProtocolTest client.
func NewInputService12ProtocolTest(config *InputService12ProtocolTestConfig) *InputService12ProtocolTest {
	if config == nil {
		config = &InputService12ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice12protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService12ProtocolTest{service}
}

// InputService12TestCaseOperation1Request generates a request for the InputService12TestCaseOperation1 operation.
func (c *InputService12ProtocolTest) InputService12TestCaseOperation1Request(input *InputService12TestShapeInputShape) (req *aws.Request) {
	if opInputService12TestCaseOperation1 == nil {
		opInputService12TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opInputService12TestCaseOperation1, input, nil)

	return
}

func (c *InputService12ProtocolTest) InputService12TestCaseOperation1(input *InputService12TestShapeInputShape) (err error) {
	req := c.InputService12TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService12TestCaseOperation1 *aws.Operation

type InputService12TestShapeInputShape struct {
	Foo *string `locationName:"foo" type:"string"`

	metadataInputService12TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService12TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Foo"`
}

// InputService13ProtocolTest is a client for InputService13ProtocolTest.
type InputService13ProtocolTest struct {
	*aws.Service
}

type InputService13ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService13ProtocolTest client.
func NewInputService13ProtocolTest(config *InputService13ProtocolTestConfig) *InputService13ProtocolTest {
	if config == nil {
		config = &InputService13ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice13protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService13ProtocolTest{service}
}

// InputService13TestCaseOperation1Request generates a request for the InputService13TestCaseOperation1 operation.
func (c *InputService13ProtocolTest) InputService13TestCaseOperation1Request(input *InputService13TestShapeInputShape) (req *aws.Request) {
	if opInputService13TestCaseOperation1 == nil {
		opInputService13TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opInputService13TestCaseOperation1, input, nil)

	return
}

func (c *InputService13ProtocolTest) InputService13TestCaseOperation1(input *InputService13TestShapeInputShape) (err error) {
	req := c.InputService13TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService13TestCaseOperation1 *aws.Operation

type InputService13TestShapeInputShape struct {
	Foo []byte `locationName:"foo" type:"blob"`

	metadataInputService13TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService13TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Foo"`
}

// InputService14ProtocolTest is a client for InputService14ProtocolTest.
type InputService14ProtocolTest struct {
	*aws.Service
}

type InputService14ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService14ProtocolTest client.
func NewInputService14ProtocolTest(config *InputService14ProtocolTestConfig) *InputService14ProtocolTest {
	if config == nil {
		config = &InputService14ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice14protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService14ProtocolTest{service}
}

// InputService14TestCaseOperation1Request generates a request for the InputService14TestCaseOperation1 operation.
func (c *InputService14ProtocolTest) InputService14TestCaseOperation1Request(input *InputService14TestShapeInputShape) (req *aws.Request) {
	if opInputService14TestCaseOperation1 == nil {
		opInputService14TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opInputService14TestCaseOperation1, input, nil)

	return
}

func (c *InputService14ProtocolTest) InputService14TestCaseOperation1(input *InputService14TestShapeInputShape) (err error) {
	req := c.InputService14TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService14TestCaseOperation1 *aws.Operation

// InputService14TestCaseOperation2Request generates a request for the InputService14TestCaseOperation2 operation.
func (c *InputService14ProtocolTest) InputService14TestCaseOperation2Request(input *InputService14TestShapeInputShape) (req *aws.Request) {
	if opInputService14TestCaseOperation2 == nil {
		opInputService14TestCaseOperation2 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opInputService14TestCaseOperation2, input, nil)

	return
}

func (c *InputService14ProtocolTest) InputService14TestCaseOperation2(input *InputService14TestShapeInputShape) (err error) {
	req := c.InputService14TestCaseOperation2Request(input)
	err = req.Send()
	return
}

var opInputService14TestCaseOperation2 *aws.Operation

type InputService14TestShapeFooShape struct {
	Baz *string `locationName:"baz" type:"string"`

	metadataInputService14TestShapeFooShape `json:"-", xml:"-"`
}

type metadataInputService14TestShapeFooShape struct {
	SDKShapeTraits bool `locationName:"foo" type:"structure"`
}

type InputService14TestShapeInputShape struct {
	Foo *InputService14TestShapeFooShape `locationName:"foo" type:"structure"`

	metadataInputService14TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService14TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Foo"`
}

// InputService15ProtocolTest is a client for InputService15ProtocolTest.
type InputService15ProtocolTest struct {
	*aws.Service
}

type InputService15ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService15ProtocolTest client.
func NewInputService15ProtocolTest(config *InputService15ProtocolTestConfig) *InputService15ProtocolTest {
	if config == nil {
		config = &InputService15ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice15protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService15ProtocolTest{service}
}

// InputService15TestCaseOperation1Request generates a request for the InputService15TestCaseOperation1 operation.
func (c *InputService15ProtocolTest) InputService15TestCaseOperation1Request(input *InputService15TestShapeInputShape) (req *aws.Request) {
	if opInputService15TestCaseOperation1 == nil {
		opInputService15TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opInputService15TestCaseOperation1, input, nil)

	return
}

func (c *InputService15ProtocolTest) InputService15TestCaseOperation1(input *InputService15TestShapeInputShape) (err error) {
	req := c.InputService15TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService15TestCaseOperation1 *aws.Operation

type InputService15TestShapeGrant struct {
	Grantee *InputService15TestShapeGrantee `type:"structure"`

	metadataInputService15TestShapeGrant `json:"-", xml:"-"`
}

type metadataInputService15TestShapeGrant struct {
	SDKShapeTraits bool `locationName:"Grant" type:"structure"`
}

type InputService15TestShapeGrantee struct {
	EmailAddress *string `type:"string"`
	Type         *string `locationName:"xsi:type" type:"string" xmlAttribute:"true"`

	metadataInputService15TestShapeGrantee `json:"-", xml:"-"`
}

type metadataInputService15TestShapeGrantee struct {
	SDKShapeTraits bool `type:"structure" xmlPrefix:"xsi" xmlURI:"http://www.w3.org/2001/XMLSchema-instance"`
}

type InputService15TestShapeInputShape struct {
	Grant *InputService15TestShapeGrant `locationName:"Grant" type:"structure"`

	metadataInputService15TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService15TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Grant"`
}

// InputService16ProtocolTest is a client for InputService16ProtocolTest.
type InputService16ProtocolTest struct {
	*aws.Service
}

type InputService16ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService16ProtocolTest client.
func NewInputService16ProtocolTest(config *InputService16ProtocolTestConfig) *InputService16ProtocolTest {
	if config == nil {
		config = &InputService16ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice16protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService16ProtocolTest{service}
}

// InputService16TestCaseOperation1Request generates a request for the InputService16TestCaseOperation1 operation.
func (c *InputService16ProtocolTest) InputService16TestCaseOperation1Request(input *InputService16TestShapeInputShape) (req *aws.Request) {
	if opInputService16TestCaseOperation1 == nil {
		opInputService16TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/{Bucket}/{Key+}",
		}
	}

	req = aws.NewRequest(c.Service, opInputService16TestCaseOperation1, input, nil)

	return
}

func (c *InputService16ProtocolTest) InputService16TestCaseOperation1(input *InputService16TestShapeInputShape) (err error) {
	req := c.InputService16TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService16TestCaseOperation1 *aws.Operation

type InputService16TestShapeInputShape struct {
	Bucket *string `location:"uri" type:"string" json:"-" xml:"-"`
	Key    *string `location:"uri" type:"string" json:"-" xml:"-"`

	metadataInputService16TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService16TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

// InputService17ProtocolTest is a client for InputService17ProtocolTest.
type InputService17ProtocolTest struct {
	*aws.Service
}

type InputService17ProtocolTestConfig struct {
	*aws.Config
}

// New returns a new InputService17ProtocolTest client.
func NewInputService17ProtocolTest(config *InputService17ProtocolTestConfig) *InputService17ProtocolTest {
	if config == nil {
		config = &InputService17ProtocolTestConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "inputservice17protocoltest",
		APIVersion:  "2014-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &InputService17ProtocolTest{service}
}

// InputService17TestCaseOperation1Request generates a request for the InputService17TestCaseOperation1 operation.
func (c *InputService17ProtocolTest) InputService17TestCaseOperation1Request(input *InputService17TestShapeInputShape) (req *aws.Request) {
	if opInputService17TestCaseOperation1 == nil {
		opInputService17TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
		}
	}

	req = aws.NewRequest(c.Service, opInputService17TestCaseOperation1, input, nil)

	return
}

func (c *InputService17ProtocolTest) InputService17TestCaseOperation1(input *InputService17TestShapeInputShape) (err error) {
	req := c.InputService17TestCaseOperation1Request(input)
	err = req.Send()
	return
}

var opInputService17TestCaseOperation1 *aws.Operation

// InputService17TestCaseOperation2Request generates a request for the InputService17TestCaseOperation2 operation.
func (c *InputService17ProtocolTest) InputService17TestCaseOperation2Request(input *InputService17TestShapeInputShape) (req *aws.Request) {
	if opInputService17TestCaseOperation2 == nil {
		opInputService17TestCaseOperation2 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path?abc=mno",
		}
	}

	req = aws.NewRequest(c.Service, opInputService17TestCaseOperation2, input, nil)

	return
}

func (c *InputService17ProtocolTest) InputService17TestCaseOperation2(input *InputService17TestShapeInputShape) (err error) {
	req := c.InputService17TestCaseOperation2Request(input)
	err = req.Send()
	return
}

var opInputService17TestCaseOperation2 *aws.Operation

type InputService17TestShapeInputShape struct {
	Foo *string `location:"querystring" locationName:"param-name" type:"string" json:"-" xml:"-"`

	metadataInputService17TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService17TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

//
// Tests begin here
//

func TestInputService1ProtocolTestBasicXMLSerializationCase1(t *testing.T) {
	svc := NewInputService1ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService1TestShapeInputShape{
		Description: aws.String("bar"),
		Name:        aws.String("foo"),
	}
	req := svc.InputService1TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><Description xmlns=\"https://foo/\">bar</Description><Name xmlns=\"https://foo/\">foo</Name></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService1ProtocolTestBasicXMLSerializationCase2(t *testing.T) {
	svc := NewInputService1ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService1TestShapeInputShape{
		Description: aws.String("bar"),
		Name:        aws.String("foo"),
	}
	req := svc.InputService1TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><Description xmlns=\"https://foo/\">bar</Description><Name xmlns=\"https://foo/\">foo</Name></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService2ProtocolTestSerializeOtherScalarTypesCase1(t *testing.T) {
	svc := NewInputService2ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService2TestShapeInputShape{
		First:  aws.Boolean(true),
		Fourth: aws.Integer(3),
		Second: aws.Boolean(false),
		Third:  aws.Float(1.2),
	}
	req := svc.InputService2TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><First xmlns=\"https://foo/\">true</First><Fourth xmlns=\"https://foo/\">3</Fourth><Second xmlns=\"https://foo/\">false</Second><Third xmlns=\"https://foo/\">1.2</Third></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService3ProtocolTestNestedStructuresCase1(t *testing.T) {
	svc := NewInputService3ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService3TestShapeInputShape{
		Description: aws.String("baz"),
		SubStructure: &InputService3TestShapeSubStructure{
			Bar: aws.String("b"),
			Foo: aws.String("a"),
		},
	}
	req := svc.InputService3TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><Description xmlns=\"https://foo/\">baz</Description><SubStructure xmlns=\"https://foo/\"><Bar xmlns=\"https://foo/\">b</Bar><Foo xmlns=\"https://foo/\">a</Foo></SubStructure></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService4ProtocolTestNestedStructuresCase1(t *testing.T) {
	svc := NewInputService4ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService4TestShapeInputShape{
		Description:  aws.String("baz"),
		SubStructure: &InputService4TestShapeSubStructure{},
	}
	req := svc.InputService4TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><Description xmlns=\"https://foo/\">baz</Description><SubStructure xmlns=\"https://foo/\"></SubStructure></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService5ProtocolTestNonFlattenedListsCase1(t *testing.T) {
	svc := NewInputService5ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService5TestShapeInputShape{
		ListParam: []*string{
			aws.String("one"),
			aws.String("two"),
			aws.String("three"),
		},
	}
	req := svc.InputService5TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><ListParam xmlns=\"https://foo/\"><member xmlns=\"https://foo/\">one</member><member xmlns=\"https://foo/\">two</member><member xmlns=\"https://foo/\">three</member></ListParam></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService6ProtocolTestNonFlattenedListsWithLocationNameCase1(t *testing.T) {
	svc := NewInputService6ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService6TestShapeInputShape{
		ListParam: []*string{
			aws.String("one"),
			aws.String("two"),
			aws.String("three"),
		},
	}
	req := svc.InputService6TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><AlternateName xmlns=\"https://foo/\"><NotMember xmlns=\"https://foo/\">one</NotMember><NotMember xmlns=\"https://foo/\">two</NotMember><NotMember xmlns=\"https://foo/\">three</NotMember></AlternateName></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService7ProtocolTestFlattenedListsCase1(t *testing.T) {
	svc := NewInputService7ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService7TestShapeInputShape{
		ListParam: []*string{
			aws.String("one"),
			aws.String("two"),
			aws.String("three"),
		},
	}
	req := svc.InputService7TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><ListParam xmlns=\"https://foo/\">one</ListParam><ListParam xmlns=\"https://foo/\">two</ListParam><ListParam xmlns=\"https://foo/\">three</ListParam></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestFlattenedListsWithLocationNameCase1(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{
		ListParam: []*string{
			aws.String("one"),
			aws.String("two"),
			aws.String("three"),
		},
	}
	req := svc.InputService8TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><item xmlns=\"https://foo/\">one</item><item xmlns=\"https://foo/\">two</item><item xmlns=\"https://foo/\">three</item></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService9ProtocolTestListOfStructuresCase1(t *testing.T) {
	svc := NewInputService9ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService9TestShapeInputShape{
		ListParam: []*InputService9TestShapeSingleFieldStruct{
			&InputService9TestShapeSingleFieldStruct{
				Element: aws.String("one"),
			},
			&InputService9TestShapeSingleFieldStruct{
				Element: aws.String("two"),
			},
			&InputService9TestShapeSingleFieldStruct{
				Element: aws.String("three"),
			},
		},
	}
	req := svc.InputService9TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><item xmlns=\"https://foo/\"><value xmlns=\"https://foo/\">one</value></item><item xmlns=\"https://foo/\"><value xmlns=\"https://foo/\">two</value></item><item xmlns=\"https://foo/\"><value xmlns=\"https://foo/\">three</value></item></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestBlobAndTimestampShapesCase1(t *testing.T) {
	svc := NewInputService10ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService10TestShapeInputShape{
		StructureParam: &InputService10TestShapeStructureShape{
			B: []byte("foo"),
			T: aws.Time(time.Unix(1422172800, 0)),
		},
	}
	req := svc.InputService10TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<OperationRequest xmlns=\"https://foo/\"><StructureParam xmlns=\"https://foo/\"><b xmlns=\"https://foo/\">Zm9v</b><t xmlns=\"https://foo/\">2015-01-25T08:00:00Z</t></StructureParam></OperationRequest>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/hostedzone", r.URL.String())

	// assert headers

}

func TestInputService11ProtocolTestHeaderMapsCase1(t *testing.T) {
	svc := NewInputService11ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService11TestShapeInputShape{
		Foo: &map[string]*string{
			"a": aws.String("b"),
			"c": aws.String("d"),
		},
	}
	req := svc.InputService11TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers
	assert.Equal(t, "b", r.Header.Get("x-foo-a"))
	assert.Equal(t, "d", r.Header.Get("x-foo-c"))

}

func TestInputService12ProtocolTestStringPayloadCase1(t *testing.T) {
	svc := NewInputService12ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService12TestShapeInputShape{
		Foo: aws.String("bar"),
	}
	req := svc.InputService12TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("bar"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService13ProtocolTestBlobPayloadCase1(t *testing.T) {
	svc := NewInputService13ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService13TestShapeInputShape{
		Foo: []byte("bar"),
	}
	req := svc.InputService13TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("bar"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService14ProtocolTestStructurePayloadCase1(t *testing.T) {
	svc := NewInputService14ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService14TestShapeInputShape{
		Foo: &InputService14TestShapeFooShape{
			Baz: aws.String("bar"),
		},
	}
	req := svc.InputService14TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<foo><baz>bar</baz></foo>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService14ProtocolTestStructurePayloadCase2(t *testing.T) {
	svc := NewInputService14ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService14TestShapeInputShape{}
	req := svc.InputService14TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService15ProtocolTestXMLAttributeCase1(t *testing.T) {
	svc := NewInputService15ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService15TestShapeInputShape{
		Grant: &InputService15TestShapeGrant{
			Grantee: &InputService15TestShapeGrantee{
				EmailAddress: aws.String("foo@example.com"),
				Type:         aws.String("CanonicalUser"),
			},
		},
	}
	req := svc.InputService15TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim("<Grant xmlns:_xmlns=\"xmlns\" _xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:XMLSchema-instance=\"http://www.w3.org/2001/XMLSchema-instance\" XMLSchema-instance:type=\"CanonicalUser\"><Grantee><EmailAddress>foo@example.com</EmailAddress></Grantee></Grant>"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService16ProtocolTestGreedyKeysCase1(t *testing.T) {
	svc := NewInputService16ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService16TestShapeInputShape{
		Bucket: aws.String("my/bucket"),
		Key:    aws.String("testing /123"),
	}
	req := svc.InputService16TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/my%2Fbucket/testing%20/123", r.URL.String())

	// assert headers

}

func TestInputService17ProtocolTestOmitsNullQueryParamsButSerializesEmptyStringsCase1(t *testing.T) {
	svc := NewInputService17ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService17TestShapeInputShape{}
	req := svc.InputService17TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService17ProtocolTestOmitsNullQueryParamsButSerializesEmptyStringsCase2(t *testing.T) {
	svc := NewInputService17ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService17TestShapeInputShape{
		Foo: aws.String(""),
	}
	req := svc.InputService17TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restxml.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body := util.SortXML(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path?abc=mno&param-name=", r.URL.String())

	// assert headers

}

