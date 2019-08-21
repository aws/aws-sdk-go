### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/request`: Fix IsErrorRetryable returning true for nil error ([#2774](https://github.com/aws/aws-sdk-go/pull/2774))
  * Fixes [#2773](https://github.com/aws/aws-sdk-go/pull/2773) where the IsErrorRetryable helper was incorrectly returning true for nil errors passed in. IsErrorRetryable will now correctly return false when the passed in error is nil like documented.

