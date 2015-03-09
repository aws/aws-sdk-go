package importexport

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/internal/protocol/query"
)

// ImportExport is a client for AWS Import/Export.
type ImportExport struct {
	*aws.Service
}

type ImportExportConfig struct {
	*aws.Config
}

// New returns a new ImportExport client.
func New(config *ImportExportConfig) *ImportExport {
	if config == nil {
		config = &ImportExportConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "importexport",
		APIVersion:  "2010-06-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return &ImportExport{service}
}
