### SDK Features

### SDK Enhancements
* `aws/session`: Add support for CSM options from shared config file ([#2768](https://github.com/aws/aws-sdk-go/pull/2768))
  * Adds support for enabling and controlling the Client Side Metrics (CSM) reporting from the shared configuration files in addition to the environment variables.

### SDK Bugs
* `private/protocol`: Add protocol tests for blob types and headers ([#2770](https://github.com/aws/aws-sdk-go/pull/2770))
  * Adds RESTJSON and RESTXML protocol tests for blob headers.
  * Related to [#750](https://github.com/aws/aws-sdk-go/issues/750)

### SDK Bugs
* `service/dynamodb/expression`: Improved reporting of bad key conditions ([#2775](https://github.com/aws/aws-sdk-go/pull/2775))
  * Improved error reporting when invalid key conditions are constructed using KeyConditionBuilder

