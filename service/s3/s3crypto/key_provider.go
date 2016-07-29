package s3crypto

import "crypto/rand"

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
