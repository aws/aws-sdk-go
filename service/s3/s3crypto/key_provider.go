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
	CipherDataGenerator
	CipherDataHandler
	// CipherDataIface is used for populating the envelope data during
	// encryption.
	CipherDataMetadata
	// MaterialDescription is used to distinguish the materials for both
	// encryption and decryption.
	MaterialDescription
}

// CipherDataGenerator is an interface that deals with key and iv generation
type CipherDataGenerator interface {
	// GenerateKey generates a key of n bytes
	GenerateKey(int) ([]byte, error)
	// GenerateIV generates an iv of n bytes
	GenerateIV(int) ([]byte, error)
}

// CipherDataHandler is an interface of getters and setters dealing with
// keys and ivs.
type CipherDataHandler interface {
	// GetEncryptedKey encrypts and returns the encrypted key
	GetEncryptedKey(key []byte) ([]byte, error)
	GetIV() []byte
	SetIV([]byte)
	GetKey() []byte
	SetKey([]byte)
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
