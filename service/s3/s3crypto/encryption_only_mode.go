package s3crypto

import (
	"io"
)

// EncryptionOnlyMode will use:
// Wrap Algorithm: AES Keywrap
// CEK Algorithm: AES CBC with PKCS5 padding
type EncryptionOnlyMode struct {
	key []byte
	iv  []byte
	// Cipher will hold either a master symmetric key or a KMS client
	// for encrypting the Envelope key
	Wrap
}

// EncryptionOnly returns a new encryption only mode structure with a specific cipher
// for the master key
func EncryptionOnly(wrap Wrap) CryptoMode {
	return &EncryptionOnlyMode{Wrap: wrap}
}

// EncryptContents will generate a random key and iv and encrypt the data using cbc
func (mode *EncryptionOnlyMode) EncryptContents(dst io.Writer, src io.Reader) error {
	kp := &SymmetricKeyProvider{}
	// Sets the key and iv to a randomly generated key and iv
	cbc, err := NewAESCBC(kp)
	if err != nil {
		return err
	}

	mode.key = kp.Key
	mode.iv = kp.IV
	reader := cbc.Encrypt(src)
	_, err = translate(dst, reader)
	return err
}

// DecryptContents placeholder
func (mode *EncryptionOnlyMode) DecryptContents(key, iv []byte, src io.ReadCloser) (io.ReadCloser, error) {
	kp := &SymmetricKeyProvider{key, iv}
	cbc, err := NewAESCBC(kp)
	if err != nil {
		return nil, err
	}

	reader := cbc.Decrypt(src)
	return &CryptoReadCloser{Body: src, Decrypter: reader}, nil
}

// GetKey returns the randomly generated key
func (mode *EncryptionOnlyMode) GetKey() []byte {
	return mode.key
}

// GetIV returns the randomly generated iv
func (mode *EncryptionOnlyMode) GetIV() []byte {
	return mode.iv
}
