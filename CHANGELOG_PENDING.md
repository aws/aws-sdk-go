### SDK Features

### SDK Enhancements
* `example/service/ecr`: Add create and delete repository examples ([#3221](https://github.com/aws/aws-sdk-go/pull/3221))
  * Adds examples demonstrating how you can create and delete repositories with the SDK.

### SDK Bugs
* `service/s3/s3crypto`: Add missing return in encryption client ([#3258](https://github.com/aws/aws-sdk-go/pull/3258))
  * Fixes a missing return in the encryption client that was causing a nil dereference panic.
