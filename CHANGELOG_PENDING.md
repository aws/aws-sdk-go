### SDK Features
* `aws/session`: Support has been added for EC2 IPv6-enabled Instance Metadata Service Endpoints ([#4006](https://github.com/aws/aws-sdk-go/pull/4006))
  * Adds support for `AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE`  environment variable which may specify `IPv6` or `IPv4` for selecting the desired endpoint.
  * Adds support for `ec2_metadata_service_endpoint_mode` AWS profile key, which may specify `IPv6` or `IPv4` for selecting the desired endpoint. Has lower precedence then `AWS_EC2_METADATA_SERVICE_ENDPOINT`.
  * Adds support for `ec2_metadata_service_endpoint` AWS profile key, which may specify an explicit endpoint URI. Has higher precedence then `ec2_metadata_service_endpoint_mode`.
* `aws/endpoints`: Supported has been added for EC2 IPv6-enabled Instance Metadata Service Endpoints ([#4006](https://github.com/aws/aws-sdk-go/pull/4006))

### SDK Enhancements

### SDK Bugs
