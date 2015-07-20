//Package efs provides gucumber integration tests suppport.
package efs

import (
	"github.com/aws/aws-sdk-go/aws/awscfg"
	"github.com/aws/aws-sdk-go/internal/features/shared"
	"github.com/aws/aws-sdk-go/service/efs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@efs", func() {
		// FIXME remove custom region
		World["client"] = efs.New(awscfg.New().WithRegion("us-west-2"))
	})
}
