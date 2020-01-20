### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/dynamodb/expression`: Allow AttributeValue as a value to BuildOperand. ([#3057](https://github.com/aws/aws-sdk-go/pull/3057))
  * This change fixes the SDK's behavior with DynamoDB Expression builder to not double marshal AttributeValues when used as BuildOperands, `Value` type. The AttributeValue will be used in the expression as the specific value set in the AttributeValue, instead of encoded as another AttributeValue.
