### SDK Features

### SDK Enhancements

### SDK Bugs
* `codegen`: Export event stream constructor for easier mocking ([#3473](https://github.com/aws/aws-sdk-go/pull/3473))
  * Fixes [#3412](https://github.com/aws/aws-sdk-go/issues/3412) by exporting the operation's EventStream type's constructor function so it can be used to fully initialize fully when mocking out behavior for API operations with event streams.
