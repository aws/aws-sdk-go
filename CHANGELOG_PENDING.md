### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/s3/s3crypto`: Fixes a bug where `gcmEncryptReader.Read` would return
an invalid number of bytes  as having been written to `data []byte`. ([#3005](https://github.com/aws/aws-sdk-go/pull/3005))
    * Fixes [#2999](https://github.com/aws/aws-sdk-go/issues/2999)
