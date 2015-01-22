package v4

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/stripe/aws-go/aws"
)

const (
	authHeaderPrefix = "AWS4-HMAC-SHA256"
	timeFormat       = "20060102T150405Z"
	shortTimeFormat  = "20060102"
)

type signer struct {
	Request         *http.Request
	Time            time.Time
	ServiceName     string
	Region          string
	AcessKeyID      string
	SecretAccessKey string
	SessionToken    string
	Body            io.ReadSeeker
	Debug           uint

	formattedTime      string
	formattedShortTime string

	signedHeaders    string
	canonicalHeaders string
	canonicalString  string
	credentialString string
	stringToSign     string
	signature        string
	authorization    string
}

// Sign requests with signature version 4.
func Sign(req *aws.Request) {
	creds, _ := req.Service.Config.Credentials.Credentials()

	s := signer{
		Request:         req.HTTPRequest,
		Time:            req.Time,
		Body:            req.Body,
		ServiceName:     req.Service.ServiceName,
		Region:          req.Service.Config.Region,
		AcessKeyID:      creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		Debug:           req.Debug,
	}
	s.sign()
}

func (v4 *signer) sign() {
	formatted := v4.Time.UTC().Format(timeFormat)

	// remove the old headers
	v4.Request.Header.Del("Date")
	v4.Request.Header.Del("Authorization")

	v4.build()

	//v4.Debug = true
	if v4.Debug > 0 {
		fmt.Printf("---[ CANONICAL STRING  ]-----------------------------\n")
		fmt.Printf("%s\n", v4.canonicalString)
		fmt.Printf("-----------------------------------------------------\n\n")
		fmt.Printf("---[ STRING TO SIGN ]--------------------------------\n")
		fmt.Printf("%s\n", v4.stringToSign)
		fmt.Printf("-----------------------------------------------------\n")
	}

	// add the new ones
	v4.Request.Header.Add("Date", formatted)
	v4.Request.Header.Add("Authorization", v4.authorization)

	if v4.SessionToken != "" {
		v4.Request.Header.Add("X-Amz-Security-Token", v4.SessionToken)
	}
}

func (v4 *signer) build() {
	v4.buildTime()
	v4.buildCanonicalHeaders()
	v4.buildCredentialString()
	v4.buildCanonicalString()
	v4.buildStringToSign()
	v4.buildSignature()
	v4.buildAuthorization()
}

func (v4 *signer) buildTime() {
	v4.formattedTime = v4.Time.UTC().Format(timeFormat)
	v4.formattedShortTime = v4.Time.UTC().Format(shortTimeFormat)
}

func (v4 *signer) buildAuthorization() {
	v4.authorization = strings.Join([]string{
		authHeaderPrefix + " Credential=" + v4.AcessKeyID + "/" + v4.credentialString,
		"SignedHeaders=" + v4.signedHeaders,
		"Signature=" + v4.signature,
	}, ",")
}

func (v4 *signer) buildCredentialString() {
	v4.credentialString = strings.Join([]string{
		v4.formattedShortTime,
		v4.Region,
		v4.ServiceName,
		"aws4_request",
	}, "/")
}

func (v4 *signer) buildCanonicalHeaders() {
	v4.signedHeaders = "host"
	v4.canonicalHeaders = "host:" + v4.Request.URL.Host
}

func (v4 *signer) buildCanonicalString() {
	v4.canonicalString = strings.Join([]string{
		v4.Request.Method,
		v4.Request.URL.Path,
		v4.Request.URL.RawQuery,
		v4.canonicalHeaders + "\n",
		v4.signedHeaders,
		v4.bodyDigest(),
	}, "\n")
}

func (v4 *signer) buildStringToSign() {
	v4.stringToSign = strings.Join([]string{
		authHeaderPrefix,
		v4.formattedTime,
		v4.credentialString,
		hexDigest(makeSha256([]byte(v4.canonicalString))),
	}, "\n")
}

func (v4 *signer) buildSignature() {
	secret := v4.SecretAccessKey
	date := makeHmac([]byte("AWS4"+secret), []byte(v4.formattedShortTime))
	region := makeHmac(date, []byte(v4.Region))
	service := makeHmac(region, []byte(v4.ServiceName))
	credentials := makeHmac(service, []byte("aws4_request"))
	signature := makeHmac(credentials, []byte(v4.stringToSign))
	v4.signature = hexDigest(signature)
}

func (v4 *signer) bodyDigest() string {
	hash := v4.Request.Header.Get("X-Amz-Content-Sha256")
	if hash == "" {
		if v4.Body == nil {
			hash = hexDigest(makeSha256([]byte{}))
		} else {
			hash = hexDigest(makeSha256Reader(v4.Body))
		}
		v4.Request.Header.Add("X-Amz-Content-Sha256", hash)
	}
	return hash
}

func makeHmac(key []byte, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func makeSha256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func makeSha256Reader(reader io.ReadSeeker) []byte {
	packet := make([]byte, 4096)
	hash := sha256.New()

	defer reader.Seek(0, 0)
	for {
		n, err := reader.Read(packet)
		if err != nil || n == 0 {
			break
		}
		hash.Write(packet[0:n])
	}

	return hash.Sum(nil)
}

func hexDigest(data []byte) string {
	var buffer bytes.Buffer
	for i := range data {
		str := strconv.FormatUint(uint64(data[i]), 16)
		if len(str) < 2 {
			buffer.WriteString("0")
		}
		buffer.WriteString(str)
	}
	return buffer.String()
}
