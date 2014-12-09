// Package route53domains provides a client for Amazon Route 53 Domains.
package route53domains

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// Route53Domains is a client for Amazon Route 53 Domains.
type Route53Domains struct {
	client *aws.JSONClient
}

// New returns a new Route53Domains client.
func New(key, secret, region string, client *http.Client) *Route53Domains {
	if client == nil {
		client = http.DefaultClient
	}

	return &Route53Domains{
		client: &aws.JSONClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "route53domains",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:       client,
			Endpoint:     endpoints.Lookup("route53domains", region),
			JSONVersion:  "1.1",
			TargetPrefix: "Route53Domains_v20140515",
		},
	}
}

// CheckDomainAvailability this operation checks the availability of one
// domain name. You can access this API without authenticating. Note that
// if the availability status of a domain is pending, you must submit
// another request to determine the availability of the domain name.
func (c *Route53Domains) CheckDomainAvailability(req CheckDomainAvailabilityRequest) (resp *CheckDomainAvailabilityResponse, err error) {
	resp = &CheckDomainAvailabilityResponse{}
	err = c.client.Do("CheckDomainAvailability", "POST", "/", req, resp)
	return
}

// DisableDomainAutoRenew this operation disables automatic renewal of
// domain registration for the specified domain.
func (c *Route53Domains) DisableDomainAutoRenew(req DisableDomainAutoRenewRequest) (resp *DisableDomainAutoRenewResponse, err error) {
	resp = &DisableDomainAutoRenewResponse{}
	err = c.client.Do("DisableDomainAutoRenew", "POST", "/", req, resp)
	return
}

// DisableDomainTransferLock this operation removes the transfer lock on
// the domain (specifically the clientTransferProhibited status) to allow
// domain transfers. We recommend you refrain from performing this action
// unless you intend to transfer the domain to a different registrar.
// Successful submission returns an operation ID that you can use to track
// the progress and completion of the action. If the request is not
// completed successfully, the domain registrant will be notified by email.
func (c *Route53Domains) DisableDomainTransferLock(req DisableDomainTransferLockRequest) (resp *DisableDomainTransferLockResponse, err error) {
	resp = &DisableDomainTransferLockResponse{}
	err = c.client.Do("DisableDomainTransferLock", "POST", "/", req, resp)
	return
}

// EnableDomainAutoRenew this operation configures Amazon Route 53 to
// automatically renew the specified domain before the domain registration
// expires. The cost of renewing your domain registration is billed to your
// AWS account. The period during which you can renew a domain name varies
// by For a list of TLDs and their renewal policies, see "Renewal,
// restoration, and deletion times"
// (http://wiki.gandi.net/en/domains/renew#renewal_restoration_and_deletion_times)
// on the website for our registrar partner, Gandi. Route 53 requires that
// you renew before the end of the renewal period that is listed on the
// Gandi website so we can complete processing before the deadline.
func (c *Route53Domains) EnableDomainAutoRenew(req EnableDomainAutoRenewRequest) (resp *EnableDomainAutoRenewResponse, err error) {
	resp = &EnableDomainAutoRenewResponse{}
	err = c.client.Do("EnableDomainAutoRenew", "POST", "/", req, resp)
	return
}

// EnableDomainTransferLock this operation sets the transfer lock on the
// domain (specifically the clientTransferProhibited status) to prevent
// domain transfers. Successful submission returns an operation ID that you
// can use to track the progress and completion of the action. If the
// request is not completed successfully, the domain registrant will be
// notified by email.
func (c *Route53Domains) EnableDomainTransferLock(req EnableDomainTransferLockRequest) (resp *EnableDomainTransferLockResponse, err error) {
	resp = &EnableDomainTransferLockResponse{}
	err = c.client.Do("EnableDomainTransferLock", "POST", "/", req, resp)
	return
}

// GetDomainDetail this operation returns detailed information about the
// domain. The domain's contact information is also returned as part of the
// output.
func (c *Route53Domains) GetDomainDetail(req GetDomainDetailRequest) (resp *GetDomainDetailResponse, err error) {
	resp = &GetDomainDetailResponse{}
	err = c.client.Do("GetDomainDetail", "POST", "/", req, resp)
	return
}

