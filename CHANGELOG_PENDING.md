### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/client`: Fix clients polluting handler list (#1197)
  * Fixes the clients potentially polluting the passed in handler list with the client's customizations. This change ensures every client always works with a clean copy of the request handlers and it cannot pollute the handlers back upstream.
  * Fixes #1184
* `aws/request`: Fix waiter error match condition (#1195)
  * Fixes the waiters's matching overwriting the request's err, effectively ignoring the error condition. This broke waiters with the FailureWaiterState matcher state.
