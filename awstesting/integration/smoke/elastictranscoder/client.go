//Package elastictranscoder provides gucumber integration tests support.
package elastictranscoder

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@elastictranscoder", func() {
		World["client"] = elastictranscoder.New(nil)
	})
}
