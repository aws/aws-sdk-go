### SDK Features

### SDK Enhancements
* `aws/credentials/stscreds`: Add optional expiry duration to WebIdentityRoleProvider ([#3356](https://github.com/aws/aws-sdk-go/pull/3356))
  * Adds a new optional field to the WebIdentityRoleProvider that allows you to specify the duration the assumed role credentials will be valid for.
* `example/service/s3/putObjectWithProgress`: Fix example for file upload with progress ([#3377](https://github.com/aws/aws-sdk-go/pull/3377))
  * Fixes [#2468](https://github.com/aws/aws-sdk-go/issues/2468) by ignoring the first read of the progress reader wrapper. Since the first read is used for signing the request, not upload progress.
  * Updated the example to write progress inline instead of newlines.
* `service/dynamodb/dynamodbattribute`: Fix typo in package docs ([#3446](https://github.com/aws/aws-sdk-go/pull/3446))
  * Fixes typo in dynamodbattribute package docs.

### SDK Bugs
* `codegen`: Export event stream constructor for easier mocking ([#3473](https://github.com/aws/aws-sdk-go/pull/3473))
  * Fixes [#3412](https://github.com/aws/aws-sdk-go/issues/3412) by exporting the operation's EventStream type's constructor function so it can be used to fully initialize fully when mocking out behavior for API operations with event streams.
