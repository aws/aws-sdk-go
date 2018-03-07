package connect

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

const startOutboundVoiceCOntact = "startOutboundVoiceContact"

type StartOutboundVoiceContactInput struct {
	_                      struct{} `type:"structure"`
	Attributes             *map[string]string
	ClientToken            *string
	ContactFlowId          *string
	DestinationPhoneNumber *string
	InstanceId             *string
	QueueId                *string
	SourcePhoneNumber      *string
}

type StartOutboundVoiceContactOutput struct {
	_         struct{} `type:"structure"`
	ContactId *string
}

func (c *Connect) StartOutboundVoiceContact(input *StartOutboundVoiceContactInput) (*StartOutboundVoiceContactOutput, error) {
	req, out := c.startOutboundVoiceContactRequest(input)
	return out, req.Send()
}

func (c *Connect) startOutboundVoiceContactRequest(input *StartOutboundVoiceContactInput) (req *request.Request, output *StartOutboundVoiceContactOutput) {
	op := &request.Operation{
		Name:       startOutboundVoiceCOntact,
		HTTPMethod: "PUT",
		HTTPPath:   "/contact/outbound-voice",
	}

	if input == nil {
		input = &StartOutboundVoiceContactInput{}
	}

	output = &StartOutboundVoiceContactOutput{}
	req = c.newRequest(op, input, output)
	req.Config.DisableSSL = aws.Bool(true)
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}
