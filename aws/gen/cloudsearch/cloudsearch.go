// Package cloudsearch provides a client for Amazon CloudSearch.
package cloudsearch

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// CloudSearch is a client for Amazon CloudSearch.
type CloudSearch struct {
	client *aws.QueryClient
}

// New returns a new CloudSearch client.
func New(key, secret, region string, client *http.Client) *CloudSearch {
	if client == nil {
		client = http.DefaultClient
	}

	service := "cloudsearch"
	endpoint, service, region := endpoints.Lookup("cloudsearch", region)

	return &CloudSearch{
		client: &aws.QueryClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: service,
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2013-01-01",
		},
	}
}

// BuildSuggesters indexes the search suggestions. For more information,
// see Configuring Suggesters in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) BuildSuggesters(req BuildSuggestersRequest) (resp *BuildSuggestersResult, err error) {
	resp = &BuildSuggestersResult{}
	err = c.client.Do("BuildSuggesters", "POST", "/", req, resp)
	return
}

// CreateDomain creates a new search domain. For more information, see
// Creating a Search Domain in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) CreateDomain(req CreateDomainRequest) (resp *CreateDomainResult, err error) {
	resp = &CreateDomainResult{}
	err = c.client.Do("CreateDomain", "POST", "/", req, resp)
	return
}

// DefineAnalysisScheme configures an analysis scheme that can be applied
// to a text or text-array field to define language-specific text
// processing options. For more information, see Configuring Analysis
// Schemes in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DefineAnalysisScheme(req DefineAnalysisSchemeRequest) (resp *DefineAnalysisSchemeResult, err error) {
	resp = &DefineAnalysisSchemeResult{}
	err = c.client.Do("DefineAnalysisScheme", "POST", "/", req, resp)
	return
}

// DefineExpression configures an Expression for the search domain. Used to
// create new expressions and modify existing ones. If the expression
// exists, the new configuration replaces the old one. For more
// information, see Configuring Expressions in the Amazon CloudSearch
// Developer Guide
func (c *CloudSearch) DefineExpression(req DefineExpressionRequest) (resp *DefineExpressionResult, err error) {
	resp = &DefineExpressionResult{}
	err = c.client.Do("DefineExpression", "POST", "/", req, resp)
	return
}

// DefineIndexField configures an IndexField for the search domain. Used to
// create new fields and modify existing ones. You must specify the name of
// the domain you are configuring and an index field configuration. The
// index field configuration specifies a unique name, the index field type,
// and the options you want to configure for the field. The options you can
// specify depend on the IndexFieldType . If the field exists, the new
// configuration replaces the old one. For more information, see
// Configuring Index Fields in the Amazon CloudSearch Developer Guide .
func (c *CloudSearch) DefineIndexField(req DefineIndexFieldRequest) (resp *DefineIndexFieldResult, err error) {
	resp = &DefineIndexFieldResult{}
	err = c.client.Do("DefineIndexField", "POST", "/", req, resp)
	return
}

// DefineSuggester configures a suggester for a domain. A suggester enables
// you to display possible matches before users finish typing their
// queries. When you configure a suggester, you must specify the name of
// the text field you want to search for possible matches and a unique name
// for the suggester. For more information, see Getting Search Suggestions
// in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DefineSuggester(req DefineSuggesterRequest) (resp *DefineSuggesterResult, err error) {
	resp = &DefineSuggesterResult{}
	err = c.client.Do("DefineSuggester", "POST", "/", req, resp)
	return
}

// DeleteAnalysisScheme deletes an analysis scheme. For more information,
// see Configuring Analysis Schemes in the Amazon CloudSearch Developer
// Guide .
func (c *CloudSearch) DeleteAnalysisScheme(req DeleteAnalysisSchemeRequest) (resp *DeleteAnalysisSchemeResult, err error) {
	resp = &DeleteAnalysisSchemeResult{}
	err = c.client.Do("DeleteAnalysisScheme", "POST", "/", req, resp)
	return
}

// DeleteDomain permanently deletes a search domain and all of its data.
// Once a domain has been deleted, it cannot be recovered. For more
// information, see Deleting a Search Domain in the Amazon CloudSearch
// Developer Guide .
func (c *CloudSearch) DeleteDomain(req DeleteDomainRequest) (resp *DeleteDomainResult, err error) {
	resp = &DeleteDomainResult{}
	err = c.client.Do("DeleteDomain", "POST", "/", req, resp)
	return
}

// DeleteExpression removes an Expression from the search domain. For more
// information, see Configuring Expressions in the Amazon CloudSearch
// Developer Guide
func (c *CloudSearch) DeleteExpression(req DeleteExpressionRequest) (resp *DeleteExpressionResult, err error) {
	resp = &DeleteExpressionResult{}
	err = c.client.Do("DeleteExpression", "POST", "/", req, resp)
	return
}

// DeleteIndexField removes an IndexField from the search domain. For more
// information, see Configuring Index Fields in the Amazon CloudSearch
// Developer Guide
func (c *CloudSearch) DeleteIndexField(req DeleteIndexFieldRequest) (resp *DeleteIndexFieldResult, err error) {
	resp = &DeleteIndexFieldResult{}
	err = c.client.Do("DeleteIndexField", "POST", "/", req, resp)
	return
}

