// Package s3 provides a client for Amazon Simple Storage Service.
package s3

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// S3 is a client for Amazon Simple Storage Service.
type S3 struct {
	client *aws.RestXMLClient
}

// New returns a new S3 client.
func New(key, secret, region string, client *http.Client) *S3 {
	if client == nil {
		client = http.DefaultClient
	}

	return &S3{
		client: &aws.RestXMLClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "s3",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoints.Lookup("s3", region),
			APIVersion: "2006-03-01",
		},
	}
}

// AbortMultipartUpload aborts a multipart upload. To verify that all parts
// have been removed, so you don't get charged for the part storage, you
// should call the List Parts operation and ensure the parts list is empty.
func (c *S3) AbortMultipartUpload(req AbortMultipartUploadRequest) (err error) {
	// NRE

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := req.UploadID; s != "" {

		q.Set("uploadId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// CompleteMultipartUpload completes a multipart upload by assembling
// previously uploaded parts.
func (c *S3) CompleteMultipartUpload(req CompleteMultipartUploadRequest) (resp *CompleteMultipartUploadOutput, err error) {
	resp = &CompleteMultipartUploadOutput{}

	var body io.Reader

	b, err := xml.Marshal(req.MultipartUpload)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := req.UploadID; s != "" {

		q.Set("uploadId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// CopyObject creates a copy of an object that is already stored in Amazon
// S3.
func (c *S3) CopyObject(req CopyObjectRequest) (resp *CopyObjectOutput, err error) {
	resp = &CopyObjectOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ACL; s != "" {

		httpReq.Header.Set("x-amz-acl", s)
	}

	if s := req.CacheControl; s != "" {

		httpReq.Header.Set("Cache-Control", s)
	}

	if s := req.ContentDisposition; s != "" {

		httpReq.Header.Set("Content-Disposition", s)
	}

	if s := req.ContentEncoding; s != "" {

		httpReq.Header.Set("Content-Encoding", s)
	}

	if s := req.ContentLanguage; s != "" {

		httpReq.Header.Set("Content-Language", s)
	}

	if s := req.ContentType; s != "" {

		httpReq.Header.Set("Content-Type", s)
	}

	if s := req.CopySource; s != "" {

		httpReq.Header.Set("x-amz-copy-source", s)
	}

	if s := req.CopySourceIfMatch; s != "" {

		httpReq.Header.Set("x-amz-copy-source-if-match", s)
	}

	if s := req.CopySourceIfModifiedSince.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("x-amz-copy-source-if-modified-since", s)
	}

	if s := req.CopySourceIfNoneMatch; s != "" {

		httpReq.Header.Set("x-amz-copy-source-if-none-match", s)
	}

	if s := req.CopySourceIfUnmodifiedSince.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("x-amz-copy-source-if-unmodified-since", s)
	}

	if s := req.CopySourceSSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-copy-source-server-side-encryption-customer-algorithm", s)
	}

	if s := req.CopySourceSSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-copy-source-server-side-encryption-customer-key", s)
	}

	if s := req.CopySourceSSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-copy-source-server-side-encryption-customer-key-MD5", s)
	}

	if s := req.Expires.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("Expires", s)
	}

	if s := req.GrantFullControl; s != "" {

		httpReq.Header.Set("x-amz-grant-full-control", s)
	}

	if s := req.GrantRead; s != "" {

		httpReq.Header.Set("x-amz-grant-read", s)
	}

	if s := req.GrantReadACP; s != "" {

		httpReq.Header.Set("x-amz-grant-read-acp", s)
	}

	if s := req.GrantWriteACP; s != "" {

		httpReq.Header.Set("x-amz-grant-write-acp", s)
	}

	if s := req.MetadataDirective; s != "" {

		httpReq.Header.Set("x-amz-metadata-directive", s)
	}

	if s := req.SSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-algorithm", s)
	}

	if s := req.SSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key", s)
	}

	if s := req.SSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key-MD5", s)
	}

	if s := req.SSEKMSKeyID; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-aws-kms-key-id", s)
	}

	if s := req.ServerSideEncryption; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption", s)
	}

	if s := req.StorageClass; s != "" {

		httpReq.Header.Set("x-amz-storage-class", s)
	}

	if s := req.WebsiteRedirectLocation; s != "" {

		httpReq.Header.Set("x-amz-website-redirect-location", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// CreateBucket is undocumented.
func (c *S3) CreateBucket(req CreateBucketRequest) (resp *CreateBucketOutput, err error) {
	resp = &CreateBucketOutput{}

	var body io.Reader

	b, err := xml.Marshal(req.CreateBucketConfiguration)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ACL; s != "" {

		httpReq.Header.Set("x-amz-acl", s)
	}

	if s := req.GrantFullControl; s != "" {

		httpReq.Header.Set("x-amz-grant-full-control", s)
	}

	if s := req.GrantRead; s != "" {

		httpReq.Header.Set("x-amz-grant-read", s)
	}

	if s := req.GrantReadACP; s != "" {

		httpReq.Header.Set("x-amz-grant-read-acp", s)
	}

	if s := req.GrantWrite; s != "" {

		httpReq.Header.Set("x-amz-grant-write", s)
	}

	if s := req.GrantWriteACP; s != "" {

		httpReq.Header.Set("x-amz-grant-write-acp", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// CreateMultipartUpload initiates a multipart upload and returns an upload
// Note: After you initiate multipart upload and upload one or more parts,
// you must either complete or abort multipart upload in order to stop
// getting charged for storage of the uploaded parts. Only after you either
// complete or abort multipart upload, Amazon S3 frees up the parts storage
// and stops charging you for the parts storage.
func (c *S3) CreateMultipartUpload(req CreateMultipartUploadRequest) (resp *CreateMultipartUploadOutput, err error) {
	resp = &CreateMultipartUploadOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}?uploads"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if s := req.ACL; s != "" {

		httpReq.Header.Set("x-amz-acl", s)
	}

	if s := req.CacheControl; s != "" {

		httpReq.Header.Set("Cache-Control", s)
	}

	if s := req.ContentDisposition; s != "" {

		httpReq.Header.Set("Content-Disposition", s)
	}

	if s := req.ContentEncoding; s != "" {

		httpReq.Header.Set("Content-Encoding", s)
	}

	if s := req.ContentLanguage; s != "" {

		httpReq.Header.Set("Content-Language", s)
	}

	if s := req.ContentType; s != "" {

		httpReq.Header.Set("Content-Type", s)
	}

	if s := req.Expires.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("Expires", s)
	}

	if s := req.GrantFullControl; s != "" {

		httpReq.Header.Set("x-amz-grant-full-control", s)
	}

	if s := req.GrantRead; s != "" {

		httpReq.Header.Set("x-amz-grant-read", s)
	}

	if s := req.GrantReadACP; s != "" {

		httpReq.Header.Set("x-amz-grant-read-acp", s)
	}

	if s := req.GrantWriteACP; s != "" {

		httpReq.Header.Set("x-amz-grant-write-acp", s)
	}

	if s := req.SSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-algorithm", s)
	}

	if s := req.SSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key", s)
	}

	if s := req.SSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key-MD5", s)
	}

	if s := req.SSEKMSKeyID; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-aws-kms-key-id", s)
	}

	if s := req.ServerSideEncryption; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption", s)
	}

	if s := req.StorageClass; s != "" {

		httpReq.Header.Set("x-amz-storage-class", s)
	}

	if s := req.WebsiteRedirectLocation; s != "" {

		httpReq.Header.Set("x-amz-website-redirect-location", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// DeleteBucket deletes the bucket. All objects (including all object
// versions and Delete Markers) in the bucket must be deleted before the
// bucket itself can be deleted.
func (c *S3) DeleteBucket(req DeleteBucketRequest) (err error) {
	// NRE

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// DeleteBucketCors deletes the cors configuration information set for the
// bucket.
func (c *S3) DeleteBucketCors(req DeleteBucketCorsRequest) (err error) {
	// NRE

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?cors"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// DeleteBucketLifecycle deletes the lifecycle configuration from the
// bucket.
func (c *S3) DeleteBucketLifecycle(req DeleteBucketLifecycleRequest) (err error) {
	// NRE

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?lifecycle"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// DeleteBucketPolicy is undocumented.
func (c *S3) DeleteBucketPolicy(req DeleteBucketPolicyRequest) (err error) {
	// NRE

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?policy"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// DeleteBucketTagging is undocumented.
func (c *S3) DeleteBucketTagging(req DeleteBucketTaggingRequest) (err error) {
	// NRE

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?tagging"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// DeleteBucketWebsite this operation removes the website configuration
// from the bucket.
func (c *S3) DeleteBucketWebsite(req DeleteBucketWebsiteRequest) (err error) {
	// NRE

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?website"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// DeleteObject removes the null version (if there is one) of an object and
// inserts a delete marker, which becomes the latest version of the object.
// If there isn't a null version, Amazon S3 does not remove any objects.
func (c *S3) DeleteObject(req DeleteObjectRequest) (resp *DeleteObjectOutput, err error) {
	resp = &DeleteObjectOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := req.VersionID; s != "" {

		q.Set("versionId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if s := req.MFA; s != "" {

		httpReq.Header.Set("x-amz-mfa", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// DeleteObjects this operation enables you to delete multiple objects from
// a bucket using a single request. You may specify up to 1000 keys.
func (c *S3) DeleteObjects(req DeleteObjectsRequest) (resp *DeleteObjectsOutput, err error) {
	resp = &DeleteObjectsOutput{}

	var body io.Reader

	b, err := xml.Marshal(req.Delete)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?delete"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if s := req.MFA; s != "" {

		httpReq.Header.Set("x-amz-mfa", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketAcl is undocumented.
func (c *S3) GetBucketAcl(req GetBucketAclRequest) (resp *GetBucketAclOutput, err error) {
	resp = &GetBucketAclOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?acl"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketCors is undocumented.
func (c *S3) GetBucketCors(req GetBucketCorsRequest) (resp *GetBucketCorsOutput, err error) {
	resp = &GetBucketCorsOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?cors"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketLifecycle returns the lifecycle configuration information set
// on the bucket.
func (c *S3) GetBucketLifecycle(req GetBucketLifecycleRequest) (resp *GetBucketLifecycleOutput, err error) {
	resp = &GetBucketLifecycleOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?lifecycle"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketLocation is undocumented.
func (c *S3) GetBucketLocation(req GetBucketLocationRequest) (resp *GetBucketLocationOutput, err error) {
	resp = &GetBucketLocationOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?location"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketLogging returns the logging status of a bucket and the
// permissions users have to view and modify that status. To use you must
// be the bucket owner.
func (c *S3) GetBucketLogging(req GetBucketLoggingRequest) (resp *GetBucketLoggingOutput, err error) {
	resp = &GetBucketLoggingOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?logging"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketNotification is undocumented.
func (c *S3) GetBucketNotification(req GetBucketNotificationRequest) (resp *GetBucketNotificationOutput, err error) {
	resp = &GetBucketNotificationOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?notification"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketPolicy is undocumented.
func (c *S3) GetBucketPolicy(req GetBucketPolicyRequest) (resp *GetBucketPolicyOutput, err error) {
	resp = &GetBucketPolicyOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?policy"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketRequestPayment returns the request payment configuration of a
// bucket.
func (c *S3) GetBucketRequestPayment(req GetBucketRequestPaymentRequest) (resp *GetBucketRequestPaymentOutput, err error) {
	resp = &GetBucketRequestPaymentOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?requestPayment"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketTagging is undocumented.
func (c *S3) GetBucketTagging(req GetBucketTaggingRequest) (resp *GetBucketTaggingOutput, err error) {
	resp = &GetBucketTaggingOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?tagging"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketVersioning is undocumented.
func (c *S3) GetBucketVersioning(req GetBucketVersioningRequest) (resp *GetBucketVersioningOutput, err error) {
	resp = &GetBucketVersioningOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?versioning"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetBucketWebsite is undocumented.
func (c *S3) GetBucketWebsite(req GetBucketWebsiteRequest) (resp *GetBucketWebsiteOutput, err error) {
	resp = &GetBucketWebsiteOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?website"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetObject is undocumented.
func (c *S3) GetObject(req GetObjectRequest) (resp *GetObjectOutput, err error) {
	resp = &GetObjectOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := req.ResponseCacheControl; s != "" {

		q.Set("response-cache-control", s)
	}

	if s := req.ResponseContentDisposition; s != "" {

		q.Set("response-content-disposition", s)
	}

	if s := req.ResponseContentEncoding; s != "" {

		q.Set("response-content-encoding", s)
	}

	if s := req.ResponseContentLanguage; s != "" {

		q.Set("response-content-language", s)
	}

	if s := req.ResponseContentType; s != "" {

		q.Set("response-content-type", s)
	}

	if s := req.ResponseExpires.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		q.Set("response-expires", s)
	}

	if s := req.VersionID; s != "" {

		q.Set("versionId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if s := req.IfMatch; s != "" {

		httpReq.Header.Set("If-Match", s)
	}

	if s := req.IfModifiedSince.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("If-Modified-Since", s)
	}

	if s := req.IfNoneMatch; s != "" {

		httpReq.Header.Set("If-None-Match", s)
	}

	if s := req.IfUnmodifiedSince.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("If-Unmodified-Since", s)
	}

	if s := req.Range; s != "" {

		httpReq.Header.Set("Range", s)
	}

	if s := req.SSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-algorithm", s)
	}

	if s := req.SSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key", s)
	}

	if s := req.SSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	resp.Body = respBody

	// TODO: extract the rest of the response

	return
}

// GetObjectAcl is undocumented.
func (c *S3) GetObjectAcl(req GetObjectAclRequest) (resp *GetObjectAclOutput, err error) {
	resp = &GetObjectAclOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}?acl"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := req.VersionID; s != "" {

		q.Set("versionId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetObjectTorrent is undocumented.
func (c *S3) GetObjectTorrent(req GetObjectTorrentRequest) (resp *GetObjectTorrentOutput, err error) {
	resp = &GetObjectTorrentOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}?torrent"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	resp.Body = respBody

	// TODO: extract the rest of the response

	return
}

// HeadBucket this operation is useful to determine if a bucket exists and
// you have permission to access it.
func (c *S3) HeadBucket(req HeadBucketRequest) (err error) {
	// NRE

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("HEAD", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// HeadObject the operation retrieves metadata from an object without
// returning the object itself. This operation is useful if you're only
// interested in an object's metadata. To use you must have access to the
// object.
func (c *S3) HeadObject(req HeadObjectRequest) (resp *HeadObjectOutput, err error) {
	resp = &HeadObjectOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := req.VersionID; s != "" {

		q.Set("versionId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("HEAD", uri, body)
	if err != nil {
		return
	}

	if s := req.IfMatch; s != "" {

		httpReq.Header.Set("If-Match", s)
	}

	if s := req.IfModifiedSince.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("If-Modified-Since", s)
	}

	if s := req.IfNoneMatch; s != "" {

		httpReq.Header.Set("If-None-Match", s)
	}

	if s := req.IfUnmodifiedSince.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("If-Unmodified-Since", s)
	}

	if s := req.Range; s != "" {

		httpReq.Header.Set("Range", s)
	}

	if s := req.SSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-algorithm", s)
	}

	if s := req.SSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key", s)
	}

	if s := req.SSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListBuckets returns a list of all buckets owned by the authenticated
// sender of the request.
func (c *S3) ListBuckets() (resp *ListBucketsOutput, err error) {
	resp = &ListBucketsOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListMultipartUploads this operation lists in-progress multipart uploads.
func (c *S3) ListMultipartUploads(req ListMultipartUploadsRequest) (resp *ListMultipartUploadsOutput, err error) {
	resp = &ListMultipartUploadsOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?uploads"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if s := req.Delimiter; s != "" {

		q.Set("delimiter", s)
	}

	if s := req.EncodingType; s != "" {

		q.Set("encoding-type", s)
	}

	if s := req.KeyMarker; s != "" {

		q.Set("key-marker", s)
	}

	if s := strconv.Itoa(req.MaxUploads); req.MaxUploads != 0 {

		q.Set("max-uploads", s)
	}

	if s := req.Prefix; s != "" {

		q.Set("prefix", s)
	}

	if s := req.UploadIDMarker; s != "" {

		q.Set("upload-id-marker", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListObjectVersions returns metadata about all of the versions of objects
// in a bucket.
func (c *S3) ListObjectVersions(req ListObjectVersionsRequest) (resp *ListObjectVersionsOutput, err error) {
	resp = &ListObjectVersionsOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}?versions"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if s := req.Delimiter; s != "" {

		q.Set("delimiter", s)
	}

	if s := req.EncodingType; s != "" {

		q.Set("encoding-type", s)
	}

	if s := req.KeyMarker; s != "" {

		q.Set("key-marker", s)
	}

	if s := strconv.Itoa(req.MaxKeys); req.MaxKeys != 0 {

		q.Set("max-keys", s)
	}

	if s := req.Prefix; s != "" {

		q.Set("prefix", s)
	}

	if s := req.VersionIDMarker; s != "" {

		q.Set("version-id-marker", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListObjects returns some or all (up to 1000) of the objects in a bucket.
// You can use the request parameters as selection criteria to return a
// subset of the objects in a bucket.
func (c *S3) ListObjects(req ListObjectsRequest) (resp *ListObjectsOutput, err error) {
	resp = &ListObjectsOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if s := req.Delimiter; s != "" {

		q.Set("delimiter", s)
	}

	if s := req.EncodingType; s != "" {

		q.Set("encoding-type", s)
	}

	if s := req.Marker; s != "" {

		q.Set("marker", s)
	}

	if s := strconv.Itoa(req.MaxKeys); req.MaxKeys != 0 {

		q.Set("max-keys", s)
	}

	if s := req.Prefix; s != "" {

		q.Set("prefix", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListParts lists the parts that have been uploaded for a specific
// multipart upload.
func (c *S3) ListParts(req ListPartsRequest) (resp *ListPartsOutput, err error) {
	resp = &ListPartsOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := strconv.Itoa(req.MaxParts); req.MaxParts != 0 {

		q.Set("max-parts", s)
	}

	if s := strconv.Itoa(req.PartNumberMarker); req.PartNumberMarker != 0 {

		q.Set("part-number-marker", s)
	}

	if s := req.UploadID; s != "" {

		q.Set("uploadId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// PutBucketAcl sets the permissions on a bucket using access control lists
func (c *S3) PutBucketAcl(req PutBucketAclRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.AccessControlPolicy)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?acl"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ACL; s != "" {

		httpReq.Header.Set("x-amz-acl", s)
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	if s := req.GrantFullControl; s != "" {

		httpReq.Header.Set("x-amz-grant-full-control", s)
	}

	if s := req.GrantRead; s != "" {

		httpReq.Header.Set("x-amz-grant-read", s)
	}

	if s := req.GrantReadACP; s != "" {

		httpReq.Header.Set("x-amz-grant-read-acp", s)
	}

	if s := req.GrantWrite; s != "" {

		httpReq.Header.Set("x-amz-grant-write", s)
	}

	if s := req.GrantWriteACP; s != "" {

		httpReq.Header.Set("x-amz-grant-write-acp", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketCors is undocumented.
func (c *S3) PutBucketCors(req PutBucketCorsRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.CORSConfiguration)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?cors"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketLifecycle sets lifecycle configuration for your bucket. If a
// lifecycle configuration exists, it replaces it.
func (c *S3) PutBucketLifecycle(req PutBucketLifecycleRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.LifecycleConfiguration)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?lifecycle"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketLogging set the logging parameters for a bucket and to specify
// permissions for who can view and modify the logging parameters. To set
// the logging status of a bucket, you must be the bucket owner.
func (c *S3) PutBucketLogging(req PutBucketLoggingRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.BucketLoggingStatus)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?logging"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketNotification enables notifications of specified events for a
// bucket.
func (c *S3) PutBucketNotification(req PutBucketNotificationRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.NotificationConfiguration)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?notification"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketPolicy replaces a policy on a bucket. If the bucket already has
// a policy, the one in this request completely replaces it.
func (c *S3) PutBucketPolicy(req PutBucketPolicyRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.Policy)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?policy"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketRequestPayment sets the request payment configuration for a
// bucket. By default, the bucket owner pays for downloads from the bucket.
// This configuration parameter enables the bucket owner (only) to specify
// that the person requesting the download will be charged for the
// download.
func (c *S3) PutBucketRequestPayment(req PutBucketRequestPaymentRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.RequestPaymentConfiguration)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?requestPayment"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketTagging is undocumented.
func (c *S3) PutBucketTagging(req PutBucketTaggingRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.Tagging)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?tagging"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketVersioning sets the versioning state of an existing bucket. To
// set the versioning state, you must be the bucket owner.
func (c *S3) PutBucketVersioning(req PutBucketVersioningRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.VersioningConfiguration)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?versioning"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	if s := req.MFA; s != "" {

		httpReq.Header.Set("x-amz-mfa", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutBucketWebsite is undocumented.
func (c *S3) PutBucketWebsite(req PutBucketWebsiteRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.WebsiteConfiguration)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}?website"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// PutObject is undocumented.
func (c *S3) PutObject(req PutObjectRequest) (resp *PutObjectOutput, err error) {
	resp = &PutObjectOutput{}

	var body io.Reader

	body = bytes.NewReader(req.Body)

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ACL; s != "" {

		httpReq.Header.Set("x-amz-acl", s)
	}

	if s := req.CacheControl; s != "" {

		httpReq.Header.Set("Cache-Control", s)
	}

	if s := req.ContentDisposition; s != "" {

		httpReq.Header.Set("Content-Disposition", s)
	}

	if s := req.ContentEncoding; s != "" {

		httpReq.Header.Set("Content-Encoding", s)
	}

	if s := req.ContentLanguage; s != "" {

		httpReq.Header.Set("Content-Language", s)
	}

	if s := strconv.Itoa(req.ContentLength); req.ContentLength != 0 {

		httpReq.Header.Set("Content-Length", s)
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	if s := req.ContentType; s != "" {

		httpReq.Header.Set("Content-Type", s)
	}

	if s := req.Expires.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("Expires", s)
	}

	if s := req.GrantFullControl; s != "" {

		httpReq.Header.Set("x-amz-grant-full-control", s)
	}

	if s := req.GrantRead; s != "" {

		httpReq.Header.Set("x-amz-grant-read", s)
	}

	if s := req.GrantReadACP; s != "" {

		httpReq.Header.Set("x-amz-grant-read-acp", s)
	}

	if s := req.GrantWriteACP; s != "" {

		httpReq.Header.Set("x-amz-grant-write-acp", s)
	}

	if s := req.SSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-algorithm", s)
	}

	if s := req.SSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key", s)
	}

	if s := req.SSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key-MD5", s)
	}

	if s := req.SSEKMSKeyID; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-aws-kms-key-id", s)
	}

	if s := req.ServerSideEncryption; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption", s)
	}

	if s := req.StorageClass; s != "" {

		httpReq.Header.Set("x-amz-storage-class", s)
	}

	if s := req.WebsiteRedirectLocation; s != "" {

		httpReq.Header.Set("x-amz-website-redirect-location", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// PutObjectAcl uses the acl subresource to set the access control list
// permissions for an object that already exists in a bucket
func (c *S3) PutObjectAcl(req PutObjectAclRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.AccessControlPolicy)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}/{Key+}?acl"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.ACL; s != "" {

		httpReq.Header.Set("x-amz-acl", s)
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	if s := req.GrantFullControl; s != "" {

		httpReq.Header.Set("x-amz-grant-full-control", s)
	}

	if s := req.GrantRead; s != "" {

		httpReq.Header.Set("x-amz-grant-read", s)
	}

	if s := req.GrantReadACP; s != "" {

		httpReq.Header.Set("x-amz-grant-read-acp", s)
	}

	if s := req.GrantWrite; s != "" {

		httpReq.Header.Set("x-amz-grant-write", s)
	}

	if s := req.GrantWriteACP; s != "" {

		httpReq.Header.Set("x-amz-grant-write-acp", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// RestoreObject restores an archived copy of an object back into Amazon S3
func (c *S3) RestoreObject(req RestoreObjectRequest) (err error) {
	// NRE

	var body io.Reader

	b, err := xml.Marshal(req.RestoreRequest)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/{Bucket}/{Key+}?restore"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := req.VersionID; s != "" {

		q.Set("versionId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// TODO: extract the rest of the response

	return
}

// UploadPart uploads a part in a multipart upload. Note: After you
// initiate multipart upload and upload one or more parts, you must either
// complete or abort multipart upload in order to stop getting charged for
// storage of the uploaded parts. Only after you either complete or abort
// multipart upload, Amazon S3 frees up the parts storage and stops
// charging you for the parts storage.
func (c *S3) UploadPart(req UploadPartRequest) (resp *UploadPartOutput, err error) {
	resp = &UploadPartOutput{}

	var body io.Reader

	body = bytes.NewReader(req.Body)

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := strconv.Itoa(req.PartNumber); req.PartNumber != 0 {

		q.Set("partNumber", s)
	}

	if s := req.UploadID; s != "" {

		q.Set("uploadId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := strconv.Itoa(req.ContentLength); req.ContentLength != 0 {

		httpReq.Header.Set("Content-Length", s)
	}

	if s := req.ContentMD5; s != "" {

		httpReq.Header.Set("Content-MD5", s)
	}

	if s := req.SSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-algorithm", s)
	}

	if s := req.SSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key", s)
	}

	if s := req.SSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// UploadPartCopy uploads a part by copying data from an existing object as
// data source.
func (c *S3) UploadPartCopy(req UploadPartCopyRequest) (resp *UploadPartCopyOutput, err error) {
	resp = &UploadPartCopyOutput{}

	var body io.Reader

	uri := c.client.Endpoint + "/{Bucket}/{Key+}"

	uri = strings.Replace(uri, "{"+"Bucket"+"}", req.Bucket, -1)
	uri = strings.Replace(uri, "{"+"Bucket+"+"}", req.Bucket, -1)

	uri = strings.Replace(uri, "{"+"Key"+"}", req.Key, -1)
	uri = strings.Replace(uri, "{"+"Key+"+"}", req.Key, -1)

	q := url.Values{}

	if s := strconv.Itoa(req.PartNumber); req.PartNumber != 0 {

		q.Set("partNumber", s)
	}

	if s := req.UploadID; s != "" {

		q.Set("uploadId", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if s := req.CopySource; s != "" {

		httpReq.Header.Set("x-amz-copy-source", s)
	}

	if s := req.CopySourceIfMatch; s != "" {

		httpReq.Header.Set("x-amz-copy-source-if-match", s)
	}

	if s := req.CopySourceIfModifiedSince.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("x-amz-copy-source-if-modified-since", s)
	}

	if s := req.CopySourceIfNoneMatch; s != "" {

		httpReq.Header.Set("x-amz-copy-source-if-none-match", s)
	}

	if s := req.CopySourceIfUnmodifiedSince.Format(time.RFC822); s != "01 Jan 01 00:00 UTC" {

		httpReq.Header.Set("x-amz-copy-source-if-unmodified-since", s)
	}

	if s := req.CopySourceRange; s != "" {

		httpReq.Header.Set("x-amz-copy-source-range", s)
	}

	if s := req.CopySourceSSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-copy-source-server-side-encryption-customer-algorithm", s)
	}

	if s := req.CopySourceSSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-copy-source-server-side-encryption-customer-key", s)
	}

	if s := req.CopySourceSSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-copy-source-server-side-encryption-customer-key-MD5", s)
	}

	if s := req.SSECustomerAlgorithm; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-algorithm", s)
	}

	if s := req.SSECustomerKey; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key", s)
	}

	if s := req.SSECustomerKeyMD5; s != "" {

		httpReq.Header.Set("x-amz-server-side-encryption-customer-key-MD5", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// AbortMultipartUploadRequest is undocumented.
type AbortMultipartUploadRequest struct {
	Bucket   string `xml:"Bucket"`
	Key      string `xml:"Key"`
	UploadID string `xml:"uploadId"`
}

// AccessControlPolicy is undocumented.
type AccessControlPolicy struct {
	Grants []Grant `xml:"AccessControlList"`
	Owner  Owner   `xml:"Owner"`
}

// Bucket is undocumented.
type Bucket struct {
	CreationDate time.Time `xml:"CreationDate"`
	Name         string    `xml:"Name"`
}

// BucketLoggingStatus is undocumented.
type BucketLoggingStatus struct {
	LoggingEnabled LoggingEnabled `xml:"LoggingEnabled"`
}

// CORSConfiguration is undocumented.
type CORSConfiguration struct {
	CORSRules []CORSRule `xml:"CORSRule"`
}

// CORSRule is undocumented.
type CORSRule struct {
	AllowedHeaders []string `xml:"AllowedHeader"`
	AllowedMethods []string `xml:"AllowedMethod"`
	AllowedOrigins []string `xml:"AllowedOrigin"`
	ExposeHeaders  []string `xml:"ExposeHeader"`
	MaxAgeSeconds  int      `xml:"MaxAgeSeconds"`
}

// CloudFunctionConfiguration is undocumented.
type CloudFunctionConfiguration struct {
	CloudFunction  string   `xml:"CloudFunction"`
	Event          string   `xml:"Event"`
	Events         []string `xml:"Event"`
	ID             string   `xml:"Id"`
	InvocationRole string   `xml:"InvocationRole"`
}

// CommonPrefix is undocumented.
type CommonPrefix struct {
	Prefix string `xml:"Prefix"`
}

// CompleteMultipartUploadOutput is undocumented.
type CompleteMultipartUploadOutput struct {
	Bucket               string    `xml:"Bucket"`
	ETag                 string    `xml:"ETag"`
	Expiration           time.Time `xml:"x-amz-expiration"`
	Key                  string    `xml:"Key"`
	Location             string    `xml:"Location"`
	SSEKMSKeyID          string    `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption string    `xml:"x-amz-server-side-encryption"`
	VersionID            string    `xml:"x-amz-version-id"`
}

// CompleteMultipartUploadRequest is undocumented.
type CompleteMultipartUploadRequest struct {
	Bucket          string                   `xml:"Bucket"`
	Key             string                   `xml:"Key"`
	MultipartUpload CompletedMultipartUpload `xml:"CompleteMultipartUpload"`
	UploadID        string                   `xml:"uploadId"`
}

// CompletedMultipartUpload is undocumented.
type CompletedMultipartUpload struct {
	Parts []CompletedPart `xml:"Part"`
}

// CompletedPart is undocumented.
type CompletedPart struct {
	ETag       string `xml:"ETag"`
	PartNumber int    `xml:"PartNumber"`
}

// Condition is undocumented.
type Condition struct {
	HttpErrorCodeReturnedEquals string `xml:"HttpErrorCodeReturnedEquals"`
	KeyPrefixEquals             string `xml:"KeyPrefixEquals"`
}

// CopyObjectOutput is undocumented.
type CopyObjectOutput struct {
	CopyObjectResult     CopyObjectResult `xml:"CopyObjectResult"`
	CopySourceVersionID  string           `xml:"x-amz-copy-source-version-id"`
	Expiration           time.Time        `xml:"x-amz-expiration"`
	SSECustomerAlgorithm string           `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5    string           `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID          string           `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption string           `xml:"x-amz-server-side-encryption"`
}

// CopyObjectRequest is undocumented.
type CopyObjectRequest struct {
	ACL                            string            `xml:"x-amz-acl"`
	Bucket                         string            `xml:"Bucket"`
	CacheControl                   string            `xml:"Cache-Control"`
	ContentDisposition             string            `xml:"Content-Disposition"`
	ContentEncoding                string            `xml:"Content-Encoding"`
	ContentLanguage                string            `xml:"Content-Language"`
	ContentType                    string            `xml:"Content-Type"`
	CopySource                     string            `xml:"x-amz-copy-source"`
	CopySourceIfMatch              string            `xml:"x-amz-copy-source-if-match"`
	CopySourceIfModifiedSince      time.Time         `xml:"x-amz-copy-source-if-modified-since"`
	CopySourceIfNoneMatch          string            `xml:"x-amz-copy-source-if-none-match"`
	CopySourceIfUnmodifiedSince    time.Time         `xml:"x-amz-copy-source-if-unmodified-since"`
	CopySourceSSECustomerAlgorithm string            `xml:"x-amz-copy-source-server-side-encryption-customer-algorithm"`
	CopySourceSSECustomerKey       string            `xml:"x-amz-copy-source-server-side-encryption-customer-key"`
	CopySourceSSECustomerKeyMD5    string            `xml:"x-amz-copy-source-server-side-encryption-customer-key-MD5"`
	Expires                        time.Time         `xml:"Expires"`
	GrantFullControl               string            `xml:"x-amz-grant-full-control"`
	GrantRead                      string            `xml:"x-amz-grant-read"`
	GrantReadACP                   string            `xml:"x-amz-grant-read-acp"`
	GrantWriteACP                  string            `xml:"x-amz-grant-write-acp"`
	Key                            string            `xml:"Key"`
	Metadata                       map[string]string `xml:"x-amz-meta-"`
	MetadataDirective              string            `xml:"x-amz-metadata-directive"`
	SSECustomerAlgorithm           string            `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKey                 string            `xml:"x-amz-server-side-encryption-customer-key"`
	SSECustomerKeyMD5              string            `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID                    string            `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption           string            `xml:"x-amz-server-side-encryption"`
	StorageClass                   string            `xml:"x-amz-storage-class"`
	WebsiteRedirectLocation        string            `xml:"x-amz-website-redirect-location"`
}

// CopyObjectResult is undocumented.
type CopyObjectResult struct {
	ETag         string    `xml:"ETag"`
	LastModified time.Time `xml:"LastModified"`
}

// CopyPartResult is undocumented.
type CopyPartResult struct {
	ETag         string    `xml:"ETag"`
	LastModified time.Time `xml:"LastModified"`
}

// CreateBucketConfiguration is undocumented.
type CreateBucketConfiguration struct {
	LocationConstraint string `xml:"LocationConstraint"`
}

// CreateBucketOutput is undocumented.
type CreateBucketOutput struct {
	Location string `xml:"Location"`
}

// CreateBucketRequest is undocumented.
type CreateBucketRequest struct {
	ACL                       string                    `xml:"x-amz-acl"`
	Bucket                    string                    `xml:"Bucket"`
	CreateBucketConfiguration CreateBucketConfiguration `xml:"CreateBucketConfiguration"`
	GrantFullControl          string                    `xml:"x-amz-grant-full-control"`
	GrantRead                 string                    `xml:"x-amz-grant-read"`
	GrantReadACP              string                    `xml:"x-amz-grant-read-acp"`
	GrantWrite                string                    `xml:"x-amz-grant-write"`
	GrantWriteACP             string                    `xml:"x-amz-grant-write-acp"`
}

// CreateMultipartUploadOutput is undocumented.
type CreateMultipartUploadOutput struct {
	Bucket               string `xml:"Bucket"`
	Key                  string `xml:"Key"`
	SSECustomerAlgorithm string `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5    string `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID          string `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption string `xml:"x-amz-server-side-encryption"`
	UploadID             string `xml:"UploadId"`
}

// CreateMultipartUploadRequest is undocumented.
type CreateMultipartUploadRequest struct {
	ACL                     string            `xml:"x-amz-acl"`
	Bucket                  string            `xml:"Bucket"`
	CacheControl            string            `xml:"Cache-Control"`
	ContentDisposition      string            `xml:"Content-Disposition"`
	ContentEncoding         string            `xml:"Content-Encoding"`
	ContentLanguage         string            `xml:"Content-Language"`
	ContentType             string            `xml:"Content-Type"`
	Expires                 time.Time         `xml:"Expires"`
	GrantFullControl        string            `xml:"x-amz-grant-full-control"`
	GrantRead               string            `xml:"x-amz-grant-read"`
	GrantReadACP            string            `xml:"x-amz-grant-read-acp"`
	GrantWriteACP           string            `xml:"x-amz-grant-write-acp"`
	Key                     string            `xml:"Key"`
	Metadata                map[string]string `xml:"x-amz-meta-"`
	SSECustomerAlgorithm    string            `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKey          string            `xml:"x-amz-server-side-encryption-customer-key"`
	SSECustomerKeyMD5       string            `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID             string            `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption    string            `xml:"x-amz-server-side-encryption"`
	StorageClass            string            `xml:"x-amz-storage-class"`
	WebsiteRedirectLocation string            `xml:"x-amz-website-redirect-location"`
}

// Delete is undocumented.
type Delete struct {
	Objects []ObjectIdentifier `xml:"Object"`
	Quiet   bool               `xml:"Quiet"`
}

// DeleteBucketCorsRequest is undocumented.
type DeleteBucketCorsRequest struct {
	Bucket string `xml:"Bucket"`
}

// DeleteBucketLifecycleRequest is undocumented.
type DeleteBucketLifecycleRequest struct {
	Bucket string `xml:"Bucket"`
}

// DeleteBucketPolicyRequest is undocumented.
type DeleteBucketPolicyRequest struct {
	Bucket string `xml:"Bucket"`
}

// DeleteBucketRequest is undocumented.
type DeleteBucketRequest struct {
	Bucket string `xml:"Bucket"`
}

// DeleteBucketTaggingRequest is undocumented.
type DeleteBucketTaggingRequest struct {
	Bucket string `xml:"Bucket"`
}

// DeleteBucketWebsiteRequest is undocumented.
type DeleteBucketWebsiteRequest struct {
	Bucket string `xml:"Bucket"`
}

// DeleteMarkerEntry is undocumented.
type DeleteMarkerEntry struct {
	IsLatest     bool      `xml:"IsLatest"`
	Key          string    `xml:"Key"`
	LastModified time.Time `xml:"LastModified"`
	Owner        Owner     `xml:"Owner"`
	VersionID    string    `xml:"VersionId"`
}

// DeleteObjectOutput is undocumented.
type DeleteObjectOutput struct {
	DeleteMarker bool   `xml:"x-amz-delete-marker"`
	VersionID    string `xml:"x-amz-version-id"`
}

// DeleteObjectRequest is undocumented.
type DeleteObjectRequest struct {
	Bucket    string `xml:"Bucket"`
	Key       string `xml:"Key"`
	MFA       string `xml:"x-amz-mfa"`
	VersionID string `xml:"versionId"`
}

// DeleteObjectsOutput is undocumented.
type DeleteObjectsOutput struct {
	Deleted []DeletedObject `xml:"Deleted"`
	Errors  []Error         `xml:"Error"`
}

// DeleteObjectsRequest is undocumented.
type DeleteObjectsRequest struct {
	Bucket string `xml:"Bucket"`
	Delete Delete `xml:"Delete"`
	MFA    string `xml:"x-amz-mfa"`
}

// DeletedObject is undocumented.
type DeletedObject struct {
	DeleteMarker          bool   `xml:"DeleteMarker"`
	DeleteMarkerVersionID string `xml:"DeleteMarkerVersionId"`
	Key                   string `xml:"Key"`
	VersionID             string `xml:"VersionId"`
}

// Error is undocumented.
type Error struct {
	Code      string `xml:"Code"`
	Key       string `xml:"Key"`
	Message   string `xml:"Message"`
	VersionID string `xml:"VersionId"`
}

// ErrorDocument is undocumented.
type ErrorDocument struct {
	Key string `xml:"Key"`
}

// GetBucketAclOutput is undocumented.
type GetBucketAclOutput struct {
	Grants []Grant `xml:"AccessControlList"`
	Owner  Owner   `xml:"Owner"`
}

// GetBucketAclRequest is undocumented.
type GetBucketAclRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketCorsOutput is undocumented.
type GetBucketCorsOutput struct {
	CORSRules []CORSRule `xml:"CORSRule"`
}

// GetBucketCorsRequest is undocumented.
type GetBucketCorsRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketLifecycleOutput is undocumented.
type GetBucketLifecycleOutput struct {
	Rules []Rule `xml:"Rule"`
}

// GetBucketLifecycleRequest is undocumented.
type GetBucketLifecycleRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketLocationOutput is undocumented.
type GetBucketLocationOutput struct {
	LocationConstraint string `xml:"LocationConstraint"`
}

// GetBucketLocationRequest is undocumented.
type GetBucketLocationRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketLoggingOutput is undocumented.
type GetBucketLoggingOutput struct {
	LoggingEnabled LoggingEnabled `xml:"LoggingEnabled"`
}

// GetBucketLoggingRequest is undocumented.
type GetBucketLoggingRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketNotificationOutput is undocumented.
type GetBucketNotificationOutput struct {
	CloudFunctionConfiguration CloudFunctionConfiguration `xml:"CloudFunctionConfiguration"`
	QueueConfiguration         QueueConfiguration         `xml:"QueueConfiguration"`
	TopicConfiguration         TopicConfiguration         `xml:"TopicConfiguration"`
}

// GetBucketNotificationRequest is undocumented.
type GetBucketNotificationRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketPolicyOutput is undocumented.
type GetBucketPolicyOutput struct {
	Policy string `xml:"Policy"`
}

// GetBucketPolicyRequest is undocumented.
type GetBucketPolicyRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketRequestPaymentOutput is undocumented.
type GetBucketRequestPaymentOutput struct {
	Payer string `xml:"Payer"`
}

// GetBucketRequestPaymentRequest is undocumented.
type GetBucketRequestPaymentRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketTaggingOutput is undocumented.
type GetBucketTaggingOutput struct {
	TagSet []Tag `xml:"TagSet"`
}

// GetBucketTaggingRequest is undocumented.
type GetBucketTaggingRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketVersioningOutput is undocumented.
type GetBucketVersioningOutput struct {
	MFADelete string `xml:"MfaDelete"`
	Status    string `xml:"Status"`
}

// GetBucketVersioningRequest is undocumented.
type GetBucketVersioningRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetBucketWebsiteOutput is undocumented.
type GetBucketWebsiteOutput struct {
	ErrorDocument         ErrorDocument         `xml:"ErrorDocument"`
	IndexDocument         IndexDocument         `xml:"IndexDocument"`
	RedirectAllRequestsTo RedirectAllRequestsTo `xml:"RedirectAllRequestsTo"`
	RoutingRules          []RoutingRule         `xml:"RoutingRules"`
}

// GetBucketWebsiteRequest is undocumented.
type GetBucketWebsiteRequest struct {
	Bucket string `xml:"Bucket"`
}

// GetObjectAclOutput is undocumented.
type GetObjectAclOutput struct {
	Grants []Grant `xml:"AccessControlList"`
	Owner  Owner   `xml:"Owner"`
}

// GetObjectAclRequest is undocumented.
type GetObjectAclRequest struct {
	Bucket    string `xml:"Bucket"`
	Key       string `xml:"Key"`
	VersionID string `xml:"versionId"`
}

// GetObjectOutput is undocumented.
type GetObjectOutput struct {
	AcceptRanges            string            `xml:"accept-ranges"`
	Body                    []byte            `xml:"Body"`
	CacheControl            string            `xml:"Cache-Control"`
	ContentDisposition      string            `xml:"Content-Disposition"`
	ContentEncoding         string            `xml:"Content-Encoding"`
	ContentLanguage         string            `xml:"Content-Language"`
	ContentLength           int               `xml:"Content-Length"`
	ContentType             string            `xml:"Content-Type"`
	DeleteMarker            bool              `xml:"x-amz-delete-marker"`
	ETag                    string            `xml:"ETag"`
	Expiration              time.Time         `xml:"x-amz-expiration"`
	Expires                 time.Time         `xml:"Expires"`
	LastModified            time.Time         `xml:"Last-Modified"`
	Metadata                map[string]string `xml:"x-amz-meta-"`
	MissingMeta             int               `xml:"x-amz-missing-meta"`
	Restore                 string            `xml:"x-amz-restore"`
	SSECustomerAlgorithm    string            `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5       string            `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID             string            `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption    string            `xml:"x-amz-server-side-encryption"`
	VersionID               string            `xml:"x-amz-version-id"`
	WebsiteRedirectLocation string            `xml:"x-amz-website-redirect-location"`
}

// GetObjectRequest is undocumented.
type GetObjectRequest struct {
	Bucket                     string    `xml:"Bucket"`
	IfMatch                    string    `xml:"If-Match"`
	IfModifiedSince            time.Time `xml:"If-Modified-Since"`
	IfNoneMatch                string    `xml:"If-None-Match"`
	IfUnmodifiedSince          time.Time `xml:"If-Unmodified-Since"`
	Key                        string    `xml:"Key"`
	Range                      string    `xml:"Range"`
	ResponseCacheControl       string    `xml:"response-cache-control"`
	ResponseContentDisposition string    `xml:"response-content-disposition"`
	ResponseContentEncoding    string    `xml:"response-content-encoding"`
	ResponseContentLanguage    string    `xml:"response-content-language"`
	ResponseContentType        string    `xml:"response-content-type"`
	ResponseExpires            time.Time `xml:"response-expires"`
	SSECustomerAlgorithm       string    `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKey             string    `xml:"x-amz-server-side-encryption-customer-key"`
	SSECustomerKeyMD5          string    `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	VersionID                  string    `xml:"versionId"`
}

// GetObjectTorrentOutput is undocumented.
type GetObjectTorrentOutput struct {
	Body []byte `xml:"Body"`
}

// GetObjectTorrentRequest is undocumented.
type GetObjectTorrentRequest struct {
	Bucket string `xml:"Bucket"`
	Key    string `xml:"Key"`
}

// Grant is undocumented.
type Grant struct {
	Grantee    Grantee `xml:"Grantee"`
	Permission string  `xml:"Permission"`
}

// Grantee is undocumented.
type Grantee struct {
	DisplayName  string `xml:"DisplayName"`
	EmailAddress string `xml:"EmailAddress"`
	ID           string `xml:"ID"`
	Type         string `xml:"Type"`
	URI          string `xml:"URI"`
}

// HeadBucketRequest is undocumented.
type HeadBucketRequest struct {
	Bucket string `xml:"Bucket"`
}

// HeadObjectOutput is undocumented.
type HeadObjectOutput struct {
	AcceptRanges            string            `xml:"accept-ranges"`
	CacheControl            string            `xml:"Cache-Control"`
	ContentDisposition      string            `xml:"Content-Disposition"`
	ContentEncoding         string            `xml:"Content-Encoding"`
	ContentLanguage         string            `xml:"Content-Language"`
	ContentLength           int               `xml:"Content-Length"`
	ContentType             string            `xml:"Content-Type"`
	DeleteMarker            bool              `xml:"x-amz-delete-marker"`
	ETag                    string            `xml:"ETag"`
	Expiration              time.Time         `xml:"x-amz-expiration"`
	Expires                 time.Time         `xml:"Expires"`
	LastModified            time.Time         `xml:"Last-Modified"`
	Metadata                map[string]string `xml:"x-amz-meta-"`
	MissingMeta             int               `xml:"x-amz-missing-meta"`
	Restore                 string            `xml:"x-amz-restore"`
	SSECustomerAlgorithm    string            `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5       string            `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID             string            `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption    string            `xml:"x-amz-server-side-encryption"`
	VersionID               string            `xml:"x-amz-version-id"`
	WebsiteRedirectLocation string            `xml:"x-amz-website-redirect-location"`
}

// HeadObjectRequest is undocumented.
type HeadObjectRequest struct {
	Bucket               string    `xml:"Bucket"`
	IfMatch              string    `xml:"If-Match"`
	IfModifiedSince      time.Time `xml:"If-Modified-Since"`
	IfNoneMatch          string    `xml:"If-None-Match"`
	IfUnmodifiedSince    time.Time `xml:"If-Unmodified-Since"`
	Key                  string    `xml:"Key"`
	Range                string    `xml:"Range"`
	SSECustomerAlgorithm string    `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKey       string    `xml:"x-amz-server-side-encryption-customer-key"`
	SSECustomerKeyMD5    string    `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	VersionID            string    `xml:"versionId"`
}

// IndexDocument is undocumented.
type IndexDocument struct {
	Suffix string `xml:"Suffix"`
}

// Initiator is undocumented.
type Initiator struct {
	DisplayName string `xml:"DisplayName"`
	ID          string `xml:"ID"`
}

// LifecycleConfiguration is undocumented.
type LifecycleConfiguration struct {
	Rules []Rule `xml:"Rule"`
}

// LifecycleExpiration is undocumented.
type LifecycleExpiration struct {
	Date time.Time `xml:"Date"`
	Days int       `xml:"Days"`
}

// ListBucketsOutput is undocumented.
type ListBucketsOutput struct {
	Buckets []Bucket `xml:"Buckets"`
	Owner   Owner    `xml:"Owner"`
}

// ListMultipartUploadsOutput is undocumented.
type ListMultipartUploadsOutput struct {
	Bucket             string            `xml:"Bucket"`
	CommonPrefixes     []CommonPrefix    `xml:"CommonPrefixes"`
	Delimiter          string            `xml:"Delimiter"`
	EncodingType       string            `xml:"EncodingType"`
	IsTruncated        bool              `xml:"IsTruncated"`
	KeyMarker          string            `xml:"KeyMarker"`
	MaxUploads         int               `xml:"MaxUploads"`
	NextKeyMarker      string            `xml:"NextKeyMarker"`
	NextUploadIDMarker string            `xml:"NextUploadIdMarker"`
	Prefix             string            `xml:"Prefix"`
	UploadIDMarker     string            `xml:"UploadIdMarker"`
	Uploads            []MultipartUpload `xml:"Upload"`
}

// ListMultipartUploadsRequest is undocumented.
type ListMultipartUploadsRequest struct {
	Bucket         string `xml:"Bucket"`
	Delimiter      string `xml:"delimiter"`
	EncodingType   string `xml:"encoding-type"`
	KeyMarker      string `xml:"key-marker"`
	MaxUploads     int    `xml:"max-uploads"`
	Prefix         string `xml:"prefix"`
	UploadIDMarker string `xml:"upload-id-marker"`
}

// ListObjectVersionsOutput is undocumented.
type ListObjectVersionsOutput struct {
	CommonPrefixes      []CommonPrefix      `xml:"CommonPrefixes"`
	DeleteMarkers       []DeleteMarkerEntry `xml:"DeleteMarker"`
	Delimiter           string              `xml:"Delimiter"`
	EncodingType        string              `xml:"EncodingType"`
	IsTruncated         bool                `xml:"IsTruncated"`
	KeyMarker           string              `xml:"KeyMarker"`
	MaxKeys             int                 `xml:"MaxKeys"`
	Name                string              `xml:"Name"`
	NextKeyMarker       string              `xml:"NextKeyMarker"`
	NextVersionIDMarker string              `xml:"NextVersionIdMarker"`
	Prefix              string              `xml:"Prefix"`
	VersionIDMarker     string              `xml:"VersionIdMarker"`
	Versions            []ObjectVersion     `xml:"Version"`
}

// ListObjectVersionsRequest is undocumented.
type ListObjectVersionsRequest struct {
	Bucket          string `xml:"Bucket"`
	Delimiter       string `xml:"delimiter"`
	EncodingType    string `xml:"encoding-type"`
	KeyMarker       string `xml:"key-marker"`
	MaxKeys         int    `xml:"max-keys"`
	Prefix          string `xml:"prefix"`
	VersionIDMarker string `xml:"version-id-marker"`
}

// ListObjectsOutput is undocumented.
type ListObjectsOutput struct {
	CommonPrefixes []CommonPrefix `xml:"CommonPrefixes"`
	Contents       []Object       `xml:"Contents"`
	Delimiter      string         `xml:"Delimiter"`
	EncodingType   string         `xml:"EncodingType"`
	IsTruncated    bool           `xml:"IsTruncated"`
	Marker         string         `xml:"Marker"`
	MaxKeys        int            `xml:"MaxKeys"`
	Name           string         `xml:"Name"`
	NextMarker     string         `xml:"NextMarker"`
	Prefix         string         `xml:"Prefix"`
}

// ListObjectsRequest is undocumented.
type ListObjectsRequest struct {
	Bucket       string `xml:"Bucket"`
	Delimiter    string `xml:"delimiter"`
	EncodingType string `xml:"encoding-type"`
	Marker       string `xml:"marker"`
	MaxKeys      int    `xml:"max-keys"`
	Prefix       string `xml:"prefix"`
}

// ListPartsOutput is undocumented.
type ListPartsOutput struct {
	Bucket               string    `xml:"Bucket"`
	Initiator            Initiator `xml:"Initiator"`
	IsTruncated          bool      `xml:"IsTruncated"`
	Key                  string    `xml:"Key"`
	MaxParts             int       `xml:"MaxParts"`
	NextPartNumberMarker int       `xml:"NextPartNumberMarker"`
	Owner                Owner     `xml:"Owner"`
	PartNumberMarker     int       `xml:"PartNumberMarker"`
	Parts                []Part    `xml:"Part"`
	StorageClass         string    `xml:"StorageClass"`
	UploadID             string    `xml:"UploadId"`
}

// ListPartsRequest is undocumented.
type ListPartsRequest struct {
	Bucket           string `xml:"Bucket"`
	Key              string `xml:"Key"`
	MaxParts         int    `xml:"max-parts"`
	PartNumberMarker int    `xml:"part-number-marker"`
	UploadID         string `xml:"uploadId"`
}

// LoggingEnabled is undocumented.
type LoggingEnabled struct {
	TargetBucket string        `xml:"TargetBucket"`
	TargetGrants []TargetGrant `xml:"TargetGrants"`
	TargetPrefix string        `xml:"TargetPrefix"`
}

// MultipartUpload is undocumented.
type MultipartUpload struct {
	Initiated    time.Time `xml:"Initiated"`
	Initiator    Initiator `xml:"Initiator"`
	Key          string    `xml:"Key"`
	Owner        Owner     `xml:"Owner"`
	StorageClass string    `xml:"StorageClass"`
	UploadID     string    `xml:"UploadId"`
}

// NoncurrentVersionExpiration is undocumented.
type NoncurrentVersionExpiration struct {
	NoncurrentDays int `xml:"NoncurrentDays"`
}

// NoncurrentVersionTransition is undocumented.
type NoncurrentVersionTransition struct {
	NoncurrentDays int    `xml:"NoncurrentDays"`
	StorageClass   string `xml:"StorageClass"`
}

// NotificationConfiguration is undocumented.
type NotificationConfiguration struct {
	CloudFunctionConfiguration CloudFunctionConfiguration `xml:"CloudFunctionConfiguration"`
	QueueConfiguration         QueueConfiguration         `xml:"QueueConfiguration"`
	TopicConfiguration         TopicConfiguration         `xml:"TopicConfiguration"`
}

// Object is undocumented.
type Object struct {
	ETag         string    `xml:"ETag"`
	Key          string    `xml:"Key"`
	LastModified time.Time `xml:"LastModified"`
	Owner        Owner     `xml:"Owner"`
	Size         int       `xml:"Size"`
	StorageClass string    `xml:"StorageClass"`
}

// ObjectIdentifier is undocumented.
type ObjectIdentifier struct {
	Key       string `xml:"Key"`
	VersionID string `xml:"VersionId"`
}

// ObjectVersion is undocumented.
type ObjectVersion struct {
	ETag         string    `xml:"ETag"`
	IsLatest     bool      `xml:"IsLatest"`
	Key          string    `xml:"Key"`
	LastModified time.Time `xml:"LastModified"`
	Owner        Owner     `xml:"Owner"`
	Size         int       `xml:"Size"`
	StorageClass string    `xml:"StorageClass"`
	VersionID    string    `xml:"VersionId"`
}

// Owner is undocumented.
type Owner struct {
	DisplayName string `xml:"DisplayName"`
	ID          string `xml:"ID"`
}

// Part is undocumented.
type Part struct {
	ETag         string    `xml:"ETag"`
	LastModified time.Time `xml:"LastModified"`
	PartNumber   int       `xml:"PartNumber"`
	Size         int       `xml:"Size"`
}

// PutBucketAclRequest is undocumented.
type PutBucketAclRequest struct {
	ACL                 string              `xml:"x-amz-acl"`
	AccessControlPolicy AccessControlPolicy `xml:"AccessControlPolicy"`
	Bucket              string              `xml:"Bucket"`
	ContentMD5          string              `xml:"Content-MD5"`
	GrantFullControl    string              `xml:"x-amz-grant-full-control"`
	GrantRead           string              `xml:"x-amz-grant-read"`
	GrantReadACP        string              `xml:"x-amz-grant-read-acp"`
	GrantWrite          string              `xml:"x-amz-grant-write"`
	GrantWriteACP       string              `xml:"x-amz-grant-write-acp"`
}

// PutBucketCorsRequest is undocumented.
type PutBucketCorsRequest struct {
	Bucket            string            `xml:"Bucket"`
	CORSConfiguration CORSConfiguration `xml:"CORSConfiguration"`
	ContentMD5        string            `xml:"Content-MD5"`
}

// PutBucketLifecycleRequest is undocumented.
type PutBucketLifecycleRequest struct {
	Bucket                 string                 `xml:"Bucket"`
	ContentMD5             string                 `xml:"Content-MD5"`
	LifecycleConfiguration LifecycleConfiguration `xml:"LifecycleConfiguration"`
}

// PutBucketLoggingRequest is undocumented.
type PutBucketLoggingRequest struct {
	Bucket              string              `xml:"Bucket"`
	BucketLoggingStatus BucketLoggingStatus `xml:"BucketLoggingStatus"`
	ContentMD5          string              `xml:"Content-MD5"`
}

// PutBucketNotificationRequest is undocumented.
type PutBucketNotificationRequest struct {
	Bucket                    string                    `xml:"Bucket"`
	ContentMD5                string                    `xml:"Content-MD5"`
	NotificationConfiguration NotificationConfiguration `xml:"NotificationConfiguration"`
}

// PutBucketPolicyRequest is undocumented.
type PutBucketPolicyRequest struct {
	Bucket     string `xml:"Bucket"`
	ContentMD5 string `xml:"Content-MD5"`
	Policy     string `xml:"Policy"`
}

// PutBucketRequestPaymentRequest is undocumented.
type PutBucketRequestPaymentRequest struct {
	Bucket                      string                      `xml:"Bucket"`
	ContentMD5                  string                      `xml:"Content-MD5"`
	RequestPaymentConfiguration RequestPaymentConfiguration `xml:"RequestPaymentConfiguration"`
}

// PutBucketTaggingRequest is undocumented.
type PutBucketTaggingRequest struct {
	Bucket     string  `xml:"Bucket"`
	ContentMD5 string  `xml:"Content-MD5"`
	Tagging    Tagging `xml:"Tagging"`
}

// PutBucketVersioningRequest is undocumented.
type PutBucketVersioningRequest struct {
	Bucket                  string                  `xml:"Bucket"`
	ContentMD5              string                  `xml:"Content-MD5"`
	MFA                     string                  `xml:"x-amz-mfa"`
	VersioningConfiguration VersioningConfiguration `xml:"VersioningConfiguration"`
}

// PutBucketWebsiteRequest is undocumented.
type PutBucketWebsiteRequest struct {
	Bucket               string               `xml:"Bucket"`
	ContentMD5           string               `xml:"Content-MD5"`
	WebsiteConfiguration WebsiteConfiguration `xml:"WebsiteConfiguration"`
}

// PutObjectAclRequest is undocumented.
type PutObjectAclRequest struct {
	ACL                 string              `xml:"x-amz-acl"`
	AccessControlPolicy AccessControlPolicy `xml:"AccessControlPolicy"`
	Bucket              string              `xml:"Bucket"`
	ContentMD5          string              `xml:"Content-MD5"`
	GrantFullControl    string              `xml:"x-amz-grant-full-control"`
	GrantRead           string              `xml:"x-amz-grant-read"`
	GrantReadACP        string              `xml:"x-amz-grant-read-acp"`
	GrantWrite          string              `xml:"x-amz-grant-write"`
	GrantWriteACP       string              `xml:"x-amz-grant-write-acp"`
	Key                 string              `xml:"Key"`
}

// PutObjectOutput is undocumented.
type PutObjectOutput struct {
	ETag                 string    `xml:"ETag"`
	Expiration           time.Time `xml:"x-amz-expiration"`
	SSECustomerAlgorithm string    `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5    string    `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID          string    `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption string    `xml:"x-amz-server-side-encryption"`
	VersionID            string    `xml:"x-amz-version-id"`
}

// PutObjectRequest is undocumented.
type PutObjectRequest struct {
	ACL                     string            `xml:"x-amz-acl"`
	Body                    []byte            `xml:"Body"`
	Bucket                  string            `xml:"Bucket"`
	CacheControl            string            `xml:"Cache-Control"`
	ContentDisposition      string            `xml:"Content-Disposition"`
	ContentEncoding         string            `xml:"Content-Encoding"`
	ContentLanguage         string            `xml:"Content-Language"`
	ContentLength           int               `xml:"Content-Length"`
	ContentMD5              string            `xml:"Content-MD5"`
	ContentType             string            `xml:"Content-Type"`
	Expires                 time.Time         `xml:"Expires"`
	GrantFullControl        string            `xml:"x-amz-grant-full-control"`
	GrantRead               string            `xml:"x-amz-grant-read"`
	GrantReadACP            string            `xml:"x-amz-grant-read-acp"`
	GrantWriteACP           string            `xml:"x-amz-grant-write-acp"`
	Key                     string            `xml:"Key"`
	Metadata                map[string]string `xml:"x-amz-meta-"`
	SSECustomerAlgorithm    string            `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKey          string            `xml:"x-amz-server-side-encryption-customer-key"`
	SSECustomerKeyMD5       string            `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID             string            `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption    string            `xml:"x-amz-server-side-encryption"`
	StorageClass            string            `xml:"x-amz-storage-class"`
	WebsiteRedirectLocation string            `xml:"x-amz-website-redirect-location"`
}

// QueueConfiguration is undocumented.
type QueueConfiguration struct {
	Event  string   `xml:"Event"`
	Events []string `xml:"Event"`
	ID     string   `xml:"Id"`
	Queue  string   `xml:"Queue"`
}

// Redirect is undocumented.
type Redirect struct {
	HostName             string `xml:"HostName"`
	HttpRedirectCode     string `xml:"HttpRedirectCode"`
	Protocol             string `xml:"Protocol"`
	ReplaceKeyPrefixWith string `xml:"ReplaceKeyPrefixWith"`
	ReplaceKeyWith       string `xml:"ReplaceKeyWith"`
}

// RedirectAllRequestsTo is undocumented.
type RedirectAllRequestsTo struct {
	HostName string `xml:"HostName"`
	Protocol string `xml:"Protocol"`
}

// RequestPaymentConfiguration is undocumented.
type RequestPaymentConfiguration struct {
	Payer string `xml:"Payer"`
}

// RestoreObjectRequest is undocumented.
type RestoreObjectRequest struct {
	Bucket         string         `xml:"Bucket"`
	Key            string         `xml:"Key"`
	RestoreRequest RestoreRequest `xml:"RestoreRequest"`
	VersionID      string         `xml:"versionId"`
}

// RestoreRequest is undocumented.
type RestoreRequest struct {
	Days int `xml:"Days"`
}

// RoutingRule is undocumented.
type RoutingRule struct {
	Condition Condition `xml:"Condition"`
	Redirect  Redirect  `xml:"Redirect"`
}

// Rule is undocumented.
type Rule struct {
	Expiration                  LifecycleExpiration         `xml:"Expiration"`
	ID                          string                      `xml:"ID"`
	NoncurrentVersionExpiration NoncurrentVersionExpiration `xml:"NoncurrentVersionExpiration"`
	NoncurrentVersionTransition NoncurrentVersionTransition `xml:"NoncurrentVersionTransition"`
	Prefix                      string                      `xml:"Prefix"`
	Status                      string                      `xml:"Status"`
	Transition                  Transition                  `xml:"Transition"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// Tagging is undocumented.
type Tagging struct {
	TagSet []Tag `xml:"TagSet"`
}

// TargetGrant is undocumented.
type TargetGrant struct {
	Grantee    Grantee `xml:"Grantee"`
	Permission string  `xml:"Permission"`
}

// TopicConfiguration is undocumented.
type TopicConfiguration struct {
	Event  string   `xml:"Event"`
	Events []string `xml:"Event"`
	ID     string   `xml:"Id"`
	Topic  string   `xml:"Topic"`
}

// Transition is undocumented.
type Transition struct {
	Date         time.Time `xml:"Date"`
	Days         int       `xml:"Days"`
	StorageClass string    `xml:"StorageClass"`
}

// UploadPartCopyOutput is undocumented.
type UploadPartCopyOutput struct {
	CopyPartResult       CopyPartResult `xml:"CopyPartResult"`
	CopySourceVersionID  string         `xml:"x-amz-copy-source-version-id"`
	SSECustomerAlgorithm string         `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5    string         `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID          string         `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption string         `xml:"x-amz-server-side-encryption"`
}

// UploadPartCopyRequest is undocumented.
type UploadPartCopyRequest struct {
	Bucket                         string    `xml:"Bucket"`
	CopySource                     string    `xml:"x-amz-copy-source"`
	CopySourceIfMatch              string    `xml:"x-amz-copy-source-if-match"`
	CopySourceIfModifiedSince      time.Time `xml:"x-amz-copy-source-if-modified-since"`
	CopySourceIfNoneMatch          string    `xml:"x-amz-copy-source-if-none-match"`
	CopySourceIfUnmodifiedSince    time.Time `xml:"x-amz-copy-source-if-unmodified-since"`
	CopySourceRange                string    `xml:"x-amz-copy-source-range"`
	CopySourceSSECustomerAlgorithm string    `xml:"x-amz-copy-source-server-side-encryption-customer-algorithm"`
	CopySourceSSECustomerKey       string    `xml:"x-amz-copy-source-server-side-encryption-customer-key"`
	CopySourceSSECustomerKeyMD5    string    `xml:"x-amz-copy-source-server-side-encryption-customer-key-MD5"`
	Key                            string    `xml:"Key"`
	PartNumber                     int       `xml:"partNumber"`
	SSECustomerAlgorithm           string    `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKey                 string    `xml:"x-amz-server-side-encryption-customer-key"`
	SSECustomerKeyMD5              string    `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	UploadID                       string    `xml:"uploadId"`
}

// UploadPartOutput is undocumented.
type UploadPartOutput struct {
	ETag                 string `xml:"ETag"`
	SSECustomerAlgorithm string `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5    string `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyID          string `xml:"x-amz-server-side-encryption-aws-kms-key-id"`
	ServerSideEncryption string `xml:"x-amz-server-side-encryption"`
}

// UploadPartRequest is undocumented.
type UploadPartRequest struct {
	Body                 []byte `xml:"Body"`
	Bucket               string `xml:"Bucket"`
	ContentLength        int    `xml:"Content-Length"`
	ContentMD5           string `xml:"Content-MD5"`
	Key                  string `xml:"Key"`
	PartNumber           int    `xml:"partNumber"`
	SSECustomerAlgorithm string `xml:"x-amz-server-side-encryption-customer-algorithm"`
	SSECustomerKey       string `xml:"x-amz-server-side-encryption-customer-key"`
	SSECustomerKeyMD5    string `xml:"x-amz-server-side-encryption-customer-key-MD5"`
	UploadID             string `xml:"uploadId"`
}

// VersioningConfiguration is undocumented.
type VersioningConfiguration struct {
	MFADelete string `xml:"MfaDelete"`
	Status    string `xml:"Status"`
}

// WebsiteConfiguration is undocumented.
type WebsiteConfiguration struct {
	ErrorDocument         ErrorDocument         `xml:"ErrorDocument"`
	IndexDocument         IndexDocument         `xml:"IndexDocument"`
	RedirectAllRequestsTo RedirectAllRequestsTo `xml:"RedirectAllRequestsTo"`
	RoutingRules          []RoutingRule         `xml:"RoutingRules"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name

var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
