package s3crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
)

// AESECB Symmetric encryption used for masterkey
type AESECB struct {
	block cipher.Block
	CipherData
}

// NewAESECB creates a new AES CBC cypto handler. It suffices
// both interfaces of Encrypter and Decrypter
func NewAESECB(key []byte) (Wrap, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AESECB{block, CipherData{Algorithm: "ecb"}}, nil
}

// Encrypt will encrypt the data using AES ECB
func (c *AESECB) Encrypt(src io.Reader) io.Reader {
	return &ecbEncryptReader{block: c.block, src: src}
}

type ecbEncryptReader struct {
	block cipher.Block
	src   io.Reader
	buffer
}

func (reader *ecbEncryptReader) Read(data []byte) (int, error) {
	blockSize := reader.block.BlockSize()
	if len(reader.data) == 0 {
		plaintext, err := ioutil.ReadAll(reader.src)
		if err != nil {
			return 0, err
		}
		ciphertext := make([]byte, len(plaintext))
		for i := 0; len(plaintext) > 0; i++ {
			reader.block.Encrypt(ciphertext[blockSize*i:blockSize*(i+1)], plaintext[:blockSize])
			plaintext = plaintext[blockSize:]
		}
		ciphertext = PadPKCS5(ciphertext, blockSize)
		reader.data = ciphertext
	}

	size := len(data)
	if size > len(reader.data) {
		size = len(reader.data)
	}

	copy(data, reader.data[:size])
	reader.data = reader.data[size:]
	if len(reader.data) == 0 {
		return size, io.EOF
	}
	return size, nil
}

// Decrypt will decrypt the data using AES ECB
func (c *AESECB) Decrypt(src io.Reader) io.Reader {
	return &ecbDecryptReader{block: c.block, src: src}
}

type ecbDecryptReader struct {
	block cipher.Block
	src   io.Reader
	buffer
}

func (reader *ecbDecryptReader) Read(data []byte) (int, error) {
	blockSize := reader.block.BlockSize()
	if len(reader.data) == 0 {
		ciphertext, err := ioutil.ReadAll(reader.src)
		if err != nil {
			return 0, err
		}
		plaintext := make([]byte, len(ciphertext))
		for i := 0; len(ciphertext) > 0; i++ {
			reader.block.Decrypt(plaintext[blockSize*i:blockSize*(i+1)], ciphertext[:blockSize])
			ciphertext = ciphertext[blockSize:]
		}
		plaintext = UnpadPKCS5(plaintext, blockSize)
		reader.data = plaintext
	}

	size := len(data)
	if size > len(reader.data) {
		size = len(reader.data)
	}

	copy(data, reader.data[:size])
	reader.data = reader.data[size:]
	if len(reader.data) == 0 {
		return size, io.EOF
	}
	return size, nil
}
