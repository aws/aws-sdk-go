### SDK Features
* `aws`: Bump minimum go version to 1.19.
  * See https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/.

### SDK Enhancements
* `aws/ec2metadata`: Added environment and shared config support for disabling IMDSv1 fallback.
  * Use env `AWS_EC2_METADATA_V1_DISABLED` or shared config `ec2_metadata_v1_disabled` accordingly.

### SDK Bugs
