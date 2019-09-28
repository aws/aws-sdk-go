### SDK Features

### SDK Enhancements
* `private/protocol/xml/xmlutil`: Support for sorting xml attributes ([#2854](https://github.com/aws/aws-sdk-go/pull/2854))

### SDK Bugs
* ` internal/ini`: Fix ini parser to handle empty values [#2860](https://github.com/aws/aws-sdk-go/pull/2860)
  * Fixes ini parser to correctly handle empty values in the ini file.
  * Adds tests for nested and empty field value parsing, along with tests suggested in [#2801](https://github.com/aws/aws-sdk-go/pull/2801)
  * Fixes [#2800](https://github.com/aws/aws-sdk-go/issues/2800)
* `private/model/api`: Write locationName for top-level shapes for rest-xml and ec2 protocols ([#2854](https://github.com/aws/aws-sdk-go/pull/2854))
* `private/mode/api`: Colliding fields should serialize with original name ([#2854](https://github.com/aws/aws-sdk-go/pull/2854))
  * Fixes [#2806](https://github.com/aws/aws-sdk-go/issues/2806) 
