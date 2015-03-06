package sns

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
)

// SNS is a client for Amazon SNS.
type SNS struct {
	*aws.Service
}

type SNSConfig struct {
	*aws.Config
}

// New returns a new SNS client.
func New(config *SNSConfig) *SNS {
	if config == nil {
		config = &SNSConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "sns",
		APIVersion:  "2010-03-31",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &SNS{service}
}