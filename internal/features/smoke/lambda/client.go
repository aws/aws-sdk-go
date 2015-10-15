//Package lambda provides gucumber integration tests support.
package lambda

import (
	"github.com/aws/aws-sdk-go/internal/features/shared"
	"github.com/aws/aws-sdk-go/service/lambda"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@lambda", func() {
		World["client"] = lambda.New(nil)
	})
}
