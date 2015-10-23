//Package iam provides gucumber integration tests support.
package iam

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/iam"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@iam", func() {
		World["client"] = iam.New(nil)
	})
}
