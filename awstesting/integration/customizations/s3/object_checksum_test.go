// +build integration

package s3

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestContentMD5Validate(t *testing.T) {
	body := []byte("really cool body content")
	bodySum := md5.Sum(body)
	sumBase64 := base64.StdEncoding.EncodeToString(bodySum[:])

	emptyBody := []byte{}
	emptyBodySum := md5.Sum(emptyBody)
	emptyBodySum64 := base64.StdEncoding.EncodeToString(emptyBodySum[:])

	cases := []struct {
		Name      string
		Body      []byte
		SumBase64 string
	}{
		{
			Body:      body,
			SumBase64: sumBase64,
			Name:      "contentMD5validation.pop",
		},
		{
			Body:      emptyBody,
			SumBase64: emptyBodySum64,
			Name:      "contentMD5validation.empty",
		},
	}

	for i, c := range cases {
		keyName := aws.String(c.Name)
		req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
			Bucket: bucketName,
			Key:    keyName,
			Body:   bytes.NewReader(c.Body),
		})

		req.Build()
		if e, a := c.SumBase64, req.HTTPRequest.Header.Get("Content-Md5"); e != a {
			t.Errorf("%d, expect %v sum, got %v", i, e, a)
		}

		if err := req.Send(); err != nil {
			t.Fatalf("%d, expect no error, got %v", i, err)
		}

		getReq, getOut := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: bucketName,
			Key:    keyName,
		})

		getReq.Build()
		if e, a := "append-md5", getReq.HTTPRequest.Header.Get("X-Amz-Te"); e != a {
			t.Errorf("%d, expect %v encoding, got %v", i, e, a)
		}
		if err := getReq.Send(); err != nil {
			t.Fatalf("%d, expect no error, got %v", i, err)
		}
		defer getOut.Body.Close()

		var readBody bytes.Buffer
		_, err := io.Copy(&readBody, getOut.Body)
		if err != nil {
			t.Fatalf("%d, expect no error, got %v", i, err)
		}

		if e, a := len(c.Body), readBody.Len(); e != a {
			t.Errorf("%d, expect %v len, got %v", i, e, a)
		}
	}
}
