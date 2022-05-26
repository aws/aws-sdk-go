### SDK Features

### SDK Enhancements

### SDK Bugs
* `service/cloudwatchevidently`: Introduces a breaking change for following parameters from a JSONValue to string type, because the SDKs JSONValue is not compatible with the service's request and response shapes.
  * `EvaluateFeatureInput.EvaluationContext` 
  * `EvaluateFeatureOutput.Details`
  * `EvaluationRequest.EvaluationContext`
  * `EvaluationResult.Details`
  * `Event.Data`
  * `ExperimentReport.Content`
  * `MetricDefinition.EventPattern`
  * `MetricDefinitionConfig.EventPattern`
