// Package iam provides a client for AWS Identity and Access Management.
package iam

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// IAM is a client for AWS Identity and Access Management.
type IAM struct {
	client aws.Client
}

// New returns a new IAM client.
func New(key, secret, region string, client *http.Client) *IAM {
	if client == nil {
		client = http.DefaultClient
	}

	return &IAM{
		client: &aws.QueryClient{
			Client: client,
			Auth: aws.Auth{
				Key:     key,
				Secret:  secret,
				Service: "iam",
				Region:  region,
			},
			Endpoint:   endpoints.Lookup("iam", region),
			APIVersion: "2010-05-08",
		},
	}
}

// AddClientIDToOpenIDConnectProvider adds a new client ID (also known as
// audience) to the list of client IDs already registered for the specified
// IAM OpenID Connect provider. This action is idempotent; it does not fail
// or return an error if you add an existing client ID to the provider.
func (c *IAM) AddClientIDToOpenIDConnectProvider(req AddClientIDToOpenIDConnectProviderRequest) (err error) {
	// NRE
	err = c.client.Do("AddClientIDToOpenIDConnectProvider", "POST", "/", req, nil)
	return
}

// AddRoleToInstanceProfile adds the specified role to the specified
// instance profile. For more information about roles, go to Working with
// Roles . For more information about instance profiles, go to About
// Instance Profiles .
func (c *IAM) AddRoleToInstanceProfile(req AddRoleToInstanceProfileRequest) (err error) {
	// NRE
	err = c.client.Do("AddRoleToInstanceProfile", "POST", "/", req, nil)
	return
}

// AddUserToGroup is undocumented.
func (c *IAM) AddUserToGroup(req AddUserToGroupRequest) (err error) {
	// NRE
	err = c.client.Do("AddUserToGroup", "POST", "/", req, nil)
	return
}

// ChangePassword changes the password of the IAM user who is calling this
// action. The root account password is not affected by this action. To
// change the password for a different user, see UpdateLoginProfile . For
// more information about modifying passwords, see Managing Passwords in
// the Using guide.
func (c *IAM) ChangePassword(req ChangePasswordRequest) (err error) {
	// NRE
	err = c.client.Do("ChangePassword", "POST", "/", req, nil)
	return
}

// CreateAccessKey creates a new AWS secret access key and corresponding
// AWS access key ID for the specified user. The default status for new
// keys is Active . If you do not specify a user name, IAM determines the
// user name implicitly based on the AWS access key ID signing the request.
// Because this action works for access keys under the AWS account, you can
// use this action to manage root credentials even if the AWS account has
// no associated users. For information about limits on the number of keys
// you can create, see Limitations on IAM Entities in the Using guide. To
// ensure the security of your AWS account, the secret access key is
// accessible only during key and user creation. You must save the key (for
// example, in a text file) if you want to be able to access it again. If a
// secret key is lost, you can delete the access keys for the associated
// user and then create new keys.
func (c *IAM) CreateAccessKey(req CreateAccessKeyRequest) (resp *CreateAccessKeyResult, err error) {
	resp = &CreateAccessKeyResult{}
	err = c.client.Do("CreateAccessKey", "POST", "/", req, resp)
	return
}

// CreateAccountAlias creates an alias for your AWS account. For
// information about using an AWS account alias, see Using an Alias for
// Your AWS Account in the Using guide.
func (c *IAM) CreateAccountAlias(req CreateAccountAliasRequest) (err error) {
	// NRE
	err = c.client.Do("CreateAccountAlias", "POST", "/", req, nil)
	return
}

// CreateGroup creates a new group. For information about the number of
// groups you can create, see Limitations on IAM Entities in the Using
// guide.
func (c *IAM) CreateGroup(req CreateGroupRequest) (resp *CreateGroupResult, err error) {
	resp = &CreateGroupResult{}
	err = c.client.Do("CreateGroup", "POST", "/", req, resp)
	return
}

// CreateInstanceProfile creates a new instance profile. For information
// about instance profiles, go to About Instance Profiles . For information
// about the number of instance profiles you can create, see Limitations on
// IAM Entities in the Using guide.
func (c *IAM) CreateInstanceProfile(req CreateInstanceProfileRequest) (resp *CreateInstanceProfileResult, err error) {
	resp = &CreateInstanceProfileResult{}
	err = c.client.Do("CreateInstanceProfile", "POST", "/", req, resp)
	return
}

// CreateLoginProfile creates a password for the specified user, giving the
// user the ability to access AWS services through the AWS Management
// Console. For more information about managing passwords, see Managing
// Passwords in the Using guide.
func (c *IAM) CreateLoginProfile(req CreateLoginProfileRequest) (resp *CreateLoginProfileResult, err error) {
	resp = &CreateLoginProfileResult{}
	err = c.client.Do("CreateLoginProfile", "POST", "/", req, resp)
	return
}

// CreateOpenIDConnectProvider creates an IAM entity to describe an
// identity provider (IdP) that supports OpenID Connect . The provider that
// you create with this operation can be used as a principal in a role's
// trust policy to establish a trust relationship between AWS and the
// provider. When you create the IAM provider, you specify the URL of the
// identity provider (IdP) to trust, a list of client IDs (also known as
// audiences) that identify the application or applications that are
// allowed to authenticate using the provider, and a list of thumbprints of
// the server certificate(s) that the IdP uses. You get all of this
// information from the IdP that you want to use for access to Because
// trust for the provider is ultimately derived from the IAM provider that
// this action creates, it is a best practice to limit access to the
// CreateOpenIDConnectProvider action to highly-privileged users.
func (c *IAM) CreateOpenIDConnectProvider(req CreateOpenIDConnectProviderRequest) (resp *CreateOpenIDConnectProviderResult, err error) {
	resp = &CreateOpenIDConnectProviderResult{}
	err = c.client.Do("CreateOpenIDConnectProvider", "POST", "/", req, resp)
	return
}

// CreateRole creates a new role for your AWS account. For more information
// about roles, go to Working with Roles . For information about
// limitations on role names and the number of roles you can create, go to
// Limitations on IAM Entities in the Using guide. The example policy
// grants permission to an EC2 instance to assume the role. The policy is
// URL-encoded according to RFC 3986. For more information about RFC 3986,
// go to http://www.faqs.org/rfcs/rfc3986.html .
func (c *IAM) CreateRole(req CreateRoleRequest) (resp *CreateRoleResult, err error) {
	resp = &CreateRoleResult{}
	err = c.client.Do("CreateRole", "POST", "/", req, resp)
	return
}

// CreateSAMLProvider creates an IAM entity to describe an identity
// provider (IdP) that supports 2.0. The provider that you create with this
// operation can be used as a principal in a role's trust policy to
// establish a trust relationship between AWS and a identity provider. You
// can create an IAM role that supports Web-based single sign-on to the AWS
// Management Console or one that supports API access to When you create
// the provider, you upload an a metadata document that you get from your
// IdP and that includes the issuer's name, expiration information, and
// keys that can be used to validate the authentication response
// (assertions) that are received from the IdP. You must generate the
// metadata document using the identity management software that is used as
// your organization's IdP. This operation requires Signature Version 4 .
// For more information, see Giving Console Access Using and Creating
// Temporary Security Credentials for Federation in the Using Temporary
// Credentials guide.
func (c *IAM) CreateSAMLProvider(req CreateSAMLProviderRequest) (resp *CreateSAMLProviderResult, err error) {
	resp = &CreateSAMLProviderResult{}
	err = c.client.Do("CreateSAMLProvider", "POST", "/", req, resp)
	return
}

// CreateUser creates a new user for your AWS account. For information
// about limitations on the number of users you can create, see Limitations
// on IAM Entities in the Using guide.
func (c *IAM) CreateUser(req CreateUserRequest) (resp *CreateUserResult, err error) {
	resp = &CreateUserResult{}
	err = c.client.Do("CreateUser", "POST", "/", req, resp)
	return
}

// CreateVirtualMFADevice creates a new virtual MFA device for the AWS
// account. After creating the virtual use EnableMFADevice to attach the
// MFA device to an IAM user. For more information about creating and
// working with virtual MFA devices, go to Using a Virtual MFA Device in
// the Using guide. For information about limits on the number of MFA
// devices you can create, see Limitations on Entities in the Using guide.
// The seed information contained in the QR code and the Base32 string
// should be treated like any other secret access information, such as your
// AWS access keys or your passwords. After you provision your virtual
// device, you should ensure that the information is destroyed following
// secure procedures.
func (c *IAM) CreateVirtualMFADevice(req CreateVirtualMFADeviceRequest) (resp *CreateVirtualMFADeviceResult, err error) {
	resp = &CreateVirtualMFADeviceResult{}
	err = c.client.Do("CreateVirtualMFADevice", "POST", "/", req, resp)
	return
}

// DeactivateMFADevice deactivates the specified MFA device and removes it
// from association with the user name for which it was originally enabled.
// For more information about creating and working with virtual MFA
// devices, go to Using a Virtual MFA Device in the Using guide.
func (c *IAM) DeactivateMFADevice(req DeactivateMFADeviceRequest) (err error) {
	// NRE
	err = c.client.Do("DeactivateMFADevice", "POST", "/", req, nil)
	return
}

// DeleteAccessKey deletes the access key associated with the specified
// user. If you do not specify a user name, IAM determines the user name
// implicitly based on the AWS access key ID signing the request. Because
// this action works for access keys under the AWS account, you can use
// this action to manage root credentials even if the AWS account has no
// associated users.
func (c *IAM) DeleteAccessKey(req DeleteAccessKeyRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteAccessKey", "POST", "/", req, nil)
	return
}

// DeleteAccountAlias deletes the specified AWS account alias. For
// information about using an AWS account alias, see Using an Alias for
// Your AWS Account in the Using guide.
func (c *IAM) DeleteAccountAlias(req DeleteAccountAliasRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteAccountAlias", "POST", "/", req, nil)
	return
}

// DeleteAccountPasswordPolicy is undocumented.
func (c *IAM) DeleteAccountPasswordPolicy() (err error) {
	// NRE
	err = c.client.Do("DeleteAccountPasswordPolicy", "POST", "/", nil, nil)
	return
}

// DeleteGroup deletes the specified group. The group must not contain any
// users or have any attached policies.
func (c *IAM) DeleteGroup(req DeleteGroupRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteGroup", "POST", "/", req, nil)
	return
}

// DeleteGroupPolicy deletes the specified policy that is associated with
// the specified group.
func (c *IAM) DeleteGroupPolicy(req DeleteGroupPolicyRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteGroupPolicy", "POST", "/", req, nil)
	return
}

// DeleteInstanceProfile deletes the specified instance profile. The
// instance profile must not have an associated role. Make sure you do not
// have any Amazon EC2 instances running with the instance profile you are
// about to delete. Deleting a role or instance profile that is associated
// with a running instance will break any applications running on the
// instance. For more information about instance profiles, go to About
// Instance Profiles .
func (c *IAM) DeleteInstanceProfile(req DeleteInstanceProfileRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteInstanceProfile", "POST", "/", req, nil)
	return
}

// DeleteLoginProfile deletes the password for the specified user, which
// terminates the user's ability to access AWS services through the AWS
// Management Console. Deleting a user's password does not prevent a user
// from accessing IAM through the command line interface or the To prevent
// all user access you must also either make the access key inactive or
// delete it. For more information about making keys inactive or deleting
// them, see UpdateAccessKey and DeleteAccessKey .
func (c *IAM) DeleteLoginProfile(req DeleteLoginProfileRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteLoginProfile", "POST", "/", req, nil)
	return
}

// DeleteOpenIDConnectProvider deletes an IAM OpenID Connect identity
// provider. Deleting an provider does not update any roles that reference
// the provider as a principal in their trust policies. Any attempt to
// assume a role that references a provider that has been deleted will
// fail. This action is idempotent; it does not fail or return an error if
// you call the action for a provider that was already deleted.
func (c *IAM) DeleteOpenIDConnectProvider(req DeleteOpenIDConnectProviderRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteOpenIDConnectProvider", "POST", "/", req, nil)
	return
}

