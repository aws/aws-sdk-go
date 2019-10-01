### SDK Features

### SDK Enhancements
* `service/s3/s3manager`: Allow reuse of Uploader buffer `sync.Pool` amongst multiple Upload calls ([#2863](https://github.com/aws/aws-sdk-go/pull/2863)) 
  * The `sync.Pool` used for the reuse of `[]byte` slices when handling streaming payloads will now be shared across multiple Upload calls when the upload part size remains constant. 

### SDK Bugs
* `internal/ini`: Fix ini parser to handle empty values [#2860](https://github.com/aws/aws-sdk-go/pull/2860)
  * Fixes incorrect modifications to the previous token value of the skipper. Adds checks for cases where a skipped statement should be marked as complete and not be ignored.
  * Adds tests for nested and empty field value parsing, along with tests suggested in [#2801](https://github.com/aws/aws-sdk-go/pull/2801)
  * Fixes [#2800](https://github.com/aws/aws-sdk-go/issues/2800)
