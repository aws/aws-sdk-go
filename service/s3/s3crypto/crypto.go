package s3crypto

import (
	"bytes"
	"io"
)

// Cipher placeholder
type Cipher interface {
	Encrypter
	Decrypter
}

// Encrypter interface with only the encrypt method
type Encrypter interface {
	Encrypt(io.Reader) (*bytes.Reader, error)
}

// Decrypter interface with only the decrypt method
type Decrypter interface {
	Decrypt(io.Reader) (*bytes.Reader, error)
}

type CipherConstructor func(key, iv []byte) (Cipher, error)
