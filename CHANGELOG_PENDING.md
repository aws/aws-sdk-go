### SDK Features

### SDK Enhancements
* `private/protocol`: Add support for parsing fractional time ([#2760](https://github.com/aws/aws-sdk-go/pull/2760))
  * Fixes the SDK's ability to parse fractional unix timestamp values and added tests.
  * Fixes [#1448](https://github.com/aws/aws-sdk-go/pull/1448)
* `aws/session`: Add support for CSM options from shared config file ([#2768](https://github.com/aws/aws-sdk-go/pull/2768))
  * Adds support for enabling and controlling the Client Side Metrics (CSM) reporting from the shared configuration files in addition to the environment variables.

### SDK Bugs
* `service/s3/s3crypto`: Fix tmp file not being deleted after upload ([#2776](https://github.com/aws/aws-sdk-go/pull/2776))
  * Fixes the s3crypto's getWriterStore utiliy's send handler not cleaning up the temporary file after Send completes.
* `private/protocol`: Add protocol tests for blob types and headers ([#2770](https://github.com/aws/aws-sdk-go/pull/2770))
  * Adds RESTJSON and RESTXML protocol tests for blob headers.
  * Related to [#750](https://github.com/aws/aws-sdk-go/issues/750)
* `service/dynamodb/expression`: Improved reporting of bad key conditions ([#2775](https://github.com/aws/aws-sdk-go/pull/2775))
  * Improved error reporting when invalid key conditions are constructed using KeyConditionBuilder
