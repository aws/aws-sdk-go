### SDK Features
* `service/dynamodb/dynamodbattribute`: Add EnableEmptyCollections flag to Encoder and Decoder ([#2834](https://github.com/aws/aws-sdk-go/pull/2834))
  * The `Encoder` and `Decoder` types have been enhanced to allow support for specifying the SDK's behavior when marshaling structures, maps, and slices to DynamoDB.
  * When `EnableEmptyCollections` is set to `True` the SDK will preserve the empty of these types in DynamoDB rather then encoding a NULL AttributeValue.
  * Fixes [#682](https://github.com/aws/aws-sdk-go/issues/682)
  * Fixes [#1890](https://github.com/aws/aws-sdk-go/issues/1890)
  * Fixes [#2746](https://github.com/aws/aws-sdk-go/issues/2746)
* `service/s3/s3manager`: Add Download Buffer Provider ([#2823](https://github.com/aws/aws-sdk-go/pull/2823))
  * Adds a new `BufferProvider` member for specifying how part data can be buffered in memory when copying from the http response body.
  * Windows platforms will now default to buffering 1MB per part to reduce contention when downloading files.
  * Non-Windows platforms will continue to employ a non-buffering behavior.
  * Fixes [#2180](https://github.com/aws/aws-sdk-go/issues/2180)
  * Fixes [#2662](https://github.com/aws/aws-sdk-go/issues/2662)

### SDK Enhancements

### SDK Bugs
