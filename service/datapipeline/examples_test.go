package datapipeline_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/datapipeline"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDataPipeline_ActivatePipeline() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ActivatePipelineInput{
		PipelineID: aws.String("id"), // Required
		ParameterValues: []*datapipeline.ParameterValue{
			&datapipeline.ParameterValue{ // Required
				ID:          aws.String("fieldNameString"),  // Required
				StringValue: aws.String("fieldStringValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.ActivatePipeline(params)

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

func ExampleDataPipeline_AddTags() {
	svc := datapipeline.New(nil)

	params := &datapipeline.AddTagsInput{
		PipelineID: aws.String("id"), // Required
		Tags: []*datapipeline.Tag{ // Required
			&datapipeline.Tag{ // Required
				Key:   aws.String("tagKey"),   // Required
				Value: aws.String("tagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTags(params)

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

func ExampleDataPipeline_CreatePipeline() {
	svc := datapipeline.New(nil)

	params := &datapipeline.CreatePipelineInput{
		Name:        aws.String("id"), // Required
		UniqueID:    aws.String("id"), // Required
		Description: aws.String("string"),
		Tags: []*datapipeline.Tag{
			&datapipeline.Tag{ // Required
				Key:   aws.String("tagKey"),   // Required
				Value: aws.String("tagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreatePipeline(params)

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

func ExampleDataPipeline_DeletePipeline() {
	svc := datapipeline.New(nil)

	params := &datapipeline.DeletePipelineInput{
		PipelineID: aws.String("id"), // Required
	}
	resp, err := svc.DeletePipeline(params)

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

func ExampleDataPipeline_DescribeObjects() {
	svc := datapipeline.New(nil)

	params := &datapipeline.DescribeObjectsInput{
		ObjectIDs: []*string{ // Required
			aws.String("id"), // Required
			// More values...
		},
		PipelineID:          aws.String("id"), // Required
		EvaluateExpressions: aws.Boolean(true),
		Marker:              aws.String("string"),
	}
	resp, err := svc.DescribeObjects(params)

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

func ExampleDataPipeline_DescribePipelines() {
	svc := datapipeline.New(nil)

	params := &datapipeline.DescribePipelinesInput{
		PipelineIDs: []*string{ // Required
			aws.String("id"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribePipelines(params)

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

func ExampleDataPipeline_EvaluateExpression() {
	svc := datapipeline.New(nil)

	params := &datapipeline.EvaluateExpressionInput{
		Expression: aws.String("longString"), // Required
		ObjectID:   aws.String("id"),         // Required
		PipelineID: aws.String("id"),         // Required
	}
	resp, err := svc.EvaluateExpression(params)

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

func ExampleDataPipeline_GetPipelineDefinition() {
	svc := datapipeline.New(nil)

	params := &datapipeline.GetPipelineDefinitionInput{
		PipelineID: aws.String("id"), // Required
		Version:    aws.String("string"),
	}
	resp, err := svc.GetPipelineDefinition(params)

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

func ExampleDataPipeline_ListPipelines() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ListPipelinesInput{
		Marker: aws.String("string"),
	}
	resp, err := svc.ListPipelines(params)

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

func ExampleDataPipeline_PollForTask() {
	svc := datapipeline.New(nil)

	params := &datapipeline.PollForTaskInput{
		WorkerGroup: aws.String("string"), // Required
		Hostname:    aws.String("id"),
		InstanceIdentity: &datapipeline.InstanceIdentity{
			Document:  aws.String("string"),
			Signature: aws.String("string"),
		},
	}
	resp, err := svc.PollForTask(params)

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

func ExampleDataPipeline_PutPipelineDefinition() {
	svc := datapipeline.New(nil)

	params := &datapipeline.PutPipelineDefinitionInput{
		PipelineID: aws.String("id"), // Required
		PipelineObjects: []*datapipeline.PipelineObject{ // Required
			&datapipeline.PipelineObject{ // Required
				Fields: []*datapipeline.Field{ // Required
					&datapipeline.Field{ // Required
						Key:         aws.String("fieldNameString"), // Required
						RefValue:    aws.String("fieldNameString"),
						StringValue: aws.String("fieldStringValue"),
					},
					// More values...
				},
				ID:   aws.String("id"), // Required
				Name: aws.String("id"), // Required
			},
			// More values...
		},
		ParameterObjects: []*datapipeline.ParameterObject{
			&datapipeline.ParameterObject{ // Required
				Attributes: []*datapipeline.ParameterAttribute{ // Required
					&datapipeline.ParameterAttribute{ // Required
						Key:         aws.String("attributeNameString"),  // Required
						StringValue: aws.String("attributeValueString"), // Required
					},
					// More values...
				},
				ID: aws.String("fieldNameString"), // Required
			},
			// More values...
		},
		ParameterValues: []*datapipeline.ParameterValue{
			&datapipeline.ParameterValue{ // Required
				ID:          aws.String("fieldNameString"),  // Required
				StringValue: aws.String("fieldStringValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.PutPipelineDefinition(params)

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

func ExampleDataPipeline_QueryObjects() {
	svc := datapipeline.New(nil)

	params := &datapipeline.QueryObjectsInput{
		PipelineID: aws.String("id"),     // Required
		Sphere:     aws.String("string"), // Required
		Limit:      aws.Long(1),
		Marker:     aws.String("string"),
		Query: &datapipeline.Query{
			Selectors: []*datapipeline.Selector{
				&datapipeline.Selector{ // Required
					FieldName: aws.String("string"),
					Operator: &datapipeline.Operator{
						Type: aws.String("OperatorType"),
						Values: []*string{
							aws.String("string"), // Required
							// More values...
						},
					},
				},
				// More values...
			},
		},
	}
	resp, err := svc.QueryObjects(params)

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

func ExampleDataPipeline_RemoveTags() {
	svc := datapipeline.New(nil)

	params := &datapipeline.RemoveTagsInput{
		PipelineID: aws.String("id"), // Required
		TagKeys: []*string{ // Required
			aws.String("string"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTags(params)

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

func ExampleDataPipeline_ReportTaskProgress() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ReportTaskProgressInput{
		TaskID: aws.String("taskId"), // Required
		Fields: []*datapipeline.Field{
			&datapipeline.Field{ // Required
				Key:         aws.String("fieldNameString"), // Required
				RefValue:    aws.String("fieldNameString"),
				StringValue: aws.String("fieldStringValue"),
			},
			// More values...
		},
	}
	resp, err := svc.ReportTaskProgress(params)

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

func ExampleDataPipeline_ReportTaskRunnerHeartbeat() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ReportTaskRunnerHeartbeatInput{
		TaskRunnerID: aws.String("id"), // Required
		Hostname:     aws.String("id"),
		WorkerGroup:  aws.String("string"),
	}
	resp, err := svc.ReportTaskRunnerHeartbeat(params)

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

func ExampleDataPipeline_SetStatus() {
	svc := datapipeline.New(nil)

	params := &datapipeline.SetStatusInput{
		ObjectIDs: []*string{ // Required
			aws.String("id"), // Required
			// More values...
		},
		PipelineID: aws.String("id"),     // Required
		Status:     aws.String("string"), // Required
	}
	resp, err := svc.SetStatus(params)

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

func ExampleDataPipeline_SetTaskStatus() {
	svc := datapipeline.New(nil)

	params := &datapipeline.SetTaskStatusInput{
		TaskID:          aws.String("taskId"),     // Required
		TaskStatus:      aws.String("TaskStatus"), // Required
		ErrorID:         aws.String("string"),
		ErrorMessage:    aws.String("errorMessage"),
		ErrorStackTrace: aws.String("string"),
	}
	resp, err := svc.SetTaskStatus(params)

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

func ExampleDataPipeline_ValidatePipelineDefinition() {
	svc := datapipeline.New(nil)

	params := &datapipeline.ValidatePipelineDefinitionInput{
		PipelineID: aws.String("id"), // Required
		PipelineObjects: []*datapipeline.PipelineObject{ // Required
			&datapipeline.PipelineObject{ // Required
				Fields: []*datapipeline.Field{ // Required
					&datapipeline.Field{ // Required
						Key:         aws.String("fieldNameString"), // Required
						RefValue:    aws.String("fieldNameString"),
						StringValue: aws.String("fieldStringValue"),
					},
					// More values...
				},
				ID:   aws.String("id"), // Required
				Name: aws.String("id"), // Required
			},
			// More values...
		},
		ParameterObjects: []*datapipeline.ParameterObject{
			&datapipeline.ParameterObject{ // Required
				Attributes: []*datapipeline.ParameterAttribute{ // Required
					&datapipeline.ParameterAttribute{ // Required
						Key:         aws.String("attributeNameString"),  // Required
						StringValue: aws.String("attributeValueString"), // Required
					},
					// More values...
				},
				ID: aws.String("fieldNameString"), // Required
			},
			// More values...
		},
		ParameterValues: []*datapipeline.ParameterValue{
			&datapipeline.ParameterValue{ // Required
				ID:          aws.String("fieldNameString"),  // Required
				StringValue: aws.String("fieldStringValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.ValidatePipelineDefinition(params)

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