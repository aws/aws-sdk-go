package codedeploy

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/codedeploy"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@codedeploy", func() {
		World["client"] = codedeploy.New(nil)
	})
}
