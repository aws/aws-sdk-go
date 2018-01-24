### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/s3/s3manager`: Fix check for nil OrigErr in Error() [#1749](https://github.com/aws/aws-sdk-go/issues/1749)
    * S3 Manager's `Error` type did not check for nil of `OrigErr` when calling `Error()`
    * Fixes [#1748](https://github.com/aws/aws-sdk-go/issues/1748)
