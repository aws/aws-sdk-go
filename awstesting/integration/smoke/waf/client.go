//Package waf provides gucumber integration tests support.
package waf

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/waf"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@waf", func() {
		World["client"] = waf.New(nil)
	})
}
