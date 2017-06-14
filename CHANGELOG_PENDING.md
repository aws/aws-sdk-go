### SDK Features

### SDK Enhancements

### SDK Bugs
* `private/model/api`: Fix RESTXML support for XML Namespace [#1343](https://github.com/aws/aws-sdk-go/pull/1343)
  * Fixes a bug with the SDK's generation of services using the REST XML protocol not annotating shape references with the XML Namespace attribute.
  * Fixes [#1334](https://github.com/aws/aws-sdk-go/pull/1334)
