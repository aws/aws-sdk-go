package cloudhsm

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/cloudhsm"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudhsm", func() {
		World["client"] = cloudhsm.New(nil)
	})
}
