// +build integration
package sqs_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/sqs"
)

var sqsClient *sqs.SQS

func TestMain(m *testing.M) {
	if os.Getenv("INTEGRATION") != "" {
		sqsClient = sqs.New(aws.DefaultCreds(), "us-east-1", nil)
		os.Exit(m.Run())
	}
}

func TestPublishAndReadMessage(t *testing.T) {

	queueName := "AWS_GO_PUBLISH_TEST_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	sqsClient := sqs.New(aws.DefaultCreds(), "us-east-1", nil)

	createQueueInput := sqs.CreateQueueRequest{
		QueueName: &queueName,
		Attributes: sqs.AttributeMap{
			"VisibilityTimeout": "40",
		},
	}

	createQueueOutput, err := sqsClient.CreateQueue(&createQueueInput)
	if err != nil {
		t.Fatal("Failed to create SQS Queue: ", err)
	}

	if *(createQueueOutput.QueueURL) == "" {
		t.Fatal("Queue URL was not marshalled back")
	}
	defer deleteQueue(createQueueOutput.QueueURL)

	sendMessageOutput, err := sqsClient.SendMessage(&sqs.SendMessageRequest{
		QueueURL:    createQueueOutput.QueueURL,
		MessageBody: aws.String("Hello Go"),
		MessageAttributes: sqs.MessageAttributeMap{
			"test_attribute_name_1": sqs.MessageAttributeValue{
				StringValue: aws.String("StringValue1"),
				DataType:    aws.String("String"),
			},
			"test_attribute_name_2": sqs.MessageAttributeValue{
				StringValue: aws.String("StringValue2"),
				DataType:    aws.String("String"),
			},
		},
	})
	if err != nil {
		t.Fatal("Failed to send message: ", err)
	}

	if *(sendMessageOutput.MessageID) == "" {
		t.Fatal("MessageId was not marshalled back")
	}
	if *(sendMessageOutput.MD5OfMessageBody) == "" {
		t.Fatal("MD5OfMessageBody was not marshalled back")
	}
	if *(sendMessageOutput.MD5OfMessageAttributes) == "" {
		t.Fatal("MD5OfMessageAttributes was not marshalled back")
	}

	var message *sqs.Message

	for i := 0; i < 5; i++ {

		receiveMessageOutput, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageRequest{
			QueueURL:              createQueueOutput.QueueURL,
			WaitTimeSeconds:       aws.Integer(10),
			MessageAttributeNames: []string{"All"},
		})
		if err != nil {
			t.Fatal("Failed to receive message: ", err)
		}

		if len(receiveMessageOutput.Messages) > 0 {
			message = &(receiveMessageOutput.Messages[0])
			break
		}
	}

	if message == nil {
		t.Fatal("Failed to receive sent message")
	}

	if *(message.Body) != "Hello Go" {
		t.Fatal("Failed to get message body")
	}

	if len(message.MessageAttributes) != 2 {
		t.Fatal("Missing message attributes: ", len(message.Attributes))
	}

	if *(message.MessageAttributes["test_attribute_name_1"].StringValue) != "StringValue1" {
		t.Fatal("Incorrect marshall of message attribute value")
	}

	if *(message.MessageAttributes["test_attribute_name_2"].StringValue) != "StringValue2" {
		t.Fatal("Incorrect marshall of message attribute value")
	}

}

func deleteQueue(queueURL *string) {

	deleteQueueInput := sqs.DeleteQueueRequest{
		QueueURL: queueURL,
	}

	sqsClient.DeleteQueue(&deleteQueueInput)
}
