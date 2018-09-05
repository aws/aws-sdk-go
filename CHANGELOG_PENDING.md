### SDK Features

### SDK Enhancements
* `private/model/api`: Add "Deprecated" to deprecated API operation and type doc strings ([#2129](https://github.com/aws/aws-sdk-go/pull/2129))
  * Updates the SDK's code generation to include `Deprecated` in the documentation string for API operations and types that are depercated by a service.
  * Related to [golang/go#10909](https://github.com/golang/go/issues/10909)
  * https://blog.golang.org/godoc-documenting-go-code

### SDK Bugs
* `service/s3/s3manager`: Fix Download Manager with iterator docs ([#2131](https://github.com/aws/aws-sdk-go/pull/2131))
  * Fixes the S3 Download manager's DownloadWithIterator documentation example.
  * Fixes [#1824](https://github.com/aws/aws-sdk-go/issues/1824)
