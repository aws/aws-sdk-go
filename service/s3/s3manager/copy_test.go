package s3manager_test

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func ExampleNewCopierWithClient() {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	copier := s3manager.NewCopierWithClient(
		svc, func(copier *s3manager.Copier) {
			copier.LeavePartsOnError = true
			copier.DiscoverSourceBucketRegion = false
		})
	_, _ = copier.CopyWithContext(
		aws.BackgroundContext(), &s3manager.CopyInput{
			Bucket:     aws.String("dest-bucket"),
			Key:        aws.String("lorem/ipsum.txt"),
			CopySource: aws.String(url.QueryEscape("src-bucket/lorem/ipsum.txt?versionId=1")),
		})
}

func ExampleCopier_CopyWithContext() {
	var svc s3iface.S3API
	copier := s3manager.NewCopierWithClient(svc)

	// Copy s3://src-bucket/lorem/ipsum.txt to s3://dest-bucket/lorem/ipsum.txt.
	// Version 1 of the source object will be copied.
	out, err := copier.Copy(
		&s3manager.CopyInput{
			Bucket:     aws.String("dest-bucket"),
			Key:        aws.String("lorem/ipsum.txt"),
			CopySource: aws.String(url.QueryEscape("src-bucket/lorem/ipsum.txt?versionId=1")),
		},
		// Optional parameter for customization.
		func(c *s3manager.Copier) {
			c.LeavePartsOnError = true
		})

	if err != nil {
		panic(err)
	}

	log.Printf("The destination object's ETag is: %s", *out.ETag)
}

type copyTestCall struct {
	method string
	input  interface{}
}

type copyTestMock struct {
	s3iface.S3API
	calls            []copyTestCall
	srcContentLength int64
	mu               sync.Mutex
	partRanges       []string // form is "num:bytes=start-end"
	errorAtNthRange  int64
}

func (m *copyTestMock) appendCall(method string, input interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = append(m.calls, copyTestCall{method, input})
}

func (m *copyTestMock) appendPartRange(num int64, partRange string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.partRanges = append(m.partRanges, fmt.Sprintf("%d:%s", num, partRange))
}

func (m *copyTestMock) getCallOrder() []string {
	var out []string
	for _, call := range m.calls {
		out = append(out, call.method)
	}
	return out
}

func (m *copyTestMock) HeadObjectWithContext(
	ctx aws.Context,
	in *s3.HeadObjectInput,
	opts ...request.Option,
) (*s3.HeadObjectOutput, error) {
	m.appendCall("HeadObject", in)
	out := s3.HeadObjectOutput{
		ContentLength: aws.Int64(m.srcContentLength),
	}
	return &out, nil
}

func (m *copyTestMock) CopyObjectWithContext(
	ctx aws.Context,
	in *s3.CopyObjectInput,
	opts ...request.Option,
) (*s3.CopyObjectOutput, error) {
	m.appendCall("CopyObject", in)
	out := s3.CopyObjectOutput{}
	out.VersionId = aws.String("VersionId-simple")
	out.CopyObjectResult = &s3.CopyObjectResult{ETag: aws.String("ETag-simple")}
	return &out, nil
}

func (m *copyTestMock) CreateMultipartUploadWithContext(
	ctx aws.Context,
	in *s3.CreateMultipartUploadInput,
	opts ...request.Option,
) (*s3.CreateMultipartUploadOutput, error) {
	m.appendCall("CreateMultipartUpload", in)
	out := s3.CreateMultipartUploadOutput{}
	out.SetUploadId("Upload123")
	return &out, nil
}

