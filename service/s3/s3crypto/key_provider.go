package s3crypto

import "crypto/rand"

// KeyProvider placeholder
type KeyProvider interface {
	EncrypterKeyProvider
	DecrypterKeyProvider
}

// EncrypterKeyProvider placeholder
type EncrypterKeyProvider interface {
	GenerateKey(int) ([]byte, error)
	GenerateIV(int) ([]byte, error)
	GetEncryptedKey(key []byte) ([]byte, error)
	GetIV() []byte
	SetIV([]byte)
	GetKey() []byte
	SetKey([]byte)
}

// DecrypterKeyProvider placeholder
type DecrypterKeyProvider interface {
	GetDecryptedKey([]byte) ([]byte, error)
}

func generateBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

// BaseKeyProvider ...
type BaseKeyProvider struct {
	KeyProvider
}

// GetKeyProvider ...
func (kp *BaseKeyProvider) GetKeyProvider() KeyProvider {
	return kp.KeyProvider
}
