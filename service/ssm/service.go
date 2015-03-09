package ssm

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
)

// SSM is a client for Amazon SSM.
type SSM struct {
	*aws.Service
}

type SSMConfig struct {
	*aws.Config
}

// New returns a new SSM client.
func New(config *SSMConfig) *SSM {
	if config == nil {
		config = &SSMConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "ssm",
		APIVersion:   "2014-11-06",
		JSONVersion:  "1.1",
		TargetPrefix: "AmazonSSM",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &SSM{service}
}
