### SDK Features

### SDK Enhancements

### SDK Bugs 
: Add retry support for RequestTimeoutException (#1158)
  * Adds support for retrying RequestTimeoutException error code that is returned by some services.

### SDK Bugs 
* `private/model/api`: Fix Waiter and Paginators panic on nil param inputs (#1157)
  * Corrects the code generation for Paginators and waiters that caused a panic if nil input parameters were used with the operations.
