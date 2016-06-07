//Package s3 provides gucumber integration tests support.
package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@s3", func() {
		World["cryptoClient"] = s3crypto.NewClient(nil, session.New(&aws.Config{
			Region:      aws.String("us-west-2"),
			Credentials: credentials.NewSharedCredentials("", "integration"),
		}))

		World["client"] = s3.New(session.New(&aws.Config{
			Region:      aws.String("us-west-2"),
			Credentials: credentials.NewSharedCredentials("", "integration"),
		}))
	})
}
