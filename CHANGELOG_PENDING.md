### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/s3`: Close http.Response.Body in CopyObject, UploadPartCopy and CompleteMultipartUpload operations
  * Fixes [#4037](https://github.com/aws/aws-sdk-go/issues/4037)
