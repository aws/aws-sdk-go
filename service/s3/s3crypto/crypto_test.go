package s3crypto_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

func TestCryptoReadCloser_Read(t *testing.T) {
	expectedStr := "HELLO WORLD "
	str := strings.NewReader(expectedStr)
	rc := &s3crypto.CryptoReadCloser{Body: ioutil.NopCloser(str), Decrypter: str}

	b, err := ioutil.ReadAll(rc)
	assert.Nil(t, err)
	assert.Equal(t, expectedStr, string(b))
}
