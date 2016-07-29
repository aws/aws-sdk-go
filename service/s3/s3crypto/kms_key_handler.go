package s3crypto

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/kms"
)

// KMSKeyHandler will make calls to KMS to get the masterkey
type KMSKeyHandler struct {
	kms          *kms.KMS
	encryptedKey []byte
	cmkID        *string

	CipherData
}

// NewKMSEncryptHandler builds a new KMS key provider using the customer key ID and material
// description.
//
// Example:
//	sess := session.New(&aws.Config{})
//	cmkID := "arn to key"
//	matdesc := s3crypto.NewJSONMatDesc()
//	kp, err := s3crypto.NewKMSKeyHandler(sess, cmkID, matdesc)
func NewKMSEncryptHandler(prov client.ConfigProvider, cmkID string, matdesc MaterialDescription) (CipherDataHandler, error) {
	if matdesc == nil {
		matdesc = MaterialDescription{}
	}
	matdesc["kms_cmk_id"] = &cmkID

	kp := &KMSKeyHandler{
		kms:   kms.New(prov),
		cmkID: &cmkID,
	}
	kp.CipherData.WrapAlgorithm = "kms"
	kp.CipherData.MaterialDescription = matdesc
	return kp, nil
}

// NewKMSDecryptHandler initializes a KMS keyprovider with a material description. This
// is used with Decrypting kms content, due to the cmkID being in the material description.
func NewKMSDecryptHandler(prov client.ConfigProvider, matdesc string) (CipherDataHandler, error) {
	m := MaterialDescription{}
	err := m.decodeDescription([]byte(matdesc))
	if err != nil {
		return nil, err
	}

	cmkID, ok := m["kms_cmk_id"]
	if !ok {
		return nil, awserr.New("MissingCMKIDError", "Material description is missing CMK ID", nil)
	}

	kp := &KMSKeyHandler{}
	kp.CipherData.MaterialDescription = m
	kp.kms = kms.New(prov)
	kp.cmkID = cmkID
	kp.WrapAlgorithm = "kms"
	return kp, nil
}

// EncryptKey getter for encrypted key. The encrypted key is set
// when GenerateKey is called.
func (kp *KMSKeyHandler) EncryptKey(key []byte) ([]byte, error) {
	return kp.encryptedKey, nil
}

// DecryptKey makes a call to KMS to decrypt the key.
func (kp *KMSKeyHandler) DecryptKey(key []byte) ([]byte, error) {
	matdesc := kp.CipherData.MaterialDescription.GetData()
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

// GenerateCipherData makes a call to KMS to generate a data key, Upon making
// the call, it also sets the encrypted key.
func (kp *KMSKeyHandler) GenerateCipherData(keySize, ivSize int) (CipherData, error) {
	out, err := kp.kms.GenerateDataKey(&kms.GenerateDataKeyInput{
		EncryptionContext: map[string]*string{
			"kms_cmk_id": kp.cmkID,
		},
		KeyId:   kp.cmkID,
		KeySpec: aws.String("AES_256"),
	})
	if err != nil {
		return CipherData{}, err
	}
	kp.encryptedKey = out.CiphertextBlob
	iv := generateBytes(ivSize)
	cd := CipherData{
		Key:                 out.Plaintext,
		IV:                  iv,
		WrapAlgorithm:       "kms",
		MaterialDescription: kp.CipherData.MaterialDescription,
	}
	return cd, nil
}
