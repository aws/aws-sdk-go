package lambdaiface

import (
	"github.com/awslabs/aws-sdk-go/service/lambda"
)

type LambdaAPI interface {
	AddEventSource(*lambda.AddEventSourceInput) (*lambda.EventSourceConfiguration, error)

	DeleteFunction(*lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)

	GetEventSource(*lambda.GetEventSourceInput) (*lambda.EventSourceConfiguration, error)

	GetFunction(*lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)

	GetFunctionConfiguration(*lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)

	InvokeAsync(*lambda.InvokeAsyncInput) (*lambda.InvokeAsyncOutput, error)

	ListEventSources(*lambda.ListEventSourcesInput) (*lambda.ListEventSourcesOutput, error)

	ListFunctions(*lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)

	RemoveEventSource(*lambda.RemoveEventSourceInput) (*lambda.RemoveEventSourceOutput, error)

	UpdateFunctionConfiguration(*lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)

	UploadFunction(*lambda.UploadFunctionInput) (*lambda.FunctionConfiguration, error)
}