### SDK Features
* `service/s3/s3manager`: Adds batch operations to s3manager [#1333](https://github.com/aws/aws-sdk-go/pull/1333)
  * Allows for batch upload, download, and delete of objects. Also adds the interface pattern to allow for easy traversal of objects. E.G `DownloadWithIterator`, `UploadWithIterator`, and `BatchDelete`. `BatchDelete` also contains a utility iterator using the `ListObjects` API to easily delete a list of objects.
  
### SDK Enhancements

### SDK Bugs
