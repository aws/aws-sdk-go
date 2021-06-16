### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/ec2metadata`: Fix client retrying 404 responses ([#3962](https://github.com/aws/aws-sdk-go/pull/3962))
  * Fixes the EC2 IMDS client to not retry 404 HTTP errors received for operations like GetMetadata.
