### SDK Features

### SDK Enhancements
* `service/kinesis`: Add support for retrying service specific API errors ([#2751](https://github.com/aws/aws-sdk-go/pull/2751)
  * Adds support for retrying the Kinesis API error, LimitExceededException.
  * Fixes [#1376](https://github.com/aws/aws-sdk-go/issues/1376)
* `aws/credentials/stscreds`: Add STS and Assume Role specific retries ([#2752](https://github.com/aws/aws-sdk-go/pull/2752))
  * Adds retries to specific STS API errors to the STS AssumeRoleWithWebIdentity credential provider, and STS API operations in general.

### SDK Bugs
