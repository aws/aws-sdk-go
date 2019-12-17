### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/ec2metadata` : Reduces request timeout for EC2Metadata client along with maximum number of retries ([#3066](https://github.com/aws/aws-sdk-go/pull/3066))
  * Reduces latency while fetching response from EC2Metadata client running in a container to around 3 seconds
  * Fixes [#2972](https://github.com/aws/aws-sdk-go/issues/2972)
