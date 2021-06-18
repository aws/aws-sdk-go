### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/signer/v4`: Add X-Amz-Object-Lock-* as unhoisted presign headers
  * Updates the SigV4 signer to exlucde the X-Amz-Object-Lock- group of headers to not be signed as a part of the query string for presigned URLs.
* `private/protocol`: Add support for UTC offset for ISO8601 datetime formats ([#3960](https://github.com/aws/aws-sdk-go/pull/3960))
  * Updates the SDK's parsing of ISO8601 date time formats to support UTC offsets.
