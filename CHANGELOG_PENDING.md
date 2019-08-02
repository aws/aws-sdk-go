### SDK Features

### SDK Enhancements
* `aws/endpoints`: Expose DNSSuffix for partitions ([#2711](https://github.com/aws/aws-sdk-go/pull/2711))
  * Exposes the underlying partition metadata's DNSSuffix value via the `DNSSuffix` method on the endpoint's `Partition` type. This allows access to the partition's DNS suffix, e.g. "amazon.com".
  * Fixes [#2710](https://github.com/aws/aws-sdk-go/issues/2710)

### SDK Bugs
