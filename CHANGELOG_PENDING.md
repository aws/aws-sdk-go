### SDK Features

### SDK Enhancements
* `aws/credentials/stscreds`: Add support for policy ARNs ([#3249](https://github.com/aws/aws-sdk-go/pull/3249))
  * Adds support for passing AWS policy ARNs to the `AssumeRoleProvider` and `WebIdentityRoleProvider` credential providers. This allows you provide policy ARNs when assuming the role that will further limit the permissions of the credentials returned.

### SDK Bugs
