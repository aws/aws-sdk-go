package csm

import (
	"encoding/json"
	"net"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
)

const (
	// DefaultPort is used when no port is specified
	DefaultPort = "31000"
)

type reporter struct {
	clientID  string
	conn      net.Conn
	metricsCh metricChan
	done      chan struct{}
}

var (
	lock   sync.Mutex
	sender *reporter
)

func connect(url string) error {
	const network = "udp"
	if err := sender.connect(network, url); err != nil {
		return err
	}

	if sender.done == nil {
		sender.done = make(chan struct{})
		go sender.start()
	}

	return nil
}

func newReporter(clientID string) *reporter {
	return &reporter{
		clientID:  clientID,
		metricsCh: newMetricChan(DefaultChannelSize),
	}
}

func (rep *reporter) SendAPICallAttemptMetric(r *request.Request) {
	if rep == nil {
		return
	}

	now := time.Now()
	creds, _ := r.Config.Credentials.Get()

	m := metric{
		ClientID:  aws.String(rep.clientID),
		API:       aws.String(r.Operation.Name),
		Service:   aws.String(r.ClientInfo.ServiceID),
		Timestamp: (*metricTime)(&now),
		UserAgent: aws.String(r.HTTPRequest.Header.Get("User-Agent")),
		Region:    r.Config.Region,
		Type:      aws.String("ApiCallAttempt"),
		Version:   aws.Int(1),

		XAmzRequestID: aws.String(r.RequestID),

		AttemptCount:   aws.Int(r.RetryCount + 1),
		AttemptLatency: aws.Int(int(now.Sub(r.AttemptTime).Nanoseconds() / int64(time.Millisecond))),
		AccessKey:      aws.String(creds.AccessKeyID),
	}

	if r.HTTPResponse != nil {
		m.HTTPStatusCode = aws.Int(r.HTTPResponse.StatusCode)
	}

	if r.Error != nil {
		if awserr, ok := r.Error.(awserr.Error); ok {
			setError(&m, awserr)
		}
	}

	rep.metricsCh.Push(m)
}

func setError(m *metric, err awserr.Error) {
	msg := err.Message()
	code := err.Code()

	switch code {
	case "RequestError",
		"SerializationError",
		request.CanceledErrorCode:

		m.SDKException = &code
		m.SDKExceptionMessage = &msg
	default:
		m.AWSException = &code
		m.AWSExceptionMessage = &msg
	}
}

func (rep *reporter) SendAPICallMetric(r *request.Request) {
	if rep == nil {
		return
	}

	now := time.Now()
	m := metric{
		ClientID:      aws.String(rep.clientID),
		API:           aws.String(r.Operation.Name),
		Service:       aws.String(r.ClientInfo.ServiceID),
		Timestamp:     (*metricTime)(&now),
		Type:          aws.String("ApiCall"),
		AttemptCount:  aws.Int(r.RetryCount + 1),
		Latency:       aws.Int(int(time.Now().Sub(r.Time) / time.Millisecond)),
		XAmzRequestID: aws.String(r.RequestID),
	}

	// TODO: Probably want to figure something out for logging dropped
	// metrics
	rep.metricsCh.Push(m)
}

func (rep *reporter) connect(network, url string) error {
	if rep.conn != nil {
		rep.conn.Close()
	}

	conn, err := net.Dial(network, url)
	if err != nil {
		return awserr.New("UDPError", "Could not connect", err)
	}

	rep.conn = conn

	return nil
}

func (rep *reporter) Close() {
	if rep.done != nil {
		close(rep.done)
		rep.done = nil
	}

	rep.metricsCh.Pause()
}

func (rep *reporter) start() {
	defer func() {
		rep.metricsCh.Pause()
	}()

Loop:
	for {
		select {
		case <-rep.done:
			rep.done = nil
			break Loop
		case m := <-rep.metricsCh.ch:
			// TODO: What to do with this error? Probably should just log
			b, err := json.Marshal(m)
			if err != nil {
				continue
			}

			rep.conn.Write(b)
		}
	}
}
