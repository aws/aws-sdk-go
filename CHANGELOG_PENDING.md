### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/session`: Fix client init not exposing endpoint resolve error ([#3059](https://github.com/aws/aws-sdk-go/pull/3059))
  * Fixes the SDK API clients not surfacing endpoint resolution errors, when the EndpointResolver is unable to resolve an endpoint for the client and region.
