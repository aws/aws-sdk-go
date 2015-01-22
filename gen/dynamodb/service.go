package dynamodb

import (
	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/protocol/jsonrpc"
	"github.com/stripe/aws-go/aws/signer/v4"
)

// DynamoDB is a client for Amazon DynamoDB.
type DynamoDB struct {
	*aws.Service
}

type DynamoDBConfig struct {
	*aws.Config
}

// New returns a new DynamoDB client.
func New(config *DynamoDBConfig) *DynamoDB {
	if config == nil {
		config = &DynamoDBConfig{}
	}

	service := &aws.Service{
		Config:       aws.MergeConfig(config.Config),
		ServiceName:  "dynamodb",
		JSONVersion:  "1.0",
		TargetPrefix: "DynamoDB_20120810",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &DynamoDB{service}
}
