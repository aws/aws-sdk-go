//Package support provides gucumber integration tests support.
package support

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/support"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@support", func() {
		World["client"] = support.New(nil)
	})
}
