package s3crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
)

// AESGCM Symmetric encryption
type AESGCM struct {
	aead  cipher.AEAD
	nonce []byte
}

// NewAESGCM creates a new authenticated crypto handler.
func NewAESGCM(key, nonce []byte) (*AESGCM, error) {
	block, err := aes.NewCipher(padAESKey(key))
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &AESGCM{aesgcm, nonce}, nil
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
	data      []byte
}

func (reader *gcmEncryptReader) Read(data []byte) (int, error) {
	rLen := len(reader.data)
	if rLen == 0 {
		b, err := ioutil.ReadAll(reader.src)
		if err != nil {
			return 0, err
		}
		reader.data = reader.encrypter.Seal(nil, reader.nonce, b, nil)
		// prevent infinite loops
		if len(reader.data) == 0 {
			return 0, io.EOF
		}
	}

	return gcmCopyBuffer(rLen, &data, &reader.data)
}

/*func (c *AESGCM) Encrypt(data io.Reader) (*bytes.Reader, error) {
	plaintext, err := ioutil.ReadAll(data)
	if err != nil {
		return bytes.NewReader([]byte{}), nil
	}

	aesgcm, err := cipher.NewGCM(c.block)
	if err != nil {
		return bytes.NewReader([]byte{}), nil
	}

	ciphertext := aesgcm.Seal(nil, c.nonce, plaintext, nil)
	return bytes.NewReader(ciphertext), nil
}*/

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
	data      []byte
}

func (reader *gcmDecryptReader) Read(data []byte) (int, error) {
	rLen := len(reader.data)
	if rLen == 0 {
		b, err := ioutil.ReadAll(reader.src)
		if err != nil {
			return 0, err
		}
		reader.data, err = reader.decrypter.Open(nil, reader.nonce, b, nil)
		if err != nil {
			return len(b), err
		}
		// prevent infinite loops
		if len(reader.data) == 0 {
			return len(b), io.EOF
		}
	}

	return gcmCopyBuffer(rLen, &data, &reader.data)
}

func gcmCopyBuffer(rLen int, dst, src *[]byte) (int, error) {
	max := len(*dst)
	if max > rLen {
		max = rLen
	}

	copy(*dst, (*src)[:max])
	(*src) = (*src)[max:]
	var err error
	if len(*src) == 0 {
		err = io.EOF
	}
	return max, err
}
