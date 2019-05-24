### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/dynamodb/expression`: Fix Builder with KeyCondition example ([#2618](https://github.com/aws/aws-sdk-go/pull/2618))
  * Fixes the ExampleBuilder_WithKeyCondition example to include the ExpressionAttributeNames member being set.
  * Related to [aws/aws-sdk-go-v2#285](https://github.com/aws/aws-sdk-go-v2/issues/285)
* `private/model/api`: Improve SDK API reference doc generation ([#2617](https://github.com/aws/aws-sdk-go/pull/2617))
  * Improves the SDK's generated documentation for API client, operation, and types. This fixes several bugs in the doc generation causing poor formatting, an difficult to read reference documentation.
  * Fixes [#2572](https://github.com/aws/aws-sdk-go/pull/2572)
  * Fixes [#2374](https://github.com/aws/aws-sdk-go/pull/2374)
