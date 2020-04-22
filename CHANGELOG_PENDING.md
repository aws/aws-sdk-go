### SDK Features

### SDK Enhancements
* `aws/credentials/stscreds`: Add support for custom web identity TokenFetcher ([#3256](https://github.com/aws/aws-sdk-go/pull/3256))
  * Adds new constructor, `NewWebIdentityRoleProviderWithToken` for `WebIdentityRoleProvider` which takes a `TokenFetcher`. Implement `TokenFetcher` to provide custom sources for web identity tokens. The `TokenFetcher` must be concurrency safe. `TokenFetcher` may return unique value each time it is called.

### SDK Bugs
