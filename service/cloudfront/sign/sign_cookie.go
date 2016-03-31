package sign

import (
	"crypto/rsa"
	"net/http"
)

//the requirements for Cookie Signer is the same for the URL signer
type CookieSigner struct {
	*URLSigner
}

type Options struct {
	Path   string
	Domain string
	Secure Bool
}

// NewCookieSigner constructs and returns a new CookieSigner to be used to for signing
// Amazon CloudFront URL resources with.
func NewCookieSigner(keyID string, privKey *rsa.PrivateKey) *CookieSigner {
	return &CookieSigner{
		keyID:   keyID,
		privKey: privKey,
	}
}

//prepares the cookies to be attached to the header. An options struct is provided
//in case people don't want to manually edit their cookies

func (c CookieSigner) CreateCookies(policy, signature []byte, o *Options) (cPolicy, cSignature, cKey *http.Cookie) {
	cPolicy := &http.Cookie{
		Name:   "CloudFront-Policy",
		Value:  string(policy),
		Path:   "/",
		Domain: ".launchpadcentral.com",
		// Secure:   true,
		HttpOnly: true,
	}

	cSig := &http.Cookie{
		Name:   "CloudFront-Signature",
		Value:  string(signature),
		Path:   "/",
		Domain: ".launchpadcentral.com",
		// Secure:   true,
		HttpOnly: true,
	}
	cKey := &http.Cookie{
		Name:   "CloudFront-Key-Pair-Id",
		Value:  key,
		Path:   "/",
		Domain: ".launchpadcentral.com",
		// Secure:   true,
		HttpOnly: true,
	}

	if Options != nil {
		if Options.Path != "" {
			cPolicy.Path = Options.Path
			cSignature.Path = Options.Path
			cKey.Path = Options.Path
		}
		if Options.Domain != "" {
			cPolicy.Domain = Options.Domain
			cSignature.Domain = Options.Domain
			cKey.Domain = Options.Domain
		}
		if Options.Secure != false {
			cPolicy.Secure = Options.Secure
			cSignature.Secure = Options.Secure
			cKey.Secure = Options.Secure
		}
	}
	return cPolicy, cSig, cKey
}
