package dynamodb

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@dynamodb", func() {
		World["client"] = dynamodb.New(nil)
	})
}
