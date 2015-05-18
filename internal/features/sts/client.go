package sts

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/sts"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@sts", func() {
		World["client"] = sts.New(nil)
	})
}
