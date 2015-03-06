package kms

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
)

// KMS is a client for KMS.
type KMS struct {
	*aws.Service
}

type KMSConfig struct {
	*aws.Config
}

// New returns a new KMS client.
func New(config *KMSConfig) *KMS {
	if config == nil {
		config = &KMSConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "kms",
		APIVersion:   "2014-11-01",
		JSONVersion:  "1.1",
		TargetPrefix: "TrentService",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &KMS{service}
}