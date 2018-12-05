// +build go1.7

package csm_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/csm"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

func TestReportingMetrics(t *testing.T) {
	sess := unit.Session.Copy(&aws.Config{
		SleepDelay: func(time.Duration) {},
	})
	sess.Handlers.Validate.Clear()
	sess.Handlers.Sign.Clear()
	sess.Handlers.Send.Clear()

	reporter := csm.Get()
	if reporter == nil {
		t.Errorf("expected non-nil reporter")
	}
	reporter.InjectHandlers(&sess.Handlers)

	cases := map[string]struct {
		Request       *request.Request
		ExpectMetrics []map[string]interface{}
	}{
		"successful request": {
			Request: func() *request.Request {
				md := metadata.ClientInfo{}
				op := &request.Operation{Name: "OperationName"}
				req := request.New(*sess.Config, md, sess.Handlers, client.DefaultRetryer{NumMaxRetries: 3}, op, nil, nil)
				req.Handlers.Send.PushBack(func(r *request.Request) {
					req.HTTPResponse = &http.Response{
						StatusCode: 200,
						Header:     http.Header{},
					}
				})
				return req
			}(),
			ExpectMetrics: []map[string]interface{}{
				{
					"Type":           "ApiCallAttempt",
					"HttpStatusCode": float64(200),
				},
				{
					"Type": "ApiCall",
				},
			},
		},
		"failed request, no retry": {
			Request: func() *request.Request {
				md := metadata.ClientInfo{}
				op := &request.Operation{Name: "OperationName"}
				req := request.New(*sess.Config, md, sess.Handlers, client.DefaultRetryer{NumMaxRetries: 3}, op, nil, nil)
				req.Handlers.Send.PushBack(func(r *request.Request) {
					req.HTTPResponse = &http.Response{
						StatusCode: 400,
						Header:     http.Header{},
					}
					req.Retryable = aws.Bool(false)
				})

				return req
			}(),
			ExpectMetrics: []map[string]interface{}{
				{
					"Type":           "ApiCallAttempt",
					"HttpStatusCode": float64(400),
				},
				{
					"Type":         "ApiCall",
					"AttemptCount": float64(1),
				},
			},
		},
		"failed request, with retry": {
			Request: func() *request.Request {
				md := metadata.ClientInfo{}
				op := &request.Operation{Name: "OperationName"}
				req := request.New(*sess.Config, md, sess.Handlers, client.DefaultRetryer{NumMaxRetries: 1}, op, nil, nil)
				resps := []*http.Response{
					{
						StatusCode: 500,
						Header:     http.Header{},
					},
					{
						StatusCode: 500,
						Header:     http.Header{},
					},
				}
				req.Handlers.Send.PushBack(func(r *request.Request) {
					req.HTTPResponse = resps[0]
					resps = resps[1:]
				})

				return req
			}(),
			ExpectMetrics: []map[string]interface{}{
				{
					"Type":           "ApiCallAttempt",
					"HttpStatusCode": float64(500),
				},
				{
					"Type":           "ApiCallAttempt",
					"HttpStatusCode": float64(500),
				},
				{
					"Type":         "ApiCall",
					"AttemptCount": float64(2),
				},
			},
		},
		"success request, with retry": {
			Request: func() *request.Request {
				md := metadata.ClientInfo{}
				op := &request.Operation{Name: "OperationName"}
				req := request.New(*sess.Config, md, sess.Handlers, client.DefaultRetryer{NumMaxRetries: 3}, op, nil, nil)
				resps := []*http.Response{
					{
						StatusCode: 500,
						Header:     http.Header{},
					},
					{
						StatusCode: 500,
						Header:     http.Header{},
					},
					{
						StatusCode: 200,
						Header:     http.Header{},
					},
				}
				req.Handlers.Send.PushBack(func(r *request.Request) {
					req.HTTPResponse = resps[0]
					resps = resps[1:]
				})

				return req
			}(),
			ExpectMetrics: []map[string]interface{}{
				{
					"Type":           "ApiCallAttempt",
					"HttpStatusCode": float64(500),
				},
				{
					"Type":           "ApiCallAttempt",
					"HttpStatusCode": float64(500),
				},
				{
					"Type":           "ApiCallAttempt",
					"HttpStatusCode": float64(200),
				},
				{
					"Type":         "ApiCall",
					"AttemptCount": float64(3),
				},
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
			defer cancelFn()

			c.Request.Send()
			for i := 0; i < len(c.ExpectMetrics); i++ {
				select {
				case m := <-csm.MetricsCh:
					for ek, ev := range c.ExpectMetrics[i] {
						if _, ok := m[ek]; !ok {
							t.Errorf("expect %v metric member", ek)
						}
						if e, a := ev, m[ek]; e != a {
							t.Errorf("expect %v:%v(%T), metric value, got %v(%T)", ek, e, e, a, a)
						}
					}
				case <-ctx.Done():
					t.Errorf("timeout waiting for metrics")
					return
				}
			}

			var extraMetrics []map[string]interface{}
		Loop:
			for {
				select {
				case m := <-csm.MetricsCh:
					extraMetrics = append(extraMetrics, m)
				default:
					break Loop
				}
			}
			if len(extraMetrics) != 0 {
				t.Fatalf("unexpected metrics, %#v", extraMetrics)
			}
		})
	}
}

