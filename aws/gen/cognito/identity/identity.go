// Package cogitoidentity provides a client for Amazon Cognito Identity.
package cogitoidentity

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// CogitoIdentity is a client for Amazon Cognito Identity.
type CogitoIdentity struct {
	client *aws.JSONClient
}

// New returns a new CogitoIdentity client.
func New(key, secret, region string, client *http.Client) *CogitoIdentity {
	if client == nil {
		client = http.DefaultClient
	}

	service := "cognito-identity"
	endpoint, service, region := endpoints.Lookup("cognito-identity", region)

	return &CogitoIdentity{
		client: &aws.JSONClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: service,
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:       client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "AWSCognitoIdentityService",
		},
	}
}

// CreateIdentityPool creates a new identity pool. The identity pool is a
// store of user identity information that is specific to your AWS account.
// The limit on identity pools is 60 per account.
func (c *CogitoIdentity) CreateIdentityPool(req CreateIdentityPoolInput) (resp *IdentityPool, err error) {
	resp = &IdentityPool{}
	err = c.client.Do("CreateIdentityPool", "POST", "/", req, resp)
	return
}

// DeleteIdentityPool deletes a user pool. Once a pool is deleted, users
// will not be able to authenticate with the pool.
func (c *CogitoIdentity) DeleteIdentityPool(req DeleteIdentityPoolInput) (err error) {
	// NRE
	err = c.client.Do("DeleteIdentityPool", "POST", "/", req, nil)
	return
}

// DescribeIdentityPool gets details about a particular identity pool,
// including the pool name, ID description, creation date, and current
// number of users.
func (c *CogitoIdentity) DescribeIdentityPool(req DescribeIdentityPoolInput) (resp *IdentityPool, err error) {
	resp = &IdentityPool{}
	err = c.client.Do("DescribeIdentityPool", "POST", "/", req, resp)
	return
}

// GetID generates (or retrieves) a Cognito ID. Supplying multiple logins
// will create an implicit linked account.
func (c *CogitoIdentity) GetID(req GetIDInput) (resp *GetIDResponse, err error) {
	resp = &GetIDResponse{}
	err = c.client.Do("GetId", "POST", "/", req, resp)
	return
}

// GetOpenIDToken gets an OpenID token, using a known Cognito ID. This
// known Cognito ID is returned by GetId . You can optionally add
// additional logins for the identity. Supplying multiple logins creates an
// implicit link. The OpenId token is valid for 15 minutes.
func (c *CogitoIdentity) GetOpenIDToken(req GetOpenIDTokenInput) (resp *GetOpenIDTokenResponse, err error) {
	resp = &GetOpenIDTokenResponse{}
	err = c.client.Do("GetOpenIdToken", "POST", "/", req, resp)
	return
}

// GetOpenIDTokenForDeveloperIdentity registers (or retrieves) a Cognito
// IdentityId and an OpenID Connect token for a user authenticated by your
// backend authentication process. Supplying multiple logins will create an
// implicit linked account. You can only specify one developer provider as
// part of the Logins map, which is linked to the identity pool. The
// developer provider is the "domain" by which Cognito will refer to your
// users. You can use GetOpenIdTokenForDeveloperIdentity to create a new
// identity and to link new logins (that is, user credentials issued by a
// public provider or developer provider) to an existing identity. When you
// want to create a new identity, the IdentityId should be null. When you
// want to associate a new login with an existing
// authenticated/unauthenticated identity, you can do so by providing the
// existing IdentityId . This API will create the identity in the specified
// IdentityPoolId
func (c *CogitoIdentity) GetOpenIDTokenForDeveloperIdentity(req GetOpenIDTokenForDeveloperIdentityInput) (resp *GetOpenIDTokenForDeveloperIdentityResponse, err error) {
	resp = &GetOpenIDTokenForDeveloperIdentityResponse{}
	err = c.client.Do("GetOpenIdTokenForDeveloperIdentity", "POST", "/", req, resp)
	return
}

// ListIdentities is undocumented.
func (c *CogitoIdentity) ListIdentities(req ListIdentitiesInput) (resp *ListIdentitiesResponse, err error) {
	resp = &ListIdentitiesResponse{}
	err = c.client.Do("ListIdentities", "POST", "/", req, resp)
	return
}

// ListIdentityPools lists all of the Cognito identity pools registered for
// your account.
func (c *CogitoIdentity) ListIdentityPools(req ListIdentityPoolsInput) (resp *ListIdentityPoolsResponse, err error) {
	resp = &ListIdentityPoolsResponse{}
	err = c.client.Do("ListIdentityPools", "POST", "/", req, resp)
	return
}

