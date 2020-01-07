### SDK Features

### SDK Enhancements
* `aws`: Add configuration option enable the SDK to unmarshal API response header maps to normalized lower case map keys. ([#3033](https://github.com/aws/aws-sdk-go/pull/3033))
  * Setting `aws.Config.LowerCaseHeaderMaps` to `true` will result in S3's X-Amz-Meta prefixed header to be unmarshaled to lower case Metadata member's map keys.

### SDK Bugs
* `aws/ec2metadata` : Reduces request timeout for EC2Metadata client along with maximum number of retries ([#3066](https://github.com/aws/aws-sdk-go/pull/3066))
  * Reduces latency while fetching response from EC2Metadata client running in a container to around 3 seconds
  * Fixes [#2972](https://github.com/aws/aws-sdk-go/issues/2972)
