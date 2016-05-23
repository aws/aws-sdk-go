package s3crypto

// DR TOOLS TEST integration test bucket
import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	ivSize  = 16
	keySize = 32
)

// Client supports client level encryption and decryption to S3.
// The crypto client uses "envelope" encryption".
type Client struct {
	s3client *s3.S3
	cfg      Config
	// masterKey will hold either a master symmetric key or a KMS client
	// for encrypting the envelope key
	masterKey Cipher
}

// Config used to customize the Client
type Config struct {
}

// NewClient will return a crypto client and set an aes cbc encrypter
// as one of the fields.
func NewClient(masterKey Cipher, p client.ConfigProvider, options ...func(*Config)) (*Client, error) {
	client := &Client{}

	client.s3client = s3.New(p)
	client.masterKey = masterKey

	for _, option := range options {
		option(&client.cfg)
	}

	return client, nil
}

// PutObjectInput for the PutObjectRequest
type PutObjectInput struct {
	SaveStrategy
	MaterialDesc     string
	S3PutObjectInput *s3.PutObjectInput
}

// PutObjectRequest will call the S3's PutObjectRequest and then save the envelope
// to the appropriate place.
func (client *Client) PutObjectRequest(f CipherConstructor, input *PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	req, out := client.s3client.PutObjectRequest(input.S3PutObjectInput)

	// TODO: Generalize this
	iv := generateRandBytes(ivSize)
	key := generateRandBytes(keySize)
	enc, err := f(key, iv)

	req.Handlers.Build.PushFront(func(r *request.Request) {
		if err != nil {
			r.Error = err
			return
		}

		body, err := enc.Encrypt(input.S3PutObjectInput.Body)
		if err != nil {
			r.Error = err
			return
		}

		input.S3PutObjectInput.Body = body
	})

	req.Handlers.Build.PushBack(func(r *request.Request) {
		env := Envelope{
			MaterialDesc: input.MaterialDesc,
			Meta: meta{ // TODO: Consider rethinking of putting this data elsewhere
				*input.S3PutObjectInput.Bucket,
				req,
				*input.S3PutObjectInput.Key,
			},
		}
		err := client.saveEnvelope(input.SaveStrategy, env, key, iv)
		r.Error = err
	})
	return req, out
}

// PutObject behaves identically to S3's PutObject, except that it calls the crypto
// PutObjectRequest method.
func (client *Client) PutObject(f CipherConstructor, input *PutObjectInput) (*s3.PutObjectOutput, error) {
	req, out := client.PutObjectRequest(f, input)
	err := req.Send()
	return out, err
}

// GetObjectInput for the GetObjectRequest
type GetObjectInput struct {
	S3GetObjectInput *s3.GetObjectInput
}

// GetObjectRequest will call the S3's GetObjectRequest. The SDK will first grab the object
// and check if the x-amz-meta-x-amz-matdesc exists. If it does, it'll use that data for decryption.
// If it isn't on the object, we will then grab the necessary instruction file.
// We will manipulate the handlers to choose which method of getting the keys
func (client *Client) GetObjectRequest(f CipherConstructor, input *GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	req, out := client.s3client.GetObjectRequest(input.S3GetObjectInput)

	// TODO: Put handler logic into own functions
	req.Handlers.Send.PushBack(func(r *request.Request) {
		matdesc := ""
		key := []byte{}
		iv := []byte{}
		encKey := &bytes.Reader{}

		if matdesc = r.HTTPResponse.Header.Get("X-Amz-Meta-X-Amz-Matdesc"); matdesc == "" {
			// TODO: Go get the {{ ObjectKey }}.instruction file
		} else {
			var err error
			key, err = base64.StdEncoding.DecodeString(r.HTTPResponse.Header.Get("X-Amz-Meta-X-Amz-Key"))
			if err != nil {
				r.Error = err
				return
			}
			key = decryptWithGCM(key)
			iv, err = base64.StdEncoding.DecodeString(r.HTTPResponse.Header.Get("X-Amz-Meta-X-Amz-Iv"))
			if err != nil {
				r.Error = err
				return
			}
			fmt.Println("KEY", key, len(key))
			encKey, err = client.masterKey.Decrypt(bytes.NewReader(key))
			if err != nil {
				r.Error = err
				return
			}
			key, err = ioutil.ReadAll(encKey)
			if err != nil {
				r.Error = err
				return
			}
		}

		fmt.Println(key, len(key))
		cipher, err := f(key, iv)
		if err != nil {
			fmt.Println(err)
			r.Error = err
			return
		}

		body, err := cipher.Decrypt(r.HTTPResponse.Body)
		if err != nil {
			r.Error = err
			return
		}
		r.HTTPResponse.Body = ioutil.NopCloser(body)
	})
	return req, out
}

// GetObject simply calls GetObjectRequest
func (client *Client) GetObject(f CipherConstructor, input *GetObjectInput) (*s3.GetObjectOutput, error) {
	req, out := client.GetObjectRequest(f, input)
	err := req.Send()
	return out, err
}

// saveEnvelope will bootstrap the envelope to either be saved as a separate object
// in S3 or add it to the request's header. The bootstrapping would encrypt the randomly
// generated symmetric key, base64 encode the iv, and lastly set the material description.
func (client *Client) saveEnvelope(strat SaveStrategy, env Envelope, envKey, envIv []byte) error {
	encKey, err := client.masterKey.Encrypt(bytes.NewReader(envKey))
	if err != nil {
		return err
	}

	key, err := ioutil.ReadAll(encKey)
	if err != nil {
		return err
	}

	if env.MaterialDesc == "" {
		env.MaterialDesc = "{}"
	}

	env.IV = []byte(base64.StdEncoding.EncodeToString(envIv))
	env.CipherKey = []byte(base64.StdEncoding.EncodeToString(key))

	strat.Save(env)
	return nil
}

func generateRandBytes(size int) []byte {
	b := make([]byte, size)
	rmd := size % 4

	for i := 0; i+4 < size; i += 4 {
		bytes := rand.Uint32()
		for j := 0; j < 4; j++ {
			b[i+j] = byte(bytes & 0xFF)
			bytes >>= 8
		}
	}

	bytes := rand.Uint32()
	for i := 0; i < rmd; i++ {
		b[i] = byte(bytes & 0xFF)
		bytes >>= 8
	}
	return b
}

//DONE Generate a random symmetric envelope key and initialization vector.
//DONE Encrypt file using a cipher created with the envelope key and initialization vector.
//Encrypt that envelope key using the master public RSA key
//DONE or the master symmetric key provided by the user.
//DONE Store the encrypted envelope key and initialization vector with the encrypted file, base64 encoded with no newlines.
//DONE Store a description of the master key (defaults to "{}")
