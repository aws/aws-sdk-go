package ec2metadata

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

// A Client is an EC2 Metadata service Client.
type Client struct {
	*aws.Service
}

// New creates a new instance of the EC2 Metadata service client.
func New(config *aws.Config) *Client {
	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config),
		ServiceName: "Client",
		Endpoint:    "http://169.254.169.254/latest",
		APIVersion:  "latest",
	}
	service.Initialize()
	service.Handlers.Unmarshal.PushBack(unmarshalHandler)
	service.Handlers.UnmarshalError.PushBack(unmarshalError)
	service.Handlers.Validate.Clear()
	service.Handlers.Validate.PushBack(validateEndpointHandler)

	return &Client{service}
}

type metadataOutput struct {
	Content string
}

func unmarshalHandler(r *aws.Request) {
	defer r.HTTPResponse.Body.Close()
	b, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Error = awserr.New("SerializationError", "unable to unmarshal EC2 metadata respose", err)
	}

	data := r.Data.(*metadataOutput)
	data.Content = string(b)
}

func unmarshalError(r *aws.Request) {
	defer r.HTTPResponse.Body.Close()
	_, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Error = awserr.New("SerializationError", "unable to unmarshal EC2 metadata error respose", err)
	}

	// TODO extract the error...
}

func validateEndpointHandler(r *aws.Request) {
	if r.Service.Endpoint == "" {
		r.Error = aws.ErrMissingEndpoint
	}
}
