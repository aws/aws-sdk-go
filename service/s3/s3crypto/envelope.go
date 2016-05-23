package s3crypto

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Envelope encryption starts off by generating a random symmetric key using AES CBC,
// AES CTR, or AES GCM. We then generate a random IV based off the encryption cipher
// chosen. The master key that was provided, whether by the user or KMS, will be used
// to encrypt the randomly generated symmetric key and base64 encode the iv. This will
// allow for decryption of that same data later.
type Envelope struct {
	// IV is the randomly generated IV base64 encoded.
	IV []byte `json:"x-amz-iv"`
	// CipherKey is the randomly generated cipher key.
	CipherKey []byte `json:"x-amz-key"`
	// MaterialDesc is a description to distinguish from other envelopes.
	MaterialDesc string `json:"x-amz-matdesc"`
	Meta         meta   `json:"-"`
}

type meta struct {
	Bucket    string
	Request   *request.Request
	ObjectKey string
}

// SaveStrategy is how the data's metadata wants to be saved
type SaveStrategy interface {
	Save(Envelope)
}

// default s3 key
const instructionKey = ".instruction"

type s3SaveStrategy struct {
	client                *s3.S3
	InstructionFileSuffix *string
}

// NewS3SaveStrategy returns a new strategy that will save to S3
func NewS3SaveStrategy(p client.ConfigProvider, suffix *string) SaveStrategy {
	return &s3SaveStrategy{s3.New(p), suffix}
}

// Save will save the envelope contents to s3.
func (strat *s3SaveStrategy) Save(env Envelope) {
	env.Meta.Request.Handlers.Send.PushFront(func(r *request.Request) {
		b, err := json.Marshal(env)
		if err != nil {
			r.Error = err
			return
		}

		instInput := s3.PutObjectInput{
			Bucket: &env.Meta.Bucket,
			Body:   bytes.NewReader(b),
		}

		if strat.InstructionFileSuffix == nil {
			instInput.Key = aws.String(env.Meta.ObjectKey + instructionKey)
		} else {
			instInput.Key = aws.String(env.Meta.ObjectKey + instructionKey + "-" + *strat.InstructionFileSuffix)
		}

		_, err = strat.client.PutObject(&instInput)
		if err != nil {
			r.Error = err
		}
	})
}

type headerSaveStrategy struct{}

// NewHeaderSaveStrategy returns a new strategy that will save to the metadata
func NewHeaderSaveStrategy() SaveStrategy {
	return &headerSaveStrategy{}
}

// Save will save the envelope to the request's header.
func (strat *headerSaveStrategy) Save(env Envelope) {
	env.Meta.Request.HTTPRequest.Header.Set("X-Amz-Meta-X-Amz-Iv", string(env.IV))
	env.Meta.Request.HTTPRequest.Header.Set("X-Amz-Meta-X-Amz-Key", string(env.CipherKey))
	env.Meta.Request.HTTPRequest.Header.Set("X-Amz-Meta-X-Amz-MatDesc", env.MaterialDesc)
}
