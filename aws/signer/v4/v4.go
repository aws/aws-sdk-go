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

var ignoredHeaders = rules{
	blacklist{
		mapRule{
			"Authorization": struct{}{},
			"User-Agent":    struct{}{},
		},
	},
}

// requiredSignedHeaders is a whitelist for build canonical headers.
var requiredSignedHeaders = rules{
	whitelist{
		mapRule{
			"Cache-Control":                                               struct{}{},
			"Content-Disposition":                                         struct{}{},
			"Content-Encoding":                                            struct{}{},
			"Content-Language":                                            struct{}{},
			"Content-Md5":                                                 struct{}{},
			"Content-Type":                                                struct{}{},
			"Expires":                                                     struct{}{},
			"If-Match":                                                    struct{}{},
			"If-Modified-Since":                                           struct{}{},
			"If-None-Match":                                               struct{}{},
			"If-Unmodified-Since":                                         struct{}{},
			"Range":                                                       struct{}{},
			"X-Amz-Acl":                                                   struct{}{},
			"X-Amz-Copy-Source":                                           struct{}{},
			"X-Amz-Copy-Source-If-Match":                                  struct{}{},
			"X-Amz-Copy-Source-If-Modified-Since":                         struct{}{},
			"X-Amz-Copy-Source-If-None-Match":                             struct{}{},
			"X-Amz-Copy-Source-If-Unmodified-Since":                       struct{}{},
			"X-Amz-Copy-Source-Range":                                     struct{}{},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm": struct{}{},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key":       struct{}{},
			"X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5":   struct{}{},
			"X-Amz-Grant-Full-control":                                    struct{}{},
			"X-Amz-Grant-Read":                                            struct{}{},
			"X-Amz-Grant-Read-Acp":                                        struct{}{},
			"X-Amz-Grant-Write":                                           struct{}{},
			"X-Amz-Grant-Write-Acp":                                       struct{}{},
			"X-Amz-Metadata-Directive":                                    struct{}{},
			"X-Amz-Mfa":                                                   struct{}{},
			"X-Amz-Request-Payer":                                         struct{}{},
			"X-Amz-Server-Side-Encryption":                                struct{}{},
			"X-Amz-Server-Side-Encryption-Aws-Kms-Key-Id":                 struct{}{},
			"X-Amz-Server-Side-Encryption-Customer-Algorithm":             struct{}{},
			"X-Amz-Server-Side-Encryption-Customer-Key":                   struct{}{},
			"X-Amz-Server-Side-Encryption-Customer-Key-Md5":               struct{}{},
			"X-Amz-Storage-Class":                                         struct{}{},
			"X-Amz-Website-Redirect-Location":                             struct{}{},
		},
	},
	patterns{"X-Amz-Meta-"},
}

// allowedHoisting is a whitelist for build query headers. The boolean value
// represents whether or not it is a pattern.
var allowedQueryHoisting = inclusiveRules{
	blacklist{requiredSignedHeaders},
	patterns{"X-Amz-"},
}

// Signer applies AWS v4 signing to given request
type Signer struct {
	Credentials           *credentials.Credentials
	Debug                 aws.LogLevelType
	Logger                aws.Logger
	DisableHeaderHoisting bool
}

// NewSigner creates a new Signer
func NewSigner(credentials *credentials.Credentials, options ...func(*Signer)) *Signer {
	v4 := &Signer{
		Credentials: credentials,
	}

	for _, option := range options {
		option(v4)
	}

	return v4
}

type signingCtx struct {
	ServiceName      string
	Region           string
	Request          *http.Request
	Body             io.ReadSeeker
	Query            url.Values
	Time             time.Time
	ExpireTime       time.Duration
	SignedHeaderVals http.Header

	credValues         credentials.Value
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
}

// Sign signs AWS v4 requests
//
// The requests body needs to be passed in as an io.ReadSeeker to support hashing the body.
func (v4 Signer) Sign(r *http.Request, body io.ReadSeeker, service, region string, signTime time.Time) (http.Header, error) {
	return v4.signWithBody(r, body, service, region, 0, signTime)
}

// Presign presigns AWS v4 requests
//
// The requests body needs to be passed in as an io.ReadSeeker to support hashing the body.
func (v4 Signer) Presign(r *http.Request, body io.ReadSeeker, service, region string, exp time.Duration, signTime time.Time) (http.Header, error) {
	return v4.signWithBody(r, body, service, region, exp, signTime)
}

func (v4 Signer) signWithBody(r *http.Request, body io.ReadSeeker, service, region string, exp time.Duration, signTime time.Time) (http.Header, error) {
	ctx := &signingCtx{
		Request:     r,
		Body:        body,
		Query:       r.URL.Query(),
		Time:        signTime,
		ExpireTime:  exp,
		isPresign:   exp != 0,
		ServiceName: service,
		Region:      region,
	}

	if ctx.isRequestSigned() {
		if !v4.Credentials.IsExpired() && time.Now().Before(ctx.Time.Add(10*time.Minute)) {
			// If the request is already signed, and the credentials have not
			// expired, and the request is not too old ignore the signing request.
			return ctx.SignedHeaderVals, nil
		}
		ctx.Time = time.Now()
		ctx.handlePresignRemoval()
	}

	var err error
	ctx.credValues, err = v4.Credentials.Get()
	if err != nil {
		return http.Header{}, err
	}

	ctx.assignAmzQueryValues()
	ctx.build(v4.DisableHeaderHoisting)

	if v4.Debug.Matches(aws.LogDebugWithSigning) {
		v4.logSigningInfo(ctx)
	}

	return ctx.SignedHeaderVals, nil
}

