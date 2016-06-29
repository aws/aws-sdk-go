package s3crypto

import (
	"github.com/aws/aws-sdk-go/service/kms"
)

// KMSMode will make calls to KMS to get the masterkey
type KMSMode struct {
	kms *kms.KMS
}

// Encrypt placeholder
//func (cipher *KMSCipher) Encrypt(src io.Reader) io.Reader {
//}

// EncryptContents placeholder
func (mode *KMSMode) EncryptContents(dst io.Writer, src io.Reader) error {
}
