package s3crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
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

// NewAESCBC creates a new AES CBC cypto handler. It suffices
// both interfaces of Encrypter and Decrypter
// If an empty key or iv is provided, a randomly generated
// key and iv is provided to the cipher
//
// TODO: See if there is a better way of randomly generating
// keys and iv
func NewAESCBC(kp *SymmetricKeyProvider) (Cipher, error) {
	if len(kp.Key) == 0 {
		kp.Key = make([]byte, cbcKeySize)
		kp.IV = make([]byte, cbcIVSize)
		rand.Read(kp.Key)
		rand.Read(kp.IV)
	}

	block, err := aes.NewCipher(padAESKey(kp.Key))
	if err != nil {
		return nil, err
	}

	encrypter := cipher.NewCBCEncrypter(block, kp.IV)
	decrypter := cipher.NewCBCDecrypter(block, kp.IV)
	return &AESCBC{block, kp.Key, kp.IV, encrypter, decrypter}, nil
}

// Encrypt will encrypt the data using AES CBC
// TODO: Return an io.Writer? This will allow Decrypt and Encrypt
// to behave in the same way
func (c *AESCBC) Encrypt(src io.Reader) io.Reader {
	reader := &cbcEncryptReader{
		c.encrypter,
		src,
	}
	return reader
}

// Decrypt will decrypt the data using AES CBC
func (c *AESCBC) Decrypt(src io.Reader) io.Reader {
	reader := &cbcDecryptReader{
		c.decrypter,
		src,
	}
	return reader
}

type cbcEncryptReader struct {
	encrypter cipher.BlockMode
	src       io.Reader
}

// Need to ensure each block read is a multiple of the block size
// TODO:
// create a blocksizeReader and wrap io.Reader in it
func (writer *cbcEncryptReader) Read(plaintext []byte) (int, error) {
	n, err := writer.src.Read(plaintext)
	if err != nil {
		return n, err
	}
	plaintext = PadPKCS5(plaintext[:n], writer.encrypter.BlockSize())

	writer.encrypter.CryptBlocks(plaintext, plaintext)
	return n, err
}

type cbcDecryptWriter struct {
	decrypter cipher.BlockMode
	dst       io.Writer
}

func (writer *cbcDecryptWriter) Write(ciphertext []byte) (int, error) {
	writer.decrypter.CryptBlocks(ciphertext, ciphertext)
	ciphertext = UnpadPKCS5(ciphertext, writer.decrypter.BlockSize())
	return writer.dst.Write(ciphertext)
}

type cbcDecryptReader struct {
	decrypter cipher.BlockMode
	src       io.Reader
}

func (reader *cbcDecryptReader) Read(ciphertext []byte) (int, error) {
	n, err := reader.src.Read(ciphertext)
	if err != nil {
		return n, err
	}

	ciphertext = ciphertext[:n]
	reader.decrypter.CryptBlocks(ciphertext, ciphertext)
	ciphertext = UnpadPKCS5(ciphertext, reader.decrypter.BlockSize())
	return n, err
}
