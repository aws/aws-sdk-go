//Package cloudhsm provides gucumber integration tests support.
package cloudhsm

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@cloudhsm", func() {
		World["client"] = cloudhsm.New(nil)
	})
}
