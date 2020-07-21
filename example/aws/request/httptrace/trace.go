package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http/httptrace"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
)

// RequestTrace provides the strucutre to store SDK request attempt latencies.
// Use TraceRequest as a API operation request option to capture trace metrics
// for the individual request.
type RequestTrace struct {
	Start, Finish time.Time

	SDKValidateStart, SDKValidateDone time.Time
	SDKBuildStart, SDKBuildDone       time.Time

	ReadResponseBody bool

	Attempts []*RequestAttemptTrace
}

// TraceRequest is a SDK request Option that will add request handlers to an
// individual request to track request latencies per attempt. Must be used only
// for a single request call per RequestTrace value.
func (t *RequestTrace) TraceRequest(r *request.Request) {
	t.Start = time.Now()
	r.Handlers.Complete.PushBack(t.onComplete)

	r.Handlers.Validate.PushFront(t.onValidateStart)
	r.Handlers.Validate.PushBack(t.onValidateDone)

	r.Handlers.Build.PushFront(t.onBuildStart)
	r.Handlers.Build.PushBack(t.onBuildDone)

	var attempt *RequestAttemptTrace

	// Signing and Start attempt
	r.Handlers.Sign.PushFront(func(rr *request.Request) {
		attempt = &RequestAttemptTrace{Start: time.Now()}
		attempt.SDKSignStart = attempt.Start
	})
	r.Handlers.Sign.PushBack(func(rr *request.Request) {
		attempt.SDKSignDone = time.Now()
	})

	// Send
	r.Handlers.Send.PushFront(func(rr *request.Request) {
		attempt.SDKSendStart = time.Now()
		attempt.HTTPTrace = NewHTTPTrace(rr.Context())
		rr.SetContext(attempt.HTTPTrace)
	})
	r.Handlers.Send.PushBack(func(rr *request.Request) {
		attempt.SDKSendDone = time.Now()
		defer func() {
			attempt.HTTPTrace.Finish = time.Now()
		}()

		if rr.Error != nil {
			return
		}

		attempt.HTTPTrace.ReadHeaderDone = time.Now()
		if t.ReadResponseBody {
			attempt.HTTPTrace.ReadBodyStart = time.Now()
			var w bytes.Buffer
			if _, err := io.Copy(&w, rr.HTTPResponse.Body); err != nil {
				rr.Error = err
				return
			}
			rr.HTTPResponse.Body.Close()
			rr.HTTPResponse.Body = ioutil.NopCloser(&w)

			attempt.HTTPTrace.ReadBodyDone = time.Now()
		}
	})

	// Unmarshal
	r.Handlers.Unmarshal.PushFront(func(rr *request.Request) {
		attempt.SDKUnmarshalStart = time.Now()
	})
	r.Handlers.Unmarshal.PushBack(func(rr *request.Request) {
		attempt.SDKUnmarshalDone = time.Now()
	})

	// Unmarshal Error
	r.Handlers.UnmarshalError.PushFront(func(rr *request.Request) {
		attempt.SDKUnmarshalErrorStart = time.Now()
	})
	r.Handlers.UnmarshalError.PushBack(func(rr *request.Request) {
		attempt.SDKUnmarshalErrorDone = time.Now()
	})

	// Retry handling and delay
	r.Handlers.Retry.PushFront(func(rr *request.Request) {
		attempt.SDKRetryStart = time.Now()
		attempt.Err = rr.Error
	})
	r.Handlers.AfterRetry.PushBack(func(rr *request.Request) {
		attempt.SDKRetryDone = time.Now()
		attempt.WillRetry = rr.WillRetry()
	})

	// Complete Attempt
	r.Handlers.CompleteAttempt.PushBack(func(rr *request.Request) {
		attempt.Finish = time.Now()
		t.Attempts = append(t.Attempts, attempt)
	})
}

func (t *RequestTrace) String() string {
	var w strings.Builder

	writeDurField(&w, "Latency", t.Start, t.Finish)
	writeDurField(&w, "Validate", t.SDKValidateStart, t.SDKValidateDone)
	writeDurField(&w, "Build", t.SDKBuildStart, t.SDKBuildDone)
	writeField(&w, "Attempts", "%d", len(t.Attempts))

	for i, a := range t.Attempts {
		fmt.Fprintf(&w, "\n\tAttempt: %d, %s", i, a)
	}

	return w.String()
}

func (t *RequestTrace) onComplete(*request.Request) {
	t.Finish = time.Now()
}
func (t *RequestTrace) onValidateStart(*request.Request) { t.SDKValidateStart = time.Now() }
func (t *RequestTrace) onValidateDone(*request.Request)  { t.SDKValidateDone = time.Now() }
func (t *RequestTrace) onBuildStart(*request.Request)    { t.SDKBuildStart = time.Now() }
func (t *RequestTrace) onBuildDone(*request.Request)     { t.SDKBuildDone = time.Now() }

// RequestAttemptTrace provides a structure for storing trace information on
// the SDK's request attempt.
type RequestAttemptTrace struct {
	Start, Finish time.Time
	Err           error

	SDKSignStart, SDKSignDone time.Time

	SDKSendStart, SDKSendDone time.Time
	HTTPTrace                 *HTTPTrace

	SDKUnmarshalStart, SDKUnmarshalDone           time.Time
	SDKUnmarshalErrorStart, SDKUnmarshalErrorDone time.Time

	WillRetry                   bool
	SDKRetryStart, SDKRetryDone time.Time
}

