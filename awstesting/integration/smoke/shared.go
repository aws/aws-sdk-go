// +build integration

// Package smoke contains shared step definitions that are used across integration tests
package smoke

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gucumber/gucumber"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3crypto"
)

// Session is a shared session for all integration smoke tests to use.
var Session = session.Must(session.NewSession())

func init() {
	logLevel := Session.Config.LogLevel
	if os.Getenv("DEBUG") != "" {
		logLevel = aws.LogLevel(aws.LogDebug)
	}
	if os.Getenv("DEBUG_SIGNING") != "" {
		logLevel = aws.LogLevel(aws.LogDebugWithSigning)
	}
	if os.Getenv("DEBUG_BODY") != "" {
		logLevel = aws.LogLevel(aws.LogDebugWithHTTPBody)
	}
	Session.Config.LogLevel = logLevel

	gucumber.When(`^I call the "(.+?)" API$`, func(op string) {
		call(op, nil, false)
	})

	gucumber.When(`^I call the "(.+?)" API with:$`, func(op string, args [][]string) {
		call(op, args, false)
	})

	gucumber.Then(`^the value at "(.+?)" should be a list$`, func(member string) {
		vals, _ := awsutil.ValuesAtPath(gucumber.World["response"], member)
		assert.NotNil(gucumber.T, vals)
	})

	gucumber.Then(`^the response should contain a "(.+?)"$`, func(member string) {
		vals, _ := awsutil.ValuesAtPath(gucumber.World["response"], member)
		assert.NotEmpty(gucumber.T, vals)
	})

	gucumber.When(`^I attempt to call the "(.+?)" API with:$`, func(op string, args [][]string) {
		call(op, args, true)
	})

	gucumber.Then(`^I expect the response error code to be "(.+?)"$`, func(code string) {
		err, ok := gucumber.World["error"].(awserr.Error)
		assert.True(gucumber.T, ok, "no error returned")
		if ok {
			assert.Equal(gucumber.T, code, err.Code(), "Error: %v", err)
		}
	})

	gucumber.And(`^I expect the response error message to include:$`, func(data string) {
		err, ok := gucumber.World["error"].(awserr.Error)
		assert.True(gucumber.T, ok, "no error returned")
		if ok {
			assert.Contains(gucumber.T, err.Error(), data)
		}
	})

	gucumber.And(`^I expect the response error message to include one of:$`, func(table [][]string) {
		err, ok := gucumber.World["error"].(awserr.Error)
		assert.True(gucumber.T, ok, "no error returned")
		if ok {
			found := false
			for _, row := range table {
				if strings.Contains(err.Error(), row[0]) {
					found = true
					break
				}
			}

			assert.True(gucumber.T, found, fmt.Sprintf("no error messages matched: \"%s\"", err.Error()))
		}
	})

	gucumber.And(`^I expect the response error message not be empty$`, func() {
		err, ok := gucumber.World["error"].(awserr.Error)
		assert.True(gucumber.T, ok, "no error returned")
		assert.NotEmpty(gucumber.T, err.Message())
	})

	gucumber.When(`^I call the "(.+?)" API with JSON:$`, func(s1 string, data string) {
		callWithJSON(s1, data, false)
	})

	gucumber.When(`^I attempt to call the "(.+?)" API with JSON:$`, func(s1 string, data string) {
		callWithJSON(s1, data, true)
	})

	gucumber.Then(`^the error code should be "(.+?)"$`, func(s1 string) {
		err, ok := gucumber.World["error"].(awserr.Error)
		assert.True(gucumber.T, ok, "no error returned")
		assert.Equal(gucumber.T, s1, err.Code())
	})

	gucumber.And(`^the error message should contain:$`, func(data string) {
		err, ok := gucumber.World["error"].(awserr.Error)
		assert.True(gucumber.T, ok, "no error returned")
		assert.Contains(gucumber.T, err.Error(), data)
	})

	gucumber.Then(`^the request should fail$`, func() {
		err, ok := gucumber.World["error"].(awserr.Error)
		assert.True(gucumber.T, ok, "no error returned")
		assert.Error(gucumber.T, err)
	})

	gucumber.Then(`^the request should be successful$`, func() {
		err, ok := gucumber.World["error"].(awserr.Error)
		assert.False(gucumber.T, ok, "error returned")
		assert.NoError(gucumber.T, err)
	})

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
			fmt.Println("Grabbing cipher key:", cipherKey)

			// To get metadata for encryption key
			ctObj, err := s3Client.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    &cipherKey,
			})
			if err != nil {
				continue
			}

			masterkeyB64 := ctObj.Metadata["Masterkey"]
			masterkey, err := base64.StdEncoding.DecodeString(*masterkeyB64)
			assert.NoError(T, err)

			s3CryptoClient.Config.MasterKey = masterkey
			ctObj, err = s3CryptoClient.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    &cipherKey,
			},
			)
			assert.NoError(T, err)

			fmt.Println("CT OBJECT", ctObj)
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
			fmt.Println("CASE ", caseKey, len(plaintexts[caseKey]), len(ciphertext))
			assert.Equal(T, len(plaintexts[caseKey]), len(ciphertext))
			assert.True(T, bytes.Equal(plaintexts[caseKey], ciphertext))
		}
	})
}

