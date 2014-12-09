// Package cloudwatch provides a client for Amazon CloudWatch.
package cloudwatch

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/bmizerany/aws4"
	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// CloudWatch is a client for Amazon CloudWatch.
type CloudWatch struct {
	client aws.Client
}

// New returns a new CloudWatch client.
func New(key, secret, region string, client *http.Client) *CloudWatch {
	if client == nil {
		client = http.DefaultClient
	}

	return &CloudWatch{
		client: &aws.QueryClient{
			Client: &aws4.Client{
				Keys: &aws4.Keys{
					AccessKey: key,
					SecretKey: secret,
				},
				Client: client,
			},
			Endpoint:   endpoints.Lookup("monitoring", region),
			APIVersion: "2010-08-01",
		},
	}
}

// DeleteAlarms deletes all specified alarms. In the event of an error, no
// alarms are deleted.
func (c *CloudWatch) DeleteAlarms(req DeleteAlarmsInput) (err error) {
	// NRE
	err = c.client.Do("DeleteAlarms", "POST", "/", req, nil)
	return
}

// DescribeAlarmHistory retrieves history for the specified alarm. Filter
// alarms by date range or item type. If an alarm name is not specified,
// Amazon CloudWatch returns histories for all of the owner's alarms.
func (c *CloudWatch) DescribeAlarmHistory(req DescribeAlarmHistoryInput) (resp *DescribeAlarmHistoryResult, err error) {
	resp = &DescribeAlarmHistoryResult{}
	err = c.client.Do("DescribeAlarmHistory", "POST", "/", req, resp)
	return
}

// DescribeAlarms retrieves alarms with the specified names. If no name is
// specified, all alarms for the user are returned. Alarms can be retrieved
// by using only a prefix for the alarm name, the alarm state, or a prefix
// for any action.
func (c *CloudWatch) DescribeAlarms(req DescribeAlarmsInput) (resp *DescribeAlarmsResult, err error) {
	resp = &DescribeAlarmsResult{}
	err = c.client.Do("DescribeAlarms", "POST", "/", req, resp)
	return
}

// DescribeAlarmsForMetric retrieves all alarms for a single metric.
// Specify a statistic, period, or unit to filter the set of alarms
// further.
func (c *CloudWatch) DescribeAlarmsForMetric(req DescribeAlarmsForMetricInput) (resp *DescribeAlarmsForMetricResult, err error) {
	resp = &DescribeAlarmsForMetricResult{}
	err = c.client.Do("DescribeAlarmsForMetric", "POST", "/", req, resp)
	return
}

// DisableAlarmActions disables actions for the specified alarms. When an
// alarm's actions are disabled the alarm's state may change, but none of
// the alarm's actions will execute.
func (c *CloudWatch) DisableAlarmActions(req DisableAlarmActionsInput) (err error) {
	// NRE
	err = c.client.Do("DisableAlarmActions", "POST", "/", req, nil)
	return
}

// EnableAlarmActions is undocumented.
func (c *CloudWatch) EnableAlarmActions(req EnableAlarmActionsInput) (err error) {
	// NRE
	err = c.client.Do("EnableAlarmActions", "POST", "/", req, nil)
	return
}

// GetMetricStatistics gets statistics for the specified metric. The
// maximum number of data points returned from a single GetMetricStatistics
// request is 1,440, wereas the maximum number of data points that can be
// queried is 50,850. If you make a request that generates more than 1,440
// data points, Amazon CloudWatch returns an error. In such a case, you can
// alter the request by narrowing the specified time range or increasing
// the specified period. Alternatively, you can make multiple requests
// across adjacent time ranges. Amazon CloudWatch aggregates data points
// based on the length of the period that you specify. For example, if you
// request statistics with a one-minute granularity, Amazon CloudWatch
// aggregates data points with time stamps that fall within the same
// one-minute period. In such a case, the data points queried can greatly
// outnumber the data points returned. The following examples show various
// statistics allowed by the data point query maximum of 50,850 when you
// call GetMetricStatistics on Amazon EC2 instances with detailed
// (one-minute) monitoring enabled: Statistics for up to 400 instances for
// a span of one hour Statistics for up to 35 instances over a span of 24
// hours Statistics for up to 2 instances over a span of 2 weeks For
// information about the namespace, metric names, and dimensions that other
// Amazon Web Services products use to send metrics to Cloudwatch, go to
// Amazon CloudWatch Metrics, Namespaces, and Dimensions Reference in the
// Amazon CloudWatch Developer Guide .
func (c *CloudWatch) GetMetricStatistics(req GetMetricStatisticsInput) (resp *GetMetricStatisticsResult, err error) {
	resp = &GetMetricStatisticsResult{}
	err = c.client.Do("GetMetricStatistics", "POST", "/", req, resp)
	return
}

