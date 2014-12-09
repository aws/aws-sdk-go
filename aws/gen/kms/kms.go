// Package kms provides a client for AWS Key Management Service.
package kms

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// KMS is a client for AWS Key Management Service.
type KMS struct {
	client *aws.JSONClient
}

// New returns a new KMS client.
func New(key, secret, region string, client *http.Client) *KMS {
	if client == nil {
		client = http.DefaultClient
	}

	service := "kms"
	endpoint, service, region := endpoints.Lookup("kms", region)

	return &KMS{
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
			TargetPrefix: "TrentService",
		},
	}
}

// CreateAlias creates a display name for a customer master key. An alias
// can be used to identify a key and should be unique. The console enforces
// a one-to-one mapping between the alias and a key. An alias name can
// contain only alphanumeric characters, forward slashes underscores and
// dashes An alias must start with the word "alias" followed by a forward
// slash (alias/). An alias that begins with "aws" after the forward slash
// (alias/aws...) is reserved by Amazon Web Services
func (c *KMS) CreateAlias(req CreateAliasRequest) (err error) {
	// NRE
	err = c.client.Do("CreateAlias", "POST", "/", req, nil)
	return
}

// CreateGrant adds a grant to a key to specify who can access the key and
// under what conditions. Grants are alternate permission mechanisms to key
// policies. If absent, access to the key is evaluated based on IAM
// policies attached to the user. By default, grants do not expire. Grants
// can be listed, retired, or revoked as indicated by the following APIs.
// Typically, when you are finished using a grant, you retire it. When you
// want to end a grant immediately, revoke it. For more information about
// grants, see Grants .
func (c *KMS) CreateGrant(req CreateGrantRequest) (resp *CreateGrantResponse, err error) {
	resp = &CreateGrantResponse{}
	err = c.client.Do("CreateGrant", "POST", "/", req, resp)
	return
}

// CreateKey creates a customer master key. Customer master keys can be
// used to encrypt small amounts of data (less than 4K) directly, but they
// are most commonly used to encrypt or envelope data keys that are then
// used to encrypt customer data. For more information about data keys, see
// GenerateDataKey and GenerateDataKeyWithoutPlaintext
func (c *KMS) CreateKey(req CreateKeyRequest) (resp *CreateKeyResponse, err error) {
	resp = &CreateKeyResponse{}
	err = c.client.Do("CreateKey", "POST", "/", req, resp)
	return
}

// Decrypt decrypts ciphertext. Ciphertext is plaintext that has been
// previously encrypted by using the Encrypt function.
func (c *KMS) Decrypt(req DecryptRequest) (resp *DecryptResponse, err error) {
	resp = &DecryptResponse{}
	err = c.client.Do("Decrypt", "POST", "/", req, resp)
	return
}

// DeleteAlias is undocumented.
func (c *KMS) DeleteAlias(req DeleteAliasRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteAlias", "POST", "/", req, nil)
	return
}

// DescribeKey provides detailed information about the specified customer
// master key.
func (c *KMS) DescribeKey(req DescribeKeyRequest) (resp *DescribeKeyResponse, err error) {
	resp = &DescribeKeyResponse{}
	err = c.client.Do("DescribeKey", "POST", "/", req, resp)
	return
}

// DisableKey marks a key as disabled, thereby preventing its use.
func (c *KMS) DisableKey(req DisableKeyRequest) (err error) {
	// NRE
	err = c.client.Do("DisableKey", "POST", "/", req, nil)
	return
}

// DisableKeyRotation is undocumented.
func (c *KMS) DisableKeyRotation(req DisableKeyRotationRequest) (err error) {
	// NRE
	err = c.client.Do("DisableKeyRotation", "POST", "/", req, nil)
	return
}

// EnableKey marks a key as enabled, thereby permitting its use. You can
// have up to 25 enabled keys at one time.
func (c *KMS) EnableKey(req EnableKeyRequest) (err error) {
	// NRE
	err = c.client.Do("EnableKey", "POST", "/", req, nil)
	return
}

