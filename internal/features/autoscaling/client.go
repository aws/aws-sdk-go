package autoscaling

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/autoscaling"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@autoscaling", func() {
		World["client"] = autoscaling.New(nil)
	})
}
