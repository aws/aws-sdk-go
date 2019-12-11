package eventstreamapi

import (
	"bytes"
	"time"

	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
)

const (
	chunkSignatureHeader = ":chunk-signature"
	chunkDateHeader      = ":date"
)

// StreamSigner defines an interface for the implementation of signing of event stream payloads
type StreamSigner interface {
	GetSignature(headers, payload []byte, date time.Time) ([]byte, error)
}

// MessageSigner encapsulates signing and attaching signatures to event stream messages
type MessageSigner struct {
	Signer StreamSigner
}

// SignMessage takes the given event stream message generates and adds signature information
// to the event stream message.
func (s MessageSigner) SignMessage(msg *eventstream.Message, date time.Time) error {
	msg.Headers.Set(chunkDateHeader, eventstream.TimestampValue(date))

	var headers bytes.Buffer
	if err := eventstream.EncodeHeaders(&headers, msg.Headers); err != nil {
		return err
	}

	sig, err := s.Signer.GetSignature(headers.Bytes(), msg.Payload, date)
	if err != nil {
		return err
	}

	msg.Headers.Set(chunkSignatureHeader, eventstream.BytesValue(sig))

	return nil
}
