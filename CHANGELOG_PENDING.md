### SDK Features

### SDK Enhancements
* Add raw error message bytes to SerializationError errors ([#2600](https://github.com/aws/aws-sdk-go/pull/2600))
  * Updates the SDK's API error message SerializationError handling to capture the original error message byte, and include it in the SerializationError error value.
  * Fixes [#2562](https://github.com/aws/aws-sdk-go/issues/2562), [#2411](https://github.com/aws/aws-sdk-go/issues/2411), [#2315](https://github.com/aws/aws-sdk-go/issues/2315)

### SDK Bugs
