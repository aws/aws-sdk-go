package csm

import (
	"github.com/aws/aws-sdk-go/aws/request"
)

// Client side metric handler names
const (
	APICallMetricHandlerName        = "awscsm.SendAPICallMetric"
	APICallAttemptMetricHandlerName = "awscsm.SendAPICallAttemptMetric"
)

// Start will start the a long running go routine to capture
// client side metrics. Calling start multiple time will only
// start the metric listener once.
//
//	Example:
//		csm.Start("clientID", "127.0.0.1:8094")
//		sess := session.NewSession()
//		csm.InjectHandlers(sess.Handlers)
//
//		svc := s3.New(sess)
//		out, err := svc.GetObject(&s3.GetObjectInput{
//			Bucket: aws.String("bucket"),
//			Key: aws.String("key"),
//		})
func Start(clientID string, url string) error {
	lock.Lock()
	defer lock.Unlock()
	if sender == nil {
		sender = newReporter(clientID)
	}

	if sender.metricsCh.IsPaused() {
		sender.metricsCh.Continue()
	}

	return connect(url)
}

// Stop will pause the metric channel preventing any new metrics from
// being added.
func Stop() {
	lock.Lock()
	defer lock.Unlock()
	sender.Close()
}

// InjectHandlers will will enable client side metrics and inject the proper
// handlers to handle how metrics are sent.
//
//	Example:
//		sess := session.NewSession()
//		csm.InjectHandlers(&sess.Handlers)
//
//		// create a new service client with our client side metric session
//		svc := s3.New(sess)
func InjectHandlers(handlers *request.Handlers) {
	apiCallHandler := request.NamedHandler{Name: APICallMetricHandlerName, Fn: sender.SendAPICallMetric}
	handlers.Complete.PushFrontNamed(apiCallHandler)

	apiCallAttemptHandler := request.NamedHandler{Name: APICallAttemptMetricHandlerName, Fn: sender.SendAPICallAttemptMetric}
	handlers.AfterRetry.PushFrontNamed(apiCallAttemptHandler)
}
