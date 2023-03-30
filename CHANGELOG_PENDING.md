### SDK Features

### SDK Enhancements

### SDK Bugs

* `aws/session`: Fix `AWS_USE_FIPS_ENDPOINT` not being inferred on resolved credentials.
  * Defer resolving default credentials chain until after other config is resolved.
