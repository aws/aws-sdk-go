// Package elastictranscoder provides a client for Amazon Elastic Transcoder.
package elastictranscoder

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

// ElasticTranscoder is a client for Amazon Elastic Transcoder.
type ElasticTranscoder struct {
	client *aws.RestClient
}

// New returns a new ElasticTranscoder client.
func New(key, secret, region string, client *http.Client) *ElasticTranscoder {
	if client == nil {
		client = http.DefaultClient
	}

	service := "elastictranscoder"
	endpoint, service, region := endpoints.Lookup("elastictranscoder", region)

	return &ElasticTranscoder{
		client: &aws.RestClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: service,
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2012-09-25",
		},
	}
}

// CancelJob the CancelJob operation cancels an unfinished job. You can
// only cancel a job that has a status of Submitted . To prevent a pipeline
// from starting to process a job while you're getting the job identifier,
// use UpdatePipelineStatus to temporarily pause the pipeline.
func (c *ElasticTranscoder) CancelJob(req CancelJobRequest) (resp *CancelJobResponse, err error) {
	resp = &CancelJobResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobs/{Id}"

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

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// CreateJob when you create a job, Elastic Transcoder returns data that
// includes the values that you specified plus information about the job
// that is created. If you have specified more than one output for your
// jobs (for example, one output for the Kindle Fire and another output for
// the Apple iPhone 4s), you currently must use the Elastic Transcoder API
// to list the jobs (as opposed to the AWS Console).
func (c *ElasticTranscoder) CreateJob(req CreateJobRequest) (resp *CreateJobResponse, err error) {
	resp = &CreateJobResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobs"

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

	return
}

// CreatePipeline the CreatePipeline operation creates a pipeline with
// settings that you specify.
func (c *ElasticTranscoder) CreatePipeline(req CreatePipelineRequest) (resp *CreatePipelineResponse, err error) {
	resp = &CreatePipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines"

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

	return
}

// CreatePreset the CreatePreset operation creates a preset with settings
// that you specify. Elastic Transcoder checks the CreatePreset settings to
// ensure that they meet Elastic Transcoder requirements and to determine
// whether they comply with H.264 standards. If your settings are not valid
// for Elastic Transcoder, Elastic Transcoder returns an 400 response
// ValidationException ) and does not create the preset. If the settings
// are valid for Elastic Transcoder but aren't strictly compliant with the
// H.264 standard, Elastic Transcoder creates the preset and returns a
// warning message in the response. This helps you determine whether your
// settings comply with the H.264 standard while giving you greater
// flexibility with respect to the video that Elastic Transcoder produces.
// Elastic Transcoder uses the H.264 video-compression format. For more
// information, see the International Telecommunication Union publication
// Recommendation H.264: Advanced video coding for generic audiovisual
// services
func (c *ElasticTranscoder) CreatePreset(req CreatePresetRequest) (resp *CreatePresetResponse, err error) {
	resp = &CreatePresetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/presets"

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

	return
}

// DeletePipeline the DeletePipeline operation removes a pipeline. You can
// only delete a pipeline that has never been used or that is not currently
// in use (doesn't contain any active jobs). If the pipeline is currently
// in use, DeletePipeline returns an error.
func (c *ElasticTranscoder) DeletePipeline(req DeletePipelineRequest) (resp *DeletePipelineResponse, err error) {
	resp = &DeletePipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}"

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

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// DeletePreset the DeletePreset operation removes a preset that you've
// added in an AWS region. You can't delete the default presets that are
// included with Elastic Transcoder.
func (c *ElasticTranscoder) DeletePreset(req DeletePresetRequest) (resp *DeletePresetResponse, err error) {
	resp = &DeletePresetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/presets/{Id}"

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

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// ListJobsByPipeline the ListJobsByPipeline operation gets a list of the
// jobs currently in a pipeline. Elastic Transcoder returns all of the jobs
// currently in the specified pipeline. The response body contains one
// element for each job that satisfies the search criteria.
func (c *ElasticTranscoder) ListJobsByPipeline(req ListJobsByPipelineRequest) (resp *ListJobsByPipelineResponse, err error) {
	resp = &ListJobsByPipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobsByPipeline/{PipelineId}"

	uri = strings.Replace(uri, "{"+"PipelineId"+"}", req.PipelineID, -1)
	uri = strings.Replace(uri, "{"+"PipelineId+"+"}", req.PipelineID, -1)

	q := url.Values{}

	if s := req.Ascending; s != "" {

		q.Set("Ascending", s)
	}

	if s := req.PageToken; s != "" {

		q.Set("PageToken", s)
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

// ListJobsByStatus the ListJobsByStatus operation gets a list of jobs that
// have a specified status. The response body contains one element for each
// job that satisfies the search criteria.
func (c *ElasticTranscoder) ListJobsByStatus(req ListJobsByStatusRequest) (resp *ListJobsByStatusResponse, err error) {
	resp = &ListJobsByStatusResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobsByStatus/{Status}"

	uri = strings.Replace(uri, "{"+"Status"+"}", req.Status, -1)
	uri = strings.Replace(uri, "{"+"Status+"+"}", req.Status, -1)

	q := url.Values{}

	if s := req.Ascending; s != "" {

		q.Set("Ascending", s)
	}

	if s := req.PageToken; s != "" {

		q.Set("PageToken", s)
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

// ListPipelines the ListPipelines operation gets a list of the pipelines
// associated with the current AWS account.
func (c *ElasticTranscoder) ListPipelines(req ListPipelinesRequest) (resp *ListPipelinesResponse, err error) {
	resp = &ListPipelinesResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines"

	q := url.Values{}

	if s := req.Ascending; s != "" {

		q.Set("Ascending", s)
	}

	if s := req.PageToken; s != "" {

		q.Set("PageToken", s)
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

// ListPresets the ListPresets operation gets a list of the default presets
// included with Elastic Transcoder and the presets that you've added in an
// AWS region.
func (c *ElasticTranscoder) ListPresets(req ListPresetsRequest) (resp *ListPresetsResponse, err error) {
	resp = &ListPresetsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/presets"

	q := url.Values{}

	if s := req.Ascending; s != "" {

		q.Set("Ascending", s)
	}

	if s := req.PageToken; s != "" {

		q.Set("PageToken", s)
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

// ReadJob the ReadJob operation returns detailed information about a job.
func (c *ElasticTranscoder) ReadJob(req ReadJobRequest) (resp *ReadJobResponse, err error) {
	resp = &ReadJobResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobs/{Id}"

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

// ReadPipeline the ReadPipeline operation gets detailed information about
// a pipeline.
func (c *ElasticTranscoder) ReadPipeline(req ReadPipelineRequest) (resp *ReadPipelineResponse, err error) {
	resp = &ReadPipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}"

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

// ReadPreset the ReadPreset operation gets detailed information about a
// preset.
func (c *ElasticTranscoder) ReadPreset(req ReadPresetRequest) (resp *ReadPresetResponse, err error) {
	resp = &ReadPresetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/presets/{Id}"

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

// TestRole the TestRole operation tests the IAM role used to create the
// pipeline. The TestRole action lets you determine whether the IAM role
// you are using has sufficient permissions to let Elastic Transcoder
// perform tasks associated with the transcoding process. The action
// attempts to assume the specified IAM role, checks read access to the
// input and output buckets, and tries to send a test notification to
// Amazon SNS topics that you specify.
func (c *ElasticTranscoder) TestRole(req TestRoleRequest) (resp *TestRoleResponse, err error) {
	resp = &TestRoleResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/roleTests"

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

	return
}

// UpdatePipeline use the UpdatePipeline operation to update settings for a
// pipeline. When you change pipeline settings, your changes take effect
// immediately. Jobs that you have already submitted and that Elastic
// Transcoder has not started to process are affected in addition to jobs
// that you submit after you change settings.
func (c *ElasticTranscoder) UpdatePipeline(req UpdatePipelineRequest) (resp *UpdatePipelineResponse, err error) {
	resp = &UpdatePipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}"

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

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// UpdatePipelineNotifications with the UpdatePipelineNotifications
// operation, you can update Amazon Simple Notification Service (Amazon
// notifications for a pipeline. When you update notifications for a
// pipeline, Elastic Transcoder returns the values that you specified in
// the request.
func (c *ElasticTranscoder) UpdatePipelineNotifications(req UpdatePipelineNotificationsRequest) (resp *UpdatePipelineNotificationsResponse, err error) {
	resp = &UpdatePipelineNotificationsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}/notifications"

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

// UpdatePipelineStatus the UpdatePipelineStatus operation pauses or
// reactivates a pipeline, so that the pipeline stops or restarts the
// processing of jobs. Changing the pipeline status is useful if you want
// to cancel one or more jobs. You can't cancel jobs after Elastic
// Transcoder has started processing them; if you pause the pipeline to
// which you submitted the jobs, you have more time to get the job IDs for
// the jobs that you want to cancel, and to send a CancelJob request.
func (c *ElasticTranscoder) UpdatePipelineStatus(req UpdatePipelineStatusRequest) (resp *UpdatePipelineStatusResponse, err error) {
	resp = &UpdatePipelineStatusResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}/status"

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

// Artwork is undocumented.
type Artwork struct {
	AlbumArtFormat string     `json:"AlbumArtFormat,omitempty"`
	Encryption     Encryption `json:"Encryption,omitempty"`
	InputKey       string     `json:"InputKey,omitempty"`
	MaxHeight      string     `json:"MaxHeight,omitempty"`
	MaxWidth       string     `json:"MaxWidth,omitempty"`
	PaddingPolicy  string     `json:"PaddingPolicy,omitempty"`
	SizingPolicy   string     `json:"SizingPolicy,omitempty"`
}

// AudioCodecOptions is undocumented.
type AudioCodecOptions struct {
	Profile string `json:"Profile,omitempty"`
}

// AudioParameters is undocumented.
type AudioParameters struct {
	BitRate      string            `json:"BitRate,omitempty"`
	Channels     string            `json:"Channels,omitempty"`
	Codec        string            `json:"Codec,omitempty"`
	CodecOptions AudioCodecOptions `json:"CodecOptions,omitempty"`
	SampleRate   string            `json:"SampleRate,omitempty"`
}

// CancelJobRequest is undocumented.
type CancelJobRequest struct {
	ID string `json:"Id"`
}

// CancelJobResponse is undocumented.
type CancelJobResponse struct {
}

// CaptionFormat is undocumented.
type CaptionFormat struct {
	Encryption Encryption `json:"Encryption,omitempty"`
	Format     string     `json:"Format,omitempty"`
	Pattern    string     `json:"Pattern,omitempty"`
}

// CaptionSource is undocumented.
type CaptionSource struct {
	Encryption Encryption `json:"Encryption,omitempty"`
	Key        string     `json:"Key,omitempty"`
	Label      string     `json:"Label,omitempty"`
	Language   string     `json:"Language,omitempty"`
	TimeOffset string     `json:"TimeOffset,omitempty"`
}

// Captions is undocumented.
type Captions struct {
	CaptionFormats []CaptionFormat `json:"CaptionFormats,omitempty"`
	CaptionSources []CaptionSource `json:"CaptionSources,omitempty"`
	MergePolicy    string          `json:"MergePolicy,omitempty"`
}

// Clip is undocumented.
type Clip struct {
	TimeSpan TimeSpan `json:"TimeSpan,omitempty"`
}

// CreateJobOutput is undocumented.
type CreateJobOutput struct {
	AlbumArt            JobAlbumArt    `json:"AlbumArt,omitempty"`
	Captions            Captions       `json:"Captions,omitempty"`
	Composition         []Clip         `json:"Composition,omitempty"`
	Encryption          Encryption     `json:"Encryption,omitempty"`
	Key                 string         `json:"Key,omitempty"`
	PresetID            string         `json:"PresetId,omitempty"`
	Rotate              string         `json:"Rotate,omitempty"`
	SegmentDuration     string         `json:"SegmentDuration,omitempty"`
	ThumbnailEncryption Encryption     `json:"ThumbnailEncryption,omitempty"`
	ThumbnailPattern    string         `json:"ThumbnailPattern,omitempty"`
	Watermarks          []JobWatermark `json:"Watermarks,omitempty"`
}

// CreateJobPlaylist is undocumented.
type CreateJobPlaylist struct {
	Format     string   `json:"Format,omitempty"`
	Name       string   `json:"Name,omitempty"`
	OutputKeys []string `json:"OutputKeys,omitempty"`
}

// CreateJobRequest is undocumented.
type CreateJobRequest struct {
	Input           JobInput            `json:"Input"`
	Output          CreateJobOutput     `json:"Output,omitempty"`
	OutputKeyPrefix string              `json:"OutputKeyPrefix,omitempty"`
	Outputs         []CreateJobOutput   `json:"Outputs,omitempty"`
	PipelineID      string              `json:"PipelineId"`
	Playlists       []CreateJobPlaylist `json:"Playlists,omitempty"`
}

// CreateJobResponse is undocumented.
type CreateJobResponse struct {
	Job Job `json:"Job,omitempty"`
}

// CreatePipelineRequest is undocumented.
type CreatePipelineRequest struct {
	AwsKmsKeyARN    string               `json:"AwsKmsKeyArn,omitempty"`
	ContentConfig   PipelineOutputConfig `json:"ContentConfig,omitempty"`
	InputBucket     string               `json:"InputBucket"`
	Name            string               `json:"Name"`
	Notifications   Notifications        `json:"Notifications,omitempty"`
	OutputBucket    string               `json:"OutputBucket,omitempty"`
	Role            string               `json:"Role"`
	ThumbnailConfig PipelineOutputConfig `json:"ThumbnailConfig,omitempty"`
}

// CreatePipelineResponse is undocumented.
type CreatePipelineResponse struct {
	Pipeline Pipeline `json:"Pipeline,omitempty"`
}

// CreatePresetRequest is undocumented.
type CreatePresetRequest struct {
	Audio       AudioParameters `json:"Audio,omitempty"`
	Container   string          `json:"Container"`
	Description string          `json:"Description,omitempty"`
	Name        string          `json:"Name"`
	Thumbnails  Thumbnails      `json:"Thumbnails,omitempty"`
	Video       VideoParameters `json:"Video,omitempty"`
}

// CreatePresetResponse is undocumented.
type CreatePresetResponse struct {
	Preset  Preset `json:"Preset,omitempty"`
	Warning string `json:"Warning,omitempty"`
}

// DeletePipelineRequest is undocumented.
type DeletePipelineRequest struct {
	ID string `json:"Id"`
}

// DeletePipelineResponse is undocumented.
type DeletePipelineResponse struct {
}

// DeletePresetRequest is undocumented.
type DeletePresetRequest struct {
	ID string `json:"Id"`
}

// DeletePresetResponse is undocumented.
type DeletePresetResponse struct {
}

// Encryption is undocumented.
type Encryption struct {
	InitializationVector string `json:"InitializationVector,omitempty"`
	Key                  string `json:"Key,omitempty"`
	KeyMd5               string `json:"KeyMd5,omitempty"`
	Mode                 string `json:"Mode,omitempty"`
}

// Job is undocumented.
type Job struct {
	ARN             string      `json:"Arn,omitempty"`
	ID              string      `json:"Id,omitempty"`
	Input           JobInput    `json:"Input,omitempty"`
	Output          JobOutput   `json:"Output,omitempty"`
	OutputKeyPrefix string      `json:"OutputKeyPrefix,omitempty"`
	Outputs         []JobOutput `json:"Outputs,omitempty"`
	PipelineID      string      `json:"PipelineId,omitempty"`
	Playlists       []Playlist  `json:"Playlists,omitempty"`
	Status          string      `json:"Status,omitempty"`
}

// JobAlbumArt is undocumented.
type JobAlbumArt struct {
	Artwork     []Artwork `json:"Artwork,omitempty"`
	MergePolicy string    `json:"MergePolicy,omitempty"`
}

// JobInput is undocumented.
type JobInput struct {
	AspectRatio string     `json:"AspectRatio,omitempty"`
	Container   string     `json:"Container,omitempty"`
	Encryption  Encryption `json:"Encryption,omitempty"`
	FrameRate   string     `json:"FrameRate,omitempty"`
	Interlaced  string     `json:"Interlaced,omitempty"`
	Key         string     `json:"Key,omitempty"`
	Resolution  string     `json:"Resolution,omitempty"`
}

// JobOutput is undocumented.
type JobOutput struct {
	AlbumArt            JobAlbumArt    `json:"AlbumArt,omitempty"`
	Captions            Captions       `json:"Captions,omitempty"`
	Composition         []Clip         `json:"Composition,omitempty"`
	Duration            int64          `json:"Duration,omitempty"`
	Encryption          Encryption     `json:"Encryption,omitempty"`
	Height              int            `json:"Height,omitempty"`
	ID                  string         `json:"Id,omitempty"`
	Key                 string         `json:"Key,omitempty"`
	PresetID            string         `json:"PresetId,omitempty"`
	Rotate              string         `json:"Rotate,omitempty"`
	SegmentDuration     string         `json:"SegmentDuration,omitempty"`
	Status              string         `json:"Status,omitempty"`
	StatusDetail        string         `json:"StatusDetail,omitempty"`
	ThumbnailEncryption Encryption     `json:"ThumbnailEncryption,omitempty"`
	ThumbnailPattern    string         `json:"ThumbnailPattern,omitempty"`
	Watermarks          []JobWatermark `json:"Watermarks,omitempty"`
	Width               int            `json:"Width,omitempty"`
}

// JobWatermark is undocumented.
type JobWatermark struct {
	Encryption        Encryption `json:"Encryption,omitempty"`
	InputKey          string     `json:"InputKey,omitempty"`
	PresetWatermarkID string     `json:"PresetWatermarkId,omitempty"`
}

// ListJobsByPipelineRequest is undocumented.
type ListJobsByPipelineRequest struct {
	Ascending  string `json:"Ascending,omitempty"`
	PageToken  string `json:"PageToken,omitempty"`
	PipelineID string `json:"PipelineId"`
}

// ListJobsByPipelineResponse is undocumented.
type ListJobsByPipelineResponse struct {
	Jobs          []Job  `json:"Jobs,omitempty"`
	NextPageToken string `json:"NextPageToken,omitempty"`
}

// ListJobsByStatusRequest is undocumented.
type ListJobsByStatusRequest struct {
	Ascending string `json:"Ascending,omitempty"`
	PageToken string `json:"PageToken,omitempty"`
	Status    string `json:"Status"`
}

// ListJobsByStatusResponse is undocumented.
type ListJobsByStatusResponse struct {
	Jobs          []Job  `json:"Jobs,omitempty"`
	NextPageToken string `json:"NextPageToken,omitempty"`
}

// ListPipelinesRequest is undocumented.
type ListPipelinesRequest struct {
	Ascending string `json:"Ascending,omitempty"`
	PageToken string `json:"PageToken,omitempty"`
}

// ListPipelinesResponse is undocumented.
type ListPipelinesResponse struct {
	NextPageToken string     `json:"NextPageToken,omitempty"`
	Pipelines     []Pipeline `json:"Pipelines,omitempty"`
}

// ListPresetsRequest is undocumented.
type ListPresetsRequest struct {
	Ascending string `json:"Ascending,omitempty"`
	PageToken string `json:"PageToken,omitempty"`
}

// ListPresetsResponse is undocumented.
type ListPresetsResponse struct {
	NextPageToken string   `json:"NextPageToken,omitempty"`
	Presets       []Preset `json:"Presets,omitempty"`
}

// Notifications is undocumented.
type Notifications struct {
	Completed   string `json:"Completed,omitempty"`
	Error       string `json:"Error,omitempty"`
	Progressing string `json:"Progressing,omitempty"`
	Warning     string `json:"Warning,omitempty"`
}

// Permission is undocumented.
type Permission struct {
	Access      []string `json:"Access,omitempty"`
	Grantee     string   `json:"Grantee,omitempty"`
	GranteeType string   `json:"GranteeType,omitempty"`
}

// Pipeline is undocumented.
type Pipeline struct {
	ARN             string               `json:"Arn,omitempty"`
	AwsKmsKeyARN    string               `json:"AwsKmsKeyArn,omitempty"`
	ContentConfig   PipelineOutputConfig `json:"ContentConfig,omitempty"`
	ID              string               `json:"Id,omitempty"`
	InputBucket     string               `json:"InputBucket,omitempty"`
	Name            string               `json:"Name,omitempty"`
	Notifications   Notifications        `json:"Notifications,omitempty"`
	OutputBucket    string               `json:"OutputBucket,omitempty"`
	Role            string               `json:"Role,omitempty"`
	Status          string               `json:"Status,omitempty"`
	ThumbnailConfig PipelineOutputConfig `json:"ThumbnailConfig,omitempty"`
}

// PipelineOutputConfig is undocumented.
type PipelineOutputConfig struct {
	Bucket       string       `json:"Bucket,omitempty"`
	Permissions  []Permission `json:"Permissions,omitempty"`
	StorageClass string       `json:"StorageClass,omitempty"`
}

// Playlist is undocumented.
type Playlist struct {
	Format       string   `json:"Format,omitempty"`
	Name         string   `json:"Name,omitempty"`
	OutputKeys   []string `json:"OutputKeys,omitempty"`
	Status       string   `json:"Status,omitempty"`
	StatusDetail string   `json:"StatusDetail,omitempty"`
}

// Preset is undocumented.
type Preset struct {
	ARN         string          `json:"Arn,omitempty"`
	Audio       AudioParameters `json:"Audio,omitempty"`
	Container   string          `json:"Container,omitempty"`
	Description string          `json:"Description,omitempty"`
	ID          string          `json:"Id,omitempty"`
	Name        string          `json:"Name,omitempty"`
	Thumbnails  Thumbnails      `json:"Thumbnails,omitempty"`
	Type        string          `json:"Type,omitempty"`
	Video       VideoParameters `json:"Video,omitempty"`
}

// PresetWatermark is undocumented.
type PresetWatermark struct {
	HorizontalAlign  string `json:"HorizontalAlign,omitempty"`
	HorizontalOffset string `json:"HorizontalOffset,omitempty"`
	ID               string `json:"Id,omitempty"`
	MaxHeight        string `json:"MaxHeight,omitempty"`
	MaxWidth         string `json:"MaxWidth,omitempty"`
	Opacity          string `json:"Opacity,omitempty"`
	SizingPolicy     string `json:"SizingPolicy,omitempty"`
	Target           string `json:"Target,omitempty"`
	VerticalAlign    string `json:"VerticalAlign,omitempty"`
	VerticalOffset   string `json:"VerticalOffset,omitempty"`
}

// ReadJobRequest is undocumented.
type ReadJobRequest struct {
	ID string `json:"Id"`
}

// ReadJobResponse is undocumented.
type ReadJobResponse struct {
	Job Job `json:"Job,omitempty"`
}

// ReadPipelineRequest is undocumented.
type ReadPipelineRequest struct {
	ID string `json:"Id"`
}

// ReadPipelineResponse is undocumented.
type ReadPipelineResponse struct {
	Pipeline Pipeline `json:"Pipeline,omitempty"`
}

// ReadPresetRequest is undocumented.
type ReadPresetRequest struct {
	ID string `json:"Id"`
}

// ReadPresetResponse is undocumented.
type ReadPresetResponse struct {
	Preset Preset `json:"Preset,omitempty"`
}

// TestRoleRequest is undocumented.
type TestRoleRequest struct {
	InputBucket  string   `json:"InputBucket"`
	OutputBucket string   `json:"OutputBucket"`
	Role         string   `json:"Role"`
	Topics       []string `json:"Topics"`
}

// TestRoleResponse is undocumented.
type TestRoleResponse struct {
	Messages []string `json:"Messages,omitempty"`
	Success  string   `json:"Success,omitempty"`
}

// Thumbnails is undocumented.
type Thumbnails struct {
	AspectRatio   string `json:"AspectRatio,omitempty"`
	Format        string `json:"Format,omitempty"`
	Interval      string `json:"Interval,omitempty"`
	MaxHeight     string `json:"MaxHeight,omitempty"`
	MaxWidth      string `json:"MaxWidth,omitempty"`
	PaddingPolicy string `json:"PaddingPolicy,omitempty"`
	Resolution    string `json:"Resolution,omitempty"`
	SizingPolicy  string `json:"SizingPolicy,omitempty"`
}

// TimeSpan is undocumented.
type TimeSpan struct {
	Duration  string `json:"Duration,omitempty"`
	StartTime string `json:"StartTime,omitempty"`
}

// UpdatePipelineNotificationsRequest is undocumented.
type UpdatePipelineNotificationsRequest struct {
	ID            string        `json:"Id"`
	Notifications Notifications `json:"Notifications"`
}

// UpdatePipelineNotificationsResponse is undocumented.
type UpdatePipelineNotificationsResponse struct {
	Pipeline Pipeline `json:"Pipeline,omitempty"`
}

// UpdatePipelineRequest is undocumented.
type UpdatePipelineRequest struct {
	AwsKmsKeyARN    string               `json:"AwsKmsKeyArn,omitempty"`
	ContentConfig   PipelineOutputConfig `json:"ContentConfig,omitempty"`
	ID              string               `json:"Id"`
	InputBucket     string               `json:"InputBucket,omitempty"`
	Name            string               `json:"Name,omitempty"`
	Notifications   Notifications        `json:"Notifications,omitempty"`
	Role            string               `json:"Role,omitempty"`
	ThumbnailConfig PipelineOutputConfig `json:"ThumbnailConfig,omitempty"`
}

// UpdatePipelineResponse is undocumented.
type UpdatePipelineResponse struct {
	Pipeline Pipeline `json:"Pipeline,omitempty"`
}

// UpdatePipelineStatusRequest is undocumented.
type UpdatePipelineStatusRequest struct {
	ID     string `json:"Id"`
	Status string `json:"Status"`
}

// UpdatePipelineStatusResponse is undocumented.
type UpdatePipelineStatusResponse struct {
	Pipeline Pipeline `json:"Pipeline,omitempty"`
}

// VideoParameters is undocumented.
type VideoParameters struct {
	AspectRatio        string            `json:"AspectRatio,omitempty"`
	BitRate            string            `json:"BitRate,omitempty"`
	Codec              string            `json:"Codec,omitempty"`
	CodecOptions       map[string]string `json:"CodecOptions,omitempty"`
	DisplayAspectRatio string            `json:"DisplayAspectRatio,omitempty"`
	FixedGOP           string            `json:"FixedGOP,omitempty"`
	FrameRate          string            `json:"FrameRate,omitempty"`
	KeyframesMaxDist   string            `json:"KeyframesMaxDist,omitempty"`
	MaxFrameRate       string            `json:"MaxFrameRate,omitempty"`
	MaxHeight          string            `json:"MaxHeight,omitempty"`
	MaxWidth           string            `json:"MaxWidth,omitempty"`
	PaddingPolicy      string            `json:"PaddingPolicy,omitempty"`
	Resolution         string            `json:"Resolution,omitempty"`
	SizingPolicy       string            `json:"SizingPolicy,omitempty"`
	Watermarks         []PresetWatermark `json:"Watermarks,omitempty"`
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
