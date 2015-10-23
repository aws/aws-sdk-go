//Package directoryservice provides gucumber integration tests support.
package directoryservice

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@directoryservice", func() {
		World["client"] = directoryservice.New(nil)
	})
}
