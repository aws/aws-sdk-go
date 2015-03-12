package elastictranscoder

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/restjson"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// ElasticTranscoder is a client for Amazon Elastic Transcoder.
type ElasticTranscoder struct {
	*aws.Service
}

type ElasticTranscoderConfig struct {
	*aws.Config
}

// New returns a new ElasticTranscoder client.
func New(config *ElasticTranscoderConfig) *ElasticTranscoder {
	if config == nil {
		config = &ElasticTranscoderConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "elastictranscoder",
		APIVersion:  "2012-09-25",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &ElasticTranscoder{service}
}