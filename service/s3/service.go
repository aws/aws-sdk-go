package s3

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/restxml"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// S3 is a client for Amazon S3.
type S3 struct {
	*aws.Service
}

// New returns a new S3 client.
func New(config *aws.Config) *S3 {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config),
		ServiceName: "s3",
		APIVersion:  "2006-03-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)

	// S3 uses a custom error parser
	service.Handlers.UnmarshalError.PushBack(unmarshalError)

	return &S3{service}
}
