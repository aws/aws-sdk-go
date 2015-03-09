package support

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// Support is a client for AWS Support.
type Support struct {
	*aws.Service
}

type SupportConfig struct {
	*aws.Config
}

// New returns a new Support client.
func New(config *SupportConfig) *Support {
	if config == nil {
		config = &SupportConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "support",
		APIVersion:   "2013-04-15",
		JSONVersion:  "1.1",
		TargetPrefix: "AWSSupport_20130415",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &Support{service}
}
