package eventstreamapi

import (
	"bytes"
	"time"

	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
)

var timeNow = time.Now

// StreamSigner defines an interface for the implementation of signing of event stream payloads
type StreamSigner interface {
	GetSignature(headers, payload []byte, date time.Time) ([]byte, error)
}

type SignEncoder struct {
	signer     StreamSigner
	encoder    Encoder
	bufEncoder *BufferEncoder

	closeErr error
	closed   bool
}

func NewSignEncoder(signer StreamSigner, encoder Encoder) *SignEncoder {
	// TODO: Need to pass down logging

	return &SignEncoder{
		signer:     signer,
		encoder:    encoder,
		bufEncoder: NewBufferEncoder(),
	}
}

func (s *SignEncoder) Close() error {
	if s.closed {
		return s.closeErr
	}

	s.closeErr = s.encode([]byte{})
	s.closed = true
	return s.closeErr
}

func (s *SignEncoder) Encode(msg eventstream.Message) error {
	payload, err := s.bufEncoder.Encode(msg)
	if err != nil {
		return err
	}

	return s.encode(payload)
}

func (s SignEncoder) encode(payload []byte) error {
	date := timeNow()

	var msg eventstream.Message
	msg.Headers.Set(DateHeader, eventstream.TimestampValue(date))
	msg.Payload = payload

	var headers bytes.Buffer
	if err := eventstream.EncodeHeaders(&headers, msg.Headers); err != nil {
		return err
	}

	sig, err := s.signer.GetSignature(headers.Bytes(), msg.Payload, date)
	if err != nil {
		return err
	}

	msg.Headers.Set(ChunkSignatureHeader, eventstream.BytesValue(sig))

	return s.encoder.Encode(msg)
}

type BufferEncoder struct {
	encoder Encoder
	buffer  *bytes.Buffer
}

func NewBufferEncoder() *BufferEncoder {
	buf := bytes.NewBuffer(make([]byte, 1024))
	return &BufferEncoder{
		encoder: eventstream.NewEncoder(buf),
		buffer:  buf,
	}
}

func (e *BufferEncoder) Encode(msg eventstream.Message) ([]byte, error) {
	e.buffer.Reset()

	if err := e.encoder.Encode(msg); err != nil {
		return nil, err
	}

	return e.buffer.Bytes(), nil
}

//// MessageSigner encapsulates signing and attaching signatures to event stream messages
//type MessageSigner struct {
//	Signer StreamSigner
//}
//
//// SignMessage takes the given event stream message generates and adds signature information
//// to the event stream message.
//func (s MessageSigner) SignMessage(msg *eventstream.Message, date time.Time) error {
//	var headers bytes.Buffer
//	if err := eventstream.EncodeHeaders(&headers, msg.Headers); err != nil {
//		return err
//	}
//
//	sig, err := s.Signer.GetSignature(headers.Bytes(), msg.Payload, date)
//	if err != nil {
//		return err
//	}
//
//	msg.Headers.Set(ChunkSignatureHeader, eventstream.BytesValue(sig))
//	return nil
//}
