package elasticbeanstalk

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
)

// ElasticBeanstalk is a client for Elastic Beanstalk.
type ElasticBeanstalk struct {
	*aws.Service
}

type ElasticBeanstalkConfig struct {
	*aws.Config
}

// New returns a new ElasticBeanstalk client.
func New(config *ElasticBeanstalkConfig) *ElasticBeanstalk {
	if config == nil {
		config = &ElasticBeanstalkConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "elasticbeanstalk",
		APIVersion:  "2010-12-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &ElasticBeanstalk{service}
}