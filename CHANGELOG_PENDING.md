### SDK Features

* `aws/session`: Add support for chaining assume IAM role from shared config ([#2579](https://github.com/aws/aws-sdk-go/pull/2579))
  * Adds support chaining assume role credentials from the shared config/credentials files. This change allows you to create an assume role chain of multiple levels of assumed IAM roles. The config profile the deepest in the chain must use static credentials, or `credential_source`. If the deepest profile doesn't have either of these the session will fail to load.
  * Fixes the SDK's shared config credential source not assuming a role with environment and ECS credentials. EC2 credentials were already supported.
  * Fix [#2528](https://github.com/aws/aws-sdk-go/issue/2528)
  * Fix [#2385](https://github.com/aws/aws-sdk-go/issue/2385)

### SDK Enhancements

### SDK Bugs