// ListMetrics returns a list of valid metrics stored for the AWS account
// owner. Returned metrics can be used with GetMetricStatistics to obtain
// statistical data for a given metric.
func (c *CloudWatch) ListMetrics(req ListMetricsInput) (resp *ListMetricsResult, err error) {
	resp = &ListMetricsResult{}
	err = c.client.Do("ListMetrics", "POST", "/", req, resp)
	return
}

// PutMetricAlarm creates or updates an alarm and associates it with the
// specified Amazon CloudWatch metric. Optionally, this operation can
// associate one or more Amazon Simple Notification Service resources with
// the alarm. When this operation creates an alarm, the alarm state is
// immediately set to . The alarm is evaluated and its StateValue is set
// appropriately. Any actions associated with the StateValue is then
// executed.
func (c *CloudWatch) PutMetricAlarm(req PutMetricAlarmInput) (err error) {
	// NRE
	err = c.client.Do("PutMetricAlarm", "POST", "/", req, nil)
	return
}

// PutMetricData publishes metric data points to Amazon CloudWatch. Amazon
// Cloudwatch associates the data points with the specified metric. If the
// specified metric does not exist, Amazon CloudWatch creates the metric.
// It can take up to fifteen minutes for a new metric to appear in calls to
// the ListMetrics action. request is limited to 8 KB for GET requests and
// 40 KB for requests. Although the Value parameter accepts numbers of type
// Double , Amazon CloudWatch truncates values with very large exponents.
// Values with base-10 exponents greater than 126 (1 x 10^126) are
// truncated. Likewise, values with base-10 exponents less than -130 (1 x
// 10^-130) are also truncated. Data that is timestamped 24 hours or more
// in the past may take in excess of 48 hours to become available from
// submission time using GetMetricStatistics
func (c *CloudWatch) PutMetricData(req PutMetricDataInput) (err error) {
	// NRE
	err = c.client.Do("PutMetricData", "POST", "/", req, nil)
	return
}

// SetAlarmState temporarily sets the state of an alarm. When the updated
// StateValue differs from the previous value, the action configured for
// the appropriate state is invoked. This is not a permanent change. The
// next periodic alarm check (in about a minute) will set the alarm to its
// actual state.
func (c *CloudWatch) SetAlarmState(req SetAlarmStateInput) (err error) {
	// NRE
	err = c.client.Do("SetAlarmState", "POST", "/", req, nil)
	return
}

// AlarmHistoryItem is undocumented.
type AlarmHistoryItem struct {
	AlarmName       string    `xml:"AlarmName"`
	HistoryData     string    `xml:"HistoryData"`
	HistoryItemType string    `xml:"HistoryItemType"`
	HistorySummary  string    `xml:"HistorySummary"`
	Timestamp       time.Time `xml:"Timestamp"`
}

// Datapoint is undocumented.
type Datapoint struct {
	Average     float64   `xml:"Average"`
	Maximum     float64   `xml:"Maximum"`
	Minimum     float64   `xml:"Minimum"`
	SampleCount float64   `xml:"SampleCount"`
	Sum         float64   `xml:"Sum"`
	Timestamp   time.Time `xml:"Timestamp"`
	Unit        string    `xml:"Unit"`
}

// DeleteAlarmsInput is undocumented.
type DeleteAlarmsInput struct {
	AlarmNames []string `xml:"AlarmNames>member"`
}

// DescribeAlarmHistoryInput is undocumented.
type DescribeAlarmHistoryInput struct {
	AlarmName       string    `xml:"AlarmName"`
	EndDate         time.Time `xml:"EndDate"`
	HistoryItemType string    `xml:"HistoryItemType"`
	MaxRecords      int       `xml:"MaxRecords"`
	NextToken       string    `xml:"NextToken"`
	StartDate       time.Time `xml:"StartDate"`
}

// DescribeAlarmHistoryOutput is undocumented.
type DescribeAlarmHistoryOutput struct {
	AlarmHistoryItems []AlarmHistoryItem `xml:"DescribeAlarmHistoryResult>AlarmHistoryItems>member"`
	NextToken         string             `xml:"DescribeAlarmHistoryResult>NextToken"`
}

