### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws`: Fix SDK's suppressing of sensitive API parameters being logged.
  * The SDK did not correctly suppress sensitive API parameters via the `String` and `GoString` methods. Updates the SDK's behavior to suppress sensitive API parameters.
