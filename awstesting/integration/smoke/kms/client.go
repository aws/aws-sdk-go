//Package kms provides gucumber integration tests support.
package kms

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/kms"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@kms", func() {
		World["client"] = kms.New(nil)
	})
}
