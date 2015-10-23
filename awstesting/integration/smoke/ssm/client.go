//Package ssm provides gucumber integration tests support.
package ssm

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/ssm"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@ssm", func() {
		World["client"] = ssm.New(nil)
	})
}
