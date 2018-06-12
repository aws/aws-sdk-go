### SDK Features

### SDK Enhancements
* `service/rds/rdsutils`: Clean up the rdsutils package and adds a new builder to construct connection strings ([#1985](https://github.com/aws/aws-sdk-go/pull/1985))
    * Rewords documentation to be more useful and provides links to prior setup needed to support authentication tokens. Introduces a builder that allows for building connection strings

### SDK Bugs
* `aws/signer/v4`: Fix X-Amz-Content-Sha256 being in to query for presign ([#1976](https://github.com/aws/aws-sdk-go/pull/1976))
    * Fixes the bug which would allow the X-Amz-Content-Sha256 header to be promoted to the query string when presigning a S3 request. This bug also was preventing users from setting their own sha256 value for a presigned URL. Presigned requests generated with the custom sha256 would of always failed with invalid signature.
    * Fixes [#1974](https://github.com/aws/aws-sdk-go/pull/1974)