// findMethod finds the op operation on the v structure using a case-insensitive
// lookup. Returns nil if no method is found.
func findMethod(v reflect.Value, op string) *reflect.Value {
	t := v.Type()
	op = strings.ToLower(op)
	for i := 0; i < t.NumMethod(); i++ {
		name := t.Method(i).Name
		if strings.ToLower(name) == op {
			m := v.MethodByName(name)
			return &m
		}
	}
	return nil
}

// call calls an operation on gucumber.World["client"] by the name op using the args
// table of arguments to set.
func call(op string, args [][]string, allowError bool) {
	v := reflect.ValueOf(gucumber.World["client"])
	if m := findMethod(v, op); m != nil {
		t := m.Type()
		in := reflect.New(t.In(0).Elem())
		fillArgs(in, args)

		resps := m.Call([]reflect.Value{in})
		gucumber.World["response"] = resps[0].Interface()
		gucumber.World["error"] = resps[1].Interface()

		if !allowError {
			err, _ := gucumber.World["error"].(error)
			assert.NoError(gucumber.T, err)
		}
	} else {
		assert.Fail(gucumber.T, "failed to find operation "+op)
	}
}

// reIsNum is a regular expression matching a numeric input (integer)
var reIsNum = regexp.MustCompile(`^\d+$`)

// reIsArray is a regular expression matching a list
var reIsArray = regexp.MustCompile(`^\['.*?'\]$`)
var reArrayElem = regexp.MustCompile(`'(.+?)'`)

// fillArgs fills arguments on the input structure using the args table of
// arguments.
func fillArgs(in reflect.Value, args [][]string) {
	if args == nil {
		return
	}

	for _, row := range args {
		path := row[0]
		var val interface{} = row[1]
		if reIsArray.MatchString(row[1]) {
			quotedStrs := reArrayElem.FindAllString(row[1], -1)
			strs := make([]*string, len(quotedStrs))
			for i, e := range quotedStrs {
				str := e[1 : len(e)-1]
				strs[i] = &str
			}
			val = strs
		} else if reIsNum.MatchString(row[1]) { // handle integer values
			num, err := strconv.ParseInt(row[1], 10, 64)
			if err == nil {
				val = num
			}
		}
		awsutil.SetValueAtPath(in.Interface(), path, val)
	}
}

func callWithJSON(op, j string, allowError bool) {
	v := reflect.ValueOf(gucumber.World["client"])
	if m := findMethod(v, op); m != nil {
		t := m.Type()
		in := reflect.New(t.In(0).Elem())
		fillJSON(in, j)

		resps := m.Call([]reflect.Value{in})
		gucumber.World["response"] = resps[0].Interface()
		gucumber.World["error"] = resps[1].Interface()

		if !allowError {
			err, _ := gucumber.World["error"].(error)
			assert.NoError(gucumber.T, err)
		}
	} else {
		assert.Fail(gucumber.T, "failed to find operation "+op)
	}
}

func fillJSON(in reflect.Value, j string) {
	d := json.NewDecoder(strings.NewReader(j))
	if err := d.Decode(in.Interface()); err != nil {
		panic(err)
	}
}
