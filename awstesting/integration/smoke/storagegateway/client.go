//Package storagegateway provides gucumber integration tests support.
package storagegateway

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@storagegateway", func() {
		World["client"] = storagegateway.New(nil)
	})
}
