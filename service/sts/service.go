package sts

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
)

// STS is a client for AWS STS.
type STS struct {
	*aws.Service
}

type STSConfig struct {
	*aws.Config
}

// New returns a new STS client.
func New(config *STSConfig) *STS {
	if config == nil {
		config = &STSConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "sts",
		APIVersion:  "2011-06-15",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)

	return &STS{service}
}