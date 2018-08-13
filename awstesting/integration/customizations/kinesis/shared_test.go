// +build integration

package kinesis

import (
	crand "crypto/rand"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/awstesting/integration"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"golang.org/x/net/http2"
)

var (
	skipTLSVerify    bool
	hUsage           string
	endpoint         string
	streamName       string
	consumerName     string
	numRecords       int
	recordSize       int
	debugEventStream bool
	mode             string

	svc     *kinesis.Kinesis
	records []*kinesis.PutRecordsRequestEntry

	startingTimestamp time.Time
)

func init() {
	flag.StringVar(
		&mode, "mode", "test",
		"Sets the mode to run in, (test,create,cleanup).",
	)
	flag.BoolVar(
		&skipTLSVerify, "skip-verify", false,
		"Skips verification of TLS certificate.",
	)
	flag.StringVar(
		&hUsage, "http", "default",
		"The HTTP `version` to use for the connection. (default,1,2)",
	)
	flag.StringVar(
		&endpoint, "endpoint", "",
		"Overrides SDK `URL` endpoint for tests.",
	)
	flag.StringVar(
		&streamName, "stream", fmt.Sprintf("awsdkgo-s%v", UniqueID()),
		"The `name` of the stream to test against.",
	)
	flag.StringVar(
		&consumerName, "consumer", fmt.Sprintf("awsdkgo-c%v", UniqueID()),
		"The `name` of the stream to test against.",
	)
	flag.IntVar(
		&numRecords, "records", 20,
		"The `number` of records per PutRecords to test with.",
	)
	flag.IntVar(
		&recordSize, "record-size", 500,
		"The size in `bytes` of each record.",
	)
	flag.BoolVar(
		&debugEventStream, "debug-eventstream", false,
		"Enables debugging of the EventStream messages",
	)
	flag.Parse()
}

func TestMain(m *testing.M) {
	svc = createClient()

	startingTimestamp = time.Now().Add(-time.Minute)

	switch mode {
	case "create":
		if err := createStream(streamName); err != nil {
			panic(err)
		}
		if err := createStreamConsumer(streamName, consumerName); err != nil {
			panic(err)
		}
		fmt.Println("Stream Ready:", streamName, consumerName)

	case "cleanup":
		if err := cleanupStreamConsumer(streamName, consumerName); err != nil {
			panic(err)
		}
		if err := cleanupStream(streamName); err != nil {
			panic(err)
		}

	case "test":
		records = createRecords(numRecords, recordSize)
		if err := putRecords(streamName, records, svc); err != nil {
			panic(err)
		}
		time.Sleep(time.Second)

		var exitCode int
		defer func() {
			os.Exit(exitCode)
		}()

		exitCode = m.Run()
	default:
		fmt.Fprintf(os.Stderr, "unknown mode, %v", mode)
		os.Exit(1)
	}
}

