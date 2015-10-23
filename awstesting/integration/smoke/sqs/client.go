//Package sqs provides gucumber integration tests support.
package sqs

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/sqs"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@sqs", func() {
		World["client"] = sqs.New(nil)
	})
}
