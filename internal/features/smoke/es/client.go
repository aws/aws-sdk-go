//Package es provides gucumber integration tests suppport.
package es

import (
	"github.com/aws/aws-sdk-go/internal/features/shared"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@es", func() {
		World["client"] = elasticsearchservice.New(nil)
	})
}
