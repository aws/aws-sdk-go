package s3crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

// AESCBC Symmetric encryption
type AESCBC struct {
	block     cipher.Block
	iv        []byte
	key       []byte
	encrypter cipher.BlockMode
	decrypter cipher.BlockMode
}

const cbcKeySize = 32
const cbcIVSize = 16

// NewAESCBCRandom ...
func NewAESCBCRandom(kp KeyProvider) (Cipher, error) {
	key, err := kp.GenerateKey(cbcKeySize)
	if err != nil {
		return nil, err
	}

	iv, err := kp.GenerateIV(cbcIVSize)
	if err != nil {
		return nil, err
	}

	kp.SetKey(key)
	kp.SetIV(iv)
	return NewAESCBC(kp)
}

// NewAESCBC creates a new AES CBC cypto handler. It suffices
// both interfaces of Encrypter and Decrypter
// If an empty key or iv is provided, a randomly generated
// key and iv is provided to the cipher
func NewAESCBC(kp KeyProvider) (Cipher, error) {
	block, err := aes.NewCipher(padAESKey(kp.GetKey()))
	if err != nil {
		return nil, err
	}

	encrypter := cipher.NewCBCEncrypter(block, kp.GetIV())
	decrypter := cipher.NewCBCDecrypter(block, kp.GetIV())
	return &AESCBC{block, kp.GetKey(), kp.GetIV(), encrypter, decrypter}, nil
}

// Encrypt will encrypt the data using AES CBC
func (c *AESCBC) Encrypt(src io.Reader) io.Reader {
	reader := &cbcEncryptReader{
		encrypter: c.encrypter,
		src:       src,
	}
	return reader
}

// Decrypt will decrypt the data using AES CBC
func (c *AESCBC) Decrypt(src io.Reader) io.Reader {
	reader := &cbcDecryptReader{
		decrypter: c.decrypter,
		src:       src,
	}
	return reader
}

type cbcEncryptReader struct {
	encrypter cipher.BlockMode
	src       io.Reader
	buffer
}

// Need to ensure each block read is a multiple of the block size
func (reader *cbcEncryptReader) Read(data []byte) (int, error) {
	// Drain body til we have one block left then append any new data
	blockSize := reader.encrypter.BlockSize()
	if reader.size > blockSize {
		return reader.drainBody(&data, blockSize), nil
	}

	ciphertext := make([]byte, bufSize)
	n, err := reader.src.Read(ciphertext)
	if err != nil && err != io.EOF {
		return n, err
	}

	ciphertext = PadPKCS5(ciphertext[:n], blockSize)
	reader.encrypter.CryptBlocks(ciphertext, ciphertext)

	if lastBlock := (n == 0 || err == io.EOF); lastBlock {
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

type cbcDecryptReader struct {
	decrypter cipher.BlockMode
	src       io.Reader
	buffer
}

func (reader *cbcDecryptReader) Read(data []byte) (int, error) {
	blockSize := reader.decrypter.BlockSize()
	if reader.size > blockSize {
		n := reader.drainBody(&data, blockSize)
		return n, nil
	}

	plaintext := make([]byte, bufSize)

	n, err := reader.src.Read(plaintext)
	if err != nil && err != io.EOF {
		return n, err
	}

	plaintext = plaintext[:n]
	reader.decrypter.CryptBlocks(plaintext, plaintext)
	plaintext = UnpadPKCS5(plaintext, reader.decrypter.BlockSize())

	if lastBlock := (n == 0 || err == io.EOF); lastBlock {
		n = reader.finalize(lastBlock, &plaintext, &data, blockSize)
		return n, err
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
