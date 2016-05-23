package s3crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
)

// AESCTR Symmetric encryption
type AESCTR struct {
	block cipher.Block
	iv    []byte
}

// NewAESCTR creates a new AES CTR cypto handler. It suffices
// both interfaces of Encrypter and Decrypter
func NewAESCTR(key, iv []byte) (*AESCTR, error) {
	block, err := aes.NewCipher(padAESKey(key))
	if err != nil {
		return nil, err
	}

	return &AESCTR{block, iv}, nil
}

// Encrypt will encrypt the data using AES CTR
func (c *AESCTR) Encrypt(data io.Reader) (*bytes.Reader, error) {
	plaintext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(c.block, c.iv)
	stream.XORKeyStream(ciphertext, plaintext)
	return bytes.NewReader(ciphertext), nil
}

// Decrypt will decrypt the data using AES CTR
func (c *AESCTR) Decrypt(data io.Reader) (*bytes.Reader, error) {
	ciphertext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(c.block, c.iv)
	stream.XORKeyStream(plaintext, ciphertext)
	return bytes.NewReader(plaintext), nil
}
