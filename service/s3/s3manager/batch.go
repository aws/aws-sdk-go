package s3manager

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// BatchError will contain the key and bucket of the object that failed to
// either upload or download.
type BatchError struct {
	Errors  Errors
	code    string
	message string
}

// Errors is a typed alias for a slice of errors to satisfy the error
// interface.
type Errors []Error

func (errs Errors) Error() string {
	buf := bytes.NewBuffer(nil)
	for i, err := range errs {
		buf.WriteString(err.Error())
		if i+1 < len(errs) {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

// Error will contain the original error, bucket, and key of the operation that failed
// during batch operations.
type Error struct {
	OrigErr error
	Bucket  *string
	Key     *string
}

func newError(err error, bucket, key *string) Error {
	return Error{
		err,
		bucket,
		key,
	}
}

func (err *Error) Error() string {
	return fmt.Sprintf("failed to upload %q to %q:\n%s", err.Key, err.Bucket, err.OrigErr.Error())
}

// NewBatchError will return a BatchError that satisfies the awserr.Error interface.
func NewBatchError(code, message string, err []Error) awserr.Error {
	return &BatchError{
		Errors:  err,
		code:    code,
		message: message,
	}
}

// Code will return the code associated with the batch error.
func (err *BatchError) Code() string {
	return err.code
}

// Message will return the message associated with the batch error.
func (err *BatchError) Message() string {
	return err.message
}

func (err *BatchError) Error() string {
	return awserr.SprintError(err.Code(), err.Message(), "", err.Errors)
}

// OrigErr will return the original error. Which, in this case, will always be nil
// for batched operations.
func (err *BatchError) OrigErr() error {
	return err.Errors
}

// BatchDeleteIterator is an interface that uses the scanner pattern to
// iterate through what needs to be deleted.
type BatchDeleteIterator interface {
	Next() bool
	Err() error
	DeleteObject() BatchDeleteObject
}

// DeleteListIterator is an alternative iterator for the BatchDelete client. This will
// iterate through a list of objects and delete the objects.
//
// Example:
//	iter := &s3manager.DeleteListIterator{
//		Client: svc,
//		Input: &s3.ListObjectsInput{
//			Bucket:  aws.String("bucket"),
//			MaxKeys: aws.Int64(5),
//		},
//		Paginator: request.Pagination{
//			NewRequest: func() (*request.Request, error) {
//				var inCpy *ListObjectsInput
//				if input != nil {
//					tmp := *input
//					inCpy = &tmp
//				}
//				req, _ := c.ListObjectsRequest(inCpy)
//				return req, nil
//			},
//		},
//	}
//
//	batcher := s3manager.NewBatchDeleteWithClient(svc)
//	if err := batcher.Delete(aws.BackgroundContext(), iter); err != nil {
//		return err
//	}
type DeleteListIterator struct {
	Bucket    *string
	Paginator request.Pagination
	objects   []*s3.Object
}

// NewDeleteListIterator will return a new DeleteListIterator.
func NewDeleteListIterator(svc s3iface.S3API, input *s3.ListObjectsInput, opts ...func(*DeleteListIterator)) BatchDeleteIterator {
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
				return req, nil
			},
		},
	}

	for _, opt := range opts {
		opt(iter)
	}
	return iter
}

// Next will use the S3API client to iterate through a list of objects.
func (iter *DeleteListIterator) Next() bool {
	if len(iter.objects) > 0 {
		iter.objects = iter.objects[1:]
	}

	if len(iter.objects) == 0 {
		return iter.Paginator.Next()
	}

	return true
}

// Err will return the last known error from Next.
func (iter *DeleteListIterator) Err() error {
	return iter.Paginator.Err()
}

// DeleteObject will return the current object to be deleted.
func (iter *DeleteListIterator) DeleteObject() BatchDeleteObject {
	if len(iter.objects) == 0 {
		p := iter.Paginator.Page().(*s3.ListObjectsOutput)
		iter.objects = p.Contents
	}

	return BatchDeleteObject{
		Object: &s3.DeleteObjectInput{
			Bucket: iter.Bucket,
			Key:    iter.objects[0].Key,
		},
	}
}

// BatchDelete will use the s3 package's service client to perform a batch
// delete.
type BatchDelete struct {
	Client s3iface.S3API
}

// NewBatchDeleteWithClient will return a new delete client that can delete a batched amount of
// objects.
//
// Example:
//	batcher := s3manager.NewBatchDeleteWithClient(client, size)
//
//	objects := []BatchDeleteObject{
//		{
//			Object:	&s3.DeleteObjectInput {
//				Key: aws.String("key"),
//				Bucket: aws.String("bucket"),
//			},
//		},
//	}
//
//	if err := batcher.Delete(&s3manager.DeleteObjectsIterator{
//		Objects: objects,
//	}); err != nil {
//		return err
//	}
func NewBatchDeleteWithClient(client s3iface.S3API, options ...func(*BatchDelete)) *BatchDelete {
	svc := &BatchDelete{
		Client: client,
	}

	for _, opt := range options {
		opt(svc)
	}

	return svc
}

