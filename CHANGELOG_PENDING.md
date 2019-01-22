### SDK Features

### SDK Enhancements
* `aws`: Update Context to be an alias of context.Context for Go 1.9 ([#2412](https://github.com/aws/aws-sdk-go/pull/2412))
  * Updates aws.Context interface to be an alias of the standard libraries context.Context type instead of redefining the interface. This will allow IDEs and utilities to interpret the aws.Context as the exactly same type as the standard libraries context.Context.

### SDK Bugs
