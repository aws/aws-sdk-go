package s3manager

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func TestBatchDelete(t *testing.T) {
	count := 0
	expected := []struct {
		bucket, key string
	}{
		{
			key:    "1",
			bucket: "bucket1",
		},
		{
			key:    "2",
			bucket: "bucket2",
		},
		{
			key:    "3",
			bucket: "bucket3",
		},
		{
			key:    "4",
			bucket: "bucket4",
		},
	}

	received := []struct {
		bucket, key string
	}{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.URL.String(), "/")
		received = append(received, struct{ bucket, key string }{urlParts[1], urlParts[2]})
		w.WriteHeader(http.StatusNoContent)
		count++
	}))

	sess := session.New(&aws.Config{
		Endpoint:         &server.URL,
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("foo"),
		Credentials:      credentials.NewStaticCredentials("AKID", "SECRET", "SESSION"),
	})

	svc := s3.New(sess)

	batcher := BatchDelete{
		Client: svc,
	}

	objects := []BatchDeleteObject{
		{
			Object: &s3.DeleteObjectInput{
				Key:    aws.String("1"),
				Bucket: aws.String("bucket1"),
			},
		},
		{
			Object: &s3.DeleteObjectInput{
				Key:    aws.String("2"),
				Bucket: aws.String("bucket2"),
			},
		},
		{
			Object: &s3.DeleteObjectInput{
				Key:    aws.String("3"),
				Bucket: aws.String("bucket3"),
			},
		},
		{
			Object: &s3.DeleteObjectInput{
				Key:    aws.String("4"),
				Bucket: aws.String("bucket4"),
			},
		},
	}

	if err := batcher.Delete(aws.BackgroundContext(), &DeleteObjectsIterator{Objects: objects}); err != nil {
		panic(err)
	}

	if count != len(objects) {
		t.Errorf("Expected %d, but received %d", len(objects), count)
	}

	if len(expected) != len(received) {
		t.Errorf("Expected %d, but received %d", len(expected), len(received))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i].key != received[i].key {
			t.Errorf("Expected %q, but received %q", expected[i].key, received[i].key)
		}

		if expected[i].bucket != received[i].bucket {
			t.Errorf("Expected %q, but received %q", expected[i].bucket, received[i].bucket)
		}
	}
}

type mockS3Client struct {
	*s3.S3
	index   int
	objects []*s3.ListObjectsOutput
}

func (client *mockS3Client) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	object := client.objects[client.index]
	client.index++
	return object, nil
}

func TestBatchDeleteList(t *testing.T) {
	count := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		count++
	}))

	sess := session.New(&aws.Config{
		Endpoint:         &server.URL,
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("foo"),
		Credentials:      credentials.NewStaticCredentials("AKID", "SECRET", "SESSION"),
	})

	objects := []*s3.ListObjectsOutput{
		{
			Contents: []*s3.Object{
				{
					Key: aws.String("1"),
				},
			},
			NextMarker:  aws.String("marker"),
			IsTruncated: aws.Bool(true),
		},
		{
			Contents: []*s3.Object{
				{
					Key: aws.String("2"),
				},
			},
			NextMarker:  aws.String("marker"),
			IsTruncated: aws.Bool(true),
		},
		{
			Contents: []*s3.Object{
				{
					Key: aws.String("3"),
				},
			},
			IsTruncated: aws.Bool(false),
		},
	}

	svc := &mockS3Client{
		s3.New(sess),
		0,
		objects,
	}

	batcher := BatchDelete{
		Client: svc,
	}

	input := &s3.ListObjectsInput{
		Bucket: aws.String("bucket"),
	}
	iter := &DeleteListIterator{
		Bucket: input.Bucket,
		Paginator: request.Pagination{
			NewRequest: func() (*request.Request, error) {
				var inCpy *s3.ListObjectsInput
				if input != nil {
					tmp := *input
					inCpy = &tmp
				}
				req, _ := svc.ListObjectsRequest(inCpy)
				req.Handlers.Clear()
				output, _ := svc.ListObjects(inCpy)
				req.Data = output
				return req, nil
			},
		},
	}

	if err := batcher.Delete(aws.BackgroundContext(), iter); err != nil {
		t.Error(err)
	}

	if count != len(objects) {
		t.Errorf("Expected %d, but received %d", len(objects), count)
	}
}

