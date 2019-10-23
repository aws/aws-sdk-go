### SDK Features

### SDK Enhancements
* `aws/endpoints`: Adds support for STS Regional Flags ([#2779](https://github.com/aws/aws-sdk-go/pull/2779))
  * Implements STS regional flag, with support for `legacy` and `regional` options. Defaults to `legacy`. Legacy, will force all regions specified in aws/endpoints/sts_legacy_regions.go to resolve to the STS global endpoint, sts.amazonaws.com. This is the SDK's current behavior.
  * When the flag's value is `regional` the SDK will resolve the endpoint based on the endpoints.json model. This allows STS to update their service's modeled endpoints to be regionalized for all regions. When `regional` turned on use `aws-global` as the region to use the global endpoint.
  * `AWS_STS_REGIONAL_ENDPOINTS=regional` for environment, or `sts_regional_endpoints=regional` in shared config file.
  * The regions the SDK defaults to the STS global endpoint in `legacy` mode are: 
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

### SDK Bugs
