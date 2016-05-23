package s3crypto

import (
	"bytes"
	"encoding/hex"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AES CBC
func TestAES_CBC_NIST_CBCGFSbox256_case_1(t *testing.T) {
	iv, _ := hex.DecodeString("00000000000000000000000000000000")
	key, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	plaintext, _ := hex.DecodeString("014730f80ac625fe84f026c60bfd547d")
	expected, _ := hex.DecodeString("5c9d844ed46f9885085e5d6a4f94c7d7")
	aescbcEncrypt(t, iv, key, plaintext, expected)
	aescbcDecrypt(t, iv, key, plaintext)
}

func TestAES_CBC_NIST_CBCVarKey256_case_254(t *testing.T) {
	iv, _ := hex.DecodeString("00000000000000000000000000000000")
	key, _ := hex.DecodeString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe")
	plaintext, _ := hex.DecodeString("00000000000000000000000000000000")
	expected, _ := hex.DecodeString("b07d4f3e2cd2ef2eb545980754dfea0f")
	aescbcEncrypt(t, iv, key, plaintext, expected)
	aescbcDecrypt(t, iv, key, plaintext)
}

func TestAES_CBC_NIST_CBCVarTxt256_case_110(t *testing.T) {
	iv, _ := hex.DecodeString("00000000000000000000000000000000")
	key, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	plaintext, _ := hex.DecodeString("fffffffffffffffffffffffffffe0000")
	expected, _ := hex.DecodeString("4b00c27e8b26da7eab9d3a88dec8b031")
	aescbcEncrypt(t, iv, key, plaintext, expected)
	aescbcDecrypt(t, iv, key, plaintext)
}

func TestAES_CBC_NIST_CBCMMT256_case_4(t *testing.T) {
	iv, _ := hex.DecodeString("11958dc6ab81e1c7f01631e9944e620f")
	key, _ := hex.DecodeString("9adc8fbd506e032af7fa20cf5343719de6d1288c158c63d6878aaf64ce26ca85")
	plaintext, _ := hex.DecodeString("c7917f84f747cd8c4b4fedc2219bdbc5f4d07588389d8248854cf2c2f89667a2d7bcf53e73d32684535f42318e24cd45793950b3825e5d5c5c8fcd3e5dda4ce9246d18337ef3052d8b21c5561c8b660e")
	expected, _ := hex.DecodeString("9c99e68236bb2e929db1089c7750f1b356d39ab9d0c40c3e2f05108ae9d0c30b04832ccdbdc08ebfa426b7f5efde986ed05784ce368193bb3699bc691065ac62e258b9aa4cc557e2b45b49ce05511e65")
	aescbcEncrypt(t, iv, key, plaintext, expected)
	aescbcDecrypt(t, iv, key, plaintext)

	aescbcEncryptParts(t, iv, key, plaintext, expected, 16)
	aescbcDecryptParts(t, iv, key, plaintext, 16)
}

func TestAES_CBC_NIST_CBCMMT256_case_9(t *testing.T) {
	iv, _ := hex.DecodeString("e49651988ebbb72eb8bb80bb9abbca34")
	key, _ := hex.DecodeString("87725bd43a45608814180773f0e7ab95a3c859d83a2130e884190e44d14c6996")
	plaintext, _ := hex.DecodeString("bfe5c6354b7a3ff3e192e05775b9b75807de12e38a626b8bf0e12d5fff78e4f1775aa7d792d885162e66d88930f9c3b2cdf8654f56972504803190386270f0aa43645db187af41fcea639b1f8026ccdd0c23e0de37094a8b941ecb7602998a4b2604e69fc04219585d854600e0ad6f99a53b2504043c08b1c3e214d17cde053cbdf91daa999ed5b47c37983ba3ee254bc5c793837daaa8c85cfc12f7f54f699f")
	expected, _ := hex.DecodeString("5b97a9d423f4b97413f388d9a341e727bb339f8e18a3fac2f2fb85abdc8f135deb30054a1afdc9b6ed7da16c55eba6b0d4d10c74e1d9a7cf8edfaeaa684ac0bd9f9d24ba674955c79dc6be32aee1c260b558ff07e3a4d49d24162011ff254db8be078e8ad07e648e6bf5679376cb4321a5ef01afe6ad8816fcc7634669c8c4389295c9241e45fff39f3225f7745032daeebe99d4b19bcb215d1bfdb36eda2c24")
	aescbcEncrypt(t, iv, key, plaintext, expected)
	aescbcDecrypt(t, iv, key, plaintext)

	aescbcEncryptParts(t, iv, key, plaintext, expected, 16)
	aescbcDecryptParts(t, iv, key, plaintext, 16)
}

func getCipherData(t *testing.T, iv, key, plaintext []byte) (io.Reader, Cipher) {
	c, err := NewAESCBC(key, iv)
	assert.Nil(t, err)
	cipherdata, err := c.Encrypt(bytes.NewBuffer(plaintext))
	assert.Nil(t, err)
	return cipherdata, c
}

func getCipherDataParts(t *testing.T, iv, key, plaintext []byte, partSize int) (io.Reader, Cipher) {
	c, err := NewAESCBC(key, iv)
	assert.Nil(t, err)
	cipherdata := bytes.NewBuffer([]byte{})
	for len(plaintext) > 0 {
		var data *bytes.Reader
		if len(plaintext) >= partSize {
			data, err = c.Encrypt(bytes.NewReader(plaintext[:partSize]))
		} else {
			data, err = c.Encrypt(bytes.NewReader(plaintext[:len(plaintext)]))
		}
		assert.Nil(t, err)

		data.WriteTo(cipherdata)
		if len(plaintext) < partSize {
			break
		}
		plaintext = plaintext[partSize:]
	}
	return cipherdata, c
}

func aescbcEncrypt(t *testing.T, iv, key, plaintext, expected []byte) {
	cipherdata, _ := getCipherData(t, iv, key, plaintext)

	ciphertext, err := ioutil.ReadAll(cipherdata)
	assert.Nil(t, err)
	assert.Equal(t, len(ciphertext), len(expected))
	assert.True(t, bytes.Equal(ciphertext, expected))
}

func aescbcEncryptParts(t *testing.T, iv, key, plaintext, expected []byte, partSize int) {
	cipherdata, _ := getCipherDataParts(t, iv, key, plaintext, partSize)

	ciphertext, err := ioutil.ReadAll(cipherdata)
	assert.Nil(t, err)
	assert.Equal(t, len(ciphertext), len(expected))
	assert.True(t, bytes.Equal(ciphertext, expected))
}

func aescbcDecrypt(t *testing.T, iv, key, plaintext []byte) {
	cipherdata, c := getCipherData(t, iv, key, plaintext)
	data, err := c.Decrypt(cipherdata)
	assert.Nil(t, err)

	text, err := ioutil.ReadAll(data)
	assert.Nil(t, err)
	assert.Equal(t, len(text), len(plaintext))
	assert.True(t, bytes.Equal(text, plaintext))
}

func aescbcDecryptParts(t *testing.T, iv, key, plaintext []byte, partSize int) {
	cipherdata, c := getCipherDataParts(t, iv, key, plaintext, partSize)
	text, err := ioutil.ReadAll(cipherdata)
	assert.Nil(t, err)

	cmpText := []byte{}
	for len(text) > 0 {
		var n int
		if len(text) >= partSize {
			n = partSize
		} else {
			n = len(text)
		}
		data, err := c.Decrypt(bytes.NewReader(text[:n]))
		assert.Nil(t, err)
		text = text[n:]
		b, err := ioutil.ReadAll(data)
		assert.Nil(t, err)
		cmpText = append(cmpText, b...)
	}

	assert.Equal(t, len(cmpText), len(plaintext))
	assert.True(t, bytes.Equal(cmpText, plaintext))
}