// DeleteRole deletes the specified role. The role must not have any
// policies attached. For more information about roles, go to Working with
// Roles . Make sure you do not have any Amazon EC2 instances running with
// the role you are about to delete. Deleting a role or instance profile
// that is associated with a running instance will break any applications
// running on the instance.
func (c *IAM) DeleteRole(req DeleteRoleRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteRole", "POST", "/", req, nil)
	return
}

// DeleteRolePolicy deletes the specified policy associated with the
// specified role.
func (c *IAM) DeleteRolePolicy(req DeleteRolePolicyRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteRolePolicy", "POST", "/", req, nil)
	return
}

// DeleteSAMLProvider deletes a provider. Deleting the provider does not
// update any roles that reference the provider as a principal in their
// trust policies. Any attempt to assume a role that references a provider
// that has been deleted will fail. This operation requires Signature
// Version 4 .
func (c *IAM) DeleteSAMLProvider(req DeleteSAMLProviderRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSAMLProvider", "POST", "/", req, nil)
	return
}

// DeleteServerCertificate deletes the specified server certificate. If you
// are using a server certificate with Elastic Load Balancing, deleting the
// certificate could have implications for your application. If Elastic
// Load Balancing doesn't detect the deletion of bound certificates, it may
// continue to use the certificates. This could cause Elastic Load
// Balancing to stop accepting traffic. We recommend that you remove the
// reference to the certificate from Elastic Load Balancing before using
// this command to delete the certificate. For more information, go to
// DeleteLoadBalancerListeners in the Elastic Load Balancing API Reference
// .
func (c *IAM) DeleteServerCertificate(req DeleteServerCertificateRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteServerCertificate", "POST", "/", req, nil)
	return
}

// DeleteSigningCertificate deletes the specified signing certificate
// associated with the specified user. If you do not specify a user name,
// IAM determines the user name implicitly based on the AWS access key ID
// signing the request. Because this action works for access keys under the
// AWS account, you can use this action to manage root credentials even if
// the AWS account has no associated users.
func (c *IAM) DeleteSigningCertificate(req DeleteSigningCertificateRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSigningCertificate", "POST", "/", req, nil)
	return
}

// DeleteUser deletes the specified user. The user must not belong to any
// groups, have any keys or signing certificates, or have any attached
// policies.
func (c *IAM) DeleteUser(req DeleteUserRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteUser", "POST", "/", req, nil)
	return
}

// DeleteUserPolicy deletes the specified policy associated with the
// specified user.
func (c *IAM) DeleteUserPolicy(req DeleteUserPolicyRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteUserPolicy", "POST", "/", req, nil)
	return
}

// DeleteVirtualMFADevice deletes a virtual MFA device. You must deactivate
// a user's virtual MFA device before you can delete it. For information
// about deactivating MFA devices, see DeactivateMFADevice .
func (c *IAM) DeleteVirtualMFADevice(req DeleteVirtualMFADeviceRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVirtualMFADevice", "POST", "/", req, nil)
	return
}

// EnableMFADevice enables the specified MFA device and associates it with
// the specified user name. When enabled, the MFA device is required for
// every subsequent login by the user name associated with the device.
func (c *IAM) EnableMFADevice(req EnableMFADeviceRequest) (err error) {
	// NRE
	err = c.client.Do("EnableMFADevice", "POST", "/", req, nil)
	return
}

// GenerateCredentialReport generates a credential report for the AWS
// account. For more information about the credential report, see Getting
// Credential Reports in the Using guide.
func (c *IAM) GenerateCredentialReport() (resp *GenerateCredentialReportResult, err error) {
	resp = &GenerateCredentialReportResult{}
	err = c.client.Do("GenerateCredentialReport", "POST", "/", nil, resp)
	return
}

// GetAccountAuthorizationDetails retrieves information about all IAM
// users, groups, and roles in your account, including their relationships
// to one another and their attached policies. Use this API to obtain a
// snapshot of the configuration of IAM permissions (users, groups, roles,
// and policies) in your account. You can optionally filter the results
// using the Filter parameter. You can paginate the results using the
// MaxItems and Marker parameters.
func (c *IAM) GetAccountAuthorizationDetails(req GetAccountAuthorizationDetailsRequest) (resp *GetAccountAuthorizationDetailsResult, err error) {
	resp = &GetAccountAuthorizationDetailsResult{}
	err = c.client.Do("GetAccountAuthorizationDetails", "POST", "/", req, resp)
	return
}

// GetAccountPasswordPolicy retrieves the password policy for the AWS
// account. For more information about using a password policy, go to
// Managing an IAM Password Policy .
func (c *IAM) GetAccountPasswordPolicy() (resp *GetAccountPasswordPolicyResult, err error) {
	resp = &GetAccountPasswordPolicyResult{}
	err = c.client.Do("GetAccountPasswordPolicy", "POST", "/", nil, resp)
	return
}

// GetAccountSummary retrieves account level information about account
// entity usage and IAM quotas. For information about limitations on IAM
// entities, see Limitations on IAM Entities in the Using guide.
func (c *IAM) GetAccountSummary() (resp *GetAccountSummaryResult, err error) {
	resp = &GetAccountSummaryResult{}
	err = c.client.Do("GetAccountSummary", "POST", "/", nil, resp)
	return
}

// GetCredentialReport retrieves a credential report for the AWS account.
// For more information about the credential report, see Getting Credential
// Reports in the Using guide.
func (c *IAM) GetCredentialReport() (resp *GetCredentialReportResult, err error) {
	resp = &GetCredentialReportResult{}
	err = c.client.Do("GetCredentialReport", "POST", "/", nil, resp)
	return
}

// GetGroup returns a list of users that are in the specified group. You
// can paginate the results using the MaxItems and Marker parameters.
func (c *IAM) GetGroup(req GetGroupRequest) (resp *GetGroupResult, err error) {
	resp = &GetGroupResult{}
	err = c.client.Do("GetGroup", "POST", "/", req, resp)
	return
}

// GetGroupPolicy retrieves the specified policy document for the specified
// group. The returned policy is URL-encoded according to RFC 3986. For
// more information about RFC 3986, go to
// http://www.faqs.org/rfcs/rfc3986.html .
func (c *IAM) GetGroupPolicy(req GetGroupPolicyRequest) (resp *GetGroupPolicyResult, err error) {
	resp = &GetGroupPolicyResult{}
	err = c.client.Do("GetGroupPolicy", "POST", "/", req, resp)
	return
}

// GetInstanceProfile retrieves information about the specified instance
// profile, including the instance profile's path, and role. For more
// information about instance profiles, go to About Instance Profiles . For
// more information about ARNs, go to ARNs .
func (c *IAM) GetInstanceProfile(req GetInstanceProfileRequest) (resp *GetInstanceProfileResult, err error) {
	resp = &GetInstanceProfileResult{}
	err = c.client.Do("GetInstanceProfile", "POST", "/", req, resp)
	return
}

// GetLoginProfile retrieves the user name and password-creation date for
// the specified user. If the user has not been assigned a password, the
// action returns a 404 NoSuchEntity ) error.
func (c *IAM) GetLoginProfile(req GetLoginProfileRequest) (resp *GetLoginProfileResult, err error) {
	resp = &GetLoginProfileResult{}
	err = c.client.Do("GetLoginProfile", "POST", "/", req, resp)
	return
}

// GetOpenIDConnectProvider returns information about the specified OpenID
// Connect provider.
func (c *IAM) GetOpenIDConnectProvider(req GetOpenIDConnectProviderRequest) (resp *GetOpenIDConnectProviderResult, err error) {
	resp = &GetOpenIDConnectProviderResult{}
	err = c.client.Do("GetOpenIDConnectProvider", "POST", "/", req, resp)
	return
}

// GetRole retrieves information about the specified role, including the
// role's path, and the policy granting permission to assume the role. For
// more information about ARNs, go to ARNs . For more information about
// roles, go to Working with Roles . The returned policy is URL-encoded
// according to RFC 3986. For more information about RFC 3986, go to
// http://www.faqs.org/rfcs/rfc3986.html .
func (c *IAM) GetRole(req GetRoleRequest) (resp *GetRoleResult, err error) {
	resp = &GetRoleResult{}
	err = c.client.Do("GetRole", "POST", "/", req, resp)
	return
}

// GetRolePolicy retrieves the specified policy document for the specified
// role. For more information about roles, go to Working with Roles . The
// returned policy is URL-encoded according to RFC 3986. For more
// information about RFC 3986, go to http://www.faqs.org/rfcs/rfc3986.html
// .
func (c *IAM) GetRolePolicy(req GetRolePolicyRequest) (resp *GetRolePolicyResult, err error) {
	resp = &GetRolePolicyResult{}
	err = c.client.Do("GetRolePolicy", "POST", "/", req, resp)
	return
}

// GetSAMLProvider returns the provider metadocument that was uploaded when
// the provider was created or updated. This operation requires Signature
// Version 4 .
func (c *IAM) GetSAMLProvider(req GetSAMLProviderRequest) (resp *GetSAMLProviderResult, err error) {
	resp = &GetSAMLProviderResult{}
	err = c.client.Do("GetSAMLProvider", "POST", "/", req, resp)
	return
}

// GetServerCertificate retrieves information about the specified server
// certificate.
func (c *IAM) GetServerCertificate(req GetServerCertificateRequest) (resp *GetServerCertificateResult, err error) {
	resp = &GetServerCertificateResult{}
	err = c.client.Do("GetServerCertificate", "POST", "/", req, resp)
	return
}

// GetUser retrieves information about the specified user, including the
// user's creation date, path, unique ID, and If you do not specify a user
// name, IAM determines the user name implicitly based on the AWS access
// key ID used to sign the request.
func (c *IAM) GetUser(req GetUserRequest) (resp *GetUserResult, err error) {
	resp = &GetUserResult{}
	err = c.client.Do("GetUser", "POST", "/", req, resp)
	return
}

// GetUserPolicy retrieves the specified policy document for the specified
// user. The returned policy is URL-encoded according to RFC 3986. For more
// information about RFC 3986, go to http://www.faqs.org/rfcs/rfc3986.html
// .
func (c *IAM) GetUserPolicy(req GetUserPolicyRequest) (resp *GetUserPolicyResult, err error) {
	resp = &GetUserPolicyResult{}
	err = c.client.Do("GetUserPolicy", "POST", "/", req, resp)
	return
}

// ListAccessKeys returns information about the access key IDs associated
// with the specified user. If there are none, the action returns an empty
// list. Although each user is limited to a small number of keys, you can
// still paginate the results using the MaxItems and Marker parameters. If
// the UserName field is not specified, the UserName is determined
// implicitly based on the AWS access key ID used to sign the request.
// Because this action works for access keys under the AWS account, you can
// use this action to manage root credentials even if the AWS account has
// no associated users. To ensure the security of your AWS account, the
// secret access key is accessible only during key and user creation.
func (c *IAM) ListAccessKeys(req ListAccessKeysRequest) (resp *ListAccessKeysResult, err error) {
	resp = &ListAccessKeysResult{}
	err = c.client.Do("ListAccessKeys", "POST", "/", req, resp)
	return
}

// ListAccountAliases lists the account aliases associated with the
// account. For information about using an AWS account alias, see Using an
// Alias for Your AWS Account in the Using guide. You can paginate the
// results using the MaxItems and Marker parameters.
func (c *IAM) ListAccountAliases(req ListAccountAliasesRequest) (resp *ListAccountAliasesResult, err error) {
	resp = &ListAccountAliasesResult{}
	err = c.client.Do("ListAccountAliases", "POST", "/", req, resp)
	return
}

// ListGroupPolicies lists the names of the policies associated with the
// specified group. If there are none, the action returns an empty list.
// You can paginate the results using the MaxItems and Marker parameters.
func (c *IAM) ListGroupPolicies(req ListGroupPoliciesRequest) (resp *ListGroupPoliciesResult, err error) {
	resp = &ListGroupPoliciesResult{}
	err = c.client.Do("ListGroupPolicies", "POST", "/", req, resp)
	return
}