func (m *copyTestMock) UploadPartCopyWithContext(
	ctx aws.Context,
	in *s3.UploadPartCopyInput,
	opts ...request.Option,
) (*s3.UploadPartCopyOutput, error) {
	if *in.PartNumber == m.errorAtNthRange {
		return nil, fmt.Errorf("intentional failure")
	}

	m.appendCall("UploadPartCopy", in)
	m.appendPartRange(*in.PartNumber, *in.CopySourceRange)
	out := s3.UploadPartCopyOutput{}
	out.CopyPartResult = &s3.CopyPartResult{
		ETag: aws.String(fmt.Sprintf("ETag-%d", *in.PartNumber)),
	}
	return &out, nil
}

func (m *copyTestMock) CompleteMultipartUploadWithContext(
	ctx aws.Context,
	in *s3.CompleteMultipartUploadInput,
	opts ...request.Option,
) (*s3.CompleteMultipartUploadOutput, error) {
	m.appendCall("CompleteMultipartUpload", in)
	out := s3.CompleteMultipartUploadOutput{}
	return &out, nil
}

func (m *copyTestMock) AbortMultipartUploadWithContext(
	ctx aws.Context,
	in *s3.AbortMultipartUploadInput,
	opts ...request.Option,
) (*s3.AbortMultipartUploadOutput, error) {
	m.appendCall("AbortMultipartUpload", in)
	out := s3.AbortMultipartUploadOutput{}
	return &out, nil
}

func assertEqual(t *testing.T, expect, actual interface{}) {
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(awstesting.SprintExpectActual(expect, actual))
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected no error; got %+v", err)
	}
}

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("expected error")
	}
}

func assertStringIn(t *testing.T, s string, slice []string) {
	for _, x := range slice {
		if x == s {
			return
		}
	}

	t.Errorf("expected to find %s in %+v", s, slice)
}

func TestCopyWhenSizeBelowThreshold(t *testing.T) {
	m := copyTestMock{
		srcContentLength: s3manager.DefaultMultipartCopyThreshold - 1,
	}
	c := s3manager.NewCopierWithClient(
		&m, func(copier *s3manager.Copier) {
			copier.DiscoverSourceBucketRegion = false
		})

	copySource := url.QueryEscape("bucket/prefix/file.txt?versionId=123")
	out, err := c.Copy(&s3manager.CopyInput{
		Bucket:     aws.String("destbucket"),
		Key:        aws.String("dest/key.txt"),
		CopySource: &copySource,
	})
	assertNoError(t, err)

	ord := m.getCallOrder()
	assertEqual(t, []string{"HeadObject", "CopyObject"}, ord)

	{
		in := m.calls[0].input.(*s3.HeadObjectInput)
		assertEqual(t, *in.Bucket, "bucket")
		assertEqual(t, *in.Key, "prefix/file.txt")
		assertEqual(t, *in.VersionId, "123")
	}

	{
		in := m.calls[1].input.(*s3.CopyObjectInput)
		assertEqual(t, *in.Bucket, "destbucket")
		assertEqual(t, *in.Key, "dest/key.txt")
		assertEqual(t, *in.CopySource, copySource)
	}

	assertEqual(t, "VersionId-simple", *out.VersionId)
	assertEqual(t, "ETag-simple", *out.ETag)
}

