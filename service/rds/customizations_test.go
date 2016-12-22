package rds

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"

	"github.com/stretchr/testify/assert"
)

func TestPresignWithPresignNotSet(t *testing.T) {
	reqs := map[string]*request.Request{}
	svc := New(unit.Session, &aws.Config{Region: aws.String("us-west-2")})

	assert.NotPanics(t, func() {
		// Doesn't panic on nil input
		req, _ := svc.CopyDBSnapshotRequest(nil)
		req.Sign()
	})

	req, _ := svc.CopyDBSnapshotRequest(&CopyDBSnapshotInput{
		SourceRegion:               aws.String("us-west-1"),
		SourceDBSnapshotIdentifier: aws.String("foo"),
		TargetDBSnapshotIdentifier: aws.String("bar"),
	})

	reqs[opCopyDBSnapshot] = req

	for op, req := range reqs {
		req.Sign()
		b, _ := ioutil.ReadAll(req.HTTPRequest.Body)
		q, _ := url.ParseQuery(string(b))

		u, _ := url.QueryUnescape(q.Get("PreSignedUrl"))
		assert.Regexp(t, fmt.Sprintf(`^https://rds.us-west-1\.amazonaws\.com/\?Action=%s.+?DestinationRegion=us-west-2.+`, op), u)
	}
}

func TestPresignWithPresignSet(t *testing.T) {
	reqs := map[string]*request.Request{}
	svc := New(unit.Session, &aws.Config{Region: aws.String("us-west-2")})

	assert.NotPanics(t, func() {
		// Doesn't panic on nil input
		req, _ := svc.CopyDBSnapshotRequest(nil)
		req.Sign()
	})

	req, _ := svc.CopyDBSnapshotRequest(&CopyDBSnapshotInput{
		SourceRegion:               aws.String("us-west-1"),
		SourceDBSnapshotIdentifier: aws.String("foo"),
		TargetDBSnapshotIdentifier: aws.String("bar"),
		PreSignedUrl:               aws.String("presignedURL"),
	})

	reqs[opCopyDBSnapshot] = req

	for _, req := range reqs {
		req.Sign()

		b, _ := ioutil.ReadAll(req.HTTPRequest.Body)
		q, _ := url.ParseQuery(string(b))

		u, _ := url.QueryUnescape(q.Get("PreSignedUrl"))
		assert.Regexp(t, `presignedURL`, u)
	}
}
