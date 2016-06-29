// Test vectors from http://csrc.nist.gov/publications/nistpubs/800-38a/sp800-38a.pdf
package s3crypto

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Page 24
func TestAESECB_Case_F_1_1(t *testing.T) {
	plaintextStr := "6bc1bee22e409f96e93d7e117393172a" // Block 1
	plaintextStr += "ae2d8a571e03ac9c9eb76fac45af8e51" // Block 2
	plaintextStr += "30c81c46a35ce411e5fbc1191a0a52ef" // Block 3
	plaintextStr += "f69f2445df4f9b17ad2b417be66c3710" // Block 4
	plaintext, _ := hex.DecodeString(plaintextStr)

	expectedStr := "3ad77bb40d7a3660a89ecaf32466ef97" // Block 1
	expectedStr += "f5d3d58503b9699de785895a96fdbaaf" // Block 2
	expectedStr += "43b1cd7f598ece23881b00e3ed030688" // Block 3
	expectedStr += "7b0c785e27e8ad3f8223207104725dd4" // Block 4
	expected, _ := hex.DecodeString(expectedStr)

	key, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	c, err := NewAESECB(key)
	assert.Nil(t, err)
	cipherdata := c.Encrypt(bytes.NewBuffer(plaintext))
	ciphertext, err := ioutil.ReadAll(cipherdata)
	assert.Nil(t, err)
	assert.Equal(t, len(ciphertext), len(expected))
	assert.True(t, bytes.Equal(ciphertext, expected))
}

func TestAESECB_Case_F_1_2(t *testing.T) {
	expectedStr := "6bc1bee22e409f96e93d7e117393172a" // Block 1
	expectedStr += "ae2d8a571e03ac9c9eb76fac45af8e51" // Block 2
	expectedStr += "30c81c46a35ce411e5fbc1191a0a52ef" // Block 3
	expectedStr += "f69f2445df4f9b17ad2b417be66c3710" // Block 4
	expected, _ := hex.DecodeString(expectedStr)

	ciphertextStr := "3ad77bb40d7a3660a89ecaf32466ef97" // Block 1
	ciphertextStr += "f5d3d58503b9699de785895a96fdbaaf" // Block 2
	ciphertextStr += "43b1cd7f598ece23881b00e3ed030688" // Block 3
	ciphertextStr += "7b0c785e27e8ad3f8223207104725dd4" // Block 4
	ciphertext, _ := hex.DecodeString(ciphertextStr)

	key, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	c, err := NewAESECB(key)
	assert.Nil(t, err)
	plaindata := c.Decrypt(bytes.NewBuffer(ciphertext))
	plaintext, err := ioutil.ReadAll(plaindata)
	assert.Nil(t, err)
	assert.Equal(t, len(plaintext), len(expected))
	assert.True(t, bytes.Equal(plaintext, expected))
}

func TestAESECB_Case_F_1_3(t *testing.T) {
	plaintextStr := "6bc1bee22e409f96e93d7e117393172a" // Block 1
	plaintextStr += "ae2d8a571e03ac9c9eb76fac45af8e51" // Block 2
	plaintextStr += "30c81c46a35ce411e5fbc1191a0a52ef" // Block 3
	plaintextStr += "f69f2445df4f9b17ad2b417be66c3710" // Block 4
	plaintext, _ := hex.DecodeString(plaintextStr)

	expectedStr := "bd334f1d6e45f25ff712a214571fa5cc" // Block 1
	expectedStr += "974104846d0ad3ad7734ecb3ecee4eef" // Block 2
	expectedStr += "ef7afd2270e2e60adce0ba2face6444e" // Block 3
	expectedStr += "9a4b41ba738d6c72fb16691603c18e0e" // Block 4
	expected, _ := hex.DecodeString(expectedStr)

	key, _ := hex.DecodeString("8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b")
	c, err := NewAESECB(key)
	assert.Nil(t, err)
	cipherdata := c.Encrypt(bytes.NewBuffer(plaintext))
	ciphertext, err := ioutil.ReadAll(cipherdata)
	assert.Nil(t, err)
	assert.Equal(t, len(ciphertext), len(expected))
	assert.True(t, bytes.Equal(ciphertext, expected))
}

