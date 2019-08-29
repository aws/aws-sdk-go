### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/client`: Updates logic for calculating the delay after which a request can be retried. Retry delay now includes the Retry-After duration specified in a request [#2796](https://github.com/aws/aws-sdk-go/pull/2796).
  * Fixes broken test for retry delays for throttled exceptions. Fixes [#2795](https://github.com/aws/aws-sdk-go/issues/2795)
  