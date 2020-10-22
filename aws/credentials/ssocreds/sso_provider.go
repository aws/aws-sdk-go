package ssocreds

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sso"
	"github.com/aws/aws-sdk-go/service/sso/ssoiface"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/internal/shareddefaults"
)

// ProviderName is the name of the credentials provider.
const (
	ErrCodeSSOCredentials = "SSOCredentialErr"
	ProviderName          = `SSOProvider`
)

// now is used to return a time.Time object representing
// the current time. This can be used to easily test and
// compare test values.
var now = time.Now

// awsSSOCachePath holds the path to the AWS SSO cache dir
var awsSSOCachePath string

// SSOProvider is used to retrieve credentials using an SSO access token
type SSOProvider struct {
	credentials.Expiry

	// Duration the STS credentials will be valid for. Truncated to seconds.
	// If unset, the assumed role will use AssumeRoleWithWebIdentity's default
	// expiry duration. See
	// https://docs.aws.amazon.com/sdk-for-go/api/service/sts/#STS.AssumeRoleWithWebIdentity
	// for more information.
	Duration time.Duration

	// The amount of time the credentials will be refreshed before they expire.
	// This is useful refresh credentials before they expire to reduce risk of
	// using credentials as they expire. If unset, will default to no expiry
	// window.
	ExpiryWindow time.Duration

	client ssoiface.SSOAPI

	accountID string
	roleName  string

	cache *SSOCache
}

// SSOCache represents an AWS SSO cache file
type SSOCache struct {
	StartURL    string `json:"startUrl"`
	Region      string `json:"region"`
	AccessToken string `json:"accessToken"`
	ExpiresAt   string `json:"expiresAt"`
}

// NewSSOCredentials will return a new set of temporary credentials based on the SSO role & token
func NewSSOCredentials(c client.ConfigProvider, ssoAccountID, ssoRoleName string) (*credentials.Credentials, error) {
	svc := sso.New(c)
	p, err := NewSSOProvider(svc, ssoAccountID, ssoRoleName)
	if err != nil {
		return nil, awserr.New(ErrCodeSSOCredentials, "failed to retrieve credentials", err)
	}
	return credentials.NewCredentials(p), nil
}

// NewSSOProvider will return a new SSOProvider configured with the
// details from the SSO cache
func NewSSOProvider(svc ssoiface.SSOAPI, accountID, roleName string) (*SSOProvider, error) {
	cache, err := getCache(filepath.Join(shareddefaults.UserHomeDir(), ".aws/sso/cache"))
	if err != nil {
		return nil, err
	}
	return &SSOProvider{
		client:    svc,
		accountID: accountID,
		roleName:  roleName,
		cache:     cache,
	}, nil
}

// Retrieve will attempt to get a set of temporary credentials
// using an AWS SSO token from the SSO Cache
func (p *SSOProvider) Retrieve() (credentials.Value, error) {
	return p.RetrieveWithContext(aws.BackgroundContext())
}

// RetrieveWithContext will attempt to get a set of temporary credentials
// using an AWS SSO token from the SSO Cache
func (p *SSOProvider) RetrieveWithContext(ctx credentials.Context) (credentials.Value, error) {
	in := &sso.GetRoleCredentialsInput{
		AccountId:   &p.accountID,
		RoleName:    &p.roleName,
		AccessToken: &p.cache.AccessToken,
	}
	req, resp := p.client.GetRoleCredentialsRequest(in)
	req.SetContext(ctx)

	if err := req.Send(); err != nil {
		return credentials.Value{}, awserr.New(ErrCodeSSOCredentials, "failed to retrieve credentials", err)
	}

	t := time.Unix(0, *resp.RoleCredentials.Expiration*int64(time.Millisecond))
	p.SetExpiration(t.UTC(), p.ExpiryWindow)

	return credentials.Value{
		ProviderName:    "SSOCredentialProvider",
		AccessKeyID:     *resp.RoleCredentials.AccessKeyId,
		SecretAccessKey: *resp.RoleCredentials.SecretAccessKey,
		SessionToken:    *resp.RoleCredentials.SessionToken,
	}, nil
}

func getCache(cacheDir string) (*SSOCache, error) {

	cache := &SSOCache{}

	err := filepath.Walk(cacheDir, func(path string, info os.FileInfo, err error) error {
		// handle failure accessing a path
		if err != nil {
			return err
		}
		// skip directories (excluding the cache dir itself)
		if info.IsDir() && path != cacheDir {
			return filepath.SkipDir
		}
		// skip anything that's not a json file
		if !strings.HasSuffix(path, ".json") {
			return nil
		}
		// skip the botocore files
		if strings.HasPrefix(filepath.Base(path), "botocore-") {
			return nil
		}
		// get the cache details from file
		cache, err = getCacheFile(path)
		if err != nil {
			return err
		}
		return io.EOF
	})

	if err != nil && err != io.EOF {
		return nil, err
	}

	return cache, nil

}

func getCacheFile(path string) (*SSOCache, error) {
	cache := &SSOCache{}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, cache)
	if err != nil {
		return nil, err
	}
	return cache, nil
}
