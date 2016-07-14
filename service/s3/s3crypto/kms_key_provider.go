package s3crypto

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

// KMSKeyProvider will make calls to KMS to get the masterkey
// TODO: Also need a KMS Wrap interface which just returns
// GetMasterKey with the correct key id
type KMSKeyProvider struct {
	kms          *kms.KMS
	key          []byte
	encryptedKey []byte
	iv           []byte
	matdesc      map[string]interface{}
}

// NewKMSKeyProvider placeholder
func NewKMSKeyProvider(sess *session.Session, matdesc string) (KeyProvider, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(matdesc), &m)
	if err != nil {
		return nil, err
	}

	kp := &KMSKeyProvider{}
	kp.matdesc = m
	kp.kms = kms.New(sess)
	return kp, nil
}

// GetKey returns key
func (kp *KMSKeyProvider) GetKey() []byte {
	return kp.key
}

// SetKey returns key
func (kp *KMSKeyProvider) SetKey(key []byte) {
	kp.key = key
}

// GetIV returns key
func (kp *KMSKeyProvider) GetIV() []byte {
	return kp.iv
}

// SetIV sets iv
func (kp *KMSKeyProvider) SetIV(iv []byte) {
	kp.iv = iv
}

// GetEncryptedKey returns key
func (kp *KMSKeyProvider) GetEncryptedKey(key []byte) ([]byte, error) {
	return kp.encryptedKey, nil
}

// GetDecryptedKey makes a call to KMS to decrypt the key
func (kp *KMSKeyProvider) GetDecryptedKey(key []byte) ([]byte, error) {
	kmsID := kp.matdesc["kms_cmk_id"].(string)
	out, err := kp.kms.Decrypt(&kms.DecryptInput{
		EncryptionContext: map[string]*string{
			"kms_cmk_id": &kmsID,
		},
		CiphertextBlob: key,
		GrantTokens:    []*string{},
	})
	if err != nil {
		return nil, err
	}
	return out.Plaintext, nil
}

// GenerateKey place holder
func (kp *KMSKeyProvider) GenerateKey(n int) ([]byte, error) {
	kmsID := aws.String(kp.matdesc["kms_cmk_id"].(string))
	out, err := kp.kms.GenerateDataKey(&kms.GenerateDataKeyInput{
		EncryptionContext: map[string]*string{
			"kms_cmk_id": kmsID,
		},
		KeyId: kmsID,
	})
	if err != nil {
		return nil, err
	}
	kp.encryptedKey = out.CiphertextBlob
	return out.Plaintext, nil
}

// GenerateIV placeholder
func (kp *KMSKeyProvider) GenerateIV(n int) ([]byte, error) {
	return generateBytes(n), nil
}
