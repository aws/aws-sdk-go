package s3manager

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestCopySourceRange(t *testing.T) {
	tests := []struct {
		expect     string
		partNum    int64
		sourceSize int64
		partSize   int64
	}{
		{"bytes=0-9", 1, 10, 999},
		{"bytes=0-4", 1, 10, 5},
		{"bytes=5-9", 2, 10, 5},
	}

	for _, test := range tests {
		rng := copySourceRange(test.sourceSize, test.partSize, test.partNum)
		if rng != test.expect {
			t.Errorf("expected %v, got %v", test.expect, rng)
		}
	}
}

func TestOptimalPartSize(t *testing.T) {
	tests := []struct {
		expect      int64
		sourceSize  int64
		concurrency int
	}{
		{MaxUploadPartSize, 2 * MaxUploadPartSize, 2},
		{MaxUploadPartSize, 3 * MaxUploadPartSize, 2},
		{MinUploadPartSize, 2 * MinUploadPartSize, 2},
		{MinUploadPartSize + 1, 2*MinUploadPartSize + 2, 2},
		{MinUploadPartSize, 1, 2},
	}

	for _, test := range tests {
		size := optimalPartSize(test.sourceSize, test.concurrency)
		if size != test.expect {
			t.Errorf("expected %v, got %v", test.expect, size)
		}
	}
}

func TestCopierInitSource(t *testing.T) {
	tests := []struct {
		input   string
		bucket  string
		key     string
		version *string
		ok      bool
	}{
		{"a/b/c.txt", "a", "b/c.txt", nil, true},
		{"a/b/c.txt?versionId=foo", "a", "b/c.txt", aws.String("foo"), true},
		{"", "", "", nil, false},
		{"a", "", "", nil, false},
		{"a/", "", "", nil, false},
	}

	for _, test := range tests {
		c := copier{
			in: &s3.CopyObjectInput{CopySource: aws.String(url.QueryEscape(test.input))},
		}

		err := c.initSource()
		if !test.ok {
			if err == nil {
				t.Errorf("expected error; got nil")
			}
		} else {
			if err != nil {
				t.Errorf("expected no error; got %+v", err)
			}

			if c.src.bucket != test.bucket {
				t.Errorf("expected bucket %v; got %v", test.bucket, c.src.bucket)
			}

			if c.src.key != test.key {
				t.Errorf("expected key %v; got %v", test.key, c.src.key)
			}

			if !reflect.DeepEqual(c.src.version, test.version) {
				t.Errorf("expected version %v; got %v", test.version, c.src.version)
			}
		}
	}
}