// DeleteSuggester deletes a suggester. For more information, see Getting
// Search Suggestions in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DeleteSuggester(req DeleteSuggesterRequest) (resp *DeleteSuggesterResult, err error) {
	resp = &DeleteSuggesterResult{}
	err = c.client.Do("DeleteSuggester", "POST", "/", req, resp)
	return
}

// DescribeAnalysisSchemes gets the analysis schemes configured for a
// domain. An analysis scheme defines language-specific text processing
// options for a text field. Can be limited to specific analysis schemes by
// name. By default, shows all analysis schemes and includes any pending
// changes to the configuration. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more
// information, see Configuring Analysis Schemes in the Amazon CloudSearch
// Developer Guide
func (c *CloudSearch) DescribeAnalysisSchemes(req DescribeAnalysisSchemesRequest) (resp *DescribeAnalysisSchemesResult, err error) {
	resp = &DescribeAnalysisSchemesResult{}
	err = c.client.Do("DescribeAnalysisSchemes", "POST", "/", req, resp)
	return
}

// DescribeAvailabilityOptions gets the availability options configured for
// a domain. By default, shows the configuration with any pending changes.
// Set the Deployed option to true to show the active configuration and
// exclude pending changes. For more information, see Configuring
// Availability Options in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DescribeAvailabilityOptions(req DescribeAvailabilityOptionsRequest) (resp *DescribeAvailabilityOptionsResult, err error) {
	resp = &DescribeAvailabilityOptionsResult{}
	err = c.client.Do("DescribeAvailabilityOptions", "POST", "/", req, resp)
	return
}

// DescribeDomains gets information about the search domains owned by this
// account. Can be limited to specific domains. Shows all domains by
// default. To get the number of searchable documents in a domain, use the
// console or submit a matchall request to your domain's search endpoint:
// q=matchall&q.parser=structured&size=0 . For more information, see
// Getting Information about a Search Domain in the Amazon CloudSearch
// Developer Guide
func (c *CloudSearch) DescribeDomains(req DescribeDomainsRequest) (resp *DescribeDomainsResult, err error) {
	resp = &DescribeDomainsResult{}
	err = c.client.Do("DescribeDomains", "POST", "/", req, resp)
	return
}

// DescribeExpressions gets the expressions configured for the search
// domain. Can be limited to specific expressions by name. By default,
// shows all expressions and includes any pending changes to the
// configuration. Set the Deployed option to true to show the active
// configuration and exclude pending changes. For more information, see
// Configuring Expressions in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DescribeExpressions(req DescribeExpressionsRequest) (resp *DescribeExpressionsResult, err error) {
	resp = &DescribeExpressionsResult{}
	err = c.client.Do("DescribeExpressions", "POST", "/", req, resp)
	return
}

// DescribeIndexFields gets information about the index fields configured
// for the search domain. Can be limited to specific fields by name. By
// default, shows all fields and includes any pending changes to the
// configuration. Set the Deployed option to true to show the active
// configuration and exclude pending changes. For more information, see
// Getting Domain Information in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DescribeIndexFields(req DescribeIndexFieldsRequest) (resp *DescribeIndexFieldsResult, err error) {
	resp = &DescribeIndexFieldsResult{}
	err = c.client.Do("DescribeIndexFields", "POST", "/", req, resp)
	return
}

// DescribeScalingParameters gets the scaling parameters configured for a
// domain. A domain's scaling parameters specify the desired search
// instance type and replication count. For more information, see
// Configuring Scaling Options in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DescribeScalingParameters(req DescribeScalingParametersRequest) (resp *DescribeScalingParametersResult, err error) {
	resp = &DescribeScalingParametersResult{}
	err = c.client.Do("DescribeScalingParameters", "POST", "/", req, resp)
	return
}

// DescribeServiceAccessPolicies gets information about the access policies
// that control access to the domain's document and search endpoints. By
// default, shows the configuration with any pending changes. Set the
// Deployed option to true to show the active configuration and exclude
// pending changes. For more information, see Configuring Access for a
// Search Domain in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DescribeServiceAccessPolicies(req DescribeServiceAccessPoliciesRequest) (resp *DescribeServiceAccessPoliciesResult, err error) {
	resp = &DescribeServiceAccessPoliciesResult{}
	err = c.client.Do("DescribeServiceAccessPolicies", "POST", "/", req, resp)
	return
}

// DescribeSuggesters gets the suggesters configured for a domain. A
// suggester enables you to display possible matches before users finish
// typing their queries. Can be limited to specific suggesters by name. By
// default, shows all suggesters and includes any pending changes to the
// configuration. Set the Deployed option to true to show the active
// configuration and exclude pending changes. For more information, see
// Getting Search Suggestions in the Amazon CloudSearch Developer Guide
func (c *CloudSearch) DescribeSuggesters(req DescribeSuggestersRequest) (resp *DescribeSuggestersResult, err error) {
	resp = &DescribeSuggestersResult{}
	err = c.client.Do("DescribeSuggesters", "POST", "/", req, resp)
	return
}

