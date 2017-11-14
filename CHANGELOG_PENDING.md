### SDK Features

### SDK Enhancements
* `aws/request`: Remove default port from HTTP host header ([#1618](https://github.com/aws/aws-sdk-go/pull/1618))
  * Updates the SDK to automatically remove default ports based on the URL's scheme when setting the HTTP Host header's value.
  * Fixes [#1537](https://github.com/aws/aws-sdk-go/issues/1537)

### SDK Bugs
