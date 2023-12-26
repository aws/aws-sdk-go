//go:build example
// +build example

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("USAGE ERROR: go run -tags example putObjWithProcess.go <bucket> <key for object>")
		return
	}

	const filename = "./base_object.txt"

	bucket := os.Args[1]
	key := os.Args[2]

	myCustomResolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == endpoints.S3ServiceID {
			return endpoints.ResolvedEndpoint{
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}

		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("ru-central1"),
		EndpointResolver: endpoints.ResolverFunc(myCustomResolver),
	}))

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %v, %v", filename, err)
	}

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024
		u.LeavePartsOnError = true
	})

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("failed to put file %v, %v", filename, err)
		return
	}
	fmt.Println()
	log.Println(output.Location)

	const patchFile = "./insertion.txt"

	file, err = os.Open(patchFile)
	if err != nil {
		log.Fatalf("failed to open file %v, %v", patchFile, err)
	}

	input := &s3.PatchObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(key),
		Body:          file,
		ContentRange:  aws.String("bytes 14-23/*"),
		ContentLength: aws.Int64(10),
	}

	assTree := s3.New(sess)
	patchOutput, err := assTree.PatchObject(input)
	if err != nil {
		log.Fatalf("could not patch: %s", err)
	}

	log.Printf("output: %#v\n", patchOutput)
}
