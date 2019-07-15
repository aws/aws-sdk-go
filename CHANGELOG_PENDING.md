### SDK Features
* `aws/session`: Add support for assuming role via Web Identity Token ([#2667](https://github.com/aws/aws-sdk-go/pull/2667))
  * Adds support for assuming an role via the Web Identity Token. Allows for OIDC token files to be used by specifying the token path through the AWS_WEB_IDENTITY_TOKEN_FILE, and AWS_ROLE_ARN environment variables.

### SDK Enhancements

### SDK Bugs
* `aws/session`: Fix SDK AWS_PROFILE and static environment credential behavior ()
  * Fixes the SDK's behavior when determining the source of credentials to load. Previously the SDK would ignore the AWS_PROFILE environment, if static environment credentials were also specified.
  * If both AWS_PROFILE and static environment credentials are defined, the SDK will load any credentials from the shared config/credentials file for the AWS_PROFILE first. Only if there are no credentials defined in the shared config/credentials file will the SDK use the static environment credentials instead.
