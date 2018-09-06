### SDK Features

### SDK Enhancements
* `private/protocol`: Serialization errors will now be wrapped in `awserr.RequestFailure` types ([#2135](https://github.com/aws/aws-sdk-go/pull/2135))
  * Updates the SDK protocol unmarshaling to handle the `SerializationError` as a request failure allowing for inspection of `requestID`s and status codes.

### SDK Bugs
