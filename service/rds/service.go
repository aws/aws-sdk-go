package rds

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/query"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// RDS is a client for Amazon RDS.
type RDS struct {
	*aws.Service
}

type RDSConfig struct {
	*aws.Config
}

// New returns a new RDS client.
func New(config *RDSConfig) *RDS {
	if config == nil {
		config = &RDSConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "rds",
		APIVersion:  "2014-10-31",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &RDS{service}
}
