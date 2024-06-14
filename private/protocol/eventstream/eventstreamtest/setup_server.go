//go:build !go1.10
// +build !go1.10

package eventstreamtest

import (
	"net/http"
	"net/http/httptest"
)

func setupServer(server *httptest.Server, useH2 bool) *http.Client {
	server.Start()

	return nil
}