// ListGroups lists the groups that have the specified path prefix. You can
// paginate the results using the MaxItems and Marker parameters.
func (c *IAM) ListGroups(req ListGroupsRequest) (resp *ListGroupsResult, err error) {
	resp = &ListGroupsResult{}
	err = c.client.Do("ListGroups", "POST", "/", req, resp)
	return
}

// ListGroupsForUser lists the groups the specified user belongs to. You
// can paginate the results using the MaxItems and Marker parameters.
func (c *IAM) ListGroupsForUser(req ListGroupsForUserRequest) (resp *ListGroupsForUserResult, err error) {
	resp = &ListGroupsForUserResult{}
	err = c.client.Do("ListGroupsForUser", "POST", "/", req, resp)
	return
}

// ListInstanceProfiles lists the instance profiles that have the specified
// path prefix. If there are none, the action returns an empty list. For
// more information about instance profiles, go to About Instance Profiles
// . You can paginate the results using the MaxItems and Marker parameters.
func (c *IAM) ListInstanceProfiles(req ListInstanceProfilesRequest) (resp *ListInstanceProfilesResult, err error) {
	resp = &ListInstanceProfilesResult{}
	err = c.client.Do("ListInstanceProfiles", "POST", "/", req, resp)
	return
}

// ListInstanceProfilesForRole lists the instance profiles that have the
// specified associated role. If there are none, the action returns an
// empty list. For more information about instance profiles, go to About
// Instance Profiles . You can paginate the results using the MaxItems and
// Marker parameters.
func (c *IAM) ListInstanceProfilesForRole(req ListInstanceProfilesForRoleRequest) (resp *ListInstanceProfilesForRoleResult, err error) {
	resp = &ListInstanceProfilesForRoleResult{}
	err = c.client.Do("ListInstanceProfilesForRole", "POST", "/", req, resp)
	return
}

// ListMFADevices lists the MFA devices. If the request includes the user
// name, then this action lists all the MFA devices associated with the
// specified user name. If you do not specify a user name, IAM determines
// the user name implicitly based on the AWS access key ID signing the
// request. You can paginate the results using the MaxItems and Marker
// parameters.
func (c *IAM) ListMFADevices(req ListMFADevicesRequest) (resp *ListMFADevicesResult, err error) {
	resp = &ListMFADevicesResult{}
	err = c.client.Do("ListMFADevices", "POST", "/", req, resp)
	return
}

// ListOpenIDConnectProviders lists information about the OpenID Connect
// providers in the AWS account.
func (c *IAM) ListOpenIDConnectProviders(req ListOpenIDConnectProvidersRequest) (resp *ListOpenIDConnectProvidersResult, err error) {
	resp = &ListOpenIDConnectProvidersResult{}
	err = c.client.Do("ListOpenIDConnectProviders", "POST", "/", req, resp)
	return
}

// ListRolePolicies lists the names of the policies associated with the
// specified role. If there are none, the action returns an empty list. You
// can paginate the results using the MaxItems and Marker parameters.
func (c *IAM) ListRolePolicies(req ListRolePoliciesRequest) (resp *ListRolePoliciesResult, err error) {
	resp = &ListRolePoliciesResult{}
	err = c.client.Do("ListRolePolicies", "POST", "/", req, resp)
	return
}

// ListRoles lists the roles that have the specified path prefix. If there
// are none, the action returns an empty list. For more information about
// roles, go to Working with Roles . You can paginate the results using the
// MaxItems and Marker parameters. The returned policy is URL-encoded
// according to RFC 3986. For more information about RFC 3986, go to
// http://www.faqs.org/rfcs/rfc3986.html .
func (c *IAM) ListRoles(req ListRolesRequest) (resp *ListRolesResult, err error) {
	resp = &ListRolesResult{}
	err = c.client.Do("ListRoles", "POST", "/", req, resp)
	return
}

// ListSAMLProviders is undocumented.
func (c *IAM) ListSAMLProviders(req ListSAMLProvidersRequest) (resp *ListSAMLProvidersResult, err error) {
	resp = &ListSAMLProvidersResult{}
	err = c.client.Do("ListSAMLProviders", "POST", "/", req, resp)
	return
}

// ListServerCertificates lists the server certificates that have the
// specified path prefix. If none exist, the action returns an empty list.
// You can paginate the results using the MaxItems and Marker parameters.
func (c *IAM) ListServerCertificates(req ListServerCertificatesRequest) (resp *ListServerCertificatesResult, err error) {
	resp = &ListServerCertificatesResult{}
	err = c.client.Do("ListServerCertificates", "POST", "/", req, resp)
	return
}

// ListSigningCertificates returns information about the signing
// certificates associated with the specified user. If there are none, the
// action returns an empty list. Although each user is limited to a small
// number of signing certificates, you can still paginate the results using
// the MaxItems and Marker parameters. If the UserName field is not
// specified, the user name is determined implicitly based on the AWS
// access key ID used to sign the request. Because this action works for
// access keys under the AWS account, you can use this action to manage
// root credentials even if the AWS account has no associated users.
func (c *IAM) ListSigningCertificates(req ListSigningCertificatesRequest) (resp *ListSigningCertificatesResult, err error) {
	resp = &ListSigningCertificatesResult{}
	err = c.client.Do("ListSigningCertificates", "POST", "/", req, resp)
	return
}

// ListUserPolicies lists the names of the policies associated with the
// specified user. If there are none, the action returns an empty list. You
// can paginate the results using the MaxItems and Marker parameters.
func (c *IAM) ListUserPolicies(req ListUserPoliciesRequest) (resp *ListUserPoliciesResult, err error) {
	resp = &ListUserPoliciesResult{}
	err = c.client.Do("ListUserPolicies", "POST", "/", req, resp)
	return
}

// ListUsers lists the IAM users that have the specified path prefix. If no
// path prefix is specified, the action returns all users in the AWS
// account. If there are none, the action returns an empty list. You can
// paginate the results using the MaxItems and Marker parameters.
func (c *IAM) ListUsers(req ListUsersRequest) (resp *ListUsersResult, err error) {
	resp = &ListUsersResult{}
	err = c.client.Do("ListUsers", "POST", "/", req, resp)
	return
}

// ListVirtualMFADevices lists the virtual MFA devices under the AWS
// account by assignment status. If you do not specify an assignment
// status, the action returns a list of all virtual MFA devices. Assignment
// status can be Assigned , Unassigned , or Any . You can paginate the
// results using the MaxItems and Marker parameters.
func (c *IAM) ListVirtualMFADevices(req ListVirtualMFADevicesRequest) (resp *ListVirtualMFADevicesResult, err error) {
	resp = &ListVirtualMFADevicesResult{}
	err = c.client.Do("ListVirtualMFADevices", "POST", "/", req, resp)
	return
}

// PutGroupPolicy adds (or updates) a policy document associated with the
// specified group. For information about policies, refer to Overview of
// Policies in the Using guide. For information about limits on the number
// of policies you can associate with a group, see Limitations on IAM
// Entities in the Using guide. Because policy documents can be large, you
// should use rather than GET when calling PutGroupPolicy . For information
// about setting up signatures and authorization through the go to Signing
// AWS API Requests in the AWS General Reference . For general information
// about using the Query API with go to Making Query Requests in the Using
// guide.
func (c *IAM) PutGroupPolicy(req PutGroupPolicyRequest) (err error) {
	// NRE
	err = c.client.Do("PutGroupPolicy", "POST", "/", req, nil)
	return
}

// PutRolePolicy adds (or updates) a policy document associated with the
// specified role. For information about policies, go to Overview of
// Policies in the Using guide. For information about limits on the
// policies you can associate with a role, see Limitations on IAM Entities
// in the Using guide. Because policy documents can be large, you should
// use rather than GET when calling PutRolePolicy . For information about
// setting up signatures and authorization through the go to Signing AWS
// API Requests in the AWS General Reference . For general information
// about using the Query API with go to Making Query Requests in the Using
// guide.
func (c *IAM) PutRolePolicy(req PutRolePolicyRequest) (err error) {
	// NRE
	err = c.client.Do("PutRolePolicy", "POST", "/", req, nil)
	return
}

// PutUserPolicy adds (or updates) a policy document associated with the
// specified user. For information about policies, refer to Overview of
// Policies in the Using guide. For information about limits on the number
// of policies you can associate with a user, see Limitations on IAM
// Entities in the Using guide. Because policy documents can be large, you
// should use rather than GET when calling PutUserPolicy . For information
// about setting up signatures and authorization through the go to Signing
// AWS API Requests in the AWS General Reference . For general information
// about using the Query API with go to Making Query Requests in the Using
// guide.
func (c *IAM) PutUserPolicy(req PutUserPolicyRequest) (err error) {
	// NRE
	err = c.client.Do("PutUserPolicy", "POST", "/", req, nil)
	return
}

// RemoveClientIDFromOpenIDConnectProvider removes the specified client ID
// (also known as audience) from the list of client IDs registered for the
// specified IAM OpenID Connect provider. This action is idempotent; it
// does not fail or return an error if you try to remove a client ID that
// was removed previously.
func (c *IAM) RemoveClientIDFromOpenIDConnectProvider(req RemoveClientIDFromOpenIDConnectProviderRequest) (err error) {
	// NRE
	err = c.client.Do("RemoveClientIDFromOpenIDConnectProvider", "POST", "/", req, nil)
	return
}

// RemoveRoleFromInstanceProfile removes the specified role from the
// specified instance profile. Make sure you do not have any Amazon EC2
// instances running with the role you are about to remove from the
// instance profile. Removing a role from an instance profile that is
// associated with a running instance will break any applications running
// on the instance. For more information about roles, go to Working with
// Roles . For more information about instance profiles, go to About
// Instance Profiles .
func (c *IAM) RemoveRoleFromInstanceProfile(req RemoveRoleFromInstanceProfileRequest) (err error) {
	// NRE
	err = c.client.Do("RemoveRoleFromInstanceProfile", "POST", "/", req, nil)
	return
}

// RemoveUserFromGroup removes the specified user from the specified group.
func (c *IAM) RemoveUserFromGroup(req RemoveUserFromGroupRequest) (err error) {
	// NRE
	err = c.client.Do("RemoveUserFromGroup", "POST", "/", req, nil)
	return
}

// ResyncMFADevice synchronizes the specified MFA device with AWS servers.
// For more information about creating and working with virtual MFA
// devices, go to Using a Virtual MFA Device in the Using guide.
func (c *IAM) ResyncMFADevice(req ResyncMFADeviceRequest) (err error) {
	// NRE
	err = c.client.Do("ResyncMFADevice", "POST", "/", req, nil)
	return
}

// UpdateAccessKey changes the status of the specified access key from
// Active to Inactive, or vice versa. This action can be used to disable a
// user's key as part of a key rotation work flow. If the UserName field is
// not specified, the UserName is determined implicitly based on the AWS
// access key ID used to sign the request. Because this action works for
// access keys under the AWS account, you can use this action to manage
// root credentials even if the AWS account has no associated users. For
// information about rotating keys, see Managing Keys and Certificates in
// the Using guide.
func (c *IAM) UpdateAccessKey(req UpdateAccessKeyRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateAccessKey", "POST", "/", req, nil)
	return
}

// UpdateAccountPasswordPolicy updates the password policy settings for the
// AWS account. This action does not support partial updates. No parameters
// are required, but if you do not specify a parameter, that parameter's
// value reverts to its default value. See the Request Parameters section
// for each parameter's default value. For more information about using a
// password policy, see Managing an IAM Password Policy in the Using guide.
func (c *IAM) UpdateAccountPasswordPolicy(req UpdateAccountPasswordPolicyRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateAccountPasswordPolicy", "POST", "/", req, nil)
	return
}

// UpdateAssumeRolePolicy updates the policy that grants an entity
// permission to assume a role. For more information about roles, go to
// Working with Roles .
func (c *IAM) UpdateAssumeRolePolicy(req UpdateAssumeRolePolicyRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateAssumeRolePolicy", "POST", "/", req, nil)
	return
}

