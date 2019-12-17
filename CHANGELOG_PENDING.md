### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/session`: Fix client init not exposing endpoint resolve error ([#3059](https://github.com/aws/aws-sdk-go/pull/3059))
  * Fixes the SDK API clients not surfacing endpoint resolution errors, when the EndpointResolver is unable to resolve an endpoint for the client and region.
* `aws/EC2Metadata` : Reduces request timeout for EC2Metadata client along with maximum number of retries ([#3028](https://github.com/aws/aws-sdk-go/pull/3028))
  * Reduces latency while fetching response from EC2Metadata client running in a container to around 3 seconds
  * Fixes [#2972](https://github.com/aws/aws-sdk-go/issues/2972)   