func TestBatchDownload(t *testing.T) {
	count := 0
	expected := []struct {
		bucket, key string
	}{
		{
			key:    "1",
			bucket: "bucket1",
		},
		{
			key:    "2",
			bucket: "bucket2",
		},
		{
			key:    "3",
			bucket: "bucket3",
		},
		{
			key:    "4",
			bucket: "bucket4",
		},
	}

	received := []struct {
		bucket, key string
	}{}

	payload := []string{
		"1",
		"2",
		"3",
		"4",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.URL.String(), "/")
		received = append(received, struct{ bucket, key string }{urlParts[1], urlParts[2]})
		w.Write([]byte(payload[count]))
		count++
	}))

	sess := session.New(&aws.Config{
		Endpoint:         &server.URL,
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("foo"),
		Credentials:      credentials.NewStaticCredentials("AKID", "SECRET", "SESSION"),
	})

	svc := NewDownloader(sess)

	objects := []BatchDownloadObject{
		{
			Object: &s3.GetObjectInput{
				Key:    aws.String("1"),
				Bucket: aws.String("bucket1"),
			},
			Writer: aws.NewWriteAtBuffer(make([]byte, 128)),
		},
		{
			Object: &s3.GetObjectInput{
				Key:    aws.String("2"),
				Bucket: aws.String("bucket2"),
			},
			Writer: aws.NewWriteAtBuffer(make([]byte, 128)),
		},
		{
			Object: &s3.GetObjectInput{
				Key:    aws.String("3"),
				Bucket: aws.String("bucket3"),
			},
			Writer: aws.NewWriteAtBuffer(make([]byte, 128)),
		},
		{
			Object: &s3.GetObjectInput{
				Key:    aws.String("4"),
				Bucket: aws.String("bucket4"),
			},
			Writer: aws.NewWriteAtBuffer(make([]byte, 128)),
		},
	}

	iter := &DownloadObjectsIterator{Objects: objects}
	if err := svc.DownloadWithIterator(aws.BackgroundContext(), iter); err != nil {
		panic(err)
	}

	if count != len(objects) {
		t.Errorf("Expected %d, but received %d", len(objects), count)
	}

	if len(expected) != len(received) {
		t.Errorf("Expected %d, but received %d", len(expected), len(received))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i].key != received[i].key {
			t.Errorf("Expected %q, but received %q", expected[i].key, received[i].key)
		}

		if expected[i].bucket != received[i].bucket {
			t.Errorf("Expected %q, but received %q", expected[i].bucket, received[i].bucket)
		}
	}

	for i, p := range payload {
		b := iter.Objects[i].Writer.(*aws.WriteAtBuffer).Bytes()
		b = bytes.Trim(b, "\x00")

		if string(b) != p {
			t.Errorf("Expected %q, but received %q", p, b)
		}
	}
}

func TestBatchUpload(t *testing.T) {
	count := 0
	expected := []struct {
		bucket, key string
		reqBody     string
	}{
		{
			key:     "1",
			bucket:  "bucket1",
			reqBody: "1",
		},
		{
			key:     "2",
			bucket:  "bucket2",
			reqBody: "2",
		},
		{
			key:     "3",
			bucket:  "bucket3",
			reqBody: "3",
		},
		{
			key:     "4",
			bucket:  "bucket4",
			reqBody: "4",
		},
	}

	received := []struct {
		bucket, key, reqBody string
	}{}

	payload := []string{
		"a",
		"b",
		"c",
		"d",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.URL.String(), "/")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}

		received = append(received, struct{ bucket, key, reqBody string }{urlParts[1], urlParts[2], string(b)})
		w.Write([]byte(payload[count]))

		count++
	}))

	sess := session.New(&aws.Config{
		Endpoint:         &server.URL,
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("foo"),
		Credentials:      credentials.NewStaticCredentials("AKID", "SECRET", "SESSION"),
	})

	svc := NewUploader(sess)

	objects := []BatchUploadObject{
		{
			Object: &UploadInput{
				Key:    aws.String("1"),
				Bucket: aws.String("bucket1"),
				Body:   bytes.NewBuffer([]byte("1")),
			},
		},
		{
			Object: &UploadInput{
				Key:    aws.String("2"),
				Bucket: aws.String("bucket2"),
				Body:   bytes.NewBuffer([]byte("2")),
			},
		},
		{
			Object: &UploadInput{
				Key:    aws.String("3"),
				Bucket: aws.String("bucket3"),
				Body:   bytes.NewBuffer([]byte("3")),
			},
		},
		{
			Object: &UploadInput{
				Key:    aws.String("4"),
				Bucket: aws.String("bucket4"),
				Body:   bytes.NewBuffer([]byte("4")),
			},
		},
	}

	iter := &UploadObjectsIterator{Objects: objects}
	if err := svc.UploadWithIterator(aws.BackgroundContext(), iter); err != nil {
		panic(err)
	}

	if count != len(objects) {
		t.Errorf("Expected %d, but received %d", len(objects), count)
	}

	if len(expected) != len(received) {
		t.Errorf("Expected %d, but received %d", len(expected), len(received))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i].key != received[i].key {
			t.Errorf("Expected %q, but received %q", expected[i].key, received[i].key)
		}

		if expected[i].bucket != received[i].bucket {
			t.Errorf("Expected %q, but received %q", expected[i].bucket, received[i].bucket)
		}

		if expected[i].reqBody != received[i].reqBody {
			t.Errorf("Expected %q, but received %q", expected[i].reqBody, received[i].reqBody)
		}
	}
}

