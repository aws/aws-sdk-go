### SDK Features

### SDK Enhancements
* `service/cloudwatch`: Add helper to send request payload as GZIP content encoding
  * Adds a new helper, `WithGzipRequest` to the `cloudwatch` package. The helper will configure the payload to be sent as `content-encoding: gzip`. It is supported by operations like `PutMetricData`. See the service's API Reference documentation for other operations supported.
### SDK Bugs
