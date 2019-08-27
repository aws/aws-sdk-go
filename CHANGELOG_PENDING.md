### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/ec2`: Fix int overflow in minTime on 386 and arm ([#2787](https://github.com/aws/aws-sdk-go/pull/2787))
  * Fixes [2786](https://github.com/aws/aws-sdk-go/issues/2786) int overflow issue on 32-bit platforms like 386 and arm.
