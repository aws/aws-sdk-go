### SDK Features

### SDK Enhancements

### SDK Bugs
* `private/protocol`: Use correct Content-Type for rest json protocol ([#2497](https://github.com/aws/aws-sdk-go/pull/2497))
  * Updates the SDK to use the correct `application/json` content type for all rest json protocol based AWS services. This fixes the bug where the jsonrpc protocol's `application/x-amz-json-X.Y` content type would be used for services like Pinpoint SMS.
