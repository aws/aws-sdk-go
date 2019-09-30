### SDK Features

### SDK Enhancements
* `service/s3/s3manager`: Allow reuse of Uploader buffer `sync.Pool` amongst multiple Upload calls ([#2863](https://github.com/aws/aws-sdk-go/pull/2863)) 
  * The `sync.Pool` used for the reuse of `[]byte` slices when handling streaming payloads will now be shared across multiple Upload calls when the upload part size remains constant. 

### SDK Bugs
