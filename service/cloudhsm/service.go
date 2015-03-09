package cloudhsm

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
)

// CloudHSM is a client for CloudHSM.
type CloudHSM struct {
	*aws.Service
}

type CloudHSMConfig struct {
	*aws.Config
}

// New returns a new CloudHSM client.
func New(config *CloudHSMConfig) *CloudHSM {
	if config == nil {
		config = &CloudHSMConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "cloudhsm",
		APIVersion:   "2014-05-30",
		JSONVersion:  "1.1",
		TargetPrefix: "CloudHsmFrontendService",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &CloudHSM{service}
}