// IndexDocuments tells the search domain to start indexing its documents
// using the latest indexing options. This operation must be invoked to
// activate options whose OptionStatus is RequiresIndexDocuments
func (c *CloudSearch) IndexDocuments(req IndexDocumentsRequest) (resp *IndexDocumentsResult, err error) {
	resp = &IndexDocumentsResult{}
	err = c.client.Do("IndexDocuments", "POST", "/", req, resp)
	return
}

// ListDomainNames is undocumented.
func (c *CloudSearch) ListDomainNames() (resp *ListDomainNamesResult, err error) {
	resp = &ListDomainNamesResult{}
	err = c.client.Do("ListDomainNames", "POST", "/", nil, resp)
	return
}

// UpdateAvailabilityOptions configures the availability options for a
// domain. Enabling the Multi-AZ option expands an Amazon CloudSearch
// domain to an additional Availability Zone in the same Region to increase
// fault tolerance in the event of a service disruption. Changes to the
// Multi-AZ option can take about half an hour to become active. For more
// information, see Configuring Availability Options in the Amazon
// CloudSearch Developer Guide
func (c *CloudSearch) UpdateAvailabilityOptions(req UpdateAvailabilityOptionsRequest) (resp *UpdateAvailabilityOptionsResult, err error) {
	resp = &UpdateAvailabilityOptionsResult{}
	err = c.client.Do("UpdateAvailabilityOptions", "POST", "/", req, resp)
	return
}

// UpdateScalingParameters configures scaling parameters for a domain. A
// domain's scaling parameters specify the desired search instance type and
// replication count. Amazon CloudSearch will still automatically scale
// your domain based on the volume of data and traffic, but not below the
// desired instance type and replication count. If the Multi-AZ option is
// enabled, these values control the resources used per Availability Zone.
// For more information, see Configuring Scaling Options in the Amazon
// CloudSearch Developer Guide .
func (c *CloudSearch) UpdateScalingParameters(req UpdateScalingParametersRequest) (resp *UpdateScalingParametersResult, err error) {
	resp = &UpdateScalingParametersResult{}
	err = c.client.Do("UpdateScalingParameters", "POST", "/", req, resp)
	return
}

// UpdateServiceAccessPolicies configures the access rules that control
// access to the domain's document and search endpoints. For more
// information, see Configuring Access for an Amazon CloudSearch Domain
func (c *CloudSearch) UpdateServiceAccessPolicies(req UpdateServiceAccessPoliciesRequest) (resp *UpdateServiceAccessPoliciesResult, err error) {
	resp = &UpdateServiceAccessPoliciesResult{}
	err = c.client.Do("UpdateServiceAccessPolicies", "POST", "/", req, resp)
	return
}

// AccessPoliciesStatus is undocumented.
type AccessPoliciesStatus struct {
	Options string       `xml:"Options"`
	Status  OptionStatus `xml:"Status"`
}

// AnalysisOptions is undocumented.
type AnalysisOptions struct {
	AlgorithmicStemming            string `xml:"AlgorithmicStemming"`
	JapaneseTokenizationDictionary string `xml:"JapaneseTokenizationDictionary"`
	StemmingDictionary             string `xml:"StemmingDictionary"`
	Stopwords                      string `xml:"Stopwords"`
	Synonyms                       string `xml:"Synonyms"`
}

// AnalysisScheme is undocumented.
type AnalysisScheme struct {
	AnalysisOptions        AnalysisOptions `xml:"AnalysisOptions"`
	AnalysisSchemeLanguage string          `xml:"AnalysisSchemeLanguage"`
	AnalysisSchemeName     string          `xml:"AnalysisSchemeName"`
}

// AnalysisSchemeStatus is undocumented.
type AnalysisSchemeStatus struct {
	Options AnalysisScheme `xml:"Options"`
	Status  OptionStatus   `xml:"Status"`
}

// AvailabilityOptionsStatus is undocumented.
type AvailabilityOptionsStatus struct {
	Options bool         `xml:"Options"`
	Status  OptionStatus `xml:"Status"`
}

// BuildSuggestersRequest is undocumented.
type BuildSuggestersRequest struct {
	DomainName string `xml:"DomainName"`
}

// BuildSuggestersResponse is undocumented.
type BuildSuggestersResponse struct {
	FieldNames []string `xml:"BuildSuggestersResult>FieldNames>member"`
}

// CreateDomainRequest is undocumented.
type CreateDomainRequest struct {
	DomainName string `xml:"DomainName"`
}

// CreateDomainResponse is undocumented.
type CreateDomainResponse struct {
	DomainStatus DomainStatus `xml:"CreateDomainResult>DomainStatus"`
}

// DateArrayOptions is undocumented.
type DateArrayOptions struct {
	DefaultValue  string `xml:"DefaultValue"`
	FacetEnabled  bool   `xml:"FacetEnabled"`
	ReturnEnabled bool   `xml:"ReturnEnabled"`
	SearchEnabled bool   `xml:"SearchEnabled"`
	SourceFields  string `xml:"SourceFields"`
}

