package kinesis

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
)

// Kinesis is a client for Kinesis.
type Kinesis struct {
	*aws.Service
}

type KinesisConfig struct {
	*aws.Config
}

// New returns a new Kinesis client.
func New(config *KinesisConfig) *Kinesis {
	if config == nil {
		config = &KinesisConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "kinesis",
		APIVersion:   "2013-12-02",
		JSONVersion:  "1.1",
		TargetPrefix: "Kinesis_20131202",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &Kinesis{service}
}