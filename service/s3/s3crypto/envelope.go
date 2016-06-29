package s3crypto

import (
	//"bytes"
	//"encoding/json"

	//"github.com/aws/aws-sdk-go/aws"
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
	IV string `json:"x-amz-iv"`
	// CipherKey is the randomly generated cipher key.
	CipherKey string `json:"x-amz-key-v2, x-amz-key"`
	// MaterialDesc is a description to distinguish from other envelopes.
	MatDesc               string `json:"x-amz-matdesc"`
	WrapAlg               string `json:"x-amz-wrap-alg"`
	CEKAlg                string `json:"x-amz-cek-alg"`
	TagLen                string `json:"x-amz-tag-len"`
	UnencryptedMD5        string `json:"x-amz-unencrypted-content-md5"`
	UnencryptedContentLen string `json:"x-amz-unencrypted-content-length"`
	version               int
}

func (client *Client) getEnvelope(input *s3.GetObjectInput, r *request.Request) (*Envelope, error) {
	if value := r.HTTPResponse.Header.Get("x-amz-meta-x-amz-key-v2"); value != "" {
		return getV2Envelope(r)
	} else if value = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-key"); value != "" {
		return getV1Envelope(r)
	} else {
		return getFromInstructionFile(input, r)
	}
}

// TODO:
// Need MD5 and SHA256 io.Reader
// Need to redo all io.Reader interface for encrypt and decrypt
// Implement KMS and Private Key RSA
func getV2Envelope(r *request.Request) (*Envelope, error) {
	env := &Envelope{}
	env.CipherKey = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-key-v2")
	env.IV = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-iv")
	env.MatDesc = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-matdesc")
	env.WrapAlg = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-wrap-alg")
	env.CEKAlg = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-cek-alg")
	env.TagLen = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-tag-len")
	env.UnencryptedMD5 = r.HTTPResponse.Header.Get("x-amz-unencrypted-content-md5")
	env.UnencryptedContentLen = r.HTTPResponse.Header.Get("x-amz-unencrypted-content-length")
	env.version = 2
	return env, nil
}

func getV1Envelope(r *request.Request) (*Envelope, error) {
	env := &Envelope{}
	env.CipherKey = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-key")
	env.IV = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-iv")
	env.MatDesc = r.HTTPResponse.Header.Get("x-amz-meta-x-amz-matdesc")
	env.version = 1
	return env, nil
}

func getFromInstructionFile(input *s3.GetObjectInput, r *request.Request) (*Envelope, error) {
	return &Envelope{}, nil
}

// SaveStrategy is how the data's metadata wants to be saved
type SaveStrategy interface {
	Save(Envelope, *s3.PutObjectInput) error
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
func (strat *s3SaveStrategy) Save(env Envelope, input *s3.PutObjectInput) error {
	/*env.Meta.Request.Handlers.Send.PushFront(func(r *request.Request) {
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
	})*/
	return nil
}

type headerSaveStrategy struct{}

// NewHeaderSaveStrategy returns a new strategy that will save to the metadata
func NewHeaderSaveStrategy() SaveStrategy {
	return &headerSaveStrategy{}
}

// Save will save the envelope to the request's header.
func (strat *headerSaveStrategy) Save(env Envelope, input *s3.PutObjectInput) error {
	if input.Metadata == nil {
		input.Metadata = make(map[string]*string)
	}

	input.Metadata["X-Amz-Key-V2"] = &env.CipherKey
	input.Metadata["X-Amz-Iv"] = &env.IV
	input.Metadata["X-Amz-MatDesc"] = &env.MatDesc
	input.Metadata["X-Amz-Wrap-Alg"] = &env.WrapAlg
	input.Metadata["X-Amz-Cek-Alg"] = &env.CEKAlg
	input.Metadata["X-Amz-Tag-Len"] = &env.TagLen
	input.Metadata["X-Amz-Unencrypted-Content-Md5"] = &env.UnencryptedMD5
	input.Metadata["X-Amz-Unencrypted-Content-Length"] = &env.UnencryptedContentLen
	return nil
}
