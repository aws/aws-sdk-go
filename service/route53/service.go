package route53

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/restxml"
)

// Route53 is a client for Route 53.
type Route53 struct {
	*aws.Service
}

type Route53Config struct {
	*aws.Config
}

// New returns a new Route53 client.
func New(config *Route53Config) *Route53 {
	if config == nil {
		config = &Route53Config{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "route53",
		APIVersion:  "2013-04-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	return &Route53{service}
}
