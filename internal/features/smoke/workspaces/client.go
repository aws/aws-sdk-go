//Package workspaces provides gucumber integration tests support.
package workspaces

import (
	"github.com/aws/aws-sdk-go/internal/features/shared"
	"github.com/aws/aws-sdk-go/service/workspaces"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@workspaces", func() {
		World["client"] = workspaces.New(nil)
	})
}
