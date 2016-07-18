package s3crypto

import "crypto/rand"

// KeyProvider is used to help generate keys and ivs. This allows for crypto based
// algorithms or services to be used.
type KeyProvider interface {
	EncrypterKeyProvider
	DecrypterKeyProvider
}

// EncrypterKeyProvider handles how to generate keys and ivs for the content encryption.
// This also satifies the `MaterialDescription` and `CipherDataIface` interfaces.
type EncrypterKeyProvider interface {
	GenerateKey(int) ([]byte, error)
	GenerateIV(int) ([]byte, error)
	GetEncryptedKey(key []byte) ([]byte, error)
	GetIV() []byte
	SetIV([]byte)
	GetKey() []byte
	SetKey([]byte)
	CipherDataIface

	MaterialDescription
}

// DecrypterKeyProvider provides an interface to which when grabbing objects
// from S3 we only need the GetDecryptedKey method.
type DecrypterKeyProvider interface {
	GetDecryptedKey([]byte) ([]byte, error)
}

func generateBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

// BaseKeyProvider a wrapper struct that returns a KeyProvider
type BaseKeyProvider struct {
	KeyProvider
}

// GetKeyProvider returns a key provider when dealing with encoding and
// decoding the envelope.
func (kp *BaseKeyProvider) GetKeyProvider() KeyProvider {
	return kp.KeyProvider
}
