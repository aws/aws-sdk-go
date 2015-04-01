package sqs

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/query"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// SQS is a client for Amazon SQS.
type SQS struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new SQS client.
func New(config *aws.Config) *SQS {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config),
		ServiceName: "sqs",
		APIVersion:  "2012-11-05",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &SQS{service}
}

// newRequest creates a new request for a SQS operation and runs any
// custom request initialization.
func (c *SQS) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
