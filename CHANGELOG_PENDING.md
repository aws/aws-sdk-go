### SDK Features

### SDK Enhancements

### SDK Bugs
* `private/protocol/query`: Fix query protocol handling of nested byte slices ([#1557](https://github.com/aws/aws-sdk-go/issues/1557))
  * Fixes the query protocol to correctly marshal nested []byte values of API operations.
* `service/s3`: Fix PutObject and UploadPart API to include ContentMD5 field ([#1559](https://github.com/aws/aws-sdk-go/pull/1559))
  * Fixes the SDK's S3 PutObject and UploadPart API code generation to correctly render the ContentMD5 field into the associated input types for these two API operations.
  * Fixes [#1553](https://github.com/aws/aws-sdk-go/pull/1553)