func TestCopyWhenSizeAboveThreshold(t *testing.T) {
	m := copyTestMock{}
	c := s3manager.NewCopierWithClient(
		&m, func(copier *s3manager.Copier) {
			copier.DiscoverSourceBucketRegion = false
		})
	c.MaxPartSize = s3manager.MinUploadPartSize
	c.MultipartCopyThreshold = s3manager.MinUploadPartSize
	m.srcContentLength = 2*c.MaxPartSize + 1

	copySource := url.QueryEscape("bucket/prefix/file.txt?versionId=123")
	_, err := c.Copy(&s3manager.CopyInput{
		Bucket:     aws.String("destbucket"),
		Key:        aws.String("dest/key.txt"),
		CopySource: &copySource,
	})
	assertNoError(t, err)

	ord := m.getCallOrder()
	assertEqual(t, []string{
		"HeadObject", "CreateMultipartUpload",
		"UploadPartCopy", "UploadPartCopy", "UploadPartCopy",
		"CompleteMultipartUpload"},
		ord)

	assertEqual(t, 3, len(m.partRanges))
	assertStringIn(t, fmt.Sprintf("1:bytes=0-%d", s3manager.MinUploadPartSize-1), m.partRanges)
	assertStringIn(t, fmt.Sprintf("2:bytes=%d-%d", s3manager.MinUploadPartSize, s3manager.MinUploadPartSize*2-1), m.partRanges)
	assertStringIn(t, fmt.Sprintf("3:bytes=%d-%d", 2*s3manager.MinUploadPartSize, 2*s3manager.MinUploadPartSize), m.partRanges)

	{
		in := m.calls[1].input.(*s3.CreateMultipartUploadInput)
		assertEqual(t, *in.Bucket, "destbucket")
		assertEqual(t, *in.Key, "dest/key.txt")
	}

	for _, call := range m.calls[2:5] {
		in := call.input.(*s3.UploadPartCopyInput)
		assertEqual(t, *in.Bucket, "destbucket")
		assertEqual(t, *in.Key, "dest/key.txt")
		assertEqual(t, *in.UploadId, "Upload123")
	}

	{
		in := m.calls[5].input.(*s3.CompleteMultipartUploadInput)
		assertEqual(t, *in.Bucket, "destbucket")
		assertEqual(t, *in.Key, "dest/key.txt")
		assertEqual(t, *in.UploadId, "Upload123")
		assertEqual(t, 3, len(in.MultipartUpload.Parts))
	}
}

func TestCopyAbortWhenUploadPartFails(t *testing.T) {
	m := copyTestMock{}
	c := s3manager.NewCopierWithClient(
		&m, func(copier *s3manager.Copier) {
			copier.DiscoverSourceBucketRegion = false
		})
	c.MaxPartSize = s3manager.MinUploadPartSize
	c.MultipartCopyThreshold = s3manager.MinUploadPartSize
	c.Concurrency = 1
	m.srcContentLength = 2*c.MaxPartSize + 1
	m.errorAtNthRange = 2

	copySource := url.QueryEscape("bucket/prefix/file.txt?versionId=123")
	_, err := c.Copy(&s3manager.CopyInput{
		Bucket:     aws.String("destbucket"),
		Key:        aws.String("dest/key.txt"),
		CopySource: &copySource,
	})
	assertError(t, err)

	ord := m.getCallOrder()
	assertEqual(t, []string{
		"HeadObject", "CreateMultipartUpload",
		"UploadPartCopy",
		"AbortMultipartUpload"},
		ord)

	{
		in := m.calls[3].input.(*s3.AbortMultipartUploadInput)
		assertEqual(t, *in.Bucket, "destbucket")
		assertEqual(t, *in.Key, "dest/key.txt")
		assertEqual(t, *in.UploadId, "Upload123")
	}
}

func TestCopyNoAbortWhenUploadPartFailsButLeavePartsIsSet(t *testing.T) {
	m := copyTestMock{}
	c := s3manager.NewCopierWithClient(
		&m, func(copier *s3manager.Copier) {
			copier.DiscoverSourceBucketRegion = false
		})
	c.MaxPartSize = s3manager.MinUploadPartSize
	c.MultipartCopyThreshold = s3manager.MinUploadPartSize
	c.Concurrency = 1
	c.LeavePartsOnError = true
	m.srcContentLength = 2*c.MaxPartSize + 1
	m.errorAtNthRange = 2

	copySource := url.QueryEscape("bucket/prefix/file.txt?versionId=123")
	_, err := c.Copy(&s3manager.CopyInput{
		Bucket:     aws.String("destbucket"),
		Key:        aws.String("dest/key.txt"),
		CopySource: &copySource,
	})
	assertError(t, err)

	ord := m.getCallOrder()
	assertEqual(t, []string{
		"HeadObject", "CreateMultipartUpload", "UploadPartCopy"},
		ord)
}
