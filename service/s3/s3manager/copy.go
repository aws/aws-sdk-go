package s3manager

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
)

// DefaultMultipartCopyThreshold is the default object size threshold (in bytes)
// for using multipart instead of simple copy. The default is 10 MB.
const DefaultMultipartCopyThreshold = 10 * 1024 * 1024

// DefaultCopyConcurrency is the default number of goroutines to spin up when
// using Copy().
const DefaultCopyConcurrency = 10

// Copier is a structure for calling Copy(). It is safe to call Copy() on this
// structure for multiple objects and across concurrent goroutines. Mutating
// the Copier's properties is not safe to be done concurrently.
type Copier struct {
	// MaxPartSize is the maximum multipart chunk size to use (in bytes). It
	// must be at least 5 MB. The actual size of the chunks will vary, but
	// remain capped at MaxPartSize. If set to 0, the value of MaxUploadPartSize
	// is used.
	MaxPartSize int64

	// MultipartCopyThreshold is the minimum size (in bytes) an object should
	// have before multipart copy is used instead of simple copy. The minimum
	// is 5 MB, and the maximum is 5 GB. If set to 0, the value of
	// DefaultMultipartCopyThreshold will be used.
	MultipartCopyThreshold int64

	// Concurrency is the number of goroutines to use when executing a multipart
	// copy. If this value is set to 0, the value of DefaultCopyConcurrency
	// will be used.
	//
	// The concurrency pool is not shared between calls to Copy.
	Concurrency int

	// Setting this value to true will cause the SDK to avoid calling
	// AbortMultipartUpload on a failure, leaving all successfully copied
	// parts on S3 for manual recovery.
	//
	// Note that storing parts of an incomplete multipart upload counts towards
	// space usage on S3 and will add additional costs if not cleaned up.
	LeavePartsOnError bool

	// S3 is the client used for executing the multipart copy.
	S3 s3iface.S3API

	// RequestOptions as a list of request options that will be passed down to
	// individual API operation requests made by the uploader.
	RequestOptions []request.Option
}

// WithCopierRequestOptions appends to the Copier's API request options.
func WithCopierRequestOptions(opts ...request.Option) func(*Copier) {
	return func(u *Copier) {
		u.RequestOptions = append(u.RequestOptions, opts...)
	}
}

// NewCopier creates a new Copier instance for copying objects between
// buckets and/or keys in S3. Customisations can be passed in options, or
// otherwise as part of the Copier's Copy method. The ConfigProvider (e.g.
// a *session.Session) will be used to instantiate a S3 service client.
//
// See NewCopierWithClient for examples.
func NewCopier(c client.ConfigProvider, options ...func(*Copier)) *Copier {
	return NewCopierWithClient(s3.New(c), options...)
}

