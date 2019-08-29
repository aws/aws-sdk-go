### SDK Features

### SDK Enhancements
* `aws/ec2metadata`: Add marketplaceProductCodes to EC2 Instance Identity Document
  * Adds `MarketplaceProductCodes` to the EC2 Instance Metadata's Identity Document. The ec2metadata client will now retrieve these values if they are available.
  * Fixes [#2781](https://github.com/aws/aws-sdk-go/issues/2781)
* `private/protocol`: Add support for parsing fractional time ([#2760](https://github.com/aws/aws-sdk-go/pull/2760))
  * Fixes the SDK's ability to parse fractional unix timestamp values and added tests.
  * Fixes [#1448](https://github.com/aws/aws-sdk-go/pull/1448)

### SDK Bugs
