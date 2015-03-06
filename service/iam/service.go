package iam

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
)

// IAM is a client for IAM.
type IAM struct {
	*aws.Service
}

type IAMConfig struct {
	*aws.Config
}

// New returns a new IAM client.
func New(config *IAMConfig) *IAM {
	if config == nil {
		config = &IAMConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "iam",
		APIVersion:  "2010-05-08",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &IAM{service}
}