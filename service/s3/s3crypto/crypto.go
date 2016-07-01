package s3crypto

import (
	"io"
)

// Cipher interface allows for either encryption and decryption of an object
type Cipher interface {
	Encrypter
	Decrypter
}

// TODO: Not used yet
const defaultWriteBufferLimit = 1024 * 1000

// Encrypter interface with only the encrypt method
type Encrypter interface {
	Encrypt(io.Reader) io.Reader
}

// Decrypter interface with only the decrypt method
type Decrypter interface {
	Decrypt(io.Reader) io.Reader
}

// Copy from package io golang stdlib
// Removes length errors since we expect dst to sometimes be smaller
// or larger depending on the cipher.
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

// CryptoReadCloser handles closing of the body and allowing reads from the decrypted
// content.
type CryptoReadCloser struct {
	Body      io.ReadCloser
	Decrypter io.Reader
	isClosed  bool
}

// Close lets the CryptoReadCloser satisfy io.ReadCloser interface
func (rc *CryptoReadCloser) Close() error {
	rc.isClosed = true
	return rc.Body.Close()
}

// Read lets the CryptoReadCloser satisfy io.ReadCloser interface
func (rc *CryptoReadCloser) Read(b []byte) (int, error) {
	if rc.isClosed {
		return 0, io.EOF
	}
	return rc.Decrypter.Read(b)
}
