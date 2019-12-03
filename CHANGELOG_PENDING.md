### SDK Features

### SDK Enhancements
* `service/s3`: Add support for Access Point resources
  * Adds support for using Access Point resource with Amazon S3 API operation calls. The Access Point resource are identified by an Amazon Resource Name (ARN).
  * To make operation calls to an S3 Access Point instead of a S3 Bucket, provide the Access Point ARN string as the value of the Bucket parameter. You can create an Access Point for your bucket with the Amazon S3 Control API. The Access Point ARN can be obtained from the S3 Control API. You should avoid building the ARN directly.

### SDK Bugs