// NewCopierWithClient creates a new Copier instance for copying objects
// between buckets and/or keys in S3. Customisations can be passed in options,
// or otherwise as part of the Copier's Copy method.
func NewCopierWithClient(svc s3iface.S3API, options ...func(*Copier)) *Copier {
	c := &Copier{
		S3:                     svc,
		MaxPartSize:            MaxUploadPartSize,
		MultipartCopyThreshold: DefaultMultipartCopyThreshold,
		Concurrency:            DefaultCopyConcurrency,
		LeavePartsOnError:      false,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// Copy copies an object between buckets and/or keys in S3. See CopyWithContext
// for more information.
func (c Copier) Copy(input *s3.CopyObjectInput, options ...func(*Copier)) (*s3.CopyObjectOutput, error) {
	return c.CopyWithContext(aws.BackgroundContext(), input, options...)
}

// CopyWithContext copies an object between buckets and/or keys in S3. When the
// size of a source object is large enough, it will use concurrent multipart
// uploads to execute on the copy of an arbitrarily-sized object.
//
// The Copier will use either the CopyObject or UploadPartCopy APIs, and no
// actual object data will pass through it. If you wish to transform the data,
// use a combination of s3manager.Downloader and s3manager.Uploader.
//
// The Copier's AWS credentials must have s3:GetObject permission on the source
// object, and s3:PutObject permission on the destination. It is advisable that
// it also have s3:AbortMultipartUpload on the destination, as otherwise a failed
// copy would leave parts abandoned. To disable aborting in case of a failed
// copy, set LeavePartsOnError to true.
//
// Additional functional options can be provided to configure the individual
// copies. These options are copies of the Copier instance Upload is called from.
// Modifying the options will not impact the original Copier instance.
//
// Use the WithCopierRequestOptions helper function to pass in request
// options that will be applied to all API operations made with this uploader.
//
// It is safe to call this method concurrently across goroutines.
func (c Copier) CopyWithContext(ctx aws.Context, input *s3.CopyObjectInput, options ...func(*Copier)) (*s3.CopyObjectOutput, error) {
	it := copier{in: input, cfg: c, ctx: ctx}

	for _, opt := range options {
		opt(&it.cfg)
	}
	it.cfg.RequestOptions = append(it.cfg.RequestOptions, request.WithAppendUserAgent("S3Manager"))

	return it.copy()
}

// optimalPartSize returns the optimal multipart copy part size. It minimizes the
// number of round trips to UploadPartCopy by maximizing the number of bytes
// copied per goroutine-part.
func optimalPartSize(sourceSize int64, concurrency int) int64 {
	return awsutil.MaxInt64(
		MinUploadPartSize, awsutil.MinInt64(
			sourceSize/int64(concurrency), MaxUploadPartSize))
}

// copySourceRange produces a value appropriate for the CopySourceRange header
// of a UploadPartCopy call. The result is "bytes=start-end", where start is
// the first, and end is the last byte to be copied. The start and end are
// inclusive byte indices into the source object.
func copySourceRange(sourceSize, partSize, partNum int64) string {
	rangeStart := (partNum - 1) * partSize
	remainingBytes := sourceSize - rangeStart
	rangeEnd := rangeStart + awsutil.MinInt64(remainingBytes, partSize) - 1
	return fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd)
}

type copier struct {
	ctx aws.Context // context from the call to CopyWithContext()
	cfg Copier      // copy of the Copier that created this copier
	in  *s3.CopyObjectInput
	src struct {
		bucket  string
		key     string
		version *string // nil if source object is not versioned
		size    int64   // size (in bytes) of the source object
	}
	partSize  int64 // derived from source object size and concurrency
	partCount int64 // derived from source object size and partSize
}

// copy executes a multipart copy based on the copier's input field. If the
// source object is small, a simple S3 copy is executed instead.
func (c *copier) copy() (*s3.CopyObjectOutput, error) {
	if err := c.init(); err != nil {
		return nil, err
	}

	c.partSize = awsutil.MinInt64(optimalPartSize(c.src.size, c.cfg.Concurrency), c.cfg.MaxPartSize)
	c.partCount = (c.src.size + c.partSize - 1) / c.partSize

	if c.src.size <= c.cfg.MultipartCopyThreshold || c.partCount < 2 {
		return c.simpleCopy()
	}

	return c.multipartCopy()
}

// init sets up the copier.
func (c *copier) init() error {
	if c.cfg.MaxPartSize == 0 {
		c.cfg.MaxPartSize = MaxUploadPartSize
	}

	if c.cfg.MaxPartSize < MinUploadPartSize || c.cfg.MaxPartSize > MaxUploadPartSize {
		return awserr.New(
			"InvalidRequest",
			fmt.Sprintf(
				"expected %d <= MaxPartSize <= %d",
				MinUploadPartSize, MaxUploadPartSize),
			nil)
	}

	if c.cfg.MultipartCopyThreshold == 0 {
		c.cfg.MultipartCopyThreshold = DefaultMultipartCopyThreshold
	}

	if c.cfg.MultipartCopyThreshold < MinUploadPartSize || c.cfg.MultipartCopyThreshold > MaxUploadPartSize {
		return awserr.New(
			"InvalidRequest",
			fmt.Sprintf(
				"expected %d <= MultipartCopyThreshold <= %d",
				MinUploadPartSize, MaxUploadPartSize),
			nil)
	}

	if c.cfg.Concurrency == 0 {
		c.cfg.Concurrency = DefaultCopyConcurrency
	}

	if c.cfg.Concurrency <= 0 {
		return awserr.New(
			"InvalidRequest",
			"expected 0 < Concurrency",
			nil)
	}

	err := c.initSource()
	if err != nil {
		return err
	}

	c.src.size, err = c.getSourceSize()
	if err != nil {
		return err
	}

	return nil
}

// initSource derives the bucket, key, and optional versionId from the
// CopySource input field.
func (c *copier) initSource() error {
	if c.in.CopySource == nil {
		return awserr.New("InvalidRequest", "expected non-nil copy source", nil)
	}

	src, err := url.QueryUnescape(*c.in.CopySource)
	if err != nil {
		return awserr.New("InvalidRequest", "invalid copy source", err)
	}

	a := strings.SplitN(src, "/", 2)
	if len(a) != 2 {
		return awserr.New(
			"InvalidRequest",
			"invalid copy source: expected bucket and key",
			nil)
	}

	b := strings.SplitN(a[1], "?versionId=", 2)
	if len(a) > 2 {
		return awserr.New(
			"InvalidRequest",
			"invalid copy source: expected at most one versionId",
			nil)
	}

	if b[0] == "" {
		return awserr.New(
			"InvalidRequest",
			"invalid copy source: expected key",
			nil)
	}

	c.src.bucket = a[0]
	c.src.key = b[0]
	if len(b) == 2 {
		c.src.version = aws.String(b[1])
	}

	return nil
}

// getSourceSize returns the size (in bytes) of the source object.
func (c *copier) getSourceSize() (int64, error) {
	out, err := c.cfg.S3.HeadObjectWithContext(
		c.ctx, &s3.HeadObjectInput{
			Bucket:    &c.src.bucket,
			Key:       &c.src.key,
			VersionId: c.src.version,
		}, c.cfg.RequestOptions...)

	if err != nil {
		return 0, err
	}

	return *out.ContentLength, nil
}

func (c *copier) simpleCopy() (*s3.CopyObjectOutput, error) {
	return c.cfg.S3.CopyObjectWithContext(c.ctx, c.in, c.cfg.RequestOptions...)
}

func (c *copier) multipartCopy() (*s3.CopyObjectOutput, error) {
	upload, err := c.createUpload()
	if err != nil {
		return nil, err
	}

	// set to false once the upload has been completed
	shouldAbortUpload := true
	defer func() {
		if shouldAbortUpload && !c.cfg.LeavePartsOnError {
			_, _ = c.abortUpload(upload)
		}
	}()

	completedParts := make([]*s3.CompletedPart, c.partCount)
	nextPartNum := int64(0)
	wg := sync.WaitGroup{}
	cctx, cancel := canceller(c.ctx)
	defer cancel()
	firstErr := make(chan error, 1)
	var firstPart *s3.UploadPartCopyOutput

	for i := int64(0); i < awsutil.MinInt64(c.partCount, int64(c.cfg.Concurrency)); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-cctx.Done():
					return

				default:
					partNum := atomic.AddInt64(&nextPartNum, 1)
					if partNum > c.partCount {
						return
					}

					in, out, err := c.uploadPartCopy(cctx, upload, partNum)
					if err != nil {
						select {
						case firstErr <- err:
							// we don't have absolute guarantees about this being
							// the first error, but it's close enough.
						default:
						}
						cancel()
						return
					}

					completedParts[partNum-1] = &s3.CompletedPart{
						PartNumber: in.PartNumber,
						ETag:       out.CopyPartResult.ETag,
					}

					if partNum == 1 {
						firstPart = out
					}
				}
			}
		}()
	}

	wg.Wait()
	select {
	case err := <-firstErr:
		return nil, &multiUploadError{
			awsError: awserr.New(
				"MultipartUpload",
				"failed to copy one or more parts",
				err),
			uploadID: *upload.UploadId,
		}
	default:
		select {
		case <-cctx.Done():
			// the cancellation must have been external to multipartCopy(),
			// and was timed such that none of the uploadPartCopy() calls
			// returned a cancellation error. this means we're responsible
			// for producing it (as well as aborting what's remained of
			// the multipart copy).
			return nil, &multiUploadError{
				awsError: awserr.New(
					"MultipartUpload",
					"multipart copy cancelled",
					cctx.Err()),
				uploadID: *upload.UploadId,
			}
		default:
		}
	}

	completed, err := c.completeUpload(upload, completedParts)
	if err != nil {
		return nil, &multiUploadError{
			awsError: awserr.New(
				"MultipartUpload",
				"failed to complete multipart upload",
				err),
			uploadID: *upload.UploadId,
		}
	}
	shouldAbortUpload = false

	out := s3.CopyObjectOutput{}
	awsutil.Copy(&out, firstPart)
	awsutil.Copy(&out, completed)
	out.CopyObjectResult = &s3.CopyObjectResult{
		ETag:         completed.ETag,
		LastModified: nil, // could use a part, but it wouldn't be exact
	}

	return &out, nil
}

