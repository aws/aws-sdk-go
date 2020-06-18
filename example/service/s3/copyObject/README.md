# Example

This is an example of copying objects of arbitrary size cross-region and cross-account
using the `github.com/aws/aws-sdk-go/service/s3/s3manager` package.

# Usage

Get the full list of parameters with:

```sh
go run -tags example ./example/service/s3/copyObject/ -help
```

Minimal example:

```sh
go run -tags example ./example/service/s3/copyObject/ \
    -copy-source /source-bucket/the/key.txt \
    -dst-bucket destination-bucket \
    -dst-key /the/key.txt
```
