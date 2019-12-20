package eventstreamtest

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
)

// ServeEventStream provides serving EventStream messages from a HTTP server to
// the client. The events are sent sequentially to the client without delay.
type ServeEventStream struct {
	T             *testing.T
	BiDirectional bool

	Events       []eventstream.Message
	ClientEvents []eventstream.Message

	ForceCloseAfter time.Duration

	requestsIdx int
}

func (s ServeEventStream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.(http.Flusher).Flush()

	if s.BiDirectional {
		s.serveBiDirectionalStream(w, r)
	} else {
		s.serveReadOnlyStream(w, r)
	}
}

func (s *ServeEventStream) serveReadOnlyStream(w http.ResponseWriter, r *http.Request) {
	encoder := eventstream.NewEncoder(flushWriter{w})

	for _, event := range s.Events {
		encoder.Encode(event)
	}
}

func (s *ServeEventStream) serveBiDirectionalStream(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	ctx := context.Background()
	if s.ForceCloseAfter > 0 {
		var cancelFunc func()
		ctx, cancelFunc = context.WithTimeout(context.Background(), s.ForceCloseAfter)
		defer cancelFunc()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.readEvents(ctx, r)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.writeEvents(ctx, w)
	}()

	wg.Wait()
}

func (s ServeEventStream) readEvents(ctx context.Context, r *http.Request) {
	messages := make(chan eventstream.Message)

	go func() {
		defer close(messages)

		payloadBuffer := make([]byte, 1024)
		decoder := eventstream.NewDecoder(r.Body, eventstream.DecodeWithLogger(defaults.Config().Logger))

		for {
			// unwrap signing envelope
			message, err := decoder.Decode(payloadBuffer)
			if err != nil {
				break
			}
			messages <- message
		}
	}()

	payloadBuffer := make([]byte, 1024)

eventLoop:
	for {
		select {
		case message, ok := <-messages:
			if !ok {
				break eventLoop
			}

			// get service event message from payload
			message, err := eventstream.Decode(bytes.NewReader(message.Payload), payloadBuffer)
			if err != nil {
				if err == io.EOF {
					break eventLoop
				}
				s.T.Errorf("expected no error decoding event, got %v", err)
				break eventLoop
			}

			// empty payload is expected for the last signing message
			if len(message.Payload) == 0 {
				break eventLoop
			}

			if len(s.ClientEvents) > 0 {
				i := s.requestsIdx
				s.requestsIdx++

				if e, a := s.ClientEvents[i], message; !reflect.DeepEqual(e, a) {
					s.T.Errorf("expected %v, got %v", e, a)
					break eventLoop
				}
			}
		case <-ctx.Done():
			break eventLoop
		}
	}

	return
}

func (s *ServeEventStream) writeEvents(ctx context.Context, w http.ResponseWriter) {
	events := make(chan eventstream.Message)
	defer close(events)

	go func() {
		encoder := eventstream.NewEncoder(flushWriter{w})
		for event := range events {
			err := encoder.Encode(event)
			if err != nil {
				if err == io.EOF {
					return
				}
				s.T.Errorf("expected no error encoding event, got %v", err)
			}
		}
	}()

	var event eventstream.Message
	pendingEvents := s.Events

eventLoop:
	for len(pendingEvents) > 0 {
		event, pendingEvents = pendingEvents[0], pendingEvents[1:]
		select {
		case events <- event:
			continue
		case <-ctx.Done():
			break eventLoop
		}
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
