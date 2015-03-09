package emr

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// EMR is a client for Amazon EMR.
type EMR struct {
	*aws.Service
}

type EMRConfig struct {
	*aws.Config
}

// New returns a new EMR client.
func New(config *EMRConfig) *EMR {
	if config == nil {
		config = &EMRConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "elasticmapreduce",
		APIVersion:   "2009-03-31",
		JSONVersion:  "1.1",
		TargetPrefix: "ElasticMapReduce",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &EMR{service}
}
