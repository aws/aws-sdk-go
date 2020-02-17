// +build go1.7

package s3manager

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	random "math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/internal/sdkio"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/internal/s3testing"
)

type recordedPartPool struct {
	allocations uint64
	gets        uint64
	outstanding int64
	*partPool
}

func newRecordedPartPool(partSize int64) *recordedPartPool {
	rp := &recordedPartPool{}
	pp := newPartPool(partSize)

	n := pp.New
	pp.New = func() interface{} {
		atomic.AddUint64(&rp.allocations, 1)
		return n()
	}

	rp.partPool = pp

	return rp
}

func (r *recordedPartPool) Get() *[]byte {
	atomic.AddUint64(&r.gets, 1)
	atomic.AddInt64(&r.outstanding, 1)
	return r.partPool.Get()
}

func (r *recordedPartPool) Put(b *[]byte) {
	atomic.AddInt64(&r.outstanding, -1)
	r.partPool.Put(b)
}

func swapByteSlicePool(f func(partSize int64) byteSlicePool) func() {
	orig := newByteSlicePool

	newByteSlicePool = f

	return func() {
		newByteSlicePool = orig
	}
}

type testReader struct {
	br *bytes.Reader
	m  sync.Mutex
}

func (r *testReader) Read(p []byte) (n int, err error) {
	r.m.Lock()
	defer r.m.Unlock()
	return r.br.Read(p)
}

func TestUploadByteSlicePool(t *testing.T) {
	cases := map[string]struct {
		PartSize int64
		FileSize int64
	}{
		"single part": {
			PartSize: sdkio.MebiByte * 5,
			FileSize: sdkio.MebiByte * 5,
		},
		"multi-part": {
			PartSize: sdkio.MebiByte * 5,
			FileSize: sdkio.MebiByte * 10,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			var p *recordedPartPool

			unswap := swapByteSlicePool(func(partSize int64) byteSlicePool {
				p = newRecordedPartPool(partSize)
				return p
			})
			defer unswap()

			sess := unit.Session.Copy()
			svc := s3.New(sess)
			svc.Handlers.Unmarshal.Clear()
			svc.Handlers.UnmarshalMeta.Clear()
			svc.Handlers.UnmarshalError.Clear()
			svc.Handlers.Send.Clear()
			svc.Handlers.Send.PushFront(func(r *request.Request) {
				if r.Body != nil {
					io.Copy(ioutil.Discard, r.Body)
				}

				r.HTTPResponse = &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
				}

				switch data := r.Data.(type) {
				case *s3.CreateMultipartUploadOutput:
					data.UploadId = aws.String("UPLOAD-ID")
				case *s3.UploadPartOutput:
					data.ETag = aws.String(fmt.Sprintf("ETAG%d", random.Int()))
				case *s3.CompleteMultipartUploadOutput:
					data.Location = aws.String("https://location")
					data.VersionId = aws.String("VERSION-ID")
				case *s3.PutObjectOutput:
					data.VersionId = aws.String("VERSION-ID")
				}
			})

			uploader := NewUploaderWithClient(svc, func(u *Uploader) {
				u.PartSize = tt.PartSize
				u.Concurrency = 50
			})

			expected := s3testing.GetTestBytes(int(tt.FileSize))
			_, err := uploader.Upload(&UploadInput{
				Bucket: aws.String("bucket"),
				Key:    aws.String("key"),
				Body:   &testReader{br: bytes.NewReader(expected)},
			})
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}

			if v := atomic.LoadInt64(&p.outstanding); v != 0 {
				t.Fatalf("expected zero outsnatding pool parts, got %d", v)
			}

			gets, allocs := atomic.LoadUint64(&p.gets), atomic.LoadUint64(&p.allocations)

			t.Logf("total gets %v, total allocations %v", gets, allocs)
		})
	}
}

func TestUploadByteSlicePool_Failures(t *testing.T) {
	cases := map[string]struct {
		PartSize   int64
		FileSize   int64
		Operations []string
	}{
		"single part": {
			PartSize: sdkio.MebiByte * 5,
			FileSize: sdkio.MebiByte * 4,
			Operations: []string{
				"PutObject",
			},
		},
		"multi-part": {
			PartSize: sdkio.MebiByte * 5,
			FileSize: sdkio.MebiByte * 10,
			Operations: []string{
				"CreateMultipartUpload",
				"UploadPart",
				"CompleteMultipartUpload",
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			for _, operation := range tt.Operations {
				t.Run(operation, func(t *testing.T) {
					var p *recordedPartPool

					unswap := swapByteSlicePool(func(partSize int64) byteSlicePool {
						p = newRecordedPartPool(partSize)
						return p
					})
					defer unswap()

					sess := unit.Session.Copy()
					svc := s3.New(sess)
					svc.Handlers.Unmarshal.Clear()
					svc.Handlers.UnmarshalMeta.Clear()
					svc.Handlers.UnmarshalError.Clear()
					svc.Handlers.Send.Clear()
					svc.Handlers.Send.PushFront(func(r *request.Request) {
						if r.Body != nil {
							io.Copy(ioutil.Discard, r.Body)
						}

						if r.Operation.Name == operation {
							r.Retryable = aws.Bool(false)
							r.Error = fmt.Errorf("request error")
							r.HTTPResponse = &http.Response{
								StatusCode: 500,
								Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
							}
							return
						}

						r.HTTPResponse = &http.Response{
							StatusCode: 200,
							Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
						}

						switch data := r.Data.(type) {
						case *s3.CreateMultipartUploadOutput:
							data.UploadId = aws.String("UPLOAD-ID")
						case *s3.UploadPartOutput:
							data.ETag = aws.String(fmt.Sprintf("ETAG%d", random.Int()))
						case *s3.CompleteMultipartUploadOutput:
							data.Location = aws.String("https://location")
							data.VersionId = aws.String("VERSION-ID")
						case *s3.PutObjectOutput:
							data.VersionId = aws.String("VERSION-ID")
						}
					})

					uploader := NewUploaderWithClient(svc, func(u *Uploader) {
						u.Concurrency = 1
						u.PartSize = tt.PartSize
					})

					expected := s3testing.GetTestBytes(int(tt.FileSize))
					_, err := uploader.Upload(&UploadInput{
						Bucket: aws.String("bucket"),
						Key:    aws.String("key"),
						Body:   &testReader{br: bytes.NewReader(expected)},
					})
					if err == nil {
						t.Fatalf("expected error but got none")
					}

					if v := atomic.LoadInt64(&p.outstanding); v != 0 {
						t.Fatalf("expected zero outsnatding pool parts, got %d", v)
					}
				})
			}
		})
	}
}
