// Package v4 implements signing for AWS V4 signer
package v4

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/rest"
)

const (
	authHeaderPrefix = "AWS4-HMAC-SHA256"
	timeFormat       = "20060102T150405Z"
	shortTimeFormat  = "20060102"
)

var ignoredHeaders = map[string]bool{
	"Authorization":  true,
	"Content-Type":   true,
	"Content-Length": true,
	"User-Agent":     true,
}

type signer struct {
	Request     *http.Request
	Time        time.Time
	ExpireTime  time.Duration
	ServiceName string
	Region      string
	CredValues  credentials.Value
	Credentials *credentials.Credentials
	Query       url.Values
	Body        io.ReadSeeker
	Debug       aws.LogLevelType
	Logger      aws.Logger

	isPresign          bool
	formattedTime      string
	formattedShortTime string

	signedHeaders    string
	canonicalHeaders string
	canonicalString  string
	credentialString string
	stringToSign     string
	signature        string
	authorization    string
	notHoist         bool
	headers          map[string][]string
}

// Sign requests with signature version 4.
//
// Will sign the requests with the service config's Credentials object
// Signing is skipped if the credentials is the credentials.AnonymousCredentials
// object.
func Sign(req *request.Request) {
	// If the request does not need to be signed ignore the signing of the
	// request if the AnonymousCredentials object is used.
	if req.Config.Credentials == credentials.AnonymousCredentials {
		return
	}

	region := req.ClientInfo.SigningRegion
	if region == "" {
		region = aws.StringValue(req.Config.Region)
	}

	name := req.ClientInfo.SigningName
	if name == "" {
		name = req.ClientInfo.ServiceName
	}

	s := signer{
		Request:     req.HTTPRequest,
		Time:        req.Time,
		ExpireTime:  req.ExpireTime,
		Query:       req.HTTPRequest.URL.Query(),
		Body:        req.Body,
		ServiceName: name,
		Region:      region,
		Credentials: req.Config.Credentials,
		Debug:       req.Config.LogLevel.Value(),
		Logger:      req.Config.Logger,
		notHoist:    req.NotHoist,
	}

	req.Error = s.sign()
	req.Headers = s.headers
}

func (v4 *signer) sign() error {
	if v4.ExpireTime != 0 {
		v4.isPresign = true
	}

	if v4.isRequestSigned() {
		if !v4.Credentials.IsExpired() {
			// If the request is already signed, and the credentials have not
			// expired yet ignore the signing request.
			return nil
		}

		// The credentials have expired for this request. The current signing
		// is invalid, and needs to be request because the request will fail.
		if v4.isPresign {
			v4.removePresign()
			// Update the request's query string to ensure the values stays in
			// sync in the case retrieving the new credentials fails.
			v4.Request.URL.RawQuery = v4.Query.Encode()
		}
	}

	var err error
	v4.CredValues, err = v4.Credentials.Get()
	if err != nil {
		return err
	}

	if v4.isPresign {
		v4.Query.Set("X-Amz-Algorithm", authHeaderPrefix)
		if v4.CredValues.SessionToken != "" {
			v4.Query.Set("X-Amz-Security-Token", v4.CredValues.SessionToken)
		} else {
			v4.Query.Del("X-Amz-Security-Token")
		}
	} else if v4.CredValues.SessionToken != "" {
		v4.Request.Header.Set("X-Amz-Security-Token", v4.CredValues.SessionToken)
	}

	v4.build()

	if v4.Debug.Matches(aws.LogDebugWithSigning) {
		v4.logSigningInfo()
	}

	return nil
}

const logSignInfoMsg = `DEBUG: Request Signiture:
---[ CANONICAL STRING  ]-----------------------------
%s
---[ STRING TO SIGN ]--------------------------------
%s%s
-----------------------------------------------------`
const logSignedURLMsg = `
---[ SIGNED URL ]------------------------------------
%s`

