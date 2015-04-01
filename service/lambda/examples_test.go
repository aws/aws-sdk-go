package lambda_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/lambda"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleLambda_AddEventSource() {
	svc := lambda.New(nil)

	params := &lambda.AddEventSourceInput{
		EventSource:  aws.String("String"),       // Required
		FunctionName: aws.String("FunctionName"), // Required
		Role:         aws.String("RoleArn"),      // Required
		BatchSize:    aws.Long(1),
		Parameters: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.AddEventSource(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_DeleteFunction() {
	svc := lambda.New(nil)

	params := &lambda.DeleteFunctionInput{
		FunctionName: aws.String("FunctionName"), // Required
	}
	resp, err := svc.DeleteFunction(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetEventSource() {
	svc := lambda.New(nil)

	params := &lambda.GetEventSourceInput{
		UUID: aws.String("String"), // Required
	}
	resp, err := svc.GetEventSource(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetFunction() {
	svc := lambda.New(nil)

	params := &lambda.GetFunctionInput{
		FunctionName: aws.String("FunctionName"), // Required
	}
	resp, err := svc.GetFunction(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetFunctionConfiguration() {
	svc := lambda.New(nil)

	params := &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String("FunctionName"), // Required
	}
	resp, err := svc.GetFunctionConfiguration(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_InvokeAsync() {
	svc := lambda.New(nil)

	params := &lambda.InvokeAsyncInput{
		FunctionName: aws.String("FunctionName"),         // Required
		InvokeArgs:   bytes.NewReader([]byte("PAYLOAD")), // Required
	}
	resp, err := svc.InvokeAsync(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_ListEventSources() {
	svc := lambda.New(nil)

	params := &lambda.ListEventSourcesInput{
		EventSourceARN: aws.String("String"),
		FunctionName:   aws.String("FunctionName"),
		Marker:         aws.String("String"),
		MaxItems:       aws.Long(1),
	}
	resp, err := svc.ListEventSources(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_ListFunctions() {
	svc := lambda.New(nil)

	params := &lambda.ListFunctionsInput{
		Marker:   aws.String("String"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListFunctions(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_RemoveEventSource() {
	svc := lambda.New(nil)

	params := &lambda.RemoveEventSourceInput{
		UUID: aws.String("String"), // Required
	}
	resp, err := svc.RemoveEventSource(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_UpdateFunctionConfiguration() {
	svc := lambda.New(nil)

	params := &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String("FunctionName"), // Required
		Description:  aws.String("Description"),
		Handler:      aws.String("Handler"),
		MemorySize:   aws.Long(1),
		Role:         aws.String("RoleArn"),
		Timeout:      aws.Long(1),
	}
	resp, err := svc.UpdateFunctionConfiguration(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_UploadFunction() {
	svc := lambda.New(nil)

	params := &lambda.UploadFunctionInput{
		FunctionName: aws.String("FunctionName"),         // Required
		FunctionZip:  bytes.NewReader([]byte("PAYLOAD")), // Required
		Handler:      aws.String("Handler"),              // Required
		Mode:         aws.String("Mode"),                 // Required
		Role:         aws.String("RoleArn"),              // Required
		Runtime:      aws.String("Runtime"),              // Required
		Description:  aws.String("Description"),
		MemorySize:   aws.Long(1),
		Timeout:      aws.Long(1),
	}
	resp, err := svc.UploadFunction(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}