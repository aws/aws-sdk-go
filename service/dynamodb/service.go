package dynamodb

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
)

// DynamoDB is a client for DynamoDB.
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
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "dynamodb",
		APIVersion:   "2012-08-10",
		JSONVersion:  "1.0",
		TargetPrefix: "DynamoDB_20120810",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &DynamoDB{service}
}