func (v4 *signer) logSigningInfo() {
	signedURLMsg := ""
	if v4.isPresign {
		signedURLMsg = fmt.Sprintf(logSignedURLMsg, v4.Request.URL.String())
	}
	msg := fmt.Sprintf(logSignInfoMsg, v4.canonicalString, v4.stringToSign, signedURLMsg)
	v4.Logger.Log(msg)
}

func (v4 *signer) build() {
	filters := make(map[string][]string)
	urlValues := url.Values{}

	v4.buildTime()             // no depends
	v4.buildCredentialString() // no depends

	var allowedSignedHeaders = map[string][]string{
		"X-Amz-Acl":               nil,
		"X-Amz-Content-Sha256":    nil,
		"X-Amz-Date":              nil,
		"X-Amz-Meta-Other-Header": nil,
		"X-Amz-Security-Token":    nil,
		"X-Amz-Target":            nil,
		/*
			"Cache-Control":                         nil,
			"Content-Disposition":                   nil,
			"Content-Encoding":                      nil,
			"Content-Language":                      nil,
			"Content-Length":                        nil,
			"Content-Md5":                           nil,
			"Content-Type":                          nil,
			"Expires":                               nil,
			"If-Match":                              nil,
			"If-Modified-Since":                     nil,
			"If-None-Match":                         nil,
			"If-Unmodified-Since":                   nil,
			"Range":                                 nil,
			"X-Amz-Copy-Source":                     nil,
			"X-Amz-Copy-Source-If-Match":            nil,
			"X-Amz-Copy-Source-If-Modified-Since":   nil,
			"X-Amz-Copy-Source-If-None-Match":       nil,
			"X-Amz-Copy-Source-If-Unmodified-Since": nil,
			"X-Amz-Copy-Source-Range":               nil,
		*/
	}
	filters = v4.buildCanonicalHeaders(allowedSignedHeaders, filters)

	if v4.isPresign {
		var allowedHoisting = map[string][]string{
			"Content-Md5":         nil,
			"Content-Disposition": nil,
		}

		urlValues, filters = buildQuery(v4.notHoist, allowedHoisting, v4.Request.Header, filters) // no depends
		for k := range urlValues {
			v4.Request.Header.Del(k)
			v4.Query.Del(k)
			v4.Query[k] = append(v4.Query[k], urlValues[k]...)
		}
	}

	v4.buildCanonicalString() // depends on canon headers / signed headers
	v4.buildStringToSign()    // depends on canon string
	v4.buildSignature()       // depends on string to sign

	if v4.isPresign {
		v4.Request.URL.RawQuery += "&X-Amz-Signature=" + v4.signature
	} else {
		parts := []string{
			authHeaderPrefix + " Credential=" + v4.CredValues.AccessKeyID + "/" + v4.credentialString,
			"SignedHeaders=" + v4.signedHeaders,
			"Signature=" + v4.signature,
		}
		v4.Request.Header.Set("Authorization", strings.Join(parts, ", "))
	}
}

func (v4 *signer) buildTime() {
	v4.formattedTime = v4.Time.UTC().Format(timeFormat)
	v4.formattedShortTime = v4.Time.UTC().Format(shortTimeFormat)

	if v4.isPresign {
		duration := int64(v4.ExpireTime / time.Second)
		v4.Query.Set("X-Amz-Date", v4.formattedTime)
		v4.Query.Set("X-Amz-Expires", strconv.FormatInt(duration, 10))
	} else {
		v4.Request.Header.Set("X-Amz-Date", v4.formattedTime)
	}
}

func (v4 *signer) buildCredentialString() {
	v4.credentialString = strings.Join([]string{
		v4.formattedShortTime,
		v4.Region,
		v4.ServiceName,
		"aws4_request",
	}, "/")

	if v4.isPresign {
		v4.Query.Set("X-Amz-Credential", v4.CredValues.AccessKeyID+"/"+v4.credentialString)
	}
}

