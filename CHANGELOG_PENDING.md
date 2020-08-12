### SDK Features

### SDK Enhancements
* `codegen`: Add XXX_Values functions for getting slice of API enums by type.
  * Fixes [#3441](https://github.com/aws/aws-sdk-go/issues/3441) by adding a new XXX_Values function for each API enum type that returns a slice of enum values, e.g `DomainStatus_Values`.
* `aws/request`: Update default retry to retry "use of closed network connection" errors ([#3476](https://github.com/aws/aws-sdk-go/pull/3476))
  * Fixes [#3406](https://github.com/aws/aws-sdk-go/issues/3406)

### SDK Bugs
* `private/protocol/json/jsonutil`: Fixes a bug that truncated millisecond precision time in API response to seconds. ([#3474](https://github.com/aws/aws-sdk-go/pull/3474))
  * Fixes [#3464](https://github.com/aws/aws-sdk-go/issues/3464)
  * Fixes [#3410](https://github.com/aws/aws-sdk-go/issues/3410)
* `codegen`: Export event stream constructor for easier mocking ([#3473](https://github.com/aws/aws-sdk-go/pull/3473))
  * Fixes [#3412](https://github.com/aws/aws-sdk-go/issues/3412) by exporting the operation's EventStream type's constructor function so it can be used to fully initialize fully when mocking out behavior for API operations with event streams.