func createClient() *kinesis.Kinesis {
	ts := &http.Transport{}

	if skipTLSVerify {
		ts.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	http2.ConfigureTransport(ts)
	switch hUsage {
	case "default":
		// Restore H2 optional support since the Transport/TLSConfig was
		// modified.
		http2.ConfigureTransport(ts)
	case "1":
		// Do nothing. Without usign ConfigureTransport h2 won't be available.
		ts.TLSClientConfig.NextProtos = []string{"http/1.1"}
	case "2":
		// Force the TLS ALPN (NextProto) to H2 only.
		ts.TLSClientConfig.NextProtos = []string{http2.NextProtoTLS}
	default:
		panic("unknown h usage, " + hUsage)
	}

	logLevel := integration.Session.Config.LogLevel.Value()
	if debugEventStream {
		logLevel |= aws.LogDebugWithEventStreamBody
	}

	sess := integration.Session.Copy(&aws.Config{
		HTTPClient: &http.Client{
			Transport: ts,
		},
		LogLevel: &logLevel,
	})
	if len(endpoint) != 0 {
		sess.Config.Endpoint = &endpoint
	}

	return kinesis.New(sess)
}

func createStream(name string) error {
	descParams := &kinesis.DescribeStreamInput{
		StreamName: &name,
	}

	_, err := svc.DescribeStream(descParams)
	if aerr, ok := err.(awserr.Error); ok && aerr.Code() == kinesis.ErrCodeResourceNotFoundException {
		_, err := svc.CreateStream(&kinesis.CreateStreamInput{
			ShardCount: aws.Int64(100),
			StreamName: &name,
		})
		if err != nil {
			return fmt.Errorf("failed to create stream, %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to describe stream, %v", err)
	}

	if err := svc.WaitUntilStreamExists(descParams); err != nil {
		return fmt.Errorf("failed to wait for stream to exist, %v", err)
	}

	return nil
}

func cleanupStream(name string) error {
	_, err := svc.DeleteStream(&kinesis.DeleteStreamInput{
		StreamName:              &name,
		EnforceConsumerDeletion: aws.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("failed to delete stream, %v", err)
	}

	return nil
}

func createStreamConsumer(streamName, consumerName string) error {
	desc, err := svc.DescribeStream(&kinesis.DescribeStreamInput{
		StreamName: &streamName,
	})
	if err != nil {
		return fmt.Errorf("failed to describe stream, %s, %v", streamName, err)
	}

	descParams := &kinesis.DescribeStreamConsumerInput{
		StreamARN:    desc.StreamDescription.StreamARN,
		ConsumerName: &consumerName,
	}
	_, err = svc.DescribeStreamConsumer(descParams)
	if aerr, ok := err.(awserr.Error); ok && aerr.Code() == kinesis.ErrCodeResourceNotFoundException {
		_, err := svc.RegisterStreamConsumer(
			&kinesis.RegisterStreamConsumerInput{
				ConsumerName: aws.String(consumerName),
				StreamARN:    desc.StreamDescription.StreamARN,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to create stream consumer %s, %v",
				consumerName, err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to describe stream consumer %s, %v",
			consumerName, err)
	}

	if err := svc.WaitUntilStreamConsumerExists(descParams); err != nil {
		return fmt.Errorf("failed to wait stream consumer to exist %s, %v",
			consumerName, err)
	}

	return nil
}

func cleanupStreamConsumer(streamName, consumerName string) error {
	desc, err := svc.DescribeStream(&kinesis.DescribeStreamInput{
		StreamName: &streamName,
	})
	if err != nil {
		return fmt.Errorf("failed to describe stream, %s, %v",
			streamName, err)
	}

	descCons, err := svc.DescribeStreamConsumer(&kinesis.DescribeStreamConsumerInput{
		StreamARN:    desc.StreamDescription.StreamARN,
		ConsumerName: &consumerName,
	})
	if err != nil {
		return fmt.Errorf("failed to describe stream consumer, %s, %v",
			consumerName, err)
	}

	_, err = svc.DeregisterStreamConsumer(
		&kinesis.DeregisterStreamConsumerInput{
			ConsumerName: descCons.ConsumerDescription.ConsumerName,
			ConsumerARN:  descCons.ConsumerDescription.ConsumerARN,
			StreamARN:    desc.StreamDescription.StreamARN,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to delete stream consumer, %s %v",
			consumerName, err)
	}

	return nil
}

func createRecords(num, size int) []*kinesis.PutRecordsRequestEntry {
	var err error
	data, err := loadRandomData(num, size)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read random data, %v", err)
		os.Exit(1)
	}

	records := make([]*kinesis.PutRecordsRequestEntry, len(data))
	for i, td := range data {
		records[i] = &kinesis.PutRecordsRequestEntry{
			Data:         td,
			PartitionKey: aws.String(UniqueID()),
		}
	}

	return records
}

func putRecords(stream string, records []*kinesis.PutRecordsRequestEntry, svc *kinesis.Kinesis) error {
	resp, err := svc.PutRecords(&kinesis.PutRecordsInput{
		StreamName: &stream,
		Records:    records,
	})
	if err != nil {
		return fmt.Errorf("failed to put records to stream %s, %v", stream, err)
	}

	if v := aws.Int64Value(resp.FailedRecordCount); v != 0 {
		return fmt.Errorf("failed to put records to stream %s, %d failed",
			stream, v)
	}

	return nil
}

func loadRandomData(m, n int) ([][]byte, error) {
	data := make([]byte, m*n)

	_, err := rand.Read(data)
	if err != nil {
		return nil, err
	}

	parts := make([][]byte, m)

	for i := 0; i < m; i++ {
		mod := (i % m)
		parts[i] = data[mod*n : (mod+1)*n]
	}

	return parts, nil
}

// UniqueID returns a unique UUID-like identifier for use in generating
// resources for integration tests.
func UniqueID() string {
	uuid := make([]byte, 16)
	io.ReadFull(crand.Reader, uuid)
	return fmt.Sprintf("%x", uuid)
}
