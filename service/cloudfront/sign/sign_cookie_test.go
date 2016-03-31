package sign

import (
	"crypto/rsa"
	"testing"
	"time"
)

const (
	policyName    = "CloudFront-Policy"
	signatureName = "CloudFront-Signature"
	keyName       = "CloudFront-Key-Pair-Id"
)

func TestSignCookie(t *testing.T) {
	privKey, err := rsa.GenerateKey(randReader, 1024)
	if err != nil {
		t.Fatalf("Unexpected priv key error, %#v", err)
	}

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
	policy, signature, key, err := signer.SignCookies(p, nil)
	if err != nil {
		t.Fatalf("Error signing cookies %#v", err)
	}

	if policy.Name != policyName {
		t.Fatalf("policy name is incorrect, expected: %s\n got: %s", policyName, policy.Name)
	}
	if signature.Name != signatureName {
		t.Fatalf("signature name is incorrect, expected: %s\n got: %s", signatureName, signature.Name)
	}
	if key.Name != keyName {
		t.Fatalf("key name is incorrect, expected: %s\n got: %s", keyName, key.Name)
	}

	//test cookie signer with additional options
	o := &CookieOptions{
		Path:   "/",
		Domain: ".myexample.com",
		Secure: true,
	}

	policy, signature, key, err = signer.SignCookies(p, o)
	if err != nil {
		t.Fatalf("Error signing cookies %#v", err)
	}

	if policy.Name != policyName {
		t.Fatalf("policy name is incorrect, expected: %s\n got: %s", policyName, policy.Name)
	}
	if signature.Name != signatureName {
		t.Fatalf("signature name is incorrect, expected: %s\n got: %s", signatureName, signature.Name)
	}
	if key.Name != keyName {
		t.Fatalf("key name is incorrect, expected: %s\n got: %s", keyName, key.Name)
	}

	if policy.Domain != o.Domain || policy.Path != o.Path || policy.Secure != o.Secure {
		t.Fatalf("extra options were not applied correctly, expected: %+v\n got: %+v", o, policy)
	}
	if signature.Domain != o.Domain || signature.Path != o.Path || signature.Secure != o.Secure {
		t.Fatalf("extra options were not applied correctly, expected: %+v\n got: %+v", o, signature)
	}
	if key.Domain != o.Domain || key.Path != o.Path || key.Secure != o.Secure {
		t.Fatalf("extra options were not applied correctly, expected: %+v\n got: %+v", o, key)
	}

}
