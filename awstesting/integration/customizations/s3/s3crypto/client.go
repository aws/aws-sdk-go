//Package s3crypto provides gucumber integration tests support.
package s3crypto

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3crypto"

	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@s3crypto", func() {
		encryptionClient := s3crypto.NewEncryptionClient(nil, nil, func(c *s3crypto.EncryptionClient) {
			c.Config.S3Session = session.New((&aws.Config{
				Region: aws.String("us-west-2"),
			}).WithLogLevel(aws.LogDebugWithRequestRetries | aws.LogDebugWithRequestErrors))
		})
		gucumber.World["encryptionClient"] = encryptionClient

		decryptionClient := s3crypto.NewDecryptionClient(session.New(), func(c *s3crypto.DecryptionClient) {
			sess := session.New(&aws.Config{
				Region: aws.String("us-east-1"),
			})
			c.Config.KMSClient = kms.New(sess)
			c.Config.S3Session = session.New(&aws.Config{Region: aws.String("us-west-2")})
		})
		gucumber.World["decryptionClient"] = decryptionClient

		gucumber.World["client"] = s3.New(session.New(&aws.Config{
			Region: aws.String("us-west-2"),
		}))
	})
}
