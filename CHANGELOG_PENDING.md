### SDK Features

### SDK Enhancements
* `aws/client`: Add PartitionID to Config ([#2902](https://github.com/aws/aws-sdk-go/pull/2902))
* `aws/client/metadata`: Add PartitionID to ClientInfo ([#2902](https://github.com/aws/aws-sdk-go/pull/2902))
* `aws/endpoints`: Add PartitionID to ResolvedEndpoint ([#2902](https://github.com/aws/aws-sdk-go/pull/2902))

### SDK Bugs
* `aws/endpoints`: Fix resolve endpoint with empty region ([#2911](https://github.com/aws/aws-sdk-go/pull/2911))
  * Fixes the SDK's behavior when attempting to resolve a service's endpoint when no region was provided. Adds legacy support for services that were able to resolve a valid endpoint. No new service will support resolving an endpoint without an region.
  * Fixes [#2909](https://github.com/aws/aws-sdk-go/issues/2909)
