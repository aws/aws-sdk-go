package directconnect

import (
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/aws"
)

// DirectConnect is a client for AWS Direct Connect.
type DirectConnect struct {
	*aws.Service
}

type DirectConnectConfig struct {
	*aws.Config
}

// New returns a new DirectConnect client.
func New(config *DirectConnectConfig) *DirectConnect {
	if config == nil {
		config = &DirectConnectConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "directconnect",
		APIVersion:   "2012-10-25",
		JSONVersion:  "1.1",
		TargetPrefix: "OvertureService",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &DirectConnect{service}
}