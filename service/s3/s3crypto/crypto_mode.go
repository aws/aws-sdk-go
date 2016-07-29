package s3crypto

import (
	"encoding/base64"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// CipherDataMetadata  is used for when populating the envelope details upon
// encryption.
type CipherDataMetadata interface {
	GetCipherName() string
	GetTagLen() string
}

func contentCipherFromEnvelope(env *Envelope, cfg Config) (ContentCipher, error) {
	wrap, err := wrapFromEnvelope(env, cfg)
	if err != nil {
		return nil, err
	}

	return cekFromEnvelope(env, wrap)
}

func wrapFromEnvelope(env *Envelope, cfg Config) (CipherDataHandler, error) {
	switch env.WrapAlg {
	case "kms":
		return NewKMSKeyProviderDecrypter(cfg.KMSSession, env.MatDesc)
	}
	return nil, awserr.New(
		"InvalidWrapAlgorithmError",
		"wrap algorithm isn't supported, "+env.WrapAlg,
		nil,
	)
}

// AESGCMNoPadding is the constant value that is used to specify
// the CEK algorithm consiting of AES GCM with no padding.
const AESGCMNoPadding = "AES/GCM/NoPadding"

func cekFromEnvelope(env *Envelope, cdh CipherDataHandler) (ContentCipher, error) {
	key, err := base64.StdEncoding.DecodeString(env.CipherKey)
	if err != nil {
		return nil, err
	}

	iv, err := base64.StdEncoding.DecodeString(env.IV)
	if err != nil {
		return nil, err
	}
	key, err = cdh.DecryptKey(key)
	if err != nil {
		return nil, err
	}

	switch env.CEKAlg {
	case AESGCMNoPadding:
		cd := CipherData{
			Key: key,
			IV:  iv,
		}
		return newAESGCMContentCipher(cd, cdh)
	}
	return nil, awserr.New(
		"InvalidCEKAlgorithmError",
		"cek algorithm isn't supported, "+env.CEKAlg,
		nil,
	)
}

func encodeMeta(reader HashReader, handler CipherDataHandler, cd *CipherData) (Envelope, error) {
	iv := base64.StdEncoding.EncodeToString(cd.IV)
	keyBytes, err := handler.EncryptKey(cd.Key)
	if err != nil {
		return Envelope{}, err
	}
	key := base64.StdEncoding.EncodeToString(keyBytes)

	md5 := reader.GetValue()
	contentLength := reader.GetContentLength()

	md5Str := base64.StdEncoding.EncodeToString(md5)
	matdesc, err := cd.encodeDescription()
	if err != nil {
		return Envelope{}, err
	}

	return Envelope{
		CipherKey:             key,
		IV:                    iv,
		MatDesc:               string(matdesc),
		WrapAlg:               cd.WrapAlgorithm,
		CEKAlg:                cd.CEKAlgorithm,
		TagLen:                cd.TagLength,
		UnencryptedMD5:        md5Str,
		UnencryptedContentLen: strconv.FormatInt(contentLength, 10),
	}, nil
}

// DecodeMeta will return the metaobject with keys as decrypted values, if they were encrypted
// or base64 encoded.
/*func DecodeMeta(env *Envelope, cd CipherData) error {
	key, err := base64.StdEncoding.DecodeString(env.CipherKey)
	if err != nil {
		return err
	}

	keyBytes, err := kp.GetDecryptedKey(key)
	if err != nil {
		return err
	}

	iv, err := base64.StdEncoding.DecodeString(env.IV)
	if err != nil {
		return err
	}

	env.CipherKey = string(keyBytes)
	env.IV = string(iv)
	return nil
}*/
