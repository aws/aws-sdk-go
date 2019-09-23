### SDK Features
* `service/dynamodb/dynamodbattribute`: Add EnableEmptyCollections flag to Encoder and Decoder ([#2834](https://github.com/aws/aws-sdk-go/pull/2834))
  * The `Encoder` and `Decoder` types have been enhanced to allow support for specifying the SDK's behavior when marshaling structures, maps, and slices to DynamoDB.
  * When `EnableEmptyCollections` is set to `True` the SDK will preserve the empty of these types in DynamoDB rather then encoding a NULL AttributeValue.
  * Fixes [#682](https://github.com/aws/aws-sdk-go/issues/682)
  * Fixes [#1890](https://github.com/aws/aws-sdk-go/issues/1890)
  * Fixes [#2746](https://github.com/aws/aws-sdk-go/issues/2746)

### SDK Enhancements

### SDK Bugs
