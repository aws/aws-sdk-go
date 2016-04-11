// Package sign provides utilities to generate signed Cookies and URLs for Amazon CloudFront.
//
// More information about signed Cookies and their structure can be found at:
// http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-setting-signed-cookie-custom-policy.html
//
// To sign a Cookie, create a CookieSigner with your private key and credential pair key ID.
// Once you have a CookieSigner instance you can call Sign or SignWithPolicy to
// sign the URLs.
//
package sign

import (
	"crypto/rsa"
	"net/http"
	"time"
)

// CookieSigner provides Cookie signing utilities to sign Cookies for Amazon CloudFront
// resources.
// Additional Options are provided and will be required depending on your set up
type CookieSigner struct {
	keyID   string
	privKey *rsa.PrivateKey
	//optional parameters
	Path   string
	Domain string
	Secure bool
}

//Example:
//
//   // Sign Cookie to be valid for 1 hour from now using Canned Policy.
//   signer := sign.NewCookieSigner(keyID, privKey)
//   cookies, err := signer.Sign(rawURL, time.Now().Add(1*time.Hour))
//   if err != nil {
//       log.Fatalf("Failed to sign cookies, err: %s\n", err.Error())
//   }

//   http.SetCookie(w, cookie[0])
// 	 http.SetCookie(w, cookie[1])
// 	 http.SetCookie(w, cookie[2])
//

//  // Sign Cookie to be valid for 1 hour from now using Canned Policy.
// func handler(w http.ResponseWriter, r *http.Request) {

//  // sign cookie to be valid for 30 minutes from now, expires one hour from now, and
//  // restricted to the 192.0.2.0/24 IP address range.

//  //http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-setting-signed-cookie-custom-policy.html
// 	p := &sign.Policy{
// 		Statements: []sign.Statement{
// 			{
//				Resource: RawCloudFrontURL, //read the provided documentation on how to set this correctly, you'll probably want to use wildcards
// 				Condition: sign.Condition{
// 					// Optional IP source address range
// 					IPAddress: &sign.IPAddress{SourceIP: "192.0.2.0/24"},
// 					// Optional date Cookie is not valid until
// 					DateGreaterThan: &sign.AWSEpochTime{time.Now().Add(30 * time.Minute)},
// 					// Required date the Cookie will expire after
// 					DateLessThan: &sign.AWSEpochTime{time.Now().Add(1 * time.Hour)},
// 				},
// 			},
// 		},
// 	}

// 	//load your private key and convert it to type rsa.PrivateKey
// 	privKey, err := sign.LoadPEMPrivKeyFile("privatekey.pem")
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}

// 	//key ID that represents the key pair associated with the private key
// 	keyID := "WOIERULSDLKEWRIU"

// 	//set credentials to the cookiesigner
// 	signer := sign.NewCookieSigner(keyID, privKey)

// 	//provide an optional struct fields to specify other options
// 	signer.Path:   "/",
// 	signer.Domain: ".cNameAssociatedWithMyDistribution.com", //http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html
// 	signer.Secure: true, //make sure your app/site can handle https payloads, otherwise set this to false
//
// 	//avoid adding an Expire or MaxAge. See provided AWS Documentation for more info

// 	cookies, err := signer.SignWithPolicy(p)
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}

//	for _, cookie := range cookies {
//	  http.SetCookie(w, cookie)
//  }
// }

// NewCookieSigner constructs and returns a new CookieSigner to be used to for signing
// Amazon CloudFront Cookie resources with.
func NewCookieSigner(keyID string, privKey *rsa.PrivateKey) *CookieSigner {
	return &CookieSigner{
		keyID:   keyID,
		privKey: privKey,
	}
}

// Sign will sign cookies to expire at the time of expires sign using the
// Amazon CloudFront default Canned Policy. The Cookie will be signed with the
// private key and Credential Key Pair Key ID previously provided to CookieSigner.
//
// If extra policy conditions are need other than expiration use SignWithPolicy instead.
func (c CookieSigner) Sign(url string, expires time.Time) ([]*http.Cookie, error) {
	scheme, _, err := cleanURLScheme(url)
	if err != nil {
		return nil, err
	}

	resource, err := CreateResource(scheme, url)
	if err != nil {
		return nil, err
	}

	return c.SignWithPolicy(NewCannedPolicy(resource, expires))
}

// SignWithPolicy will sign cookies with the Policy provided.  The cookies will be
// signed with the private key and Credential Key Pair Key ID previously provided to CookieSigner.
//
// Use this signing method if you are looking to sign a Cookie with more than just
// the Policy's expiry time, or reusing Policies between multiple Cookie signings.
// If only the expiry time is needed you can use Sign and provide just the
// Cookies expiry time. A minimum of at least one policy statement is required for signed Cookies.
//
// Note: It is not safe to use Polices between multiple signers concurrently
func (c CookieSigner) SignWithPolicy(p *Policy) ([]*http.Cookie, error) {
	b64Sig, b64Policy, err := p.Sign(c.privKey)
	if err != nil {
		return nil, err
	}

	return c.createCookies(b64Policy, b64Sig), nil
}

//prepares the cookies to be attached to the header.
func (c CookieSigner) createCookies(policy, signature []byte) []*http.Cookie {
	//creates proper cookies
	p := &http.Cookie{
		Name:     "CloudFront-Policy",
		Value:    string(policy),
		HttpOnly: true,
		Secure:   false,
	}

	s := &http.Cookie{
		Name:     "CloudFront-Signature",
		Value:    string(signature),
		HttpOnly: true,
		Secure:   false,
	}
	k := &http.Cookie{
		Name:     "CloudFront-Key-Pair-Id",
		Value:    c.keyID,
		HttpOnly: true,
		Secure:   false,
	}

	//if options are included, assign them to the cookies
	if c.Path != "" {
		p.Path = c.Path
		s.Path = c.Path
		k.Path = c.Path
	}
	if c.Domain != "" {
		p.Domain = c.Domain
		s.Domain = c.Domain
		k.Domain = c.Domain
	}
	if c.Secure != false {
		p.Secure = c.Secure
		s.Secure = c.Secure
		k.Secure = c.Secure
	}

	//return 3 cookies to be attached to the response
	return []*http.Cookie{p, s, k}
}
