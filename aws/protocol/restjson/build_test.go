package restjson_test

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/restjson"
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)

	return &InputService1ProtocolTest{service}
}

// InputService1TestCaseOperation1Request generates a request for the InputService1TestCaseOperation1 operation.
func (c *InputService1ProtocolTest) InputService1TestCaseOperation1Request(input *InputService1TestShapeInputShape) (req *aws.Request) {
	if opInputService1TestCaseOperation1 == nil {
		opInputService1TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
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
	PipelineId *string `location:"uri" type:"string" json:"-" xml:"-"`

	metadataInputService1TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService1TestShapeInputShape struct {
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

	return &InputService2ProtocolTest{service}
}

// InputService2TestCaseOperation1Request generates a request for the InputService2TestCaseOperation1 operation.
func (c *InputService2ProtocolTest) InputService2TestCaseOperation1Request(input *InputService2TestShapeInputShape) (req *aws.Request) {
	if opInputService2TestCaseOperation1 == nil {
		opInputService2TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
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
	Foo *string `location:"uri" locationName:"PipelineId" type:"string" json:"-" xml:"-"`

	metadataInputService2TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService2TestShapeInputShape struct {
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

	return &InputService3ProtocolTest{service}
}

// InputService3TestCaseOperation1Request generates a request for the InputService3TestCaseOperation1 operation.
func (c *InputService3ProtocolTest) InputService3TestCaseOperation1Request(input *InputService3TestShapeInputShape) (req *aws.Request) {
	if opInputService3TestCaseOperation1 == nil {
		opInputService3TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
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
	Ascending  *string `location:"querystring" locationName:"Ascending" type:"string" json:"-" xml:"-"`
	PageToken  *string `location:"querystring" locationName:"PageToken" type:"string" json:"-" xml:"-"`
	PipelineId *string `location:"uri" locationName:"PipelineId" type:"string" json:"-" xml:"-"`

	metadataInputService3TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService3TestShapeInputShape struct {
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

	return &InputService4ProtocolTest{service}
}

// InputService4TestCaseOperation1Request generates a request for the InputService4TestCaseOperation1 operation.
func (c *InputService4ProtocolTest) InputService4TestCaseOperation1Request(input *InputService4TestShapeInputShape) (req *aws.Request) {
	if opInputService4TestCaseOperation1 == nil {
		opInputService4TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
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
	Ascending  *string                           `location:"querystring" locationName:"Ascending" type:"string" json:"-" xml:"-"`
	Config     *InputService4TestShapeStructType `type:"structure" json:",omitempty"`
	PageToken  *string                           `location:"querystring" locationName:"PageToken" type:"string" json:"-" xml:"-"`
	PipelineId *string                           `location:"uri" locationName:"PipelineId" type:"string" json:"-" xml:"-"`

	metadataInputService4TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService4TestShapeInputShape struct {
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

	return &InputService5ProtocolTest{service}
}

// InputService5TestCaseOperation1Request generates a request for the InputService5TestCaseOperation1 operation.
func (c *InputService5ProtocolTest) InputService5TestCaseOperation1Request(input *InputService5TestShapeInputShape) (req *aws.Request) {
	if opInputService5TestCaseOperation1 == nil {
		opInputService5TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
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
	Ascending  *string                           `location:"querystring" locationName:"Ascending" type:"string" json:"-" xml:"-"`
	Checksum   *string                           `location:"header" locationName:"x-amz-checksum" type:"string" json:"-" xml:"-"`
	Config     *InputService5TestShapeStructType `type:"structure" json:",omitempty"`
	PageToken  *string                           `location:"querystring" locationName:"PageToken" type:"string" json:"-" xml:"-"`
	PipelineId *string                           `location:"uri" locationName:"PipelineId" type:"string" json:"-" xml:"-"`

	metadataInputService5TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService5TestShapeInputShape struct {
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

	return &InputService6ProtocolTest{service}
}

// InputService6TestCaseOperation1Request generates a request for the InputService6TestCaseOperation1 operation.
func (c *InputService6ProtocolTest) InputService6TestCaseOperation1Request(input *InputService6TestShapeInputShape) (req *aws.Request) {
	if opInputService6TestCaseOperation1 == nil {
		opInputService6TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-01-01/vaults/{vaultName}/archives",
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
	Body      []byte  `locationName:"body" type:"blob" json:",omitempty"`
	Checksum  *string `location:"header" locationName:"x-amz-sha256-tree-hash" type:"string" json:"-" xml:"-"`
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" json:"-" xml:"-"`

	metadataInputService6TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService6TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Body" required:"vaultName" json:",omitempty"`
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

	return &InputService7ProtocolTest{service}
}

// InputService7TestCaseOperation1Request generates a request for the InputService7TestCaseOperation1 operation.
func (c *InputService7ProtocolTest) InputService7TestCaseOperation1Request(input *InputService7TestShapeInputShape) (req *aws.Request) {
	if opInputService7TestCaseOperation1 == nil {
		opInputService7TestCaseOperation1 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path",
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

// InputService7TestCaseOperation2Request generates a request for the InputService7TestCaseOperation2 operation.
func (c *InputService7ProtocolTest) InputService7TestCaseOperation2Request(input *InputService7TestShapeInputShape) (req *aws.Request) {
	if opInputService7TestCaseOperation2 == nil {
		opInputService7TestCaseOperation2 = &aws.Operation{
			Name:       "OperationName",
			HTTPMethod: "POST",
			HTTPPath:   "/path?abc=mno",
		}
	}

	req = aws.NewRequest(c.Service, opInputService7TestCaseOperation2, input, nil)

	return
}

func (c *InputService7ProtocolTest) InputService7TestCaseOperation2(input *InputService7TestShapeInputShape) (err error) {
	req := c.InputService7TestCaseOperation2Request(input)
	err = req.Send()
	return
}

var opInputService7TestCaseOperation2 *aws.Operation

type InputService7TestShapeInputShape struct {
	Foo *string `location:"querystring" locationName:"param-name" type:"string" json:"-" xml:"-"`

	metadataInputService7TestShapeInputShape `json:"-", xml:"-"`
}

type metadataInputService7TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

//
// Tests begin here
//

func TestInputService1ProtocolTestURIParameterOnlyWithNoLocationNameCase1(t *testing.T) {
	svc := NewInputService1ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService1TestShapeInputShape{
		PipelineId: aws.String("foo"),
	}
	req := svc.InputService1TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo", r.URL.String())

	// assert headers

}

func TestInputService2ProtocolTestURIParameterOnlyWithLocationNameCase1(t *testing.T) {
	svc := NewInputService2ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService2TestShapeInputShape{
		Foo: aws.String("bar"),
	}
	req := svc.InputService2TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/bar", r.URL.String())

	// assert headers

}

func TestInputService3ProtocolTestURIParameterAndQuerystringParamsCase1(t *testing.T) {
	svc := NewInputService3ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService3TestShapeInputShape{
		Ascending:  aws.String("true"),
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req := svc.InputService3TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers

}

func TestInputService4ProtocolTestURIParameterQuerystringParamsAndJSONBodyCase1(t *testing.T) {
	svc := NewInputService4ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService4TestShapeInputShape{
		Ascending: aws.String("true"),
		Config: &InputService4TestShapeStructType{
			A: aws.String("one"),
			B: aws.String("two"),
		},
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req := svc.InputService4TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("{\"Config\":{\"A\":\"one\",\"B\":\"two\"}}"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers

}

func TestInputService5ProtocolTestURIParameterQuerystringParamsHeadersAndJSONBodyCase1(t *testing.T) {
	svc := NewInputService5ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService5TestShapeInputShape{
		Ascending: aws.String("true"),
		Checksum:  aws.String("12345"),
		Config: &InputService5TestShapeStructType{
			A: aws.String("one"),
			B: aws.String("two"),
		},
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req := svc.InputService5TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("{\"Config\":{\"A\":\"one\",\"B\":\"two\"}}"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers
	assert.Equal(t, "12345", r.Header.Get("x-amz-checksum"))

}

func TestInputService6ProtocolTestStreamingPayloadCase1(t *testing.T) {
	svc := NewInputService6ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService6TestShapeInputShape{
		Body:      []byte("contents"),
		Checksum:  aws.String("foo"),
		VaultName: aws.String("name"),
	}
	req := svc.InputService6TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim("contents"), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/vaults/name/archives", r.URL.String())

	// assert headers
	assert.Equal(t, "foo", r.Header.Get("x-amz-sha256-tree-hash"))

}

func TestInputService7ProtocolTestOmitsNullQueryParamsButSerializesEmptyStringsCase1(t *testing.T) {
	svc := NewInputService7ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService7TestShapeInputShape{}
	req := svc.InputService7TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

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
	req := svc.InputService7TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(""), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path?abc=mno&param-name=", r.URL.String())

	// assert headers

}

