package ec2metadata_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
)

func initTestServer(path string, resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("TestGetRegion: URL:", r.URL.String())
		if r.RequestURI != path {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		w.Write([]byte(resp))
	}))
}

func TestGetMetadata(t *testing.T) {
	server := initTestServer(
		"/latest/meta-data/some/path",
		"success", // real response includes suffix
	)
	defer server.Close()
	c := ec2metadata.New(aws.NewConfig().WithEndpoint(server.URL + "/latest"))

	resp, err := c.GetMetadata("some/path")

	assert.NoError(t, err)
	assert.Equal(t, "success", resp)
}
func TestGetRegion(t *testing.T) {
	server := initTestServer(
		"/latest/meta-data/placement/availability-zone",
		"us-west-2a", // real response includes suffix
	)
	defer server.Close()
	c := ec2metadata.New(aws.NewConfig().WithEndpoint(server.URL + "/latest"))

	region, err := c.Region()

	assert.NoError(t, err)
	assert.Equal(t, "us-west-2", region)
}

func TestMetadataAvailable(t *testing.T) {
	server := initTestServer(
		"/latest/meta-data/instance-id",
		"instance-id",
	)
	defer server.Close()
	c := ec2metadata.New(aws.NewConfig().WithEndpoint(server.URL + "/latest"))

	available := c.Available()

	assert.True(t, available)
}
func TestMetadataNotAvailable(t *testing.T) {
	c := ec2metadata.New(aws.NewConfig())

	available := c.Available()

	assert.False(t, available)
}
