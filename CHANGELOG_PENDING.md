### SDK Features
* `service/s3/s3manager`: Update S3 Upload Multipart location ([#2453](https://github.com/aws/aws-sdk-go/pull/2453))
  * Updates the Location returned value of S3 Upload's Multipart UploadOutput type to be consistent with single part upload URL. This update also brings the multipart upload Location inline with the S3 object URLs created by the SDK
  * Fix [#1385](https://github.com/aws/aws-sdk-go/issues/1385)

### SDK Enhancements
* `service/s3`: Update BucketRegionError message to include more information ([#2451](https://github.com/aws/aws-sdk-go/pull/2451))
  * Updates the BucketRegionError error message to include information about the endpoint and actual region the bucket is in if known. This error message is created by the SDK, but could produce a confusing error message if the user provided a region that doesn't match the endpoint.
  * Fix [#2426](https://github.com/aws/aws-sdk-go/pull/2451)

### SDK Bugs
