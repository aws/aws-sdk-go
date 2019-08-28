### SDK Features

### SDK Enhancements
* `aws/ec2metadata`: Add marketplaceProductCodes to EC2 Instance Identity Document
  * Adds `MarketplaceProductCodes` to the EC2 Instance Metadata's Identity Document. The ec2metadata client will now retrieve these values if they are available.
  * Fixes [#2781](https://github.com/aws/aws-sdk-go/issues/2781)

### SDK Bugs
