### SDK Features

### SDK Enhancements
* `aws/session`: Add Assume role for credential process from aws shared config ([#2674](https://github.com/aws/aws-sdk-go/pull/2674))
  * Adds support for assuming role using credential process from the shared config file. Also updated SDK's environment testing and added SDK's CI testing with Windows.
* `aws/csm`: Add support for AWS_CSM_HOST env option ([#2677](https://github.com/aws/aws-sdk-go/pull/2677))
  * Adds support for a host to be configured for the SDK's metric reporting Client Side Metrics (CSM) client via the AWS_CSM_HOST environment variable.

### SDK Bugs
