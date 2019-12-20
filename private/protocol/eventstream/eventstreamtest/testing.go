package eventstreamtest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
)

// ServeEventStream provides serving EventStream messages from a HTTP server to
// the client. The events are sent sequentially to the client without delay.
type ServeEventStream struct {
	T      *testing.T
	Events []eventstream.Message
}

func (s ServeEventStream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	encoder := eventstream.NewEncoder(flushWriter{w})

	go io.Copy(ioutil.Discard, r.Body)

	for _, event := range s.Events {
		encoder.Encode(event)
	}
}

// SetupEventStreamSession creates a HTTP server SDK session for communicating
// with that server to be used for EventStream APIs. If HTTP/2 is enabled the
// server/client will only attempt to use HTTP/2.
func SetupEventStreamSession(
	t *testing.T, handler http.Handler, h2 bool,
) (sess *session.Session, cleanupFn func(), err error) {
	server := httptest.NewUnstartedServer(handler)

	client := setupServer(server, h2)

	cleanupFn = func() {
		server.Close()
	}

	sess, err = session.NewSession(unit.Session.Config, &aws.Config{
		Endpoint:               &server.URL,
		DisableParamValidation: aws.Bool(true),
		HTTPClient:             client,
		//		LogLevel:               aws.LogLevel(aws.LogDebugWithEventStreamBody),
	})
	if err != nil {
		return nil, nil, err
	}

	return sess, cleanupFn, nil
}

type flushWriter struct {
	w io.Writer
}

func (fw flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if f, ok := fw.w.(http.Flusher); ok {
		f.Flush()
	}
	return
}

// MarshalEventPayload marshals a SDK API shape into its associated wire
// protocol payload.
func MarshalEventPayload(
	payloadMarshaler protocol.PayloadMarshaler,
	v interface{},
) []byte {
	var w bytes.Buffer
	err := payloadMarshaler.MarshalPayload(&w, v)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal event %T, %v", v, v))
	}

	return w.Bytes()
}

// Prevent circular dependencies on eventstreamapi redefine these here.
const (
	messageTypeHeader    = `:message-type` // Identifies type of message.
	eventMessageType     = `event`
	exceptionMessageType = `exception`
)

// EventMessageTypeHeader is an event message type header for specifying an
// event is an message type.
var EventMessageTypeHeader = eventstream.Header{
	Name:  messageTypeHeader,
	Value: eventstream.StringValue(eventMessageType),
}

// EventExceptionTypeHeader is an event exception type header for specifying an
// event is an exception type.
var EventExceptionTypeHeader = eventstream.Header{
	Name:  messageTypeHeader,
	Value: eventstream.StringValue(exceptionMessageType),
}

// AssertMessageEqual compares to event stream messages, and determines if they
// are equal. Will trigger an testing Error if components of the message are
// not equal.
func AssertMessageEqual(t testing.TB, a, b eventstream.Message, msg ...interface{}) {
	getHelper(t)()

	ah, err := bytesEncodeHeader(a.Headers)
	if err != nil {
		t.Fatalf("unable to encode a's headers, %v", err)
	}

	bh, err := bytesEncodeHeader(b.Headers)
	if err != nil {
		t.Fatalf("unable to encode b's headers, %v", err)
	}

	if !bytes.Equal(ah, bh) {
		aj, err := json.Marshal(ah)
		if err != nil {
			t.Fatalf("unable to json encode a's headers, %v", err)
		}
		bj, err := json.Marshal(bh)
		if err != nil {
			t.Fatalf("unable to json encode b's headers, %v", err)
		}
		t.Errorf("%s\nexpect headers: %v\n\t%v\nactual headers: %v\n\t%v\n",
			fmt.Sprint(msg...),
			base64.StdEncoding.EncodeToString(ah), aj,
			base64.StdEncoding.EncodeToString(bh), bj,
		)
	}

	if !bytes.Equal(a.Payload, b.Payload) {
		t.Errorf("%s\nexpect payload: %v\nactual payload: %v\n",
			fmt.Sprint(msg...),
			base64.StdEncoding.EncodeToString(a.Payload),
			base64.StdEncoding.EncodeToString(b.Payload),
		)
	}
}

func bytesEncodeHeader(v eventstream.Headers) ([]byte, error) {
	var buf bytes.Buffer
	if err := eventstream.EncodeHeaders(&buf, v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