// GetOperationDetail this operation returns the current status of an
// operation that is not completed.
func (c *Route53Domains) GetOperationDetail(req GetOperationDetailRequest) (resp *GetOperationDetailResponse, err error) {
	resp = &GetOperationDetailResponse{}
	err = c.client.Do("GetOperationDetail", "POST", "/", req, resp)
	return
}

// ListDomains this operation returns all the domain names registered with
// Amazon Route 53 for the current AWS account.
func (c *Route53Domains) ListDomains(req ListDomainsRequest) (resp *ListDomainsResponse, err error) {
	resp = &ListDomainsResponse{}
	err = c.client.Do("ListDomains", "POST", "/", req, resp)
	return
}

// ListOperations this operation returns the operation IDs of operations
// that are not yet complete.
func (c *Route53Domains) ListOperations(req ListOperationsRequest) (resp *ListOperationsResponse, err error) {
	resp = &ListOperationsResponse{}
	err = c.client.Do("ListOperations", "POST", "/", req, resp)
	return
}

// RegisterDomain this operation registers a domain. Domains are registered
// by the AWS registrar partner, Gandi. For some top-level domains (TLDs),
// this operation requires extra parameters. When you register a domain,
// Amazon Route 53 does the following: Creates a Amazon Route 53 hosted
// zone that has the same name as the domain. Amazon Route 53 assigns four
// name servers to your hosted zone and automatically updates your domain
// registration with the names of these name servers. Enables autorenew, so
// your domain registration will renew automatically each year. We'll
// notify you in advance of the renewal date so you can choose whether to
// renew the registration. Optionally enables privacy protection, so
// queries return contact information for our registrar partner, Gandi,
// instead of the information you entered for registrant, admin, and tech
// contacts. If registration is successful, returns an operation ID that
// you can use to track the progress and completion of the action. If the
// request is not completed successfully, the domain registrant is notified
// by email. Charges your AWS account an amount based on the top-level
// domain. For more information, see Amazon Route 53 Pricing
func (c *Route53Domains) RegisterDomain(req RegisterDomainRequest) (resp *RegisterDomainResponse, err error) {
	resp = &RegisterDomainResponse{}
	err = c.client.Do("RegisterDomain", "POST", "/", req, resp)
	return
}

// RetrieveDomainAuthCode this operation returns the AuthCode for the
// domain. To transfer a domain to another registrar, you provide this
// value to the new registrar.
func (c *Route53Domains) RetrieveDomainAuthCode(req RetrieveDomainAuthCodeRequest) (resp *RetrieveDomainAuthCodeResponse, err error) {
	resp = &RetrieveDomainAuthCodeResponse{}
	err = c.client.Do("RetrieveDomainAuthCode", "POST", "/", req, resp)
	return
}

// TransferDomain this operation transfers a domain from another registrar
// to Amazon Route 53. Domains are registered by the AWS registrar, Gandi
// upon transfer. To transfer a domain, you need to meet all the domain
// transfer criteria, including the following: You must supply nameservers
// to transfer a domain. You must disable the domain transfer lock (if any)
// before transferring the domain. A minimum of 60 days must have elapsed
// since the domain's registration or last transfer. We recommend you use
// the Amazon Route 53 as the DNS service for your domain. You can create a
// hosted zone in Amazon Route 53 for your current domain before
// transferring your domain. Note that upon transfer, the domain duration
// is extended for a year if not otherwise specified. Autorenew is enabled
// by default. If the transfer is successful, this method returns an
// operation ID that you can use to track the progress and completion of
// the action. If the request is not completed successfully, the domain
// registrant will be notified by email. Transferring domains charges your
// AWS account an amount based on the top-level domain. For more
// information, see Amazon Route 53 Pricing .
func (c *Route53Domains) TransferDomain(req TransferDomainRequest) (resp *TransferDomainResponse, err error) {
	resp = &TransferDomainResponse{}
	err = c.client.Do("TransferDomain", "POST", "/", req, resp)
	return
}

