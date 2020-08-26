### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/ec2metadata`: Add support for EC2 IMDS endpoint from environment variable ([#3504](https://github.com/aws/aws-sdk-go/pull/3504))
  * Adds support for specifying a custom EC2 IMDS endpoint from the environment variable, `AWS_EC2_METADATA_SERVICE_ENDPOINT`.
  * The `aws/session#Options` struct also has a new field, `EC2IMDSEndpoint`. This field can be used to configure the custom endpoint of the EC2 IMDS client. The option only applies to EC2 IMDS clients created after the Session with `EC2IMDSEndpoint` is specified.
