### SDK Features
* `aws/request`: Add support for context.Context to SDK API operation requests (#1132)
  * Adds support for context.Context to the SDK by adding `WithContext` methods for each API operation, Paginators and Waiters. e.g `PutObjectWithContext`. This change also adds the ability to provide request functional options to the method calls instead of requiring you to use the `Request` API operation method (e.g `PutObjectRequest`).
  * Adds a `Complete` Request handler list that will be called ever time a request is completed. This includes both success and failure. Complete will only be called once per API operation request.
  * `private/waiter` package moved from the private group to `aws/request/waiter` and made publicly available.
  * Adds Context support to all API operations, Waiters(WaitUntil) and Paginators(Pages) methods.
  * Adds Context support for s3manager and s3crypto clients.

### SDK Enhancements
* `aws/signer/v4`: Adds support for unsigned payload signer config (#1130)
  * Adds configuration option to the v4.Signer to specify the request's body should not be signed. This will only correclty function on services that support unsigned payload. e.g. S3, Glacier. 

### SDK Bug Fixes
* `service/s3`: Fix S3 HostID to be available in S3 request error message (#1131)
  * Adds a new type s3.RequestFailure which exposes the S3 HostID value from a S3 API operation response. This is helpful when you have an error with S3, and need to contact support. Both RequestID and HostID are needed.
* `private/model/api`: Do not return a link if uid is empty (#1133)
  * Fixes SDK's doc generation to not generate API reference doc links if the SDK us unable to create a valid link.
* `aws/request`: Optimization to handler list copy to prevent multiple alloc calls. (#1134)
  * Removes unneeded allocations when making API operation requests.
