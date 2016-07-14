package s3crypto

import (
	"io"
)

// Cipher interface allows for either encryption and decryption of an object
type Cipher interface {
	Encrypter
	Decrypter
}

// Wrap placeholder
type Wrap interface {
	Cipher
	GetCipherName() string
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

// CipherName ...
type CipherName struct {
	Algorithm string
}

// GetCipherName ...
func (da *CipherName) GetCipherName() string {
	return da.Algorithm
}
