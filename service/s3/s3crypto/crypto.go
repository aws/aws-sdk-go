package s3crypto

import (
	"io"
)

// Cipher placeholder
type Cipher interface {
	Encrypter
	Decrypter
}

const defaultWriteBufferLimit = 1024 * 1000

// Encrypter interface with only the encrypt method
type Encrypter interface {
	Encrypt(io.Reader) io.Reader
}

// Decrypter interface with only the decrypt method
type Decrypter interface {
	Decrypt(io.Reader) io.Reader
}

// CipherConstructor is constructors for symmetric keys
type CipherConstructor func(key, iv []byte) (Cipher, error)

// Copy from package io golang stdlib
func translate(dst io.Writer, src io.Reader) (written int64, err error) {
	buf := make([]byte, 32*1024)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				break
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}

// CryptoReadCloser placeholder
type CryptoReadCloser struct {
	Body      io.ReadCloser
	Decrypter io.Reader
	Pipe      io.Reader
	isClosed  bool
}

// Close placeholder
func (rc *CryptoReadCloser) Close() error {
	rc.isClosed = true
	return rc.Body.Close()
}

func (rc *CryptoReadCloser) Read(b []byte) (int, error) {
	if rc.isClosed {
		return 0, io.EOF
	}
	return rc.Decrypter.Read(b)
}
