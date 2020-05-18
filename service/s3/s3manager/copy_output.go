package s3manager

// CopyOutput contains the result of an object copy.
type CopyOutput struct {
	// The name of the bucket that contains the newly created object.
	Bucket *string

	// The version of the source object that was copied, if you have enabled versioning
	// on the source bucket.
	CopySourceVersionId *string `location:"header" locationName:"x-amz-copy-source-version-id" type:"string"`

	// Entity tag that identifies the newly created object's data. Objects with
	// different object data will have different entity tags. The entity tag is
	// an opaque string. The entity tag may or may not be an MD5 digest of the object
	// data. If the entity tag is not an MD5 digest of the object data, it will
	// contain one or more nonhexadecimal characters and/or will consist of less
	// than 32 or more than 32 hexadecimal digits.
	ETag *string

	// If the object expiration is configured, this will contain the expiration
	// date (expiry-date) and rule ID (rule-id). The value of rule-id is URL encoded.
	Expiration *string

	// The object key of the newly created object.
	Key *string

	// The URI that identifies the newly created object.
	Location *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged *string

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// If you specified server-side encryption either with an Amazon S3-managed
	// encryption key or an AWS KMS customer master key (CMK) in your initiate multipart
	// upload request, the response includes this header. It confirms the encryption
	// algorithm that Amazon S3 used to encrypt the object.
	ServerSideEncryption *string

	// Version ID of the newly created object, in case the bucket has versioning
	// turned on.
	VersionId *string
}
