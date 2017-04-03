### SDK Features

### SDK Enhancements
* `aws/request`: Improve handler copy, push back, push front performance (#1171)
  * Minor optimization to the handler list's handling of copying and pushing request handlers to the handler list.
* Update codegen header to use Go std wording (#1172)
  * Go recently accepted the proposal for standard generated file header wording in, https://golang.org/s/generatedcode.

### SDK Bugs
* `service/dynamodb`: Fix DynamoDB using custom retryer (#1170)
  * Fixes (#1139) the DynamoDB service client clobbering any custom retryer that was passed into the service client or Session's config.
