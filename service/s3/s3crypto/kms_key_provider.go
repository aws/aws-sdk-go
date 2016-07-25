package s3crypto

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/kms"
)

// KMSKeyProvider will make calls to KMS to get the masterkey
type KMSKeyProvider struct {
	kms          *kms.KMS
	key          []byte
	encryptedKey []byte
	iv           []byte
	cmkID        *string

	CipherData
	MaterialDescription
}

// NewKMSKeyProvider builds a new KMS key provider using the customer key ID and material
// description.
//
// Example:
//	sess := session.New(&aws.Config{})
//	cmkID := "arn to key"
//	matdesc := s3crypto.NewJSONMatDesc()
//	kp, err := s3crypto.NewKMSKeyProvider(sess, cmkID, matdesc)
func NewKMSKeyProvider(prov client.ConfigProvider, cmkID string, matdesc MaterialDescription) (KeyProvider, error) {
	if matdesc == nil {
		matdesc = &JSONMatDesc{}
	}
	matdesc.Set("kms_cmk_id", cmkID)

	kp := &KMSKeyProvider{
		MaterialDescription: matdesc,
		kms:                 kms.New(prov),
		cmkID:               &cmkID,
	}
	kp.Algorithm = "kms"
	return kp, nil
}

// NewKMSKeyProviderWithMatDesc initializes a KMS keyprovider with a material description. This
// is used with Decrypting kms content, due to the cmkID being in the material description.
func NewKMSKeyProviderWithMatDesc(prov client.ConfigProvider, matdesc string) (KeyProvider, error) {
	m := &JSONMatDesc{}
	err := m.DecodeDescription([]byte(matdesc))
	if err != nil {
		return nil, err
	}

	cmkID, ok := m.Get("kms_cmk_id")
	if !ok {
		return nil, awserr.New("MissingCMKID", "Material description is missing CMK ID", nil)
	}

	kp := &KMSKeyProvider{}
	kp.MaterialDescription = m
	kp.kms = kms.New(prov)
	kp.cmkID = &cmkID
	kp.Algorithm = "kms"
	return kp, nil
}

// GetKey getter for key
func (kp *KMSKeyProvider) GetKey() []byte {
	return kp.key
}

// SetKey setter for key
func (kp *KMSKeyProvider) SetKey(key []byte) {
	kp.key = key
}

// GetIV getter for IV
func (kp *KMSKeyProvider) GetIV() []byte {
	return kp.iv
}

// SetIV setter for IV
func (kp *KMSKeyProvider) SetIV(iv []byte) {
	kp.iv = iv
}

// GetEncryptedKey getter for encrypted key. The encrypted key is set
// when GenerateKey is called.
func (kp *KMSKeyProvider) GetEncryptedKey(key []byte) ([]byte, error) {
	return kp.encryptedKey, nil
}

// GetDecryptedKey makes a call to KMS to decrypt the key.
func (kp *KMSKeyProvider) GetDecryptedKey(key []byte) ([]byte, error) {
	matdesc := kp.MaterialDescription.GetData()
	out, err := kp.kms.Decrypt(&kms.DecryptInput{
		EncryptionContext: matdesc,
		CiphertextBlob:    key,
		GrantTokens:       []*string{},
	})
	if err != nil {
		return nil, err
	}
	return out.Plaintext, nil
}

// GenerateKey makes a call to KMS to generate a data key, Upon making
// the call, it also sets the encrypted key.
func (kp *KMSKeyProvider) GenerateKey(n int) ([]byte, error) {
	out, err := kp.kms.GenerateDataKey(&kms.GenerateDataKeyInput{
		EncryptionContext: map[string]*string{
			"kms_cmk_id": kp.cmkID,
		},
		KeyId:   kp.cmkID,
		KeySpec: aws.String("AES_256"),
	})
	if err != nil {
		return nil, err
	}
	kp.encryptedKey = out.CiphertextBlob
	return out.Plaintext, nil
}

// GenerateIV generates an IV of n bytes
func (kp *KMSKeyProvider) GenerateIV(n int) ([]byte, error) {
	return generateBytes(n), nil
}
