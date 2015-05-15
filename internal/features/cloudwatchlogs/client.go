package cloudwatchlogs

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/cloudwatchlogs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudwatchlogs", func() {
		World["client"] = cloudwatchlogs.New(nil)
	})
}
