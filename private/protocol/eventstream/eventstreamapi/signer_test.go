// +build go1.7

package eventstreamapi

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
)

func TestMessageSigner(t *testing.T) {
	currentTime := time.Date(2019, 1, 27, 22, 37, 54, 0, time.UTC)

	cases := map[string]struct {
		signer        StreamSigner
		input         eventstream.Message
		expected      eventstream.Message
		expectedError string
	}{
		"sign message": {
			signer: mockChunkSigner{signature: "524f1d03d1d81e94a099042736d40bd9681b867321443ff58a4568e274dbd83bff"},
			input: eventstream.Message{
				Headers: []eventstream.Header{
					{
						Name:  "header_name",
						Value: eventstream.StringValue("header value"),
					},
				},
				Payload: []byte("payload"),
			},
			expected: eventstream.Message{
				Headers: []eventstream.Header{
					{
						Name:  "header_name",
						Value: eventstream.StringValue("header value"),
					},
					{
						Name:  ":date",
						Value: eventstream.TimestampValue(currentTime),
					},
					{
						Name:  ":chunk-signature",
						Value: eventstream.BytesValue(mustDecodeHex(hex.DecodeString("524f1d03d1d81e94a099042736d40bd9681b867321443ff58a4568e274dbd83bff"))),
					},
				},
				Payload: []byte("payload"),
			},
		},
		"signing error": {
			signer: mockChunkSigner{err: fmt.Errorf("signing error")},
			input: eventstream.Message{
				Headers: []eventstream.Header{
					{
						Name:  "header_name",
						Value: eventstream.StringValue("header value"),
					},
				},
				Payload: []byte("payload"),
			},
			expectedError: "signing error",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			messageSigner := MessageSigner{Signer: tt.signer}

			err := messageSigner.SignMessage(&tt.input, currentTime)
			if err == nil && len(tt.expectedError) > 0 {
				t.Fatalf("expected error, but got nil")
			} else if err != nil && len(tt.expectedError) == 0 {
				t.Fatalf("expected no error, but got %v", err)
			} else if err != nil && len(tt.expectedError) > 0 && !strings.Contains(err.Error(), tt.expectedError) {
				t.Fatalf("expected %v, but got %v", tt.expectedError, err)
			} else if len(tt.expectedError) > 0 {
				return
			}

			if e, a := tt.expected, tt.input; !reflect.DeepEqual(e, a) {
				t.Errorf("expected %v, got %v", e, a)
			}
		})
	}
}
