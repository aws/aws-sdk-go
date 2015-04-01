package dynamodb

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// DynamoDB is a client for DynamoDB.
type DynamoDB struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new DynamoDB client.
func New(config *aws.Config) *DynamoDB {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "dynamodb",
		APIVersion:   "2012-08-10",
		JSONVersion:  "1.0",
		TargetPrefix: "DynamoDB_20120810",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &DynamoDB{service}
}

// newRequest creates a new request for a DynamoDB operation and runs any
// custom request initialization.
func (c *DynamoDB) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
