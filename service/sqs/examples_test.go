package sqs_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/sqs"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSQS_AddPermission() {
	svc := sqs.New(nil)

	params := &sqs.AddPermissionInput{
		AWSAccountIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Actions: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Label:    aws.String("String"), // Required
		QueueURL: aws.String("String"), // Required
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

func ExampleSQS_ChangeMessageVisibility() {
	svc := sqs.New(nil)

	params := &sqs.ChangeMessageVisibilityInput{
		QueueURL:          aws.String("String"), // Required
		ReceiptHandle:     aws.String("String"), // Required
		VisibilityTimeout: aws.Long(1),          // Required
	}
	resp, err := svc.ChangeMessageVisibility(params)

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

func ExampleSQS_ChangeMessageVisibilityBatch() {
	svc := sqs.New(nil)

	params := &sqs.ChangeMessageVisibilityBatchInput{
		Entries: []*sqs.ChangeMessageVisibilityBatchRequestEntry{ // Required
			&sqs.ChangeMessageVisibilityBatchRequestEntry{ // Required
				ID:                aws.String("String"), // Required
				ReceiptHandle:     aws.String("String"), // Required
				VisibilityTimeout: aws.Long(1),
			},
			// More values...
		},
		QueueURL: aws.String("String"), // Required
	}
	resp, err := svc.ChangeMessageVisibilityBatch(params)

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

func ExampleSQS_CreateQueue() {
	svc := sqs.New(nil)

	params := &sqs.CreateQueueInput{
		QueueName: aws.String("String"), // Required
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateQueue(params)

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

func ExampleSQS_DeleteMessage() {
	svc := sqs.New(nil)

	params := &sqs.DeleteMessageInput{
		QueueURL:      aws.String("String"), // Required
		ReceiptHandle: aws.String("String"), // Required
	}
	resp, err := svc.DeleteMessage(params)

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

func ExampleSQS_DeleteMessageBatch() {
	svc := sqs.New(nil)

	params := &sqs.DeleteMessageBatchInput{
		Entries: []*sqs.DeleteMessageBatchRequestEntry{ // Required
			&sqs.DeleteMessageBatchRequestEntry{ // Required
				ID:            aws.String("String"), // Required
				ReceiptHandle: aws.String("String"), // Required
			},
			// More values...
		},
		QueueURL: aws.String("String"), // Required
	}
	resp, err := svc.DeleteMessageBatch(params)

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

func ExampleSQS_DeleteQueue() {
	svc := sqs.New(nil)

	params := &sqs.DeleteQueueInput{
		QueueURL: aws.String("String"), // Required
	}
	resp, err := svc.DeleteQueue(params)

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

func ExampleSQS_GetQueueAttributes() {
	svc := sqs.New(nil)

	params := &sqs.GetQueueAttributesInput{
		QueueURL: aws.String("String"), // Required
		AttributeNames: []*string{
			aws.String("QueueAttributeName"), // Required
			// More values...
		},
	}
	resp, err := svc.GetQueueAttributes(params)

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

func ExampleSQS_GetQueueURL() {
	svc := sqs.New(nil)

	params := &sqs.GetQueueURLInput{
		QueueName:              aws.String("String"), // Required
		QueueOwnerAWSAccountID: aws.String("String"),
	}
	resp, err := svc.GetQueueURL(params)

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

func ExampleSQS_ListDeadLetterSourceQueues() {
	svc := sqs.New(nil)

	params := &sqs.ListDeadLetterSourceQueuesInput{
		QueueURL: aws.String("String"), // Required
	}
	resp, err := svc.ListDeadLetterSourceQueues(params)

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

func ExampleSQS_ListQueues() {
	svc := sqs.New(nil)

	params := &sqs.ListQueuesInput{
		QueueNamePrefix: aws.String("String"),
	}
	resp, err := svc.ListQueues(params)

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

func ExampleSQS_PurgeQueue() {
	svc := sqs.New(nil)

	params := &sqs.PurgeQueueInput{
		QueueURL: aws.String("String"), // Required
	}
	resp, err := svc.PurgeQueue(params)

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

func ExampleSQS_ReceiveMessage() {
	svc := sqs.New(nil)

	params := &sqs.ReceiveMessageInput{
		QueueURL: aws.String("String"), // Required
		AttributeNames: []*string{
			aws.String("QueueAttributeName"), // Required
			// More values...
		},
		MaxNumberOfMessages: aws.Long(1),
		MessageAttributeNames: []*string{
			aws.String("MessageAttributeName"), // Required
			// More values...
		},
		VisibilityTimeout: aws.Long(1),
		WaitTimeSeconds:   aws.Long(1),
	}
	resp, err := svc.ReceiveMessage(params)

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

func ExampleSQS_RemovePermission() {
	svc := sqs.New(nil)

	params := &sqs.RemovePermissionInput{
		Label:    aws.String("String"), // Required
		QueueURL: aws.String("String"), // Required
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

func ExampleSQS_SendMessage() {
	svc := sqs.New(nil)

	params := &sqs.SendMessageInput{
		MessageBody:  aws.String("String"), // Required
		QueueURL:     aws.String("String"), // Required
		DelaySeconds: aws.Long(1),
		MessageAttributes: &map[string]*sqs.MessageAttributeValue{
			"Key": &sqs.MessageAttributeValue{ // Required
				DataType: aws.String("String"), // Required
				BinaryListValues: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				BinaryValue: []byte("PAYLOAD"),
				StringListValues: []*string{
					aws.String("String"), // Required
					// More values...
				},
				StringValue: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.SendMessage(params)

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

func ExampleSQS_SendMessageBatch() {
	svc := sqs.New(nil)

	params := &sqs.SendMessageBatchInput{
		Entries: []*sqs.SendMessageBatchRequestEntry{ // Required
			&sqs.SendMessageBatchRequestEntry{ // Required
				ID:           aws.String("String"), // Required
				MessageBody:  aws.String("String"), // Required
				DelaySeconds: aws.Long(1),
				MessageAttributes: &map[string]*sqs.MessageAttributeValue{
					"Key": &sqs.MessageAttributeValue{ // Required
						DataType: aws.String("String"), // Required
						BinaryListValues: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						BinaryValue: []byte("PAYLOAD"),
						StringListValues: []*string{
							aws.String("String"), // Required
							// More values...
						},
						StringValue: aws.String("String"),
					},
					// More values...
				},
			},
			// More values...
		},
		QueueURL: aws.String("String"), // Required
	}
	resp, err := svc.SendMessageBatch(params)

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

func ExampleSQS_SetQueueAttributes() {
	svc := sqs.New(nil)

	params := &sqs.SetQueueAttributesInput{
		Attributes: &map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
		QueueURL: aws.String("String"), // Required
	}
	resp, err := svc.SetQueueAttributes(params)

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