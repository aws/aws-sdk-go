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

func TestPresignWithPresignNotSet(t *testing.T) {
	svc := rds.New(unit.Session, &aws.Config{Region: aws.String("us-west-2")})

	assert.NotPanics(t, func() {
		// Doesn't panic on nil input
		req, _ := svc.CopyDBSnapshotRequest(nil)
		req.Sign()
	})

	req, _ := svc.CopyDBSnapshotRequest(&rds.CopyDBSnapshotInput{
		SourceRegion:               aws.String("us-west-1"),
		SourceDBSnapshotIdentifier: aws.String("foo"),
		TargetDBSnapshotIdentifier: aws.String("bar"),
	})
	req.Sign()

	b, _ := ioutil.ReadAll(req.HTTPRequest.Body)
	q, _ := url.ParseQuery(string(b))

	u, _ := url.QueryUnescape(q.Get("PreSignedUrl"))
	assert.Regexp(t, `^https://rds.us-west-1\.amazonaws\.com/.+`, u)
}

func TestPresignWithPresignSet(t *testing.T) {
	svc := rds.New(unit.Session, &aws.Config{Region: aws.String("us-west-2")})

	assert.NotPanics(t, func() {
		// Doesn't panic on nil input
		req, _ := svc.CopyDBSnapshotRequest(nil)
		req.Sign()
	})

	req, _ := svc.CopyDBSnapshotRequest(&rds.CopyDBSnapshotInput{
		SourceRegion:               aws.String("us-west-1"),
		SourceDBSnapshotIdentifier: aws.String("foo"),
		TargetDBSnapshotIdentifier: aws.String("bar"),
		PreSignedUrl:               aws.String("presignedURL"),
	})
	req.Sign()

	b, _ := ioutil.ReadAll(req.HTTPRequest.Body)
	q, _ := url.ParseQuery(string(b))

	u, _ := url.QueryUnescape(q.Get("PreSignedUrl"))
	assert.Regexp(t, `presignedURL`, u)
}
