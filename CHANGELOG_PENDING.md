### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/request`: Ensure New request handles nil retryer ([#2934](https://github.com/aws/aws-sdk-go/pull/2934))
  * Adds additional default behavior to the SDK's New request constructor, to handle the case where a nil Retryer was passed in. This error could occur when the SDK's Request type was being used to create requests directly, not through one of the SDK's client.
  * Fixes [#2889](https://github.com/aws/aws-sdk-go/issues/2889)
