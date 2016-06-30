package s3crypto

import (
	"io"
)

type decryptionMode struct {
	kp  KeyProvider
	cek Decrypter
}

// DecryptContents does not use key and iv because the decrypter was instantiated earlier
func (mode *decryptionMode) DecryptContents(key, iv []byte, src io.ReadCloser) (io.ReadCloser, error) {
	reader := mode.cek.Decrypt(src)
	return &CryptoReadCloser{Body: src, Decrypter: reader}, nil
}
