### SDK Features
* Support has been added for configuring endpoints with requirements such as FIPS and DualStack. ([#3938](https://github.com/aws/aws-sdk-go/pull/3938))
  * `AWS_USE_FIPS_ENDPOINT` and `AWS_USE_DUALSTACK_ENDPOINT` can be set to `true` or `false` in the environment to indicate an endpoint with the respective characteristic must be resolved.
  * `use_fips_endpoint` and `use_dualstack_endpoint` can be set to `true` or `false` in the shared config file to indicate an endpoint with the respective characteristic must be resolved.
  * Programmatic configuration of FIPS and DualStack endpoint resolution.
  * For more information see the `aws/session` package documentation.

### SDK Enhancements

### SDK Bugs
