# Example

concatenateObjects is an example using the AWS SDK for Go to concatenate two objects together.


# Usage

The example uses the the bucket name provided, two keys for each object, and lastly the output key.

```sh
AWS_REGION=<region> go run concatenateObjects.go <bucket> <key for object 1> <key for object 2> <key for output>
```
