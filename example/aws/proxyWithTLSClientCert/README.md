# Example

Example of using the AWS SDK for Go with an HTTPS_PROXY that requires client
TLS certificates. The example will use the proxy configured via the environment
variable `HTTPS_PROXY` proxy a request for the Amazon S3 `ListBuckets` API
operation call.

The example assumes credentials are provided in the environment, or shared
credentials file `~/.aws/credentials`. The `certificate` and `key` files paths
are required to be specified when the example is run. An certificate authority
(CA) file path can also be optionally specified.

## Usage:

```sh
export HTTPS_PROXY=127.0.0.1:8443
export AWS_REGION=us-west-2
go run -cert <certfile> -key <keyfile> [-ca <cafile>]
```

