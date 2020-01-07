### SDK Features
* `service/dynamodb/dynamodbattribute`: Add caching of struct serialization ([#3070](https://github.com/aws/aws-sdk-go/pull/3070))
  * Adds struct field caching to the SDK's DynamoDB AttributeValue marshalers and unmarshalers. This significantly reduces time, and overall allocations of the (un)marshalers by caching the reflected structure's fields. This should improve the performance of applications using DynamoDB AttributeValue (un)marshallers.

### SDK Enhancements

### SDK Bugs
