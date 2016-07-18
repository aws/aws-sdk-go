package s3crypto

import (
	"io"
)

type decryptionMode struct {
	kp  KeyProvider
	cek Decrypter
}

// DecryptContents does not use key and iv because the decrypter was instantiated earlier
func (mode *decryptionMode) DecryptContents(kp KeyProvider, src io.ReadCloser) (io.ReadCloser, error) {
	reader := mode.cek.Decrypt(src)
	return &CryptoReadCloser{Body: src, Decrypter: reader}, nil
}

func (mode *decryptionMode) GetKeyProvider() KeyProvider {
	return mode.kp
}
