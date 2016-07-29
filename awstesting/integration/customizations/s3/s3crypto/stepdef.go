// Package s3crypto contains shared step definitions that are used across integration tests
package s3crypto

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"strings"

	. "github.com/gucumber/gucumber"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

func init() {
	When(`^I get all fixtures for "(.+?)" from "(.+?)"$`,
		func(cekAlg, bucket string) {
			prefix := "plaintext_test_case_"
			baseFolder := "crypto_tests/" + cekAlg
			s3Client := World["client"].(*s3.S3)

			out, err := s3Client.ListObjects(&s3.ListObjectsInput{
				Bucket: aws.String(bucket),
				Prefix: aws.String(baseFolder + "/" + prefix),
			})
			assert.NoError(T, err)

			plaintexts := make(map[string][]byte)
			for _, obj := range out.Contents {
				plaintextKey := obj.Key
				ptObj, err := s3Client.GetObject(&s3.GetObjectInput{
					Bucket: aws.String(bucket),
					Key:    plaintextKey,
				})
				assert.NoError(T, err)
				caseKey := strings.TrimPrefix(*plaintextKey, baseFolder+"/"+prefix)
				plaintext, err := ioutil.ReadAll(ptObj.Body)
				assert.NoError(T, err)

				plaintexts[caseKey] = plaintext
			}
			World["baseFolder"] = baseFolder
			World["bucket"] = bucket
			World["plaintexts"] = plaintexts
		})

	Then(`^I decrypt each fixture against "(.+?)" "(.+?)"$`, func(lang, version string) {
		plaintexts := World["plaintexts"].(map[string][]byte)
		baseFolder := World["baseFolder"].(string)
		bucket := World["bucket"].(string)
		prefix := "ciphertext_test_case_"
		s3Client := World["client"].(*s3.S3)
		s3CryptoClient := World["cryptoClient"].(*s3crypto.Client)
		language := "language_" + lang

		ciphertexts := make(map[string][]byte)
		for caseKey := range plaintexts {
			cipherKey := baseFolder + "/" + version + "/" + language + "/" + prefix + caseKey

			// To get metadata for encryption key
			ctObj, err := s3Client.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    &cipherKey,
			})
			if err != nil {
				continue
			}

			// We don't support wrap, so skip it
			if *ctObj.Metadata["X-Amz-Wrap-Alg"] == "AESWrap" {
				continue
			}
			//masterkeyB64 := ctObj.Metadata["Masterkey"]
			//masterkey, err := base64.StdEncoding.DecodeString(*masterkeyB64)
			//assert.NoError(T, err)

			//s3CryptoClient.Config.MasterKey = masterkey
			ctObj, err = s3CryptoClient.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    &cipherKey,
			},
			)
			assert.NoError(T, err)

			ciphertext, err := ioutil.ReadAll(ctObj.Body)
			assert.NoError(T, err)
			ciphertexts[caseKey] = ciphertext
		}
		World["ciphertexts"] = ciphertexts
	})

	And(`^I compare the decrypted ciphertext to the plaintext$`, func() {
		plaintexts := World["plaintexts"].(map[string][]byte)
		ciphertexts := World["ciphertexts"].(map[string][]byte)
		for caseKey, ciphertext := range ciphertexts {
			assert.Equal(T, len(plaintexts[caseKey]), len(ciphertext))
			assert.True(T, bytes.Equal(plaintexts[caseKey], ciphertext))
		}
	})

	Then(`^I encrypt each fixture with "(.+?)" "(.+?)" "(.+?)" and "(.+?)"$`, func(kek, v1, v2, cek string) {
		var handler s3crypto.CipherDataHandler
		var builder s3crypto.ContentCipherBuilder
		switch kek {
		case "kms":
			m := s3crypto.MaterialDescription{}
			arn, err := getAliasInformation(v1, v2)
			assert.Nil(T, err)

			b64Arn := base64.StdEncoding.EncodeToString([]byte(arn))
			assert.Nil(T, err)
			World["Masterkey"] = b64Arn

			handler, err = s3crypto.NewKMSKeyProvider(session.New(&aws.Config{
				Region: &v2,
			}), arn, m)
			assert.Nil(T, err)
		default:
			T.Skip()
		}

		switch cek {
		case "aes_gcm":
			builder = s3crypto.AESGCMContentCipherBuilder(handler)
		default:
			T.Skip()
		}

		c := s3crypto.New(nil, builder, func(c *s3crypto.Client) {
			c.Config.S3Session = session.New(&aws.Config{
				Region: aws.String("us-west-2"),
			})
			session.New(&aws.Config{
				Region: aws.String("us-east-1"),
			})
		})
		World["cryptoClient"] = c
		World["cek"] = cek
	})

	And(`^upload "(.+?)" data with folder "(.+?)"$`, func(language, folder string) {
		c := World["cryptoClient"].(*s3crypto.Client)
		cek := World["cek"].(string)
		bucket := World["bucket"].(string)
		plaintexts := World["plaintexts"].(map[string][]byte)
		key := World["Masterkey"].(string)
		for caseKey, plaintext := range plaintexts {
			input := &s3.PutObjectInput{
				Bucket: &bucket,
				Key:    aws.String("crypto_tests/" + cek + "/" + folder + "/language_" + language + "/ciphertext_test_case_" + caseKey),
				Body:   bytes.NewReader(plaintext),
				Metadata: map[string]*string{
					"Masterkey": &key,
				},
			}

			_, err := c.PutObject(input)
			assert.Nil(T, err)
		}
	})
}

func getAliasInformation(alias, region string) (string, error) {
	arn := ""
	svc := kms.New(session.New(&aws.Config{
		Region: &region,
	}))

	truncated := true
	var marker *string
	for truncated {
		out, err := svc.ListAliases(&kms.ListAliasesInput{
			Marker: marker,
		})
		if err != nil {
			return arn, err
		}
		for _, aliasEntry := range out.Aliases {
			if *aliasEntry.AliasName == "alias/"+alias {
				return *aliasEntry.AliasArn, nil
			}
		}
		truncated = *out.Truncated
		marker = out.NextMarker
	}

	return "", errors.New("The alias " + alias + " does not exist in your account. Please add the proper alias to a key")
}
