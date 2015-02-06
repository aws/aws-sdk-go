package dynamodb

import (
	"math"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
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
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "dynamodb",
		APIVersion:   "2012-08-10",
		JSONVersion:  "1.0",
		TargetPrefix: "DynamoDB_20120810",
	}
	service.Initialize()

	service.DefaultMaxRetries = 10
	service.RetryRules = func(r *aws.Request) time.Duration {
		delay := time.Duration(math.Pow(2, float64(r.RetryCount))) * 50
		return delay * time.Millisecond
	}

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &DynamoDB{service}
}