// LookupDeveloperIdentity retrieves the IdentityID associated with a
// DeveloperUserIdentifier or the list of DeveloperUserIdentifier s
// associated with an IdentityId for an existing identity. Either
// IdentityID or DeveloperUserIdentifier must not be null. If you supply
// only one of these values, the other value will be searched in the
// database and returned as a part of the response. If you supply both,
// DeveloperUserIdentifier will be matched against IdentityID . If the
// values are verified against the database, the response returns both
// values and is the same as the request. Otherwise a
// ResourceConflictException is thrown.
func (c *CogitoIdentity) LookupDeveloperIdentity(req LookupDeveloperIdentityInput) (resp *LookupDeveloperIdentityResponse, err error) {
	resp = &LookupDeveloperIdentityResponse{}
	err = c.client.Do("LookupDeveloperIdentity", "POST", "/", req, resp)
	return
}

// MergeDeveloperIdentities merges two users having different IdentityId s,
// existing in the same identity pool, and identified by the same developer
// provider. You can use this action to request that discrete users be
// merged and identified as a single user in the Cognito environment.
// Cognito associates the given source user SourceUserIdentifier ) with the
// IdentityId of the DestinationUserIdentifier . Only
// developer-authenticated users can be merged. If the users to be merged
// are associated with the same public provider, but as two different
// users, an exception will be thrown.
func (c *CogitoIdentity) MergeDeveloperIdentities(req MergeDeveloperIdentitiesInput) (resp *MergeDeveloperIdentitiesResponse, err error) {
	resp = &MergeDeveloperIdentitiesResponse{}
	err = c.client.Do("MergeDeveloperIdentities", "POST", "/", req, resp)
	return
}

// UnlinkDeveloperIdentity unlinks a DeveloperUserIdentifier from an
// existing identity. Unlinked developer users will be considered new
// identities next time they are seen. If, for a given Cognito identity,
// you remove all federated identities as well as the developer user
// identifier, the Cognito identity becomes inaccessible.
func (c *CogitoIdentity) UnlinkDeveloperIdentity(req UnlinkDeveloperIdentityInput) (err error) {
	// NRE
	err = c.client.Do("UnlinkDeveloperIdentity", "POST", "/", req, nil)
	return
}

// UnlinkIdentity unlinks a federated identity from an existing account.
// Unlinked logins will be considered new identities next time they are
// seen. Removing the last linked login will make this identity
// inaccessible.
func (c *CogitoIdentity) UnlinkIdentity(req UnlinkIdentityInput) (err error) {
	// NRE
	err = c.client.Do("UnlinkIdentity", "POST", "/", req, nil)
	return
}

// UpdateIdentityPool is undocumented.
func (c *CogitoIdentity) UpdateIdentityPool(req IdentityPool) (resp *IdentityPool, err error) {
	resp = &IdentityPool{}
	err = c.client.Do("UpdateIdentityPool", "POST", "/", req, resp)
	return
}

// CreateIdentityPoolInput is undocumented.
type CreateIdentityPoolInput struct {
	AllowUnauthenticatedIdentities bool              `json:"AllowUnauthenticatedIdentities"`
	DeveloperProviderName          string            `json:"DeveloperProviderName,omitempty"`
	IdentityPoolName               string            `json:"IdentityPoolName"`
	OpenIDConnectProviderARNs      []string          `json:"OpenIdConnectProviderARNs,omitempty"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}

// DeleteIdentityPoolInput is undocumented.
type DeleteIdentityPoolInput struct {
	IdentityPoolID string `json:"IdentityPoolId"`
}

// DescribeIdentityPoolInput is undocumented.
type DescribeIdentityPoolInput struct {
	IdentityPoolID string `json:"IdentityPoolId"`
}

// GetIDInput is undocumented.
type GetIDInput struct {
	AccountID      string            `json:"AccountId"`
	IdentityPoolID string            `json:"IdentityPoolId"`
	Logins         map[string]string `json:"Logins,omitempty"`
}

// GetIDResponse is undocumented.
type GetIDResponse struct {
	IdentityID string `json:"IdentityId,omitempty"`
}

// GetOpenIDTokenForDeveloperIdentityInput is undocumented.
type GetOpenIDTokenForDeveloperIdentityInput struct {
	IdentityID     string            `json:"IdentityId,omitempty"`
	IdentityPoolID string            `json:"IdentityPoolId"`
	Logins         map[string]string `json:"Logins"`
	TokenDuration  int64             `json:"TokenDuration,omitempty"`
}

// GetOpenIDTokenForDeveloperIdentityResponse is undocumented.
type GetOpenIDTokenForDeveloperIdentityResponse struct {
	IdentityID string `json:"IdentityId,omitempty"`
	Token      string `json:"Token,omitempty"`
}

// GetOpenIDTokenInput is undocumented.
type GetOpenIDTokenInput struct {
	IdentityID string            `json:"IdentityId"`
	Logins     map[string]string `json:"Logins,omitempty"`
}

// GetOpenIDTokenResponse is undocumented.
type GetOpenIDTokenResponse struct {
	IdentityID string `json:"IdentityId,omitempty"`
	Token      string `json:"Token,omitempty"`
}

// IdentityDescription is undocumented.
type IdentityDescription struct {
	IdentityID string   `json:"IdentityId,omitempty"`
	Logins     []string `json:"Logins,omitempty"`
}

// IdentityPool is undocumented.
type IdentityPool struct {
	AllowUnauthenticatedIdentities bool              `json:"AllowUnauthenticatedIdentities"`
	DeveloperProviderName          string            `json:"DeveloperProviderName,omitempty"`
	IdentityPoolID                 string            `json:"IdentityPoolId"`
	IdentityPoolName               string            `json:"IdentityPoolName"`
	OpenIDConnectProviderARNs      []string          `json:"OpenIdConnectProviderARNs,omitempty"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}