func buildQuery(notHoist bool, allowed, header, filters map[string][]string) (url.Values, map[string][]string) {
	query := url.Values{}
	for k, h := range header {
		_, allow := allowed[k]
		_, filter := filters[k]

		if allow && !(filter || notHoist) {
			filters[k] = h
			query[k] = append(query[k], h...)
		}
	}

	return query, filters
}
func (v4 *signer) buildCanonicalHeaders(allowed, filters map[string][]string) map[string][]string {
	var headers []string
	headers = append(headers, "host")
	for k, v := range v4.Request.Header {
		if _, ok := allowed[http.CanonicalHeaderKey(k)]; !ok {
			continue // ignored header
		}
		filters[k] = nil
		lowerCaseKey := strings.ToLower(k)
		headers = append(headers, lowerCaseKey)

		if v4.headers == nil {
			v4.headers = make(map[string][]string)
		}
		v4.headers[lowerCaseKey] = v
	}
	sort.Strings(headers)

	v4.signedHeaders = strings.Join(headers, ";")

	if v4.isPresign {
		v4.Query.Set("X-Amz-SignedHeaders", v4.signedHeaders)
	}

	headerValues := make([]string, len(headers))
	for i, k := range headers {
		if k == "host" {
			headerValues[i] = "host:" + v4.Request.URL.Host
		} else {
			headerValues[i] = k + ":" +
				strings.Join(v4.Request.Header[http.CanonicalHeaderKey(k)], ",")
		}
	}

	v4.canonicalHeaders = strings.Join(headerValues, "\n")
	return filters
}

func (v4 *signer) buildCanonicalString() {
	v4.Request.URL.RawQuery = strings.Replace(v4.Query.Encode(), "+", "%20", -1)
	uri := v4.Request.URL.Opaque
	if uri != "" {
		uri = "/" + strings.Join(strings.Split(uri, "/")[3:], "/")
	} else {
		uri = v4.Request.URL.Path
	}
	if uri == "" {
		uri = "/"
	}

	if v4.ServiceName != "s3" {
		uri = rest.EscapePath(uri, false)
	}

	v4.canonicalString = strings.Join([]string{
		v4.Request.Method,
		uri,
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
		hex.EncodeToString(makeSha256([]byte(v4.canonicalString))),
	}, "\n")
}

func (v4 *signer) buildSignature() {
	secret := v4.CredValues.SecretAccessKey
	date := makeHmac([]byte("AWS4"+secret), []byte(v4.formattedShortTime))
	region := makeHmac(date, []byte(v4.Region))
	service := makeHmac(region, []byte(v4.ServiceName))
	credentials := makeHmac(service, []byte("aws4_request"))
	signature := makeHmac(credentials, []byte(v4.stringToSign))
	v4.signature = hex.EncodeToString(signature)
}

func (v4 *signer) bodyDigest() string {
	hash := v4.Request.Header.Get("X-Amz-Content-Sha256")
	if hash == "" {
		if v4.isPresign && v4.ServiceName == "s3" {
			hash = "UNSIGNED-PAYLOAD"
		} else if v4.Body == nil {
			hash = hex.EncodeToString(makeSha256([]byte{}))
		} else {
			hash = hex.EncodeToString(makeSha256Reader(v4.Body))
		}
		v4.Request.Header.Add("X-Amz-Content-Sha256", hash)
	}
	return hash
}

// isRequestSigned returns if the request is currently signed or presigned
func (v4 *signer) isRequestSigned() bool {
	if v4.isPresign && v4.Query.Get("X-Amz-Signature") != "" {
		return true
	}
	if v4.Request.Header.Get("Authorization") != "" {
		return true
	}

	return false
}

// unsign removes signing flags for both signed and presigned requests.
func (v4 *signer) removePresign() {
	v4.Query.Del("X-Amz-Algorithm")
	v4.Query.Del("X-Amz-Signature")
	v4.Query.Del("X-Amz-Security-Token")
	v4.Query.Del("X-Amz-Date")
	v4.Query.Del("X-Amz-Expires")
	v4.Query.Del("X-Amz-Credential")
	v4.Query.Del("X-Amz-SignedHeaders")
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
	hash := sha256.New()
	start, _ := reader.Seek(0, 1)
	defer reader.Seek(start, 0)

	io.Copy(hash, reader)
	return hash.Sum(nil)
}
