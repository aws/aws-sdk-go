//Package swf provides gucumber integration tests support.
package swf

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/swf"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@swf", func() {
		World["client"] = swf.New(nil)
	})
}
