package util

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/internal/sdkio"
)

const contentMD5Header = "Content-Md5"

// ContentMD5 computes and sets the HTTP Content-MD5 header for requests that
// require it.
func ContentMD5(r *request.Request) {
	// if Content-MD5 header is already present, return
	if v := r.HTTPRequest.Header.Get(contentMD5Header); len(v) != 0 {
		return
	}

	// if S3DisableContentMD5Validation flag is set, return
	if aws.BoolValue(r.Config.S3DisableContentMD5Validation) {
		return
	}

	// if request is presigned, return
	if r.IsPresigned() {
		return
	}

	// if body is not seekable, return
	if !aws.IsReaderSeekable(r.Body) {
		if r.Config.Logger != nil {
			r.Config.Logger.Log(fmt.Sprintf(
				"Unable to compute Content-MD5 for unseekable body, S3.%s",
				r.Operation.Name))
		}
		return
	}

	h := md5.New()

	if _, err := CopySeekableBody(h, r.Body); err != nil {
		r.Error = awserr.New("ContentMD5", "failed to compute body MD5", err)
		return
	}

	// encode the md5 checksum in base64 and set the request header.
	v := base64.StdEncoding.EncodeToString(h.Sum(nil))
	r.HTTPRequest.Header.Set(contentMD5Header, v)
}

// CopySeekableBody copies the seekable body to an io.Writer
func CopySeekableBody(dst io.Writer, src io.ReadSeeker) (int64, error) {
	curPos, err := src.Seek(0, sdkio.SeekCurrent)
	if err != nil {
		return 0, err
	}

	// hash the body.  seek back to the first position after reading to reset
	// the body for transmission.  copy errors may be assumed to be from the
	// body.
	n, err := io.Copy(dst, src)
	if err != nil {
		return n, err
	}

	_, err = src.Seek(curPos, sdkio.SeekStart)
	if err != nil {
		return n, err
	}

	return n, nil
}
