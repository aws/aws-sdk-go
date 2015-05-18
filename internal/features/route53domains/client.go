package route53domains

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/route53domains"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@route53domains", func() {
		World["client"] = route53domains.New(nil)
	})
}
