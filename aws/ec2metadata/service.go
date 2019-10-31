// Package ec2metadata provides the client for making API calls to the
// EC2 Metadata service.
//
// This package's client can be disabled completely by setting the environment
// variable "AWS_EC2_METADATA_DISABLED=true". This environment variable set to
// true instructs the SDK to disable the EC2 Metadata client. The client cannot
// be used while the environment variable is set to true, (case insensitive).
package ec2metadata

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
)

// ServiceName is the name of the service.
const ServiceName = "ec2metadata"
const disableServiceEnvVar = "AWS_EC2_METADATA_DISABLED"

const TokenHeader = "x-aws-ec2-metadata-token"
const TTLHeader = "x-aws-ec2-metadata-token-ttl-seconds"
const defaultTTL = "21600"

// NamedHandler for EC2Metadata client used to fetch token for IMDS
var fetchTokenHandler request.NamedHandler

// A EC2Metadata is an EC2 Metadata service Client.
type EC2Metadata struct {
	*client.Client
	tokenTTL string
}

// A ec2Token struct helps use of token in EC2 Metadata service ops
type ec2Token struct {
	Token string
	credentials.Expiry
}

// token points to an ec2Token
var token *ec2Token

// New creates a new instance of the EC2Metadata client with a session.
// This client is safe to use across multiple goroutines.
//
//
// Example:
//     // Create a EC2Metadata client from just a session.
//     svc := ec2metadata.New(mySession)
//
//     // Create a EC2Metadata client with additional configuration
//     svc := ec2metadata.New(mySession, aws.NewConfig().WithLogLevel(aws.LogDebugHTTPBody))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *EC2Metadata {
	c := p.ClientConfig(ServiceName, cfgs...)
	return NewClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// NewClient returns a new EC2Metadata client. Should be used to create
// a client when not using a session. Generally using just New with a session
// is preferred.
//
// If an unmodified HTTP client is provided from the stdlib default, or no client
// the EC2RoleProvider's EC2Metadata HTTP client's timeout will be shortened.
// To disable this set Config.EC2MetadataDisableTimeoutOverride to false. Enabled by default.
func NewClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string, opts ...func(*client.Client)) *EC2Metadata {
	if !aws.BoolValue(cfg.EC2MetadataDisableTimeoutOverride) && httpClientZero(cfg.HTTPClient) {
		// If the http client is unmodified and this feature is not disabled
		// set custom timeouts for EC2Metadata requests.
		cfg.HTTPClient = &http.Client{
			// use a shorter timeout than default because the metadata
			// service is local if it is running, and to fail faster
			// if not running on an ec2 instance.
			Timeout: 5 * time.Second,
		}
	}

	svc := &EC2Metadata{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName: ServiceName,
				ServiceID:   ServiceName,
				Endpoint:    endpoint,
				APIVersion:  "latest",
			},
			handlers,
		),
	}

	svc.tokenTTL = defaultTTL

	token = &ec2Token{}
	fetchTokenHandler = request.NamedHandler{
		Name: "FetchTokenHandler",
		Fn:   svc.fetchTokenHandler,
	}

	svc.Handlers.Build.PushBackNamed(fetchTokenHandler)
	svc.Handlers.Unmarshal.PushBackNamed(unmarshalHandler)
	svc.Handlers.UnmarshalError.PushBack(unmarshalError)
	svc.Handlers.Validate.Clear()
	svc.Handlers.Validate.PushBack(validateEndpointHandler)

	// Disable the EC2 Metadata service if the environment variable is set.
	// This short-circuits the service's functionality to always fail to send
	// requests.
	if strings.ToLower(os.Getenv(disableServiceEnvVar)) == "true" {
		svc.Handlers.Send.SwapNamed(request.NamedHandler{
			Name: corehandlers.SendHandler.Name,
			Fn: func(r *request.Request) {
				r.HTTPResponse = &http.Response{
					Header: http.Header{},
				}
				r.Error = awserr.New(
					request.CanceledErrorCode,
					"EC2 IMDS access disabled via "+disableServiceEnvVar+" env var",
					nil)
			},
		})
	}

	// Add additional options to the service config
	for _, option := range opts {
		option(svc.Client)
	}
	return svc
}

