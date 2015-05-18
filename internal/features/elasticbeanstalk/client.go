package elasticbeanstalk

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/elasticbeanstalk"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@elasticbeanstalk", func() {
		World["client"] = elasticbeanstalk.New(nil)
	})
}
