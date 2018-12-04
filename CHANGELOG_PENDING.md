### SDK Features
* `aws/credential`: Add credential_process provider ([#2217](https://github.com/aws/aws-sdk-go/pull/2217))
  * Adds support for the shared configuration file's `credential_process` property. This property allows the application to execute a command in order to retrieve AWS credentials for AWS service API request.  In order to use this feature your application must enable the SDK's support of the shared configuration file. See, https://docs.aws.amazon.com/sdk-for-go/api/aws/session/#hdr-Sessions_from_Shared_Config for more information on enabling shared config support.

### SDK Enhancements
* `service/sqs`: Add batch checksum validation test ([#2307](https://github.com/aws/aws-sdk-go/pull/2307))
  * Adds additional test of the SQS batch checksum validation.

### SDK Bugs
* aws/session: Fix formatting bug in doc. ([#2294](https://github.com/aws/aws-sdk-go/pull/2294))
  * Fixes a minor issue in aws/session/doc.go where mistakenly used format specifiers in logger.Println.
* Fix SDK model cleanup to remove old model folder ([#2324](https://github.com/aws/aws-sdk-go/pull/2324))
  * Fixes the SDK's model cleanup to remove the entire old model folder not just the api-2.json file.
* Fix SDK's vet usage to use go vet with build tags ([#2300](https://github.com/aws/aws-sdk-go/pull/2300))
  * Updates the SDK's usage of vet to use go vet instead of go tool vet. This allows the SDK to pass build tags and packages instead of just folder paths to the tool.
