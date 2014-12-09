// Package importexport provides a client for AWS Import/Export.
package importexport

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// ImportExport is a client for AWS Import/Export.
type ImportExport struct {
	client aws.Client
}

// New returns a new ImportExport client.
func New(key, secret, region string, client *http.Client) *ImportExport {
	if client == nil {
		client = http.DefaultClient
	}

	return &ImportExport{
		client: &aws.QueryClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "importexport",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoints.Lookup("importexport", region),
			APIVersion: "2010-06-01",
		},
	}
}

// CancelJob this operation cancels a specified job. Only the job owner can
// cancel it. The operation fails if the job has already started or is
// complete.
func (c *ImportExport) CancelJob(req CancelJobInput) (resp *CancelJobResult, err error) {
	resp = &CancelJobResult{}
	err = c.client.Do("CancelJob", "POST", "/?Operation=CancelJob", req, resp)
	return
}

// CreateJob this operation initiates the process of scheduling an upload
// or download of your data. You include in the request a manifest that
// describes the data transfer specifics. The response to the request
// includes a job ID, which you can use in other operations, a signature
// that you use to identify your storage device, and the address where you
// should ship your storage device.
func (c *ImportExport) CreateJob(req CreateJobInput) (resp *CreateJobResult, err error) {
	resp = &CreateJobResult{}
	err = c.client.Do("CreateJob", "POST", "/?Operation=CreateJob", req, resp)
	return
}

// GetStatus this operation returns information about a job, including
// where the job is in the processing pipeline, the status of the results,
// and the signature value associated with the job. You can only return
// information about jobs you own.
func (c *ImportExport) GetStatus(req GetStatusInput) (resp *GetStatusResult, err error) {
	resp = &GetStatusResult{}
	err = c.client.Do("GetStatus", "POST", "/?Operation=GetStatus", req, resp)
	return
}

// ListJobs this operation returns the jobs associated with the requester.
// AWS Import/Export lists the jobs in reverse chronological order based on
// the date of creation. For example if Job Test1 was created 2009Dec30 and
// Test2 was created 2010Feb05, the ListJobs operation would return Test2
// followed by Test1.
func (c *ImportExport) ListJobs(req ListJobsInput) (resp *ListJobsResult, err error) {
	resp = &ListJobsResult{}
	err = c.client.Do("ListJobs", "POST", "/?Operation=ListJobs", req, resp)
	return
}

// UpdateJob you use this operation to change the parameters specified in
// the original manifest file by supplying a new manifest file. The
// manifest file attached to this request replaces the original manifest
// file. You can only use the operation after a CreateJob request but
// before the data transfer starts and you can only use it on jobs you own.
func (c *ImportExport) UpdateJob(req UpdateJobInput) (resp *UpdateJobResult, err error) {
	resp = &UpdateJobResult{}
	err = c.client.Do("UpdateJob", "POST", "/?Operation=UpdateJob", req, resp)
	return
}

// CancelJobInput is undocumented.
type CancelJobInput struct {
	JobID string `xml:"JobId"`
}

// CancelJobOutput is undocumented.
type CancelJobOutput struct {
	Success bool `xml:"CancelJobResult>Success"`
}

// CreateJobInput is undocumented.
type CreateJobInput struct {
	JobType          string `xml:"JobType"`
	Manifest         string `xml:"Manifest"`
	ManifestAddendum string `xml:"ManifestAddendum"`
	ValidateOnly     bool   `xml:"ValidateOnly"`
}

// CreateJobOutput is undocumented.
type CreateJobOutput struct {
	AwsShippingAddress    string `xml:"CreateJobResult>AwsShippingAddress"`
	JobID                 string `xml:"CreateJobResult>JobId"`
	JobType               string `xml:"CreateJobResult>JobType"`
	Signature             string `xml:"CreateJobResult>Signature"`
	SignatureFileContents string `xml:"CreateJobResult>SignatureFileContents"`
	WarningMessage        string `xml:"CreateJobResult>WarningMessage"`
}

// GetStatusInput is undocumented.
type GetStatusInput struct {
	JobID string `xml:"JobId"`
}

