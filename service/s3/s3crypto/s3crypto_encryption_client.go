package s3crypto

import (
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// EncryptionClient is an S3 crypto client. By default the SDK will use Authentication mode which
// will use KMS for key wrapping and AES GCM for content encryption.
// AES GCM will load all data into memory. However, the rest of the content algorithms
// do not load the entire contents into memory.
type EncryptionClient struct {
	S3API  s3iface.S3API
	Config EncryptionConfig
}

// EncryptionConfig used to customize the Client
type EncryptionConfig struct {
	ContentCipherBuilder ContentCipherBuilder
	// SaveStrategy will dictate where the envelope is saved.
	//
	// Defaults to the object's metadata
	SaveStrategy SaveStrategy
	// InstructionFileSuffix is the instruction file name suffix when using get requests.
	// If it is empty, then the item key will be used followed by .instruction
	InstructionFileSuffix string
	S3Session             client.ConfigProvider
	// TempFolderPath is used to store temp files when calling PutObject
	TempFolderPath string
}

// NewEncryptionClient instantiates a new S3 crypto client
//
// Example:
//	cmkID := "some key id to kms"
//	sess := session.New()
//	handler, err = s3crypto.NewKMSEncryptHandler(sess, cmkID, s3crypto.MaterialDescription{})
//	if err != nil {
//	  return err
//	}
//	svc := s3crypto.New(sess, s3crypto.AESGCMContentCipherBuilder(handler))
func NewEncryptionClient(prov client.ConfigProvider, builder ContentCipherBuilder, options ...func(*EncryptionClient)) *EncryptionClient {
	client := &EncryptionClient{
		Config: EncryptionConfig{
			ContentCipherBuilder: builder,
			SaveStrategy:         headerSaveStrategy{},
			S3Session:            prov,
		},
	}

	for _, option := range options {
		option(client)
	}

	client.S3API = s3.New(client.Config.S3Session)
	return client
}

// PutObjectRequest creates a temp file to encrypt the contents into. It then streams
// that data to S3.
//
// Example:
//	svc := s3crypto.New(session.New(), s3crypto.AESGCMContentCipherBuilder(handler))
//	req, out := svc.PutObjectRequest(&s3.PutObjectInput {
//	  Key: aws.String("testKey"),
//	  Bucket: aws.String("testBucket"),
//	  Body: bytes.NewBuffer("test data"),
//	})
//	err := req.Send()
func (c *EncryptionClient) PutObjectRequest(input *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	req, out := c.S3API.PutObjectRequest(input)

	// Create temp file to be used later for calculating the SHA256 header
	f, err := ioutil.TempFile(c.Config.TempFolderPath, "")
	if err != nil {
		req.Error = err
		return req, out
	}

	encryptor, err := c.Config.ContentCipherBuilder.NewEncryptor()
	req.Handlers.Build.PushFront(func(r *request.Request) {
		if err != nil {
			r.Error = err
			return
		}

		md5 := newMD5Reader(input.Body)
		sha := newSHA256Writer(f)
		reader, err := encryptor.EncryptContents(md5)
		if err != nil {
			r.Error = err
			return
		}

		_, err = io.Copy(sha, reader)
		if err != nil {
			r.Error = err
			return
		}

		data := encryptor.GetCipherData()
		env, err := encodeMeta(md5, data)
		if err != nil {
			r.Error = err
			return
		}

		shaHex := hex.EncodeToString(sha.GetValue())
		req.HTTPRequest.Header.Set("X-Amz-Content-Sha256", shaHex)

		f.Seek(0, 0)
		input.Body = f

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
func (c *EncryptionClient) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	req, out := c.PutObjectRequest(input)
	return out, req.Send()
}
