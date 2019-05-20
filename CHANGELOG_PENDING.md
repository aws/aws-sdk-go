### SDK Features

### SDK Enhancements
* Add raw error message bytes to SerializationError errors ([#2600](https://github.com/aws/aws-sdk-go/pull/2600))
  * Updates the SDK's API error message SerializationError handling to capture the original error message byte, and include it in the SerializationError error value.
  * Fixes [#2562](https://github.com/aws/aws-sdk-go/issues/2562), [#2411](https://github.com/aws/aws-sdk-go/issues/2411), [#2315](https://github.com/aws/aws-sdk-go/issues/2315)

### SDK Bugs
* service/s3/s3manager: Fix uploader to check for empty part before max parts check (#2556)
  * Fixes the S3 Upload manager's behavior for uploading exactly MaxUploadParts * PartSize to S3. The uploader would previously return an error after the full content was uploaded, because the assert on max upload parts was occurring before the check if there were any more parts to upload.
  * Fixes [#2557](https://github.com/aws/aws-sdk-go/issues/2557)