// UpdateGroup updates the name and/or the path of the specified group. You
// should understand the implications of changing a group's path or name.
// For more information, see Renaming Users and Groups in the Using guide.
// To change a group name the requester must have appropriate permissions
// on both the source object and the target object. For example, to change
// Managers to MGRs, the entity making the request must have permission on
// Managers and MGRs, or must have permission on all For more information
// about permissions, see Permissions and Policies .
func (c *IAM) UpdateGroup(req UpdateGroupRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateGroup", "POST", "/", req, nil)
	return
}

// UpdateLoginProfile changes the password for the specified user. Users
// can change their own passwords by calling ChangePassword . For more
// information about modifying passwords, see Managing Passwords in the
// Using guide.
func (c *IAM) UpdateLoginProfile(req UpdateLoginProfileRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateLoginProfile", "POST", "/", req, nil)
	return
}

// UpdateOpenIDConnectProviderThumbprint replaces the existing list of
// server certificate thumbprints with a new list. The list that you pass
// with this action completely replaces the existing list of thumbprints.
// (The lists are not merged.) Typically, you need to update a thumbprint
// only when the identity provider's certificate changes, which occurs
// rarely. However, if the provider's certificate does change, any attempt
// to assume an IAM role that specifies the IAM provider as a principal
// will fail until the certificate thumbprint is updated. Because trust for
// the OpenID Connect provider is ultimately derived from the provider's
// certificate and is validated by the thumbprint, it is a best practice to
// limit access to the UpdateOpenIDConnectProviderThumbprint action to
// highly-privileged users.
func (c *IAM) UpdateOpenIDConnectProviderThumbprint(req UpdateOpenIDConnectProviderThumbprintRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateOpenIDConnectProviderThumbprint", "POST", "/", req, nil)
	return
}

// UpdateSAMLProvider updates the metadata document for an existing
// provider. This operation requires Signature Version 4 .
func (c *IAM) UpdateSAMLProvider(req UpdateSAMLProviderRequest) (resp *UpdateSAMLProviderResult, err error) {
	resp = &UpdateSAMLProviderResult{}
	err = c.client.Do("UpdateSAMLProvider", "POST", "/", req, resp)
	return
}

// UpdateServerCertificate updates the name and/or the path of the
// specified server certificate. You should understand the implications of
// changing a server certificate's path or name. For more information, see
// Managing Server Certificates in the Using guide. To change a server
// certificate name the requester must have appropriate permissions on both
// the source object and the target object. For example, to change the name
// from ProductionCert to ProdCert, the entity making the request must have
// permission on ProductionCert and ProdCert, or must have permission on
// all For more information about permissions, see Permissions and Policies
// .
func (c *IAM) UpdateServerCertificate(req UpdateServerCertificateRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateServerCertificate", "POST", "/", req, nil)
	return
}

// UpdateSigningCertificate changes the status of the specified signing
// certificate from active to disabled, or vice versa. This action can be
// used to disable a user's signing certificate as part of a certificate
// rotation work flow. If the UserName field is not specified, the UserName
// is determined implicitly based on the AWS access key ID used to sign the
// request. Because this action works for access keys under the AWS
// account, you can use this action to manage root credentials even if the
// AWS account has no associated users. For information about rotating
// certificates, see Managing Keys and Certificates in the Using guide.
func (c *IAM) UpdateSigningCertificate(req UpdateSigningCertificateRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateSigningCertificate", "POST", "/", req, nil)
	return
}

// UpdateUser updates the name and/or the path of the specified user. You
// should understand the implications of changing a user's path or name.
// For more information, see Renaming Users and Groups in the Using guide.
// To change a user name the requester must have appropriate permissions on
// both the source object and the target object. For example, to change Bob
// to Robert, the entity making the request must have permission on Bob and
// Robert, or must have permission on all For more information about
// permissions, see Permissions and Policies .
func (c *IAM) UpdateUser(req UpdateUserRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateUser", "POST", "/", req, nil)
	return
}

// UploadServerCertificate uploads a server certificate entity for the AWS
// account. The server certificate entity includes a public key
// certificate, a private key, and an optional certificate chain, which
// should all be PEM-encoded. For information about the number of server
// certificates you can upload, see Limitations on IAM Entities in the
// Using guide. Because the body of the public key certificate, private
// key, and the certificate chain can be large, you should use rather than
// GET when calling UploadServerCertificate . For information about setting
// up signatures and authorization through the go to Signing AWS API
// Requests in the AWS General Reference . For general information about
// using the Query API with go to Making Query Requests in the Using guide.
func (c *IAM) UploadServerCertificate(req UploadServerCertificateRequest) (resp *UploadServerCertificateResult, err error) {
	resp = &UploadServerCertificateResult{}
	err = c.client.Do("UploadServerCertificate", "POST", "/", req, resp)
	return
}

// UploadSigningCertificate uploads an X.509 signing certificate and
// associates it with the specified user. Some AWS services use X.509
// signing certificates to validate requests that are signed with a
// corresponding private key. When you upload the certificate, its default
// status is Active . If the UserName field is not specified, the user name
// is determined implicitly based on the AWS access key ID used to sign the
// request. Because this action works for access keys under the AWS
// account, you can use this action to manage root credentials even if the
// AWS account has no associated users. Because the body of a X.509
// certificate can be large, you should use rather than GET when calling
// UploadSigningCertificate . For information about setting up signatures
// and authorization through the go to Signing AWS API Requests in the AWS
// General Reference . For general information about using the Query API
// with go to Making Query Requests in the Using guide.
func (c *IAM) UploadSigningCertificate(req UploadSigningCertificateRequest) (resp *UploadSigningCertificateResult, err error) {
	resp = &UploadSigningCertificateResult{}
	err = c.client.Do("UploadSigningCertificate", "POST", "/", req, resp)
	return
}

// AccessKey is undocumented.
type AccessKey struct {
	AccessKeyID     string    `xml:"AccessKeyId"`
	CreateDate      time.Time `xml:"CreateDate"`
	SecretAccessKey string    `xml:"SecretAccessKey"`
	Status          string    `xml:"Status"`
	UserName        string    `xml:"UserName"`
}

// AccessKeyMetadata is undocumented.
type AccessKeyMetadata struct {
	AccessKeyID string    `xml:"AccessKeyId"`
	CreateDate  time.Time `xml:"CreateDate"`
	Status      string    `xml:"Status"`
	UserName    string    `xml:"UserName"`
}

// AddClientIDToOpenIDConnectProviderRequest is undocumented.
type AddClientIDToOpenIDConnectProviderRequest struct {
	ClientID                 string `xml:"ClientID"`
	OpenIDConnectProviderARN string `xml:"OpenIDConnectProviderArn"`
}

// AddRoleToInstanceProfileRequest is undocumented.
type AddRoleToInstanceProfileRequest struct {
	InstanceProfileName string `xml:"InstanceProfileName"`
	RoleName            string `xml:"RoleName"`
}

// AddUserToGroupRequest is undocumented.
type AddUserToGroupRequest struct {
	GroupName string `xml:"GroupName"`
	UserName  string `xml:"UserName"`
}

// ChangePasswordRequest is undocumented.
type ChangePasswordRequest struct {
	NewPassword string `xml:"NewPassword"`
	OldPassword string `xml:"OldPassword"`
}

// CreateAccessKeyRequest is undocumented.
type CreateAccessKeyRequest struct {
	UserName string `xml:"UserName"`
}

// CreateAccessKeyResponse is undocumented.
type CreateAccessKeyResponse struct {
	AccessKey AccessKey `xml:"CreateAccessKeyResult>AccessKey"`
}

// CreateAccountAliasRequest is undocumented.
type CreateAccountAliasRequest struct {
	AccountAlias string `xml:"AccountAlias"`
}

// CreateGroupRequest is undocumented.
type CreateGroupRequest struct {
	GroupName string `xml:"GroupName"`
	Path      string `xml:"Path"`
}

// CreateGroupResponse is undocumented.
type CreateGroupResponse struct {
	Group Group `xml:"CreateGroupResult>Group"`
}

// CreateInstanceProfileRequest is undocumented.
type CreateInstanceProfileRequest struct {
	InstanceProfileName string `xml:"InstanceProfileName"`
	Path                string `xml:"Path"`
}

// CreateInstanceProfileResponse is undocumented.
type CreateInstanceProfileResponse struct {
	InstanceProfile InstanceProfile `xml:"CreateInstanceProfileResult>InstanceProfile"`
}

// CreateLoginProfileRequest is undocumented.
type CreateLoginProfileRequest struct {
	Password              string `xml:"Password"`
	PasswordResetRequired bool   `xml:"PasswordResetRequired"`
	UserName              string `xml:"UserName"`
}

// CreateLoginProfileResponse is undocumented.
type CreateLoginProfileResponse struct {
	LoginProfile LoginProfile `xml:"CreateLoginProfileResult>LoginProfile"`
}

// CreateOpenIDConnectProviderRequest is undocumented.
type CreateOpenIDConnectProviderRequest struct {
	ClientIDList   []string `xml:"ClientIDList>member"`
	ThumbprintList []string `xml:"ThumbprintList>member"`
	URL            string   `xml:"Url"`
}

// CreateOpenIDConnectProviderResponse is undocumented.
type CreateOpenIDConnectProviderResponse struct {
	OpenIDConnectProviderARN string `xml:"CreateOpenIDConnectProviderResult>OpenIDConnectProviderArn"`
}

// CreateRoleRequest is undocumented.
type CreateRoleRequest struct {
	AssumeRolePolicyDocument string `xml:"AssumeRolePolicyDocument"`
	Path                     string `xml:"Path"`
	RoleName                 string `xml:"RoleName"`
}

// CreateRoleResponse is undocumented.
type CreateRoleResponse struct {
	Role Role `xml:"CreateRoleResult>Role"`
}

// CreateSAMLProviderRequest is undocumented.
type CreateSAMLProviderRequest struct {
	Name                 string `xml:"Name"`
	SAMLMetadataDocument string `xml:"SAMLMetadataDocument"`
}

// CreateSAMLProviderResponse is undocumented.
type CreateSAMLProviderResponse struct {
	SAMLProviderARN string `xml:"CreateSAMLProviderResult>SAMLProviderArn"`
}

// CreateUserRequest is undocumented.
type CreateUserRequest struct {
	Path     string `xml:"Path"`
	UserName string `xml:"UserName"`
}

// CreateUserResponse is undocumented.
type CreateUserResponse struct {
	User User `xml:"CreateUserResult>User"`
}

// CreateVirtualMFADeviceRequest is undocumented.
type CreateVirtualMFADeviceRequest struct {
	Path                 string `xml:"Path"`
	VirtualMFADeviceName string `xml:"VirtualMFADeviceName"`
}

// CreateVirtualMFADeviceResponse is undocumented.
type CreateVirtualMFADeviceResponse struct {
	VirtualMFADevice VirtualMFADevice `xml:"CreateVirtualMFADeviceResult>VirtualMFADevice"`
}

// DeactivateMFADeviceRequest is undocumented.
type DeactivateMFADeviceRequest struct {
	SerialNumber string `xml:"SerialNumber"`
	UserName     string `xml:"UserName"`
}

// DeleteAccessKeyRequest is undocumented.
type DeleteAccessKeyRequest struct {
	AccessKeyID string `xml:"AccessKeyId"`
	UserName    string `xml:"UserName"`
}

// DeleteAccountAliasRequest is undocumented.
type DeleteAccountAliasRequest struct {
	AccountAlias string `xml:"AccountAlias"`
}

// DeleteGroupPolicyRequest is undocumented.
type DeleteGroupPolicyRequest struct {
	GroupName  string `xml:"GroupName"`
	PolicyName string `xml:"PolicyName"`
}

// DeleteGroupRequest is undocumented.
type DeleteGroupRequest struct {
	GroupName string `xml:"GroupName"`
}

