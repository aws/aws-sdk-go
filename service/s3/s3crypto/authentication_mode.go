package s3crypto

import (
	"io"
)

// AuthenticationMode will use:
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
		CipherData:      CipherData{"AES/GCM/NoPadding", "128"},
	}
}

// EncryptContents will generate a random key and iv and encrypt the data using cbc
func (mode *AuthenticationMode) EncryptContents(dst io.Writer, src io.Reader) error {
	// Sets the key and iv to a randomly generated key and iv
	cbc, err := NewAESGCMRandom(mode)
	if err != nil {
		return err
	}

	reader := cbc.Encrypt(src)
	_, err = io.Copy(dst, reader)
	return err
}

// DecryptContents placeholder
func (mode *AuthenticationMode) DecryptContents(key, iv []byte, src io.ReadCloser) (io.ReadCloser, error) {
	kp := &SymmetricKeyProvider{key: key, iv: iv}
	cbc, err := NewAESGCM(kp)
	if err != nil {
		return nil, err
	}

	reader := cbc.Decrypt(src)
	return &CryptoReadCloser{Body: src, Decrypter: reader}, nil
}
