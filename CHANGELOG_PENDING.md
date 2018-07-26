### SDK Features
* `private/model/api`: SDK APIs input/output are not consistently generated ([#2073](https://github.com/aws/aws-sdk-go/pull/2073))
  * Updates the SDK's API code generation to generate the API input and output types consistently. This ensures that the SDK will no longer rename input/output types unexpectedly as in [#2070](https://github.com/aws/aws-sdk-go/issues/2070). SDK API input and output parameter types will always be the API name with a suffix of Input and Output.
  * Existing service APIs which were incorrectly modeled have been preserved to ensure they do not break.
  * Fixes [#2070](https://github.com/aws/aws-sdk-go/issues/2070)

### SDK Enhancements
* `service/s3/s3manager`: Document default behavior for Upload's MaxNumParts ([#2077](https://github.com/aws/aws-sdk-go/issues/2077))
  * Updates the S3 Upload Manager's default behavior for MaxNumParts, and ensures that the Uploader.MaxNumPart's member value is initialized properly if the type was created via struct initialization instead of using the NewUploader function.
  * Fixes [#2015](https://github.com/aws/aws-sdk-go/issues/2015)

### SDK Bugs
* `private/model/api`: SDK APIs input/output are not consistently generated ([#2073](https://github.com/aws/aws-sdk-go/pull/2073))
  * Fixes EFS service breaking change in v1.14.26 where `FileSystemDescription` was incorrectly renamed to `UpdateFileSystemOutput.
  * Fixes [#2070](https://github.com/aws/aws-sdk-go/issues/2070)