// DescribeAlarmsForMetricInput is undocumented.
type DescribeAlarmsForMetricInput struct {
	Dimensions []Dimension `xml:"Dimensions>member"`
	MetricName string      `xml:"MetricName"`
	Namespace  string      `xml:"Namespace"`
	Period     int         `xml:"Period"`
	Statistic  string      `xml:"Statistic"`
	Unit       string      `xml:"Unit"`
}

// DescribeAlarmsForMetricOutput is undocumented.
type DescribeAlarmsForMetricOutput struct {
	MetricAlarms []MetricAlarm `xml:"DescribeAlarmsForMetricResult>MetricAlarms>member"`
}

// DescribeAlarmsInput is undocumented.
type DescribeAlarmsInput struct {
	ActionPrefix    string   `xml:"ActionPrefix"`
	AlarmNamePrefix string   `xml:"AlarmNamePrefix"`
	AlarmNames      []string `xml:"AlarmNames>member"`
	MaxRecords      int      `xml:"MaxRecords"`
	NextToken       string   `xml:"NextToken"`
	StateValue      string   `xml:"StateValue"`
}

// DescribeAlarmsOutput is undocumented.
type DescribeAlarmsOutput struct {
	MetricAlarms []MetricAlarm `xml:"DescribeAlarmsResult>MetricAlarms>member"`
	NextToken    string        `xml:"DescribeAlarmsResult>NextToken"`
}

// Dimension is undocumented.
type Dimension struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

// DimensionFilter is undocumented.
type DimensionFilter struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

// DisableAlarmActionsInput is undocumented.
type DisableAlarmActionsInput struct {
	AlarmNames []string `xml:"AlarmNames>member"`
}

// EnableAlarmActionsInput is undocumented.
type EnableAlarmActionsInput struct {
	AlarmNames []string `xml:"AlarmNames>member"`
}

// GetMetricStatisticsInput is undocumented.
type GetMetricStatisticsInput struct {
	Dimensions []Dimension `xml:"Dimensions>member"`
	EndTime    time.Time   `xml:"EndTime"`
	MetricName string      `xml:"MetricName"`
	Namespace  string      `xml:"Namespace"`
	Period     int         `xml:"Period"`
	StartTime  time.Time   `xml:"StartTime"`
	Statistics []string    `xml:"Statistics>member"`
	Unit       string      `xml:"Unit"`
}

// GetMetricStatisticsOutput is undocumented.
type GetMetricStatisticsOutput struct {
	Datapoints []Datapoint `xml:"GetMetricStatisticsResult>Datapoints>member"`
	Label      string      `xml:"GetMetricStatisticsResult>Label"`
}

// ListMetricsInput is undocumented.
type ListMetricsInput struct {
	Dimensions []DimensionFilter `xml:"Dimensions>member"`
	MetricName string            `xml:"MetricName"`
	Namespace  string            `xml:"Namespace"`
	NextToken  string            `xml:"NextToken"`
}

// ListMetricsOutput is undocumented.
type ListMetricsOutput struct {
	Metrics   []Metric `xml:"ListMetricsResult>Metrics>member"`
	NextToken string   `xml:"ListMetricsResult>NextToken"`
}

// Metric is undocumented.
type Metric struct {
	Dimensions []Dimension `xml:"Dimensions>member"`
	MetricName string      `xml:"MetricName"`
	Namespace  string      `xml:"Namespace"`
}

// MetricAlarm is undocumented.
type MetricAlarm struct {
	ActionsEnabled                     bool        `xml:"ActionsEnabled"`
	AlarmActions                       []string    `xml:"AlarmActions>member"`
	AlarmARN                           string      `xml:"AlarmArn"`
	AlarmConfigurationUpdatedTimestamp time.Time   `xml:"AlarmConfigurationUpdatedTimestamp"`
	AlarmDescription                   string      `xml:"AlarmDescription"`
	AlarmName                          string      `xml:"AlarmName"`
	ComparisonOperator                 string      `xml:"ComparisonOperator"`
	Dimensions                         []Dimension `xml:"Dimensions>member"`
	EvaluationPeriods                  int         `xml:"EvaluationPeriods"`
	InsufficientDataActions            []string    `xml:"InsufficientDataActions>member"`
	MetricName                         string      `xml:"MetricName"`
	Namespace                          string      `xml:"Namespace"`
	OKActions                          []string    `xml:"OKActions>member"`
	Period                             int         `xml:"Period"`
	StateReason                        string      `xml:"StateReason"`
	StateReasonData                    string      `xml:"StateReasonData"`
	StateUpdatedTimestamp              time.Time   `xml:"StateUpdatedTimestamp"`
	StateValue                         string      `xml:"StateValue"`
	Statistic                          string      `xml:"Statistic"`
	Threshold                          float64     `xml:"Threshold"`
	Unit                               string      `xml:"Unit"`
}

