package s3crypto

import "bytes"
import "io/ioutil"
import "fmt"

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
func PadPKCS7(data []byte, blocksize int) []byte {
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
func UnpadPKCS7(src []byte, blocksize int) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func decryptWithGCM(data []byte) []byte {
	gcm, err := NewAESGCM(make([]byte, 32), make([]byte, 12))
	fmt.Println(err)
	out, err := gcm.Decrypt(bytes.NewReader(data))
	fmt.Println(err)
	b, err := ioutil.ReadAll(out)
	fmt.Println(err)
	return b
}
