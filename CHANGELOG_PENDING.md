### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/EC2Metadata` : Reduces request timeout for EC2Metadata client along with maximum number of retries ([#3028](https://github.com/aws/aws-sdk-go/pull/3028))
  * Reduces latency while fetching response from EC2Metadata client running in a container to around 3 seconds 
