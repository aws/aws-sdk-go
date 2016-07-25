package s3crypto

import (
	"encoding/base64"
	"io"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// CryptoMode is an abstraction layer that deals with encryption of
// the contents and which key provider to use.
type CryptoMode interface {
	EncryptContents(io.Writer, io.Reader) error
	DecryptContents(KeyProvider, io.ReadCloser) (io.ReadCloser, error)
	GetKeyProvider() KeyProvider
	CipherDataMetadata
}

// CipherDataMetadata  is used for when populating the envelope details upon
// encryption.
type CipherDataMetadata interface {
	GetCipherName() string
	GetTagLen() string
}

// DecryptMode is meant to used only in reading objects from s3
type DecryptMode interface {
	DecryptContents(KeyProvider, io.ReadCloser) (io.ReadCloser, error)
	GetKeyProvider() KeyProvider
}

func modeForEnvelope(env *Envelope, cfg Config) (DecryptMode, error) {
	kp, err := keyProviderForEnvelope(env, cfg)
	if err != nil {
		return nil, err
	}

	err = DecodeMeta(env, kp)
	if err != nil {
		return nil, err
	}

	kp.SetKey([]byte(env.CipherKey))
	kp.SetIV([]byte(env.IV))

	cek, err := cekForEnvelope(env, kp)
	if err != nil {
		return nil, err
	}
	return &decryptionMode{kp: kp, cek: cek}, nil
}

// wrapFactory will build a new CryptoMode based off the wrapping algorithm
// TODO: Have the Cipher constructors return errs instead of panicing on invalid
// key and iv lengths
func keyProviderForEnvelope(env *Envelope, cfg Config) (KeyProvider, error) {

	switch env.WrapAlg {
	case "kms":
		return NewKMSKeyProviderWithMatDesc(cfg.KMSSession, env.MatDesc)
	}
	return nil, awserr.New(
		"Invali1WrapAlgorithmError",
		"wrap algorithm isn't supported, "+env.WrapAlg,
		nil,
	)
}

// AESGCMNoPadding is the constant value that is used to specify
// the CEK algorithm consiting of AES GCM with no padding.
const AESGCMNoPadding = "AES/GCM/NoPadding"

func cekForEnvelope(env *Envelope, kp KeyProvider) (Decrypter, error) {
	switch env.CEKAlg {
	case AESGCMNoPadding:
		return NewAESGCM(kp)
	}
	return nil, awserr.New(
		"InvalidCEK",
		"cek algorithm isn't supported, "+env.CEKAlg,
		nil,
	)
}

// EncodeMeta will return the meta object to be saved
func EncodeMeta(reader HashReader, mode CryptoMode) (Envelope, error) {
	kp := mode.GetKeyProvider()
	iv := base64.StdEncoding.EncodeToString(kp.GetIV())
	keyBytes, err := kp.GetEncryptedKey(kp.GetKey())
	if err != nil {
		return Envelope{}, err
	}
	key := base64.StdEncoding.EncodeToString(keyBytes)

	md5 := reader.GetValue()
	contentLength := reader.GetContentLength()

	md5Str := base64.StdEncoding.EncodeToString(md5)
	matdesc, err := kp.EncodeDescription()
	if err != nil {
		return Envelope{}, err
	}

	return Envelope{
		CipherKey:             key,
		IV:                    iv,
		MatDesc:               string(matdesc),
		WrapAlg:               kp.GetCipherName(),
		CEKAlg:                mode.GetCipherName(),
		TagLen:                mode.GetTagLen(),
		UnencryptedMD5:        md5Str,
		UnencryptedContentLen: strconv.FormatInt(contentLength, 10),
	}, nil
}

// DecodeMeta will return the metaobject with keys as decrypted values, if they were encrypted
// or base64 encoded.
func DecodeMeta(env *Envelope, kp KeyProvider) error {
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
}