func (ctx *signingCtx) handlePresignRemoval() {
	if !ctx.isPresign {
		return
	}

	// The credentials have expired for this request. The current signing
	// is invalid, and needs to be request because the request will fail.
	ctx.removePresign()

	// Update the request's query string to ensure the values stays in
	// sync in the case retrieving the new credentials fails.
	ctx.Request.URL.RawQuery = ctx.Query.Encode()
}

func (ctx *signingCtx) assignAmzQueryValues() {
	if ctx.isPresign {
		ctx.Query.Set("X-Amz-Algorithm", authHeaderPrefix)
		if ctx.credValues.SessionToken != "" {
			ctx.Query.Set("X-Amz-Security-Token", ctx.credValues.SessionToken)
		} else {
			ctx.Query.Del("X-Amz-Security-Token")
		}

		return
	}

	if ctx.credValues.SessionToken != "" {
		ctx.Request.Header.Set("X-Amz-Security-Token", ctx.credValues.SessionToken)
	}
}

// Sign AWS SDK service requests with version 4.
//
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

	v4 := NewSigner(req.Config.Credentials, func(v4 *Signer) {
		v4.Debug = req.Config.LogLevel.Value()
		v4.Logger = req.Config.Logger
		v4.DisableHeaderHoisting = req.NotHoist
	})

	signedHeaders, err := v4.signWithBody(req.HTTPRequest, req.Body, name, region, req.ExpireTime, req.Time)
	req.SignedHeaderVals = signedHeaders
	req.Error = err
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

func (v4 *Signer) logSigningInfo(ctx *signingCtx) {
	signedURLMsg := ""
	if ctx.isPresign {
		signedURLMsg = fmt.Sprintf(logSignedURLMsg, ctx.Request.URL.String())
	}
	msg := fmt.Sprintf(logSignInfoMsg, ctx.canonicalString, ctx.stringToSign, signedURLMsg)
	v4.Logger.Log(msg)
}

func (ctx *signingCtx) build(disableHeaderHoisting bool) {
	ctx.buildTime()             // no depends
	ctx.buildCredentialString() // no depends

	unsignedHeaders := ctx.Request.Header
	if ctx.isPresign {
		if !disableHeaderHoisting {
			urlValues := url.Values{}
			urlValues, unsignedHeaders = buildQuery(allowedQueryHoisting, unsignedHeaders) // no depends
			for k := range urlValues {
				ctx.Query[k] = urlValues[k]
			}
		}
	}

	ctx.buildCanonicalHeaders(ignoredHeaders, unsignedHeaders)
	ctx.buildCanonicalString() // depends on canon headers / signed headers
	ctx.buildStringToSign()    // depends on canon string
	ctx.buildSignature()       // depends on string to sign

	if ctx.isPresign {
		ctx.Request.URL.RawQuery += "&X-Amz-Signature=" + ctx.signature
	} else {
		parts := []string{
			authHeaderPrefix + " Credential=" + ctx.credValues.AccessKeyID + "/" + ctx.credentialString,
			"SignedHeaders=" + ctx.signedHeaders,
			"Signature=" + ctx.signature,
		}
		ctx.Request.Header.Set("Authorization", strings.Join(parts, ", "))
	}
}

func (ctx *signingCtx) buildTime() {
	ctx.formattedTime = ctx.Time.UTC().Format(timeFormat)
	ctx.formattedShortTime = ctx.Time.UTC().Format(shortTimeFormat)

	if ctx.isPresign {
		duration := int64(ctx.ExpireTime / time.Second)
		ctx.Query.Set("X-Amz-Date", ctx.formattedTime)
		ctx.Query.Set("X-Amz-Expires", strconv.FormatInt(duration, 10))
	} else {
		ctx.Request.Header.Set("X-Amz-Date", ctx.formattedTime)
	}
}

func (ctx *signingCtx) buildCredentialString() {
	ctx.credentialString = strings.Join([]string{
		ctx.formattedShortTime,
		ctx.Region,
		ctx.ServiceName,
		"aws4_request",
	}, "/")

	if ctx.isPresign {
		ctx.Query.Set("X-Amz-Credential", ctx.credValues.AccessKeyID+"/"+ctx.credentialString)
	}
}

