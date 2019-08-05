### SDK Features

### SDK Enhancements
* `service/kinesis`: Add support for retrying service specific API errors ([#2751](https://github.com/aws/aws-sdk-go/pull/2751)
  * Adds support for retrying the Kinesis API error, LimitExceededException.
  * Fixes [#1376](https://github.com/aws/aws-sdk-go/issues/1376)
* `aws/credentials/stscreds`: Add STS and Assume Role specific retries ([#2752](https://github.com/aws/aws-sdk-go/pull/2752))
  * Adds retries to specific STS API errors to the STS AssumeRoleWithWebIdentity credential provider, and STS API operations in general.
* `aws/session`: Ignore invalid shared config file when not used ([#2731](https://github.com/aws/aws-sdk-go/pull/2731))
  * Updates the Session to not fail to load when credentials are provided via the environment variable, the AWS_PROFILE/Option.Profile have not been specified, and the shared config has not been enabled.
  * Fix [#2455](https://github.com/aws/aws-sdk-go/issues/2727)

### SDK Bugs
