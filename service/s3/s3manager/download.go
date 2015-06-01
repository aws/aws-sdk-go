package s3manager

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/s3"
)

// The default range of bytes to get at a time when using Download().
var DefaultDownloadPartSize int64 = 1024 * 1024 * 5

// The default number of goroutines to spin up when using Download().
var DefaultDownloadConcurrency = 5

// The default set of options used when opts is nil in Download().
var DefaultDownloadOptions = &DownloadOptions{
	PartSize:    DefaultDownloadPartSize,
	Concurrency: DefaultDownloadConcurrency,
}

// DownloadOptions keeps tracks of extra options to pass to an Download() call.
type DownloadOptions struct {
	// The buffer size (in bytes) to use when buffering data into chunks and
	// sending them as parts to S3. The minimum allowed part size is 5MB, and
	// if this value is set to zero, the DefaultPartSize value will be used.
	PartSize int64

	// The number of goroutines to spin up in parallel when sending parts.
	// If this is set to zero, the DefaultConcurrency value will be used.
	Concurrency int

	// An S3 client to use when performing downloads. Leave this as nil to use
	// a default client.
	S3 *s3.S3
}

func NewDownloader(opts *DownloadOptions) *Downloader {
	if opts == nil {
		opts = DefaultDownloadOptions
	}
	return &Downloader{opts: *opts}
}

type Downloader struct {
	opts DownloadOptions
}

func (d *Downloader) Download(w io.WriterAt, input *s3.GetObjectInput) error {
	impl := downloadimpl{w: w, in: input, opts: &d.opts}
	return impl.download()
}

type downloadimpl struct {
	opts *DownloadOptions
	in   *s3.GetObjectInput
	w    io.WriterAt

	wg sync.WaitGroup
	m  sync.Mutex

	pos        int64
	totalBytes int64
	err        error
}

func (d *downloadimpl) init() {
	d.totalBytes = -1

	if d.opts.Concurrency == 0 {
		d.opts.Concurrency = DefaultDownloadConcurrency
	}

	if d.opts.PartSize == 0 {
		d.opts.PartSize = DefaultDownloadPartSize
	}

	if d.opts.S3 == nil {
		d.opts.S3 = s3.New(nil)
	}
}

func (d *downloadimpl) download() error {
	d.init()

	// Spin up workers
	ch := make(chan dlchunk, d.opts.Concurrency)
	d.wg.Add(d.opts.Concurrency)
	for i := 0; i < cap(ch); i++ {
		go d.downloadPart(ch)
	}

	// Assign work
	for d.geterr() == nil {
		if d.pos != 0 {
			total := d.getTotalBytes()
			if total < 0 {
				time.Sleep(10 * time.Millisecond)
				continue
			} else if d.pos >= total {
				break
			}
		}

		ch <- dlchunk{
			dlchunkcounter: &dlchunkcounter{},
			w:              d.w,
			start:          d.pos,
			size:           d.opts.PartSize,
		}
		d.pos += d.opts.PartSize
	}

	// Wait for completion
	close(ch)
	d.wg.Wait()

	// Return error
	return d.err
}

func (d *downloadimpl) downloadPart(ch chan dlchunk) {
	defer d.wg.Done()

	for {
		chunk, ok := <-ch

		if !ok {
			break
		}

		if d.geterr() == nil {
			in := &s3.GetObjectInput{}
			awsutil.Copy(in, d.in)
			rng := fmt.Sprintf("bytes=%d-%d",
				chunk.start, chunk.start+chunk.size-1)
			in.Range = &rng

			resp, err := d.opts.S3.GetObject(in)
			if err != nil {
				d.seterr(err)
			} else {
				d.setTotalBytes(resp)

				_, err := io.Copy(chunk, resp.Body)
				resp.Body.Close()

				if err != nil {
					d.seterr(err)
				}

			}
		}
	}
}

func (d *downloadimpl) getTotalBytes() int64 {
	d.m.Lock()
	defer d.m.Unlock()

	return d.totalBytes
}

func (d *downloadimpl) setTotalBytes(resp *s3.GetObjectOutput) {
	d.m.Lock()
	defer d.m.Unlock()

	if d.totalBytes >= 0 {
		return
	}

	parts := strings.Split(*resp.ContentRange, "/")
	total, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	if err != nil {
		d.err = err
		return
	}

	d.totalBytes = total
}

// geterr is a thread-safe getter for the error object
func (d *downloadimpl) geterr() error {
	d.m.Lock()
	defer d.m.Unlock()

	return d.err
}

// seterr is a thread-safe setter for the error object
func (d *downloadimpl) seterr(e error) {
	d.m.Lock()
	defer d.m.Unlock()

	d.err = e
}

type dlchunk struct {
	*dlchunkcounter
	w     io.WriterAt
	start int64
	size  int64
}

type dlchunkcounter struct {
	cur int64
}

func (c dlchunk) Write(p []byte) (n int, err error) {
	if c.cur >= c.size {
		return 0, io.EOF
	}

	n, err = c.w.WriteAt(p, c.start+c.cur)
	c.cur += int64(n)

	return
}
