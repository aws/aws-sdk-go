package s3crypto

import (
	"io"
)

// AuthenticationMode will use AES GCM for the main cipher.
type AuthenticationMode struct {
	// Cipher will hold either a master symmetric key or a KMS client
	// for encrypting the Envelope key
	BaseKeyProvider

	// This is the CEK algorithm used
	CipherData
}

// Authentication returns a new encryption only mode structure with a specific cipher
// for the master key
func Authentication(kp KeyProvider) CryptoMode {
	return &AuthenticationMode{
		BaseKeyProvider: BaseKeyProvider{kp},
		CipherData:      CipherData{AESGCMNoPadding, "128"},
	}
}

// EncryptContents will generate a random key and iv and encrypt the data using cbc
/*func (mode *AuthenticationMode) EncryptContents(dst io.Writer, src io.Reader) error {
	// Sets the key and iv to a randomly generated key and iv
	cbc, err := NewAESGCMRandom(mode)
	if err != nil {
		return err
	}

	reader := cbc.Encrypt(src)
	_, err = io.Copy(dst, reader)
	return err
}*/

// EncryptContents will generate a random key and iv and encrypt the data using cbc
func (mode *AuthenticationMode) EncryptContents(src io.Reader) (io.Reader, error) {
	// Sets the key and iv to a randomly generated key and iv
	cbc, err := NewAESGCMRandom(mode)
	if err != nil {
		return nil, err
	}

	return cbc.Encrypt(src), err
}

// DecryptContents will use the symmetric key provider to instantiate a new GCM cipher.
// We grab a decrypt reader from gcm and wrap it in a CryptoReadCloser. The only error
// expected here is when the key or iv is of invalid length.
func (mode *AuthenticationMode) DecryptContents(kp KeyProvider, src io.ReadCloser) (io.ReadCloser, error) {
	gcm, err := NewAESGCM(kp)
	if err != nil {
		return nil, err
	}

	reader := gcm.Decrypt(src)
	return &CryptoReadCloser{Body: src, Decrypter: reader}, nil
}
