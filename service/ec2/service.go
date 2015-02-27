package ec2

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/ec2query"
)

// EC2 is a client for Amazon EC2.
type EC2 struct {
	*aws.Service
}

type EC2Config struct {
	*aws.Config
}

// New returns a new EC2 client.
func New(config *EC2Config) *EC2 {
	if config == nil {
		config = &EC2Config{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "ec2",
		APIVersion:  "2014-10-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)

	return &EC2{service}
}