// DateOptions is undocumented.
type DateOptions struct {
	DefaultValue  string `xml:"DefaultValue"`
	FacetEnabled  bool   `xml:"FacetEnabled"`
	ReturnEnabled bool   `xml:"ReturnEnabled"`
	SearchEnabled bool   `xml:"SearchEnabled"`
	SortEnabled   bool   `xml:"SortEnabled"`
	SourceField   string `xml:"SourceField"`
}

// DefineAnalysisSchemeRequest is undocumented.
type DefineAnalysisSchemeRequest struct {
	AnalysisScheme AnalysisScheme `xml:"AnalysisScheme"`
	DomainName     string         `xml:"DomainName"`
}

// DefineAnalysisSchemeResponse is undocumented.
type DefineAnalysisSchemeResponse struct {
	AnalysisScheme AnalysisSchemeStatus `xml:"DefineAnalysisSchemeResult>AnalysisScheme"`
}

// DefineExpressionRequest is undocumented.
type DefineExpressionRequest struct {
	DomainName string     `xml:"DomainName"`
	Expression Expression `xml:"Expression"`
}

// DefineExpressionResponse is undocumented.
type DefineExpressionResponse struct {
	Expression ExpressionStatus `xml:"DefineExpressionResult>Expression"`
}

// DefineIndexFieldRequest is undocumented.
type DefineIndexFieldRequest struct {
	DomainName string     `xml:"DomainName"`
	IndexField IndexField `xml:"IndexField"`
}

// DefineIndexFieldResponse is undocumented.
type DefineIndexFieldResponse struct {
	IndexField IndexFieldStatus `xml:"DefineIndexFieldResult>IndexField"`
}

// DefineSuggesterRequest is undocumented.
type DefineSuggesterRequest struct {
	DomainName string    `xml:"DomainName"`
	Suggester  Suggester `xml:"Suggester"`
}

// DefineSuggesterResponse is undocumented.
type DefineSuggesterResponse struct {
	Suggester SuggesterStatus `xml:"DefineSuggesterResult>Suggester"`
}

// DeleteAnalysisSchemeRequest is undocumented.
type DeleteAnalysisSchemeRequest struct {
	AnalysisSchemeName string `xml:"AnalysisSchemeName"`
	DomainName         string `xml:"DomainName"`
}

// DeleteAnalysisSchemeResponse is undocumented.
type DeleteAnalysisSchemeResponse struct {
	AnalysisScheme AnalysisSchemeStatus `xml:"DeleteAnalysisSchemeResult>AnalysisScheme"`
}

// DeleteDomainRequest is undocumented.
type DeleteDomainRequest struct {
	DomainName string `xml:"DomainName"`
}

// DeleteDomainResponse is undocumented.
type DeleteDomainResponse struct {
	DomainStatus DomainStatus `xml:"DeleteDomainResult>DomainStatus"`
}

// DeleteExpressionRequest is undocumented.
type DeleteExpressionRequest struct {
	DomainName     string `xml:"DomainName"`
	ExpressionName string `xml:"ExpressionName"`
}

// DeleteExpressionResponse is undocumented.
type DeleteExpressionResponse struct {
	Expression ExpressionStatus `xml:"DeleteExpressionResult>Expression"`
}

// DeleteIndexFieldRequest is undocumented.
type DeleteIndexFieldRequest struct {
	DomainName     string `xml:"DomainName"`
	IndexFieldName string `xml:"IndexFieldName"`
}

// DeleteIndexFieldResponse is undocumented.
type DeleteIndexFieldResponse struct {
	IndexField IndexFieldStatus `xml:"DeleteIndexFieldResult>IndexField"`
}

// DeleteSuggesterRequest is undocumented.
type DeleteSuggesterRequest struct {
	DomainName    string `xml:"DomainName"`
	SuggesterName string `xml:"SuggesterName"`
}

// DeleteSuggesterResponse is undocumented.
type DeleteSuggesterResponse struct {
	Suggester SuggesterStatus `xml:"DeleteSuggesterResult>Suggester"`
}

// DescribeAnalysisSchemesRequest is undocumented.
type DescribeAnalysisSchemesRequest struct {
	AnalysisSchemeNames []string `xml:"AnalysisSchemeNames>member"`
	Deployed            bool     `xml:"Deployed"`
	DomainName          string   `xml:"DomainName"`
}

// DescribeAnalysisSchemesResponse is undocumented.
type DescribeAnalysisSchemesResponse struct {
	AnalysisSchemes []AnalysisSchemeStatus `xml:"DescribeAnalysisSchemesResult>AnalysisSchemes>member"`
}

// DescribeAvailabilityOptionsRequest is undocumented.
type DescribeAvailabilityOptionsRequest struct {
	Deployed   bool   `xml:"Deployed"`
	DomainName string `xml:"DomainName"`
}

// DescribeAvailabilityOptionsResponse is undocumented.
type DescribeAvailabilityOptionsResponse struct {
	AvailabilityOptions AvailabilityOptionsStatus `xml:"DescribeAvailabilityOptionsResult>AvailabilityOptions"`
}