// EnableKeyRotation enables rotation of the specified customer master key.
func (c *KMS) EnableKeyRotation(req EnableKeyRotationRequest) (err error) {
	// NRE
	err = c.client.Do("EnableKeyRotation", "POST", "/", req, nil)
	return
}

// Encrypt encrypts plaintext into ciphertext by using a customer master
// key.
func (c *KMS) Encrypt(req EncryptRequest) (resp *EncryptResponse, err error) {
	resp = &EncryptResponse{}
	err = c.client.Do("Encrypt", "POST", "/", req, resp)
	return
}

// GenerateDataKey generates a secure data key. Data keys are used to
// encrypt and decrypt data. They are wrapped by customer master keys.
func (c *KMS) GenerateDataKey(req GenerateDataKeyRequest) (resp *GenerateDataKeyResponse, err error) {
	resp = &GenerateDataKeyResponse{}
	err = c.client.Do("GenerateDataKey", "POST", "/", req, resp)
	return
}

// GenerateDataKeyWithoutPlaintext returns a key wrapped by a customer
// master key without the plaintext copy of that key. To retrieve the
// plaintext, see GenerateDataKey .
func (c *KMS) GenerateDataKeyWithoutPlaintext(req GenerateDataKeyWithoutPlaintextRequest) (resp *GenerateDataKeyWithoutPlaintextResponse, err error) {
	resp = &GenerateDataKeyWithoutPlaintextResponse{}
	err = c.client.Do("GenerateDataKeyWithoutPlaintext", "POST", "/", req, resp)
	return
}

// GenerateRandom is undocumented.
func (c *KMS) GenerateRandom(req GenerateRandomRequest) (resp *GenerateRandomResponse, err error) {
	resp = &GenerateRandomResponse{}
	err = c.client.Do("GenerateRandom", "POST", "/", req, resp)
	return
}

// GetKeyPolicy is undocumented.
func (c *KMS) GetKeyPolicy(req GetKeyPolicyRequest) (resp *GetKeyPolicyResponse, err error) {
	resp = &GetKeyPolicyResponse{}
	err = c.client.Do("GetKeyPolicy", "POST", "/", req, resp)
	return
}

// GetKeyRotationStatus retrieves a Boolean value that indicates whether
// key rotation is enabled for the specified key.
func (c *KMS) GetKeyRotationStatus(req GetKeyRotationStatusRequest) (resp *GetKeyRotationStatusResponse, err error) {
	resp = &GetKeyRotationStatusResponse{}
	err = c.client.Do("GetKeyRotationStatus", "POST", "/", req, resp)
	return
}

// ListAliases is undocumented.
func (c *KMS) ListAliases(req ListAliasesRequest) (resp *ListAliasesResponse, err error) {
	resp = &ListAliasesResponse{}
	err = c.client.Do("ListAliases", "POST", "/", req, resp)
	return
}

// ListGrants is undocumented.
func (c *KMS) ListGrants(req ListGrantsRequest) (resp *ListGrantsResponse, err error) {
	resp = &ListGrantsResponse{}
	err = c.client.Do("ListGrants", "POST", "/", req, resp)
	return
}

// ListKeyPolicies is undocumented.
func (c *KMS) ListKeyPolicies(req ListKeyPoliciesRequest) (resp *ListKeyPoliciesResponse, err error) {
	resp = &ListKeyPoliciesResponse{}
	err = c.client.Do("ListKeyPolicies", "POST", "/", req, resp)
	return
}

// ListKeys is undocumented.
func (c *KMS) ListKeys(req ListKeysRequest) (resp *ListKeysResponse, err error) {
	resp = &ListKeysResponse{}
	err = c.client.Do("ListKeys", "POST", "/", req, resp)
	return
}

// PutKeyPolicy is undocumented.
func (c *KMS) PutKeyPolicy(req PutKeyPolicyRequest) (err error) {
	// NRE
	err = c.client.Do("PutKeyPolicy", "POST", "/", req, nil)
	return
}

