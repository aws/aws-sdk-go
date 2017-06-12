### SDK Features

### SDK Enhancements

### SDK Bugs
* `aws/request`: Fix NewErrParamMinLen to use correct ParamMinLenErrCode [#1336](https://github.com/aws/aws-sdk-go/issues/1336)
  * Fixes the `NewErrParamMinLen` function returning the wrong error code. `ParamMinLenErrCode` should be returned not `ParamMinValueErrCode`.
  * Fixes [#1335](https://github.com/aws/aws-sdk-go/issues/1335)
