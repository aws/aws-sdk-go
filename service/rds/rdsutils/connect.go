package rdsutils

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

// BuildAuthToken is used for generate a autherized presigned URL to connect to
// the database.
//
// The endpoint consists of the scheme, hostname, and port. IE {scheme}://{hostname}[:port]. The
// region is the region of database that the auth token would be generated for. The dbUser is the user
// that the request would be authenticated with. The creds are the credentials the auth token is signed
// with.
func BuildAuthToken(endpoint, region, dbUser string, creds *credentials.Credentials) (string, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	values := req.URL.Query()
	values.Set("Action", "connect")
	values.Set("DBUser", dbUser)
	req.URL.RawQuery = values.Encode()

	signer := v4.Signer{
		Credentials: creds,
	}
	_, err = signer.Presign(req, nil, "rds", region, 15*time.Minute, time.Now())
	if err != nil {
		return "", err
	}

	return req.URL.String(), nil
}
