### SDK Features

### SDK Enhancements
* `aws/endpoints`: Adds support for STS Regional Flags ([#2779](https://github.com/aws/aws-sdk-go/pull/2779))
  * Adds unit test cases for STS Regional flags to test STS endpoint resolver when flag is set to `legacy` or `regional` or is `unset`.
  * Adds error handling when loading environment config from sharedConfig or envConfig. 

### SDK Bugs
