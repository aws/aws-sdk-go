package s3crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
)

const (
	gcmKeySize   = 32
	gcmNonceSize = 12
)

// AESGCM Symmetric encryption algorithm. Since Golang designed this
// with only TLS in mind. We have to load it all into memory meaning
// this isn't streamed.
type AESGCM struct {
	aead  cipher.AEAD
	nonce []byte
}

// NewAESGCMRandom create a new AES GCM cipher, but also randomly
// generates a key and iv.
//
// Example:
//
//	cmkID := "arn to key"
//	kp, _ := s3crypto.NewKMSKeyProvider(session.New(), cmkID, s3crypto.NewJSONMatDesc())
//	cipher, _ := s3crypto.NewAESGCMRandom(kp)
func NewAESGCMRandom(kp KeyProvider) (Cipher, error) {
	key, err := kp.GenerateKey(gcmKeySize)
	if err != nil {
		return nil, err
	}

	nonce, err := kp.GenerateIV(gcmNonceSize)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	kp.SetKey(key)
	kp.SetIV(nonce)
	return &AESGCM{aesgcm, nonce}, nil
}

// NewAESGCM creates a new AES GCM cipher. Expects keys to be of
// the correct size.
//
// Example:
//
//	kp := &s3crypto.SymmetricKeyProvider{}
//	kp.SetKey(key)
//	kp.SetIV(iv)
//	cipher, _ := s3crypto.NewAESGCM(kp)
func NewAESGCM(kp KeyProvider) (Cipher, error) {
	block, err := aes.NewCipher(kp.GetKey())
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &AESGCM{aesgcm, kp.GetIV()}, nil
}

// Encrypt will encrypt the data using AES GCM
// Tag will be included as the last 16 bytes of the slice
func (c *AESGCM) Encrypt(src io.Reader) io.Reader {
	reader := &gcmEncryptReader{
		encrypter: c.aead,
		nonce:     c.nonce,
		src:       src,
	}
	return reader
}

type gcmEncryptReader struct {
	encrypter cipher.AEAD
	nonce     []byte
	src       io.Reader
	buf       *bytes.Buffer
}

func (reader *gcmEncryptReader) Read(data []byte) (int, error) {
	if reader.buf == nil {
		b, err := ioutil.ReadAll(reader.src)
		if err != nil {
			return len(b), err
		}
		b = reader.encrypter.Seal(nil, reader.nonce, b, nil)
		reader.buf = bytes.NewBuffer(b)
	}

	return reader.buf.Read(data)
}

// Decrypt will decrypt the data using AES GCM
func (c *AESGCM) Decrypt(src io.Reader) io.Reader {
	return &gcmDecryptReader{
		decrypter: c.aead,
		nonce:     c.nonce,
		src:       src,
	}
}

type gcmDecryptReader struct {
	decrypter cipher.AEAD
	nonce     []byte
	src       io.Reader
	buf       *bytes.Buffer
}

func (reader *gcmDecryptReader) Read(data []byte) (int, error) {
	if reader.buf == nil {
		b, err := ioutil.ReadAll(reader.src)
		if err != nil {
			return len(b), err
		}
		b, err = reader.decrypter.Open(nil, reader.nonce, b, nil)
		if err != nil {
			return len(b), err
		}

		reader.buf = bytes.NewBuffer(b)
	}

	return reader.buf.Read(data)
}
