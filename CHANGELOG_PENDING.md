### SDK Features

### SDK Enhancements

* `aws/ec2metadata`: Added an option to disable fallback to IMDSv1. 
  * When set the SDK will no longer fallback to IMDSv1 when fetching a token fails. Use `aws.WithEC2MetadataDisableFallback` to enable.

### SDK Bugs
