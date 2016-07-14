package s3crypto

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Client is an S3 crypto client. By default we will use EncryptionOnly mode which
// will use AES Wrap for key wrapping and AES GCM for content encryption.
// AES GCM will load all data into memory. However, the rest of the content algorithms
// do not load the entire contents into memory.
type Client struct {
	S3     *s3.S3
	Config Config
}

// Config used to customize the Client
type Config struct {
	// Mode is used to dictate how we encrypt and decrypt the data.
	//
	// Defaults to AES GCM with the suggested pairing of AES Wrap/PRSA/KMS
	Mode CryptoMode
	// SaveStrategy will dictate where the envelope is saved.
	//
	// Defaults to the object's metadata
	SaveStrategy
	// InstructionSuffix is the instruction file name suffix. If it is empty, then
	// the item key will be used followed by .instruction
	InstructionSuffix string
	// This is used to instantiate new kms clients when calling GetObject
	KMSSession *session.Session
	S3Session  *session.Session

	// Used for instantiating non-kms key providers when getting objects
	MasterKey []byte
}

// TODO: Add default minimum file and if the contents is less than that, then just
// load it all into memory

// New placeholder
// TODO: Change master cipher to be not ECB
// cipher := NewAESECB(masterkey)
// svc := New(EncryptionOnly(NewSymmetricKeyProvider(cipher))
func New(mode CryptoMode, options ...func(*Client)) *Client {
	sess := session.New()
	// TODO: Change this to strict authenticaton mode
	client := &Client{
		Config: Config{
			Mode:         mode,
			SaveStrategy: NewHeaderSaveStrategy(),
			KMSSession:   sess,
			S3Session:    sess,
		},
	}

	for _, option := range options {
		option(client)
	}

	client.S3 = s3.New(client.Config.S3Session)
	return client
}

// PutObjectRequest creates a temp file to encrypt the contents into. It then streams
// that data to S3.
//
// cipher := NewAESECB(masterkey)
// svc := s3cryto.New(s3crypto.EncryptionOnly(s3crypto.NewSymmetricKeyProvider(cipher))
// req, out := svc.PutObjectRequest(&s3.PutObjectInput {
//	Bucket: aws.String("my_bucket"),
//	Key: aws.String("object_key"),
//	Body: strings.NewReader("WHATEVER"),
// })
// err := req.Send()
func (c *Client) PutObjectRequest(input *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	req, out := c.S3.PutObjectRequest(input)

	// Create temp file to be used later for calculating the SHA256 header
	f, err := ioutil.TempFile("./", *input.Key)
	if err != nil {
		req.Error = err
		return req, out
	}

	req.Handlers.Build.PushFront(func(r *request.Request) {
		md5 := newMD5Reader(input.Body)
		sha := newSHA256Writer(f)
		err = c.Config.Mode.EncryptContents(sha, md5)
		if err != nil {
			r.Error = err
			return
		}

		req.HTTPRequest.Header.Set("X-Amz-Content-Sha256", fmt.Sprintf("%d", sha.GetValue()))

		f.Seek(0, 0)
		input.Body = f

		env, err := EncodeMeta(md5, c.Config.Mode)
		if err != nil {
			r.Error = err
			return
		}
		err = c.Config.SaveStrategy.Save(env, input)
		r.Error = err
	})

	fn := func(r *request.Request) {
		// Close the temp file and cleanup
		f.Close()
		os.Remove(f.Name())
	}
	req.Handlers.Send.PushBack(fn)
	req.Handlers.ValidateResponse.PushBack(fn)
	return req, out
}

// PutObject is a wrapper for PutObjectRequest
func (c *Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	req, out := c.PutObjectRequest(input)
	return out, req.Send()
}

// GetObjectRequest placeholder
func (c *Client) GetObjectRequest(input *s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	req, out := c.S3.GetObjectRequest(input)
	req.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		env, err := c.getEnvelope(input, r)
		if err != nil {
			r.Error = err
			return
		}

		// If KMS should return the correct CEK algorithm with the proper
		// KMS key provider
		mode, err := modeFactory(env, c.Config)
		if err != nil {
			r.Error = err
			return
		}

		reader, err := mode.DecryptContents([]byte(env.CipherKey), []byte(env.IV), out.Body)
		if err != nil {
			r.Error = err
			return
		}
		out.Body = reader
	})
	return req, out
}

// GetObject placeholder
func (c *Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	req, out := c.GetObjectRequest(input)
	return out, req.Send()
}