type mockClient struct {
	s3iface.S3API
	index     int
	responses []response
}

type response struct {
	out interface{}
	err error
}

func (client *mockClient) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	resp := client.responses[client.index]
	client.index++
	return resp.out.(*s3.PutObjectOutput), resp.err
}

func (client *mockClient) PutObjectRequest(input *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	resp := client.responses[client.index]
	req, _ := client.S3API.PutObjectRequest(input)
	req.Handlers.Clear()
	req.Data = resp.out
	req.Error = resp.err

	client.index++
	return req, resp.out.(*s3.PutObjectOutput)
}

func (client *mockClient) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	resp := client.responses[client.index]
	client.index++
	return resp.out.(*s3.ListObjectsOutput), resp.err
}

func (client *mockClient) ListObjectsRequest(input *s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	resp := client.responses[client.index]
	req, _ := client.S3API.ListObjectsRequest(input)
	req.Handlers.Clear()
	req.Data = resp.out
	req.Error = resp.err

	client.index++
	return req, resp.out.(*s3.ListObjectsOutput)
}

func TestBatchError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	sess := session.New(&aws.Config{
		Endpoint:         &server.URL,
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("foo"),
		Credentials:      credentials.NewStaticCredentials("AKID", "SECRET", "SESSION"),
	})
	svc := &mockClient{
		s3.New(sess),
		0,
		[]response{
			{
				&s3.PutObjectOutput{},
				errors.New("Foo"),
			},
			{
				&s3.PutObjectOutput{},
				nil,
			},
			{
				&s3.PutObjectOutput{},
				nil,
			},
			{
				&s3.PutObjectOutput{},
				errors.New("Bar"),
			},
		},
	}
	uploader := NewUploaderWithClient(svc)

	objects := []BatchUploadObject{
		{
			Object: &UploadInput{
				Key:    aws.String("1"),
				Bucket: aws.String("bucket1"),
				Body:   bytes.NewBuffer([]byte("1")),
			},
		},
		{
			Object: &UploadInput{
				Key:    aws.String("2"),
				Bucket: aws.String("bucket2"),
				Body:   bytes.NewBuffer([]byte("2")),
			},
		},
		{
			Object: &UploadInput{
				Key:    aws.String("3"),
				Bucket: aws.String("bucket3"),
				Body:   bytes.NewBuffer([]byte("3")),
			},
		},
		{
			Object: &UploadInput{
				Key:    aws.String("4"),
				Bucket: aws.String("bucket4"),
				Body:   bytes.NewBuffer([]byte("4")),
			},
		},
	}

	iter := &UploadObjectsIterator{Objects: objects}
	if err := uploader.UploadWithIterator(aws.BackgroundContext(), iter); err != nil {
		if bErr, ok := err.(*BatchError); !ok {
			t.Error("Expected BatchError, but received other")
		} else {
			if len(bErr.Errors) != 2 {
				t.Errorf("Expected 2 errors, but received %d", len(bErr.Errors))
			}

			expected := []struct {
				bucket, key string
			}{
				{
					"bucket1",
					"1",
				},
				{
					"bucket4",
					"4",
				},
			}
			for i, expect := range expected {
				if *bErr.Errors[i].Bucket != expect.bucket {
					t.Errorf("Case %d: Invalid bucket expected %s, but received %s", i, expect.bucket, *bErr.Errors[i].Bucket)
				}

				if *bErr.Errors[i].Key != expect.key {
					t.Errorf("Case %d: Invalid key expected %s, but received %s", i, expect.key, *bErr.Errors[i].Key)
				}
			}
		}
	} else {
		t.Error("Expected error, but received nil")
	}

	if svc.index != len(objects) {
		t.Errorf("Expected %d, but received %d", len(objects), svc.index)
	}

}
