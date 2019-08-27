### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/csm`: Fix metricChan's unsafe atomic operations ([#2785](https://github.com/aws/aws-sdk-go/pull/2785))
  * Fixes [#2784](https://github.com/aws/aws-sdk-go/issues/2784) test failure caused by the metricChan.paused member being a value instead of a pointer. If the metricChan value was ever copied the atomic operations performed on paused would be invalid.