// MetricDatum is undocumented.
type MetricDatum struct {
	Dimensions      []Dimension  `xml:"Dimensions>member"`
	MetricName      string       `xml:"MetricName"`
	StatisticValues StatisticSet `xml:"StatisticValues"`
	Timestamp       time.Time    `xml:"Timestamp"`
	Unit            string       `xml:"Unit"`
	Value           float64      `xml:"Value"`
}

// PutMetricAlarmInput is undocumented.
type PutMetricAlarmInput struct {
	ActionsEnabled          bool        `xml:"ActionsEnabled"`
	AlarmActions            []string    `xml:"AlarmActions>member"`
	AlarmDescription        string      `xml:"AlarmDescription"`
	AlarmName               string      `xml:"AlarmName"`
	ComparisonOperator      string      `xml:"ComparisonOperator"`
	Dimensions              []Dimension `xml:"Dimensions>member"`
	EvaluationPeriods       int         `xml:"EvaluationPeriods"`
	InsufficientDataActions []string    `xml:"InsufficientDataActions>member"`
	MetricName              string      `xml:"MetricName"`
	Namespace               string      `xml:"Namespace"`
	OKActions               []string    `xml:"OKActions>member"`
	Period                  int         `xml:"Period"`
	Statistic               string      `xml:"Statistic"`
	Threshold               float64     `xml:"Threshold"`
	Unit                    string      `xml:"Unit"`
}

// PutMetricDataInput is undocumented.
type PutMetricDataInput struct {
	MetricData []MetricDatum `xml:"MetricData>member"`
	Namespace  string        `xml:"Namespace"`
}

// SetAlarmStateInput is undocumented.
type SetAlarmStateInput struct {
	AlarmName       string `xml:"AlarmName"`
	StateReason     string `xml:"StateReason"`
	StateReasonData string `xml:"StateReasonData"`
	StateValue      string `xml:"StateValue"`
}

// StatisticSet is undocumented.
type StatisticSet struct {
	Maximum     float64 `xml:"Maximum"`
	Minimum     float64 `xml:"Minimum"`
	SampleCount float64 `xml:"SampleCount"`
	Sum         float64 `xml:"Sum"`
}

// DescribeAlarmHistoryResult is a wrapper for DescribeAlarmHistoryOutput.
type DescribeAlarmHistoryResult struct {
	XMLName xml.Name `xml:"DescribeAlarmHistoryResponse"`

	AlarmHistoryItems []AlarmHistoryItem `xml:"DescribeAlarmHistoryResult>AlarmHistoryItems>member"`
	NextToken         string             `xml:"DescribeAlarmHistoryResult>NextToken"`
}

// DescribeAlarmsForMetricResult is a wrapper for DescribeAlarmsForMetricOutput.
type DescribeAlarmsForMetricResult struct {
	XMLName xml.Name `xml:"DescribeAlarmsForMetricResponse"`

	MetricAlarms []MetricAlarm `xml:"DescribeAlarmsForMetricResult>MetricAlarms>member"`
}

// DescribeAlarmsResult is a wrapper for DescribeAlarmsOutput.
type DescribeAlarmsResult struct {
	XMLName xml.Name `xml:"DescribeAlarmsResponse"`

	MetricAlarms []MetricAlarm `xml:"DescribeAlarmsResult>MetricAlarms>member"`
	NextToken    string        `xml:"DescribeAlarmsResult>NextToken"`
}

// GetMetricStatisticsResult is a wrapper for GetMetricStatisticsOutput.
type GetMetricStatisticsResult struct {
	XMLName xml.Name `xml:"GetMetricStatisticsResponse"`

	Datapoints []Datapoint `xml:"GetMetricStatisticsResult>Datapoints>member"`
	Label      string      `xml:"GetMetricStatisticsResult>Label"`
}

// ListMetricsResult is a wrapper for ListMetricsOutput.
type ListMetricsResult struct {
	XMLName xml.Name `xml:"ListMetricsResponse"`

	Metrics   []Metric `xml:"ListMetricsResult>Metrics>member"`
	NextToken string   `xml:"ListMetricsResult>NextToken"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
