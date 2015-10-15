//Package simpledb provides gucumber integration tests support.
package simpledb

import (
	"github.com/aws/aws-sdk-go/internal/features/shared"
	"github.com/aws/aws-sdk-go/service/simpledb"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@simpledb", func() {
		World["client"] = simpledb.New(nil)
	})
}
