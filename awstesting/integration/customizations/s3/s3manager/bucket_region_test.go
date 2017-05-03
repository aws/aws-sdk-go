// +build integration

package s3manager

import (
	"testing"

	"github.com/EMCECS/aws-sdk-go/aws"
	"github.com/EMCECS/aws-sdk-go/awstesting/integration"
	"github.com/EMCECS/aws-sdk-go/service/s3/s3manager"
)

func TestGetBucketRegion(t *testing.T) {
	expectRegion := aws.StringValue(integration.Session.Config.Region)

	ctx := aws.BackgroundContext()
	region, err := s3manager.GetBucketRegion(ctx, integration.Session,
		aws.StringValue(bucketName), expectRegion)

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := expectRegion, region; e != a {
		t.Errorf("expect %s bucket region, got %s", e, a)
	}
}
