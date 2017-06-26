### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/signer/v4`: checking length on `stripExcessSpaces` [#1372](https://github.com/aws/aws-sdk-go/issues/1372)
  * Fixes a bug where `stripExcessSpaces` did not check length against the slice.
  * Fixes: [#1371](https://github.com/aws/aws-sdk-go/issues/1371)
