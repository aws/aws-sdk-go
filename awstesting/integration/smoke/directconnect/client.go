//Package directconnect provides gucumber integration tests support.
package directconnect

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/directconnect"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@directconnect", func() {
		World["client"] = directconnect.New(nil)
	})
}
