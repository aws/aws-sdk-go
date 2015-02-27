package ecs

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
)

// ECS is a client for Amazon ECS.
type ECS struct {
	*aws.Service
}

type ECSConfig struct {
	*aws.Config
}

// New returns a new ECS client.
func New(config *ECSConfig) *ECS {
	if config == nil {
		config = &ECSConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "ecs",
		APIVersion:   "2014-11-13",
		JSONVersion:  "1.1",
		TargetPrefix: "AmazonEC2ContainerServiceV20141113",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &ECS{service}
}