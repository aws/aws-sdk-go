package sign

import (
	"crypto/rsa"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

const (
	policyName    = "CloudFront-Policy"
	signatureName = "CloudFront-Signature"
	keyName       = "CloudFront-Key-Pair-Id"
	testDomain    = ".myexample.com"
)

func RunCookieSetup(t *testing.T, f func(*rsa.PrivateKey)) {
	origRandReader := randReader
	randReader = newRandomReader(rand.New(rand.NewSource(1)))
	defer func() {
		randReader = origRandReader
	}()

	privKey, err := rsa.GenerateKey(randReader, 1024)
	if err != nil {
		t.Fatalf("Unexpected priv key error, %#v", err)
	}
	f(privKey)
}
func TestCookieWithPolicy(t *testing.T) {
	RunCookieSetup(t, func(privKey *rsa.PrivateKey) {

		signer := NewCookieSigner("keyID", privKey)

		if signer.keyID != "keyID" || signer.privKey != privKey {
			t.Fatalf("NewCookieSigner does not properly assign values %+v", signer)
		}

		p := &Policy{
			Statements: []Statement{
				{
					Resource: "*",
					Condition: Condition{
						DateLessThan: &AWSEpochTime{time.Now().Add(1 * time.Hour)},
					},
				},
			},
		}
		//test cookie signer without any additional options
		cookies, err := signer.SignWithPolicy(p)
		if err != nil {
			t.Fatalf("Error signing cookies %#v", err)
		}
		validateCookies(t, cookies, signer)
	})
}

//test cookie signer with additional options
func TestCookieWithOptions(t *testing.T) {
	RunCookieSetup(t, func(privKey *rsa.PrivateKey) {

		signer := NewCookieSigner("keyID", privKey)
		signer.Path = "/"
		signer.Domain = testDomain
		signer.Secure = false

		if signer.keyID != "keyID" || signer.privKey != privKey {
			t.Fatalf("NewCookieSigner does not properly assign values %+v", signer)
		}

		p := &Policy{
			Statements: []Statement{
				{
					Resource: "*",
					Condition: Condition{
						DateLessThan: &AWSEpochTime{time.Now().Add(1 * time.Hour)},
					},
				},
			},
		}

		cookies, err := signer.SignWithPolicy(p)
		if err != nil {
			t.Fatalf("Error signing cookies %#v", err)
		}
		validateCookies(t, cookies, signer)

	})
}

func TestCannedCookie(t *testing.T) {
	RunCookieSetup(t, func(privKey *rsa.PrivateKey) {
		signer := NewCookieSigner("keyID", privKey)
		cookies, err := signer.Sign("http://sub.cloudfront.com/path", time.Now().Add(1*time.Hour))
		if err != nil {
			t.Fatalf("Error signing cookies %#v", err)
		}
		validateCookies(t, cookies, signer)
	})
}

func validateCookies(t *testing.T, c []*http.Cookie, s *CookieSigner) {
	if c[0].Name != policyName {
		t.Fatalf("policy name is incorrect, expected: %s\n got: %s", policyName, c[0].Name)
	}
	if c[1].Name != signatureName {
		t.Fatalf("signature name is incorrect, expected: %s\n got: %s", signatureName, c[1].Name)
	}
	if c[2].Name != keyName {
		t.Fatalf("key name is incorrect, expected: %s\n got: %s", keyName, c[2].Name)
	}

	if c[0].Domain != s.Domain || c[0].Path != s.Path || c[0].Secure != s.Secure {
		t.Fatalf("extra options were not applied correctly, expected: %+v\n got: %+v", s, c[0])
	}
	if c[1].Domain != s.Domain || c[1].Path != s.Path || c[1].Secure != s.Secure {
		t.Fatalf("extra options were not applied correctly, expected: %+v\n got: %+v", s, c[1])
	}
	if c[2].Domain != s.Domain || c[2].Path != s.Path || c[2].Secure != s.Secure {
		t.Fatalf("extra options were not applied correctly, expected: %+v\n got: %+v", s, c[2])
	}
}
