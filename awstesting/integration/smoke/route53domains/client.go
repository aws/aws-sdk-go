//Package route53domains provides gucumber integration tests support.
package route53domains

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/route53domains"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@route53domains", func() {
		World["client"] = route53domains.New(nil)
	})
}