func buildQuery(r rule, header http.Header) (url.Values, http.Header) {
	query := url.Values{}
	unsignedHeaders := http.Header{}
	for k, h := range header {
		if r.IsValid(k) {
			query[k] = h
		} else {
			unsignedHeaders[k] = h
		}
	}

	return query, unsignedHeaders
}
func (ctx *signingCtx) buildCanonicalHeaders(r rule, header http.Header) {
	var headers []string
	headers = append(headers, "host")
	for k, v := range header {
		canonicalKey := http.CanonicalHeaderKey(k)
		if !r.IsValid(canonicalKey) {
			continue // ignored header
		}
		if ctx.SignedHeaderVals == nil {
			ctx.SignedHeaderVals = make(http.Header)
		}

		lowerCaseKey := strings.ToLower(k)
		if _, ok := ctx.SignedHeaderVals[lowerCaseKey]; ok {
			// include additional values
			ctx.SignedHeaderVals[lowerCaseKey] = append(ctx.SignedHeaderVals[lowerCaseKey], v...)
			continue
		}

		headers = append(headers, lowerCaseKey)
		ctx.SignedHeaderVals[lowerCaseKey] = v
	}
	sort.Strings(headers)

	ctx.signedHeaders = strings.Join(headers, ";")

	if ctx.isPresign {
		ctx.Query.Set("X-Amz-SignedHeaders", ctx.signedHeaders)
	}

	headerValues := make([]string, len(headers))
	for i, k := range headers {
		if k == "host" {
			headerValues[i] = "host:" + ctx.Request.URL.Host
		} else {
			headerValues[i] = k + ":" +
				strings.Join(ctx.SignedHeaderVals[k], ",")
		}
	}

	ctx.canonicalHeaders = strings.Join(stripExcessSpaces(headerValues), "\n")
}

func (ctx *signingCtx) buildCanonicalString() {
	ctx.Request.URL.RawQuery = strings.Replace(ctx.Query.Encode(), "+", "%20", -1)
	uri := ctx.Request.URL.Opaque
	if uri != "" {
		uri = "/" + strings.Join(strings.Split(uri, "/")[3:], "/")
	} else {
		uri = ctx.Request.URL.Path
	}
	if uri == "" {
		uri = "/"
	}

	if ctx.ServiceName != "s3" {
		uri = rest.EscapePath(uri, false)
	}

	ctx.canonicalString = strings.Join([]string{
		ctx.Request.Method,
		uri,
		ctx.Request.URL.RawQuery,
		ctx.canonicalHeaders + "\n",
		ctx.signedHeaders,
		ctx.bodyDigest(),
	}, "\n")
}

func (ctx *signingCtx) buildStringToSign() {
	ctx.stringToSign = strings.Join([]string{
		authHeaderPrefix,
		ctx.formattedTime,
		ctx.credentialString,
		hex.EncodeToString(makeSha256([]byte(ctx.canonicalString))),
	}, "\n")
}

func (ctx *signingCtx) buildSignature() {
	secret := ctx.credValues.SecretAccessKey
	date := makeHmac([]byte("AWS4"+secret), []byte(ctx.formattedShortTime))
	region := makeHmac(date, []byte(ctx.Region))
	service := makeHmac(region, []byte(ctx.ServiceName))
	credentials := makeHmac(service, []byte("aws4_request"))
	signature := makeHmac(credentials, []byte(ctx.stringToSign))
	ctx.signature = hex.EncodeToString(signature)
}

func (ctx *signingCtx) bodyDigest() string {
	hash := ctx.Request.Header.Get("X-Amz-Content-Sha256")
	if hash == "" {
		if ctx.isPresign && ctx.ServiceName == "s3" {
			hash = "UNSIGNED-PAYLOAD"
		} else if ctx.Body == nil {
			hash = hex.EncodeToString(makeSha256([]byte{}))
		} else {
			hash = hex.EncodeToString(makeSha256Reader(ctx.Body))
		}
		ctx.Request.Header.Add("X-Amz-Content-Sha256", hash)
	}
	return hash
}

// isRequestSigned returns if the request is currently signed or presigned
func (ctx *signingCtx) isRequestSigned() bool {
	if ctx.isPresign && ctx.Query.Get("X-Amz-Signature") != "" {
		return true
	}
	if ctx.Request.Header.Get("Authorization") != "" {
		return true
	}

	return false
}

// unsign removes signing flags for both signed and presigned requests.
func (ctx *signingCtx) removePresign() {
	ctx.Query.Del("X-Amz-Algorithm")
	ctx.Query.Del("X-Amz-Signature")
	ctx.Query.Del("X-Amz-Security-Token")
	ctx.Query.Del("X-Amz-Date")
	ctx.Query.Del("X-Amz-Expires")
	ctx.Query.Del("X-Amz-Credential")
	ctx.Query.Del("X-Amz-SignedHeaders")
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

func stripExcessSpaces(headerVals []string) []string {
	vals := make([]string, len(headerVals))
	for i, str := range headerVals {
		stripped := ""
		found := false
		str = strings.TrimSpace(str)
		for _, c := range str {
			if !found && c == ' ' {
				stripped += string(c)
				found = true
			} else if c != ' ' {
				stripped += string(c)
				found = false
			}
		}
		vals[i] = stripped
	}
	return vals
}
