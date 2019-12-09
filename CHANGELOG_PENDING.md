### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/s3`: Fix SDK support for Accesspoint ARNs with slash in resource ([#3001](https://github.com/aws/aws-sdk-go/pull/3001))
  * Fixes the SDK's handling of S3 Accesspoint ARNs to correctly parse ARNs with slashes in the resource component as valid. Previously the SDK's ARN parsing incorrectly identify ARN resources with slash delimiters as invalid ARNs.
