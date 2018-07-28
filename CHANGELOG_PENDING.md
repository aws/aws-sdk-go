### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/endpoints`: Add workaround for AWS China Application Autoscaling ([#2080](https://github.com/aws/aws-sdk-go/pull/2080))
  * Adds workaround to correct the endpoint for Application Autoscaling running in AWS China. This will allow your application to make API calls to Application Autoscaling service in AWS China.
  * Fixes [#2079](https://github.com/aws/aws-sdk-go/issues/2079)
  * Fixes [#1957](https://github.com/aws/aws-sdk-go/issues/1957)
