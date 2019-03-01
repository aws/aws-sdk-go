### SDK Features

### SDK Enhancements
* `example/service/s3`: Add example of S3 download with progress ([#2456](https://github.com/aws/aws-sdk-go/pull/2456))
  * Adds a new example to the S3 service's examples. This example shows how you could use the S3's GetObject API call in conjunction with a custom writer keeping track of progress.
  * Related to [#1868](https://github.com/aws/aws-sdk-go/pull/1868), [#2468](https://github.com/aws/aws-sdk-go/pull/2468)

### SDK Bugs
* `aws/session`: Allow HTTP Proxy with custom CA bundle ([#2343](https://github.com/aws/aws-sdk-go/pull/2343))
  * Ensures Go HTTP Client's  `ProxyFromEnvironment` functionality is still enabled when  custom CA bundles are used with the SDK.
  * Fix [#2287](https://github.com/aws/aws-sdk-go/pull/2287)
