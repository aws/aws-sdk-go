### SDK Features

### SDK Enhancements
* `aws/credentials`: `ProviderWithContext` optional interface has been added to support passing contexts on credential retrieval ([#3223](https://github.com/aws/aws-sdk-go/pull/3223))
  * Credential providers that implement the optional `ProviderWithContext` will have context passed to them
  * `ec2rolecreds.EC2RoleProvider`, `endpointcreds.Provider`, `stscreds.AssumeRoleProvider`, `stscreds.WebIdentityRoleProvider` have been updated to support the `ProviderWithContext` interface
  * Fixes [#3213](https://github.com/aws/aws-sdk-go/issues/3213)
* `aws/ec2metadata`: Context aware operations have been added `EC2Metadata` client ([#3223](https://github.com/aws/aws-sdk-go/pull/3223))

### SDK Bugs
