package s3crypto

import "crypto/rand"

/*// KeyProvider is used to help generate keys and ivs. This allows for crypto based
// algorithms or services to be used.
type KeyProvider interface {
	EncrypterKeyProvider
}*/

// KeyProviderEncrypter test
type KeyProviderEncrypter interface {
}

/*// EncrypterKeyProvider handles how to generate keys and ivs for the content encryption.
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
}*/

// CipherDataHandler is an interface that deals with key and iv generation
type CipherDataHandler interface {
	GenerateCipherData(int, int) (CipherData, error)
	EncryptKey(key []byte) ([]byte, error)
	DecryptKey([]byte) ([]byte, error)
}

func generateBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

// BaseKeyProvider a wrapper struct that returns a KeyProvider
/*type BaseKeyProvider struct {
	KeyProvider
}

// GetKeyProvider returns a key provider when dealing with encoding and
// decoding the envelope.
func (kp *BaseKeyProvider) GetKeyProvider() KeyProvider {
	return kp.KeyProvider
}*/
