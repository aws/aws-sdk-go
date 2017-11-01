### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/request`: Fix bug in request presign creating invalide URL ([#1624](https://github.com/aws/aws-sdk-go/pull/1624))
  * Fixes a bug the Request Presign and PresignRequest methods that would allow a invalid expire duration as input. A expire time of 0 would be interpreted by the SDK to generate a normal request signature, not a presigned URL. This caused the returned URL unusable.
  * Fixes [#1617](https://github.com/aws/aws-sdk-go/issues/1617)
