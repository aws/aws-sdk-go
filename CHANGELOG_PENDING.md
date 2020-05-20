### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/s3/s3crypto`: Add missing return in encryption client ([#3258](https://github.com/aws/aws-sdk-go/pull/3258))
  * Fixes a missing return in the encryption client that was causing a nil dereference panic.
