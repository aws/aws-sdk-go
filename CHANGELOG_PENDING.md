### SDK Features

### SDK Enhancements
* `private/protocol`: Add support for parsing fractional time ([#2760](https://github.com/aws/aws-sdk-go/pull/2760))
  * Fixes the SDK's ability to parse fractional unix timestamp values and added tests.
  * Fixes [#1448](https://github.com/aws/aws-sdk-go/pull/1448)
* `aws/session`: Ignore invalid shared config file when not used ([#2731](https://github.com/aws/aws-sdk-go/pull/2731))
  * Updates the Session to not fail to load when credentials are provided via the environment variable, the AWS_PROFILE/Option.Profile have not been specified, and the shared config has not been enabled.
  * Fixes [#2455](https://github.com/aws/aws-sdk-go/issues/2455)

### SDK Bugs
