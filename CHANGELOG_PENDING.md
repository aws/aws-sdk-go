### SDK Features
* `service/s3/s3manager`: Add Download Buffer Provider ([#2823](https://github.com/aws/aws-sdk-go/pull/2823))
  * Adds a new `BufferProvider` member for specifying how part data can be buffered in memory when copying from the http response body.
  * Windows platforms will now default to buffering 1MB per part to reduce contention when downloading files.
  * Non-Windows platforms will continue to employ a non-buffering behavior.
  * Fixes [#2180](https://github.com/aws/aws-sdk-go/issues/2180)
  * Fixes [#2662](https://github.com/aws/aws-sdk-go/issues/2662)

### SDK Enhancements

### SDK Bugs
