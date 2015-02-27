package cloudtrail

import (
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/aws"
)

// CloudTrail is a client for CloudTrail.
type CloudTrail struct {
	*aws.Service
}

type CloudTrailConfig struct {
	*aws.Config
}

// New returns a new CloudTrail client.
func New(config *CloudTrailConfig) *CloudTrail {
	if config == nil {
		config = &CloudTrailConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
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

	return &CloudTrail{service}
}