// DeleteInstanceProfileRequest is undocumented.
type DeleteInstanceProfileRequest struct {
	InstanceProfileName string `xml:"InstanceProfileName"`
}

// DeleteLoginProfileRequest is undocumented.
type DeleteLoginProfileRequest struct {
	UserName string `xml:"UserName"`
}

// DeleteOpenIDConnectProviderRequest is undocumented.
type DeleteOpenIDConnectProviderRequest struct {
	OpenIDConnectProviderARN string `xml:"OpenIDConnectProviderArn"`
}

// DeleteRolePolicyRequest is undocumented.
type DeleteRolePolicyRequest struct {
	PolicyName string `xml:"PolicyName"`
	RoleName   string `xml:"RoleName"`
}

// DeleteRoleRequest is undocumented.
type DeleteRoleRequest struct {
	RoleName string `xml:"RoleName"`
}

// DeleteSAMLProviderRequest is undocumented.
type DeleteSAMLProviderRequest struct {
	SAMLProviderARN string `xml:"SAMLProviderArn"`
}

// DeleteServerCertificateRequest is undocumented.
type DeleteServerCertificateRequest struct {
	ServerCertificateName string `xml:"ServerCertificateName"`
}

// DeleteSigningCertificateRequest is undocumented.
type DeleteSigningCertificateRequest struct {
	CertificateID string `xml:"CertificateId"`
	UserName      string `xml:"UserName"`
}

// DeleteUserPolicyRequest is undocumented.
type DeleteUserPolicyRequest struct {
	PolicyName string `xml:"PolicyName"`
	UserName   string `xml:"UserName"`
}

// DeleteUserRequest is undocumented.
type DeleteUserRequest struct {
	UserName string `xml:"UserName"`
}

// DeleteVirtualMFADeviceRequest is undocumented.
type DeleteVirtualMFADeviceRequest struct {
	SerialNumber string `xml:"SerialNumber"`
}

// EnableMFADeviceRequest is undocumented.
type EnableMFADeviceRequest struct {
	AuthenticationCode1 string `xml:"AuthenticationCode1"`
	AuthenticationCode2 string `xml:"AuthenticationCode2"`
	SerialNumber        string `xml:"SerialNumber"`
	UserName            string `xml:"UserName"`
}

// GenerateCredentialReportResponse is undocumented.
type GenerateCredentialReportResponse struct {
	Description string `xml:"GenerateCredentialReportResult>Description"`
	State       string `xml:"GenerateCredentialReportResult>State"`
}

// GetAccountAuthorizationDetailsRequest is undocumented.
type GetAccountAuthorizationDetailsRequest struct {
	Filter   []string `xml:"Filter>member"`
	Marker   string   `xml:"Marker"`
	MaxItems int      `xml:"MaxItems"`
}

// GetAccountAuthorizationDetailsResponse is undocumented.
type GetAccountAuthorizationDetailsResponse struct {
	GroupDetailList []GroupDetail `xml:"GetAccountAuthorizationDetailsResult>GroupDetailList>member"`
	IsTruncated     bool          `xml:"GetAccountAuthorizationDetailsResult>IsTruncated"`
	Marker          string        `xml:"GetAccountAuthorizationDetailsResult>Marker"`
	RoleDetailList  []RoleDetail  `xml:"GetAccountAuthorizationDetailsResult>RoleDetailList>member"`
	UserDetailList  []UserDetail  `xml:"GetAccountAuthorizationDetailsResult>UserDetailList>member"`
}

// GetAccountPasswordPolicyResponse is undocumented.
type GetAccountPasswordPolicyResponse struct {
	PasswordPolicy PasswordPolicy `xml:"GetAccountPasswordPolicyResult>PasswordPolicy"`
}

// GetAccountSummaryResponse is undocumented.
type GetAccountSummaryResponse struct {
	SummaryMap map[string]int `xml:"GetAccountSummaryResult>SummaryMap"`
}

// GetCredentialReportResponse is undocumented.
type GetCredentialReportResponse struct {
	Content       []byte    `xml:"GetCredentialReportResult>Content"`
	GeneratedTime time.Time `xml:"GetCredentialReportResult>GeneratedTime"`
	ReportFormat  string    `xml:"GetCredentialReportResult>ReportFormat"`
}

// GetGroupPolicyRequest is undocumented.
type GetGroupPolicyRequest struct {
	GroupName  string `xml:"GroupName"`
	PolicyName string `xml:"PolicyName"`
}

// GetGroupPolicyResponse is undocumented.
type GetGroupPolicyResponse struct {
	GroupName      string `xml:"GetGroupPolicyResult>GroupName"`
	PolicyDocument string `xml:"GetGroupPolicyResult>PolicyDocument"`
	PolicyName     string `xml:"GetGroupPolicyResult>PolicyName"`
}

// GetGroupRequest is undocumented.
type GetGroupRequest struct {
	GroupName string `xml:"GroupName"`
	Marker    string `xml:"Marker"`
	MaxItems  int    `xml:"MaxItems"`
}

// GetGroupResponse is undocumented.
type GetGroupResponse struct {
	Group       Group  `xml:"GetGroupResult>Group"`
	IsTruncated bool   `xml:"GetGroupResult>IsTruncated"`
	Marker      string `xml:"GetGroupResult>Marker"`
	Users       []User `xml:"GetGroupResult>Users>member"`
}

// GetInstanceProfileRequest is undocumented.
type GetInstanceProfileRequest struct {
	InstanceProfileName string `xml:"InstanceProfileName"`
}

// GetInstanceProfileResponse is undocumented.
type GetInstanceProfileResponse struct {
	InstanceProfile InstanceProfile `xml:"GetInstanceProfileResult>InstanceProfile"`
}

// GetLoginProfileRequest is undocumented.
type GetLoginProfileRequest struct {
	UserName string `xml:"UserName"`
}

// GetLoginProfileResponse is undocumented.
type GetLoginProfileResponse struct {
	LoginProfile LoginProfile `xml:"GetLoginProfileResult>LoginProfile"`
}

// GetOpenIDConnectProviderRequest is undocumented.
type GetOpenIDConnectProviderRequest struct {
	OpenIDConnectProviderARN string `xml:"OpenIDConnectProviderArn"`
}

// GetOpenIDConnectProviderResponse is undocumented.
type GetOpenIDConnectProviderResponse struct {
	ClientIDList   []string  `xml:"GetOpenIDConnectProviderResult>ClientIDList>member"`
	CreateDate     time.Time `xml:"GetOpenIDConnectProviderResult>CreateDate"`
	ThumbprintList []string  `xml:"GetOpenIDConnectProviderResult>ThumbprintList>member"`
	URL            string    `xml:"GetOpenIDConnectProviderResult>Url"`
}

// GetRolePolicyRequest is undocumented.
type GetRolePolicyRequest struct {
	PolicyName string `xml:"PolicyName"`
	RoleName   string `xml:"RoleName"`
}

// GetRolePolicyResponse is undocumented.
type GetRolePolicyResponse struct {
	PolicyDocument string `xml:"GetRolePolicyResult>PolicyDocument"`
	PolicyName     string `xml:"GetRolePolicyResult>PolicyName"`
	RoleName       string `xml:"GetRolePolicyResult>RoleName"`
}

// GetRoleRequest is undocumented.
type GetRoleRequest struct {
	RoleName string `xml:"RoleName"`
}

// GetRoleResponse is undocumented.
type GetRoleResponse struct {
	Role Role `xml:"GetRoleResult>Role"`
}

// GetSAMLProviderRequest is undocumented.
type GetSAMLProviderRequest struct {
	SAMLProviderARN string `xml:"SAMLProviderArn"`
}

// GetSAMLProviderResponse is undocumented.
type GetSAMLProviderResponse struct {
	CreateDate           time.Time `xml:"GetSAMLProviderResult>CreateDate"`
	SAMLMetadataDocument string    `xml:"GetSAMLProviderResult>SAMLMetadataDocument"`
	ValidUntil           time.Time `xml:"GetSAMLProviderResult>ValidUntil"`
}

// GetServerCertificateRequest is undocumented.
type GetServerCertificateRequest struct {
	ServerCertificateName string `xml:"ServerCertificateName"`
}

// GetServerCertificateResponse is undocumented.
type GetServerCertificateResponse struct {
	ServerCertificate ServerCertificate `xml:"GetServerCertificateResult>ServerCertificate"`
}

// GetUserPolicyRequest is undocumented.
type GetUserPolicyRequest struct {
	PolicyName string `xml:"PolicyName"`
	UserName   string `xml:"UserName"`
}

// GetUserPolicyResponse is undocumented.
type GetUserPolicyResponse struct {
	PolicyDocument string `xml:"GetUserPolicyResult>PolicyDocument"`
	PolicyName     string `xml:"GetUserPolicyResult>PolicyName"`
	UserName       string `xml:"GetUserPolicyResult>UserName"`
}

// GetUserRequest is undocumented.
type GetUserRequest struct {
	UserName string `xml:"UserName"`
}

// GetUserResponse is undocumented.
type GetUserResponse struct {
	User User `xml:"GetUserResult>User"`
}

// Group is undocumented.
type Group struct {
	ARN        string    `xml:"Arn"`
	CreateDate time.Time `xml:"CreateDate"`
	GroupID    string    `xml:"GroupId"`
	GroupName  string    `xml:"GroupName"`
	Path       string    `xml:"Path"`
}

// GroupDetail is undocumented.
type GroupDetail struct {
	ARN             string         `xml:"Arn"`
	CreateDate      time.Time      `xml:"CreateDate"`
	GroupID         string         `xml:"GroupId"`
	GroupName       string         `xml:"GroupName"`
	GroupPolicyList []PolicyDetail `xml:"GroupPolicyList>member"`
	Path            string         `xml:"Path"`
}

// InstanceProfile is undocumented.
type InstanceProfile struct {
	ARN                 string    `xml:"Arn"`
	CreateDate          time.Time `xml:"CreateDate"`
	InstanceProfileID   string    `xml:"InstanceProfileId"`
	InstanceProfileName string    `xml:"InstanceProfileName"`
	Path                string    `xml:"Path"`
	Roles               []Role    `xml:"Roles>member"`
}

// ListAccessKeysRequest is undocumented.
type ListAccessKeysRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems int    `xml:"MaxItems"`
	UserName string `xml:"UserName"`
}

// ListAccessKeysResponse is undocumented.
type ListAccessKeysResponse struct {
	AccessKeyMetadata []AccessKeyMetadata `xml:"ListAccessKeysResult>AccessKeyMetadata>member"`
	IsTruncated       bool                `xml:"ListAccessKeysResult>IsTruncated"`
	Marker            string              `xml:"ListAccessKeysResult>Marker"`
}

// ListAccountAliasesRequest is undocumented.
type ListAccountAliasesRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems int    `xml:"MaxItems"`
}

// ListAccountAliasesResponse is undocumented.
type ListAccountAliasesResponse struct {
	AccountAliases []string `xml:"ListAccountAliasesResult>AccountAliases>member"`
	IsTruncated    bool     `xml:"ListAccountAliasesResult>IsTruncated"`
	Marker         string   `xml:"ListAccountAliasesResult>Marker"`
}

// ListGroupPoliciesRequest is undocumented.
type ListGroupPoliciesRequest struct {
	GroupName string `xml:"GroupName"`
	Marker    string `xml:"Marker"`
	MaxItems  int    `xml:"MaxItems"`
}

// ListGroupPoliciesResponse is undocumented.
type ListGroupPoliciesResponse struct {
	IsTruncated bool     `xml:"ListGroupPoliciesResult>IsTruncated"`
	Marker      string   `xml:"ListGroupPoliciesResult>Marker"`
	PolicyNames []string `xml:"ListGroupPoliciesResult>PolicyNames>member"`
}

// ListGroupsForUserRequest is undocumented.
type ListGroupsForUserRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems int    `xml:"MaxItems"`
	UserName string `xml:"UserName"`
}

