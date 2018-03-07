package connectiface

import "github.com/aws/aws-sdk-go/service/connect"

type ConnectAPI interface {
	StartOutboundVoiceContact(*connect.StartOutboundVoiceContactInput) (connect.StartOutboundVoiceContactOutput, error)
}

var _ ConnectAPI = (*connect.Connect)(nil)
