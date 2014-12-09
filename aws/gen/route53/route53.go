// Package route53 provides a client for Amazon Route 53.
package route53

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// Route53 is a client for Amazon Route 53.
type Route53 struct {
	client *aws.RestXMLClient
}

// New returns a new Route53 client.
func New(key, secret, region string, client *http.Client) *Route53 {
	if client == nil {
		client = http.DefaultClient
	}

	return &Route53{
		client: &aws.RestXMLClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "route53",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoints.Lookup("route53", region),
			APIVersion: "2013-04-01",
		},
	}
}

// AssociateVPCWithHostedZone this action associates a VPC with an hosted
// zone. To associate a VPC with an hosted zone, send a request to the
// 2013-04-01/hostedzone/ hosted zone /associatevpc resource. The request
// body must include an XML document with a
// AssociateVPCWithHostedZoneRequest element. The response returns the
// AssociateVPCWithHostedZoneResponse element that contains ChangeInfo for
// you to track the progress of the AssociateVPCWithHostedZoneRequest you
// made. See GetChange operation for how to track the progress of your
// change.
func (c *Route53) AssociateVPCWithHostedZone(req AssociateVPCWithHostedZoneRequest) (resp *AssociateVPCWithHostedZoneResponse, err error) {
	resp = &AssociateVPCWithHostedZoneResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone/{Id}/associatevpc"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.HostedZoneID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.HostedZoneID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ChangeResourceRecordSets use this action to create or change your
// authoritative DNS information. To use this action, send a request to the
// 2013-04-01/hostedzone/ hosted Zone /rrset resource. The request body
// must include an XML document with a ChangeResourceRecordSetsRequest
// element. Changes are a list of change items and are considered
// transactional. For more information on transactional changes, also known
// as change batches, see Creating, Changing, and Deleting Resource Record
// Sets Using the Route 53 in the Amazon Route 53 Developer Guide Due to
// the nature of transactional changes, you cannot delete the same resource
// record set more than once in a single change batch. If you attempt to
// delete the same change batch more than once, Route 53 returns an
// InvalidChangeBatch error. In response to a ChangeResourceRecordSets
// request, your DNS data is changed on all Route 53 DNS servers.
// Initially, the status of a change is . This means the change has not yet
// propagated to all the authoritative Route 53 DNS servers. When the
// change is propagated to all hosts, the change returns a status of Note
// the following limitations on a ChangeResourceRecordSets request: - A
// request cannot contain more than 100 Change elements. - A request cannot
// contain more than 1000 ResourceRecord elements. The sum of the number of
// characters (including spaces) in all Value elements in a request cannot
// exceed 32,000 characters.
func (c *Route53) ChangeResourceRecordSets(req ChangeResourceRecordSetsRequest) (resp *ChangeResourceRecordSetsResponse, err error) {
	resp = &ChangeResourceRecordSetsResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone/{Id}/rrset/"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.HostedZoneID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.HostedZoneID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ChangeTagsForResource <nil>
func (c *Route53) ChangeTagsForResource(req ChangeTagsForResourceRequest) (resp *ChangeTagsForResourceResponse, err error) {
	resp = &ChangeTagsForResourceResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/tags/{ResourceType}/{ResourceId}"

	uri = strings.Replace(uri, "{"+"ResourceId"+"}", req.ResourceID, -1)
	uri = strings.Replace(uri, "{"+"ResourceId+"+"}", req.ResourceID, -1)

	uri = strings.Replace(uri, "{"+"ResourceType"+"}", req.ResourceType, -1)
	uri = strings.Replace(uri, "{"+"ResourceType+"+"}", req.ResourceType, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// CreateHealthCheck this action creates a new health check. To create a
// new health check, send a request to the 2013-04-01/healthcheck resource.
// The request body must include an XML document with a
// CreateHealthCheckRequest element. The response returns the
// CreateHealthCheckResponse element that contains metadata about the
// health check.
func (c *Route53) CreateHealthCheck(req CreateHealthCheckRequest) (resp *CreateHealthCheckResponse, err error) {
	resp = &CreateHealthCheckResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/healthcheck"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// CreateHostedZone this action creates a new hosted zone. To create a new
// hosted zone, send a request to the 2013-04-01/hostedzone resource. The
// request body must include an XML document with a CreateHostedZoneRequest
// element. The response returns the CreateHostedZoneResponse element that
// contains metadata about the hosted zone. Route 53 automatically creates
// a default SOA record and four NS records for the zone. The NS records in
// the hosted zone are the name servers you give your registrar to delegate
// your domain to. For more information about SOA and NS records, see NS
// and SOA Records that Route 53 Creates for a Hosted Zone in the Amazon
// Route 53 Developer Guide When you create a zone, its initial status is .
// This means that it is not yet available on all DNS servers. The status
// of the zone changes to when the NS and SOA records are available on all
// Route 53 DNS servers. When trying to create a hosted zone using a
// reusable delegation set, you could specify an optional DelegationSetId,
// and Route53 would assign those 4 NS records for the zone, instead of
// alloting a new one.
func (c *Route53) CreateHostedZone(req CreateHostedZoneRequest) (resp *CreateHostedZoneResponse, err error) {
	resp = &CreateHostedZoneResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// CreateReusableDelegationSet this action creates a reusable
// delegationSet. To create a new reusable delegationSet, send a request to
// the 2013-04-01/delegationset resource. The request body must include an
// XML document with a CreateReusableDelegationSetRequest element. The
// response returns the CreateReusableDelegationSetResponse element that
// contains metadata about the delegationSet. If the optional parameter
// HostedZoneId is specified, it marks the delegationSet associated with
// that particular hosted zone as reusable.
func (c *Route53) CreateReusableDelegationSet(req CreateReusableDelegationSetRequest) (resp *CreateReusableDelegationSetResponse, err error) {
	resp = &CreateReusableDelegationSetResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/delegationset"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// DeleteHealthCheck this action deletes a health check. To delete a health
// check, send a request to the 2013-04-01/healthcheck/ health check
// resource. You can delete a health check only if there are no resource
// record sets associated with this health check. If resource record sets
// are associated with this health check, you must disassociate them before
// you can delete your health check. If you try to delete a health check
// that is associated with resource record sets, Route 53 will deny your
// request with a HealthCheckInUse error. For information about
// disassociating the records from your health check, see
// ChangeResourceRecordSets
func (c *Route53) DeleteHealthCheck(req DeleteHealthCheckRequest) (resp *DeleteHealthCheckResponse, err error) {
	resp = &DeleteHealthCheckResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/healthcheck/{HealthCheckId}"

	uri = strings.Replace(uri, "{"+"HealthCheckId"+"}", req.HealthCheckID, -1)
	uri = strings.Replace(uri, "{"+"HealthCheckId+"+"}", req.HealthCheckID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// DeleteHostedZone this action deletes a hosted zone. To delete a hosted
// zone, send a request to the 2013-04-01/hostedzone/ hosted zone resource.
// For more information about deleting a hosted zone, see Deleting a Hosted
// Zone in the Amazon Route 53 Developer Guide You can delete a hosted zone
// only if there are no resource record sets other than the default SOA
// record and NS resource record sets. If your hosted zone contains other
// resource record sets, you must delete them before you can delete your
// hosted zone. If you try to delete a hosted zone that contains other
// resource record sets, Route 53 will deny your request with a
// HostedZoneNotEmpty error. For information about deleting records from
// your hosted zone, see ChangeResourceRecordSets
func (c *Route53) DeleteHostedZone(req DeleteHostedZoneRequest) (resp *DeleteHostedZoneResponse, err error) {
	resp = &DeleteHostedZoneResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone/{Id}"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// DeleteReusableDelegationSet this action deletes a reusable delegation
// set. To delete a reusable delegation set, send a request to the
// 2013-04-01/delegationset/ delegation set resource. You can delete a
// reusable delegation set only if there are no associated hosted zones. If
// your reusable delegation set contains associated hosted zones, you must
// delete them before you can delete your reusable delegation set. If you
// try to delete a reusable delegation set that contains associated hosted
// zones, Route 53 will deny your request with a DelegationSetInUse error.
func (c *Route53) DeleteReusableDelegationSet(req DeleteReusableDelegationSetRequest) (resp *DeleteReusableDelegationSetResponse, err error) {
	resp = &DeleteReusableDelegationSetResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/delegationset/{Id}"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// DisassociateVPCFromHostedZone this action disassociates a VPC from an
// hosted zone. To disassociate a VPC to a hosted zone, send a request to
// the 2013-04-01/hostedzone/ hosted zone /disassociatevpc resource. The
// request body must include an XML document with a
// DisassociateVPCFromHostedZoneRequest element. The response returns the
// DisassociateVPCFromHostedZoneResponse element that contains ChangeInfo
// for you to track the progress of the
// DisassociateVPCFromHostedZoneRequest you made. See GetChange operation
// for how to track the progress of your change.
func (c *Route53) DisassociateVPCFromHostedZone(req DisassociateVPCFromHostedZoneRequest) (resp *DisassociateVPCFromHostedZoneResponse, err error) {
	resp = &DisassociateVPCFromHostedZoneResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone/{Id}/disassociatevpc"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.HostedZoneID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.HostedZoneID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetChange this action returns the current status of a change batch
// request. The status is one of the following values: - indicates that the
// changes in this request have not replicated to all Route 53 DNS servers.
// This is the initial status of all change batch requests. - indicates
// that the changes have replicated to all Amazon Route 53 DNS servers.
func (c *Route53) GetChange(req GetChangeRequest) (resp *GetChangeResponse, err error) {
	resp = &GetChangeResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/change/{Id}"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetCheckerIPRanges to retrieve a list of the IP ranges used by Amazon
// Route 53 health checkers to check the health of your resources, send a
// request to the 2013-04-01/checkeripranges resource. You can use these IP
// addresses to configure router and firewall rules to allow health
// checkers to check the health of your resources.
func (c *Route53) GetCheckerIPRanges(req GetCheckerIPRangesRequest) (resp *GetCheckerIPRangesResponse, err error) {
	resp = &GetCheckerIPRangesResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/checkeripranges"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetGeoLocation to retrieve a single geo location, send a request to the
// 2013-04-01/geolocation resource with one of these options: continentcode
// | countrycode | countrycode and subdivisioncode.
func (c *Route53) GetGeoLocation(req GetGeoLocationRequest) (resp *GetGeoLocationResponse, err error) {
	resp = &GetGeoLocationResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/geolocation"

	q := url.Values{}

	if s := req.ContinentCode; s != "" {

		q.Set("continentcode", s)
	}

	if s := req.CountryCode; s != "" {

		q.Set("countrycode", s)
	}

	if s := req.SubdivisionCode; s != "" {

		q.Set("subdivisioncode", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetHealthCheck to retrieve the health check, send a request to the
// 2013-04-01/healthcheck/ health check resource.
func (c *Route53) GetHealthCheck(req GetHealthCheckRequest) (resp *GetHealthCheckResponse, err error) {
	resp = &GetHealthCheckResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/healthcheck/{HealthCheckId}"

	uri = strings.Replace(uri, "{"+"HealthCheckId"+"}", req.HealthCheckID, -1)
	uri = strings.Replace(uri, "{"+"HealthCheckId+"+"}", req.HealthCheckID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetHealthCheckCount to retrieve a count of all your health checks, send
// a request to the 2013-04-01/healthcheckcount resource.
func (c *Route53) GetHealthCheckCount(req GetHealthCheckCountRequest) (resp *GetHealthCheckCountResponse, err error) {
	resp = &GetHealthCheckCountResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/healthcheckcount"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetHealthCheckLastFailureReason if you want to learn why a health check
// is currently failing or why it failed most recently (if at all), you can
// get the failure reason for the most recent failure. Send a request to
// the 2013-04-01/healthcheck/ health check /lastfailurereason resource.
func (c *Route53) GetHealthCheckLastFailureReason(req GetHealthCheckLastFailureReasonRequest) (resp *GetHealthCheckLastFailureReasonResponse, err error) {
	resp = &GetHealthCheckLastFailureReasonResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/healthcheck/{HealthCheckId}/lastfailurereason"

	uri = strings.Replace(uri, "{"+"HealthCheckId"+"}", req.HealthCheckID, -1)
	uri = strings.Replace(uri, "{"+"HealthCheckId+"+"}", req.HealthCheckID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetHealthCheckStatus to retrieve the health check status, send a request
// to the 2013-04-01/healthcheck/ health check /status resource. You can
// use this call to get a health check's current status.
func (c *Route53) GetHealthCheckStatus(req GetHealthCheckStatusRequest) (resp *GetHealthCheckStatusResponse, err error) {
	resp = &GetHealthCheckStatusResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/healthcheck/{HealthCheckId}/status"

	uri = strings.Replace(uri, "{"+"HealthCheckId"+"}", req.HealthCheckID, -1)
	uri = strings.Replace(uri, "{"+"HealthCheckId+"+"}", req.HealthCheckID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetHostedZone to retrieve the delegation set for a hosted zone, send a
// request to the 2013-04-01/hostedzone/ hosted zone resource. The
// delegation set is the four Route 53 name servers that were assigned to
// the hosted zone when you created it.
func (c *Route53) GetHostedZone(req GetHostedZoneRequest) (resp *GetHostedZoneResponse, err error) {
	resp = &GetHostedZoneResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone/{Id}"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// GetReusableDelegationSet to retrieve the reusable delegation set, send a
// request to the 2013-04-01/delegationset/ delegation set resource.
func (c *Route53) GetReusableDelegationSet(req GetReusableDelegationSetRequest) (resp *GetReusableDelegationSetResponse, err error) {
	resp = &GetReusableDelegationSetResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/delegationset/{Id}"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListGeoLocations to retrieve a list of supported geo locations, send a
// request to the 2013-04-01/geolocations resource. The response to this
// request includes a GeoLocationDetailsList element with zero, one, or
// multiple GeoLocationDetails child elements. The list is sorted by
// country code, and then subdivision code, followed by continents at the
// end of the list. By default, the list of geo locations is displayed on a
// single page. You can control the length of the page that is displayed by
// using the MaxItems parameter. If the list is truncated, IsTruncated will
// be set to true and a combination of NextContinentCode, NextCountryCode,
// NextSubdivisionCode will be populated. You can pass these as parameters
// to StartContinentCode, StartCountryCode, StartSubdivisionCode to control
// the geo location that the list begins with.
func (c *Route53) ListGeoLocations(req ListGeoLocationsRequest) (resp *ListGeoLocationsResponse, err error) {
	resp = &ListGeoLocationsResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/geolocations"

	q := url.Values{}

	if s := req.MaxItems; s != "" {

		q.Set("maxitems", s)
	}

	if s := req.StartContinentCode; s != "" {

		q.Set("startcontinentcode", s)
	}

	if s := req.StartCountryCode; s != "" {

		q.Set("startcountrycode", s)
	}

	if s := req.StartSubdivisionCode; s != "" {

		q.Set("startsubdivisioncode", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListHealthChecks to retrieve a list of your health checks, send a
// request to the 2013-04-01/healthcheck resource. The response to this
// request includes a HealthChecks element with zero, one, or multiple
// HealthCheck child elements. By default, the list of health checks is
// displayed on a single page. You can control the length of the page that
// is displayed by using the MaxItems parameter. You can use the Marker
// parameter to control the health check that the list begins with. Amazon
// Route 53 returns a maximum of 100 items. If you set MaxItems to a value
// greater than 100, Amazon Route 53 returns only the first 100.
func (c *Route53) ListHealthChecks(req ListHealthChecksRequest) (resp *ListHealthChecksResponse, err error) {
	resp = &ListHealthChecksResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/healthcheck"

	q := url.Values{}

	if s := req.Marker; s != "" {

		q.Set("marker", s)
	}

	if s := req.MaxItems; s != "" {

		q.Set("maxitems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListHostedZones to retrieve a list of your hosted zones, send a request
// to the 2013-04-01/hostedzone resource. The response to this request
// includes a HostedZones element with zero, one, or multiple HostedZone
// child elements. By default, the list of hosted zones is displayed on a
// single page. You can control the length of the page that is displayed by
// using the MaxItems parameter. You can use the Marker parameter to
// control the hosted zone that the list begins with. Amazon Route 53
// returns a maximum of 100 items. If you set MaxItems to a value greater
// than 100, Amazon Route 53 returns only the first 100.
func (c *Route53) ListHostedZones(req ListHostedZonesRequest) (resp *ListHostedZonesResponse, err error) {
	resp = &ListHostedZonesResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone"

	q := url.Values{}

	if s := req.DelegationSetID; s != "" {

		q.Set("delegationsetid", s)
	}

	if s := req.Marker; s != "" {

		q.Set("marker", s)
	}

	if s := req.MaxItems; s != "" {

		q.Set("maxitems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListResourceRecordSets imagine all the resource record sets in a zone
// listed out in front of you. Imagine them sorted lexicographically first
// by DNS name (with the labels reversed, like "com.amazon.www" for
// example), and secondarily, lexicographically by record type. This
// operation retrieves at most MaxItems resource record sets from this
// list, in order, starting at a position specified by the Name and Type
// arguments: If both Name and Type are omitted, this means start the
// results at the first in the HostedZone. If Name is specified but Type is
// omitted, this means start the results at the first in the list whose
// name is greater than or equal to Name. If both Name and Type are
// specified, this means start the results at the first in the list whose
// name is greater than or equal to Name and whose type is greater than or
// equal to Type. It is an error to specify the Type but not the Name. Use
// ListResourceRecordSets to retrieve a single known record set by
// specifying the record set's name and type, and setting MaxItems = 1 To
// retrieve all the records in a HostedZone, first pause any processes
// making calls to ChangeResourceRecordSets. Initially call
// ListResourceRecordSets without a Name and Type to get the first page of
// record sets. For subsequent calls, set Name and Type to the NextName and
// NextType values returned by the previous response. In the presence of
// concurrent ChangeResourceRecordSets calls, there is no consistency of
// results across calls to ListResourceRecordSets. The only way to get a
// consistent multi-page snapshot of all RRSETs in a zone is to stop making
// changes while pagination is in progress. However, the results from
// ListResourceRecordSets are consistent within a page. If MakeChange calls
// are taking place concurrently, the result of each one will either be
// completely visible in your results or not at all. You will not see
// partial changes, or changes that do not ultimately succeed. (This
// follows from the fact that MakeChange is atomic) The results from
// ListResourceRecordSets are strongly consistent with
// ChangeResourceRecordSets. To be precise, if a single process makes a
// call to ChangeResourceRecordSets and receives a successful response, the
// effects of that change will be visible in a subsequent call to
// ListResourceRecordSets by that process.
func (c *Route53) ListResourceRecordSets(req ListResourceRecordSetsRequest) (resp *ListResourceRecordSetsResponse, err error) {
	resp = &ListResourceRecordSetsResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone/{Id}/rrset"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.HostedZoneID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.HostedZoneID, -1)

	q := url.Values{}

	if s := req.MaxItems; s != "" {

		q.Set("maxitems", s)
	}

	if s := req.StartRecordIdentifier; s != "" {

		q.Set("identifier", s)
	}

	if s := req.StartRecordName; s != "" {

		q.Set("name", s)
	}

	if s := req.StartRecordType; s != "" {

		q.Set("type", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListReusableDelegationSets to retrieve a list of your reusable
// delegation sets, send a request to the 2013-04-01/delegationset
// resource. The response to this request includes a DelegationSets element
// with zero, one, or multiple DelegationSet child elements. By default,
// the list of delegation sets is displayed on a single page. You can
// control the length of the page that is displayed by using the MaxItems
// parameter. You can use the Marker parameter to control the delegation
// set that the list begins with. Amazon Route 53 returns a maximum of 100
// items. If you set MaxItems to a value greater than 100, Amazon Route 53
// returns only the first 100.
func (c *Route53) ListReusableDelegationSets(req ListReusableDelegationSetsRequest) (resp *ListReusableDelegationSetsResponse, err error) {
	resp = &ListReusableDelegationSetsResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/delegationset"

	q := url.Values{}

	if s := req.Marker; s != "" {

		q.Set("marker", s)
	}

	if s := req.MaxItems; s != "" {

		q.Set("maxitems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListTagsForResource <nil>
func (c *Route53) ListTagsForResource(req ListTagsForResourceRequest) (resp *ListTagsForResourceResponse, err error) {
	resp = &ListTagsForResourceResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/tags/{ResourceType}/{ResourceId}"

	uri = strings.Replace(uri, "{"+"ResourceId"+"}", req.ResourceID, -1)
	uri = strings.Replace(uri, "{"+"ResourceId+"+"}", req.ResourceID, -1)

	uri = strings.Replace(uri, "{"+"ResourceType"+"}", req.ResourceType, -1)
	uri = strings.Replace(uri, "{"+"ResourceType+"+"}", req.ResourceType, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// ListTagsForResources <nil>
func (c *Route53) ListTagsForResources(req ListTagsForResourcesRequest) (resp *ListTagsForResourcesResponse, err error) {
	resp = &ListTagsForResourcesResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/tags/{ResourceType}"

	uri = strings.Replace(uri, "{"+"ResourceType"+"}", req.ResourceType, -1)
	uri = strings.Replace(uri, "{"+"ResourceType+"+"}", req.ResourceType, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// UpdateHealthCheck this action updates an existing health check. To
// update a health check, send a request to the 2013-04-01/healthcheck/
// health check resource. The request body must include an XML document
// with an UpdateHealthCheckRequest element. The response returns an
// UpdateHealthCheckResponse element, which contains metadata about the
// health check.
func (c *Route53) UpdateHealthCheck(req UpdateHealthCheckRequest) (resp *UpdateHealthCheckResponse, err error) {
	resp = &UpdateHealthCheckResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/healthcheck/{HealthCheckId}"

	uri = strings.Replace(uri, "{"+"HealthCheckId"+"}", req.HealthCheckID, -1)
	uri = strings.Replace(uri, "{"+"HealthCheckId+"+"}", req.HealthCheckID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// UpdateHostedZoneComment to update the hosted zone comment, send a
// request to the 2013-04-01/hostedzone/ hosted zone resource. The request
// body must include an XML document with a UpdateHostedZoneCommentRequest
// element. The response to this request includes the modified HostedZone
// element. The comment can have a maximum length of 256 characters.
func (c *Route53) UpdateHostedZoneComment(req UpdateHostedZoneCommentRequest) (resp *UpdateHostedZoneCommentResponse, err error) {
	resp = &UpdateHostedZoneCommentResponse{}

	var body io.Reader

	uri := c.client.Endpoint + "/2013-04-01/hostedzone/{Id}"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	// TODO: extract the rest of the response

	return
}

// AliasTarget is undocumented.
type AliasTarget struct {
	DNSName              string `xml:"DNSName"`
	EvaluateTargetHealth bool   `xml:"EvaluateTargetHealth"`
	HostedZoneID         string `xml:"HostedZoneId"`
}

// AssociateVPCWithHostedZoneRequest is undocumented.
type AssociateVPCWithHostedZoneRequest struct {
	Comment      string `xml:"Comment"`
	HostedZoneID string `xml:"Id"`
	VPC          VPC    `xml:"VPC"`
}

// AssociateVPCWithHostedZoneResponse is undocumented.
type AssociateVPCWithHostedZoneResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

// Change is undocumented.
type Change struct {
	Action            string            `xml:"Action"`
	ResourceRecordSet ResourceRecordSet `xml:"ResourceRecordSet"`
}

// ChangeBatch is undocumented.
type ChangeBatch struct {
	Changes []Change `xml:"Changes"`
	Comment string   `xml:"Comment"`
}

// ChangeInfo is undocumented.
type ChangeInfo struct {
	Comment     string    `xml:"Comment"`
	ID          string    `xml:"Id"`
	Status      string    `xml:"Status"`
	SubmittedAt time.Time `xml:"SubmittedAt"`
}

// ChangeResourceRecordSetsRequest is undocumented.
type ChangeResourceRecordSetsRequest struct {
	ChangeBatch  ChangeBatch `xml:"ChangeBatch"`
	HostedZoneID string      `xml:"Id"`
}

// ChangeResourceRecordSetsResponse is undocumented.
type ChangeResourceRecordSetsResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

// ChangeTagsForResourceRequest is undocumented.
type ChangeTagsForResourceRequest struct {
	AddTags       []Tag    `xml:"AddTags"`
	RemoveTagKeys []string `xml:"RemoveTagKeys"`
	ResourceID    string   `xml:"ResourceId"`
	ResourceType  string   `xml:"ResourceType"`
}

// ChangeTagsForResourceResponse is undocumented.
type ChangeTagsForResourceResponse struct {
}

// CreateHealthCheckRequest is undocumented.
type CreateHealthCheckRequest struct {
	CallerReference   string            `xml:"CallerReference"`
	HealthCheckConfig HealthCheckConfig `xml:"HealthCheckConfig"`
}

// CreateHealthCheckResponse is undocumented.
type CreateHealthCheckResponse struct {
	HealthCheck HealthCheck `xml:"HealthCheck"`
	Location    string      `xml:"Location"`
}

// CreateHostedZoneRequest is undocumented.
type CreateHostedZoneRequest struct {
	CallerReference  string           `xml:"CallerReference"`
	DelegationSetID  string           `xml:"DelegationSetId"`
	HostedZoneConfig HostedZoneConfig `xml:"HostedZoneConfig"`
	Name             string           `xml:"Name"`
	VPC              VPC              `xml:"VPC"`
}

// CreateHostedZoneResponse is undocumented.
type CreateHostedZoneResponse struct {
	ChangeInfo    ChangeInfo    `xml:"ChangeInfo"`
	DelegationSet DelegationSet `xml:"DelegationSet"`
	HostedZone    HostedZone    `xml:"HostedZone"`
	Location      string        `xml:"Location"`
	VPC           VPC           `xml:"VPC"`
}

// CreateReusableDelegationSetRequest is undocumented.
type CreateReusableDelegationSetRequest struct {
	CallerReference string `xml:"CallerReference"`
	HostedZoneID    string `xml:"HostedZoneId"`
}

// CreateReusableDelegationSetResponse is undocumented.
type CreateReusableDelegationSetResponse struct {
	DelegationSet DelegationSet `xml:"DelegationSet"`
	Location      string        `xml:"Location"`
}

// DelegationSet is undocumented.
type DelegationSet struct {
	CallerReference string   `xml:"CallerReference"`
	ID              string   `xml:"Id"`
	NameServers     []string `xml:"NameServers"`
}

// DeleteHealthCheckRequest is undocumented.
type DeleteHealthCheckRequest struct {
	HealthCheckID string `xml:"HealthCheckId"`
}

// DeleteHealthCheckResponse is undocumented.
type DeleteHealthCheckResponse struct {
}

// DeleteHostedZoneRequest is undocumented.
type DeleteHostedZoneRequest struct {
	ID string `xml:"Id"`
}

// DeleteHostedZoneResponse is undocumented.
type DeleteHostedZoneResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

// DeleteReusableDelegationSetRequest is undocumented.
type DeleteReusableDelegationSetRequest struct {
	ID string `xml:"Id"`
}

// DeleteReusableDelegationSetResponse is undocumented.
type DeleteReusableDelegationSetResponse struct {
}

// DisassociateVPCFromHostedZoneRequest is undocumented.
type DisassociateVPCFromHostedZoneRequest struct {
	Comment      string `xml:"Comment"`
	HostedZoneID string `xml:"Id"`
	VPC          VPC    `xml:"VPC"`
}

// DisassociateVPCFromHostedZoneResponse is undocumented.
type DisassociateVPCFromHostedZoneResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

// GeoLocation is undocumented.
type GeoLocation struct {
	ContinentCode   string `xml:"ContinentCode"`
	CountryCode     string `xml:"CountryCode"`
	SubdivisionCode string `xml:"SubdivisionCode"`
}

// GeoLocationDetails is undocumented.
type GeoLocationDetails struct {
	ContinentCode   string `xml:"ContinentCode"`
	ContinentName   string `xml:"ContinentName"`
	CountryCode     string `xml:"CountryCode"`
	CountryName     string `xml:"CountryName"`
	SubdivisionCode string `xml:"SubdivisionCode"`
	SubdivisionName string `xml:"SubdivisionName"`
}

// GetChangeRequest is undocumented.
type GetChangeRequest struct {
	ID string `xml:"Id"`
}

// GetChangeResponse is undocumented.
type GetChangeResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

// GetCheckerIPRangesRequest is undocumented.
type GetCheckerIPRangesRequest struct {
}

// GetCheckerIPRangesResponse is undocumented.
type GetCheckerIPRangesResponse struct {
	CheckerIPRanges []string `xml:"CheckerIpRanges"`
}

// GetGeoLocationRequest is undocumented.
type GetGeoLocationRequest struct {
	ContinentCode   string `xml:"continentcode"`
	CountryCode     string `xml:"countrycode"`
	SubdivisionCode string `xml:"subdivisioncode"`
}

// GetGeoLocationResponse is undocumented.
type GetGeoLocationResponse struct {
	GeoLocationDetails GeoLocationDetails `xml:"GeoLocationDetails"`
}

// GetHealthCheckCountRequest is undocumented.
type GetHealthCheckCountRequest struct {
}

// GetHealthCheckCountResponse is undocumented.
type GetHealthCheckCountResponse struct {
	HealthCheckCount int64 `xml:"HealthCheckCount"`
}

// GetHealthCheckLastFailureReasonRequest is undocumented.
type GetHealthCheckLastFailureReasonRequest struct {
	HealthCheckID string `xml:"HealthCheckId"`
}

// GetHealthCheckLastFailureReasonResponse is undocumented.
type GetHealthCheckLastFailureReasonResponse struct {
	HealthCheckObservations []HealthCheckObservation `xml:"HealthCheckObservations"`
}

// GetHealthCheckRequest is undocumented.
type GetHealthCheckRequest struct {
	HealthCheckID string `xml:"HealthCheckId"`
}

// GetHealthCheckResponse is undocumented.
type GetHealthCheckResponse struct {
	HealthCheck HealthCheck `xml:"HealthCheck"`
}

// GetHealthCheckStatusRequest is undocumented.
type GetHealthCheckStatusRequest struct {
	HealthCheckID string `xml:"HealthCheckId"`
}

// GetHealthCheckStatusResponse is undocumented.
type GetHealthCheckStatusResponse struct {
	HealthCheckObservations []HealthCheckObservation `xml:"HealthCheckObservations"`
}

// GetHostedZoneRequest is undocumented.
type GetHostedZoneRequest struct {
	ID string `xml:"Id"`
}

// GetHostedZoneResponse is undocumented.
type GetHostedZoneResponse struct {
	DelegationSet DelegationSet `xml:"DelegationSet"`
	HostedZone    HostedZone    `xml:"HostedZone"`
	VPCs          []VPC         `xml:"VPCs"`
}

// GetReusableDelegationSetRequest is undocumented.
type GetReusableDelegationSetRequest struct {
	ID string `xml:"Id"`
}

// GetReusableDelegationSetResponse is undocumented.
type GetReusableDelegationSetResponse struct {
	DelegationSet DelegationSet `xml:"DelegationSet"`
}

// HealthCheck is undocumented.
type HealthCheck struct {
	CallerReference    string            `xml:"CallerReference"`
	HealthCheckConfig  HealthCheckConfig `xml:"HealthCheckConfig"`
	HealthCheckVersion int64             `xml:"HealthCheckVersion"`
	ID                 string            `xml:"Id"`
}

// HealthCheckConfig is undocumented.
type HealthCheckConfig struct {
	FailureThreshold         int    `xml:"FailureThreshold"`
	FullyQualifiedDomainName string `xml:"FullyQualifiedDomainName"`
	IPAddress                string `xml:"IPAddress"`
	Port                     int    `xml:"Port"`
	RequestInterval          int    `xml:"RequestInterval"`
	ResourcePath             string `xml:"ResourcePath"`
	SearchString             string `xml:"SearchString"`
	Type                     string `xml:"Type"`
}

// HealthCheckObservation is undocumented.
type HealthCheckObservation struct {
	IPAddress    string       `xml:"IPAddress"`
	StatusReport StatusReport `xml:"StatusReport"`
}

// HostedZone is undocumented.
type HostedZone struct {
	CallerReference        string           `xml:"CallerReference"`
	Config                 HostedZoneConfig `xml:"Config"`
	ID                     string           `xml:"Id"`
	Name                   string           `xml:"Name"`
	ResourceRecordSetCount int64            `xml:"ResourceRecordSetCount"`
}

// HostedZoneConfig is undocumented.
type HostedZoneConfig struct {
	Comment     string `xml:"Comment"`
	PrivateZone bool   `xml:"PrivateZone"`
}

// ListGeoLocationsRequest is undocumented.
type ListGeoLocationsRequest struct {
	MaxItems             string `xml:"maxitems"`
	StartContinentCode   string `xml:"startcontinentcode"`
	StartCountryCode     string `xml:"startcountrycode"`
	StartSubdivisionCode string `xml:"startsubdivisioncode"`
}

// ListGeoLocationsResponse is undocumented.
type ListGeoLocationsResponse struct {
	GeoLocationDetailsList []GeoLocationDetails `xml:"GeoLocationDetailsList"`
	IsTruncated            bool                 `xml:"IsTruncated"`
	MaxItems               string               `xml:"MaxItems"`
	NextContinentCode      string               `xml:"NextContinentCode"`
	NextCountryCode        string               `xml:"NextCountryCode"`
	NextSubdivisionCode    string               `xml:"NextSubdivisionCode"`
}

// ListHealthChecksRequest is undocumented.
type ListHealthChecksRequest struct {
	Marker   string `xml:"marker"`
	MaxItems string `xml:"maxitems"`
}

// ListHealthChecksResponse is undocumented.
type ListHealthChecksResponse struct {
	HealthChecks []HealthCheck `xml:"HealthChecks"`
	IsTruncated  bool          `xml:"IsTruncated"`
	Marker       string        `xml:"Marker"`
	MaxItems     string        `xml:"MaxItems"`
	NextMarker   string        `xml:"NextMarker"`
}

// ListHostedZonesRequest is undocumented.
type ListHostedZonesRequest struct {
	DelegationSetID string `xml:"delegationsetid"`
	Marker          string `xml:"marker"`
	MaxItems        string `xml:"maxitems"`
}

// ListHostedZonesResponse is undocumented.
type ListHostedZonesResponse struct {
	HostedZones []HostedZone `xml:"HostedZones"`
	IsTruncated bool         `xml:"IsTruncated"`
	Marker      string       `xml:"Marker"`
	MaxItems    string       `xml:"MaxItems"`
	NextMarker  string       `xml:"NextMarker"`
}

// ListResourceRecordSetsRequest is undocumented.
type ListResourceRecordSetsRequest struct {
	HostedZoneID          string `xml:"Id"`
	MaxItems              string `xml:"maxitems"`
	StartRecordIdentifier string `xml:"identifier"`
	StartRecordName       string `xml:"name"`
	StartRecordType       string `xml:"type"`
}

// ListResourceRecordSetsResponse is undocumented.
type ListResourceRecordSetsResponse struct {
	IsTruncated          bool                `xml:"IsTruncated"`
	MaxItems             string              `xml:"MaxItems"`
	NextRecordIdentifier string              `xml:"NextRecordIdentifier"`
	NextRecordName       string              `xml:"NextRecordName"`
	NextRecordType       string              `xml:"NextRecordType"`
	ResourceRecordSets   []ResourceRecordSet `xml:"ResourceRecordSets"`
}

// ListReusableDelegationSetsRequest is undocumented.
type ListReusableDelegationSetsRequest struct {
	Marker   string `xml:"marker"`
	MaxItems string `xml:"maxitems"`
}

// ListReusableDelegationSetsResponse is undocumented.
type ListReusableDelegationSetsResponse struct {
	DelegationSets []DelegationSet `xml:"DelegationSets"`
	IsTruncated    bool            `xml:"IsTruncated"`
	Marker         string          `xml:"Marker"`
	MaxItems       string          `xml:"MaxItems"`
	NextMarker     string          `xml:"NextMarker"`
}

// ListTagsForResourceRequest is undocumented.
type ListTagsForResourceRequest struct {
	ResourceID   string `xml:"ResourceId"`
	ResourceType string `xml:"ResourceType"`
}

// ListTagsForResourceResponse is undocumented.
type ListTagsForResourceResponse struct {
	ResourceTagSet ResourceTagSet `xml:"ResourceTagSet"`
}

// ListTagsForResourcesRequest is undocumented.
type ListTagsForResourcesRequest struct {
	ResourceIds  []string `xml:"ResourceIds"`
	ResourceType string   `xml:"ResourceType"`
}

// ListTagsForResourcesResponse is undocumented.
type ListTagsForResourcesResponse struct {
	ResourceTagSets []ResourceTagSet `xml:"ResourceTagSets"`
}

// ResourceRecord is undocumented.
type ResourceRecord struct {
	Value string `xml:"Value"`
}

// ResourceRecordSet is undocumented.
type ResourceRecordSet struct {
	AliasTarget     AliasTarget      `xml:"AliasTarget"`
	Failover        string           `xml:"Failover"`
	GeoLocation     GeoLocation      `xml:"GeoLocation"`
	HealthCheckID   string           `xml:"HealthCheckId"`
	Name            string           `xml:"Name"`
	Region          string           `xml:"Region"`
	ResourceRecords []ResourceRecord `xml:"ResourceRecords"`
	SetIdentifier   string           `xml:"SetIdentifier"`
	TTL             int64            `xml:"TTL"`
	Type            string           `xml:"Type"`
	Weight          int64            `xml:"Weight"`
}

// ResourceTagSet is undocumented.
type ResourceTagSet struct {
	ResourceID   string `xml:"ResourceId"`
	ResourceType string `xml:"ResourceType"`
	Tags         []Tag  `xml:"Tags"`
}

// StatusReport is undocumented.
type StatusReport struct {
	CheckedTime time.Time `xml:"CheckedTime"`
	Status      string    `xml:"Status"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// UpdateHealthCheckRequest is undocumented.
type UpdateHealthCheckRequest struct {
	FailureThreshold         int    `xml:"FailureThreshold"`
	FullyQualifiedDomainName string `xml:"FullyQualifiedDomainName"`
	HealthCheckID            string `xml:"HealthCheckId"`
	HealthCheckVersion       int64  `xml:"HealthCheckVersion"`
	IPAddress                string `xml:"IPAddress"`
	Port                     int    `xml:"Port"`
	ResourcePath             string `xml:"ResourcePath"`
	SearchString             string `xml:"SearchString"`
}

// UpdateHealthCheckResponse is undocumented.
type UpdateHealthCheckResponse struct {
	HealthCheck HealthCheck `xml:"HealthCheck"`
}

// UpdateHostedZoneCommentRequest is undocumented.
type UpdateHostedZoneCommentRequest struct {
	Comment string `xml:"Comment"`
	ID      string `xml:"Id"`
}

// UpdateHostedZoneCommentResponse is undocumented.
type UpdateHostedZoneCommentResponse struct {
	HostedZone HostedZone `xml:"HostedZone"`
}

// VPC is undocumented.
type VPC struct {
	VPCID     string `xml:"VPCId"`
	VPCRegion string `xml:"VPCRegion"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name

var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