// UpdateDomainContact this operation updates the contact information for a
// particular domain. Information for at least one contact (registrant,
// administrator, or technical) must be supplied for update. If the update
// is successful, this method returns an operation ID that you can use to
// track the progress and completion of the action. If the request is not
// completed successfully, the domain registrant will be notified by email.
func (c *Route53Domains) UpdateDomainContact(req UpdateDomainContactRequest) (resp *UpdateDomainContactResponse, err error) {
	resp = &UpdateDomainContactResponse{}
	err = c.client.Do("UpdateDomainContact", "POST", "/", req, resp)
	return
}

// UpdateDomainContactPrivacy this operation updates the specified domain
// contact's privacy setting. When the privacy option is enabled, personal
// information such as postal or email address is hidden from the results
// of a public query. The privacy services are provided by the AWS
// registrar, Gandi. For more information, see the Gandi privacy features
// This operation only affects the privacy of the specified contact type
// (registrant, administrator, or tech). Successful acceptance returns an
// operation ID that you can use with GetOperationDetail to track the
// progress and completion of the action. If the request is not completed
// successfully, the domain registrant will be notified by email.
func (c *Route53Domains) UpdateDomainContactPrivacy(req UpdateDomainContactPrivacyRequest) (resp *UpdateDomainContactPrivacyResponse, err error) {
	resp = &UpdateDomainContactPrivacyResponse{}
	err = c.client.Do("UpdateDomainContactPrivacy", "POST", "/", req, resp)
	return
}

// UpdateDomainNameservers this operation replaces the current set of name
// servers for the domain with the specified set of name servers. If you
// use Amazon Route 53 as your DNS service, specify the four name servers
// in the delegation set for the hosted zone for the domain. If successful,
// this operation returns an operation ID that you can use to track the
// progress and completion of the action. If the request is not completed
// successfully, the domain registrant will be notified by email.
func (c *Route53Domains) UpdateDomainNameservers(req UpdateDomainNameserversRequest) (resp *UpdateDomainNameserversResponse, err error) {
	resp = &UpdateDomainNameserversResponse{}
	err = c.client.Do("UpdateDomainNameservers", "POST", "/", req, resp)
	return
}

// CheckDomainAvailabilityRequest is undocumented.
type CheckDomainAvailabilityRequest struct {
	DomainName  string `json:"DomainName"`
	IdnLangCode string `json:"IdnLangCode,omitempty"`
}

// CheckDomainAvailabilityResponse is undocumented.
type CheckDomainAvailabilityResponse struct {
	Availability string `json:"Availability"`
}

// ContactDetail is undocumented.
type ContactDetail struct {
	AddressLine1     string       `json:"AddressLine1,omitempty"`
	AddressLine2     string       `json:"AddressLine2,omitempty"`
	City             string       `json:"City,omitempty"`
	ContactType      string       `json:"ContactType,omitempty"`
	CountryCode      string       `json:"CountryCode,omitempty"`
	Email            string       `json:"Email,omitempty"`
	ExtraParams      []ExtraParam `json:"ExtraParams,omitempty"`
	Fax              string       `json:"Fax,omitempty"`
	FirstName        string       `json:"FirstName,omitempty"`
	LastName         string       `json:"LastName,omitempty"`
	OrganizationName string       `json:"OrganizationName,omitempty"`
	PhoneNumber      string       `json:"PhoneNumber,omitempty"`
	State            string       `json:"State,omitempty"`
	ZipCode          string       `json:"ZipCode,omitempty"`
}

// DisableDomainAutoRenewRequest is undocumented.
type DisableDomainAutoRenewRequest struct {
	DomainName string `json:"DomainName"`
}

// DisableDomainAutoRenewResponse is undocumented.
type DisableDomainAutoRenewResponse struct {
}

// DisableDomainTransferLockRequest is undocumented.
type DisableDomainTransferLockRequest struct {
	DomainName string `json:"DomainName"`
}

// DisableDomainTransferLockResponse is undocumented.
type DisableDomainTransferLockResponse struct {
	OperationID string `json:"OperationId"`
}

// DomainSummary is undocumented.
type DomainSummary struct {
	AutoRenew    bool      `json:"AutoRenew,omitempty"`
	DomainName   string    `json:"DomainName"`
	Expiry       time.Time `json:"Expiry,omitempty"`
	TransferLock bool      `json:"TransferLock,omitempty"`
}

