### SDK Features

### SDK Enhancements

### SDK Bugs
* `private/model/api`: Fix SDK's unmarshaling of unmodeled response payload ([#2340](https://github.com/aws/aws-sdk-go/pull/2340))
  * Fixes the SDK's unmarshaling of API operation response payloads for operations that are unmodeled. Prevents the SDK due to unexpected response payloads causing errors in the API protocol unmarshaler.
  * Fixes [#2332](https://github.com/aws/aws-sdk-go/issues/2332)
