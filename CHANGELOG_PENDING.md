### SDK Features

### SDK Enhancements
* `service/s3`:  Use interfaces assertions instead of ValuesAtPath for S3 field lookups. [#1401](https://github.com/aws/aws-sdk-go/pull/1401)
  * Improves the performance across the board for all S3 API calls by removing the usage of `ValuesAtPath` being used for every S3 API call.

### SDK Bugs
*`aws/request`: waiter test bug
  * waiters_test.go file would sometimes fail due to travis hiccups. This occurs because a test would sometimes fail the cancel check and succeed the timeout. However, the timeout check should never occur in that test. This fix introduces a new field that dictates how waiters will sleep.
