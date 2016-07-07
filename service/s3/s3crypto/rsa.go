package s3crypto

import (
	"crypto/rsa"
	"hash"
	"io"
	"io/ioutil"
)

// OAEPRSA placeholder
type OAEPRSA struct {
	key  rsa.PrivateKey
	hash hash.Hash
	rng  io.Reader
}

// NewOAEPRSA placeholder
func NewOAEPRSA(key rsa.PrivateKey, h hash.Hash, rng io.Reader) Cipher {
	return &OAEPRSA{
		key:  key,
		hash: h,
		rng:  rng,
	}
}

// Encrypt placeholder
func (cipher *OAEPRSA) Encrypt(src io.Reader) io.Reader {
	return &oaepRSAEncryptReader{
		src:     src,
		OAEPRSA: cipher,
	}
}

type oaepRSAEncryptReader struct {
	*OAEPRSA
	src  io.Reader
	data []byte
}

func (reader *oaepRSAEncryptReader) Read(data []byte) (int, error) {
	rLen := len(reader.data)
	if rLen == 0 {
		b, err := ioutil.ReadAll(reader.src)
		if err != nil {
			return len(b), err
		}
		reader.data, err = rsa.EncryptOAEP(reader.hash, reader.rng, &reader.key.PublicKey, b, []byte(""))
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

type oaepRSADecryptReader struct {
	src io.Reader
	*OAEPRSA
	data []byte
}

// Decrypt placeholder
func (cipher *OAEPRSA) Decrypt(src io.Reader) io.Reader {
	return &oaepRSADecryptReader{
		src:     src,
		OAEPRSA: cipher,
	}
}

func (reader *oaepRSADecryptReader) Read(data []byte) (int, error) {
	rLen := len(reader.data)
	if rLen == 0 {
		b, err := ioutil.ReadAll(reader.src)
		if err != nil {
			return len(b), err
		}
		reader.data, err = rsa.DecryptOAEP(reader.hash, reader.rng, &reader.key, b, []byte(""))
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
