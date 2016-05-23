package s3crypto

import (
	"encoding/hex"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
)

func TestNewClient(t *testing.T) {
	key, _ := hex.DecodeString("31bdadd96698c204aa9ce1448ea94ae1fb4a9a0b3c9d773b51bb1822666b8f22")
	cipher, err := NewAESECB(key)
	assert.Nil(t, err)

	sess := session.New(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	svc, err := NewClient(cipher, sess)

	assert.Nil(t, err)
	assert.NotNil(t, svc)
}

func TestEnvelopeSaveToHeader(t *testing.T) {
	iv, _ := hex.DecodeString("11958dc6ab81e1c7f01631e9944e620f")
	key, _ := hex.DecodeString("31bdadd96698c204aa9ce1448ea94ae1fb4a9a0b3c9d773b51bb1822666b8f22")

	req := request.Request{
		HTTPRequest: &http.Request{},
	}
	req.HTTPRequest.Header = make(http.Header)
	env := Envelope{
		IV:           iv,
		CipherKey:    key,
		MaterialDesc: "Testing123",
		Meta: meta{
			Request: &req,
		},
	}

	strat := headerSaveStrategy{}
	strat.Save(env)

	assert.Equal(t, req.HTTPRequest.Header.Get("X-Amz-Meta-X-Amz-Iv"), string(iv))
	assert.Equal(t, req.HTTPRequest.Header.Get("X-Amz-Meta-X-Amz-Key"), string(key))
	assert.Equal(t, req.HTTPRequest.Header.Get("X-Amz-Meta-X-Amz-MatDesc"), env.MaterialDesc)
}

/*
func TestSaveToS3(t *testing.T) {
	iv, _ := hex.DecodeString("11958dc6ab81e1c7f01631e9944e620f")
	key, _ := hex.DecodeString("31bdadd96698c204aa9ce1448ea94ae1fb4a9a0b3c9d773b51bb1822666b8f22")

	//flags := aws.LogDebugWithHTTPBody
	cipher, err := NewAESECB(key, iv)
	assert.Nil(t, err)
	sess := session.New(&aws.Config{
		Region: aws.String("us-west-2"),
		//LogLevel: &flags,
	})

	svc, err := NewClient(cipher,sess)

	//strat := NewS3SaveStrategy(sess, nil)
	strat := NewHeaderSaveStrategy()
	assert.Nil(t, err)
	out, err := svc.PutObject(&PutObjectInput{
		SaveStrategy: strat,
		S3PutObjectInput: &s3.PutObjectInput{
			Bucket: aws.String("bucketmesilly"),
			Key:    aws.String("yar"),
			Body:   bytes.NewReader(bytes.Repeat([]byte{byte('A')}, 1024)),
		},
	})
	_ = out
	t.Log(out, err)
	assert.Nil(t, err)

	tout, err := svc.GetObject(&GetObjectInput{
		S3GetObjectInput: &s3.GetObjectInput{
			Bucket: aws.String("bucketmesilly"),
			Key:    aws.String("yar"),
		},
	})
	t.Log(tout, err)
}*/

func TestGenerateBytes(t *testing.T) {
	iv := generateRandBytes(16)
	assert.Equal(t, len(iv), 16)

	iv = generateRandBytes(15)
	assert.Equal(t, len(iv), 15)

	iv = generateRandBytes(17)
	assert.Equal(t, len(iv), 17)
}
