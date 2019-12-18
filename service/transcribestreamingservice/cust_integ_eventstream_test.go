// +build integration

package transcribestreamingservice

import (
	"bytes"
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/awstesting/integration"
)

var (
	audioFilename   string
	audioFormat     string
	audioLang       string
	audioSampleRate int
	audioFrameSize  int
	withDebug       bool
)

func init() {
	flag.BoolVar(&withDebug, "debug", false, "Include debug logging with test.")
	flag.StringVar(&audioFilename, "audio-file", "", "Audio file filename to perform test with.")
	flag.StringVar(&audioLang, "audio-lang", LanguageCodeEnUs, "Language of audio speech.")
	flag.StringVar(&audioFormat, "audio-format", MediaEncodingPcm, "Format of audio.")
	flag.IntVar(&audioSampleRate, "audio-sample", 16000, "Sample rate of the audio.")
	flag.IntVar(&audioFrameSize, "audio-frame", 15*1024, "Size of frames of audio uploaded.")
}

func TestInteg_StartStreamTranscription(t *testing.T) {
	var audio io.Reader
	if len(audioFilename) != 0 {
		audioFile, err := os.Open(audioFilename)
		if err != nil {
			t.Fatalf("expect to open file, %v", err)
		}
		defer audioFile.Close()
		audio = audioFile
	} else {
		b, err := base64.StdEncoding.DecodeString(
			`UklGRjzxPQBXQVZFZm10IBAAAAABAAEAgD4AAAB9AAACABAAZGF0YVTwPQAAAAAAAAAAAAAAAAD//wIA/f8EAA==`,
		)
		if err != nil {
			t.Fatalf("expect decode audio bytes, %v", err)
		}
		audio = bytes.NewReader(b)
	}

	sess := integration.SessionWithDefaultRegion("us-west-2")
	var cfgs []*aws.Config
	if withDebug {
		cfgs = append(cfgs, &aws.Config{
			Logger:   t,
			LogLevel: aws.LogLevel(aws.LogDebugWithEventStreamBody),
		})
	}

	client := New(sess, cfgs...)
	resp, err := client.StartStreamTranscription(&StartStreamTranscriptionInput{
		LanguageCode:         aws.String(audioLang),
		MediaEncoding:        aws.String(audioFormat),
		MediaSampleRateHertz: aws.Int64(int64(audioSampleRate)),
	})
	if err != nil {
		t.Fatalf("failed to start streaming, %v", err)
	}
	stream := resp.GetStream()
	defer stream.Close()

	go StreamAudio(stream.Writer, audioFrameSize, audio)

	for event := range stream.Events() {
		switch e := event.(type) {
		case *TranscriptEvent:
			t.Logf("got event, %v results", len(e.Transcript.Results))
			for _, res := range e.Transcript.Results {
				for _, alt := range res.Alternatives {
					t.Logf("* %s", aws.StringValue(alt.Transcript))
				}
			}
		default:
			t.Fatalf("unexpected event, %T", event)
		}
	}

	if err := stream.Err(); err != nil {
		t.Fatalf("expect no error from stream, got %v", err)
	}
}

func StreamAudio(stream AudioStreamWriter, frameSize int, input io.Reader) (err error) {
	defer func() {
		if closeErr := stream.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close stream, %v", closeErr)
		}
	}()

	frame := make([]byte, frameSize)
	for {
		var n int
		n, err = input.Read(frame)
		if n > 0 {
			err = stream.Send(context.Background(), &AudioEvent{
				AudioChunk: frame[:n],
			})
			if err != nil {
				return fmt.Errorf("failed to send audio event, %v", err)
			}
		}

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("failed to read audio, %v", err)
		}
	}
}
