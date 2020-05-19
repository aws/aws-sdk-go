package s3crypto

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

const (
	// KMSWrap is a constant used during decryption to build a KMS key handler.
	KMSWrap = "kms"

	// KMSContextWrap is a constant used during decryption to build a kms+context key handler
	KMSContextWrap = "kms+context"
)

// kmsKeyHandler will make calls to KMS to get the masterkey
type kmsKeyHandler struct {
	kms         kmsiface.KMSAPI
	cmkID       *string
	withContext bool

	CipherData
}

// NewKMSKeyGenerator builds a new KMS key provider using the customer key ID and material
// description.
//
// Example:
//	sess := session.New(&aws.Config{})
//	cmkID := "arn to key"
//	matdesc := s3crypto.MaterialDescription{}
//	handler := s3crypto.NewKMSKeyGenerator(kms.New(sess), cmkID)
//
// deprecated: See NewKMSContextKeyGenerator
func NewKMSKeyGenerator(kmsClient kmsiface.KMSAPI, cmkID string) CipherDataGenerator {
	return NewKMSKeyGeneratorWithMatDesc(kmsClient, cmkID, MaterialDescription{})
}

// NewKMSContextKeyGenerator builds a new kms+context key provider using the customer key ID and material
// description.
//
// Example:
//	sess := session.New(&aws.Config{})
//	cmkID := "arn to key"
//	matdesc := s3crypto.MaterialDescription{}
//	handler := s3crypto.NewKMSContextKeyGenerator(kms.New(sess), cmkID)
func NewKMSContextKeyGenerator(client kmsiface.KMSAPI, cmkID string) CipherDataGeneratorWithCEKAlg {
	return NewKMSContextKeyGeneratorWithMatDesc(client, cmkID, MaterialDescription{})
}

func newKMSKeyHandler(client kmsiface.KMSAPI, cmkID string, withContext bool, matdesc MaterialDescription) *kmsKeyHandler {
	// These values are read only making them thread safe
	kp := &kmsKeyHandler{
		kms:         client,
		cmkID:       &cmkID,
		withContext: withContext,
	}

	if matdesc == nil {
		matdesc = MaterialDescription{}
	}

	// These values are read only making them thread safe
	if kp.withContext {
		kp.CipherData.WrapAlgorithm = KMSContextWrap
	} else {
		matdesc["kms_cmk_id"] = &cmkID
		kp.CipherData.WrapAlgorithm = KMSWrap
	}
	kp.CipherData.MaterialDescription = matdesc
	return kp
}

// NewKMSKeyGeneratorWithMatDesc builds a new KMS key provider using the customer key ID and material
// description.
//
// Example:
//	sess := session.New(&aws.Config{})
//	cmkID := "arn to key"
//	matdesc := s3crypto.MaterialDescription{}
//	handler := s3crypto.NewKMSKeyGeneratorWithMatDesc(kms.New(sess), cmkID, matdesc)
//
// deprecated: See NewKMSContextKeyGeneratorWithMatDesc
func NewKMSKeyGeneratorWithMatDesc(kmsClient kmsiface.KMSAPI, cmkID string, matdesc MaterialDescription) CipherDataGenerator {
	return newKMSKeyHandler(kmsClient, cmkID, false, matdesc)
}

// NewKMSContextKeyGeneratorWithMatDesc builds a new kms+context key provider using the customer key ID and material
// description.
//
// Example:
//	sess := session.New(&aws.Config{})
//	cmkID := "arn to key"
//	matdesc := s3crypto.MaterialDescription{}
//	handler := s3crypto.NewKMSKeyGeneratorWithMatDesc(kms.New(sess), cmkID, matdesc)
func NewKMSContextKeyGeneratorWithMatDesc(kmsClient kmsiface.KMSAPI, cmkID string, matdesc MaterialDescription) CipherDataGeneratorWithCEKAlg {
	return newKMSKeyHandler(kmsClient, cmkID, true, matdesc)
}

// NewKMSWrapEntry builds returns a new KMS key provider and its decrypt handler.
//
// Example:
//	sess := session.New(&aws.Config{})
//	customKMSClient := kms.New(sess)
//	decryptHandler := s3crypto.NewKMSWrapEntry(customKMSClient)
//
//	svc := s3crypto.NewDecryptionClient(sess, func(svc *s3crypto.DecryptionClient) {
//		svc.WrapRegistry[s3crypto.KMSWrap] = decryptHandler
//	}))
//
// deprecated: See NewKMSContextWrapEntry
func NewKMSWrapEntry(kmsClient kmsiface.KMSAPI) WrapEntry {
	// These values are read only making them thread safe
	kp := &kmsKeyHandler{
		kms: kmsClient,
	}

	return kp.decryptHandler
}