// GetStatusOutput is undocumented.
type GetStatusOutput struct {
	AwsShippingAddress    string    `xml:"GetStatusResult>AwsShippingAddress"`
	Carrier               string    `xml:"GetStatusResult>Carrier"`
	CreationDate          time.Time `xml:"GetStatusResult>CreationDate"`
	CurrentManifest       string    `xml:"GetStatusResult>CurrentManifest"`
	ErrorCount            int       `xml:"GetStatusResult>ErrorCount"`
	JobID                 string    `xml:"GetStatusResult>JobId"`
	JobType               string    `xml:"GetStatusResult>JobType"`
	LocationCode          string    `xml:"GetStatusResult>LocationCode"`
	LocationMessage       string    `xml:"GetStatusResult>LocationMessage"`
	LogBucket             string    `xml:"GetStatusResult>LogBucket"`
	LogKey                string    `xml:"GetStatusResult>LogKey"`
	ProgressCode          string    `xml:"GetStatusResult>ProgressCode"`
	ProgressMessage       string    `xml:"GetStatusResult>ProgressMessage"`
	Signature             string    `xml:"GetStatusResult>Signature"`
	SignatureFileContents string    `xml:"GetStatusResult>SignatureFileContents"`
	TrackingNumber        string    `xml:"GetStatusResult>TrackingNumber"`
}

// Job is undocumented.
type Job struct {
	CreationDate time.Time `xml:"CreationDate"`
	IsCanceled   bool      `xml:"IsCanceled"`
	JobID        string    `xml:"JobId"`
	JobType      string    `xml:"JobType"`
}

// ListJobsInput is undocumented.
type ListJobsInput struct {
	Marker  string `xml:"Marker"`
	MaxJobs int    `xml:"MaxJobs"`
}

// ListJobsOutput is undocumented.
type ListJobsOutput struct {
	IsTruncated bool  `xml:"ListJobsResult>IsTruncated"`
	Jobs        []Job `xml:"ListJobsResult>Jobs>member"`
}

// UpdateJobInput is undocumented.
type UpdateJobInput struct {
	JobID        string `xml:"JobId"`
	JobType      string `xml:"JobType"`
	Manifest     string `xml:"Manifest"`
	ValidateOnly bool   `xml:"ValidateOnly"`
}

// UpdateJobOutput is undocumented.
type UpdateJobOutput struct {
	Success        bool   `xml:"UpdateJobResult>Success"`
	WarningMessage string `xml:"UpdateJobResult>WarningMessage"`
}

// CancelJobResult is a wrapper for CancelJobOutput.
type CancelJobResult struct {
	XMLName xml.Name `xml:"CancelJobResponse"`

	Success bool `xml:"CancelJobResult>Success"`
}

// CreateJobResult is a wrapper for CreateJobOutput.
type CreateJobResult struct {
	XMLName xml.Name `xml:"CreateJobResponse"`

	AwsShippingAddress    string `xml:"CreateJobResult>AwsShippingAddress"`
	JobID                 string `xml:"CreateJobResult>JobId"`
	JobType               string `xml:"CreateJobResult>JobType"`
	Signature             string `xml:"CreateJobResult>Signature"`
	SignatureFileContents string `xml:"CreateJobResult>SignatureFileContents"`
	WarningMessage        string `xml:"CreateJobResult>WarningMessage"`
}

// GetStatusResult is a wrapper for GetStatusOutput.
type GetStatusResult struct {
	XMLName xml.Name `xml:"GetStatusResponse"`

	AwsShippingAddress    string    `xml:"GetStatusResult>AwsShippingAddress"`
	Carrier               string    `xml:"GetStatusResult>Carrier"`
	CreationDate          time.Time `xml:"GetStatusResult>CreationDate"`
	CurrentManifest       string    `xml:"GetStatusResult>CurrentManifest"`
	ErrorCount            int       `xml:"GetStatusResult>ErrorCount"`
	JobID                 string    `xml:"GetStatusResult>JobId"`
	JobType               string    `xml:"GetStatusResult>JobType"`
	LocationCode          string    `xml:"GetStatusResult>LocationCode"`
	LocationMessage       string    `xml:"GetStatusResult>LocationMessage"`
	LogBucket             string    `xml:"GetStatusResult>LogBucket"`
	LogKey                string    `xml:"GetStatusResult>LogKey"`
	ProgressCode          string    `xml:"GetStatusResult>ProgressCode"`
	ProgressMessage       string    `xml:"GetStatusResult>ProgressMessage"`
	Signature             string    `xml:"GetStatusResult>Signature"`
	SignatureFileContents string    `xml:"GetStatusResult>SignatureFileContents"`
	TrackingNumber        string    `xml:"GetStatusResult>TrackingNumber"`
}

// ListJobsResult is a wrapper for ListJobsOutput.
type ListJobsResult struct {
	XMLName xml.Name `xml:"ListJobsResponse"`

	IsTruncated bool  `xml:"ListJobsResult>IsTruncated"`
	Jobs        []Job `xml:"ListJobsResult>Jobs>member"`
}

// UpdateJobResult is a wrapper for UpdateJobOutput.
type UpdateJobResult struct {
	XMLName xml.Name `xml:"UpdateJobResponse"`

	Success        bool   `xml:"UpdateJobResult>Success"`
	WarningMessage string `xml:"UpdateJobResult>WarningMessage"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
