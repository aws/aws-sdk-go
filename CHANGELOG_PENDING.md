### SDK Features

### SDK Enhancements
* `aws/defaults`: Exports shared credentials and config default filenames used by the SDK. [#1308](https://github.com/aws/aws-sdk-go/pull/1308)
  * Adds SharedCredentialsFilename and SharedConfigFilename functions to defaults package.

### SDK Bugs
* `aws/credentials`: Fixes shared credential provider's default filename on Windows. [#1308](https://github.com/aws/aws-sdk-go/pull/1308)
  * The shared credentials provider would attempt to use the wrong filename on Windows if the `HOME` environment variable was defined.
