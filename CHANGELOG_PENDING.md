### SDK Features

### SDK Enhancements

### SDK Bugs
* `private/mode/api`: Fix idempotency members not to require validation [#2353](https://github.com/aws/aws-sdk-go/pull/2353)
  * Fixes the SDK's usage of API operation request members marked as idempotency tokens to not require validation. These fields will be auto populated by the SDK if the user does not provide a value. The SDK was requiring the user to provide a value or disable validation to use these APIs.
* deps: Update Go Deps lock file to correct tracking hash [#2354](https://github.com/aws/aws-sdk-go/pull/2354)
