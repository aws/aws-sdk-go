### SDK Features

### SDK Enhancements
* `aws/signer/v4`: Always sign a request with the current time. ([#2336](https://github.com/aws/aws-sdk-go/pull/2336))
  * Updates the SDK's v4 request signer to always sign requests with the current time. For the first request attempt, the request's creation time was used in the request's signature. In edge cases this allowed the signature to expire before the request was sent if there was significant delay between creating the request and sending it, (e.g. rate limiting).
* `aws/endpoints`: Deprecate endpoint service ID generation. ([#2338](https://github.com/aws/aws-sdk-go/pull/2338))
  * Deprecates the service ID generation. The list of service IDs do not directly 1:1 relate to a AWS service. The set of ServiceIDs is confusing, and inaccurate. Instead users should use the EndpointID value defined in each service client's package

### SDK Bugs
