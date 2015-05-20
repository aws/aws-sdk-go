package lambda_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/lambda"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleLambda_AddPermission() {
	svc := lambda.New(nil)

	params := &lambda.AddPermissionInput{
		Action:        aws.String("Action"),       // Required
		FunctionName:  aws.String("FunctionName"), // Required
		Principal:     aws.String("Principal"),    // Required
		StatementID:   aws.String("StatementId"),  // Required
		SourceARN:     aws.String("Arn"),
		SourceAccount: aws.String("SourceOwner"),
	}
	resp, err := svc.AddPermission(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_CreateEventSourceMapping() {
	svc := lambda.New(nil)

	params := &lambda.CreateEventSourceMappingInput{
		EventSourceARN:   aws.String("Arn"),                 // Required
		FunctionName:     aws.String("FunctionName"),        // Required
		StartingPosition: aws.String("EventSourcePosition"), // Required
		BatchSize:        aws.Long(1),
		Enabled:          aws.Boolean(true),
	}
	resp, err := svc.CreateEventSourceMapping(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_CreateFunction() {
	svc := lambda.New(nil)

	params := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{ // Required
			ZipFile: []byte("PAYLOAD"),
		},
		FunctionName: aws.String("FunctionName"), // Required
		Handler:      aws.String("Handler"),      // Required
		Role:         aws.String("RoleArn"),      // Required
		Runtime:      aws.String("Runtime"),      // Required
		Description:  aws.String("Description"),
		MemorySize:   aws.Long(1),
		Timeout:      aws.Long(1),
	}
	resp, err := svc.CreateFunction(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_DeleteEventSourceMapping() {
	svc := lambda.New(nil)

	params := &lambda.DeleteEventSourceMappingInput{
		UUID: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEventSourceMapping(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
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

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetEventSourceMapping() {
	svc := lambda.New(nil)

	params := &lambda.GetEventSourceMappingInput{
		UUID: aws.String("String"), // Required
	}
	resp, err := svc.GetEventSourceMapping(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
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

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
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

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetPolicy() {
	svc := lambda.New(nil)

	params := &lambda.GetPolicyInput{
		FunctionName: aws.String("FunctionName"), // Required
	}
	resp, err := svc.GetPolicy(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_Invoke() {
	svc := lambda.New(nil)

	params := &lambda.InvokeInput{
		FunctionName:   aws.String("FunctionName"), // Required
		ClientContext:  aws.String("String"),
		InvocationType: aws.String("InvocationType"),
		LogType:        aws.String("LogType"),
		Payload:        []byte("PAYLOAD"),
	}
	resp, err := svc.Invoke(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
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

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_ListEventSourceMappings() {
	svc := lambda.New(nil)

	params := &lambda.ListEventSourceMappingsInput{
		EventSourceARN: aws.String("Arn"),
		FunctionName:   aws.String("FunctionName"),
		Marker:         aws.String("String"),
		MaxItems:       aws.Long(1),
	}
	resp, err := svc.ListEventSourceMappings(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
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

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_RemovePermission() {
	svc := lambda.New(nil)

	params := &lambda.RemovePermissionInput{
		FunctionName: aws.String("FunctionName"), // Required
		StatementID:  aws.String("StatementId"),  // Required
	}
	resp, err := svc.RemovePermission(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_UpdateEventSourceMapping() {
	svc := lambda.New(nil)

	params := &lambda.UpdateEventSourceMappingInput{
		UUID:         aws.String("String"), // Required
		BatchSize:    aws.Long(1),
		Enabled:      aws.Boolean(true),
		FunctionName: aws.String("FunctionName"),
	}
	resp, err := svc.UpdateEventSourceMapping(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_UpdateFunctionCode() {
	svc := lambda.New(nil)

	params := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String("FunctionName"), // Required
		ZipFile:      []byte("PAYLOAD"),          // Required
	}
	resp, err := svc.UpdateFunctionCode(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
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

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}