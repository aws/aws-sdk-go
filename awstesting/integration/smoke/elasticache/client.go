//Package elasticache provides gucumber integration tests support.
package elasticache

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/elasticache"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@elasticache", func() {
		World["client"] = elasticache.New(nil)
	})
}
