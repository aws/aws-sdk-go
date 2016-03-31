package sign

import (
	"crypto/rsa"
	"net/http"
)

//the requirements for Cookie Signer is the same for the URL signer
type CookieSigner struct {
	keyID   string
	privKey *rsa.PrivateKey
}

type CookieOptions struct {
	Path   string
	Domain string
	Secure bool
}

//Example:

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
// 					// Optional date URL is not valid until
// 					DateGreaterThan: &sign.AWSEpochTime{time.Now().Add(30 * time.Minute)},
// 					// Required date the URL will expire after
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

// 	//creates 3 signed cookies, provide an optional Cookie Options struct to specify other options
// 	o := &sign.CookieOptions{
// 		Path:   "/",
// 		Domain: ".cNameAssociatedWithMyDistribution.com",
// 		Secure: true, //make sure your app/site can handle https payloads, otherwise set this to false
// 	}
// 	//http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-setting-signed-cookie-custom-policy.html#private-content-custom-policy-statement-signed-cookies-examples
//  //http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html
// 	//avoid adding an Expire or MaxAge. See provided AWS Documentation for more info

// 	policy, signature, key, err := signer.SignCookies(p)
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}

// 	http.SetCookie(w, policy)
// 	http.SetCookie(w, signature)
// 	http.SetCookie(w, key)

// }

// NewCookieSigner constructs and returns a new CookieSigner to be used to for signing
// Amazon CloudFront URL resources with.
func NewCookieSigner(keyID string, privKey *rsa.PrivateKey) *CookieSigner {
	return &CookieSigner{
		keyID:   keyID,
		privKey: privKey,
	}
}

//CookieOptions are optional
func (c CookieSigner) SignCookies(p *Policy, o *CookieOptions) (cPolicy, cSignature, cKey *http.Cookie, err error) {
	b64Sig, b64Policy, err := p.Sign(c.privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	cPolicy, cSignature, cKey = c.CreateCookies(b64Policy, b64Sig, o)
	return cPolicy, cSignature, cKey, nil

}

//prepares the cookies to be attached to the header. An (optional) options struct is provided
//in case people don't want to manually edit their cookies

func (c CookieSigner) CreateCookies(policy, signature []byte, o *CookieOptions) (cPolicy, cSignature, cKey *http.Cookie) {
	//creates proper cookies
	cPolicy = &http.Cookie{
		Name:     "CloudFront-Policy",
		Value:    string(policy),
		HttpOnly: true,
	}

	cSignature = &http.Cookie{
		Name:     "CloudFront-Signature",
		Value:    string(signature),
		HttpOnly: true,
	}
	cKey = &http.Cookie{
		Name:     "CloudFront-Key-Pair-Id",
		Value:    c.keyID,
		HttpOnly: true,
	}

	//if options are included, assign them to the cookies
	if o != nil {
		if o.Path != "" {
			cPolicy.Path = o.Path
			cSignature.Path = o.Path
			cKey.Path = o.Path
		}
		if o.Domain != "" {
			cPolicy.Domain = o.Domain
			cSignature.Domain = o.Domain
			cKey.Domain = o.Domain
		}
		if o.Secure != false {
			cPolicy.Secure = o.Secure
			cSignature.Secure = o.Secure
			cKey.Secure = o.Secure
		}
	}

	//return 3 cookies to be attached to the response
	return cPolicy, cSignature, cKey
}
