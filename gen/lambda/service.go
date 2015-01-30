package lambda

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/restjson"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
)

// Lambda is a client for Amazon Lambda.
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
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)

	return &Lambda{service}
}
