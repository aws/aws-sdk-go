### SDK Features

### SDK Enhancements
* `service/s3/s3manager`: adding cleanup function to batch objects [#1375](https://github.com/aws/aws-sdk-go/issues/1375)
  * This enhancement will add an After field that will be called after each iteration of the batch operation.

### SDK Bugs
* `aws/signer/v4`: checking length on `stripExcessSpaces` [#1372](https://github.com/aws/aws-sdk-go/issues/1372)
  * Fixes a bug where `stripExcessSpaces` did not check length against the slice.
  * Fixes: [#1371](https://github.com/aws/aws-sdk-go/issues/1371)
