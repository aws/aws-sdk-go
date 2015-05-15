package route53

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/route53"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@route53", func() {
		World["client"] = route53.New(nil)
	})
}
