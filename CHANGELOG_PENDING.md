### SDK Features

### SDK Enhancements
* `aws/credentials`: Clarify `token` usage in `NewStaticCredentials` documentation.
  * Related to [#3162](https://github.com/aws/aws-sdk-go/issues/3162).
* `service/s3/s3manager`: Improve memory allocation behavior by replacing sync.Pool with custom pool implementation ([#3183](https://github.com/aws/aws-sdk-go/pull/3183))
  * Improves memory allocations that occur when the provided `io.Reader` to upload does not satisfy both the `io.ReaderAt` and `io.ReadSeeker` interfaces.
  * Fixes [#3075](https://github.com/aws/aws-sdk-go/issues/3075)

### SDK Bugs
