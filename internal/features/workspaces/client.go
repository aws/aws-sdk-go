package workspaces

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/workspaces"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@workspaces", func() {
		World["client"] = workspaces.New(nil)
	})
}
