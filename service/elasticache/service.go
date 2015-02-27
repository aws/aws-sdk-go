package elasticache

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
)

// ElastiCache is a client for Amazon ElastiCache.
type ElastiCache struct {
	*aws.Service
}

type ElastiCacheConfig struct {
	*aws.Config
}

// New returns a new ElastiCache client.
func New(config *ElastiCacheConfig) *ElastiCache {
	if config == nil {
		config = &ElastiCacheConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "elasticache",
		APIVersion:  "2015-02-02",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)

	return &ElastiCache{service}
}