### SDK Features
* Add generated error types for JSONRPC and RESTJSON APIs
  * Adds generated error types for APIs using JSONRPC and RESTJSON protocols. This allows you to retrieve additional error metadata within an error message that was previously unavailable. For example, Amazon DynamoDB's TransactWriteItems operation can return a `TransactionCanceledException` continuing detailed `CancellationReasons` member. This data is now available by type asserting the error returned from the operation call to `TransactionCanceledException` type.

### SDK Enhancements

### SDK Bugs