// ReEncrypt encrypts data on the server side with a new customer master
// key without exposing the plaintext of the data on the client side. The
// data is first decrypted and then encrypted. This operation can also be
// used to change the encryption context of a ciphertext.
func (c *KMS) ReEncrypt(req ReEncryptRequest) (resp *ReEncryptResponse, err error) {
	resp = &ReEncryptResponse{}
	err = c.client.Do("ReEncrypt", "POST", "/", req, resp)
	return
}

// RetireGrant retires a grant. You can retire a grant when you're done
// using it to clean up. You should revoke a grant when you intend to
// actively deny operations that depend on it.
func (c *KMS) RetireGrant(req RetireGrantRequest) (err error) {
	// NRE
	err = c.client.Do("RetireGrant", "POST", "/", req, nil)
	return
}

// RevokeGrant revokes a grant. You can revoke a grant to actively deny
// operations that depend on it.
func (c *KMS) RevokeGrant(req RevokeGrantRequest) (err error) {
	// NRE
	err = c.client.Do("RevokeGrant", "POST", "/", req, nil)
	return
}

// UpdateKeyDescription <nil>
func (c *KMS) UpdateKeyDescription(req UpdateKeyDescriptionRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateKeyDescription", "POST", "/", req, nil)
	return
}

// AliasListEntry is undocumented.
type AliasListEntry struct {
	AliasARN    string `json:"AliasArn,omitempty"`
	AliasName   string `json:"AliasName,omitempty"`
	TargetKeyID string `json:"TargetKeyId,omitempty"`
}

// CreateAliasRequest is undocumented.
type CreateAliasRequest struct {
	AliasName   string `json:"AliasName"`
	TargetKeyID string `json:"TargetKeyId"`
}

// CreateGrantRequest is undocumented.
type CreateGrantRequest struct {
	Constraints       GrantConstraints `json:"Constraints,omitempty"`
	GrantTokens       []string         `json:"GrantTokens,omitempty"`
	GranteePrincipal  string           `json:"GranteePrincipal"`
	KeyID             string           `json:"KeyId"`
	Operations        []string         `json:"Operations,omitempty"`
	RetiringPrincipal string           `json:"RetiringPrincipal,omitempty"`
}

// CreateGrantResponse is undocumented.
type CreateGrantResponse struct {
	GrantID    string `json:"GrantId,omitempty"`
	GrantToken string `json:"GrantToken,omitempty"`
}

// CreateKeyRequest is undocumented.
type CreateKeyRequest struct {
	Description string `json:"Description,omitempty"`
	KeyUsage    string `json:"KeyUsage,omitempty"`
	Policy      string `json:"Policy,omitempty"`
}

// CreateKeyResponse is undocumented.
type CreateKeyResponse struct {
	KeyMetadata KeyMetadata `json:"KeyMetadata,omitempty"`
}

// DecryptRequest is undocumented.
type DecryptRequest struct {
	CiphertextBlob    []byte            `json:"CiphertextBlob"`
	EncryptionContext map[string]string `json:"EncryptionContext,omitempty"`
	GrantTokens       []string          `json:"GrantTokens,omitempty"`
}

// DecryptResponse is undocumented.
type DecryptResponse struct {
	KeyID     string `json:"KeyId,omitempty"`
	Plaintext []byte `json:"Plaintext,omitempty"`
}

// DeleteAliasRequest is undocumented.
type DeleteAliasRequest struct {
	AliasName string `json:"AliasName"`
}

// DescribeKeyRequest is undocumented.
type DescribeKeyRequest struct {
	KeyID string `json:"KeyId"`
}

// DescribeKeyResponse is undocumented.
type DescribeKeyResponse struct {
	KeyMetadata KeyMetadata `json:"KeyMetadata,omitempty"`
}

// DisableKeyRequest is undocumented.
type DisableKeyRequest struct {
	KeyID string `json:"KeyId"`
}

// DisableKeyRotationRequest is undocumented.
type DisableKeyRotationRequest struct {
	KeyID string `json:"KeyId"`
}

// EnableKeyRequest is undocumented.
type EnableKeyRequest struct {
	KeyID string `json:"KeyId"`
}

