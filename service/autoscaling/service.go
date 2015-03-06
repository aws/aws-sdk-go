package autoscaling

import (
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
	"github.com/awslabs/aws-sdk-go/aws"
)

// AutoScaling is a client for Auto Scaling.
type AutoScaling struct {
	*aws.Service
}

type AutoScalingConfig struct {
	*aws.Config
}

// New returns a new AutoScaling client.
func New(config *AutoScalingConfig) *AutoScaling {
	if config == nil {
		config = &AutoScalingConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "autoscaling",
		APIVersion:  "2011-01-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &AutoScaling{service}
}