// DescribeDomainsRequest is undocumented.
type DescribeDomainsRequest struct {
	DomainNames []string `xml:"DomainNames>member"`
}

// DescribeDomainsResponse is undocumented.
type DescribeDomainsResponse struct {
	DomainStatusList []DomainStatus `xml:"DescribeDomainsResult>DomainStatusList>member"`
}

// DescribeExpressionsRequest is undocumented.
type DescribeExpressionsRequest struct {
	Deployed        bool     `xml:"Deployed"`
	DomainName      string   `xml:"DomainName"`
	ExpressionNames []string `xml:"ExpressionNames>member"`
}

// DescribeExpressionsResponse is undocumented.
type DescribeExpressionsResponse struct {
	Expressions []ExpressionStatus `xml:"DescribeExpressionsResult>Expressions>member"`
}

// DescribeIndexFieldsRequest is undocumented.
type DescribeIndexFieldsRequest struct {
	Deployed   bool     `xml:"Deployed"`
	DomainName string   `xml:"DomainName"`
	FieldNames []string `xml:"FieldNames>member"`
}

// DescribeIndexFieldsResponse is undocumented.
type DescribeIndexFieldsResponse struct {
	IndexFields []IndexFieldStatus `xml:"DescribeIndexFieldsResult>IndexFields>member"`
}

// DescribeScalingParametersRequest is undocumented.
type DescribeScalingParametersRequest struct {
	DomainName string `xml:"DomainName"`
}

// DescribeScalingParametersResponse is undocumented.
type DescribeScalingParametersResponse struct {
	ScalingParameters ScalingParametersStatus `xml:"DescribeScalingParametersResult>ScalingParameters"`
}

// DescribeServiceAccessPoliciesRequest is undocumented.
type DescribeServiceAccessPoliciesRequest struct {
	Deployed   bool   `xml:"Deployed"`
	DomainName string `xml:"DomainName"`
}

// DescribeServiceAccessPoliciesResponse is undocumented.
type DescribeServiceAccessPoliciesResponse struct {
	AccessPolicies AccessPoliciesStatus `xml:"DescribeServiceAccessPoliciesResult>AccessPolicies"`
}

// DescribeSuggestersRequest is undocumented.
type DescribeSuggestersRequest struct {
	Deployed       bool     `xml:"Deployed"`
	DomainName     string   `xml:"DomainName"`
	SuggesterNames []string `xml:"SuggesterNames>member"`
}

// DescribeSuggestersResponse is undocumented.
type DescribeSuggestersResponse struct {
	Suggesters []SuggesterStatus `xml:"DescribeSuggestersResult>Suggesters>member"`
}

// DocumentSuggesterOptions is undocumented.
type DocumentSuggesterOptions struct {
	FuzzyMatching  string `xml:"FuzzyMatching"`
	SortExpression string `xml:"SortExpression"`
	SourceField    string `xml:"SourceField"`
}

// DomainStatus is undocumented.
type DomainStatus struct {
	ARN                    string          `xml:"ARN"`
	Created                bool            `xml:"Created"`
	Deleted                bool            `xml:"Deleted"`
	DocService             ServiceEndpoint `xml:"DocService"`
	DomainID               string          `xml:"DomainId"`
	DomainName             string          `xml:"DomainName"`
	Limits                 Limits          `xml:"Limits"`
	Processing             bool            `xml:"Processing"`
	RequiresIndexDocuments bool            `xml:"RequiresIndexDocuments"`
	SearchInstanceCount    int             `xml:"SearchInstanceCount"`
	SearchInstanceType     string          `xml:"SearchInstanceType"`
	SearchPartitionCount   int             `xml:"SearchPartitionCount"`
	SearchService          ServiceEndpoint `xml:"SearchService"`
}

// DoubleArrayOptions is undocumented.
type DoubleArrayOptions struct {
	DefaultValue  float64 `xml:"DefaultValue"`
	FacetEnabled  bool    `xml:"FacetEnabled"`
	ReturnEnabled bool    `xml:"ReturnEnabled"`
	SearchEnabled bool    `xml:"SearchEnabled"`
	SourceFields  string  `xml:"SourceFields"`
}

// DoubleOptions is undocumented.
type DoubleOptions struct {
	DefaultValue  float64 `xml:"DefaultValue"`
	FacetEnabled  bool    `xml:"FacetEnabled"`
	ReturnEnabled bool    `xml:"ReturnEnabled"`
	SearchEnabled bool    `xml:"SearchEnabled"`
	SortEnabled   bool    `xml:"SortEnabled"`
	SourceField   string  `xml:"SourceField"`
}

// Expression is undocumented.
type Expression struct {
	ExpressionName  string `xml:"ExpressionName"`
	ExpressionValue string `xml:"ExpressionValue"`
}

// ExpressionStatus is undocumented.
type ExpressionStatus struct {
	Options Expression   `xml:"Options"`
	Status  OptionStatus `xml:"Status"`
}

// IndexDocumentsRequest is undocumented.
type IndexDocumentsRequest struct {
	DomainName string `xml:"DomainName"`
}

