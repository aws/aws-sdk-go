package s3crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
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
func (c *AESECB) Encrypt(src io.Reader) io.Reader {
	return &ecbEncryptReader{c.block, src}
}

type ecbEncryptReader struct {
	block cipher.Block
	src   io.Reader
}

func (reader *ecbEncryptReader) Read(plaintext []byte) (int, error) {
	data := make([]byte, len(plaintext))
	n, err := reader.src.Read(data)
	if err != nil {
		return n, err
	}

	ciphertext := make([]byte, n)
	blockSize := reader.block.BlockSize()
	data = PadPKCS5(data[:n], blockSize)
	for i := 0; len(data) > 0; i++ {
		reader.block.Encrypt(ciphertext[blockSize*i:blockSize*(i+1)], data[:blockSize])
		data = data[blockSize:]
	}
	plaintext = append(plaintext[:0], ciphertext...)
	return len(plaintext), err
}

// Decrypt will decrypt the data using AES ECB
func (c *AESECB) Decrypt(src io.Reader) io.Reader {
	return &ecbDecryptReader{c.block, src}
}

type ecbDecryptReader struct {
	block cipher.Block
	src   io.Reader
}

func (reader *ecbDecryptReader) Read(ciphertext []byte) (int, error) {
	data := make([]byte, len(ciphertext))
	n, err := reader.src.Read(data)
	if err != nil {
		return n, err
	}

	data = data[:n]

	plaintext := make([]byte, n)
	blockSize := reader.block.BlockSize()
	for i := 0; len(data) > 0; i++ {
		reader.block.Decrypt(plaintext[blockSize*i:blockSize*(i+1)], data[:blockSize])
		data = data[blockSize:]
	}
	plaintext = UnpadPKCS5(plaintext, blockSize)
	ciphertext = append(ciphertext[:0], plaintext...)
	return len(ciphertext), err
}
