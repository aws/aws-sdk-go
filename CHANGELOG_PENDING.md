### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/endpoints` Add a workaround for AWS China EC2
* Add a workaround to correct the endpoint for EC2 running in AWS China. This will allow your application to make API calls to EC2 service in AWS China.
* Fix problems like [#2079](https://github.com/aws/aws-sdk-go/issues/2079), [#1957](https://github.com/aws/aws-sdk-go/issues/1957), but the corrected endpoint is `ec2.{region}.amazonaws.com`