func TestAESECB_Case_F_1_4(t *testing.T) {
	expectedStr := "6bc1bee22e409f96e93d7e117393172a" // Block 1
	expectedStr += "ae2d8a571e03ac9c9eb76fac45af8e51" // Block 2
	expectedStr += "30c81c46a35ce411e5fbc1191a0a52ef" // Block 3
	expectedStr += "f69f2445df4f9b17ad2b417be66c3710" // Block 4
	expected, _ := hex.DecodeString(expectedStr)

	ciphertextStr := "bd334f1d6e45f25ff712a214571fa5cc" // Block 1
	ciphertextStr += "974104846d0ad3ad7734ecb3ecee4eef" // Block 2
	ciphertextStr += "ef7afd2270e2e60adce0ba2face6444e" // Block 3
	ciphertextStr += "9a4b41ba738d6c72fb16691603c18e0e" // Block 4
	ciphertext, _ := hex.DecodeString(ciphertextStr)

	key, _ := hex.DecodeString("8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b")
	c, err := NewAESECB(key)
	assert.Nil(t, err)
	plaindata := c.Decrypt(bytes.NewBuffer(ciphertext))
	plaintext, err := ioutil.ReadAll(plaindata)
	assert.Nil(t, err)
	assert.Equal(t, len(plaintext), len(expected))
	assert.True(t, bytes.Equal(plaintext, expected))
}

func TestAESECB_Case_F_1_5(t *testing.T) {
	plaintextStr := "6bc1bee22e409f96e93d7e117393172a" // Block 1
	plaintextStr += "ae2d8a571e03ac9c9eb76fac45af8e51" // Block 2
	plaintextStr += "30c81c46a35ce411e5fbc1191a0a52ef" // Block 3
	plaintextStr += "f69f2445df4f9b17ad2b417be66c3710" // Block 4
	plaintext, _ := hex.DecodeString(plaintextStr)

	expectedStr := "f3eed1bdb5d2a03c064b5a7e3db181f8" // Block 1
	expectedStr += "591ccb10d410ed26dc5ba74a31362870" // Block 2
	expectedStr += "b6ed21b99ca6f4f9f153e7b1beafed1d" // Block 3
	expectedStr += "23304b7a39f9f3ff067d8d8f9e24ecc7" // Block 4
	expected, _ := hex.DecodeString(expectedStr)

	key, _ := hex.DecodeString("603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4")
	c, err := NewAESECB(key)
	assert.Nil(t, err)
	cipherdata := c.Encrypt(bytes.NewBuffer(plaintext))
	ciphertext, err := ioutil.ReadAll(cipherdata)
	assert.Nil(t, err)
	assert.Equal(t, len(ciphertext), len(expected))
	assert.True(t, bytes.Equal(ciphertext, expected))
}

func TestAESECB_Case_F_1_6(t *testing.T) {
	expectedStr := "6bc1bee22e409f96e93d7e117393172a" // Block 1
	expectedStr += "ae2d8a571e03ac9c9eb76fac45af8e51" // Block 2
	expectedStr += "30c81c46a35ce411e5fbc1191a0a52ef" // Block 3
	expectedStr += "f69f2445df4f9b17ad2b417be66c3710" // Block 4
	expected, _ := hex.DecodeString(expectedStr)

	ciphertextStr := "f3eed1bdb5d2a03c064b5a7e3db181f8" // Block 1
	ciphertextStr += "591ccb10d410ed26dc5ba74a31362870" // Block 2
	ciphertextStr += "b6ed21b99ca6f4f9f153e7b1beafed1d" // Block 3
	ciphertextStr += "23304b7a39f9f3ff067d8d8f9e24ecc7" // Block 4
	ciphertext, _ := hex.DecodeString(ciphertextStr)

	key, _ := hex.DecodeString("603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4")
	c, err := NewAESECB(key)
	assert.Nil(t, err)
	plaindata := c.Decrypt(bytes.NewBuffer(ciphertext))
	plaintext, err := ioutil.ReadAll(plaindata)
	assert.Nil(t, err)
	assert.Equal(t, len(plaintext), len(expected))
	assert.True(t, bytes.Equal(plaintext, expected))
}
