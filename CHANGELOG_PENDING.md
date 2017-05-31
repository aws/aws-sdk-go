### SDK Features

### SDK Enhancements
* `aws/defaults`: Exports shared credentials and config default filenames used by the SDK. [#1308](https://github.com/aws/aws-sdk-go/pull/1308)
  * Adds SharedCredentialsFilename and SharedConfigFilename functions to defaults package.

### SDK Bugs
* `aws/credentials`: Fixes shared credential provider's default filename on Windows. [#1308](https://github.com/aws/aws-sdk-go/pull/1308)
  * The shared credentials provider would attempt to use the wrong filename on Windows if the `HOME` environment variable was defined.
* `service/s3/s3manager`: service/s3/s3manager: Fix Downloader ignoring Range get parameter [#1311](https://github.com/aws/aws-sdk-go/pull/1311)
  * Fixes the S3 Download Manager ignoring the GetObjectInput's Range parameter. If this parameter is provided it will force the downloader to fallback to a single GetObject request disabling concurrency and automatic part size gets.
  * Fixes [#1296](https://github.com/aws/aws-sdk-go/issues/1296)
