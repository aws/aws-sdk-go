// +build integration,perftest

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/awstesting/integration"
	"github.com/aws/aws-sdk-go/internal/sdkio"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var benchConfig BenchmarkConfig

type BenchmarkConfig struct {
	bucket       string
	tempdir      string
	clientConfig ClientConfig
}

func (b *BenchmarkConfig) SetupFlags(prefix string, flagSet *flag.FlagSet) {
	flagSet.StringVar(&b.bucket, "bucket", "", "Bucket to use for benchmark")
	flagSet.StringVar(&b.tempdir, "temp", os.TempDir(), "location to create temporary files")
	b.clientConfig.SetupFlags(prefix, flagSet)
}

var benchStrategies = map[string]s3manager.WriterReadFromProvider{
	"Unbuffered": nil,
	"Buffered":   s3manager.NewPooledBufferedWriterReadFromProvider(int(sdkio.MebiByte)),
}

func BenchmarkDownload(b *testing.B) {
	baseSdkConfig := SDKConfig{}

	// FileSizes: 5 MB, 1 GB
	for _, fileSize := range []int64{5 * sdkio.MebiByte, 1 * sdkio.GibiByte} {
		key, err := setupDownloadTest(benchConfig.bucket, fileSize)
		if err != nil {
			b.Fatalf("failed to setup download test: %v", err)
		}
		f, err := ioutil.TempFile(benchConfig.tempdir, "BenchmarkDownload")
		if err != nil {
			b.Fatalf("failed to create temporary file: %v", err)
		}
		b.Run(fmt.Sprintf("%s File", integration.SizeToName(int(fileSize))), func(b *testing.B) {
			// Concurrency: 5, 10, 100
			for _, concurrency := range []int{s3manager.DefaultDownloadConcurrency, 2 * s3manager.DefaultDownloadConcurrency, 100} {
				b.Run(fmt.Sprintf("%d Concurrency", concurrency), func(b *testing.B) {
					// PartSize: 5 MB, 25 MB, 100 MB
					for _, partSize := range []int64{s3manager.DefaultDownloadPartSize, 25 * sdkio.MebiByte, 100 * sdkio.MebiByte} {
						if partSize > fileSize {
							continue
						}
						b.Run(fmt.Sprintf("%s PartSize", integration.SizeToName(int(partSize))), func(b *testing.B) {
							for name, strat := range benchStrategies {
								b.Run(name, func(b *testing.B) {
									sdkConfig := baseSdkConfig
									sdkConfig.Concurrency = concurrency
									sdkConfig.PartSize = partSize
									sdkConfig.BufferProvider = strat

									b.ResetTimer()
									for i := 0; i < b.N; i++ {
										benchDownload(b, benchConfig.bucket, key, f, sdkConfig, benchConfig.clientConfig)
										_, err := f.Seek(0, io.SeekStart)
										if err != nil {
											b.Fatalf("failed to seek file back to beginning: %v", err)
										}
									}
								})
							}
						})
					}
				})
			}
		})

		err = teardownDownloadTest(benchConfig.bucket, key)
		if err != nil {
			b.Fatalf("failed to cleanup test file: %v", err)
		}
		if err = f.Close(); err != nil {
			b.Errorf("failed to close file: %v", err)
		}
		if err = os.Remove(f.Name()); err != nil {
			b.Errorf("failed to remove file: %v", err)
		}
	}
}

func benchDownload(b *testing.B, bucket, key string, body io.WriterAt, sdkConfig SDKConfig, clientConfig ClientConfig) {
	downloader := newDownloader(clientConfig, sdkConfig)
	_, err := downloader.Download(body, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		b.Fatalf("failed to download object, %v", err)
	}
}

func TestMain(m *testing.M) {
	benchConfig.SetupFlags("", flag.CommandLine)
	flag.Parse()
	os.Exit(m.Run())
}
