### SDK Features

### SDK Enhancements

### SDK Bugs
* `codegen`: Fix SDK not draining HTTP response body for operations without modeled output parameters.
  * Fixes a bug in the SDK that was causing HTTP response bodies to not be drained, preventing the connection being reused in the connection pool. For some operations this could contributed to an increase in `close_wait` connection states.
