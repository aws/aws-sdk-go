// Package cloudfront provides a client for Amazon CloudFront.
package cloudfront

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

// CloudFront is a client for Amazon CloudFront.
type CloudFront struct {
	client *aws.RestClient
}

// New returns a new CloudFront client.
func New(creds aws.Credentials, region string, client *http.Client) *CloudFront {
	if client == nil {
		client = http.DefaultClient
	}

	service := "cloudfront"
	endpoint, service, region := endpoints.Lookup("cloudfront", region)

	return &CloudFront{
		client: &aws.RestClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2014-10-21",
		},
	}
}

// CreateCloudFrontOriginAccessIdentity is undocumented.
func (c *CloudFront) CreateCloudFrontOriginAccessIdentity(req CreateCloudFrontOriginAccessIdentityRequest) (resp *CreateCloudFrontOriginAccessIdentityResult, err error) {
	resp = &CreateCloudFrontOriginAccessIdentityResult{}

	var body io.Reader
	var contentType string

	contentType = "application/xml"
	b, err := xml.Marshal(req.CloudFrontOriginAccessIdentityConfig)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-10-21/origin-access-identity/cloudfront"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	if s := httpResp.Header.Get("Location"); s != "" {

		resp.Location = s

	}

	return
}

// CreateDistribution is undocumented.
func (c *CloudFront) CreateDistribution(req CreateDistributionRequest) (resp *CreateDistributionResult, err error) {
	resp = &CreateDistributionResult{}

	var body io.Reader
	var contentType string

	contentType = "application/xml"
	b, err := xml.Marshal(req.DistributionConfig)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-10-21/distribution"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	if s := httpResp.Header.Get("Location"); s != "" {

		resp.Location = s

	}

	return
}

// CreateInvalidation is undocumented.
func (c *CloudFront) CreateInvalidation(req CreateInvalidationRequest) (resp *CreateInvalidationResult, err error) {
	resp = &CreateInvalidationResult{}

	var body io.Reader
	var contentType string

	contentType = "application/xml"
	b, err := xml.Marshal(req.InvalidationBatch)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-10-21/distribution/{DistributionId}/invalidation"

	uri = strings.Replace(uri, "{"+"DistributionId"+"}", req.DistributionID, -1)
	uri = strings.Replace(uri, "{"+"DistributionId+"+"}", req.DistributionID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("Location"); s != "" {

		resp.Location = s

	}

	return
}

// CreateStreamingDistribution is undocumented.
func (c *CloudFront) CreateStreamingDistribution(req CreateStreamingDistributionRequest) (resp *CreateStreamingDistributionResult, err error) {
	resp = &CreateStreamingDistributionResult{}

	var body io.Reader
	var contentType string

	contentType = "application/xml"
	b, err := xml.Marshal(req.StreamingDistributionConfig)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-10-21/streaming-distribution"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	if s := httpResp.Header.Get("Location"); s != "" {

		resp.Location = s

	}

	return
}

