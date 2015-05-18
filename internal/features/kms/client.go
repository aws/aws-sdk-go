package kms

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/kms"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@kms", func() {
		World["client"] = kms.New(nil)
	})
}
