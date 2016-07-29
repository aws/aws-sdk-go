package s3crypto_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
	"github.com/stretchr/testify/assert"
)

func TestGetCipherData(t *testing.T) {
	cd := s3crypto.CipherData{
		Key:                 []byte("testkey"),
		IV:                  []byte("testiv"),
		WrapAlgorithm:       "wrap",
		CEKAlgorithm:        "cek",
		TagLength:           "taglen",
		MaterialDescription: s3crypto.MaterialDescription{},
	}
	expected := cd
	data := cd.GetCipherData()
	assert.NotNil(t, data)
	assert.Equal(t, expected, *data)
}
