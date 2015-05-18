package cognitosync

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/cognitosync"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cognitosync", func() {
		World["client"] = cognitosync.New(nil)
	})
}
