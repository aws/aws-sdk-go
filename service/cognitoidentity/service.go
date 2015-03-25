package cognitoidentity

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// CognitoIdentity is a client for Amazon Cognito Identity.
type CognitoIdentity struct {
	*aws.Service
}

// New returns a new CognitoIdentity client.
func New(config *aws.Config) *CognitoIdentity {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "cognito-identity",
		APIVersion:   "2014-06-30",
		JSONVersion:  "1.1",
		TargetPrefix: "AWSCognitoIdentityService",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &CognitoIdentity{service}
}
