### SDK Features

### SDK Enhancements
* `private/model/api`: Skip unsupported API models during code generation ([#2849](https://github.com/aws/aws-sdk-go/pull/2849))
  * Adds support for removing API modeled operations that use unsupported features. If a API model results in having no operations it will be skipped.

### SDK Bugs
* `private/model/api` : Fixes broken test for code generation example ([#2855](https://github.com/aws/aws-sdk-go/pull/2855))
    
