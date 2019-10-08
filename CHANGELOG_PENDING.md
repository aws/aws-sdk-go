### SDK Features

### SDK Enhancements
* `aws/endpoints`: Adds support for STS Regional Flags ([#2779](https://github.com/aws/aws-sdk-go/pull/2779))
  * STS will default to use the STS global endpoint for most regions when STS regional flag is not set to `regional`. 
  * The regions that default to the STS global endpoint are: 
    * ap-northeast-1
    * ap-south-1
    * ap-southeast-1
    * ap-southeast-2
    * aws-global
    * ca-central-1
    * eu-central-1
    * eu-north-1
    * eu-west-1
    * eu-west-2
    * eu-west-3
    * sa-east-1
    * us-east-1
    * us-east-2
    * us-west-1
    * us-west-2
  * Adds unit test cases for STS Regional flags to test STS endpoint resolver when flag is set to `legacy` or `regional`. Also adds test to verify behavior when STS Regional Flag is not set.
  * Adds error handling when loading environment config from sharedConfig or envConfig. 

### SDK Bugs
