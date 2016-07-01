package s3crypto

import (
	"bytes"
	"io/ioutil"
)

// SymmetricKeyProvider placeholder
type SymmetricKeyProvider struct {
	key          []byte
	encryptedKey []byte
	iv           []byte
	Cipher
}

// NewSymmetricKeyProvider ...
func NewSymmetricKeyProvider(master Cipher) KeyProvider {
	return &SymmetricKeyProvider{Cipher: master}
}

// GenerateKey placeholder
func (kp *SymmetricKeyProvider) GenerateKey(n int) ([]byte, error) {
	return generateBytes(n), nil
}

// GenerateIV placeholder
func (kp *SymmetricKeyProvider) GenerateIV(n int) ([]byte, error) {
	return generateBytes(n), nil
}

// GetKey placeholder
func (kp *SymmetricKeyProvider) GetKey() []byte {
	return kp.key
}

// SetKey placeholder
func (kp *SymmetricKeyProvider) SetKey(key []byte) {
	kp.key = key
}

// GetEncryptedKey placeholder
func (kp *SymmetricKeyProvider) GetEncryptedKey() ([]byte, error) {
	dst := kp.Encrypt(bytes.NewBuffer(kp.key))
	return ioutil.ReadAll(dst)
}

// GetDecryptedKey placeholder
func (kp *SymmetricKeyProvider) GetDecryptedKey() ([]byte, error) {
	dst := kp.Decrypt(bytes.NewBuffer(kp.encryptedKey))
	b, err := ioutil.ReadAll(dst)
	return b, err
}

// GetIV placeholder
func (kp *SymmetricKeyProvider) GetIV() []byte {
	return kp.iv
}

// SetIV placeholder
func (kp *SymmetricKeyProvider) SetIV(b []byte) {
	kp.iv = b
}

// EncryptKey will also set the encrypted key?
func (kp *SymmetricKeyProvider) EncryptKey(cipher Encrypter) {
}

// SetEncryptedKey placeholder
func (kp *SymmetricKeyProvider) SetEncryptedKey(key []byte) {
	kp.encryptedKey = key
}
