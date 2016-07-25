package s3crypto

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// DefaultInstructionKeySuffix is appended to the end of the instruction file key when
// grabbing or saving to S3
const DefaultInstructionKeySuffix = ".instruction"

const (
	metaHeader                     = "x-amz-meta"
	keyV1Header                    = "x-amz-key"
	keyV2Header                    = keyV1Header + "-v2"
	ivHeader                       = "x-amz-iv"
	matDescHeader                  = "x-amz-matdesc"
	cekAlgorithmHeader             = "x-amz-cek-alg"
	wrapAlgorithmHeader            = "x-amz-wrap-alg"
	tagLengthHeader                = "x-amz-tag-len"
	unencryptedMD5Header           = "x-amz-unencrypted-content-md5"
	unencryptedContentLengthHeader = "x-amz-unencrypted-content-length"
)

// Envelope encryption starts off by generating a random symmetric key using
// AES GCM. The SDK generates a random IV based off the encryption cipher
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
	if value := r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, keyV2Header}, "-")); value != "" {
		return getV2Envelope(r)
	} else if value = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, keyV1Header}, "-")); value != "" {
		return getV1Envelope(r)
	} else {
		return getFromInstructionFile(client.S3, input, r, client.Config.InstructionFileSuffix)
	}
}

func getV2Envelope(r *request.Request) (*Envelope, error) {
	env := &Envelope{}
	env.CipherKey = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, keyV2Header}, "-"))
	env.IV = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, ivHeader}, "-"))
	env.MatDesc = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, matDescHeader}, "-"))
	env.WrapAlg = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, wrapAlgorithmHeader}, "-"))
	env.CEKAlg = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, cekAlgorithmHeader}, "-"))
	env.TagLen = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, tagLengthHeader}, "-"))
	env.UnencryptedMD5 = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, unencryptedMD5Header}, "-"))
	env.UnencryptedContentLen = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, unencryptedContentLengthHeader}, "-"))
	env.version = 2
	return env, nil
}

func getV1Envelope(r *request.Request) (*Envelope, error) {
	env := &Envelope{}
	env.CipherKey = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, keyV1Header}, "-"))
	env.IV = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, ivHeader}, "-"))
	env.MatDesc = r.HTTPResponse.Header.Get(strings.Join([]string{metaHeader, matDescHeader}, "-"))
	env.version = 1
	return env, nil
}

// TODO: write test
func getFromInstructionFile(svc s3iface.S3API, input *s3.GetObjectInput, r *request.Request, suffix string) (*Envelope, error) {
	if suffix == "" {
		suffix = DefaultInstructionKeySuffix
	}
	out, err := svc.GetObject(&s3.GetObjectInput{
		Key:    aws.String(strings.Join([]string{*input.Key, suffix}, "")),
		Bucket: input.Bucket,
	})
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return nil, err
	}
	env := &Envelope{}
	err = json.Unmarshal(b, env)
	return env, err
}

// SaveStrategy is how the data's metadata wants to be saved
type SaveStrategy interface {
	Save(Envelope, *s3.PutObjectInput) error
}

type s3SaveStrategy struct {
	client                *s3.S3
	InstructionFileSuffix *string
}

// NewS3SaveStrategy returns a new strategy that will save to S3
func NewS3SaveStrategy(p client.ConfigProvider, suffix *string) SaveStrategy {
	return &s3SaveStrategy{s3.New(p), suffix}
}

// Save will save the envelope contents to s3.
func (strat s3SaveStrategy) Save(env Envelope, input *s3.PutObjectInput) error {
	b, err := json.Marshal(env)
	if err != nil {
		return err
	}

	instInput := s3.PutObjectInput{
		Bucket: input.Bucket,
		Body:   bytes.NewReader(b),
	}

	if strat.InstructionFileSuffix == nil {
		instInput.Key = aws.String(*input.Key + DefaultInstructionKeySuffix)
	} else {
		instInput.Key = aws.String(*input.Key + *strat.InstructionFileSuffix)
	}

	_, err = strat.client.PutObject(&instInput)
	return err
}

type headerSaveStrategy struct{}

// NewHeaderSaveStrategy returns a new strategy that will save to the metadata
func NewHeaderSaveStrategy() SaveStrategy {
	return &headerSaveStrategy{}
}

// Save will save the envelope to the request's header.
func (strat headerSaveStrategy) Save(env Envelope, input *s3.PutObjectInput) error {
	if input.Metadata == nil {
		input.Metadata = map[string]*string{}
	}

	input.Metadata[http.CanonicalHeaderKey(keyV2Header)] = &env.CipherKey
	input.Metadata[http.CanonicalHeaderKey(ivHeader)] = &env.IV
	input.Metadata[http.CanonicalHeaderKey(matDescHeader)] = &env.MatDesc
	input.Metadata[http.CanonicalHeaderKey(wrapAlgorithmHeader)] = &env.WrapAlg
	input.Metadata[http.CanonicalHeaderKey(cekAlgorithmHeader)] = &env.CEKAlg
	input.Metadata[http.CanonicalHeaderKey(tagLengthHeader)] = &env.TagLen
	input.Metadata[http.CanonicalHeaderKey(unencryptedMD5Header)] = &env.UnencryptedMD5
	input.Metadata[http.CanonicalHeaderKey(unencryptedContentLengthHeader)] = &env.UnencryptedContentLen
	return nil
}