// SetTokenTTL exposes tokenTTL config on client
func (c *EC2Metadata) SetTokenTTL (TTL string){
	c.tokenTTL = TTL
}

func httpClientZero(c *http.Client) bool {
	return c == nil || (c.Transport == nil && c.CheckRedirect == nil && c.Jar == nil && c.Timeout == 0)
}

type metadataOutput struct {
	Content string
}

type tokenOutput struct {
	Token string
	TTL   time.Duration
}

// unmarshal handler for token data
var unmarshalTokenHandler = request.NamedHandler{
	Name: "unmarshalTokenHandler",
	Fn: func(r *request.Request) {
		var b bytes.Buffer
		if _, err := io.Copy(&b, r.HTTPResponse.Body); err != nil {
			r.Error = awserr.New(request.ErrCodeSerialization, "unable to unmarshal EC2 metadata response", err)
			return
		}

		v := r.HTTPResponse.Header.Get(TTLHeader)
		if data, ok := r.Data.(*tokenOutput); ok {
			data.Token = b.String()
			// TTL is in seconds
			i, err := strconv.Atoi(v)
			if err != nil {
				r.Error = awserr.New(request.ParamFormatErrCode, "unable to parse EC2 token TTL response", err)
				return
			}
			t := time.Duration(i) * time.Second
			data.TTL = t
		}
	},
}

var unmarshalHandler = request.NamedHandler{
	Name: "unmarshalMetadataHandler",
	Fn: func(r *request.Request) {
		defer r.HTTPResponse.Body.Close()
		var b bytes.Buffer
		if _, err := io.Copy(&b, r.HTTPResponse.Body); err != nil {
			r.Error = awserr.New(request.ErrCodeSerialization, "unable to unmarshal EC2 metadata response", err)
			return
		}

		if data, ok := r.Data.(*metadataOutput); ok {
			data.Content = b.String()
		}
	},
}

func unmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	var b bytes.Buffer
	if _, err := io.Copy(&b, r.HTTPResponse.Body); err != nil {
		e:= awserr.New(request.ErrCodeSerialization, "unable to unmarshal EC2 metadata error response", err)
		r.Error = awserr.NewRequestFailure(e, r.HTTPResponse.StatusCode, r.RequestID)
		return
	}

	// Response body format is not consistent between metadata endpoints.
	// Grab the error message as a string and include that as the source error
	r.Error = awserr.New("EC2MetadataError", "failed to make EC2Metadata request", errors.New(b.String()))
}

func validateEndpointHandler(r *request.Request) {
	if r.ClientInfo.Endpoint == "" {
		r.Error = aws.ErrMissingEndpoint
	}
}

// fetchTokenHandler fetches token for EC2Metadata service client by default.
func (c *EC2Metadata) fetchTokenHandler(r *request.Request) {

	if token != nil && !token.IsExpired() {
		r.HTTPRequest.Header.Set("x-aws-ec2-metadata-token", token.Token)
		return
	}

	c.Handlers.Build.Remove(fetchTokenHandler)
	defer c.Handlers.Build.PushBackNamed(fetchTokenHandler)

	op := &request.Operation{
		Name:       "GetToken",
		HTTPMethod: "PUT",
		HTTPPath:   "/api/token",
	}

	var output tokenOutput
	req := c.NewRequest(op, nil, &output)

	req.Handlers.Unmarshal.Swap("unmarshalMetadataHandler", unmarshalTokenHandler)
	defer req.Handlers.Unmarshal.Swap("unmarshalTokenHandler", unmarshalHandler)

	req.HTTPRequest.Header.Set(TTLHeader,c.tokenTTL)
	err := req.Send()

	if err != nil {
		// Errors with bad request status should be returned.
		if req.HTTPResponse.StatusCode == http.StatusBadRequest {
			e := awserr.New(req.HTTPResponse.Status, "Fetch token failed", err)
			r.Error = awserr.NewRequestFailure(e, req.HTTPResponse.StatusCode, req.RequestID)
		}
		return
	}

	token.Token = output.Token
	token.SetExpiration(time.Now().Add(output.TTL), 10*time.Second)

	// Inject token header to the request.
	r.HTTPRequest.Header.Set(TokenHeader, token.Token)
}
