### SDK Features
* `aws/session`: Corrected order of SDK environment and shared config loading.
  * Environment credentials have precedence over shared config credentials even if the AWS_PROFILE environment credentials are present. The session.Options.Profile value needs to be used to specify a profile for shared config to have precedence over environment credentials. #2694 incorrectly gave AWS_PROFILE for shared config precedence over environment credentials as well.

### SDK Enhancements

### SDK Bugs
* `aws/session`: Fix credential loading order for env and shared config ([#2729](https://github.com/aws/aws-sdk-go/pull/2729))
  * Fixes the credential loading order for environment credentials, when the presence of an AWS_PROFILE value is also provided. The environment credentials have precedence over the AWS_PROFILE.
  * Fixes [#2727](https://github.com/aws/aws-sdk-go/issues/2727)
