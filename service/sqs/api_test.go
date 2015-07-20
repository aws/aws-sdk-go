// +build integration

package sqs_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
)

func TestFlattenedTraits(t *testing.T) {
	s := sqs.New(nil)
	_, err := s.DeleteMessageBatch(&sqs.DeleteMessageBatchInput{
		QueueURL: awsconv.String("QUEUE"),
		Entries: []*sqs.DeleteMessageBatchRequestEntry{
			{
				ID:            awsconv.String("TEST"),
				ReceiptHandle: awsconv.String("RECEIPT"),
			},
		},
	})

	assert.Error(t, err)
	assert.Equal(t, "InvalidAddress", err.Code())
	assert.Equal(t, "The address QUEUE is not valid for this endpoint.", err.Message())
}
