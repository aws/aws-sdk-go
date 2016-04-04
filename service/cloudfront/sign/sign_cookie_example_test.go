package sign

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// /Example:
func TestSignedCookie(t *testing.T) {
	w := httptest.NewRecorder()

	HandlerExample(w)

	if len(w.HeaderMap["Set-Cookie"]) != 3 {
		t.Fatalf("Cookies were not properly set, expected %d   \n got: %d", 3, len(w.HeaderMap["Set-Cookie"]))
	}

}
func HandlerExample(w http.ResponseWriter) {

	// sign cookie to be valid for 30 minutes from now, expires one hour from now, and
	// restricted to the 192.0.2.0/24 IP address range.

	//http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-setting-signed-cookie-custom-policy.html
	p := &Policy{
		Statements: []Statement{
			{
				Resource: "http://sub.cloudfront.com", //read the provided documentation on how to set this correctly, you'll probably want to use wildcards
				Condition: Condition{
					// Optional IP source address range
					IPAddress: &IPAddress{SourceIP: "192.0.2.0/24"},
					// Optional date URL is not valid until
					DateGreaterThan: &AWSEpochTime{time.Now().Add(30 * time.Minute)},
					// Required date the URL will expire after
					DateLessThan: &AWSEpochTime{time.Now().Add(1 * time.Hour)},
				},
			},
		},
	}

	//load your private key and convert it to type rsa.PrivateKey
	// privKey, err := sign.LoadPEMPrivKeyFile("privatekey.pem")
	// if err != nil {
	// 	fmt.Println("error", err)
	// }

	//we create a random key for the example
	privKey, err := rsa.GenerateKey(randReader, 1024)

	//key ID that represents the key pair associated with the private key
	keyID := "WOIERULSDLKEWRIU"

	//set credentials to the cookiesigner
	signer := NewCookieSigner(keyID, privKey)

	//provide an optional struct fields to specify other options
	signer.Path = "/"
	signer.Domain = ".cNameAssociatedWithMyDistribution.com" //http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html
	signer.Secure = true                                     //make sure your app/site can handle https payloads, otherwise set this to false

	//avoid adding an Expire or MaxAge. See provided AWS Documentation for more info

	cookies, err := signer.SignWithPolicy(p)
	if err != nil {
		fmt.Println("error", err)
	}

	http.SetCookie(w, cookies[0])
	http.SetCookie(w, cookies[1])
	http.SetCookie(w, cookies[2])

}
