### SDK Features

### SDK Enhancements
* `private/model/api`: Add EventStream support over RPC protocl ([#1998](https://github.com/aws/aws-sdk-go/pull/1998))
  * Adds support for EventStream over JSON PRC protocol. This adds support for the EventStream's initial-response event, EventStream headers, and EventStream modeled exceptions. Also replaces the hand written tests with generated tests for EventStream usage.

### SDK Bugs
