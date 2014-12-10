// Package cloudsearchdomain provides a client for Amazon CloudSearch Domain.
package cloudsearchdomain

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// CloudSearchDomain is a client for Amazon CloudSearch Domain.
type CloudSearchDomain struct {
	client *aws.RestClient
}

// New returns a new CloudSearchDomain client.
func New(creds aws.Credentials, region string, client *http.Client) *CloudSearchDomain {
	if client == nil {
		client = http.DefaultClient
	}

	service := "cloudsearchdomain"
	endpoint, service, region := endpoints.Lookup("cloudsearchdomain", region)

	return &CloudSearchDomain{
		client: &aws.RestClient{
			Credentials: creds,
			Service:     service,
			Region:      region,
			Client:      client,
			Endpoint:    endpoint,
			APIVersion:  "2013-01-01",
		},
	}
}

// Search retrieves a list of documents that match the specified search
// criteria. How you specify the search criteria depends on which query
// parser you use. Amazon CloudSearch supports four query parsers: simple :
// search all text and text-array fields for the specified string. Search
// for phrases, individual terms, and prefixes. structured : search
// specific fields, construct compound queries using Boolean operators, and
// use advanced features such as term boosting and proximity searching.
// lucene : specify search criteria using the Apache Lucene query parser
// syntax. dismax : specify search criteria using the simplified subset of
// the Apache Lucene query parser syntax defined by the DisMax query
// parser. For more information, see Searching Your Data in the Amazon
// CloudSearch Developer Guide The endpoint for submitting Search requests
// is domain-specific. You submit search requests to a domain's search
// endpoint. To get the search endpoint for your domain, use the Amazon
// CloudSearch configuration service DescribeDomains action. A domain's
// endpoints are also displayed on the domain dashboard in the Amazon
// CloudSearch console.
func (c *CloudSearchDomain) Search(req SearchRequest) (resp *SearchResponse, err error) {
	resp = &SearchResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2013-01-01/search?format=sdk&pretty=true"

	q := url.Values{}

	if s := req.Cursor; s != "" {

		q.Set("cursor", s)
	}

	if s := req.Expr; s != "" {

		q.Set("expr", s)
	}

	if s := req.Facet; s != "" {

		q.Set("facet", s)
	}

	if s := req.FilterQuery; s != "" {

		q.Set("fq", s)
	}

	if s := req.Highlight; s != "" {

		q.Set("highlight", s)
	}

	if s := fmt.Sprintf("%v", req.Partial); s != "" {

		q.Set("partial", s)
	}

	if s := req.Query; s != "" {

		q.Set("q", s)
	}

	if s := req.QueryOptions; s != "" {

		q.Set("q.options", s)
	}

	if s := req.QueryParser; s != "" {

		q.Set("q.parser", s)
	}

	if s := req.Return; s != "" {

		q.Set("return", s)
	}

	if s := fmt.Sprintf("%v", req.Size); s != "" {

		q.Set("size", s)
	}

	if s := req.Sort; s != "" {

		q.Set("sort", s)
	}

	if s := fmt.Sprintf("%v", req.Start); s != "" {

		q.Set("start", s)
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

// Suggest retrieves autocomplete suggestions for a partial query string.
// You can use suggestions enable you to display likely matches before
// users finish typing. In Amazon CloudSearch, suggestions are based on the
// contents of a particular text field. When you request suggestions,
// Amazon CloudSearch finds all of the documents whose values in the
// suggester field start with the specified query string. The beginning of
// the field must match the query string to be considered a match. For more
// information about configuring suggesters and retrieving suggestions, see
// Getting Suggestions in the Amazon CloudSearch Developer Guide . The
// endpoint for submitting Suggest requests is domain-specific. You submit
// suggest requests to a domain's search endpoint. To get the search
// endpoint for your domain, use the Amazon CloudSearch configuration
// service DescribeDomains action. A domain's endpoints are also displayed
// on the domain dashboard in the Amazon CloudSearch console.
func (c *CloudSearchDomain) Suggest(req SuggestRequest) (resp *SuggestResponse, err error) {
	resp = &SuggestResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2013-01-01/suggest?format=sdk&pretty=true"

	q := url.Values{}

	if s := req.Query; s != "" {

		q.Set("q", s)
	}

	if s := fmt.Sprintf("%v", req.Size); s != "" {

		q.Set("size", s)
	}

	if s := req.Suggester; s != "" {

		q.Set("suggester", s)
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

// UploadDocuments posts a batch of documents to a search domain for
// indexing. A document batch is a collection of add and delete operations
// that represent the documents you want to add, update, or delete from
// your domain. Batches can be described in either or Each item that you
// want Amazon CloudSearch to return as a search result (such as a product)
// is represented as a document. Every document has a unique ID and one or
// more fields that contain the data that you want to search and return in
// results. Individual documents cannot contain more than 1 MB of data. The
// entire batch cannot exceed 5 MB. To get the best possible upload
// performance, group add and delete operations in batches that are close
// the 5 MB limit. Submitting a large volume of single-document batches can
// overload a domain's document service. The endpoint for submitting
// UploadDocuments requests is domain-specific. To get the document
// endpoint for your domain, use the Amazon CloudSearch configuration
// service DescribeDomains action. A domain's endpoints are also displayed
// on the domain dashboard in the Amazon CloudSearch console. For more
// information about formatting your data for Amazon CloudSearch, see
// Preparing Your Data in the Amazon CloudSearch Developer Guide . For more
// information about uploading data for indexing, see Uploading Data in the
// Amazon CloudSearch Developer Guide .
func (c *CloudSearchDomain) UploadDocuments(req UploadDocumentsRequest) (resp *UploadDocumentsResponse, err error) {
	resp = &UploadDocumentsResponse{}

	var body io.Reader
	var contentType string

	contentType = "application/json"
	b, err := json.Marshal(req.Documents)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2013-01-01/documents/batch?format=sdk"

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

	if s := req.ContentType; s != "" {

		httpReq.Header.Set("Content-Type", s)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// Bucket is undocumented.
type Bucket struct {
	Count int64  `json:"count,omitempty"`
	Value string `json:"value,omitempty"`
}

// BucketInfo is undocumented.
type BucketInfo struct {
	Buckets []Bucket `json:"buckets,omitempty"`
}

// DocumentServiceWarning is undocumented.
type DocumentServiceWarning struct {
	Message string `json:"message,omitempty"`
}

// Hit is undocumented.
type Hit struct {
	Fields     map[string][]string `json:"fields,omitempty"`
	Highlights map[string]string   `json:"highlights,omitempty"`
	ID         string              `json:"id,omitempty"`
}

// Hits is undocumented.
type Hits struct {
	Cursor string `json:"cursor,omitempty"`
	Found  int64  `json:"found,omitempty"`
	Hit    []Hit  `json:"hit,omitempty"`
	Start  int64  `json:"start,omitempty"`
}

// SearchRequest is undocumented.
type SearchRequest struct {
	Cursor       string `json:"cursor,omitempty"`
	Expr         string `json:"expr,omitempty"`
	Facet        string `json:"facet,omitempty"`
	FilterQuery  string `json:"filterQuery,omitempty"`
	Highlight    string `json:"highlight,omitempty"`
	Partial      bool   `json:"partial,omitempty"`
	Query        string `json:"query"`
	QueryOptions string `json:"queryOptions,omitempty"`
	QueryParser  string `json:"queryParser,omitempty"`
	Return       string `json:"return,omitempty"`
	Size         int64  `json:"size,omitempty"`
	Sort         string `json:"sort,omitempty"`
	Start        int64  `json:"start,omitempty"`
}

// SearchResponse is undocumented.
type SearchResponse struct {
	Facets map[string]BucketInfo `json:"facets,omitempty"`
	Hits   Hits                  `json:"hits,omitempty"`
	Status SearchStatus          `json:"status,omitempty"`
}

// SearchStatus is undocumented.
type SearchStatus struct {
	Rid    string `json:"rid,omitempty"`
	Timems int64  `json:"timems,omitempty"`
}

// SuggestModel is undocumented.
type SuggestModel struct {
	Found       int64             `json:"found,omitempty"`
	Query       string            `json:"query,omitempty"`
	Suggestions []SuggestionMatch `json:"suggestions,omitempty"`
}

// SuggestRequest is undocumented.
type SuggestRequest struct {
	Query     string `json:"query"`
	Size      int64  `json:"size,omitempty"`
	Suggester string `json:"suggester"`
}

// SuggestResponse is undocumented.
type SuggestResponse struct {
	Status  SuggestStatus `json:"status,omitempty"`
	Suggest SuggestModel  `json:"suggest,omitempty"`
}

// SuggestStatus is undocumented.
type SuggestStatus struct {
	Rid    string `json:"rid,omitempty"`
	Timems int64  `json:"timems,omitempty"`
}

// SuggestionMatch is undocumented.
type SuggestionMatch struct {
	ID         string `json:"id,omitempty"`
	Score      int64  `json:"score,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
}

// UploadDocumentsRequest is undocumented.
type UploadDocumentsRequest struct {
	ContentType string `json:"contentType"`
	Documents   []byte `json:"documents"`
}

// UploadDocumentsResponse is undocumented.
type UploadDocumentsResponse struct {
	Adds     int64                    `json:"adds,omitempty"`
	Deletes  int64                    `json:"deletes,omitempty"`
	Status   string                   `json:"status,omitempty"`
	Warnings []DocumentServiceWarning `json:"warnings,omitempty"`
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
var _ json.RawMessage
