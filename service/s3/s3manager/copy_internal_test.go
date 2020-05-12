package s3manager

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
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
		region  string
	}{
		{"a/b/c.txt", "a", "b/c.txt", nil, true, ""},
		{"a/b/c.txt?versionId=foo", "a", "b/c.txt", aws.String("foo"), true, ""},
		{"", "", "", nil, false, ""},
		{"a", "", "", nil, false, ""},
		{"a/", "", "", nil, false, ""},
	}

	for _, test := range tests {
		c := copier{
			in: &CopyInput{
				CopySource: aws.String(url.QueryEscape(test.input)),
			},
		}

		if test.region != "" {
			c.in.SourceRegion = &test.region
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

			if c.src.region != test.region {
				t.Errorf("expected region %v; got %v", test.region, c.src.region)
			}
		}
	}
}

type copierHeadObjectMock struct {
	s3iface.S3API
	r request.Request
}

func (mock *copierHeadObjectMock) HeadObjectWithContext(
	ctx aws.Context,
	input *s3.HeadObjectInput,
	opts ...request.Option,
) (*s3.HeadObjectOutput, error) {
	for _, opt := range opts {
		opt(&mock.r)
	}

	return nil, nil
}

func TestCopierHeadObjectRegion(t *testing.T) {
	t.Run("NoRegion", func(t *testing.T) {
		m := copierHeadObjectMock{}
		c := copier{}
		c.in = &CopyInput{}
		c.cfg.S3 = &m
		c.src.region = ""
		_, _ = c.getHeadObject()

		if m.r.Config.Region != nil {
			t.Errorf("expected request region to remain nil")
		}
	})

	t.Run("SourceRegionGiven", func(t *testing.T) {
		m := copierHeadObjectMock{}
		c := copier{}
		c.in = &CopyInput{}
		c.cfg.S3 = &m
		c.src.region = "x-central-1"
		_, _ = c.getHeadObject()

		switch actual := m.r.Config.Region; {
		case actual == nil:
			t.Errorf("expected request region; got nil")
		case *actual != "x-central-1":
			t.Errorf("expected request region to equal x-central-1; got %v", *actual)
		}
	})
}
