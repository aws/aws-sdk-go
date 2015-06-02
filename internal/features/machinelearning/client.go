package machinelearning

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/machinelearning"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@machinelearning", func() {
		World["client"] = machinelearning.New(nil)
	})
}
