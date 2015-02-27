package datapipeline

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
)

// DataPipeline is a client for AWS Data Pipeline.
type DataPipeline struct {
	*aws.Service
}

type DataPipelineConfig struct {
	*aws.Config
}

// New returns a new DataPipeline client.
func New(config *DataPipelineConfig) *DataPipeline {
	if config == nil {
		config = &DataPipelineConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "datapipeline",
		APIVersion:   "2012-10-29",
		JSONVersion:  "1.1",
		TargetPrefix: "DataPipeline",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &DataPipeline{service}
}