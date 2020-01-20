### SDK Features
* `services/transcribestreamingservice`: Support for Amazon Transcribe Streaming ([#3048](https://github.com/aws/aws-sdk-go/pull/3048))
  * The SDK now supports the Amazon Transcribe Streaming APIs by utilizing event stream encoding over HTTP/2
  * See [Amazon Transcribe Developer Guide](https://docs.aws.amazon.com/transcribe/latest/dg)
  * Fixes [#2487](https://github.com/aws/aws-sdk-go/issues/2487)

### SDK Enhancements

### SDK Bugs
* `service/dynamodb/expression`: Allow AttributeValue as a value to BuildOperand. ([#3057](https://github.com/aws/aws-sdk-go/pull/3057))
  * This change fixes the SDK's behavior with DynamoDB Expression builder to not double marshal AttributeValues when used as BuildOperands, `Value` type. The AttributeValue will be used in the expression as the specific value set in the AttributeValue, instead of encoded as another AttributeValue.