type mockService struct {
	*client.Client
}

type input struct{}
type output struct{}

func (s *mockService) Request(i input) *request.Request {
	op := &request.Operation{
		Name:       "foo",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	o := output{}
	req := s.NewRequest(op, &i, &o)
	return req
}

func BenchmarkWithCSM(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("{}")))
	}))

	cfg := aws.Config{
		Endpoint: aws.String(server.URL),
	}

	sess := unit.Session.Copy(&cfg)
	r := csm.Get()

	r.InjectHandlers(&sess.Handlers)

	c := sess.ClientConfig("id", &cfg)

	svc := mockService{
		client.New(
			*c.Config,
			metadata.ClientInfo{
				ServiceName:   "service",
				ServiceID:     "id",
				SigningName:   "signing",
				SigningRegion: "region",
				Endpoint:      server.URL,
				APIVersion:    "0",
				JSONVersion:   "1.1",
				TargetPrefix:  "prefix",
			},
			c.Handlers,
		),
	}

	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	for i := 0; i < b.N; i++ {
		req := svc.Request(input{})
		req.Send()
	}
}

func BenchmarkWithCSMNoUDPConnection(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("{}")))
	}))

	cfg := aws.Config{
		Endpoint: aws.String(server.URL),
	}

	sess := unit.Session.Copy(&cfg)
	r := csm.Get()
	r.Pause()
	r.InjectHandlers(&sess.Handlers)
	defer r.Pause()

	c := sess.ClientConfig("id", &cfg)

	svc := mockService{
		client.New(
			*c.Config,
			metadata.ClientInfo{
				ServiceName:   "service",
				ServiceID:     "id",
				SigningName:   "signing",
				SigningRegion: "region",
				Endpoint:      server.URL,
				APIVersion:    "0",
				JSONVersion:   "1.1",
				TargetPrefix:  "prefix",
			},
			c.Handlers,
		),
	}

	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	for i := 0; i < b.N; i++ {
		req := svc.Request(input{})
		req.Send()
	}
}

func BenchmarkWithoutCSM(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("{}")))
	}))

	cfg := aws.Config{
		Endpoint: aws.String(server.URL),
	}
	sess := unit.Session.Copy(&cfg)
	c := sess.ClientConfig("id", &cfg)

	svc := mockService{
		client.New(
			*c.Config,
			metadata.ClientInfo{
				ServiceName:   "service",
				ServiceID:     "id",
				SigningName:   "signing",
				SigningRegion: "region",
				Endpoint:      server.URL,
				APIVersion:    "0",
				JSONVersion:   "1.1",
				TargetPrefix:  "prefix",
			},
			c.Handlers,
		),
	}

	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	for i := 0; i < b.N; i++ {
		req := svc.Request(input{})
		req.Send()
	}
}
