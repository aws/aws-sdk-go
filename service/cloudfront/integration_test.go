// +build integration

package cloudfront_test

import (
	"regexp"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/cloudfront"
	"github.com/stretchr/testify/assert"
)

func assertMatches(t *testing.T, regex, expected string) {
	if !regexp.MustCompile(regex).Match([]byte(expected)) {
		t.Errorf("%q\n\tdoes not match /%s/", expected, regex)
	}
}

func TestListDistributions(t *testing.T) {
	client := cloudfront.New(nil)
	resp, err := client.ListDistributions(nil)

	assert.Nil(t, err)
	assert.True(t, *resp.DistributionList.Quantity >= 0)
}

func TestCreateDistribution(t *testing.T) {
	client := cloudfront.New(nil)
	_, serr := client.CreateDistribution(&cloudfront.CreateDistributionInput{
		DistributionConfig: &cloudfront.DistributionConfig{
			CallerReference: aws.String("ID1"),
			Enabled:         aws.Boolean(true),
			Comment:         aws.String("A comment"),
			Origins:         &cloudfront.Origins{Quantity: aws.Long(0)},
			DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{
				ForwardedValues: &cloudfront.ForwardedValues{
					Cookies:     &cloudfront.CookiePreference{Forward: aws.String("cookie")},
					QueryString: aws.Boolean(true),
				},
				TargetOriginID: aws.String("origin"),
				TrustedSigners: &cloudfront.TrustedSigners{
					Enabled:  aws.Boolean(true),
					Quantity: aws.Long(0),
				},
				ViewerProtocolPolicy: aws.String("policy"),
				MinTTL:               aws.Long(0),
			},
		},
	})

	err := aws.Error(serr)
	assert.NotNil(t, err)
	assert.Equal(t, "MalformedXML", err.Code)
	assertMatches(t, "validation errors detected", err.Message)
}
