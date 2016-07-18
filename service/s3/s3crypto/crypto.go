package s3crypto

import (
	"io"
)

// Cipher interface allows for either encryption and decryption of an object
type Cipher interface {
	Encrypter
	Decrypter
}

// Wrap interface includes an additional method that allows us to get the cipher
// name. This is used during encryption when populating the envelope information.
type Wrap interface {
	Cipher
	CipherDataIface
}

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

// CipherData is used to populate the envelope during encryption calls
type CipherData struct {
	Algorithm string
	TagLen    string
}

// GetCipherName handles populating x-amz-wrap-alg and x-amz-cek-alg when
// adding the envelope.
func (cd *CipherData) GetCipherName() string {
	return cd.Algorithm
}

// GetTagLen handles populating x-amz-tag-len when adding the envelope.
func (cd *CipherData) GetTagLen() string {
	return cd.TagLen
}
