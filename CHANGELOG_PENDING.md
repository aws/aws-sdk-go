### SDK Features

### SDK Enhancements
* `service/s3/s3manager`: Clarify documentation and behavior of GetBucketRegion ([#3428](https://github.com/aws/aws-sdk-go/pull/3428))
  * Updates the documentation for GetBucketRegion's behavior with regard to default configuration for path style addressing. Provides examples how to override this behavior.
  * Updates the GetBucketRegion utility to not require a region hint when the session or client was configured with a custom endpoint URL.
  * Related to [#3115](https://github.com/aws/aws-sdk-go/issues/3115)
* `service/s3`: Add failsafe handling for unknown stream messages
  * Adds failsafe handling for receiving unknown stream messages from an API. A `<streamName>UnknownEvent` type will encapsulate the unknown message received from the API. Where `<streamName>` is the name of the API's stream, (e.g. S3's `SelectObjectContentEventStreamUnknownEvent`).

### SDK Bugs

