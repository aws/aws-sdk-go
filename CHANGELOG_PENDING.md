### SDK Features

### SDK Enhancements
* `service/s3`: Add utilities to make getting a bucket's region easier (#1207)
  * Adds two features which make it easier to get a bucket's region, `s3.NormalizeBucketLocation` and `s3manager.GetBucketRegion`.

### SDK Bugs
* `service/s3`: Fix HeadObject's incorrect documented error codes (#1213)
  * The HeadObject's model incorrectly states that the operation can return the NoSuchKey error code.
  * Fixes #1208
