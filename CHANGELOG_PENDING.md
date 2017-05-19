### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/request`: Add support for PUT temporary redirects (307) [#1283](https://github.com/aws/aws-sdk-go/issues/1283)
  * Adds support for Go 1.8's GetBody function allowing the SDK's http request using PUT and POST methods to be redirected with temporary redirects with 307 status code.
  * Fixes: [#1267](https://github.com/aws/aws-sdk-go/issues/1267)
* `aws/request`: Add handling for retrying temporary errors during unmarshal [#1289](https://github.com/aws/aws-sdk-go/issues/1289)
  * Adds support for retrying temporary errors that occur during unmarshaling of a request's response body.
  * Fixes: [#1275](https://github.com/aws/aws-sdk-go/issues/1275)
