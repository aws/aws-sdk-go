### SDK Features

### SDK Enhancements
* `aws`: Add `WithLowerCaseHeaderMaps` and `WithDisableRestProtocolURICleaning` to `aws.Config`. ([#3671](https://github.com/aws/aws-sdk-go/pull/3671))

### SDK Bugs
* `aws`: Fixed a case where `LowerCaseHeaderMaps` would not be merged when merging`aws.Config` types. ([#3671](https://github.com/aws/aws-sdk-go/pull/3671))