func (at *RequestAttemptTrace) String() string {
	var w strings.Builder

	writeDurField(&w, "Latency", at.Start, at.Finish)
	writeDurField(&w, "Sign", at.SDKSignStart, at.SDKSignDone)
	writeDurField(&w, "Send", at.SDKSendStart, at.SDKSendDone)

	writeDurField(&w, "Unmarshal", at.SDKUnmarshalStart, at.SDKUnmarshalDone)
	writeDurField(&w, "UnmarshalError", at.SDKUnmarshalErrorStart, at.SDKUnmarshalErrorDone)

	writeField(&w, "WillRetry", "%t", at.WillRetry)
	writeDurField(&w, "Retry", at.SDKRetryStart, at.SDKRetryDone)

	fmt.Fprintf(&w, "\n\t\tHTTP: %s", at.HTTPTrace)
	if at.Err != nil {
		fmt.Fprintf(&w, "\n\t\tError: %v", at.Err)
	}

	return w.String()
}

// HTTPTrace provides the trace time stamps of the HTTP request's segments.
type HTTPTrace struct {
	context.Context

	Start, Finish time.Time

	GetConnStart, GetConnDone time.Time
	Reused                    bool

	DNSStart, DNSDone                   time.Time
	ConnectStart, ConnectDone           time.Time
	TLSHandshakeStart, TLSHandshakeDone time.Time
	WriteHeaderDone                     time.Time
	WriteRequestDone                    time.Time
	FirstResponseByte                   time.Time

	ReadHeaderStart, ReadHeaderDone time.Time
	ReadBodyStart, ReadBodyDone     time.Time
}

// NewHTTPTrace returns a initialized HTTPTrace for an
// httptrace.ClientTrace, based on the context passed.
func NewHTTPTrace(ctx context.Context) *HTTPTrace {
	t := &HTTPTrace{
		Start: time.Now(),
	}

	trace := &httptrace.ClientTrace{
		GetConn:              t.getConn,
		GotConn:              t.gotConn,
		PutIdleConn:          t.putIdleConn,
		GotFirstResponseByte: t.gotFirstResponseByte,
		Got100Continue:       t.got100Continue,
		DNSStart:             t.dnsStart,
		DNSDone:              t.dnsDone,
		ConnectStart:         t.connectStart,
		ConnectDone:          t.connectDone,
		TLSHandshakeStart:    t.tlsHandshakeStart,
		TLSHandshakeDone:     t.tlsHandshakeDone,
		WroteHeaders:         t.wroteHeaders,
		Wait100Continue:      t.wait100Continue,
		WroteRequest:         t.wroteRequest,
	}

	t.Context = httptrace.WithClientTrace(ctx, trace)

	return t
}

func (t *HTTPTrace) String() string {
	var w strings.Builder

	writeDurField(&w, "Latency", t.Start, t.Finish)
	writeField(&w, "ConnReused", "%t", t.Reused)

	if !t.Reused {
		writeDurField(&w, "GetConn", t.GetConnStart, t.GetConnDone)
		writeDurField(&w, "DNS", t.DNSStart, t.DNSDone)
		writeDurField(&w, "Connect", t.ConnectStart, t.ConnectDone)
		writeDurField(&w, "TLS", t.TLSHandshakeStart, t.TLSHandshakeDone)
	} else {
		writeDurField(&w, "GetConn", t.Start, t.GetConnDone)
	}

	writeDurField(&w, "WriteRequest", t.GetConnDone, t.WriteRequestDone)
	writeDurField(&w, "WaitResponseFirstByte", t.Start, t.FirstResponseByte)
	writeDurField(&w, "ReadResponseHeader", t.ReadHeaderStart, t.ReadHeaderDone)
	writeDurField(&w, "ReadResponseBody", t.ReadBodyStart, t.ReadBodyDone)

	return w.String()
}

func (t *HTTPTrace) getConn(hostPort string) {
	t.GetConnStart = time.Now()
}
func (t *HTTPTrace) gotConn(info httptrace.GotConnInfo) {
	t.GetConnDone = time.Now()
	t.Reused = info.Reused
}
func (t *HTTPTrace) putIdleConn(err error) {}
func (t *HTTPTrace) gotFirstResponseByte() {
	t.FirstResponseByte = time.Now()
	t.ReadHeaderStart = t.FirstResponseByte
}
func (t *HTTPTrace) got100Continue() {}
func (t *HTTPTrace) dnsStart(info httptrace.DNSStartInfo) {
	t.DNSStart = time.Now()
}
func (t *HTTPTrace) dnsDone(info httptrace.DNSDoneInfo) {
	t.DNSDone = time.Now()
}
func (t *HTTPTrace) connectStart(network, addr string) {
	t.ConnectStart = time.Now()
}
func (t *HTTPTrace) connectDone(network, addr string, err error) {
	t.ConnectDone = time.Now()
}
func (t *HTTPTrace) tlsHandshakeStart() {
	t.TLSHandshakeStart = time.Now()
}
func (t *HTTPTrace) tlsHandshakeDone(state tls.ConnectionState, err error) {
	t.TLSHandshakeDone = time.Now()
}
func (t *HTTPTrace) wroteHeaders() {
	t.WriteHeaderDone = time.Now()
}
func (t *HTTPTrace) wait100Continue() {}
func (t *HTTPTrace) wroteRequest(info httptrace.WroteRequestInfo) {
	t.WriteRequestDone = time.Now()
}

func writeField(w io.Writer, field string, format string, args ...interface{}) error {
	_, err := fmt.Fprintf(w, "%s: "+format+", ", append([]interface{}{field}, args...)...)
	return err
}

func writeDurField(w io.Writer, field string, start, stop time.Time) error {
	if start.IsZero() || stop.IsZero() {
		return nil
	}

	_, err := fmt.Fprintf(w, "%s: %s, ", field, stop.Sub(start))
	return err
}