// EnableDomainAutoRenewRequest is undocumented.
type EnableDomainAutoRenewRequest struct {
	DomainName string `json:"DomainName"`
}

// EnableDomainAutoRenewResponse is undocumented.
type EnableDomainAutoRenewResponse struct {
}

// EnableDomainTransferLockRequest is undocumented.
type EnableDomainTransferLockRequest struct {
	DomainName string `json:"DomainName"`
}

// EnableDomainTransferLockResponse is undocumented.
type EnableDomainTransferLockResponse struct {
	OperationID string `json:"OperationId"`
}

// ExtraParam is undocumented.
type ExtraParam struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// GetDomainDetailRequest is undocumented.
type GetDomainDetailRequest struct {
	DomainName string `json:"DomainName"`
}

// GetDomainDetailResponse is undocumented.
type GetDomainDetailResponse struct {
	AbuseContactEmail string        `json:"AbuseContactEmail,omitempty"`
	AbuseContactPhone string        `json:"AbuseContactPhone,omitempty"`
	AdminContact      ContactDetail `json:"AdminContact"`
	AdminPrivacy      bool          `json:"AdminPrivacy,omitempty"`
	AutoRenew         bool          `json:"AutoRenew,omitempty"`
	CreationDate      time.Time     `json:"CreationDate,omitempty"`
	DNSSec            string        `json:"DnsSec,omitempty"`
	DomainName        string        `json:"DomainName"`
	ExpirationDate    time.Time     `json:"ExpirationDate,omitempty"`
	Nameservers       []Nameserver  `json:"Nameservers"`
	RegistrantContact ContactDetail `json:"RegistrantContact"`
	RegistrantPrivacy bool          `json:"RegistrantPrivacy,omitempty"`
	RegistrarName     string        `json:"RegistrarName,omitempty"`
	RegistrarURL      string        `json:"RegistrarUrl,omitempty"`
	RegistryDomainID  string        `json:"RegistryDomainId,omitempty"`
	Reseller          string        `json:"Reseller,omitempty"`
	StatusList        []string      `json:"StatusList,omitempty"`
	TechContact       ContactDetail `json:"TechContact"`
	TechPrivacy       bool          `json:"TechPrivacy,omitempty"`
	UpdatedDate       time.Time     `json:"UpdatedDate,omitempty"`
	WhoIsServer       string        `json:"WhoIsServer,omitempty"`
}

// GetOperationDetailRequest is undocumented.
type GetOperationDetailRequest struct {
	OperationID string `json:"OperationId"`
}

// GetOperationDetailResponse is undocumented.
type GetOperationDetailResponse struct {
	DomainName    string    `json:"DomainName,omitempty"`
	Message       string    `json:"Message,omitempty"`
	OperationID   string    `json:"OperationId,omitempty"`
	Status        string    `json:"Status,omitempty"`
	SubmittedDate time.Time `json:"SubmittedDate,omitempty"`
	Type          string    `json:"Type,omitempty"`
}

// ListDomainsRequest is undocumented.
type ListDomainsRequest struct {
	Marker   string `json:"Marker,omitempty"`
	MaxItems int    `json:"MaxItems,omitempty"`
}

// ListDomainsResponse is undocumented.
type ListDomainsResponse struct {
	Domains        []DomainSummary `json:"Domains"`
	NextPageMarker string          `json:"NextPageMarker,omitempty"`
}

// ListOperationsRequest is undocumented.
type ListOperationsRequest struct {
	Marker   string `json:"Marker,omitempty"`
	MaxItems int    `json:"MaxItems,omitempty"`
}

// ListOperationsResponse is undocumented.
type ListOperationsResponse struct {
	NextPageMarker string             `json:"NextPageMarker,omitempty"`
	Operations     []OperationSummary `json:"Operations"`
}

// Nameserver is undocumented.
type Nameserver struct {
	GlueIPs []string `json:"GlueIps,omitempty"`
	Name    string   `json:"Name"`
}

