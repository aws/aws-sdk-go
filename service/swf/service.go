package swf

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// SWF is a client for Amazon SWF.
type SWF struct {
	*aws.Service
}

type SWFConfig struct {
	*aws.Config
}

// New returns a new SWF client.
func New(config *SWFConfig) *SWF {
	if config == nil {
		config = &SWFConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "swf",
		APIVersion:   "2012-01-25",
		JSONVersion:  "1.0",
		TargetPrefix: "SimpleWorkflowService",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &SWF{service}
}
