//Package glacier provides gucumber integration tests support.
package glacier

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/glacier"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@glacier", func() {
		World["client"] = glacier.New(nil)
	})
}
