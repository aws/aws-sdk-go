// +build example

package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	log.SetFlags(0)

	var args struct {
		region            string // optional
		copySource        string // required
		srcRegion         string // optional
		dstBucket         string // required
		dstKey            string // required
		metadataDirective string // optional

		copierMaxPartSize      int64 // optional
		copierThreshold        int64 // optional
		copierConcurrency      int   // optional
		copierNoDiscoverRegion bool  // optional
	}

	flag.StringVar(&args.region, "region", "", "s3 client region (optional)")
	flag.StringVar(&args.copySource, "copy-source", "", "copy source (required)")
	flag.StringVar(&args.srcRegion, "src-region", "", "source region (optional)")
	flag.StringVar(&args.dstBucket, "dst-bucket", "", "destination bucket (required)")
	flag.StringVar(&args.dstKey, "dst-key", "", "destination key (required)")
	flag.StringVar(&args.metadataDirective, "metadata-directive", "", "metadata directive (optional)")
	flag.Int64Var(&args.copierMaxPartSize, "copier-max-part-size", s3manager.MaxUploadPartSize, "copier max part size (optional)")
	flag.Int64Var(&args.copierThreshold, "copier-threshold", s3manager.DefaultMultipartCopyThreshold, "copier multipart threshold (optional)")
	flag.IntVar(&args.copierConcurrency, "copier-concurrency", s3manager.DefaultCopyConcurrency, "copier concurrency (optional)")
	flag.BoolVar(&args.copierNoDiscoverRegion, "copier-no-discover-region", false, "should the copier automatically discover the region? (optional)")
	flag.Parse()

	copier := newCopier(args.region)

	if args.copierNoDiscoverRegion {
		copier.DiscoverSourceBucketRegion = false
	}

	input := s3manager.CopyInput{
		Bucket:     aws.String(args.dstBucket),
		CopySource: aws.String(args.copySource),
		Key:        aws.String(args.dstKey),
	}

	if args.srcRegion != "" {
		input.SourceRegion = aws.String(args.srcRegion)
	}

	if args.metadataDirective != "" {
		input.MetadataDirective = aws.String(args.metadataDirective)
	}

	out, err := copier.Copy(&input)
	if err != nil {
		log.Fatalf("copy failed: %v", err)
	}

	if out.CopySourceVersionId != nil {
		log.Printf("CopySourceVersionId: %s", aws.StringValue(out.CopySourceVersionId))
	}

	log.Printf("ETag: %s", aws.StringValue(out.ETag))
	log.Printf("Location: %s", aws.StringValue(out.Location))
	log.Printf("VersionId: %s", aws.StringValue(out.VersionId))
}

func newCopier(region string) *s3manager.Copier {
	opts := session.Options{
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}

	if region != "" {
		opts.Config.Region = aws.String(region)
	}

	sess := session.Must(session.NewSessionWithOptions(opts))
	return s3manager.NewCopierWithClient(s3.New(sess))
}