// OperationSummary is undocumented.
type OperationSummary struct {
	OperationID   string    `json:"OperationId"`
	Status        string    `json:"Status"`
	SubmittedDate time.Time `json:"SubmittedDate"`
	Type          string    `json:"Type"`
}

// RegisterDomainRequest is undocumented.
type RegisterDomainRequest struct {
	AdminContact                    ContactDetail `json:"AdminContact"`
	AutoRenew                       bool          `json:"AutoRenew,omitempty"`
	DomainName                      string        `json:"DomainName"`
	DurationInYears                 int           `json:"DurationInYears"`
	IdnLangCode                     string        `json:"IdnLangCode,omitempty"`
	PrivacyProtectAdminContact      bool          `json:"PrivacyProtectAdminContact,omitempty"`
	PrivacyProtectRegistrantContact bool          `json:"PrivacyProtectRegistrantContact,omitempty"`
	PrivacyProtectTechContact       bool          `json:"PrivacyProtectTechContact,omitempty"`
	RegistrantContact               ContactDetail `json:"RegistrantContact"`
	TechContact                     ContactDetail `json:"TechContact"`
}

// RegisterDomainResponse is undocumented.
type RegisterDomainResponse struct {
	OperationID string `json:"OperationId"`
}

// RetrieveDomainAuthCodeRequest is undocumented.
type RetrieveDomainAuthCodeRequest struct {
	DomainName string `json:"DomainName"`
}

// RetrieveDomainAuthCodeResponse is undocumented.
type RetrieveDomainAuthCodeResponse struct {
	AuthCode string `json:"AuthCode"`
}

// TransferDomainRequest is undocumented.
type TransferDomainRequest struct {
	AdminContact                    ContactDetail `json:"AdminContact"`
	AuthCode                        string        `json:"AuthCode,omitempty"`
	AutoRenew                       bool          `json:"AutoRenew,omitempty"`
	DomainName                      string        `json:"DomainName"`
	DurationInYears                 int           `json:"DurationInYears"`
	IdnLangCode                     string        `json:"IdnLangCode,omitempty"`
	Nameservers                     []Nameserver  `json:"Nameservers"`
	PrivacyProtectAdminContact      bool          `json:"PrivacyProtectAdminContact,omitempty"`
	PrivacyProtectRegistrantContact bool          `json:"PrivacyProtectRegistrantContact,omitempty"`
	PrivacyProtectTechContact       bool          `json:"PrivacyProtectTechContact,omitempty"`
	RegistrantContact               ContactDetail `json:"RegistrantContact"`
	TechContact                     ContactDetail `json:"TechContact"`
}

// TransferDomainResponse is undocumented.
type TransferDomainResponse struct {
	OperationID string `json:"OperationId"`
}

// UpdateDomainContactPrivacyRequest is undocumented.
type UpdateDomainContactPrivacyRequest struct {
	AdminPrivacy      bool   `json:"AdminPrivacy,omitempty"`
	DomainName        string `json:"DomainName"`
	RegistrantPrivacy bool   `json:"RegistrantPrivacy,omitempty"`
	TechPrivacy       bool   `json:"TechPrivacy,omitempty"`
}

// UpdateDomainContactPrivacyResponse is undocumented.
type UpdateDomainContactPrivacyResponse struct {
	OperationID string `json:"OperationId"`
}

// UpdateDomainContactRequest is undocumented.
type UpdateDomainContactRequest struct {
	AdminContact      ContactDetail `json:"AdminContact,omitempty"`
	DomainName        string        `json:"DomainName"`
	RegistrantContact ContactDetail `json:"RegistrantContact,omitempty"`
	TechContact       ContactDetail `json:"TechContact,omitempty"`
}

// UpdateDomainContactResponse is undocumented.
type UpdateDomainContactResponse struct {
	OperationID string `json:"OperationId"`
}

// UpdateDomainNameserversRequest is undocumented.
type UpdateDomainNameserversRequest struct {
	DomainName  string       `json:"DomainName"`
	Nameservers []Nameserver `json:"Nameservers"`
}

// UpdateDomainNameserversResponse is undocumented.
type UpdateDomainNameserversResponse struct {
	OperationID string `json:"OperationId"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
