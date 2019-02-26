### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/request`: Fix RequestUserAgent tests to be stable ([#2462](https://github.com/aws/aws-sdk-go/pull/2462))
  * Fixes the request User-Agent unit tests to be stable across all platforms and environments.
  * Fixes [#2366](https://github.com/aws/aws-sdk-go/issues/2366)
* `aws/ec2metadata`: Fix EC2 Metadata client panic with debug logging ([#2461](https://github.com/aws/aws-sdk-go/pull/2461))
  * Fixes a panic that could occur witihin the EC2 Metadata client when both `AWS_EC2_METADATA_DISABLED` env var is set and log level is LogDebugWithHTTPBody.
* `private/protocol/rest`: Trim space in header key and value ([#2460](https://github.com/aws/aws-sdk-go/pull/2460))
  * Updates the REST protocol marshaler to trip leading and trailing space from header keys and values before setting the HTTP request header. Fixes a bug when using S3 metadata where metadata values with leading spaces would trigger request signature validation errors when the request is received by the service.
  * Fixes [#2448](https://github.com/aws/aws-sdk-go/issues/2448)
