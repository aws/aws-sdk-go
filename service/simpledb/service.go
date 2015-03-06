package simpledb

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
)

// SimpleDB is a client for Amazon SimpleDB.
type SimpleDB struct {
	*aws.Service
}

type SimpleDBConfig struct {
	*aws.Config
}

// New returns a new SimpleDB client.
func New(config *SimpleDBConfig) *SimpleDB {
	if config == nil {
		config = &SimpleDBConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "sdb",
		APIVersion:  "2009-04-15",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &SimpleDB{service}
}