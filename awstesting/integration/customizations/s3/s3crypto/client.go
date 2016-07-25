//Package s3crypto provides gucumber integration tests support.
package s3crypto

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@s3crypto", func() {
		c := s3crypto.New(nil, func(c *s3crypto.Client) {
			c.Config.S3Session = session.New((&aws.Config{
				Region: aws.String("us-west-2"),
			}).WithLogLevel(aws.LogDebugWithRequestRetries | aws.LogDebugWithRequestErrors))
			c.Config.KMSSession = session.New((&aws.Config{
				Region: aws.String("us-east-1"),
			}).WithLogLevel(aws.LogDebugWithRequestRetries | aws.LogDebugWithRequestErrors))
		})
		World["cryptoClient"] = c

		World["client"] = s3.New(session.New(&aws.Config{
			Region: aws.String("us-west-2"),
		}))
	})
}
