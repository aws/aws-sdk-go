package s3crypto

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
)

// CryptoMode placeholder
type CryptoMode interface {
	EncryptContents(io.Writer, io.Reader) error
	DecryptMode
	GetKey() []byte
	GetIV() []byte
	Wrap
}

// DecryptMode is meant to used only in reading objects from s3
type DecryptMode interface {
	DecryptContents([]byte, []byte, io.ReadCloser) (io.ReadCloser, error)
}

func modeFactory(env *Envelope, mode CryptoMode) (DecryptMode, error) {
	// If the envelope version is 1, then we
	// default to whatever mode the user specified when constructing the
	// client.
	if env.version == 1 {
		return mode, nil
	}
	wrap, err := wrapFactory(env, mode)
	if err != nil {
		return nil, err
	}
	cek, err := cekFactory(env, &SymmetricKeyProvider{[]byte(env.CipherKey), []byte(env.IV)})
	if err != nil {
		return nil, err
	}
	return &decryptionMode{wrap: wrap, cek: cek}, nil
}

// wrapFactory will build a new CryptoMode based off the wrapping algorithm
func wrapFactory(env *Envelope, mode CryptoMode) (Wrap, error) {

	switch env.WrapAlg {
	case "kms":
	case "rsa":
	case "ecb":
		return NewAESECB(mode.GetMasterKey())
	case "aeswrap":
	}
	return nil, nil
}

func cekFactory(env *Envelope, kp *SymmetricKeyProvider) (Decrypter, error) {
	switch env.CEKAlg {
	case "AES/CBC/PKCS5Padding":
		return NewAESCBC(kp)
	}
	return nil, nil
}

// EncodeMeta will return the meta object to be saved
func EncodeMeta(reader HashReader, mode CryptoMode) (Envelope, error) {
	iv := base64.StdEncoding.EncodeToString(mode.GetIV())
	dst := mode.Encrypt(bytes.NewReader(mode.GetKey()))

	b, err := ioutil.ReadAll(dst)
	if err != nil {
		return Envelope{}, err
	}
	key := base64.StdEncoding.EncodeToString(b)
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
func DecodeMeta(env *Envelope, mode CryptoMode) error {
	key, err := base64.StdEncoding.DecodeString(env.CipherKey)
	if err != nil {
		return err
	}
	dst := mode.Decrypt(bytes.NewReader(key))

	b, err := ioutil.ReadAll(dst)
	if err != nil {
		return err
	}

	env.CipherKey = string(b)
	iv, err := base64.StdEncoding.DecodeString(env.IV)
	if err != nil {
		return err
	}
	env.IV = string(iv)
	return nil
}