// DeleteCloudFrontOriginAccessIdentity is undocumented.
func (c *CloudFront) DeleteCloudFrontOriginAccessIdentity(req DeleteCloudFrontOriginAccessIdentityRequest) (err error) {
	// NRE

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/origin-access-identity/cloudfront/{Id}"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if s := req.IfMatch; s != "" {

		httpReq.Header.Set("If-Match", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	return
}

// DeleteDistribution is undocumented.
func (c *CloudFront) DeleteDistribution(req DeleteDistributionRequest) (err error) {
	// NRE

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/distribution/{Id}"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if s := req.IfMatch; s != "" {

		httpReq.Header.Set("If-Match", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	return
}

// DeleteStreamingDistribution is undocumented.
func (c *CloudFront) DeleteStreamingDistribution(req DeleteStreamingDistributionRequest) (err error) {
	// NRE

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/streaming-distribution/{Id}"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if s := req.IfMatch; s != "" {

		httpReq.Header.Set("If-Match", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	return
}

// GetCloudFrontOriginAccessIdentity get the information about an origin
// access identity.
func (c *CloudFront) GetCloudFrontOriginAccessIdentity(req GetCloudFrontOriginAccessIdentityRequest) (resp *GetCloudFrontOriginAccessIdentityResult, err error) {
	resp = &GetCloudFrontOriginAccessIdentityResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/origin-access-identity/cloudfront/{Id}"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// GetCloudFrontOriginAccessIdentityConfig get the configuration
// information about an origin access identity.
func (c *CloudFront) GetCloudFrontOriginAccessIdentityConfig(req GetCloudFrontOriginAccessIdentityConfigRequest) (resp *GetCloudFrontOriginAccessIdentityConfigResult, err error) {
	resp = &GetCloudFrontOriginAccessIdentityConfigResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/origin-access-identity/cloudfront/{Id}/config"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// GetDistribution is undocumented.
func (c *CloudFront) GetDistribution(req GetDistributionRequest) (resp *GetDistributionResult, err error) {
	resp = &GetDistributionResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/distribution/{Id}"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// GetDistributionConfig get the configuration information about a
// distribution.
func (c *CloudFront) GetDistributionConfig(req GetDistributionConfigRequest) (resp *GetDistributionConfigResult, err error) {
	resp = &GetDistributionConfigResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/distribution/{Id}/config"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// GetInvalidation is undocumented.
func (c *CloudFront) GetInvalidation(req GetInvalidationRequest) (resp *GetInvalidationResult, err error) {
	resp = &GetInvalidationResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/distribution/{DistributionId}/invalidation/{Id}"

	uri = strings.Replace(uri, "{"+"DistributionId"+"}", req.DistributionID, -1)
	uri = strings.Replace(uri, "{"+"DistributionId+"+"}", req.DistributionID, -1)

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// GetStreamingDistribution get the information about a streaming
// distribution.
func (c *CloudFront) GetStreamingDistribution(req GetStreamingDistributionRequest) (resp *GetStreamingDistributionResult, err error) {
	resp = &GetStreamingDistributionResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/streaming-distribution/{Id}"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// GetStreamingDistributionConfig get the configuration information about a
// streaming distribution.
func (c *CloudFront) GetStreamingDistributionConfig(req GetStreamingDistributionConfigRequest) (resp *GetStreamingDistributionConfigResult, err error) {
	resp = &GetStreamingDistributionConfigResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/streaming-distribution/{Id}/config"

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

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// ListCloudFrontOriginAccessIdentities is undocumented.
func (c *CloudFront) ListCloudFrontOriginAccessIdentities(req ListCloudFrontOriginAccessIdentitiesRequest) (resp *ListCloudFrontOriginAccessIdentitiesResult, err error) {
	resp = &ListCloudFrontOriginAccessIdentitiesResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/origin-access-identity/cloudfront"

	q := url.Values{}

	if s := req.Marker; s != "" {

		q.Set("Marker", s)
	}

	if s := req.MaxItems; s != "" {

		q.Set("MaxItems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// ListDistributions is undocumented.
func (c *CloudFront) ListDistributions(req ListDistributionsRequest) (resp *ListDistributionsResult, err error) {
	resp = &ListDistributionsResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/distribution"

	q := url.Values{}

	if s := req.Marker; s != "" {

		q.Set("Marker", s)
	}

	if s := req.MaxItems; s != "" {

		q.Set("MaxItems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// ListInvalidations is undocumented.
func (c *CloudFront) ListInvalidations(req ListInvalidationsRequest) (resp *ListInvalidationsResult, err error) {
	resp = &ListInvalidationsResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/distribution/{DistributionId}/invalidation"

	uri = strings.Replace(uri, "{"+"DistributionId"+"}", req.DistributionID, -1)
	uri = strings.Replace(uri, "{"+"DistributionId+"+"}", req.DistributionID, -1)

	q := url.Values{}

	if s := req.Marker; s != "" {

		q.Set("Marker", s)
	}

	if s := req.MaxItems; s != "" {

		q.Set("MaxItems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// ListStreamingDistributions is undocumented.
func (c *CloudFront) ListStreamingDistributions(req ListStreamingDistributionsRequest) (resp *ListStreamingDistributionsResult, err error) {
	resp = &ListStreamingDistributionsResult{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-10-21/streaming-distribution"

	q := url.Values{}

	if s := req.Marker; s != "" {

		q.Set("Marker", s)
	}

	if s := req.MaxItems; s != "" {

		q.Set("MaxItems", s)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// UpdateCloudFrontOriginAccessIdentity is undocumented.
func (c *CloudFront) UpdateCloudFrontOriginAccessIdentity(req UpdateCloudFrontOriginAccessIdentityRequest) (resp *UpdateCloudFrontOriginAccessIdentityResult, err error) {
	resp = &UpdateCloudFrontOriginAccessIdentityResult{}

	var body io.Reader
	var contentType string

	contentType = "application/xml"
	b, err := xml.Marshal(req.CloudFrontOriginAccessIdentityConfig)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-10-21/origin-access-identity/cloudfront/{Id}/config"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if s := req.IfMatch; s != "" {

		httpReq.Header.Set("If-Match", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// UpdateDistribution is undocumented.
func (c *CloudFront) UpdateDistribution(req UpdateDistributionRequest) (resp *UpdateDistributionResult, err error) {
	resp = &UpdateDistributionResult{}

	var body io.Reader
	var contentType string

	contentType = "application/xml"
	b, err := xml.Marshal(req.DistributionConfig)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-10-21/distribution/{Id}/config"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if s := req.IfMatch; s != "" {

		httpReq.Header.Set("If-Match", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// UpdateStreamingDistribution is undocumented.
func (c *CloudFront) UpdateStreamingDistribution(req UpdateStreamingDistributionRequest) (resp *UpdateStreamingDistributionResult, err error) {
	resp = &UpdateStreamingDistributionResult{}

	var body io.Reader
	var contentType string

	contentType = "application/xml"
	b, err := xml.Marshal(req.StreamingDistributionConfig)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-10-21/streaming-distribution/{Id}/config"

	uri = strings.Replace(uri, "{"+"Id"+"}", req.ID, -1)
	uri = strings.Replace(uri, "{"+"Id+"+"}", req.ID, -1)

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if s := req.IfMatch; s != "" {

		httpReq.Header.Set("If-Match", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	if s := httpResp.Header.Get("ETag"); s != "" {

		resp.ETag = s

	}

	return
}

// ActiveTrustedSigners is undocumented.
type ActiveTrustedSigners struct {
	Enabled  bool     `xml:"Enabled"`
	Items    []Signer `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// Aliases is undocumented.
type Aliases struct {
	Items    []string `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// AllowedMethods is undocumented.
type AllowedMethods struct {
	CachedMethods CachedMethods `xml:"CachedMethods"`
	Items         []string      `xml:"Items"`
	Quantity      int           `xml:"Quantity"`
}

// CacheBehavior is undocumented.
type CacheBehavior struct {
	AllowedMethods       AllowedMethods  `xml:"AllowedMethods"`
	ForwardedValues      ForwardedValues `xml:"ForwardedValues"`
	MinTTL               int64           `xml:"MinTTL"`
	PathPattern          string          `xml:"PathPattern"`
	SmoothStreaming      bool            `xml:"SmoothStreaming"`
	TargetOriginID       string          `xml:"TargetOriginId"`
	TrustedSigners       TrustedSigners  `xml:"TrustedSigners"`
	ViewerProtocolPolicy string          `xml:"ViewerProtocolPolicy"`
}

// CacheBehaviors is undocumented.
type CacheBehaviors struct {
	Items    []CacheBehavior `xml:"Items"`
	Quantity int             `xml:"Quantity"`
}

// CachedMethods is undocumented.
type CachedMethods struct {
	Items    []string `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// CloudFrontOriginAccessIdentity is undocumented.
type CloudFrontOriginAccessIdentity struct {
	CloudFrontOriginAccessIdentityConfig CloudFrontOriginAccessIdentityConfig `xml:"CloudFrontOriginAccessIdentityConfig"`
	ID                                   string                               `xml:"Id"`
	S3CanonicalUserID                    string                               `xml:"S3CanonicalUserId"`
}

// CloudFrontOriginAccessIdentityConfig is undocumented.
type CloudFrontOriginAccessIdentityConfig struct {
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"Comment"`
}

// CloudFrontOriginAccessIdentityList is undocumented.
type CloudFrontOriginAccessIdentityList struct {
	IsTruncated bool                                    `xml:"IsTruncated"`
	Items       []CloudFrontOriginAccessIdentitySummary `xml:"Items"`
	Marker      string                                  `xml:"Marker"`
	MaxItems    int                                     `xml:"MaxItems"`
	NextMarker  string                                  `xml:"NextMarker"`
	Quantity    int                                     `xml:"Quantity"`
}

// CloudFrontOriginAccessIdentitySummary is undocumented.
type CloudFrontOriginAccessIdentitySummary struct {
	Comment           string `xml:"Comment"`
	ID                string `xml:"Id"`
	S3CanonicalUserID string `xml:"S3CanonicalUserId"`
}

// CookieNames is undocumented.
type CookieNames struct {
	Items    []string `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// CookiePreference is undocumented.
type CookiePreference struct {
	Forward          string      `xml:"Forward"`
	WhitelistedNames CookieNames `xml:"WhitelistedNames"`
}

// CreateCloudFrontOriginAccessIdentityRequest is undocumented.
type CreateCloudFrontOriginAccessIdentityRequest struct {
	CloudFrontOriginAccessIdentityConfig CloudFrontOriginAccessIdentityConfig `xml:"CloudFrontOriginAccessIdentityConfig"`
}

// CreateCloudFrontOriginAccessIdentityResult is undocumented.
type CreateCloudFrontOriginAccessIdentityResult struct {
	CloudFrontOriginAccessIdentity CloudFrontOriginAccessIdentity `xml:"CloudFrontOriginAccessIdentity"`
	ETag                           string                         `xml:"ETag"`
	Location                       string                         `xml:"Location"`
}

// CreateDistributionRequest is undocumented.
type CreateDistributionRequest struct {
	DistributionConfig DistributionConfig `xml:"DistributionConfig"`
}

// CreateDistributionResult is undocumented.
type CreateDistributionResult struct {
	Distribution Distribution `xml:"Distribution"`
	ETag         string       `xml:"ETag"`
	Location     string       `xml:"Location"`
}

// CreateInvalidationRequest is undocumented.
type CreateInvalidationRequest struct {
	DistributionID    string            `xml:"DistributionId"`
	InvalidationBatch InvalidationBatch `xml:"InvalidationBatch"`
}

// CreateInvalidationResult is undocumented.
type CreateInvalidationResult struct {
	Invalidation Invalidation `xml:"Invalidation"`
	Location     string       `xml:"Location"`
}

// CreateStreamingDistributionRequest is undocumented.
type CreateStreamingDistributionRequest struct {
	StreamingDistributionConfig StreamingDistributionConfig `xml:"StreamingDistributionConfig"`
}

// CreateStreamingDistributionResult is undocumented.
type CreateStreamingDistributionResult struct {
	ETag                  string                `xml:"ETag"`
	Location              string                `xml:"Location"`
	StreamingDistribution StreamingDistribution `xml:"StreamingDistribution"`
}

// CustomErrorResponse is undocumented.
type CustomErrorResponse struct {
	ErrorCachingMinTTL int64  `xml:"ErrorCachingMinTTL"`
	ErrorCode          int    `xml:"ErrorCode"`
	ResponseCode       string `xml:"ResponseCode"`
	ResponsePagePath   string `xml:"ResponsePagePath"`
}

// CustomErrorResponses is undocumented.
type CustomErrorResponses struct {
	Items    []CustomErrorResponse `xml:"Items"`
	Quantity int                   `xml:"Quantity"`
}

// CustomOriginConfig is undocumented.
type CustomOriginConfig struct {
	HTTPPort             int    `xml:"HTTPPort"`
	HTTPSPort            int    `xml:"HTTPSPort"`
	OriginProtocolPolicy string `xml:"OriginProtocolPolicy"`
}

// DefaultCacheBehavior is undocumented.
type DefaultCacheBehavior struct {
	AllowedMethods       AllowedMethods  `xml:"AllowedMethods"`
	ForwardedValues      ForwardedValues `xml:"ForwardedValues"`
	MinTTL               int64           `xml:"MinTTL"`
	SmoothStreaming      bool            `xml:"SmoothStreaming"`
	TargetOriginID       string          `xml:"TargetOriginId"`
	TrustedSigners       TrustedSigners  `xml:"TrustedSigners"`
	ViewerProtocolPolicy string          `xml:"ViewerProtocolPolicy"`
}

// DeleteCloudFrontOriginAccessIdentityRequest is undocumented.
type DeleteCloudFrontOriginAccessIdentityRequest struct {
	ID      string `xml:"Id"`
	IfMatch string `xml:"If-Match"`
}

// DeleteDistributionRequest is undocumented.
type DeleteDistributionRequest struct {
	ID      string `xml:"Id"`
	IfMatch string `xml:"If-Match"`
}

// DeleteStreamingDistributionRequest is undocumented.
type DeleteStreamingDistributionRequest struct {
	ID      string `xml:"Id"`
	IfMatch string `xml:"If-Match"`
}

// Distribution is undocumented.
type Distribution struct {
	ActiveTrustedSigners          ActiveTrustedSigners `xml:"ActiveTrustedSigners"`
	DistributionConfig            DistributionConfig   `xml:"DistributionConfig"`
	DomainName                    string               `xml:"DomainName"`
	ID                            string               `xml:"Id"`
	InProgressInvalidationBatches int                  `xml:"InProgressInvalidationBatches"`
	LastModifiedTime              time.Time            `xml:"LastModifiedTime"`
	Status                        string               `xml:"Status"`
}

// DistributionConfig is undocumented.
type DistributionConfig struct {
	Aliases              Aliases              `xml:"Aliases"`
	CacheBehaviors       CacheBehaviors       `xml:"CacheBehaviors"`
	CallerReference      string               `xml:"CallerReference"`
	Comment              string               `xml:"Comment"`
	CustomErrorResponses CustomErrorResponses `xml:"CustomErrorResponses"`
	DefaultCacheBehavior DefaultCacheBehavior `xml:"DefaultCacheBehavior"`
	DefaultRootObject    string               `xml:"DefaultRootObject"`
	Enabled              bool                 `xml:"Enabled"`
	Logging              LoggingConfig        `xml:"Logging"`
	Origins              Origins              `xml:"Origins"`
	PriceClass           string               `xml:"PriceClass"`
	Restrictions         Restrictions         `xml:"Restrictions"`
	ViewerCertificate    ViewerCertificate    `xml:"ViewerCertificate"`
}

// DistributionList is undocumented.
type DistributionList struct {
	IsTruncated bool                  `xml:"IsTruncated"`
	Items       []DistributionSummary `xml:"Items"`
	Marker      string                `xml:"Marker"`
	MaxItems    int                   `xml:"MaxItems"`
	NextMarker  string                `xml:"NextMarker"`
	Quantity    int                   `xml:"Quantity"`
}

// DistributionSummary is undocumented.
type DistributionSummary struct {
	Aliases              Aliases              `xml:"Aliases"`
	CacheBehaviors       CacheBehaviors       `xml:"CacheBehaviors"`
	Comment              string               `xml:"Comment"`
	CustomErrorResponses CustomErrorResponses `xml:"CustomErrorResponses"`
	DefaultCacheBehavior DefaultCacheBehavior `xml:"DefaultCacheBehavior"`
	DomainName           string               `xml:"DomainName"`
	Enabled              bool                 `xml:"Enabled"`
	ID                   string               `xml:"Id"`
	LastModifiedTime     time.Time            `xml:"LastModifiedTime"`
	Origins              Origins              `xml:"Origins"`
	PriceClass           string               `xml:"PriceClass"`
	Restrictions         Restrictions         `xml:"Restrictions"`
	Status               string               `xml:"Status"`
	ViewerCertificate    ViewerCertificate    `xml:"ViewerCertificate"`
}

// ForwardedValues is undocumented.
type ForwardedValues struct {
	Cookies     CookiePreference `xml:"Cookies"`
	Headers     Headers          `xml:"Headers"`
	QueryString bool             `xml:"QueryString"`
}

// GeoRestriction is undocumented.
type GeoRestriction struct {
	Items           []string `xml:"Items"`
	Quantity        int      `xml:"Quantity"`
	RestrictionType string   `xml:"RestrictionType"`
}

// GetCloudFrontOriginAccessIdentityConfigRequest is undocumented.
type GetCloudFrontOriginAccessIdentityConfigRequest struct {
	ID string `xml:"Id"`
}

// GetCloudFrontOriginAccessIdentityConfigResult is undocumented.
type GetCloudFrontOriginAccessIdentityConfigResult struct {
	CloudFrontOriginAccessIdentityConfig CloudFrontOriginAccessIdentityConfig `xml:"CloudFrontOriginAccessIdentityConfig"`
	ETag                                 string                               `xml:"ETag"`
}

// GetCloudFrontOriginAccessIdentityRequest is undocumented.
type GetCloudFrontOriginAccessIdentityRequest struct {
	ID string `xml:"Id"`
}

// GetCloudFrontOriginAccessIdentityResult is undocumented.
type GetCloudFrontOriginAccessIdentityResult struct {
	CloudFrontOriginAccessIdentity CloudFrontOriginAccessIdentity `xml:"CloudFrontOriginAccessIdentity"`
	ETag                           string                         `xml:"ETag"`
}

// GetDistributionConfigRequest is undocumented.
type GetDistributionConfigRequest struct {
	ID string `xml:"Id"`
}

// GetDistributionConfigResult is undocumented.
type GetDistributionConfigResult struct {
	DistributionConfig DistributionConfig `xml:"DistributionConfig"`
	ETag               string             `xml:"ETag"`
}

// GetDistributionRequest is undocumented.
type GetDistributionRequest struct {
	ID string `xml:"Id"`
}

// GetDistributionResult is undocumented.
type GetDistributionResult struct {
	Distribution Distribution `xml:"Distribution"`
	ETag         string       `xml:"ETag"`
}

// GetInvalidationRequest is undocumented.
type GetInvalidationRequest struct {
	DistributionID string `xml:"DistributionId"`
	ID             string `xml:"Id"`
}

// GetInvalidationResult is undocumented.
type GetInvalidationResult struct {
	Invalidation Invalidation `xml:"Invalidation"`
}

// GetStreamingDistributionConfigRequest is undocumented.
type GetStreamingDistributionConfigRequest struct {
	ID string `xml:"Id"`
}

// GetStreamingDistributionConfigResult is undocumented.
type GetStreamingDistributionConfigResult struct {
	ETag                        string                      `xml:"ETag"`
	StreamingDistributionConfig StreamingDistributionConfig `xml:"StreamingDistributionConfig"`
}

// GetStreamingDistributionRequest is undocumented.
type GetStreamingDistributionRequest struct {
	ID string `xml:"Id"`
}

// GetStreamingDistributionResult is undocumented.
type GetStreamingDistributionResult struct {
	ETag                  string                `xml:"ETag"`
	StreamingDistribution StreamingDistribution `xml:"StreamingDistribution"`
}

// Headers is undocumented.
type Headers struct {
	Items    []string `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// Invalidation is undocumented.
type Invalidation struct {
	CreateTime        time.Time         `xml:"CreateTime"`
	ID                string            `xml:"Id"`
	InvalidationBatch InvalidationBatch `xml:"InvalidationBatch"`
	Status            string            `xml:"Status"`
}

// InvalidationBatch is undocumented.
type InvalidationBatch struct {
	CallerReference string `xml:"CallerReference"`
	Paths           Paths  `xml:"Paths"`
}

// InvalidationList is undocumented.
type InvalidationList struct {
	IsTruncated bool                  `xml:"IsTruncated"`
	Items       []InvalidationSummary `xml:"Items"`
	Marker      string                `xml:"Marker"`
	MaxItems    int                   `xml:"MaxItems"`
	NextMarker  string                `xml:"NextMarker"`
	Quantity    int                   `xml:"Quantity"`
}

// InvalidationSummary is undocumented.
type InvalidationSummary struct {
	CreateTime time.Time `xml:"CreateTime"`
	ID         string    `xml:"Id"`
	Status     string    `xml:"Status"`
}

// KeyPairIds is undocumented.
type KeyPairIds struct {
	Items    []string `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// ListCloudFrontOriginAccessIdentitiesRequest is undocumented.
type ListCloudFrontOriginAccessIdentitiesRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems string `xml:"MaxItems"`
}

// ListCloudFrontOriginAccessIdentitiesResult is undocumented.
type ListCloudFrontOriginAccessIdentitiesResult struct {
	CloudFrontOriginAccessIdentityList CloudFrontOriginAccessIdentityList `xml:"CloudFrontOriginAccessIdentityList"`
}

// ListDistributionsRequest is undocumented.
type ListDistributionsRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems string `xml:"MaxItems"`
}

// ListDistributionsResult is undocumented.
type ListDistributionsResult struct {
	DistributionList DistributionList `xml:"DistributionList"`
}

// ListInvalidationsRequest is undocumented.
type ListInvalidationsRequest struct {
	DistributionID string `xml:"DistributionId"`
	Marker         string `xml:"Marker"`
	MaxItems       string `xml:"MaxItems"`
}

// ListInvalidationsResult is undocumented.
type ListInvalidationsResult struct {
	InvalidationList InvalidationList `xml:"InvalidationList"`
}

// ListStreamingDistributionsRequest is undocumented.
type ListStreamingDistributionsRequest struct {
	Marker   string `xml:"Marker"`
	MaxItems string `xml:"MaxItems"`
}

// ListStreamingDistributionsResult is undocumented.
type ListStreamingDistributionsResult struct {
	StreamingDistributionList StreamingDistributionList `xml:"StreamingDistributionList"`
}

// LoggingConfig is undocumented.
type LoggingConfig struct {
	Bucket         string `xml:"Bucket"`
	Enabled        bool   `xml:"Enabled"`
	IncludeCookies bool   `xml:"IncludeCookies"`
	Prefix         string `xml:"Prefix"`
}

// Origin is undocumented.
type Origin struct {
	CustomOriginConfig CustomOriginConfig `xml:"CustomOriginConfig"`
	DomainName         string             `xml:"DomainName"`
	ID                 string             `xml:"Id"`
	S3OriginConfig     S3OriginConfig     `xml:"S3OriginConfig"`
}

// Origins is undocumented.
type Origins struct {
	Items    []Origin `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// Paths is undocumented.
type Paths struct {
	Items    []string `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// Restrictions is undocumented.
type Restrictions struct {
	GeoRestriction GeoRestriction `xml:"GeoRestriction"`
}

// S3Origin is undocumented.
type S3Origin struct {
	DomainName           string `xml:"DomainName"`
	OriginAccessIdentity string `xml:"OriginAccessIdentity"`
}

// S3OriginConfig is undocumented.
type S3OriginConfig struct {
	OriginAccessIdentity string `xml:"OriginAccessIdentity"`
}

// Signer is undocumented.
type Signer struct {
	AwsAccountNumber string     `xml:"AwsAccountNumber"`
	KeyPairIds       KeyPairIds `xml:"KeyPairIds"`
}

// StreamingDistribution is undocumented.
type StreamingDistribution struct {
	ActiveTrustedSigners        ActiveTrustedSigners        `xml:"ActiveTrustedSigners"`
	DomainName                  string                      `xml:"DomainName"`
	ID                          string                      `xml:"Id"`
	LastModifiedTime            time.Time                   `xml:"LastModifiedTime"`
	Status                      string                      `xml:"Status"`
	StreamingDistributionConfig StreamingDistributionConfig `xml:"StreamingDistributionConfig"`
}

// StreamingDistributionConfig is undocumented.
type StreamingDistributionConfig struct {
	Aliases         Aliases                `xml:"Aliases"`
	CallerReference string                 `xml:"CallerReference"`
	Comment         string                 `xml:"Comment"`
	Enabled         bool                   `xml:"Enabled"`
	Logging         StreamingLoggingConfig `xml:"Logging"`
	PriceClass      string                 `xml:"PriceClass"`
	S3Origin        S3Origin               `xml:"S3Origin"`
	TrustedSigners  TrustedSigners         `xml:"TrustedSigners"`
}

// StreamingDistributionList is undocumented.
type StreamingDistributionList struct {
	IsTruncated bool                           `xml:"IsTruncated"`
	Items       []StreamingDistributionSummary `xml:"Items"`
	Marker      string                         `xml:"Marker"`
	MaxItems    int                            `xml:"MaxItems"`
	NextMarker  string                         `xml:"NextMarker"`
	Quantity    int                            `xml:"Quantity"`
}

// StreamingDistributionSummary is undocumented.
type StreamingDistributionSummary struct {
	Aliases          Aliases        `xml:"Aliases"`
	Comment          string         `xml:"Comment"`
	DomainName       string         `xml:"DomainName"`
	Enabled          bool           `xml:"Enabled"`
	ID               string         `xml:"Id"`
	LastModifiedTime time.Time      `xml:"LastModifiedTime"`
	PriceClass       string         `xml:"PriceClass"`
	S3Origin         S3Origin       `xml:"S3Origin"`
	Status           string         `xml:"Status"`
	TrustedSigners   TrustedSigners `xml:"TrustedSigners"`
}

// StreamingLoggingConfig is undocumented.
type StreamingLoggingConfig struct {
	Bucket  string `xml:"Bucket"`
	Enabled bool   `xml:"Enabled"`
	Prefix  string `xml:"Prefix"`
}

// TrustedSigners is undocumented.
type TrustedSigners struct {
	Enabled  bool     `xml:"Enabled"`
	Items    []string `xml:"Items"`
	Quantity int      `xml:"Quantity"`
}

// UpdateCloudFrontOriginAccessIdentityRequest is undocumented.
type UpdateCloudFrontOriginAccessIdentityRequest struct {
	CloudFrontOriginAccessIdentityConfig CloudFrontOriginAccessIdentityConfig `xml:"CloudFrontOriginAccessIdentityConfig"`
	ID                                   string                               `xml:"Id"`
	IfMatch                              string                               `xml:"If-Match"`
}

// UpdateCloudFrontOriginAccessIdentityResult is undocumented.
type UpdateCloudFrontOriginAccessIdentityResult struct {
	CloudFrontOriginAccessIdentity CloudFrontOriginAccessIdentity `xml:"CloudFrontOriginAccessIdentity"`
	ETag                           string                         `xml:"ETag"`
}

// UpdateDistributionRequest is undocumented.
type UpdateDistributionRequest struct {
	DistributionConfig DistributionConfig `xml:"DistributionConfig"`
	ID                 string             `xml:"Id"`
	IfMatch            string             `xml:"If-Match"`
}

// UpdateDistributionResult is undocumented.
type UpdateDistributionResult struct {
	Distribution Distribution `xml:"Distribution"`
	ETag         string       `xml:"ETag"`
}

// UpdateStreamingDistributionRequest is undocumented.
type UpdateStreamingDistributionRequest struct {
	ID                          string                      `xml:"Id"`
	IfMatch                     string                      `xml:"If-Match"`
	StreamingDistributionConfig StreamingDistributionConfig `xml:"StreamingDistributionConfig"`
}

// UpdateStreamingDistributionResult is undocumented.
type UpdateStreamingDistributionResult struct {
	ETag                  string                `xml:"ETag"`
	StreamingDistribution StreamingDistribution `xml:"StreamingDistribution"`
}

// ViewerCertificate is undocumented.
type ViewerCertificate struct {
	CloudFrontDefaultCertificate bool   `xml:"CloudFrontDefaultCertificate"`
	IAMCertificateID             string `xml:"IAMCertificateId"`
	MinimumProtocolVersion       string `xml:"MinimumProtocolVersion"`
	SSLSupportMethod             string `xml:"SSLSupportMethod"`
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
