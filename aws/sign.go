package aws

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"
)

type V4Signer struct {
	Key                      string
	Secret                   string
	Service                  string
	Region                   string
	IncludeXAmzContentSha256 bool // Add the x-amz-content-sha256 header
}

func (s *V4Signer) sign(req *http.Request) {
	req.Header.Set("host", req.Host) // host header must be included as a signed header
	payloadHash := s.payloadHash(req)
	if s.IncludeXAmzContentSha256 {
		req.Header.Set("x-amz-content-sha256", payloadHash) // x-amz-content-sha256 contains the payload hash
	}
	t := s.requestTime(req)                           // Get request time
	creq := s.canonicalRequest(req, payloadHash)      // Build canonical request
	sts := s.stringToSign(t, creq)                    // Build string to sign
	signature := s.signature(t, sts)                  // Calculate the AWS Signature Version 4
	auth := s.authorization(req.Header, t, signature) // Create Authorization header value
	req.Header.Set("Authorization", auth)             // Add Authorization header to request
	return
}

const (
	iso8601BasicFormat      = "20060102T150405Z"
	iso8601BasicFormatShort = "20060102"
)

func (s *V4Signer) requestTime(req *http.Request) time.Time {

	// Get "x-amz-date" header
	date := req.Header.Get("x-amz-date")

	// Attempt to parse as ISO8601BasicFormat
	t, err := time.Parse(iso8601BasicFormat, date)
	if err == nil {
		return t
	}

	// Attempt to parse as http.TimeFormat
	t, err = time.Parse(http.TimeFormat, date)
	if err == nil {
		req.Header.Set("x-amz-date", t.Format(iso8601BasicFormat))
		return t
	}

	// Get "date" header
	date = req.Header.Get("date")

	// Attempt to parse as http.TimeFormat
	t, err = time.Parse(http.TimeFormat, date)
	if err == nil {
		return t
	}

	// Create a current time header to be used
	t = time.Now().UTC()
	req.Header.Set("x-amz-date", t.Format(iso8601BasicFormat))
	return t
}

func (s *V4Signer) canonicalRequest(req *http.Request, payloadHash string) string {
	if payloadHash == "" {
		payloadHash = s.payloadHash(req)
	}
	c := new(bytes.Buffer)
	fmt.Fprintf(c, "%s\n", req.Method)
	fmt.Fprintf(c, "%s\n", s.canonicalURI(req.URL))
	fmt.Fprintf(c, "%s\n", s.canonicalQueryString(req.URL))
	fmt.Fprintf(c, "%s\n\n", s.canonicalHeaders(req.Header))
	fmt.Fprintf(c, "%s\n", s.signedHeaders(req.Header))
	fmt.Fprintf(c, "%s", payloadHash)
	return c.String()
}

func (s *V4Signer) canonicalURI(u *url.URL) string {
	u = &url.URL{Path: u.Path}
	canonicalPath := u.String()
	slash := strings.HasSuffix(canonicalPath, "/")
	canonicalPath = path.Clean(canonicalPath)
	if canonicalPath != "/" && slash {
		canonicalPath += "/"
	}

	return canonicalPath
}

func (s *V4Signer) canonicalQueryString(u *url.URL) string {
	var a []string
	for k, vs := range u.Query() {
		k = url.QueryEscape(k)
		for _, v := range vs {
			if v == "" {
				a = append(a, k+"=")
			} else {
				v = url.QueryEscape(v)
				a = append(a, k+"="+v)
			}
		}
	}
	sort.Strings(a)
	return strings.Join(a, "&")
}

func (s *V4Signer) canonicalHeaders(h http.Header) string {
	i, a := 0, make([]string, len(h))
	for k, v := range h {
		for j, w := range v {
			v[j] = strings.Trim(w, " ")
		}
		sort.Strings(v)
		a[i] = strings.ToLower(k) + ":" + strings.Join(v, ",")
		i++
	}
	sort.Strings(a)
	return strings.Join(a, "\n")
}

func (s *V4Signer) signedHeaders(h http.Header) string {
	i, a := 0, make([]string, len(h))
	for k, _ := range h {
		a[i] = strings.ToLower(k)
		i++
	}
	sort.Strings(a)
	return strings.Join(a, ";")
}

func (s *V4Signer) payloadHash(req *http.Request) string {
	var b []byte
	if req.Body == nil {
		b = []byte("")
	} else {
		var err error
		b, err = ioutil.ReadAll(req.Body)
		if err != nil {
			// TODO: I REALLY DON'T LIKE THIS PANIC!!!!
			panic(err)
		}
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	return s.hash(string(b))
}

func (s *V4Signer) stringToSign(t time.Time, creq string) string {
	w := new(bytes.Buffer)
	fmt.Fprint(w, "AWS4-HMAC-SHA256\n")
	fmt.Fprintf(w, "%s\n", t.Format(iso8601BasicFormat))
	fmt.Fprintf(w, "%s\n", s.credentialScope(t))
	fmt.Fprintf(w, "%s", s.hash(creq))
	return w.String()
}

func (s *V4Signer) credentialScope(t time.Time) string {
	return fmt.Sprintf("%s/%s/%s/aws4_request", t.Format(iso8601BasicFormatShort), s.Region, s.Service)
}

func (s *V4Signer) signature(t time.Time, sts string) string {
	h := s.hmac(s.derivedKey(t), []byte(sts))
	return fmt.Sprintf("%x", h)
}

func (s *V4Signer) derivedKey(t time.Time) []byte {
	h := s.hmac([]byte("AWS4"+s.Secret), []byte(t.Format(iso8601BasicFormatShort)))
	h = s.hmac(h, []byte(s.Region))
	h = s.hmac(h, []byte(s.Service))
	h = s.hmac(h, []byte("aws4_request"))
	return h
}

func (s *V4Signer) authorization(header http.Header, t time.Time, signature string) string {
	w := new(bytes.Buffer)
	fmt.Fprint(w, "AWS4-HMAC-SHA256 ")
	fmt.Fprintf(w, "Credential=%s/%s, ", s.Key, s.credentialScope(t))
	fmt.Fprintf(w, "SignedHeaders=%s, ", s.signedHeaders(header))
	fmt.Fprintf(w, "Signature=%s", signature)
	return w.String()
}

func (s *V4Signer) hash(in string) string {
	h := sha256.New()
	fmt.Fprintf(h, "%s", in)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (s *V4Signer) hmac(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
