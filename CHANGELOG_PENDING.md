### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/credentials`: Monotonic clock readings will now be cleared when setting credential expiry time. ([#3573](https://github.com/aws/aws-sdk-go/pull/3573))
  * Prevents potential issues when the host system is hibernated / slept and the monotonic clock readings don't match the wall-clock time.