// IndexDocumentsResponse is undocumented.
type IndexDocumentsResponse struct {
	FieldNames []string `xml:"IndexDocumentsResult>FieldNames>member"`
}

// IndexField is undocumented.
type IndexField struct {
	DateArrayOptions    DateArrayOptions    `xml:"DateArrayOptions"`
	DateOptions         DateOptions         `xml:"DateOptions"`
	DoubleArrayOptions  DoubleArrayOptions  `xml:"DoubleArrayOptions"`
	DoubleOptions       DoubleOptions       `xml:"DoubleOptions"`
	IndexFieldName      string              `xml:"IndexFieldName"`
	IndexFieldType      string              `xml:"IndexFieldType"`
	IntArrayOptions     IntArrayOptions     `xml:"IntArrayOptions"`
	IntOptions          IntOptions          `xml:"IntOptions"`
	LatLonOptions       LatLonOptions       `xml:"LatLonOptions"`
	LiteralArrayOptions LiteralArrayOptions `xml:"LiteralArrayOptions"`
	LiteralOptions      LiteralOptions      `xml:"LiteralOptions"`
	TextArrayOptions    TextArrayOptions    `xml:"TextArrayOptions"`
	TextOptions         TextOptions         `xml:"TextOptions"`
}

// IndexFieldStatus is undocumented.
type IndexFieldStatus struct {
	Options IndexField   `xml:"Options"`
	Status  OptionStatus `xml:"Status"`
}

// IntArrayOptions is undocumented.
type IntArrayOptions struct {
	DefaultValue  int64  `xml:"DefaultValue"`
	FacetEnabled  bool   `xml:"FacetEnabled"`
	ReturnEnabled bool   `xml:"ReturnEnabled"`
	SearchEnabled bool   `xml:"SearchEnabled"`
	SourceFields  string `xml:"SourceFields"`
}

// IntOptions is undocumented.
type IntOptions struct {
	DefaultValue  int64  `xml:"DefaultValue"`
	FacetEnabled  bool   `xml:"FacetEnabled"`
	ReturnEnabled bool   `xml:"ReturnEnabled"`
	SearchEnabled bool   `xml:"SearchEnabled"`
	SortEnabled   bool   `xml:"SortEnabled"`
	SourceField   string `xml:"SourceField"`
}

// LatLonOptions is undocumented.
type LatLonOptions struct {
	DefaultValue  string `xml:"DefaultValue"`
	FacetEnabled  bool   `xml:"FacetEnabled"`
	ReturnEnabled bool   `xml:"ReturnEnabled"`
	SearchEnabled bool   `xml:"SearchEnabled"`
	SortEnabled   bool   `xml:"SortEnabled"`
	SourceField   string `xml:"SourceField"`
}

// Limits is undocumented.
type Limits struct {
	MaximumPartitionCount   int `xml:"MaximumPartitionCount"`
	MaximumReplicationCount int `xml:"MaximumReplicationCount"`
}

// ListDomainNamesResponse is undocumented.
type ListDomainNamesResponse struct {
	DomainNames map[string]string `xml:"ListDomainNamesResult>DomainNames"`
}

// LiteralArrayOptions is undocumented.
type LiteralArrayOptions struct {
	DefaultValue  string `xml:"DefaultValue"`
	FacetEnabled  bool   `xml:"FacetEnabled"`
	ReturnEnabled bool   `xml:"ReturnEnabled"`
	SearchEnabled bool   `xml:"SearchEnabled"`
	SourceFields  string `xml:"SourceFields"`
}

// LiteralOptions is undocumented.
type LiteralOptions struct {
	DefaultValue  string `xml:"DefaultValue"`
	FacetEnabled  bool   `xml:"FacetEnabled"`
	ReturnEnabled bool   `xml:"ReturnEnabled"`
	SearchEnabled bool   `xml:"SearchEnabled"`
	SortEnabled   bool   `xml:"SortEnabled"`
	SourceField   string `xml:"SourceField"`
}

// OptionStatus is undocumented.
type OptionStatus struct {
	CreationDate    time.Time `xml:"CreationDate"`
	PendingDeletion bool      `xml:"PendingDeletion"`
	State           string    `xml:"State"`
	UpdateDate      time.Time `xml:"UpdateDate"`
	UpdateVersion   int       `xml:"UpdateVersion"`
}

// ScalingParameters is undocumented.
type ScalingParameters struct {
	DesiredInstanceType     string `xml:"DesiredInstanceType"`
	DesiredPartitionCount   int    `xml:"DesiredPartitionCount"`
	DesiredReplicationCount int    `xml:"DesiredReplicationCount"`
}

// ScalingParametersStatus is undocumented.
type ScalingParametersStatus struct {
	Options ScalingParameters `xml:"Options"`
	Status  OptionStatus      `xml:"Status"`
}

// ServiceEndpoint is undocumented.
type ServiceEndpoint struct {
	Endpoint string `xml:"Endpoint"`
}

// Suggester is undocumented.
type Suggester struct {
	DocumentSuggesterOptions DocumentSuggesterOptions `xml:"DocumentSuggesterOptions"`
	SuggesterName            string                   `xml:"SuggesterName"`
}

