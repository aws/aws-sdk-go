package redshift

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/query"
)

// Redshift is a client for Amazon Redshift.
type Redshift struct {
	*aws.Service
}

type RedshiftConfig struct {
	*aws.Config
}

// New returns a new Redshift client.
func New(config *RedshiftConfig) *Redshift {
	if config == nil {
		config = &RedshiftConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "redshift",
		APIVersion:  "2012-12-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &Redshift{service}
}