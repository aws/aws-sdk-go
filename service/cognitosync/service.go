package cognitosync

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/restjson"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// CognitoSync is a client for Amazon Cognito Sync.
type CognitoSync struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new CognitoSync client.
func New(config *aws.Config) *CognitoSync {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config),
		ServiceName: "cognito-sync",
		APIVersion:  "2014-06-30",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &CognitoSync{service}
}

// newRequest creates a new request for a CognitoSync operation and runs any
// custom request initialization.
func (c *CognitoSync) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