// EnableKeyRotationRequest is undocumented.
type EnableKeyRotationRequest struct {
	KeyID string `json:"KeyId"`
}

// EncryptRequest is undocumented.
type EncryptRequest struct {
	EncryptionContext map[string]string `json:"EncryptionContext,omitempty"`
	GrantTokens       []string          `json:"GrantTokens,omitempty"`
	KeyID             string            `json:"KeyId"`
	Plaintext         []byte            `json:"Plaintext"`
}

// EncryptResponse is undocumented.
type EncryptResponse struct {
	CiphertextBlob []byte `json:"CiphertextBlob,omitempty"`
	KeyID          string `json:"KeyId,omitempty"`
}

// GenerateDataKeyRequest is undocumented.
type GenerateDataKeyRequest struct {
	EncryptionContext map[string]string `json:"EncryptionContext,omitempty"`
	GrantTokens       []string          `json:"GrantTokens,omitempty"`
	KeyID             string            `json:"KeyId"`
	KeySpec           string            `json:"KeySpec,omitempty"`
	NumberOfBytes     int               `json:"NumberOfBytes,omitempty"`
}

// GenerateDataKeyResponse is undocumented.
type GenerateDataKeyResponse struct {
	CiphertextBlob []byte `json:"CiphertextBlob,omitempty"`
	KeyID          string `json:"KeyId,omitempty"`
	Plaintext      []byte `json:"Plaintext,omitempty"`
}

// GenerateDataKeyWithoutPlaintextRequest is undocumented.
type GenerateDataKeyWithoutPlaintextRequest struct {
	EncryptionContext map[string]string `json:"EncryptionContext,omitempty"`
	GrantTokens       []string          `json:"GrantTokens,omitempty"`
	KeyID             string            `json:"KeyId"`
	KeySpec           string            `json:"KeySpec,omitempty"`
	NumberOfBytes     int               `json:"NumberOfBytes,omitempty"`
}

// GenerateDataKeyWithoutPlaintextResponse is undocumented.
type GenerateDataKeyWithoutPlaintextResponse struct {
	CiphertextBlob []byte `json:"CiphertextBlob,omitempty"`
	KeyID          string `json:"KeyId,omitempty"`
}

// GenerateRandomRequest is undocumented.
type GenerateRandomRequest struct {
	NumberOfBytes int `json:"NumberOfBytes,omitempty"`
}

// GenerateRandomResponse is undocumented.
type GenerateRandomResponse struct {
	Plaintext []byte `json:"Plaintext,omitempty"`
}

// GetKeyPolicyRequest is undocumented.
type GetKeyPolicyRequest struct {
	KeyID      string `json:"KeyId"`
	PolicyName string `json:"PolicyName"`
}

// GetKeyPolicyResponse is undocumented.
type GetKeyPolicyResponse struct {
	Policy string `json:"Policy,omitempty"`
}

// GetKeyRotationStatusRequest is undocumented.
type GetKeyRotationStatusRequest struct {
	KeyID string `json:"KeyId"`
}

// GetKeyRotationStatusResponse is undocumented.
type GetKeyRotationStatusResponse struct {
	KeyRotationEnabled bool `json:"KeyRotationEnabled,omitempty"`
}

// GrantConstraints is undocumented.
type GrantConstraints struct {
	EncryptionContextEquals map[string]string `json:"EncryptionContextEquals,omitempty"`
	EncryptionContextSubset map[string]string `json:"EncryptionContextSubset,omitempty"`
}

// GrantListEntry is undocumented.
type GrantListEntry struct {
	Constraints       GrantConstraints `json:"Constraints,omitempty"`
	GrantID           string           `json:"GrantId,omitempty"`
	GranteePrincipal  string           `json:"GranteePrincipal,omitempty"`
	IssuingAccount    string           `json:"IssuingAccount,omitempty"`
	Operations        []string         `json:"Operations,omitempty"`
	RetiringPrincipal string           `json:"RetiringPrincipal,omitempty"`
}

// KeyListEntry is undocumented.
type KeyListEntry struct {
	KeyARN string `json:"KeyArn,omitempty"`
	KeyID  string `json:"KeyId,omitempty"`
}

