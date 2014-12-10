package aws

import (
	"net/http"
	"net/http/httputil"
	"strings"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	currentTime = func() time.Time {
		return time.Date(2014, 12, 10, 13, 40, 0, 0, time.FixedZone("PST", -8))
	}
	defer func() {
		currentTime = time.Now
	}()

	req, err := http.NewRequest(
		"POST",
		"https://wobble.awsamazon.com/fliff?twimble=boo&flaff=hojo&plum",
		strings.NewReader("this is a wonderful place to be"),
	)
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("X-Amz-Target", "Weeble.Wobble")
	if err != nil {
		t.Fatal(err)
	}

	c := Context{
		Service:     "wobble",
		Region:      "us-east-1",
		Credentials: Creds("accessKeyID", "secretAccessKey", ""),
	}
	c.sign(req)

	actual, err := httputil.DumpRequest(req, true)
	if err != nil {
		t.Fatal(err)
	}

	expected := "" +
		"POST /fliff?twimble=boo&flaff=hojo&plum HTTP/1.1\r\n" +
		"Host: wobble.awsamazon.com\r\n" +
		"Authorization: AWS4-HMAC-SHA256 Credential=accessKeyID/20141210/us-east-1/wobble/aws4_request, SignedHeaders=content-type;host;x-amz-content-sha256;x-amz-date;x-amz-target, Signature=9ee43779e8904f65972480295581e34413480b0695c537e058b5837b1bc08688\r\n" +
		"Content-Type: text/plain\r\n" +
		"X-Amz-Content-Sha256: 7e5e76a11da3174202305c75bfe0963cf2b1b63d7944851ac3c0b84e88105678\r\n" +
		"X-Amz-Date: 20141210T134008Z\r\n" +
		"X-Amz-Target: Weeble.Wobble\r\n" +
		"\r\n" +
		"this is a wonderful place to be"

	if string(actual) != expected {
		t.Errorf("Signed request was \n%s\n but expected\n%s", actual, expected)
	}
}
