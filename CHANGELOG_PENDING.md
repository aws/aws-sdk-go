### SDK Features

### SDK Enhancements
* `aws/credentials`: Add grouping of concurrent refresh of credentials ([#3127](https://github.com/aws/aws-sdk-go/pull/3127/)
  * Concurrent calls to `Credentials.Get` are now grouped in order to prevent numerous synchronous calls to refresh the credentials. Replacing the mutex with a singleflight reduces the overall amount of time request signatures need to wait while retrieving credentials. This is improvement becomes pronounced when many requests are being made concurrently.

### SDK Bugs