// KeyMetadata is undocumented.
type KeyMetadata struct {
	AWSAccountID string    `json:"AWSAccountId,omitempty"`
	ARN          string    `json:"Arn,omitempty"`
	CreationDate time.Time `json:"CreationDate,omitempty"`
	Description  string    `json:"Description,omitempty"`
	Enabled      bool      `json:"Enabled,omitempty"`
	KeyID        string    `json:"KeyId"`
	KeyUsage     string    `json:"KeyUsage,omitempty"`
}

// ListAliasesRequest is undocumented.
type ListAliasesRequest struct {
	Limit  int    `json:"Limit,omitempty"`
	Marker string `json:"Marker,omitempty"`
}

// ListAliasesResponse is undocumented.
type ListAliasesResponse struct {
	Aliases    []AliasListEntry `json:"Aliases,omitempty"`
	NextMarker string           `json:"NextMarker,omitempty"`
	Truncated  bool             `json:"Truncated,omitempty"`
}

// ListGrantsRequest is undocumented.
type ListGrantsRequest struct {
	KeyID  string `json:"KeyId"`
	Limit  int    `json:"Limit,omitempty"`
	Marker string `json:"Marker,omitempty"`
}

// ListGrantsResponse is undocumented.
type ListGrantsResponse struct {
	Grants     []GrantListEntry `json:"Grants,omitempty"`
	NextMarker string           `json:"NextMarker,omitempty"`
	Truncated  bool             `json:"Truncated,omitempty"`
}

// ListKeyPoliciesRequest is undocumented.
type ListKeyPoliciesRequest struct {
	KeyID  string `json:"KeyId"`
	Limit  int    `json:"Limit,omitempty"`
	Marker string `json:"Marker,omitempty"`
}

// ListKeyPoliciesResponse is undocumented.
type ListKeyPoliciesResponse struct {
	NextMarker  string   `json:"NextMarker,omitempty"`
	PolicyNames []string `json:"PolicyNames,omitempty"`
	Truncated   bool     `json:"Truncated,omitempty"`
}

// ListKeysRequest is undocumented.
type ListKeysRequest struct {
	Limit  int    `json:"Limit,omitempty"`
	Marker string `json:"Marker,omitempty"`
}

// ListKeysResponse is undocumented.
type ListKeysResponse struct {
	Keys       []KeyListEntry `json:"Keys,omitempty"`
	NextMarker string         `json:"NextMarker,omitempty"`
	Truncated  bool           `json:"Truncated,omitempty"`
}

// PutKeyPolicyRequest is undocumented.
type PutKeyPolicyRequest struct {
	KeyID      string `json:"KeyId"`
	Policy     string `json:"Policy"`
	PolicyName string `json:"PolicyName"`
}

// ReEncryptRequest is undocumented.
type ReEncryptRequest struct {
	CiphertextBlob               []byte            `json:"CiphertextBlob"`
	DestinationEncryptionContext map[string]string `json:"DestinationEncryptionContext,omitempty"`
	DestinationKeyID             string            `json:"DestinationKeyId"`
	GrantTokens                  []string          `json:"GrantTokens,omitempty"`
	SourceEncryptionContext      map[string]string `json:"SourceEncryptionContext,omitempty"`
}

// ReEncryptResponse is undocumented.
type ReEncryptResponse struct {
	CiphertextBlob []byte `json:"CiphertextBlob,omitempty"`
	KeyID          string `json:"KeyId,omitempty"`
	SourceKeyID    string `json:"SourceKeyId,omitempty"`
}

// RetireGrantRequest is undocumented.
type RetireGrantRequest struct {
	GrantToken string `json:"GrantToken"`
}

// RevokeGrantRequest is undocumented.
type RevokeGrantRequest struct {
	GrantID string `json:"GrantId"`
	KeyID   string `json:"KeyId"`
}

// UpdateKeyDescriptionRequest is undocumented.
type UpdateKeyDescriptionRequest struct {
	Description string `json:"Description"`
	KeyID       string `json:"KeyId"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