func (c *copier) createUpload() (*s3.CreateMultipartUploadOutput, error) {
	in := s3.CreateMultipartUploadInput{}
	awsutil.Copy(&in, c.in)
	return c.cfg.S3.CreateMultipartUploadWithContext(c.ctx, &in, c.cfg.RequestOptions...)
}

func (c *copier) abortUpload(upload *s3.CreateMultipartUploadOutput) (*s3.AbortMultipartUploadOutput, error) {
	in := s3.AbortMultipartUploadInput{}
	awsutil.Copy(&in, c.in)
	in.UploadId = upload.UploadId
	return c.cfg.S3.AbortMultipartUploadWithContext(c.ctx, &in, c.cfg.RequestOptions...)
}

func (c *copier) uploadPartCopy(
	ctx aws.Context,
	upload *s3.CreateMultipartUploadOutput,
	partNum int64,
) (*s3.UploadPartCopyInput, *s3.UploadPartCopyOutput, error) {
	in := s3.UploadPartCopyInput{}
	awsutil.Copy(&in, c.in)

	// disable these options, as the CopySourceIn* headers from c.in would
	// pertain to the full object and not the individual parts.
	in.CopySourceIfNoneMatch = nil
	in.CopySourceIfMatch = nil
	in.CopySourceIfModifiedSince = nil
	in.CopySourceIfUnmodifiedSince = nil

	in.UploadId = upload.UploadId
	in.PartNumber = &partNum
	in.CopySourceRange = aws.String(copySourceRange(c.src.size, c.partSize, partNum))

	out, err := c.cfg.S3.UploadPartCopyWithContext(ctx, &in, c.cfg.RequestOptions...)
	if err != nil {
		return nil, nil, err
	}

	return &in, out, nil
}

func (c *copier) completeUpload(
	upload *s3.CreateMultipartUploadOutput,
	parts []*s3.CompletedPart,
) (*s3.CompleteMultipartUploadOutput, error) {
	in := s3.CompleteMultipartUploadInput{}
	awsutil.Copy(&in, c.in)
	in.UploadId = upload.UploadId
	completed := s3.CompletedMultipartUpload{Parts: parts}
	in.MultipartUpload = &completed
	return c.cfg.S3.CompleteMultipartUploadWithContext(c.ctx, &in, c.cfg.RequestOptions...)
}
