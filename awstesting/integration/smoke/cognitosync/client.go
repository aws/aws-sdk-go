//Package cognitosync provides gucumber integration tests support.
package cognitosync

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/cognitosync"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@cognitosync", func() {
		World["client"] = cognitosync.New(nil)
	})
}
