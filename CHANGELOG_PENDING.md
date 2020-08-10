### SDK Features

### SDK Enhancements
* `example/service/s3/putObjectWithProgress`: Fix example for file upload with progress ([#3377](https://github.com/aws/aws-sdk-go/pull/3377))
  * Fixes [#2468](https://github.com/aws/aws-sdk-go/issues/2468) by ignoring the first read of the progress reader wrapper. Since the first read is used for signing the request, not upload progress.
  * Updated the example to write progress inline instead of newlines.
* `service/dynamodb/dynamodbattribute`: Fix typo in package docs ([#3446](https://github.com/aws/aws-sdk-go/pull/3446))
  * Fixes typo in dynamodbattribute package docs.

### SDK Bugs
