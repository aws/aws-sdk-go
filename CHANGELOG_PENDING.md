### SDK Features

### SDK Enhancements
* `service/s3`: Update SelectObjectContent doc example to be on the API not nested type. ([#1991](https://github.com/aws/aws-sdk-go/pull/1991))

### SDK Bugs
* `aws/client`: Fix HTTP debug log EventStream payloads ([#2000](https://github.com/aws/aws-sdk-go/pull/2000))
  * Fixes the SDK's HTTP client debug logging to not log the HTTP response body for EventStreams. This prevents the SDK from buffering a very large amount of data to be logged at once. The aws.LogDebugWithEventStreamBody should be used to log the event stream events.
  * Fixes a bug in the SDK's response logger which will buffer the response body's content if LogDebug is enabled but LogDebugWithHTTPBody is not.
