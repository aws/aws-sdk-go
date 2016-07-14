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
	BaseKeyProvider

	// This is the CEK algorithm used
	CipherName
}

// EncryptionOnly returns a new encryption only mode structure with a specific cipher
// for the master key
func EncryptionOnly(kp KeyProvider) CryptoMode {
	return &EncryptionOnlyMode{
		BaseKeyProvider: BaseKeyProvider{kp},
		CipherName:      CipherName{"AES/CBC/PKCS5Padding"},
	}
}

// EncryptContents will generate a random key and iv and encrypt the data using cbc
func (mode *EncryptionOnlyMode) EncryptContents(dst io.Writer, src io.Reader) error {
	// Sets the key and iv to a randomly generated key and iv
	cbc, err := NewAESCBCRandom(mode)
	if err != nil {
		return err
	}

	reader := cbc.Encrypt(src)
	_, err = io.Copy(dst, reader)
	return err
}

// DecryptContents placeholder
func (mode *EncryptionOnlyMode) DecryptContents(key, iv []byte, src io.ReadCloser) (io.ReadCloser, error) {
	kp := &SymmetricKeyProvider{key: key, iv: iv}
	cbc, err := NewAESCBC(kp)
	if err != nil {
		return nil, err
	}

	reader := cbc.Decrypt(src)
	return &CryptoReadCloser{Body: src, Decrypter: reader}, nil
}
