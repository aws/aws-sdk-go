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
	return &ecbEncryptReader{block: c.block, src: src}
}

type ecbEncryptReader struct {
	block cipher.Block
	src   io.Reader
	buffer
}

func (reader *ecbEncryptReader) Read(data []byte) (int, error) {
	blockSize := reader.block.BlockSize()
	// Drain body til we have one block left then append any new data
	if reader.size > blockSize {
		return reader.drainBody(&data, blockSize), nil
	}

	plaintext := make([]byte, bufSize)
	n, err := reader.src.Read(plaintext)
	if err != nil && err != io.EOF {
		return n, err
	}

	plaintext = plaintext[:n]
	ciphertext := make([]byte, n)

	for i := 0; len(plaintext) > 0; i++ {
		reader.block.Encrypt(ciphertext[blockSize*i:blockSize*(i+1)], plaintext[:blockSize])
		plaintext = plaintext[blockSize:]
	}
	ciphertext = PadPKCS5(ciphertext, blockSize)

	// Nothing has been read, unpad and return EOF.
	if lastBlock := n == 0 || err == io.EOF; lastBlock {
		return reader.finalize(lastBlock, &ciphertext, &data, blockSize), err
	}

	cLen := len(ciphertext)
	// Buffer has too much data in it.
	if reader.size+cLen > bufSize {
		return reader.appendToBuffer(&data, ciphertext), nil
	}

	for i := 0; i < cLen; i++ {
		reader.data[reader.size+i] = ciphertext[i]
	}
	reader.size += cLen
	return 0, nil
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
	// Drain body til we have one block left then append any new data
	if reader.size > blockSize {
		return reader.drainBody(&data, blockSize), nil
	}

	ciphertext := make([]byte, bufSize)
	n, err := reader.src.Read(ciphertext)
	if err != nil && err != io.EOF {
		return n, err
	}

	ciphertext = ciphertext[:n]
	plaintext := make([]byte, n)
	for i := 0; len(ciphertext) > 0; i++ {
		reader.block.Decrypt(plaintext[blockSize*i:blockSize*(i+1)], ciphertext[:blockSize])
		ciphertext = ciphertext[blockSize:]
	}
	plaintext = UnpadPKCS5(plaintext, blockSize)

	if lastBlock := n == 0 || err == io.EOF; lastBlock {
		return reader.finalize(lastBlock, &ciphertext, &data, blockSize), err
	}

	pLen := len(plaintext)
	// Buffer has too much data in it.
	if reader.size+pLen > bufSize {
		return reader.appendToBuffer(&data, plaintext), nil
	}

	for i := 0; i < pLen; i++ {
		reader.data[reader.size+i] = plaintext[i]
	}
	reader.size += pLen
	return 0, nil
}
