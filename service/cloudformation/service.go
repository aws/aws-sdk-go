package cloudformation

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
)

// CloudFormation is a client for AWS CloudFormation.
type CloudFormation struct {
	*aws.Service
}

type CloudFormationConfig struct {
	*aws.Config
}

// New returns a new CloudFormation client.
func New(config *CloudFormationConfig) *CloudFormation {
	if config == nil {
		config = &CloudFormationConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "cloudformation",
		APIVersion:  "2010-05-15",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)

	return &CloudFormation{service}
}