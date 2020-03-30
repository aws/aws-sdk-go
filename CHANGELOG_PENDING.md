### SDK Features

### SDK Enhancements
* Update SDK's `go-jmespath` dependency to latest tagged version `0.3.0` ([#3205](https://github.com/aws/aws-sdk-go/pull/3205))

### SDK Bugs
* Fix generated SDK errors to use pointer receivers
  * Fixes the generated SDK API errors to use pointer function receivers instead of value. This fixes potential confusion writing code and not casting to the correct type. The SDK will always return the API error as a pointer, not value.
  * Code that did type assertions from the operation's returned error to the value type would never be satisfied. Leading to errors being missed. Changing the function receiver to a pointer prevents this error. Highlighting it in code bases.
