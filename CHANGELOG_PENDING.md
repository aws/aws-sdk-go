### SDK Features

### SDK Enhancements
* `private/model/api`: Add validation to ensure there is no duplication of services in models/apis ([#1758](https://github.com/aws/aws-sdk-go/pull/1758))
    * Prevents the SDK from mistakenly generating code a single service multiple times with different model versions.
* `example/service/ec2/instancesbyRegion`: Fix typos in example ([#1762](https://github.com/aws/aws-sdk-go/pull/1762))
* `private/model/api`: removing SDK API reference crosslinks from input/output shapes. (#1765) 

### SDK Bugs
* `aws/session`: Fix bug in session.New not supporting AWS_SDK_LOAD_CONFIG ([#1770](https://github.com/aws/aws-sdk-go/pull/1770))
    * Fixes a bug in the session.New function that was not correctly sourcing the shared configuration files' path.
    * Fixes [#1771](https://github.com/aws/aws-sdk-go/pull/1771)
