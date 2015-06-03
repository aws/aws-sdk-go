package sqs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/internal/apierr"
)

var (
	errChecksumMissingBody = fmt.Errorf("cannot compute checksum. missing body")
	errChecksumMissingMD5  = fmt.Errorf("cannot verify checksum. missing response MD5")
)

func setupChecksumValidation(r *aws.Request) {
	if r.Config.DisableComputeChecksums {
		return
	}

	switch r.Operation {
	case opSendMessage:
		r.Handlers.Unmarshal.PushBack(verifySendMessage)
	case opSendMessageBatch:
		r.Handlers.Unmarshal.PushBack(verifySendMessageBatch)
	case opReceiveMessage:
		r.Handlers.Unmarshal.PushBack(verifyReceiveMessage)
	}
}

func verifySendMessage(r *aws.Request) {
	if r.DataFilled() && r.ParamsFilled() {
		in := r.Params.(*SendMessageInput)
		out := r.Data.(*SendMessageOutput)
		err := checksumsMatch(in.MessageBody, out.MD5OfMessageBody)
		if err != nil {
			setChecksumError(r, err.Error())
		}
	}
}

func verifySendMessageBatch(r *aws.Request) {
	if r.DataFilled() && r.ParamsFilled() {
		entries := map[string]*SendMessageBatchResultEntry{}
		ids := []string{}

		out := r.Data.(*SendMessageBatchOutput)
		for _, entry := range out.Successful {
			entries[*entry.ID] = entry
		}

		in := r.Params.(*SendMessageBatchInput)
		for _, entry := range in.Entries {
			if e := entries[*entry.ID]; e != nil {
				err := checksumsMatch(entry.MessageBody, e.MD5OfMessageBody)
				if err != nil {
					ids = append(ids, *e.MessageID)
				}
			}
		}
		if len(ids) > 0 {
			setChecksumError(r, "invalid messages: %s", strings.Join(ids, ", "))
		}
	}
}

func verifyReceiveMessage(r *aws.Request) {
	if r.DataFilled() && r.ParamsFilled() {
		ids := []string{}
		out := r.Data.(*ReceiveMessageOutput)
		for _, msg := range out.Messages {
			err := checksumsMatch(msg.Body, msg.MD5OfBody)
			if err != nil {
				ids = append(ids, *msg.MessageID)
			}
		}
		if len(ids) > 0 {
			setChecksumError(r, "invalid messages: %s", strings.Join(ids, ", "))
		}
	}
}

func checksumsMatch(body, expectedMD5 *string) error {
	if body == nil {
		return errChecksumMissingBody
	} else if expectedMD5 == nil {
		return errChecksumMissingMD5
	}

	msum := md5.Sum([]byte(*body))
	sum := hex.EncodeToString(msum[:])
	if sum != *expectedMD5 {
		return fmt.Errorf("expected MD5 checksum '%s', got '%s'", *expectedMD5, sum)
	}

	return nil
}

func setChecksumError(r *aws.Request, format string, args ...interface{}) {
	r.Retryable.Set(true)
	r.Error = apierr.New("InvalidChecksum", fmt.Sprintf(format, args...), nil)
}
