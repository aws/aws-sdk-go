### SDK Features
* Remove SDK's `vendor` directory of vendored dependencies
  * Updates the SDK's Go module definition to enumerate all dependencies of the SDK and its components. 
  * SDK's repository root package has been updated to refer to runtime dependencies like `go-jmespath` for `go get` the SDK with Go without modules.
* Deletes the deprecated `awsmigrate` utility from the SDK's repository.
  * This utility is no longer relevant. The utility allowed users the beta pre-release v0 SDK to update to the v1.0 released version of the SDK.

### SDK Enhancements

### SDK Bugs
