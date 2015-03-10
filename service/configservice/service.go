package configservice

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
)

// ConfigService is a client for AWS Config.
type ConfigService struct {
	*aws.Service
}

type ConfigServiceConfig struct {
	*aws.Config
}

// New returns a new ConfigService client.
func New(config *ConfigServiceConfig) *ConfigService {
	if config == nil {
		config = &ConfigServiceConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "config",
		APIVersion:   "2014-11-12",
		JSONVersion:  "1.1",
		TargetPrefix: "StarlingDoveService",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &ConfigService{service}
}