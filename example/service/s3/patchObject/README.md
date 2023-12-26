# Example

This is an example using the AWS SDK for Go to patch object.


# Usage

The example uses the bucket name provided, one key for object, uploads `base_object.txt` file to S3, then edits
this file, replacing it's 14th to 23rd bytes with `"a"` characters.

```sh
go run -tags example patchObject.go <bucket> <key for object>
```
