package s3crypto_test

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

const (
	bucketName       = "aws-s3-shared-tests"
	plaintextPrefix  = "/plaintext_test_case_"
	ciphertextPrefix = "/ciphertext_test_case_"
	languagePrefix   = "/language_"
	versionOne       = "version_1"
	versionTwo       = "version_2"
)

var languages = []string{"Java"}

func TestIntegration(t *testing.T) {
	sess := session.New(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	s3CryptoClient := &s3crypto.Client{}

	s3Client := s3.New(sess)
	folders := []string{"/aes_cbc"}

	for _, folder := range folders {
		out, err := s3Client.ListObjects(&s3.ListObjectsInput{
			Bucket: aws.String(bucketName),
			Prefix: aws.String(versionOne + folder + plaintextPrefix),
		})
		assert.Nil(t, err)

		for _, obj := range out.Contents {
			plaintextKey := obj.Key
			ptObj, err := s3Client.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(bucketName),
				Key:    plaintextKey,
			})
			assert.Nil(t, err)

			caseKey := strings.TrimPrefix(*plaintextKey, folder+plaintextPrefix)
			for _, language := range languages {
				cipherKey := versionOne + folder + languagePrefix + language + ciphertextPrefix + caseKey

				// To get metadata for encryption key
				ctObj, err := s3Client.GetObject(&s3.GetObjectInput{
					Bucket: aws.String(bucketName),
					Key:    &cipherKey,
				})

				masterkeyB64 := ctObj.Metadata["Masterkey"]
				masterkey, err := base64.StdEncoding.DecodeString(*masterkeyB64)
				assert.Nil(t, err)
				t.Log("masterKey", *masterkey, masterkey)
				cipher, err := s3crypto.NewAESECB(masterkey)

				s3CryptoClient, err = s3crypto.NewClient(cipher, sess)
				ctObj, err = s3CryptoClient.GetObject(s3crypto.NewAESCBC, &s3crypto.GetObjectInput{
					S3GetObjectInput: &s3.GetObjectInput{
						Bucket: aws.String(bucketName),
						Key:    &cipherKey,
					},
				})
				assert.Nil(t, err)

				plaintext, err := ioutil.ReadAll(ptObj.Body)
				assert.Nil(t, err)
				ciphertext, err := ioutil.ReadAll(ctObj.Body)
				assert.Nil(t, err)
				assert.True(t, bytes.Equal(ciphertext, plaintext))
				t.Log(ciphertext, plaintext)
			}
		}
	}
}
