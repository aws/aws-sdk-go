package opsworks

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
)

// OpsWorks is a client for AWS OpsWorks.
type OpsWorks struct {
	*aws.Service
}

type OpsWorksConfig struct {
	*aws.Config
}

// New returns a new OpsWorks client.
func New(config *OpsWorksConfig) *OpsWorks {
	if config == nil {
		config = &OpsWorksConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "opsworks",
		APIVersion:   "2013-02-18",
		JSONVersion:  "1.1",
		TargetPrefix: "OpsWorks_20130218",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &OpsWorks{service}
}