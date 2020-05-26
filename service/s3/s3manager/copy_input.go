package s3manager

import "time"

// CopyInput provides input parameters for copying an object from the CopySource
// to the destination Bucket and Key. It is based on the CopyObjectInput, but some
// fields may be omitted.
type CopyInput struct {
	_ struct{}

	// The canned ACL to apply to the object.
	ACL *string

	// The name of the destination bucket.
	//
	// Bucket is a required field
	Bucket *string

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// The name of the source bucket and key name of the source object, separated
	// by a slash (/). Must be URL-encoded.
	//
	// CopySource is a required field
	CopySource *string

	// Specifies the algorithm to use when decrypting the source object (for example,
	// AES256).
	CopySourceSSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one
	// that was used when the source object was created.
	CopySourceSSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	CopySourceSSECustomerKeyMD5 *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string

	// Allows grantee to read the object data and its metadata.
	GrantRead *string

	// Allows grantee to read the object ACL.
	GrantReadACP *string

	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string

	// The key of the destination object.
	//
	// Key is a required field
	Key *string

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string

	// Specifies whether the metadata is copied from the source object or replaced
	// with metadata provided in the request.
	MetadataDirective *string

	// Specifies whether you want to apply a Legal Hold to the copied object.
	ObjectLockLegalHoldStatus *string

	// The Object Lock mode that you want to apply to the copied object.
	ObjectLockMode *string

	// The date and time when you want the copied object's Object Lock to expire.
	ObjectLockRetainUntilDate *time.Time

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer *string

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// S3 does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// Specifies the AWS KMS Encryption Context to use for object encryption. The
	// value of this header is a base64-encoded UTF-8 string holding JSON with the
	// encryption context key-value pairs.
	SSEKMSEncryptionContext *string

	// Specifies the AWS KMS key ID to use for object encryption. All GET and PUT
	// requests for an object protected by AWS KMS will fail if not made via SSL
	// or using SigV4. For information about configuring using any of the officially
	// supported AWS SDKs and AWS CLI, see Specifying the Signature Version in Request
	// Authentication (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version)
	// in the Amazon S3 Developer Guide.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon
	// S3 (for example, AES256, aws:kms).
	ServerSideEncryption *string

	// The region of the source bucket if it differs from that of the S3 client.
	SourceRegion *string

	// The type of storage to use for the object. Defaults to 'STANDARD'.
	StorageClass *string

	// The tag-set for the object destination object this value must be used in
	// conjunction with the TaggingDirective. The tag-set must be encoded as URL
	// Query parameters.
	Tagging *string

	// Specifies whether the object tag-set are copied from the source object or
	// replaced with tag-set provided in the request.
	TaggingDirective *string

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string
}
