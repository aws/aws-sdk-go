package s3crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
)

// AESGCM Symmetric encryption
type AESGCM struct {
	block  cipher.Block
	nounce []byte
}

// NewAESGCM creates a new authenticated crypto handler.
func NewAESGCM(key, nounce []byte) (*AESGCM, error) {
	block, err := aes.NewCipher(padAESKey(key))
	if err != nil {
		return nil, err
	}

	return &AESGCM{block, nounce}, nil
}

// Encrypt will encrypt the data using AES GCM
// Tag will be included as the last 16 bytes of the slice
func (c *AESGCM) Encrypt(data io.Reader) (*bytes.Reader, error) {
	plaintext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), nil
	}

	aesgcm, err := cipher.NewGCM(c.block)
	if err != nil {
		return bytes.NewReader([]byte{}), nil
	}

	ciphertext := aesgcm.Seal(nil, c.nounce, plaintext, nil)
	return bytes.NewReader(ciphertext), nil
}

// Decrypt will decrypt the data using AES GCM
func (c *AESGCM) Decrypt(data io.Reader) (*bytes.Reader, error) {
	ciphertext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}

	aesgcm, err := cipher.NewGCM(c.block)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}

	plaintext, err := aesgcm.Open(nil, c.nounce, ciphertext, nil)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}

	return bytes.NewReader(plaintext), nil
}
