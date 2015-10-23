//Package sns provides gucumber integration tests support.
package sns

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/sns"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@sns", func() {
		World["client"] = sns.New(nil)
	})
}
