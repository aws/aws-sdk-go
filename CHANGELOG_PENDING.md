### SDK Features

### SDK Enhancements
* `aws/client`: Adds configurations to the default retryer ([#2830](https://github.com/aws/aws-sdk-go/pull/2830))
  * Exposes members of the default retryer. Adds NoOpRetryer to support no retry behavior. 
  * Updates the underlying logic used by the default retryer to calculate jittered delay for retry. 
  * Fixes [#2829](https://github.com/aws/aws-sdk-go/issues/2829)

### SDK Bugs
