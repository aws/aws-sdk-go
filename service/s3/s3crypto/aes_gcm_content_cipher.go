package s3crypto

import (
	"io"
)

const (
	gcmKeySize   = 32
	gcmNonceSize = 12
)

type gcmContentCipherBuilder struct {
	handler CipherDataHandler
}

// AESGCMContentCipherBuilder returns a new encryption only mode structure with a specific cipher
// for the master key
func AESGCMContentCipherBuilder(handler CipherDataHandler) ContentCipherBuilder {
	return gcmContentCipherBuilder{handler}
}

func (builder gcmContentCipherBuilder) NewEncryptor() (ContentCipher, error) {
	cd, err := builder.handler.GenerateCipherData(gcmKeySize, gcmNonceSize)
	if err != nil {
		return nil, err
	}

	return newAESGCMContentCipher(cd, builder.handler)
}

func newAESGCMContentCipher(cd CipherData, handler CipherDataHandler) (ContentCipher, error) {
	cd.CEKAlgorithm = AESGCMNoPadding
	cd.TagLength = "128"

	cipher, err := NewAESGCM(cd)
	if err != nil {
		return nil, err
	}

	return &AESGCMContentCipher{
		CipherData: cd,
		Cipher:     cipher,
		handler:    handler,
	}, nil
}

// AESGCMContentCipher will use AES GCM for the main cipher.
type AESGCMContentCipher struct {
	CipherData // Embedding so we don't need duplicate GetCipherData methods
	Cipher     Cipher
	handler    CipherDataHandler
}

// EncryptContents will generate a random key and iv and encrypt the data using cbc
func (cc *AESGCMContentCipher) EncryptContents(src io.Reader) (io.Reader, error) {
	return cc.Cipher.Encrypt(src), nil
}

// DecryptContents will use the symmetric key provider to instantiate a new GCM cipher.
// We grab a decrypt reader from gcm and wrap it in a CryptoReadCloser. The only error
// expected here is when the key or iv is of invalid length.
func (cc *AESGCMContentCipher) DecryptContents(src io.ReadCloser) (io.ReadCloser, error) {
	reader := cc.Cipher.Decrypt(src)
	return &CryptoReadCloser{Body: src, Decrypter: reader}, nil
}

// GetHandler returns the handler
func (cc AESGCMContentCipher) GetHandler() CipherDataHandler {
	return cc.handler
}
