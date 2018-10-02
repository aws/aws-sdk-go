### SDK Features

### SDK Enhancements
* `aws/config`: fix typo in Config struct documentation (#2169)
  * fix typo in Config struct documentation in aws-sdk-go/aws/config.go
* `internal/csm`: Add region to api call metrics (#2175)
* `private/model/api`: Use modeled service signing version in code generation (#2162)
  * Updates the SDK's code generate to make use of the model's service signature version when generating the client for the service. This allows the SDK to generate a client using the correct signature version, e.g v4 vs s3v4 without the need for additional customizations.

### SDK Bugs
* `service/cloudfront/sign`: Do not Escape HTML when encode the cloudfront sign policy (#2164)
  * Fixes the signer escaping HTML elements `<`, `>`, and `&` in the signature policy incorrectly. Allows use of multiple query parameters in the URL to be signed.
  * Fixes #2163