// IdentityPoolShortDescription is undocumented.
type IdentityPoolShortDescription struct {
	IdentityPoolID   string `json:"IdentityPoolId,omitempty"`
	IdentityPoolName string `json:"IdentityPoolName,omitempty"`
}

// ListIdentitiesInput is undocumented.
type ListIdentitiesInput struct {
	IdentityPoolID string `json:"IdentityPoolId"`
	MaxResults     int    `json:"MaxResults"`
	NextToken      string `json:"NextToken,omitempty"`
}

// ListIdentitiesResponse is undocumented.
type ListIdentitiesResponse struct {
	Identities     []IdentityDescription `json:"Identities,omitempty"`
	IdentityPoolID string                `json:"IdentityPoolId,omitempty"`
	NextToken      string                `json:"NextToken,omitempty"`
}

// ListIdentityPoolsInput is undocumented.
type ListIdentityPoolsInput struct {
	MaxResults int    `json:"MaxResults"`
	NextToken  string `json:"NextToken,omitempty"`
}

// ListIdentityPoolsResponse is undocumented.
type ListIdentityPoolsResponse struct {
	IdentityPools []IdentityPoolShortDescription `json:"IdentityPools,omitempty"`
	NextToken     string                         `json:"NextToken,omitempty"`
}

// LookupDeveloperIdentityInput is undocumented.
type LookupDeveloperIdentityInput struct {
	DeveloperUserIdentifier string `json:"DeveloperUserIdentifier,omitempty"`
	IdentityID              string `json:"IdentityId,omitempty"`
	IdentityPoolID          string `json:"IdentityPoolId"`
	MaxResults              int    `json:"MaxResults,omitempty"`
	NextToken               string `json:"NextToken,omitempty"`
}

// LookupDeveloperIdentityResponse is undocumented.
type LookupDeveloperIdentityResponse struct {
	DeveloperUserIdentifierList []string `json:"DeveloperUserIdentifierList,omitempty"`
	IdentityID                  string   `json:"IdentityId,omitempty"`
	NextToken                   string   `json:"NextToken,omitempty"`
}

// MergeDeveloperIdentitiesInput is undocumented.
type MergeDeveloperIdentitiesInput struct {
	DestinationUserIdentifier string `json:"DestinationUserIdentifier"`
	DeveloperProviderName     string `json:"DeveloperProviderName"`
	IdentityPoolID            string `json:"IdentityPoolId"`
	SourceUserIdentifier      string `json:"SourceUserIdentifier"`
}

// MergeDeveloperIdentitiesResponse is undocumented.
type MergeDeveloperIdentitiesResponse struct {
	IdentityID string `json:"IdentityId,omitempty"`
}

// UnlinkDeveloperIdentityInput is undocumented.
type UnlinkDeveloperIdentityInput struct {
	DeveloperProviderName   string `json:"DeveloperProviderName"`
	DeveloperUserIdentifier string `json:"DeveloperUserIdentifier"`
	IdentityID              string `json:"IdentityId"`
	IdentityPoolID          string `json:"IdentityPoolId"`
}

// UnlinkIdentityInput is undocumented.
type UnlinkIdentityInput struct {
	IdentityID     string            `json:"IdentityId"`
	Logins         map[string]string `json:"Logins"`
	LoginsToRemove []string          `json:"LoginsToRemove"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