// NewKMSContextWrapEntry builds returns a new KMS key provider and its decrypt handler.
//
// Example:
//	sess := session.New(&aws.Config{})
//	customKMSClient := kms.New(sess)
//	decryptHandler := s3crypto.NewKMSContextWrapEntry(customKMSClient)
//
//	svc := s3crypto.NewDecryptionClient(sess, func(svc *s3crypto.DecryptionClient) {
//		svc.WrapRegistry[s3crypto.KMSContextWrap] = decryptHandler
//	}))
func NewKMSContextWrapEntry(kmsClient kmsiface.KMSAPI) WrapEntry {
	// These values are read only making them thread safe
	kp := &kmsKeyHandler{
		kms:         kmsClient,
		withContext: true,
	}

	return kp.decryptHandler
}

// decryptHandler initializes a KMS keyprovider with a material description. This
// is used with Decrypting kms content, due to the cmkID being in the material description.
func (kp kmsKeyHandler) decryptHandler(env Envelope) (CipherDataDecrypter, error) {
	m := MaterialDescription{}
	err := m.decodeDescription([]byte(env.MatDesc))
	if err != nil {
		return nil, err
	}

	cmkID, ok := m["kms_cmk_id"]
	if !kp.withContext && !ok {
		return nil, awserr.New("MissingCMKIDError", "Material description is missing CMK ID", nil)
	}

	kp.CipherData.MaterialDescription = m
	kp.cmkID = cmkID
	kp.WrapAlgorithm = KMSWrap
	if kp.withContext {
		kp.WrapAlgorithm = KMSContextWrap
	}
	return &kp, nil
}

// DecryptKey makes a call to KMS to decrypt the key.
func (kp *kmsKeyHandler) DecryptKey(key []byte) ([]byte, error) {
	return kp.DecryptKeyWithContext(aws.BackgroundContext(), key)
}

// DecryptKeyWithContext makes a call to KMS to decrypt the key with request context.
func (kp *kmsKeyHandler) DecryptKeyWithContext(ctx aws.Context, key []byte) ([]byte, error) {
	out, err := kp.kms.DecryptWithContext(ctx,
		&kms.DecryptInput{
			EncryptionContext: kp.CipherData.MaterialDescription,
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
func (kp *kmsKeyHandler) GenerateCipherData(keySize, ivSize int) (CipherData, error) {
	return kp.GenerateCipherDataWithContext(aws.BackgroundContext(), keySize, ivSize)
}

func (kp kmsKeyHandler) GenerateCipherDataWithCEKAlg(keySize, ivSize int, cekAlgorithm string) (CipherData, error) {
	return kp.GenerateCipherDataWithCEKAlgWithContext(aws.BackgroundContext(), keySize, ivSize, cekAlgorithm)
}

// GenerateCipherDataWithContext makes a call to KMS to generate a data key,
// Upon making the call, it also sets the encrypted key.
func (kp *kmsKeyHandler) GenerateCipherDataWithContext(ctx aws.Context, keySize, ivSize int) (CipherData, error) {
	return kp.GenerateCipherDataWithCEKAlgWithContext(ctx, keySize, ivSize, "")
}

func (kp kmsKeyHandler) GenerateCipherDataWithCEKAlgWithContext(ctx aws.Context, keySize int, ivSize int, cekAlgorithm string) (CipherData, error) {
	md := kp.CipherData.MaterialDescription

	wrapAlgorithm := KMSWrap
	if kp.withContext {
		wrapAlgorithm = KMSContextWrap
		if len(cekAlgorithm) == 0 {
			return CipherData{}, fmt.Errorf("CEK algorithm identifier must not be empty")
		}
		md["aws:"+cekAlgorithmHeader] = &cekAlgorithm
	}

	out, err := kp.kms.GenerateDataKeyWithContext(ctx,
		&kms.GenerateDataKeyInput{
			EncryptionContext: md,
			KeyId:             kp.cmkID,
			KeySpec:           aws.String("AES_256"),
		})
	if err != nil {
		return CipherData{}, err
	}

	iv, err := generateBytes(ivSize)
	if err != nil {
		return CipherData{}, err
	}

	cd := CipherData{
		Key:                 out.Plaintext,
		IV:                  iv,
		WrapAlgorithm:       wrapAlgorithm,
		MaterialDescription: md,
		EncryptedKey:        out.CiphertextBlob,
	}
	return cd, nil
}

func (kp *kmsKeyHandler) isUsingDeprecatedFeatures() error {
	if !kp.withContext {
		return errDeprecatedCipherDataGenerator
	}
	return nil
}
