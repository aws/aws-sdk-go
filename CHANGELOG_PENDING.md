### SDK Features

### SDK Enhancements
* `private/protocol`: Add protocol tests for blob types and headers ([#2770](https://github.com/aws/aws-sdk-go/pull/2770))
  * Adds RESTJSON and RESTXML protocol tests for blob headers.
  * Related to [#750](https://github.com/aws/aws-sdk-go/issues/750)

### SDK Bugs

* `service/dynamodb/expression`: Improved reporting of bad key conditions ([#2775](https://github.com/aws/aws-sdk-go/pull/2775))
  * Improved error reporting when invalid key conditions are constructed using KeyConditionBuilder
