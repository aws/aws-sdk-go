package s3crypto

import (
	"encoding/base64"
	"fmt"
	"io"
)

// CryptoMode placeholder
type CryptoMode interface {
	EncryptContents(io.Writer, io.Reader) error
	DecryptContents([]byte, []byte, io.ReadCloser) (io.ReadCloser, error)
	GetKeyProvider() KeyProvider
}

// DecryptMode is meant to used only in reading objects from s3
type DecryptMode interface {
	DecryptContents([]byte, []byte, io.ReadCloser) (io.ReadCloser, error)
	GetKeyProvider() KeyProvider
}

func modeFactory(env *Envelope, cfg Config) (DecryptMode, error) {
	kp, err := keyProviderFactory(env, cfg)
	fmt.Println("ERROR 1", err)
	if err != nil {
		return nil, err
	}

	err = DecodeMeta(env, kp)
	fmt.Println("ERROR 2", err)
	if err != nil {
		return nil, err
	}

	kp.SetKey([]byte(env.CipherKey))
	kp.SetIV([]byte(env.IV))

	cek, err := cekFactory(env, kp)
	if err != nil {
		return nil, err
	}
	return &decryptionMode{kp: kp, cek: cek}, nil
}

// wrapFactory will build a new CryptoMode based off the wrapping algorithm
// TODO: Have the Cipher constructors return errs instead of panicing on invalid
// key and iv lengths
func keyProviderFactory(env *Envelope, cfg Config) (KeyProvider, error) {

	switch env.WrapAlg {
	case "kms":
		return NewKMSKeyProvider(cfg.KMSSession)
	case "rsa":
	case "ecb", "":
		fmt.Println("MASTER", cfg.MasterKey)
		cipher, err := NewAESECB(cfg.MasterKey)
		if err != nil {
			return nil, err
		}
		return NewSymmetricKeyProvider(cipher), nil
	case "aeswrap":
	}
	return nil, nil
}

func cekFactory(env *Envelope, kp KeyProvider) (Decrypter, error) {
	switch env.CEKAlg {
	case "AES/CBC/PKCS5Padding", "":
		return NewAESCBC(kp)
	}
	return nil, nil
}

// EncodeMeta will return the meta object to be saved
func EncodeMeta(reader HashReader, kp KeyProvider) (Envelope, error) {
	iv := base64.StdEncoding.EncodeToString(kp.GetIV())
	keyBytes, err := kp.GetEncryptedKey()
	if err != nil {
		return Envelope{}, nil
	}
	key := base64.StdEncoding.EncodeToString(keyBytes)

	md5 := []byte{}
	contentLength := 0
	if reader != nil {
		md5 = reader.GetValue()
		contentLength = reader.GetContentLength()
	}

	md5Str := base64.StdEncoding.EncodeToString(md5)

	return Envelope{
		CipherKey:             key,
		IV:                    iv,
		MatDesc:               "{}",
		WrapAlg:               "keywrap", // TODO: Remove this hard coded value and replace with MasterKey.DescribeAlg()
		CEKAlg:                "cbc",
		TagLen:                "0",
		UnencryptedMD5:        md5Str,
		UnencryptedContentLen: fmt.Sprintf("%d", contentLength),
	}, nil
}

// DecodeMeta will return the metaobject with keys as decrypted values, if they were encrypted
// or base64 encoded.
func DecodeMeta(env *Envelope, kp KeyProvider) error {
	key, err := base64.StdEncoding.DecodeString(env.CipherKey)
	if err != nil {
		return err
	}
	kp.SetEncryptedKey(key)

	keyBytes, err := kp.GetDecryptedKey()
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
