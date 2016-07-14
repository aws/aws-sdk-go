package s3crypto

import (
	"fmt"
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
}

// EncryptionOnly returns a new encryption only mode structure with a specific cipher
// for the master key
// TODO: Wrap may not be needed and just put master key stuff on KeyProvider
func EncryptionOnly(kp KeyProvider) CryptoMode {
	return &EncryptionOnlyMode{BaseKeyProvider: BaseKeyProvider{kp}}
}

// EncryptContents will generate a random key and iv and encrypt the data using cbc
func (mode *EncryptionOnlyMode) EncryptContents(dst io.Writer, src io.Reader) error {
	// Sets the key and iv to a randomly generated key and iv
	cbc, err := NewAESCBCRandom(mode)
	if err != nil {
		return err
	}

	// TODO: Don't think this is needed
	// kp.SetEncryptedKey(mode.Wrap)
	reader := cbc.Encrypt(src)
	n, err := io.Copy(dst, reader)
	fmt.Println("TET", n, err)
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
