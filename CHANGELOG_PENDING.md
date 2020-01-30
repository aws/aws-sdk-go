### SDK Features

### SDK Enhancements
* `aws/request`: Add support for EC2 specific throttle exception code
  * Adds support for the EC2ThrottledException throttling exception code. The SDK will now treat this error code as throttling.

### SDK Bugs
* `aws/request`: Fixes an issue where the HTTP host header did not reflect changes to the endpoint URL ([#3102](https://github.com/aws/aws-sdk-go/pull/3102))
  * Fixes [#3093](https://github.com/aws/aws-sdk-go/issues/3093)
