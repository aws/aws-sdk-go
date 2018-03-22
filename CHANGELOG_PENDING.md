### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/endpoints`: Use service metadata for fallback signing name ([#1854](https://github.com/aws/aws-sdk-go/pull/1854))
  * Updates the SDK's endpoint resolution to fallback deriving the service's signing name from the service's modeled metadata in addition the the endpoints modeled data.
  * Fixes [#1850](https://github.com/aws/aws-sdk-go/issues/1850)
