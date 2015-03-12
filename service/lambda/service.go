package lambda

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/restjson"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// Lambda is a client for AWS Lambda.
type Lambda struct {
	*aws.Service
}

type LambdaConfig struct {
	*aws.Config
}

// New returns a new Lambda client.
func New(config *LambdaConfig) *Lambda {
	if config == nil {
		config = &LambdaConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "lambda",
		APIVersion:  "2014-11-11",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &Lambda{service}
}