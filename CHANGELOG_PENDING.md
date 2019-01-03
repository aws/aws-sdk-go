### SDK Features

### SDK Enhancements
* `aws/credentials`: Add support for getting credential's ExpiresAt. ([#2375](https://github.com/aws/aws-sdk-go/pull/2375))
  * Adds an Expirer interface that Providers can implement, and add a suitable implementation to Expiry class used by most Providers. Add a method on Credentials to get the expiration time of the underlying Provider, if Expirer is supported, without exposing Provider to callers.
  * Fix [#1329](https://github.com/aws/aws-sdk-go/pull/1329)

### SDK Bugs
* `aws/ec2metadata`: bounds check region identifier before split ([#2380](https://github.com/aws/aws-sdk-go/pull/2380))
  * Adds empty response checking to ec2metadata's Region request to prevent a out of bounds panic if empty response received.
* Fix SDK's generated API reference doc page's constants section links ([#2373](https://github.com/aws/aws-sdk-go/pull/2373))
  * Fixes the SDK's generated API reference documentation page's constants section links to to be clickable.
