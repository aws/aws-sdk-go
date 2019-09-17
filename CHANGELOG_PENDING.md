### SDK Features
* `service/s3/s3manager`: Add Upload Buffer Provider ([#2792](https://github.com/aws/aws-sdk-go/pull/2792))
  * Adds a new `BufferProvider` member for specifying how part data can be buffered in memory.
  * Windows platforms will now default to buffering 1MB per part to reduce contention when uploading files.
  * Non-Windows platforms will continue to employ a non-buffering behavior.

### SDK Enhancements
* `awstesting/integration/performance/s3UploadManager`: Extended to support benchmarking and usage of new `BufferProvider` ([#2792](https://github.com/aws/aws-sdk-go/pull/2792))

### SDK Bugs
