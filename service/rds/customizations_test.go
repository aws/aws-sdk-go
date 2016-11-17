package rds_test

import (
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/rds"

	"github.com/stretchr/testify/assert"
)

func TestCopyDBClusterSnapshotWithPresignNotSet(t *testing.T) {
	svc := rds.New(unit.Session, &aws.Config{Region: aws.String("us-west-2")})

	assert.NotPanics(t, func() {
		// Doesn't panic on nil input
		req, _ := svc.CopyDBClusterSnapshotRequest(nil)
		req.Sign()
	})

	req, _ := svc.CopyDBClusterSnapshotRequest(&rds.CopyDBClusterSnapshotInput{
		SourceRegion:                      aws.String("us-west-1"),
		SourceDBClusterSnapshotIdentifier: aws.String("foo"),
		TargetDBClusterSnapshotIdentifier: aws.String("bar"),
	})
	req.Sign()

	b, _ := ioutil.ReadAll(req.HTTPRequest.Body)
	q, _ := url.ParseQuery(string(b))

	u, _ := url.QueryUnescape(q.Get("PreSignedUrl"))
	assert.Regexp(t, `^https://rds.us-west-1\.amazonaws\.com/\?Action=CopyDBClusterSnapshot.+?DestinationRegion=us-west-2.+`, u)
}

func TestCopyDBClusterSnapshotWithPresignSet(t *testing.T) {
	svc := rds.New(unit.Session, &aws.Config{Region: aws.String("us-west-2")})

	assert.NotPanics(t, func() {
		// Doesn't panic on nil input
		req, _ := svc.CopyDBClusterSnapshotRequest(nil)
		req.Sign()
	})

	req, _ := svc.CopyDBClusterSnapshotRequest(&rds.CopyDBClusterSnapshotInput{
		SourceRegion:                      aws.String("us-west-1"),
		SourceDBClusterSnapshotIdentifier: aws.String("foo"),
		TargetDBClusterSnapshotIdentifier: aws.String("bar"),
		PreSignedUrl:                      aws.String("presignedURL"),
	})
	req.Sign()

	b, _ := ioutil.ReadAll(req.HTTPRequest.Body)
	q, _ := url.ParseQuery(string(b))

	u, _ := url.QueryUnescape(q.Get("PreSignedUrl"))
	assert.Regexp(t, `presignedURL`, u)
}
