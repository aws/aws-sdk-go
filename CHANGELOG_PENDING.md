### SDK Features

### SDK Enhancements
* `aws/session`: Ignore invalid shared config file when not used ([#2731](https://github.com/aws/aws-sdk-go/pull/2731))
  * Updates the Session to not fail to load when credentials are provided via the environment variable, the AWS_PROFILE/Option.Profile have not been specified, and the shared config has not been enabled.
  * Fix [#2455](https://github.com/aws/aws-sdk-go/issues/2727)

### SDK Bugs
