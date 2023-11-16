### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/session`: SDK should now trigger a credential refresh for shared config files if we receive an expired token exception.
  * This change updates the SDK to use a `SharedCredentialProvider` instead of a `StaticProvider` so that the credentials file will be loaded again whenever we mark the credential as expired. If you're using the default `AfterRetryHandler`, credentials should be reloaded automatically whenever we receive an `ExpiredToken`, `ExpiredTokenException`, or `RequestExpired` error.
