package rds

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/rds"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@rds", func() {
		World["client"] = rds.New(nil)
	})
}
