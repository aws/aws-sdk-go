### SDK Features

### SDK Enhancements

### SDK Bugs
* Fix improper use of printf-style functions.
  * Required for Go 1.24.
* `service/s3/s3manager`: Abort multipart download if object is modified during download
  * Fixes [4986](https://github.com/aws/aws-sdk-go/issues/4986)
