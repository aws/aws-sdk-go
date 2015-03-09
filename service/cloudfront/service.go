package cloudfront

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/restxml"
)

// CloudFront is a client for CloudFront.
type CloudFront struct {
	*aws.Service
}

type CloudFrontConfig struct {
	*aws.Config
}

// New returns a new CloudFront client.
func New(config *CloudFrontConfig) *CloudFront {
	if config == nil {
		config = &CloudFrontConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "cloudfront",
		APIVersion:  "2014-11-06",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &CloudFront{service}
}