// NewBatchDelete will return a new delete client that can delete a batched amount of
// objects.
//
// Example:
//	batcher := s3manager.NewBatchDelete(sess, size)
//
//	objects := []BatchDeleteObject{
//		{
//			Object:	&s3.DeleteObjectInput {
//				Key: aws.String("key"),
//				Bucket: aws.String("bucket"),
//			},
//		},
//	}
//
//	if err := batcher.Delete(&s3manager.DeleteObjectsIterator{
//		Objects: objects,
//	}); err != nil {
//		return err
//	}
func NewBatchDelete(c client.ConfigProvider, options ...func(*BatchDelete)) *BatchDelete {
	client := s3.New(c)
	return NewBatchDeleteWithClient(client, options...)
}

// BatchDeleteObject is a wrapper object for calling the batch delete operation.
type BatchDeleteObject struct {
	Object *s3.DeleteObjectInput
	// After will run after each iteration during the batch process. This function will
	// be executed whether or not the request was successful.
	After func() error
}

// DeleteObjectsIterator is an interface that uses the scanner pattern to iterate
// through a series of objects to be deleted.
type DeleteObjectsIterator struct {
	Objects []BatchDeleteObject
	index   int
	inc     bool
}

// Next will increment the default iterator's index and and ensure that there
// is another object to iterator to.
func (iter *DeleteObjectsIterator) Next() bool {
	if iter.inc {
		iter.index++
	} else {
		iter.inc = true
	}
	return iter.index < len(iter.Objects)
}

// Err will return an error. Since this is just used to satisfy the BatchDeleteIterator interface
// this will only return nil.
func (iter *DeleteObjectsIterator) Err() error {
	return nil
}

// DeleteObject will return the BatchDeleteObject at the current batched index.
func (iter *DeleteObjectsIterator) DeleteObject() BatchDeleteObject {
	object := iter.Objects[iter.index]
	return object
}

// Delete will use the iterator to queue up objects that need to be deleted.
// Once the batch size is met, this will call the deleteBatch function.
func (d *BatchDelete) Delete(ctx aws.Context, iter BatchDeleteIterator) error {
	var errs []Error
	for iter.Next() {
		object := iter.DeleteObject()
		if _, err := d.Client.DeleteObjectWithContext(ctx, object.Object); err != nil {
			errs = append(errs, newError(err, object.Object.Bucket, object.Object.Key))
		}

		if object.After == nil {
			continue
		}

		if err := object.After(); err != nil {
			errs = append(errs, newError(err, object.Object.Bucket, object.Object.Key))
		}
	}

	if len(errs) > 0 {
		return NewBatchError("BatchedDeleteIncomplete", "some objects have failed to be deleted.", errs)
	}
	return nil
}

// BatchDownloadIterator is an interface that uses the scanner pattern to iterate
// through a series of objects to be downloaded.
type BatchDownloadIterator interface {
	Next() bool
	Err() error
	DownloadObject() BatchDownloadObject
}

// BatchDownloadObject contains all necessary information to run a batch operation once.
type BatchDownloadObject struct {
	Object *s3.GetObjectInput
	Writer io.WriterAt
	// After will run after each iteration during the batch process. This function will
	// be executed whether or not the request was successful.
	After func() error
}

// DownloadObjectsIterator implements the BatchDownloadIterator interface and allows for batched
// download of objects.
type DownloadObjectsIterator struct {
	Objects []BatchDownloadObject
	index   int
	inc     bool
}

// Next will increment the default iterator's index and and ensure that there
// is another object to iterator to.
func (batcher *DownloadObjectsIterator) Next() bool {
	if batcher.inc {
		batcher.index++
	} else {
		batcher.inc = true
	}
	return batcher.index < len(batcher.Objects)
}

// DownloadObject will return the BatchDownloadObject at the current batched index.
func (batcher *DownloadObjectsIterator) DownloadObject() BatchDownloadObject {
	object := batcher.Objects[batcher.index]
	return object
}

// Err will return an error. Since this is just used to satisfy the BatchDeleteIterator interface
// this will only return nil.
func (batcher *DownloadObjectsIterator) Err() error {
	return nil
}

// BatchUploadIterator is an interface that uses the scanner pattern to
// iterate through what needs to be uploaded.
type BatchUploadIterator interface {
	Next() bool
	Err() error
	UploadObject() BatchUploadObject
}

// UploadObjectsIterator implements the BatchUploadIterator interface and allows for batched
// upload of objects.
type UploadObjectsIterator struct {
	Objects []BatchUploadObject
	index   int
	inc     bool
}

// Next will increment the default iterator's index and and ensure that there
// is another object to iterator to.
func (batcher *UploadObjectsIterator) Next() bool {
	if batcher.inc {
		batcher.index++
	} else {
		batcher.inc = true
	}
	return batcher.index < len(batcher.Objects)
}

// Err will return an error. Since this is just used to satisfy the BatchUploadIterator interface
// this will only return nil.
func (batcher *UploadObjectsIterator) Err() error {
	return nil
}

// UploadObject will return the BatchUploadObject at the current batched index.
func (batcher *UploadObjectsIterator) UploadObject() BatchUploadObject {
	object := batcher.Objects[batcher.index]
	return object
}

// BatchUploadObject contains all necessary information to run a batch operation once.
type BatchUploadObject struct {
	Object *UploadInput
	// After will run after each iteration during the batch process. This function will
	// be executed whether or not the request was successful.
	After func() error
}
