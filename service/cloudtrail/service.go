package cloudtrail

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// CloudTrail is a client for CloudTrail.
type CloudTrail struct {
	*aws.Service
}

// New returns a new CloudTrail client.
func New(config *aws.Config) *CloudTrail {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "cloudtrail",
		APIVersion:   "2013-11-01",
		JSONVersion:  "1.1",
		TargetPrefix: "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &CloudTrail{service}
}