// SuggesterStatus is undocumented.
type SuggesterStatus struct {
	Options Suggester    `xml:"Options"`
	Status  OptionStatus `xml:"Status"`
}

// TextArrayOptions is undocumented.
type TextArrayOptions struct {
	AnalysisScheme   string `xml:"AnalysisScheme"`
	DefaultValue     string `xml:"DefaultValue"`
	HighlightEnabled bool   `xml:"HighlightEnabled"`
	ReturnEnabled    bool   `xml:"ReturnEnabled"`
	SourceFields     string `xml:"SourceFields"`
}

// TextOptions is undocumented.
type TextOptions struct {
	AnalysisScheme   string `xml:"AnalysisScheme"`
	DefaultValue     string `xml:"DefaultValue"`
	HighlightEnabled bool   `xml:"HighlightEnabled"`
	ReturnEnabled    bool   `xml:"ReturnEnabled"`
	SortEnabled      bool   `xml:"SortEnabled"`
	SourceField      string `xml:"SourceField"`
}

// UpdateAvailabilityOptionsRequest is undocumented.
type UpdateAvailabilityOptionsRequest struct {
	DomainName string `xml:"DomainName"`
	MultiAZ    bool   `xml:"MultiAZ"`
}

// UpdateAvailabilityOptionsResponse is undocumented.
type UpdateAvailabilityOptionsResponse struct {
	AvailabilityOptions AvailabilityOptionsStatus `xml:"UpdateAvailabilityOptionsResult>AvailabilityOptions"`
}

// UpdateScalingParametersRequest is undocumented.
type UpdateScalingParametersRequest struct {
	DomainName        string            `xml:"DomainName"`
	ScalingParameters ScalingParameters `xml:"ScalingParameters"`
}

// UpdateScalingParametersResponse is undocumented.
type UpdateScalingParametersResponse struct {
	ScalingParameters ScalingParametersStatus `xml:"UpdateScalingParametersResult>ScalingParameters"`
}

// UpdateServiceAccessPoliciesRequest is undocumented.
type UpdateServiceAccessPoliciesRequest struct {
	AccessPolicies string `xml:"AccessPolicies"`
	DomainName     string `xml:"DomainName"`
}

// UpdateServiceAccessPoliciesResponse is undocumented.
type UpdateServiceAccessPoliciesResponse struct {
	AccessPolicies AccessPoliciesStatus `xml:"UpdateServiceAccessPoliciesResult>AccessPolicies"`
}

// BuildSuggestersResult is a wrapper for BuildSuggestersResponse.
type BuildSuggestersResult struct {
	XMLName xml.Name `xml:"BuildSuggestersResponse"`

	FieldNames []string `xml:"BuildSuggestersResult>FieldNames>member"`
}

// CreateDomainResult is a wrapper for CreateDomainResponse.
type CreateDomainResult struct {
	XMLName xml.Name `xml:"CreateDomainResponse"`

	DomainStatus DomainStatus `xml:"CreateDomainResult>DomainStatus"`
}

// DefineAnalysisSchemeResult is a wrapper for DefineAnalysisSchemeResponse.
type DefineAnalysisSchemeResult struct {
	XMLName xml.Name `xml:"DefineAnalysisSchemeResponse"`

	AnalysisScheme AnalysisSchemeStatus `xml:"DefineAnalysisSchemeResult>AnalysisScheme"`
}

// DefineExpressionResult is a wrapper for DefineExpressionResponse.
type DefineExpressionResult struct {
	XMLName xml.Name `xml:"DefineExpressionResponse"`

	Expression ExpressionStatus `xml:"DefineExpressionResult>Expression"`
}

// DefineIndexFieldResult is a wrapper for DefineIndexFieldResponse.
type DefineIndexFieldResult struct {
	XMLName xml.Name `xml:"DefineIndexFieldResponse"`

	IndexField IndexFieldStatus `xml:"DefineIndexFieldResult>IndexField"`
}

// DefineSuggesterResult is a wrapper for DefineSuggesterResponse.
type DefineSuggesterResult struct {
	XMLName xml.Name `xml:"DefineSuggesterResponse"`

	Suggester SuggesterStatus `xml:"DefineSuggesterResult>Suggester"`
}

// DeleteAnalysisSchemeResult is a wrapper for DeleteAnalysisSchemeResponse.
type DeleteAnalysisSchemeResult struct {
	XMLName xml.Name `xml:"DeleteAnalysisSchemeResponse"`

	AnalysisScheme AnalysisSchemeStatus `xml:"DeleteAnalysisSchemeResult>AnalysisScheme"`
}

// DeleteDomainResult is a wrapper for DeleteDomainResponse.
type DeleteDomainResult struct {
	XMLName xml.Name `xml:"DeleteDomainResponse"`

	DomainStatus DomainStatus `xml:"DeleteDomainResult>DomainStatus"`
}

// DeleteExpressionResult is a wrapper for DeleteExpressionResponse.
type DeleteExpressionResult struct {
	XMLName xml.Name `xml:"DeleteExpressionResponse"`

	Expression ExpressionStatus `xml:"DeleteExpressionResult>Expression"`
}