// ListGroupsForUserResponse is undocumented.
type ListGroupsForUserResponse struct {
	Groups      []Group `xml:"ListGroupsForUserResult>Groups>member"`
	IsTruncated bool    `xml:"ListGroupsForUserResult>IsTruncated"`
	Marker      string  `xml:"ListGroupsForUserResult>Marker"`
}

// ListGroupsRequest is undocumented.
type ListGroupsRequest struct {
	Marker     string `xml:"Marker"`
	MaxItems   int    `xml:"MaxItems"`
	PathPrefix string `xml:"PathPrefix"`
}

// ListGroupsResponse is undocumented.
type ListGroupsResponse struct {
	Groups      []Group `xml:"ListGroupsResult>Groups>member"`
	IsTruncated bool    `xml:"ListGroupsResult>IsTruncated"`
	Marker      string  `xml:"ListGroupsResult>Marker"`
}

// ListInstanceProfilesForRoleRequest is undocumented.
type ListInstanceProfilesForRoleRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems int    `xml:"MaxItems"`
	RoleName string `xml:"RoleName"`
}

// ListInstanceProfilesForRoleResponse is undocumented.
type ListInstanceProfilesForRoleResponse struct {
	InstanceProfiles []InstanceProfile `xml:"ListInstanceProfilesForRoleResult>InstanceProfiles>member"`
	IsTruncated      bool              `xml:"ListInstanceProfilesForRoleResult>IsTruncated"`
	Marker           string            `xml:"ListInstanceProfilesForRoleResult>Marker"`
}

// ListInstanceProfilesRequest is undocumented.
type ListInstanceProfilesRequest struct {
	Marker     string `xml:"Marker"`
	MaxItems   int    `xml:"MaxItems"`
	PathPrefix string `xml:"PathPrefix"`
}

// ListInstanceProfilesResponse is undocumented.
type ListInstanceProfilesResponse struct {
	InstanceProfiles []InstanceProfile `xml:"ListInstanceProfilesResult>InstanceProfiles>member"`
	IsTruncated      bool              `xml:"ListInstanceProfilesResult>IsTruncated"`
	Marker           string            `xml:"ListInstanceProfilesResult>Marker"`
}

// ListMFADevicesRequest is undocumented.
type ListMFADevicesRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems int    `xml:"MaxItems"`
	UserName string `xml:"UserName"`
}

// ListMFADevicesResponse is undocumented.
type ListMFADevicesResponse struct {
	IsTruncated bool        `xml:"ListMFADevicesResult>IsTruncated"`
	MFADevices  []MFADevice `xml:"ListMFADevicesResult>MFADevices>member"`
	Marker      string      `xml:"ListMFADevicesResult>Marker"`
}

// ListOpenIDConnectProvidersRequest is undocumented.
type ListOpenIDConnectProvidersRequest struct {
}

// ListOpenIDConnectProvidersResponse is undocumented.
type ListOpenIDConnectProvidersResponse struct {
	OpenIDConnectProviderList []OpenIDConnectProviderListEntry `xml:"ListOpenIDConnectProvidersResult>OpenIDConnectProviderList>member"`
}

// ListRolePoliciesRequest is undocumented.
type ListRolePoliciesRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems int    `xml:"MaxItems"`
	RoleName string `xml:"RoleName"`
}

// ListRolePoliciesResponse is undocumented.
type ListRolePoliciesResponse struct {
	IsTruncated bool     `xml:"ListRolePoliciesResult>IsTruncated"`
	Marker      string   `xml:"ListRolePoliciesResult>Marker"`
	PolicyNames []string `xml:"ListRolePoliciesResult>PolicyNames>member"`
}

// ListRolesRequest is undocumented.
type ListRolesRequest struct {
	Marker     string `xml:"Marker"`
	MaxItems   int    `xml:"MaxItems"`
	PathPrefix string `xml:"PathPrefix"`
}

// ListRolesResponse is undocumented.
type ListRolesResponse struct {
	IsTruncated bool   `xml:"ListRolesResult>IsTruncated"`
	Marker      string `xml:"ListRolesResult>Marker"`
	Roles       []Role `xml:"ListRolesResult>Roles>member"`
}

// ListSAMLProvidersRequest is undocumented.
type ListSAMLProvidersRequest struct {
}

// ListSAMLProvidersResponse is undocumented.
type ListSAMLProvidersResponse struct {
	SAMLProviderList []SAMLProviderListEntry `xml:"ListSAMLProvidersResult>SAMLProviderList>member"`
}

// ListServerCertificatesRequest is undocumented.
type ListServerCertificatesRequest struct {
	Marker     string `xml:"Marker"`
	MaxItems   int    `xml:"MaxItems"`
	PathPrefix string `xml:"PathPrefix"`
}

// ListServerCertificatesResponse is undocumented.
type ListServerCertificatesResponse struct {
	IsTruncated                   bool                        `xml:"ListServerCertificatesResult>IsTruncated"`
	Marker                        string                      `xml:"ListServerCertificatesResult>Marker"`
	ServerCertificateMetadataList []ServerCertificateMetadata `xml:"ListServerCertificatesResult>ServerCertificateMetadataList>member"`
}

// ListSigningCertificatesRequest is undocumented.
type ListSigningCertificatesRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems int    `xml:"MaxItems"`
	UserName string `xml:"UserName"`
}

// ListSigningCertificatesResponse is undocumented.
type ListSigningCertificatesResponse struct {
	Certificates []SigningCertificate `xml:"ListSigningCertificatesResult>Certificates>member"`
	IsTruncated  bool                 `xml:"ListSigningCertificatesResult>IsTruncated"`
	Marker       string               `xml:"ListSigningCertificatesResult>Marker"`
}

// ListUserPoliciesRequest is undocumented.
type ListUserPoliciesRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems int    `xml:"MaxItems"`
	UserName string `xml:"UserName"`
}

// ListUserPoliciesResponse is undocumented.
type ListUserPoliciesResponse struct {
	IsTruncated bool     `xml:"ListUserPoliciesResult>IsTruncated"`
	Marker      string   `xml:"ListUserPoliciesResult>Marker"`
	PolicyNames []string `xml:"ListUserPoliciesResult>PolicyNames>member"`
}

// ListUsersRequest is undocumented.
type ListUsersRequest struct {
	Marker     string `xml:"Marker"`
	MaxItems   int    `xml:"MaxItems"`
	PathPrefix string `xml:"PathPrefix"`
}

// ListUsersResponse is undocumented.
type ListUsersResponse struct {
	IsTruncated bool   `xml:"ListUsersResult>IsTruncated"`
	Marker      string `xml:"ListUsersResult>Marker"`
	Users       []User `xml:"ListUsersResult>Users>member"`
}

// ListVirtualMFADevicesRequest is undocumented.
type ListVirtualMFADevicesRequest struct {
	AssignmentStatus string `xml:"AssignmentStatus"`
	Marker           string `xml:"Marker"`
	MaxItems         int    `xml:"MaxItems"`
}

// ListVirtualMFADevicesResponse is undocumented.
type ListVirtualMFADevicesResponse struct {
	IsTruncated       bool               `xml:"ListVirtualMFADevicesResult>IsTruncated"`
	Marker            string             `xml:"ListVirtualMFADevicesResult>Marker"`
	VirtualMFADevices []VirtualMFADevice `xml:"ListVirtualMFADevicesResult>VirtualMFADevices>member"`
}

// LoginProfile is undocumented.
type LoginProfile struct {
	CreateDate            time.Time `xml:"CreateDate"`
	PasswordResetRequired bool      `xml:"PasswordResetRequired"`
	UserName              string    `xml:"UserName"`
}

// MFADevice is undocumented.
type MFADevice struct {
	EnableDate   time.Time `xml:"EnableDate"`
	SerialNumber string    `xml:"SerialNumber"`
	UserName     string    `xml:"UserName"`
}

// OpenIDConnectProviderListEntry is undocumented.
type OpenIDConnectProviderListEntry struct {
	ARN string `xml:"Arn"`
}

// PasswordPolicy is undocumented.
type PasswordPolicy struct {
	AllowUsersToChangePassword bool `xml:"AllowUsersToChangePassword"`
	ExpirePasswords            bool `xml:"ExpirePasswords"`
	HardExpiry                 bool `xml:"HardExpiry"`
	MaxPasswordAge             int  `xml:"MaxPasswordAge"`
	MinimumPasswordLength      int  `xml:"MinimumPasswordLength"`
	PasswordReusePrevention    int  `xml:"PasswordReusePrevention"`
	RequireLowercaseCharacters bool `xml:"RequireLowercaseCharacters"`
	RequireNumbers             bool `xml:"RequireNumbers"`
	RequireSymbols             bool `xml:"RequireSymbols"`
	RequireUppercaseCharacters bool `xml:"RequireUppercaseCharacters"`
}

// PolicyDetail is undocumented.
type PolicyDetail struct {
	PolicyDocument string `xml:"PolicyDocument"`
	PolicyName     string `xml:"PolicyName"`
}

// PutGroupPolicyRequest is undocumented.
type PutGroupPolicyRequest struct {
	GroupName      string `xml:"GroupName"`
	PolicyDocument string `xml:"PolicyDocument"`
	PolicyName     string `xml:"PolicyName"`
}

// PutRolePolicyRequest is undocumented.
type PutRolePolicyRequest struct {
	PolicyDocument string `xml:"PolicyDocument"`
	PolicyName     string `xml:"PolicyName"`
	RoleName       string `xml:"RoleName"`
}

// PutUserPolicyRequest is undocumented.
type PutUserPolicyRequest struct {
	PolicyDocument string `xml:"PolicyDocument"`
	PolicyName     string `xml:"PolicyName"`
	UserName       string `xml:"UserName"`
}

// RemoveClientIDFromOpenIDConnectProviderRequest is undocumented.
type RemoveClientIDFromOpenIDConnectProviderRequest struct {
	ClientID                 string `xml:"ClientID"`
	OpenIDConnectProviderARN string `xml:"OpenIDConnectProviderArn"`
}

// RemoveRoleFromInstanceProfileRequest is undocumented.
type RemoveRoleFromInstanceProfileRequest struct {
	InstanceProfileName string `xml:"InstanceProfileName"`
	RoleName            string `xml:"RoleName"`
}

// RemoveUserFromGroupRequest is undocumented.
type RemoveUserFromGroupRequest struct {
	GroupName string `xml:"GroupName"`
	UserName  string `xml:"UserName"`
}

// ResyncMFADeviceRequest is undocumented.
type ResyncMFADeviceRequest struct {
	AuthenticationCode1 string `xml:"AuthenticationCode1"`
	AuthenticationCode2 string `xml:"AuthenticationCode2"`
	SerialNumber        string `xml:"SerialNumber"`
	UserName            string `xml:"UserName"`
}

// Role is undocumented.
type Role struct {
	ARN                      string    `xml:"Arn"`
	AssumeRolePolicyDocument string    `xml:"AssumeRolePolicyDocument"`
	CreateDate               time.Time `xml:"CreateDate"`
	Path                     string    `xml:"Path"`
	RoleID                   string    `xml:"RoleId"`
	RoleName                 string    `xml:"RoleName"`
}

// RoleDetail is undocumented.
type RoleDetail struct {
	ARN                      string            `xml:"Arn"`
	AssumeRolePolicyDocument string            `xml:"AssumeRolePolicyDocument"`
	CreateDate               time.Time         `xml:"CreateDate"`
	InstanceProfileList      []InstanceProfile `xml:"InstanceProfileList>member"`
	Path                     string            `xml:"Path"`
	RoleID                   string            `xml:"RoleId"`
	RoleName                 string            `xml:"RoleName"`
	RolePolicyList           []PolicyDetail    `xml:"RolePolicyList>member"`
}

// SAMLProviderListEntry is undocumented.
type SAMLProviderListEntry struct {
	ARN        string    `xml:"Arn"`
	CreateDate time.Time `xml:"CreateDate"`
	ValidUntil time.Time `xml:"ValidUntil"`
}

