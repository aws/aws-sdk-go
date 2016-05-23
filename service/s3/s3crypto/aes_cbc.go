package s3crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
)

// AESCBC Symmetric encryption
type AESCBC struct {
	block     cipher.Block
	iv        []byte
	encrypter cipher.BlockMode
	decrypter cipher.BlockMode
}

// NewAESCBC creates a new AES CBC cypto handler. It suffices
// both interfaces of Encrypter and Decrypter
func NewAESCBC(key, iv []byte) (Cipher, error) {
	block, err := aes.NewCipher(padAESKey(key))
	if err != nil {
		return nil, err
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	decrypter := cipher.NewCBCDecrypter(block, iv)
	return &AESCBC{block, iv, encrypter, decrypter}, nil
}

// Encrypt will encrypt the data using AES CBC
func (c *AESCBC) Encrypt(data io.Reader) (*bytes.Reader, error) {
	plaintext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}
	//encrypter := cipher.NewCBCEncrypter(c.block, c.iv)

	ciphertext := make([]byte, len(plaintext))
	plaintext = PadPKCS7(plaintext, c.encrypter.BlockSize())
	c.encrypter.CryptBlocks(ciphertext, plaintext)
	return bytes.NewReader(ciphertext), nil
}

// Decrypt will decrypt the data using AES CBC
func (c *AESCBC) Decrypt(data io.Reader) (*bytes.Reader, error) {
	ciphertext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}
	c.decrypter.CryptBlocks(ciphertext, ciphertext)
	return bytes.NewReader(ciphertext), nil
}
