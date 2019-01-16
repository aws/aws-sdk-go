### SDK Features

### SDK Enhancements
* `aws/request: Improve error handling in shouldRetryCancel ([#2298](https://github.com/aws/aws-sdk-go/pull/2298))
  * Simplifies and improves SDK's detection of HTTP request errors that should be retried. Previously the SDK would incorrectly attempt to retry `EHOSTDOWN` connection errors. This change fixes this, by using the `Temporary` interface when available.

### SDK Bugs
