// +build integration

package cloudsearch_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/cloudsearch"
)

func TestCreateDomains(t *testing.T) {

	name := "awsgo" + strconv.FormatInt(time.Now().UnixNano(), 10)
	client := cloudsearch.New(aws.DefaultCreds(), "us-east-1", nil)

	createOutput, err := client.CreateDomain(&cloudsearch.CreateDomainRequest{
		DomainName: &name,
	})
	if err != nil {
		t.Fatal("Failed to create domain: ", err)
	}
	defer deleteDomain(t, client, &name)

	if *(createOutput.DomainStatus.ARN) == "" {
		t.Fatal("Create response not marshalled correctly")
	}
}

func deleteDomain(t *testing.T, client *cloudsearch.CloudSearch, name *string) {

	output, err := client.DeleteDomain(&cloudsearch.DeleteDomainRequest{
		DomainName: name,
	})
	if err != nil {
		t.Fatal("Failed to delete cloud search domain: ", err)
	}
	if *(output.DomainStatus.ARN) == "" {
		t.Fatal("Delete response not marshalled correctly")
	}
}
