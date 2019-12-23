// +build go1.6

package transcribestreamingservice

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream/eventstreamtest"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

func TestStartStreamTranscription_Write(t *testing.T) {
	writeRequests, requestMessages := mockStartStreamTranscriptionWriteEvents()

	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		&eventstreamtest.ServeEventStream{
			T:             t,
			ClientEvents:  requestMessages,
			BiDirectional: true,
		},
		true)
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartStreamTranscription(nil)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	stream := resp.GetStream()

	for _, event := range writeRequests {
		err = stream.Send(context.Background(), &event)
		if err != nil {
			t.Fatalf("expect no error, got %v", err)
		}
	}

	if err := stream.Close(); err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestStartStreamTranscription_WriteClose(t *testing.T) {
	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{T: t, BiDirectional: true},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartStreamTranscription(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}

	// Assert calling Err before close does not close the stream.
	resp.GetStream().Err()
	err = resp.GetStream().Send(context.Background(), &AudioEvent{})
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	resp.GetStream().Close()

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func TestStartStreamTranscription_WriteError(t *testing.T) {
	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:               t,
			BiDirectional:   true,
			ForceCloseAfter: time.Millisecond * 500,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartStreamTranscription(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}

	defer resp.GetStream().Close()

	for {
		err = resp.GetStream().Send(context.Background(), &AudioEvent{})
		if err != nil {
			if strings.Contains("unable to send event", err.Error()) {
				t.Errorf("expected stream closed error, got %v", err)
			}
			break
		}
	}
}

func TestStartStreamTranscription_ReadWrite(t *testing.T) {
	expectedServiceEvents, serviceEvents := mockStartStreamTranscriptionReadEvents()
	clientEvents, expectedClientEvents := mockStartStreamTranscriptionWriteEvents()

	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		&eventstreamtest.ServeEventStream{
			T:             t,
			ClientEvents:  expectedClientEvents,
			Events:        serviceEvents,
			BiDirectional: true,
		},
		true)
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartStreamTranscription(nil)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	stream := resp.GetStream()
	defer stream.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		var i int
		for event := range resp.GetStream().Events() {
			if event == nil {
				t.Errorf("%d, expect event, got nil", i)
			}
			if e, a := expectedServiceEvents[i], event; !reflect.DeepEqual(e, a) {
				t.Errorf("%d, expect %T %v, got %T %v", i, e, e, a, a)
			}
			i++
		}
	}()

	for _, event := range clientEvents {
		err = stream.Send(context.Background(), &event)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}

	resp.GetStream().Close()

	wg.Wait()

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func mockStartStreamTranscriptionWriteEvents() (
	[]AudioEvent,
	[]eventstream.Message,
) {
	audioEvents := []AudioEvent{
		{AudioChunk: make([]byte, 64)},
	}

	var eventMessages []eventstream.Message

	var marshalers request.HandlerList
	marshalers.PushBackNamed(restjson.BuildHandler)

	pm := protocol.HandlerPayloadMarshal{Marshalers: marshalers}

	for _, audioEvent := range audioEvents {
		event, err := audioEvent.MarshalEvent(pm)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal event %v", err))
		}

		event.Headers.Set(":event-type", eventstream.StringValue("AudioEvent"))

		eventMessages = append(eventMessages, event)
	}

	return audioEvents, eventMessages
}
