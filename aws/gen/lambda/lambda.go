// Package lambda provides a client for Amazon Lambda.
package lambda

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// Lambda is a client for Amazon Lambda.
type Lambda struct {
	client *aws.RestClient
}

// New returns a new Lambda client.
func New(key, secret, region string, client *http.Client) *Lambda {
	if client == nil {
		client = http.DefaultClient
	}

	service := "lambda"
	endpoint, service, region := endpoints.Lookup("lambda", region)

	return &Lambda{
		client: &aws.RestClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: service,
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2014-11-11",
		},
	}
}

// AddEventSource identifies an Amazon Kinesis stream as the event source
// for an AWS Lambda function. AWS Lambda invokes the specified function
// when records are posted to the stream. This is the pull model, where AWS
// Lambda invokes the function. For more information, go to AWS LambdaL How
// it Works in the AWS Lambda Developer Guide. This association between an
// Amazon Kinesis stream and an AWS Lambda function is called the event
// source mapping. You provide the configuration information (for example,
// which stream to read from and which AWS Lambda function to invoke) for
// the event source mapping in the request body. This operation requires
// permission for the iam:PassRole action for the IAM role. It also
// requires permission for the lambda:AddEventSource action.
func (c *Lambda) AddEventSource(req AddEventSourceRequest) (resp *EventSourceConfiguration, err error) {
	resp = &EventSourceConfiguration{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/event-source-mappings/"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// DeleteFunction deletes the specified Lambda function code and
// configuration. This operation requires permission for the
// lambda:DeleteFunction action.
func (c *Lambda) DeleteFunction(req DeleteFunctionRequest) (err error) {
	// NRE

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}"

	uri = strings.Replace(uri, "{"+"FunctionName"+"}", req.FunctionName, -1)
	uri = strings.Replace(uri, "{"+"FunctionName+"+"}", req.FunctionName, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	return
}

// GetEventSource returns configuration information for the specified event
// source mapping (see AddEventSource This operation requires permission
// for the lambda:GetEventSource action.
func (c *Lambda) GetEventSource(req GetEventSourceRequest) (resp *EventSourceConfiguration, err error) {
	resp = &EventSourceConfiguration{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/event-source-mappings/{UUID}"

	uri = strings.Replace(uri, "{"+"UUID"+"}", req.UUID, -1)
	uri = strings.Replace(uri, "{"+"UUID+"+"}", req.UUID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// GetFunction returns the configuration information of the Lambda function
// and a presigned URL link to the .zip file you uploaded with
// UploadFunction so you can download the .zip file. Note that the URL is
// valid for up to 10 minutes. The configuration information is the same
// information you provided as parameters when uploading the function. This
// operation requires permission for the lambda:GetFunction action.
func (c *Lambda) GetFunction(req GetFunctionRequest) (resp *GetFunctionResponse, err error) {
	resp = &GetFunctionResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}"

	uri = strings.Replace(uri, "{"+"FunctionName"+"}", req.FunctionName, -1)
	uri = strings.Replace(uri, "{"+"FunctionName+"+"}", req.FunctionName, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// GetFunctionConfiguration returns the configuration information of the
// Lambda function. This the same information you provided as parameters
// when uploading the function by using UploadFunction This operation
// requires permission for the lambda:GetFunctionConfiguration operation.
func (c *Lambda) GetFunctionConfiguration(req GetFunctionConfigurationRequest) (resp *FunctionConfiguration, err error) {
	resp = &FunctionConfiguration{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}/configuration"

	uri = strings.Replace(uri, "{"+"FunctionName"+"}", req.FunctionName, -1)
	uri = strings.Replace(uri, "{"+"FunctionName+"+"}", req.FunctionName, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// InvokeAsync submits an invocation request to AWS Lambda. Upon receiving
// the request, Lambda executes the specified function asynchronously. To
// see the logs generated by the Lambda function execution, see the
// CloudWatch logs console. This operation requires permission for the
// lambda:InvokeAsync action.
func (c *Lambda) InvokeAsync(req InvokeAsyncRequest) (resp *InvokeAsyncResponse, err error) {
	resp = &InvokeAsyncResponse{}

	var body io.Reader
	var contentType string

	contentType = "application/json"
	b, err := json.Marshal(req.InvokeArgs)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}/invoke-async/"

	uri = strings.Replace(uri, "{"+"FunctionName"+"}", req.FunctionName, -1)
	uri = strings.Replace(uri, "{"+"FunctionName+"+"}", req.FunctionName, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: add support for extracting output members from statusCode to support Status

	return
}

// ListEventSources returns a list of event source mappings. For each
// mapping, the API returns configuration information (see AddEventSource
// ). You can optionally specify filters to retrieve specific event source
// mappings. This operation requires permission for the
// lambda:ListEventSources action.
func (c *Lambda) ListEventSources(req ListEventSourcesRequest) (resp *ListEventSourcesResponse, err error) {
	resp = &ListEventSourcesResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/event-source-mappings/"

	q := url.Values{}

	if s := req.EventSourceARN; s != "" {

		q.Set("EventSource", s)
	}

	if s := req.FunctionName; s != "" {

		q.Set("FunctionName", s)
	}

	if s := req.Marker; s != "" {

		q.Set("Marker", s)
	}

	if s := strconv.Itoa(req.MaxItems); req.MaxItems != 0 {

		q.Set("MaxItems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// ListFunctions returns a list of your Lambda functions. For each
// function, the response includes the function configuration information.
// You must use GetFunction to retrieve the code for your function. This
// operation requires permission for the lambda:ListFunctions action.
func (c *Lambda) ListFunctions(req ListFunctionsRequest) (resp *ListFunctionsResponse, err error) {
	resp = &ListFunctionsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/"

	q := url.Values{}

	if s := req.Marker; s != "" {

		q.Set("Marker", s)
	}

	if s := strconv.Itoa(req.MaxItems); req.MaxItems != 0 {

		q.Set("MaxItems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// RemoveEventSource removes an event source mapping. This means AWS Lambda
// will no longer invoke the function for events in the associated source.
// This operation requires permission for the lambda:RemoveEventSource
// action.
func (c *Lambda) RemoveEventSource(req RemoveEventSourceRequest) (err error) {
	// NRE

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/event-source-mappings/{UUID}"

	uri = strings.Replace(uri, "{"+"UUID"+"}", req.UUID, -1)
	uri = strings.Replace(uri, "{"+"UUID+"+"}", req.UUID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	return
}

// UpdateFunctionConfiguration updates the configuration parameters for the
// specified Lambda function by using the values provided in the request.
// You provide only the parameters you want to change. This operation must
// only be used on an existing Lambda function and cannot be used to update
// the function's code. This operation requires permission for the
// lambda:UpdateFunctionConfiguration action.
func (c *Lambda) UpdateFunctionConfiguration(req UpdateFunctionConfigurationRequest) (resp *FunctionConfiguration, err error) {
	resp = &FunctionConfiguration{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}/configuration"

	uri = strings.Replace(uri, "{"+"FunctionName"+"}", req.FunctionName, -1)
	uri = strings.Replace(uri, "{"+"FunctionName+"+"}", req.FunctionName, -1)

	q := url.Values{}

	if s := req.Description; s != "" {

		q.Set("Description", s)
	}

	if s := req.Handler; s != "" {

		q.Set("Handler", s)
	}

	if s := strconv.Itoa(req.MemorySize); req.MemorySize != 0 {

		q.Set("MemorySize", s)
	}

	if s := req.Role; s != "" {

		q.Set("Role", s)
	}

	if s := strconv.Itoa(req.Timeout); req.Timeout != 0 {

		q.Set("Timeout", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// UploadFunction creates a new Lambda function or updates an existing
// function. The function metadata is created from the request parameters,
// and the code for the function is provided by a .zip file in the request
// body. If the function name already exists, the existing Lambda function
// is updated with the new code and metadata. This operation requires
// permission for the lambda:UploadFunction action.
func (c *Lambda) UploadFunction(req UploadFunctionRequest) (resp *FunctionConfiguration, err error) {
	resp = &FunctionConfiguration{}

	var body io.Reader
	var contentType string

	contentType = "application/json"
	b, err := json.Marshal(req.FunctionZip)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}"

	uri = strings.Replace(uri, "{"+"FunctionName"+"}", req.FunctionName, -1)
	uri = strings.Replace(uri, "{"+"FunctionName+"+"}", req.FunctionName, -1)

	q := url.Values{}

	if s := req.Description; s != "" {

		q.Set("Description", s)
	}

	if s := req.Handler; s != "" {

		q.Set("Handler", s)
	}

	if s := strconv.Itoa(req.MemorySize); req.MemorySize != 0 {

		q.Set("MemorySize", s)
	}

	if s := req.Mode; s != "" {

		q.Set("Mode", s)
	}

	if s := req.Role; s != "" {

		q.Set("Role", s)
	}

	if s := req.Runtime; s != "" {

		q.Set("Runtime", s)
	}

	if s := strconv.Itoa(req.Timeout); req.Timeout != 0 {

		q.Set("Timeout", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// AddEventSourceRequest is undocumented.
type AddEventSourceRequest struct {
	BatchSize    int               `json:"BatchSize,omitempty"`
	EventSource  string            `json:"EventSource"`
	FunctionName string            `json:"FunctionName"`
	Parameters   map[string]string `json:"Parameters,omitempty"`
	Role         string            `json:"Role"`
}

// DeleteFunctionRequest is undocumented.
type DeleteFunctionRequest struct {
	FunctionName string `json:"FunctionName"`
}

// EventSourceConfiguration is undocumented.
type EventSourceConfiguration struct {
	BatchSize    int               `json:"BatchSize,omitempty"`
	EventSource  string            `json:"EventSource,omitempty"`
	FunctionName string            `json:"FunctionName,omitempty"`
	IsActive     bool              `json:"IsActive,omitempty"`
	LastModified time.Time         `json:"LastModified,omitempty"`
	Parameters   map[string]string `json:"Parameters,omitempty"`
	Role         string            `json:"Role,omitempty"`
	Status       string            `json:"Status,omitempty"`
	UUID         string            `json:"UUID,omitempty"`
}

// FunctionCodeLocation is undocumented.
type FunctionCodeLocation struct {
	Location       string `json:"Location,omitempty"`
	RepositoryType string `json:"RepositoryType,omitempty"`
}

// FunctionConfiguration is undocumented.
type FunctionConfiguration struct {
	CodeSize        int64     `json:"CodeSize,omitempty"`
	ConfigurationID string    `json:"ConfigurationId,omitempty"`
	Description     string    `json:"Description,omitempty"`
	FunctionARN     string    `json:"FunctionARN,omitempty"`
	FunctionName    string    `json:"FunctionName,omitempty"`
	Handler         string    `json:"Handler,omitempty"`
	LastModified    time.Time `json:"LastModified,omitempty"`
	MemorySize      int       `json:"MemorySize,omitempty"`
	Mode            string    `json:"Mode,omitempty"`
	Role            string    `json:"Role,omitempty"`
	Runtime         string    `json:"Runtime,omitempty"`
	Timeout         int       `json:"Timeout,omitempty"`
}

// GetEventSourceRequest is undocumented.
type GetEventSourceRequest struct {
	UUID string `json:"UUID"`
}

// GetFunctionConfigurationRequest is undocumented.
type GetFunctionConfigurationRequest struct {
	FunctionName string `json:"FunctionName"`
}

// GetFunctionRequest is undocumented.
type GetFunctionRequest struct {
	FunctionName string `json:"FunctionName"`
}

// GetFunctionResponse is undocumented.
type GetFunctionResponse struct {
	Code          FunctionCodeLocation  `json:"Code,omitempty"`
	Configuration FunctionConfiguration `json:"Configuration,omitempty"`
}

// InvokeAsyncRequest is undocumented.
type InvokeAsyncRequest struct {
	FunctionName string `json:"FunctionName"`
	InvokeArgs   []byte `json:"InvokeArgs"`
}

// InvokeAsyncResponse is undocumented.
type InvokeAsyncResponse struct {
	Status int `json:"Status,omitempty"`
}

// ListEventSourcesRequest is undocumented.
type ListEventSourcesRequest struct {
	EventSourceARN string `json:"EventSourceArn,omitempty"`
	FunctionName   string `json:"FunctionName,omitempty"`
	Marker         string `json:"Marker,omitempty"`
	MaxItems       int    `json:"MaxItems,omitempty"`
}

// ListEventSourcesResponse is undocumented.
type ListEventSourcesResponse struct {
	EventSources []EventSourceConfiguration `json:"EventSources,omitempty"`
	NextMarker   string                     `json:"NextMarker,omitempty"`
}

// ListFunctionsRequest is undocumented.
type ListFunctionsRequest struct {
	Marker   string `json:"Marker,omitempty"`
	MaxItems int    `json:"MaxItems,omitempty"`
}

// ListFunctionsResponse is undocumented.
type ListFunctionsResponse struct {
	Functions  []FunctionConfiguration `json:"Functions,omitempty"`
	NextMarker string                  `json:"NextMarker,omitempty"`
}

// RemoveEventSourceRequest is undocumented.
type RemoveEventSourceRequest struct {
	UUID string `json:"UUID"`
}

// UpdateFunctionConfigurationRequest is undocumented.
type UpdateFunctionConfigurationRequest struct {
	Description  string `json:"Description,omitempty"`
	FunctionName string `json:"FunctionName"`
	Handler      string `json:"Handler,omitempty"`
	MemorySize   int    `json:"MemorySize,omitempty"`
	Role         string `json:"Role,omitempty"`
	Timeout      int    `json:"Timeout,omitempty"`
}

// UploadFunctionRequest is undocumented.
type UploadFunctionRequest struct {
	Description  string `json:"Description,omitempty"`
	FunctionName string `json:"FunctionName"`
	FunctionZip  []byte `json:"FunctionZip"`
	Handler      string `json:"Handler"`
	MemorySize   int    `json:"MemorySize,omitempty"`
	Mode         string `json:"Mode"`
	Role         string `json:"Role"`
	Runtime      string `json:"Runtime"`
	Timeout      int    `json:"Timeout,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name

var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
var _ json.RawMessage
