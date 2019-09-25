// Package behaviortests contains automatically generated AWS clients for behavior testing.
package behaviortests

//go:generate go run -tags codegen ../../private/model/cli/gen-api/main.go -path=../../awstesting/behavior -svc-import-path=github.com/aws/aws-sdk-go/awstesting/behavior ./*/*/api-2.json
//go:generate gofmt -s -w ../../awstesting/behavior
