package rest

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanPath(t *testing.T) {
	uri := &url.URL{
		Path:   "//foo//bar",
		Scheme: "https",
		Host:   "host",
	}
	cleanPath(uri)

	expected := "https://host/foo/bar"
	assert.Equal(t, expected, uri.String())
}
