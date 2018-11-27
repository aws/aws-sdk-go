### SDK Features

### SDK Enhancements
* `service/s3/s3manager`: Generate Upload Manager's UploadInput structure ([#2296](https://github.com/aws/aws-sdk-go/pull/2296))
  * Updates the SDK's code generation to also generate the S3 Upload Manager's UploadInput structure type based on the modeled S3 PutObjectInput. This ensures parity between the two types, and the S3 manager does not fall behind the capabilities of PutObject.

### SDK Bugs
* `private/model/api`: Fix model loading to not require docs model. ([#2303](https://github.com/aws/aws-sdk-go/pull/2303))
  * Fixes the SDK's model loading to not require that the docs model be present. This model isn't explicitly required.
* Fixup endpoint discovery unit test to be stable ([#2305](https://github.com/aws/aws-sdk-go/pull/2305))
  * Fixes the SDK's endpoint discovery async unit test to be stable, and produce consistent unit test results.
