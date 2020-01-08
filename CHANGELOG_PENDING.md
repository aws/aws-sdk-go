### SDK Features
* `service/dynamodb/dynamodbattribute`: Go 1.9+, Add caching of struct serialization ([#3070](https://github.com/aws/aws-sdk-go/pull/3070))
  * For Go 1.9 and above, adds struct field caching to the SDK's DynamoDB AttributeValue marshalers and unmarshalers. This significantly reduces time, and overall allocations of the (un)marshalers by caching the reflected structure's fields. This should improve the performance of applications using DynamoDB AttributeValue (un)marshalers.

### SDK Enhancements

### SDK Bugs
