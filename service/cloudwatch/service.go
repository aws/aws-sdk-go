package cloudwatch

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/query"
)

// CloudWatch is a client for CloudWatch.
type CloudWatch struct {
	*aws.Service
}

type CloudWatchConfig struct {
	*aws.Config
}

// New returns a new CloudWatch client.
func New(config *CloudWatchConfig) *CloudWatch {
	if config == nil {
		config = &CloudWatchConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "monitoring",
		APIVersion:  "2010-08-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &CloudWatch{service}
}