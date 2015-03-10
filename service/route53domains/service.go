package route53domains

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
)

// Route53Domains is a client for Amazon Route 53 Domains.
type Route53Domains struct {
	*aws.Service
}

type Route53DomainsConfig struct {
	*aws.Config
}

// New returns a new Route53Domains client.
func New(config *Route53DomainsConfig) *Route53Domains {
	if config == nil {
		config = &Route53DomainsConfig{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config.Config),
		ServiceName:  "route53domains",
		APIVersion:   "2014-05-15",
		JSONVersion:  "1.1",
		TargetPrefix: "Route53Domains_v20140515",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &Route53Domains{service}
}