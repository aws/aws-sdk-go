package s3crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
)

// AESECB Symmetric encryption used for masterkey
type AESECB struct {
	block cipher.Block
}

// NewAESECB creates a new AES CBC cypto handler. It suffices
// both interfaces of Encrypter and Decrypter
func NewAESECB(key []byte) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AESECB{block}, nil
}

// Encrypt will encrypt the data using AES ECB
func (c *AESECB) Encrypt(data io.Reader) (*bytes.Reader, error) {
	plaintext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}

	ciphertext := make([]byte, len(plaintext))
	blockSize := c.block.BlockSize()
	plaintext = PadPKCS5(plaintext, blockSize)
	for i := 0; len(plaintext) > 0; i++ {
		c.block.Encrypt(ciphertext[blockSize*i:blockSize*(i+1)], plaintext[:blockSize])
		plaintext = plaintext[blockSize:]
	}
	return bytes.NewReader(ciphertext), nil
}

// Decrypt will decrypt the data using AES ECB
func (c *AESECB) Decrypt(data io.Reader) (*bytes.Reader, error) {
	ciphertext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}

	plaintext := make([]byte, len(ciphertext))
	blockSize := c.block.BlockSize()
	for i := 0; len(ciphertext) > 0; i++ {
		c.block.Decrypt(plaintext[blockSize*i:blockSize*(i+1)], ciphertext[:blockSize])
		ciphertext = ciphertext[blockSize:]
	}
	plaintext = UnpadPKCS5(plaintext, blockSize)
	return bytes.NewReader(plaintext), nil
}