// ServerCertificate is undocumented.
type ServerCertificate struct {
	CertificateBody           string                    `xml:"CertificateBody"`
	CertificateChain          string                    `xml:"CertificateChain"`
	ServerCertificateMetadata ServerCertificateMetadata `xml:"ServerCertificateMetadata"`
}

// ServerCertificateMetadata is undocumented.
type ServerCertificateMetadata struct {
	ARN                   string    `xml:"Arn"`
	Expiration            time.Time `xml:"Expiration"`
	Path                  string    `xml:"Path"`
	ServerCertificateID   string    `xml:"ServerCertificateId"`
	ServerCertificateName string    `xml:"ServerCertificateName"`
	UploadDate            time.Time `xml:"UploadDate"`
}

// SigningCertificate is undocumented.
type SigningCertificate struct {
	CertificateBody string    `xml:"CertificateBody"`
	CertificateID   string    `xml:"CertificateId"`
	Status          string    `xml:"Status"`
	UploadDate      time.Time `xml:"UploadDate"`
	UserName        string    `xml:"UserName"`
}

// UpdateAccessKeyRequest is undocumented.
type UpdateAccessKeyRequest struct {
	AccessKeyID string `xml:"AccessKeyId"`
	Status      string `xml:"Status"`
	UserName    string `xml:"UserName"`
}

// UpdateAccountPasswordPolicyRequest is undocumented.
type UpdateAccountPasswordPolicyRequest struct {
	AllowUsersToChangePassword bool `xml:"AllowUsersToChangePassword"`
	HardExpiry                 bool `xml:"HardExpiry"`
	MaxPasswordAge             int  `xml:"MaxPasswordAge"`
	MinimumPasswordLength      int  `xml:"MinimumPasswordLength"`
	PasswordReusePrevention    int  `xml:"PasswordReusePrevention"`
	RequireLowercaseCharacters bool `xml:"RequireLowercaseCharacters"`
	RequireNumbers             bool `xml:"RequireNumbers"`
	RequireSymbols             bool `xml:"RequireSymbols"`
	RequireUppercaseCharacters bool `xml:"RequireUppercaseCharacters"`
}

// UpdateAssumeRolePolicyRequest is undocumented.
type UpdateAssumeRolePolicyRequest struct {
	PolicyDocument string `xml:"PolicyDocument"`
	RoleName       string `xml:"RoleName"`
}

// UpdateGroupRequest is undocumented.
type UpdateGroupRequest struct {
	GroupName    string `xml:"GroupName"`
	NewGroupName string `xml:"NewGroupName"`
	NewPath      string `xml:"NewPath"`
}

// UpdateLoginProfileRequest is undocumented.
type UpdateLoginProfileRequest struct {
	Password              string `xml:"Password"`
	PasswordResetRequired bool   `xml:"PasswordResetRequired"`
	UserName              string `xml:"UserName"`
}

// UpdateOpenIDConnectProviderThumbprintRequest is undocumented.
type UpdateOpenIDConnectProviderThumbprintRequest struct {
	OpenIDConnectProviderARN string   `xml:"OpenIDConnectProviderArn"`
	ThumbprintList           []string `xml:"ThumbprintList>member"`
}

// UpdateSAMLProviderRequest is undocumented.
type UpdateSAMLProviderRequest struct {
	SAMLMetadataDocument string `xml:"SAMLMetadataDocument"`
	SAMLProviderARN      string `xml:"SAMLProviderArn"`
}

// UpdateSAMLProviderResponse is undocumented.
type UpdateSAMLProviderResponse struct {
	SAMLProviderARN string `xml:"UpdateSAMLProviderResult>SAMLProviderArn"`
}

// UpdateServerCertificateRequest is undocumented.
type UpdateServerCertificateRequest struct {
	NewPath                  string `xml:"NewPath"`
	NewServerCertificateName string `xml:"NewServerCertificateName"`
	ServerCertificateName    string `xml:"ServerCertificateName"`
}

// UpdateSigningCertificateRequest is undocumented.
type UpdateSigningCertificateRequest struct {
	CertificateID string `xml:"CertificateId"`
	Status        string `xml:"Status"`
	UserName      string `xml:"UserName"`
}

// UpdateUserRequest is undocumented.
type UpdateUserRequest struct {
	NewPath     string `xml:"NewPath"`
	NewUserName string `xml:"NewUserName"`
	UserName    string `xml:"UserName"`
}

// UploadServerCertificateRequest is undocumented.
type UploadServerCertificateRequest struct {
	CertificateBody       string `xml:"CertificateBody"`
	CertificateChain      string `xml:"CertificateChain"`
	Path                  string `xml:"Path"`
	PrivateKey            string `xml:"PrivateKey"`
	ServerCertificateName string `xml:"ServerCertificateName"`
}

// UploadServerCertificateResponse is undocumented.
type UploadServerCertificateResponse struct {
	ServerCertificateMetadata ServerCertificateMetadata `xml:"UploadServerCertificateResult>ServerCertificateMetadata"`
}

// UploadSigningCertificateRequest is undocumented.
type UploadSigningCertificateRequest struct {
	CertificateBody string `xml:"CertificateBody"`
	UserName        string `xml:"UserName"`
}

// UploadSigningCertificateResponse is undocumented.
type UploadSigningCertificateResponse struct {
	Certificate SigningCertificate `xml:"UploadSigningCertificateResult>Certificate"`
}

// User is undocumented.
type User struct {
	ARN              string    `xml:"Arn"`
	CreateDate       time.Time `xml:"CreateDate"`
	PasswordLastUsed time.Time `xml:"PasswordLastUsed"`
	Path             string    `xml:"Path"`
	UserID           string    `xml:"UserId"`
	UserName         string    `xml:"UserName"`
}

// UserDetail is undocumented.
type UserDetail struct {
	ARN            string         `xml:"Arn"`
	CreateDate     time.Time      `xml:"CreateDate"`
	GroupList      []string       `xml:"GroupList>member"`
	Path           string         `xml:"Path"`
	UserID         string         `xml:"UserId"`
	UserName       string         `xml:"UserName"`
	UserPolicyList []PolicyDetail `xml:"UserPolicyList>member"`
}

// VirtualMFADevice is undocumented.
type VirtualMFADevice struct {
	Base32StringSeed []byte    `xml:"Base32StringSeed"`
	EnableDate       time.Time `xml:"EnableDate"`
	QRCodePNG        []byte    `xml:"QRCodePNG"`
	SerialNumber     string    `xml:"SerialNumber"`
	User             User      `xml:"User"`
}

// CreateAccessKeyResult is a wrapper for CreateAccessKeyResponse.
type CreateAccessKeyResult struct {
	XMLName xml.Name `xml:"CreateAccessKeyResponse"`

	AccessKey AccessKey `xml:"CreateAccessKeyResult>AccessKey"`
}

// CreateGroupResult is a wrapper for CreateGroupResponse.
type CreateGroupResult struct {
	XMLName xml.Name `xml:"CreateGroupResponse"`

	Group Group `xml:"CreateGroupResult>Group"`
}

// CreateInstanceProfileResult is a wrapper for CreateInstanceProfileResponse.
type CreateInstanceProfileResult struct {
	XMLName xml.Name `xml:"CreateInstanceProfileResponse"`

	InstanceProfile InstanceProfile `xml:"CreateInstanceProfileResult>InstanceProfile"`
}

// CreateLoginProfileResult is a wrapper for CreateLoginProfileResponse.
type CreateLoginProfileResult struct {
	XMLName xml.Name `xml:"CreateLoginProfileResponse"`

	LoginProfile LoginProfile `xml:"CreateLoginProfileResult>LoginProfile"`
}

// CreateOpenIDConnectProviderResult is a wrapper for CreateOpenIDConnectProviderResponse.
type CreateOpenIDConnectProviderResult struct {
	XMLName xml.Name `xml:"CreateOpenIDConnectProviderResponse"`

	OpenIDConnectProviderARN string `xml:"CreateOpenIDConnectProviderResult>OpenIDConnectProviderArn"`
}

// CreateRoleResult is a wrapper for CreateRoleResponse.
type CreateRoleResult struct {
	XMLName xml.Name `xml:"CreateRoleResponse"`

	Role Role `xml:"CreateRoleResult>Role"`
}

// CreateSAMLProviderResult is a wrapper for CreateSAMLProviderResponse.
type CreateSAMLProviderResult struct {
	XMLName xml.Name `xml:"CreateSAMLProviderResponse"`

	SAMLProviderARN string `xml:"CreateSAMLProviderResult>SAMLProviderArn"`
}

// CreateUserResult is a wrapper for CreateUserResponse.
type CreateUserResult struct {
	XMLName xml.Name `xml:"CreateUserResponse"`

	User User `xml:"CreateUserResult>User"`
}

// CreateVirtualMFADeviceResult is a wrapper for CreateVirtualMFADeviceResponse.
type CreateVirtualMFADeviceResult struct {
	XMLName xml.Name `xml:"CreateVirtualMFADeviceResponse"`

	VirtualMFADevice VirtualMFADevice `xml:"CreateVirtualMFADeviceResult>VirtualMFADevice"`
}

// GenerateCredentialReportResult is a wrapper for GenerateCredentialReportResponse.
type GenerateCredentialReportResult struct {
	XMLName xml.Name `xml:"GenerateCredentialReportResponse"`

	Description string `xml:"GenerateCredentialReportResult>Description"`
	State       string `xml:"GenerateCredentialReportResult>State"`
}

// GetAccountAuthorizationDetailsResult is a wrapper for GetAccountAuthorizationDetailsResponse.
type GetAccountAuthorizationDetailsResult struct {
	XMLName xml.Name `xml:"GetAccountAuthorizationDetailsResponse"`

	GroupDetailList []GroupDetail `xml:"GetAccountAuthorizationDetailsResult>GroupDetailList>member"`
	IsTruncated     bool          `xml:"GetAccountAuthorizationDetailsResult>IsTruncated"`
	Marker          string        `xml:"GetAccountAuthorizationDetailsResult>Marker"`
	RoleDetailList  []RoleDetail  `xml:"GetAccountAuthorizationDetailsResult>RoleDetailList>member"`
	UserDetailList  []UserDetail  `xml:"GetAccountAuthorizationDetailsResult>UserDetailList>member"`
}

// GetAccountPasswordPolicyResult is a wrapper for GetAccountPasswordPolicyResponse.
type GetAccountPasswordPolicyResult struct {
	XMLName xml.Name `xml:"GetAccountPasswordPolicyResponse"`

	PasswordPolicy PasswordPolicy `xml:"GetAccountPasswordPolicyResult>PasswordPolicy"`
}

// GetAccountSummaryResult is a wrapper for GetAccountSummaryResponse.
type GetAccountSummaryResult struct {
	XMLName xml.Name `xml:"GetAccountSummaryResponse"`

	SummaryMap map[string]int `xml:"GetAccountSummaryResult>SummaryMap"`
}

// GetCredentialReportResult is a wrapper for GetCredentialReportResponse.
type GetCredentialReportResult struct {
	XMLName xml.Name `xml:"GetCredentialReportResponse"`

	Content       []byte    `xml:"GetCredentialReportResult>Content"`
	GeneratedTime time.Time `xml:"GetCredentialReportResult>GeneratedTime"`
	ReportFormat  string    `xml:"GetCredentialReportResult>ReportFormat"`
}

// GetGroupPolicyResult is a wrapper for GetGroupPolicyResponse.
type GetGroupPolicyResult struct {
	XMLName xml.Name `xml:"GetGroupPolicyResponse"`

	GroupName      string `xml:"GetGroupPolicyResult>GroupName"`
	PolicyDocument string `xml:"GetGroupPolicyResult>PolicyDocument"`
	PolicyName     string `xml:"GetGroupPolicyResult>PolicyName"`
}

// GetGroupResult is a wrapper for GetGroupResponse.
type GetGroupResult struct {
	XMLName xml.Name `xml:"GetGroupResponse"`

	Group       Group  `xml:"GetGroupResult>Group"`
	IsTruncated bool   `xml:"GetGroupResult>IsTruncated"`
	Marker      string `xml:"GetGroupResult>Marker"`
	Users       []User `xml:"GetGroupResult>Users>member"`
}

