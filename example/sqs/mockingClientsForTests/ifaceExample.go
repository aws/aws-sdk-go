package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Queue URL required.")
		os.Exit(1)
	}

	sess := session.Must(session.NewSession())

	q := Queue{
		Client: sqs.New(sess),
		URL:    os.Args[1],
	}

	msgs, err := q.GetMessages(20)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("Messages:")
	for _, msg := range msgs {
		fmt.Println("%s>%s: %s", msg.From, msg.To, msg.Msg)
	}
}

type Queue struct {
	Client sqsiface.SQSAPI
	URL    string
}

type Message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

func (q *Queue) GetMessages(waitTimeout int64) ([]Message, error) {
	params := sqs.ReceiveMessageInput{
		QueueUrl: aws.String(q.URL),
	}
	if waitTimeout > 0 {
		params.WaitTimeSeconds = aws.Int64(waitTimeout)
	}
	resp, err := q.Client.ReceiveMessage(&params)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages, %v", err)
	}

	msgs := make([]Message, len(resp.Messages))
	for i, msg := range resp.Messages {
		parsedMsg := Message{}
		if err := json.Unmarshal([]byte(aws.StringValue(msg.Body)), &parsedMsg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal message, %v", err)
		}

		msgs[i] = parsedMsg
	}

	return msgs, nil
}
