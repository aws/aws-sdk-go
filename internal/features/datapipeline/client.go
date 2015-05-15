package datapipeline

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/datapipeline"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@datapipeline", func() {
		World["client"] = datapipeline.New(nil)
	})
}
