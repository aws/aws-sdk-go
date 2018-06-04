### SDK Features
* Add support for EventStream based APIs (S3 SelectObjectContent) ([#1941](https://github.com/aws/aws-sdk-go/pull/1941))
  * Adds support for EventStream asynchronous APIs such as S3 SelectObjectContents API. This API allows your application to receiving multiple events asynchronously from the API response. Your application recieves these events from a channel on the API response.
  * See PR [#1941](https://github.com/aws/aws-sdk-go/pull/1941) for example.
  * Fixes [#1895](https://github.com/aws/aws-sdk-go/issues/1895)

### SDK Enhancements

### SDK Bugs
