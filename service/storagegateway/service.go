package storagegateway

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// StorageGateway is a client for AWS Storage Gateway.
type StorageGateway struct {
	*aws.Service
}

// New returns a new StorageGateway client.
func New(config *aws.Config) *StorageGateway {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "storagegateway",
		APIVersion:   "2013-06-30",
		JSONVersion:  "1.1",
		TargetPrefix: "StorageGateway_20130630",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &StorageGateway{service}
}
