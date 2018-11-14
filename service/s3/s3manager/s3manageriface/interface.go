// Package s3manageriface provides an interface for the s3manager package
package s3manageriface

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// DownloaderAPI is the interface type for s3manager.Downloader.
type DownloaderAPI interface {
	Download(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error)
	DownloadWithContext(aws.Context, io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error)
	DownloadWithIterator(ctx aws.Context, iter s3manager.BatchDownloadIterator, opts ...func(*s3manager.Downloader)) error
}

var _ DownloaderAPI = (*s3manager.Downloader)(nil)

// UploaderAPI is the interface type for s3manager.Uploader.
type UploaderAPI interface {
	Upload(*s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
	UploadWithContext(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
	UploadWithIterator(ctx aws.Context, iter s3manager.BatchUploadIterator, opts ...func(*s3manager.Uploader)) error
}

var _ UploaderAPI = (*s3manager.Uploader)(nil)

// BatchDeleteAPI is the interface type for s3manager.BatchDelete.
type BatchDeleteAPI interface {
	Delete(ctx aws.Context, iter s3manager.BatchDeleteIterator) error
}

var _ BatchDeleteAPI = (*s3manager.BatchDelete)(nil)
