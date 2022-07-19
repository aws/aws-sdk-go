### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/pricing`: Fixes a bug that caused `GetProductsOutput.PriceList` to be generated incorrectly. ([#4486](https://github.com/aws/aws-sdk-go/pull/4486))
  * The [v1.44.46](https://github.com/aws/aws-sdk-go/releases/tag/v1.44.46) release incorrectly resulted in the `PriceList` field's type changing from `[]aws.JSONValue` to `[]*string`.
  * This release reverts this change, with the field now correctly updated to `[]aws.JSONValue`.
  * Fixes [#4480](https://github.com/aws/aws-sdk-go/issues/4480)
