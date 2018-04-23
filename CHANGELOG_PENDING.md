### SDK Features

### SDK Enhancements

### SDK Bugs
* Fix XML unmarshaler not correctly unmarshaling list of timestamp values ([#1894](https://github.com/aws/aws-sdk-go/pull/1894))
  * Fixes a bug in the XML unmarshaler that would incorrectly try to unmarshal "time.Time" parameters that did not have the struct tag type on them. This would occur for nested lists like CloudWatch's GetMetricDataResponse MetricDataResults timestamp parameters.
  * Fixes [#1892](https://github.com/aws/aws-sdk-go/issues/1892)
