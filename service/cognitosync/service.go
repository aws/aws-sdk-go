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

type CognitoSyncConfig struct {
	*aws.Config
}

// New returns a new CognitoSync client.
func New(config *CognitoSyncConfig) *CognitoSync {
	if config == nil {
		config = &CognitoSyncConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
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

	return &CognitoSync{service}
}