package s3crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

// AESCBC Symmetric encryption
// AESCBC is susceptible to Oracle Padding attacks. It is strongly
// preferred the use of AES GCM or AES CBC+Mac or AES CTR+Mac
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
	blockSize := reader.encrypter.BlockSize()
	ciphertext := make([]byte, bufSize)
	n, err := reader.src.Read(ciphertext)
	reader.data = append(reader.data, ciphertext[:n]...)

	if v := len(reader.data); v > 0 {
		size := len(data)
		if len(reader.data) < size {
			size = len(reader.data)
		}

		// Incase padded data is larger than len(data)
		if size > blockSize {
			size -= size % blockSize
		}

		ciphertext := make([]byte, size)
		ciphertext = PadPKCS5(reader.data[:size], blockSize)
		reader.encrypter.CryptBlocks(ciphertext, ciphertext)
		copy(data, ciphertext)
		reader.data = reader.data[size:]
		if len(reader.data) == 0 && err == io.EOF {
			return len(ciphertext), err
		}
		return len(ciphertext), nil
	}
	return 0, err
}

type cbcDecryptReader struct {
	decrypter cipher.BlockMode
	src       io.Reader
	buffer
}

func (reader *cbcDecryptReader) Read(data []byte) (int, error) {
	plaintext := make([]byte, bufSize)
	n, err := reader.src.Read(plaintext)
	reader.data = append(reader.data, plaintext[:n]...)

	blockSize := reader.decrypter.BlockSize()
	// we can check if v%blockSize, because data is always padded on encrypt
	if v := len(reader.data); v > 0 && v%blockSize == 0 {
		size := len(data)
		if len(reader.data) < size {
			size = len(reader.data)
		}
		plaintext := make([]byte, size)
		reader.decrypter.CryptBlocks(plaintext, reader.data[:size])
		plaintext = UnpadPKCS5(plaintext, blockSize)
		copy(data, plaintext)
		reader.data = reader.data[size:]
		if len(reader.data) == 0 && err == io.EOF {
			return len(plaintext), err
		}
		return len(plaintext), nil
	}
	return 0, err
}
