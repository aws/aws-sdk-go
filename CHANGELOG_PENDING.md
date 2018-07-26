### SDK Features

### SDK Enhancements
* `service/s3/s3manager`: Document default behavior for Upload's MaxNumParts ([#2077](https://github.com/aws/aws-sdk-go/issues/2077))
  * Updates the S3 Upload Manager's default behavior for MaxNumParts, and ensures that the Uploader.MaxNumPart's member value is initialized properly if the type was created via struct initialization instead of using the NewUploader function.
  * Fixes [#2015](https://github.com/aws/aws-sdk-go/issues/2015)

### SDK Bugs
