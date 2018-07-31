### SDK Features

### SDK Enhancements
* `service/dynamodb/dynamodbattribute`: Add support for custom struct tag keys([#2054](https://github.com/aws/aws-sdk-go/pull/2054))
  * Adds support for (un)marshaling Go types using custom struct tag keys. The new `MarshalOptions.TagKey` allows the user to specify the tag key to use when (un)marshaling struct fields.  Adds support for struct tags such as `yaml`, `toml`, etc. Support for these keys are in name only, and require the tag value format and values to be supported by the package's Marshalers.

### SDK Bugs
* `aws/endpoints`: Add workaround for AWS China Application Autoscaling ([#2080](https://github.com/aws/aws-sdk-go/pull/2080))
  * Adds workaround to correct the endpoint for Application Autoscaling running in AWS China. This will allow your application to make API calls to Application Autoscaling service in AWS China.
  * Fixes [#2079](https://github.com/aws/aws-sdk-go/issues/2079)
  * Fixes [#1957](https://github.com/aws/aws-sdk-go/issues/1957)
* `private/protocol/xml/xmlutil`: Fix SDK marshaling of empty types ([#2081](https://github.com/aws/aws-sdk-go/pull/2081))
  * Fixes the SDK's marshaling of types without members. This corrects the issue where the SDK would not marshal an XML tag for a type, if that type did not have any exported members.
  * Fixes [#2015](https://github.com/aws/aws-sdk-go/issues/2015)
