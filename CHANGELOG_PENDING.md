### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/csm`: Fix metricChan's unsafe atomic operations ([#2785](https://github.com/aws/aws-sdk-go/pull/2785))
  * Fixes [#2784](https://github.com/aws/aws-sdk-go/issues/2784) test failure caused by the metricChan.paused member being a value instead of a pointer. If the metricChan value was ever copied the atomic operations performed on paused would be invalid.
* `aws/client`: Updates logic for request retry delay calculation ([#2796](https://github.com/aws/aws-sdk-go/pull/2796))
  * Updates logic for calculating the delay after which a request can be retried. Retry delay now includes the Retry-After duration specified in a request. Fixes broken test for retry delays for throttled exceptions.
  * Fixes [#2795](https://github.com/aws/aws-sdk-go/issues/2795)
