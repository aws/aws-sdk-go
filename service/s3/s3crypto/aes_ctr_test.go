package s3crypto

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AES CTR
func TestAES_CTR_RFC3686_Case_7(t *testing.T) {
	iv, _ := hex.DecodeString("00000060DB5672C97AA8F0B200000001")
	key, _ := hex.DecodeString("776BEFF2851DB06F4C8A0542C8696F6C6A81AF1EEC96B4D37FC1D689E6C1C104")
	plaintext := []byte("Single block msg")
	expected, _ := hex.DecodeString("145AD01DBF824EC7560863DC71E3E0C0")
	aesctrTest(t, iv, key, plaintext, expected)
}

func TestAES_CTR_RFC3686_Case_8(t *testing.T) {
	iv, _ := hex.DecodeString("00FAAC24C1585EF15A43D87500000001")
	key, _ := hex.DecodeString("F6D66D6BD52D59BB0796365879EFF886C66DD51A5B6A99744B50590C87A23884")
	plaintext, _ := hex.DecodeString("000102030405060708090A0B0C0D0E0F101112131415161718191A1B1C1D1E1F")
	expected, _ := hex.DecodeString("F05E231B3894612C49EE000B804EB2A9B8306B508F839D6A5530831D9344AF1C")
	aesctrTest(t, iv, key, plaintext, expected)
}

func TestAES_CTR_RFC3686_Case_9(t *testing.T) {
	iv, _ := hex.DecodeString("001CC5B751A51D70A1C1114800000001")
	key, _ := hex.DecodeString("FF7A617CE69148E4F1726E2F43581DE2AA62D9F805532EDFF1EED687FB54153D")
	plaintext, _ := hex.DecodeString("000102030405060708090A0B0C0D0E0F101112131415161718191A1B1C1D1E1F20212223")
	expected, _ := hex.DecodeString("EB6C52821D0BBBF7CE7594462ACA4FAAB407DF866569FD07F48CC0B583D6071F1EC0E6B8")
	aesctrTest(t, iv, key, plaintext, expected)
}

func aesctrTest(t *testing.T, iv, key, plaintext, expected []byte) {
	ctr, err := NewAESCTR(key, iv)
	assert.Nil(t, err)
	cipherdata, err := ctr.Encrypt(bytes.NewBuffer(plaintext))
	assert.Nil(t, err)

	ciphertext, err := ioutil.ReadAll(cipherdata)
	assert.Nil(t, err)
	assert.Equal(t, len(ciphertext), len(expected))
	assert.True(t, bytes.Equal(ciphertext, expected))

	data, err := ctr.Decrypt(bytes.NewBuffer(ciphertext))
	assert.Nil(t, err)
	text, err := ioutil.ReadAll(data)
	assert.Nil(t, err)
	assert.Equal(t, len(text), len(plaintext))
	assert.True(t, bytes.Equal(text, plaintext))
}
