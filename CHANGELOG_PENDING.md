### SDK Features
* `services/transcribestreamingservice`: Support for Amazon Transcribe Streaming ([#3048](https://github.com/aws/aws-sdk-go/pull/3048))
  * The SDK now supports the Amazon Transcribe Streaming APIs by utilizing event stream encoding over HTTP/2
  * See [Amazon Transcribe Developer Guide](https://docs.aws.amazon.com/transcribe/latest/dg)
  * Fixes [#2487](https://github.com/aws/aws-sdk-go/issues/2487)

### SDK Enhancements

### SDK Bugs
* `aws/session`: Fix client init not exposing endpoint resolve error ([#3059](https://github.com/aws/aws-sdk-go/pull/3059))
  * Fixes the SDK API clients not surfacing endpoint resolution errors, when the EndpointResolver is unable to resolve an endpoint for the client and region.