// GetInstanceProfileResult is a wrapper for GetInstanceProfileResponse.
type GetInstanceProfileResult struct {
	XMLName xml.Name `xml:"GetInstanceProfileResponse"`

	InstanceProfile InstanceProfile `xml:"GetInstanceProfileResult>InstanceProfile"`
}

// GetLoginProfileResult is a wrapper for GetLoginProfileResponse.
type GetLoginProfileResult struct {
	XMLName xml.Name `xml:"GetLoginProfileResponse"`

	LoginProfile LoginProfile `xml:"GetLoginProfileResult>LoginProfile"`
}

// GetOpenIDConnectProviderResult is a wrapper for GetOpenIDConnectProviderResponse.
type GetOpenIDConnectProviderResult struct {
	XMLName xml.Name `xml:"GetOpenIDConnectProviderResponse"`

	ClientIDList   []string  `xml:"GetOpenIDConnectProviderResult>ClientIDList>member"`
	CreateDate     time.Time `xml:"GetOpenIDConnectProviderResult>CreateDate"`
	ThumbprintList []string  `xml:"GetOpenIDConnectProviderResult>ThumbprintList>member"`
	URL            string    `xml:"GetOpenIDConnectProviderResult>Url"`
}

// GetRolePolicyResult is a wrapper for GetRolePolicyResponse.
type GetRolePolicyResult struct {
	XMLName xml.Name `xml:"GetRolePolicyResponse"`

	PolicyDocument string `xml:"GetRolePolicyResult>PolicyDocument"`
	PolicyName     string `xml:"GetRolePolicyResult>PolicyName"`
	RoleName       string `xml:"GetRolePolicyResult>RoleName"`
}

// GetRoleResult is a wrapper for GetRoleResponse.
type GetRoleResult struct {
	XMLName xml.Name `xml:"GetRoleResponse"`

	Role Role `xml:"GetRoleResult>Role"`
}

// GetSAMLProviderResult is a wrapper for GetSAMLProviderResponse.
type GetSAMLProviderResult struct {
	XMLName xml.Name `xml:"GetSAMLProviderResponse"`

	CreateDate           time.Time `xml:"GetSAMLProviderResult>CreateDate"`
	SAMLMetadataDocument string    `xml:"GetSAMLProviderResult>SAMLMetadataDocument"`
	ValidUntil           time.Time `xml:"GetSAMLProviderResult>ValidUntil"`
}

// GetServerCertificateResult is a wrapper for GetServerCertificateResponse.
type GetServerCertificateResult struct {
	XMLName xml.Name `xml:"GetServerCertificateResponse"`

	ServerCertificate ServerCertificate `xml:"GetServerCertificateResult>ServerCertificate"`
}

// GetUserPolicyResult is a wrapper for GetUserPolicyResponse.
type GetUserPolicyResult struct {
	XMLName xml.Name `xml:"GetUserPolicyResponse"`

	PolicyDocument string `xml:"GetUserPolicyResult>PolicyDocument"`
	PolicyName     string `xml:"GetUserPolicyResult>PolicyName"`
	UserName       string `xml:"GetUserPolicyResult>UserName"`
}

// GetUserResult is a wrapper for GetUserResponse.
type GetUserResult struct {
	XMLName xml.Name `xml:"GetUserResponse"`

	User User `xml:"GetUserResult>User"`
}

// ListAccessKeysResult is a wrapper for ListAccessKeysResponse.
type ListAccessKeysResult struct {
	XMLName xml.Name `xml:"ListAccessKeysResponse"`

	AccessKeyMetadata []AccessKeyMetadata `xml:"ListAccessKeysResult>AccessKeyMetadata>member"`
	IsTruncated       bool                `xml:"ListAccessKeysResult>IsTruncated"`
	Marker            string              `xml:"ListAccessKeysResult>Marker"`
}

// ListAccountAliasesResult is a wrapper for ListAccountAliasesResponse.
type ListAccountAliasesResult struct {
	XMLName xml.Name `xml:"ListAccountAliasesResponse"`

	AccountAliases []string `xml:"ListAccountAliasesResult>AccountAliases>member"`
	IsTruncated    bool     `xml:"ListAccountAliasesResult>IsTruncated"`
	Marker         string   `xml:"ListAccountAliasesResult>Marker"`
}

// ListGroupPoliciesResult is a wrapper for ListGroupPoliciesResponse.
type ListGroupPoliciesResult struct {
	XMLName xml.Name `xml:"ListGroupPoliciesResponse"`

	IsTruncated bool     `xml:"ListGroupPoliciesResult>IsTruncated"`
	Marker      string   `xml:"ListGroupPoliciesResult>Marker"`
	PolicyNames []string `xml:"ListGroupPoliciesResult>PolicyNames>member"`
}

// ListGroupsForUserResult is a wrapper for ListGroupsForUserResponse.
type ListGroupsForUserResult struct {
	XMLName xml.Name `xml:"ListGroupsForUserResponse"`

	Groups      []Group `xml:"ListGroupsForUserResult>Groups>member"`
	IsTruncated bool    `xml:"ListGroupsForUserResult>IsTruncated"`
	Marker      string  `xml:"ListGroupsForUserResult>Marker"`
}

// ListGroupsResult is a wrapper for ListGroupsResponse.
type ListGroupsResult struct {
	XMLName xml.Name `xml:"ListGroupsResponse"`

	Groups      []Group `xml:"ListGroupsResult>Groups>member"`
	IsTruncated bool    `xml:"ListGroupsResult>IsTruncated"`
	Marker      string  `xml:"ListGroupsResult>Marker"`
}

// ListInstanceProfilesForRoleResult is a wrapper for ListInstanceProfilesForRoleResponse.
type ListInstanceProfilesForRoleResult struct {
	XMLName xml.Name `xml:"ListInstanceProfilesForRoleResponse"`

	InstanceProfiles []InstanceProfile `xml:"ListInstanceProfilesForRoleResult>InstanceProfiles>member"`
	IsTruncated      bool              `xml:"ListInstanceProfilesForRoleResult>IsTruncated"`
	Marker           string            `xml:"ListInstanceProfilesForRoleResult>Marker"`
}

// ListInstanceProfilesResult is a wrapper for ListInstanceProfilesResponse.
type ListInstanceProfilesResult struct {
	XMLName xml.Name `xml:"ListInstanceProfilesResponse"`

	InstanceProfiles []InstanceProfile `xml:"ListInstanceProfilesResult>InstanceProfiles>member"`
	IsTruncated      bool              `xml:"ListInstanceProfilesResult>IsTruncated"`
	Marker           string            `xml:"ListInstanceProfilesResult>Marker"`
}

// ListMFADevicesResult is a wrapper for ListMFADevicesResponse.
type ListMFADevicesResult struct {
	XMLName xml.Name `xml:"ListMFADevicesResponse"`

	IsTruncated bool        `xml:"ListMFADevicesResult>IsTruncated"`
	MFADevices  []MFADevice `xml:"ListMFADevicesResult>MFADevices>member"`
	Marker      string      `xml:"ListMFADevicesResult>Marker"`
}

// ListOpenIDConnectProvidersResult is a wrapper for ListOpenIDConnectProvidersResponse.
type ListOpenIDConnectProvidersResult struct {
	XMLName xml.Name `xml:"ListOpenIDConnectProvidersResponse"`

	OpenIDConnectProviderList []OpenIDConnectProviderListEntry `xml:"ListOpenIDConnectProvidersResult>OpenIDConnectProviderList>member"`
}

// ListRolePoliciesResult is a wrapper for ListRolePoliciesResponse.
type ListRolePoliciesResult struct {
	XMLName xml.Name `xml:"ListRolePoliciesResponse"`

	IsTruncated bool     `xml:"ListRolePoliciesResult>IsTruncated"`
	Marker      string   `xml:"ListRolePoliciesResult>Marker"`
	PolicyNames []string `xml:"ListRolePoliciesResult>PolicyNames>member"`
}

// ListRolesResult is a wrapper for ListRolesResponse.
type ListRolesResult struct {
	XMLName xml.Name `xml:"ListRolesResponse"`

	IsTruncated bool   `xml:"ListRolesResult>IsTruncated"`
	Marker      string `xml:"ListRolesResult>Marker"`
	Roles       []Role `xml:"ListRolesResult>Roles>member"`
}

// ListSAMLProvidersResult is a wrapper for ListSAMLProvidersResponse.
type ListSAMLProvidersResult struct {
	XMLName xml.Name `xml:"ListSAMLProvidersResponse"`

	SAMLProviderList []SAMLProviderListEntry `xml:"ListSAMLProvidersResult>SAMLProviderList>member"`
}

// ListServerCertificatesResult is a wrapper for ListServerCertificatesResponse.
type ListServerCertificatesResult struct {
	XMLName xml.Name `xml:"ListServerCertificatesResponse"`

	IsTruncated                   bool                        `xml:"ListServerCertificatesResult>IsTruncated"`
	Marker                        string                      `xml:"ListServerCertificatesResult>Marker"`
	ServerCertificateMetadataList []ServerCertificateMetadata `xml:"ListServerCertificatesResult>ServerCertificateMetadataList>member"`
}

// ListSigningCertificatesResult is a wrapper for ListSigningCertificatesResponse.
type ListSigningCertificatesResult struct {
	XMLName xml.Name `xml:"ListSigningCertificatesResponse"`

	Certificates []SigningCertificate `xml:"ListSigningCertificatesResult>Certificates>member"`
	IsTruncated  bool                 `xml:"ListSigningCertificatesResult>IsTruncated"`
	Marker       string               `xml:"ListSigningCertificatesResult>Marker"`
}

// ListUserPoliciesResult is a wrapper for ListUserPoliciesResponse.
type ListUserPoliciesResult struct {
	XMLName xml.Name `xml:"ListUserPoliciesResponse"`

	IsTruncated bool     `xml:"ListUserPoliciesResult>IsTruncated"`
	Marker      string   `xml:"ListUserPoliciesResult>Marker"`
	PolicyNames []string `xml:"ListUserPoliciesResult>PolicyNames>member"`
}

// ListUsersResult is a wrapper for ListUsersResponse.
type ListUsersResult struct {
	XMLName xml.Name `xml:"ListUsersResponse"`

	IsTruncated bool   `xml:"ListUsersResult>IsTruncated"`
	Marker      string `xml:"ListUsersResult>Marker"`
	Users       []User `xml:"ListUsersResult>Users>member"`
}

// ListVirtualMFADevicesResult is a wrapper for ListVirtualMFADevicesResponse.
type ListVirtualMFADevicesResult struct {
	XMLName xml.Name `xml:"ListVirtualMFADevicesResponse"`

	IsTruncated       bool               `xml:"ListVirtualMFADevicesResult>IsTruncated"`
	Marker            string             `xml:"ListVirtualMFADevicesResult>Marker"`
	VirtualMFADevices []VirtualMFADevice `xml:"ListVirtualMFADevicesResult>VirtualMFADevices>member"`
}

// UpdateSAMLProviderResult is a wrapper for UpdateSAMLProviderResponse.
type UpdateSAMLProviderResult struct {
	XMLName xml.Name `xml:"UpdateSAMLProviderResponse"`

	SAMLProviderARN string `xml:"UpdateSAMLProviderResult>SAMLProviderArn"`
}

// UploadServerCertificateResult is a wrapper for UploadServerCertificateResponse.
type UploadServerCertificateResult struct {
	XMLName xml.Name `xml:"UploadServerCertificateResponse"`

	ServerCertificateMetadata ServerCertificateMetadata `xml:"UploadServerCertificateResult>ServerCertificateMetadata"`
}

// UploadSigningCertificateResult is a wrapper for UploadSigningCertificateResponse.
type UploadSigningCertificateResult struct {
	XMLName xml.Name `xml:"UploadSigningCertificateResponse"`

	Certificate SigningCertificate `xml:"UploadSigningCertificateResult>Certificate"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
