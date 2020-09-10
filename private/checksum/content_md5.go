package checksum

import (
	"crypto/md5"
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/internal/awslog"
)

const contentMD5Header = "Content-Md5"

// AddBodyContentMD5Handler computes and sets the HTTP Content-MD5 header for requests that
// require it.
func AddBodyContentMD5Handler(r *request.Request) {
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
		awslog.Errorf(r.Context(), &r.Config,
			"Unable to compute Content-MD5 for unseekable body, S3.%s",
			r.Operation.Name)
		return
	}

	h := md5.New()

	if _, err := aws.CopySeekableBody(h, r.Body); err != nil {
		r.Error = awserr.New("ContentMD5", "failed to compute body MD5", err)
		return
	}

	// encode the md5 checksum in base64 and set the request header.
	v := base64.StdEncoding.EncodeToString(h.Sum(nil))
	r.HTTPRequest.Header.Set(contentMD5Header, v)
}
