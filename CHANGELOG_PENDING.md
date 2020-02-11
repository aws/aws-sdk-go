### SDK Features
* Remove SDK's `vendor` directory of vendored dependencies
  * Updates the SDK's Go module definition to enumerate all dependencies of the SDK and its components. 
  * SDK's repository root package has been updated to refer to runtime dependencies like `go-jmespath` for `go get` the SDK with Go without modules.
* Deletes the deprecated `awsmigrate` utility from the SDK's repository.
  * This utility is no longer relevant. The utility allowed users the beta pre-release v0 SDK to update to the v1.0 released version of the SDK.

### SDK Enhancements
* `aws/credentials`: Add grouping of concurrent refresh of credentials ([#3127](https://github.com/aws/aws-sdk-go/pull/3127/)
  * Concurrent calls to `Credentials.Get` are now grouped in order to prevent numerous synchronous calls to refresh the credentials. Replacing the mutex with a singleflight reduces the overall amount of time request signatures need to wait while retrieving credentials. This is improvement becomes pronounced when many requests are being made concurrently.

### SDK Bugs
