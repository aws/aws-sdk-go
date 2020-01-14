### SDK Features

### SDK Enhancements
* `service/s3/s3crypto`: Added X-Ray support to encrypt/decrypt clients ([#2912](https://github.com/aws/aws-sdk-go/pull/2912))
  * Adds support for passing Context down to the crypto client's KMS client enabling tracing for tools like X-Ray, and metrics.

### SDK Bugs
