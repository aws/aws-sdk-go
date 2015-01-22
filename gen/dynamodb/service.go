package dynamodb

import (
	"net/http"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/protocol/jsonrpc"
	"github.com/stripe/aws-go/aws/signer/v4"
)

// DynamoDB is a client for Amazon DynamoDB.
type DynamoDB struct {
	*aws.Service
}

// New returns a new DynamoDB client.
func New(creds aws.CredentialsProvider, region string, client *http.Client, manualSend bool) *DynamoDB {
	service := &aws.Service{
		Context: aws.Context{
			Credentials: creds,
			Service:     "dynamodb",
			Region:      region,
		},
		HTTPClient:   client,
		ManualSend:   manualSend,
		JSONVersion:  "1.0",
		TargetPrefix: "DynamoDB_20120810",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)

	return &DynamoDB{service}
}