// DeleteIndexFieldResult is a wrapper for DeleteIndexFieldResponse.
type DeleteIndexFieldResult struct {
	XMLName xml.Name `xml:"DeleteIndexFieldResponse"`

	IndexField IndexFieldStatus `xml:"DeleteIndexFieldResult>IndexField"`
}

// DeleteSuggesterResult is a wrapper for DeleteSuggesterResponse.
type DeleteSuggesterResult struct {
	XMLName xml.Name `xml:"DeleteSuggesterResponse"`

	Suggester SuggesterStatus `xml:"DeleteSuggesterResult>Suggester"`
}

// DescribeAnalysisSchemesResult is a wrapper for DescribeAnalysisSchemesResponse.
type DescribeAnalysisSchemesResult struct {
	XMLName xml.Name `xml:"DescribeAnalysisSchemesResponse"`

	AnalysisSchemes []AnalysisSchemeStatus `xml:"DescribeAnalysisSchemesResult>AnalysisSchemes>member"`
}

// DescribeAvailabilityOptionsResult is a wrapper for DescribeAvailabilityOptionsResponse.
type DescribeAvailabilityOptionsResult struct {
	XMLName xml.Name `xml:"DescribeAvailabilityOptionsResponse"`

	AvailabilityOptions AvailabilityOptionsStatus `xml:"DescribeAvailabilityOptionsResult>AvailabilityOptions"`
}

// DescribeDomainsResult is a wrapper for DescribeDomainsResponse.
type DescribeDomainsResult struct {
	XMLName xml.Name `xml:"DescribeDomainsResponse"`

	DomainStatusList []DomainStatus `xml:"DescribeDomainsResult>DomainStatusList>member"`
}

// DescribeExpressionsResult is a wrapper for DescribeExpressionsResponse.
type DescribeExpressionsResult struct {
	XMLName xml.Name `xml:"DescribeExpressionsResponse"`

	Expressions []ExpressionStatus `xml:"DescribeExpressionsResult>Expressions>member"`
}

// DescribeIndexFieldsResult is a wrapper for DescribeIndexFieldsResponse.
type DescribeIndexFieldsResult struct {
	XMLName xml.Name `xml:"DescribeIndexFieldsResponse"`

	IndexFields []IndexFieldStatus `xml:"DescribeIndexFieldsResult>IndexFields>member"`
}

// DescribeScalingParametersResult is a wrapper for DescribeScalingParametersResponse.
type DescribeScalingParametersResult struct {
	XMLName xml.Name `xml:"DescribeScalingParametersResponse"`

	ScalingParameters ScalingParametersStatus `xml:"DescribeScalingParametersResult>ScalingParameters"`
}

// DescribeServiceAccessPoliciesResult is a wrapper for DescribeServiceAccessPoliciesResponse.
type DescribeServiceAccessPoliciesResult struct {
	XMLName xml.Name `xml:"DescribeServiceAccessPoliciesResponse"`

	AccessPolicies AccessPoliciesStatus `xml:"DescribeServiceAccessPoliciesResult>AccessPolicies"`
}

// DescribeSuggestersResult is a wrapper for DescribeSuggestersResponse.
type DescribeSuggestersResult struct {
	XMLName xml.Name `xml:"DescribeSuggestersResponse"`

	Suggesters []SuggesterStatus `xml:"DescribeSuggestersResult>Suggesters>member"`
}

// IndexDocumentsResult is a wrapper for IndexDocumentsResponse.
type IndexDocumentsResult struct {
	XMLName xml.Name `xml:"IndexDocumentsResponse"`

	FieldNames []string `xml:"IndexDocumentsResult>FieldNames>member"`
}

// ListDomainNamesResult is a wrapper for ListDomainNamesResponse.
type ListDomainNamesResult struct {
	XMLName xml.Name `xml:"ListDomainNamesResponse"`

	DomainNames map[string]string `xml:"ListDomainNamesResult>DomainNames"`
}

// UpdateAvailabilityOptionsResult is a wrapper for UpdateAvailabilityOptionsResponse.
type UpdateAvailabilityOptionsResult struct {
	XMLName xml.Name `xml:"UpdateAvailabilityOptionsResponse"`

	AvailabilityOptions AvailabilityOptionsStatus `xml:"UpdateAvailabilityOptionsResult>AvailabilityOptions"`
}

// UpdateScalingParametersResult is a wrapper for UpdateScalingParametersResponse.
type UpdateScalingParametersResult struct {
	XMLName xml.Name `xml:"UpdateScalingParametersResponse"`

	ScalingParameters ScalingParametersStatus `xml:"UpdateScalingParametersResult>ScalingParameters"`
}

// UpdateServiceAccessPoliciesResult is a wrapper for UpdateServiceAccessPoliciesResponse.
type UpdateServiceAccessPoliciesResult struct {
	XMLName xml.Name `xml:"UpdateServiceAccessPoliciesResponse"`

	AccessPolicies AccessPoliciesStatus `xml:"UpdateServiceAccessPoliciesResult>AccessPolicies"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
