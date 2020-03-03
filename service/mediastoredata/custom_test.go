// +build go1.8

package mediastoredata_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastoredata"
)

func Test_DescribeEndpoint(t *testing.T) {
	const containerName = "awsgosdkteamintegcontainer"

	sess := unit.Session
	if v := aws.StringValue(sess.Config.Region); len(v) == 0 {
		sess.Config.Region = aws.String("us-east-1")
	}

	ctrlSvc := mediastore.New(sess)
	input := &mediastore.DescribeContainerInput{
		ContainerName: aws.String(containerName),
	}
	req, _ := ctrlSvc.DescribeContainerRequest(input)
	if e, a := "/", req.HTTPRequest.URL.Path; e != a {
		t.Fatalf("expected %v, got %v", e, a)
	}
	if e, a := input, req.Params; e != a {
		t.Fatalf("expected %v, got %v", e, a)
	}

	dataSvc := mediastoredata.New(sess, &aws.Config{
		Endpoint: aws.String("mockEndpoint"),
	})
	req, _ = dataSvc.ListItemsRequest(&mediastoredata.ListItemsInput{})
	if e, a := "mockEndpoint", req.HTTPRequest.URL.Host; e != a {
		t.Fatalf("expected %v, got %v", e, a)
	}
	if e, a := "/", req.HTTPRequest.URL.Path; e != a {
		t.Fatalf("expected %v, got %v", e, a)
	}

}
