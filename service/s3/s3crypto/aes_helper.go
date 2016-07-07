package s3crypto

import (
	"bytes"
)

func padAESKey(key []byte) []byte {
	padded := []byte{}

	if v := len(key); v < 16 {
		padded = make([]byte, 16-v)
	} else if v < 24 {
		padded = make([]byte, 24-v)
	} else if v < 32 {
		padded = make([]byte, 32-v)
	}

	padded = append(padded, key...)
	return padded
}

// PadPKCS5 uses the RFC2898 standard described here:
// https://www.ietf.org/rfc/rfc2898.txt
//
// Given data is blocksize-6
// 0 0 0 0 0 0 0 ... 0 0 0
// padding the bytes with the number of bytes missing, len(data)%blocksize
// 0 0 0 0 0 0 0 ... 0 0 0 6 6 6 6 6 6
func PadPKCS5(data []byte, blocksize int) []byte {
	padAmount := len(data) % blocksize

	// Here we do not want to pad something when not necessary
	if padAmount == 0 {
		return data
	}

	padding := blocksize - padAmount
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// UnpadPKCS5 unpad using PKCS5
func UnpadPKCS5(src []byte, blocksize int) []byte {
	length := len(src)
	if length < 1 {
		return src
	}
	count := src[length-1]

	// Verify correct padding. If it isnt, we assume that it hasn't been padded
	for i := 1; i <= int(count); i++ {
		if src[length-i] != count {
			return src
		}
	}
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
