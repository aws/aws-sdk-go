package elb

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/query"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// ELB is a client for Elastic Load Balancing.
type ELB struct {
	*aws.Service
}

type ELBConfig struct {
	*aws.Config
}

// New returns a new ELB client.
func New(config *ELBConfig) *ELB {
	if config == nil {
		config = &ELBConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "elasticloadbalancing",
		APIVersion:  "2012-06-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &ELB{service}
}
