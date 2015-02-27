package storagegateway

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
)

// StorageGateway is a client for AWS Storage Gateway.
type StorageGateway struct {
	*aws.Service
}

type StorageGatewayConfig struct {
	*aws.Config
}

// New returns a new StorageGateway client.
func New(config *StorageGatewayConfig) *StorageGateway {
	if config == nil {
		config = &StorageGatewayConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
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

	return &StorageGateway{service}
}