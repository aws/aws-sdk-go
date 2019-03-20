package defaults

import (
	"net/http"
	"reflect"
	"time"
)

// HTTP Client default configuration values
var (
	// net/http.DefaultTransport defaults to 0, i.e. wait for ever
	defaultResponseHeaderTimeout = 60 * time.Second

	// net/http.DefaultTransport defaults to DefaultMaxIdleConnsPerHost(2)
	defaultMaxIdleConnsPerHost = 10
)

func getDefaultClient() *http.Client {
	c := copyClient(http.DefaultClient)

	// Apply SDK default overrides to transport.
	if tr, ok := c.Transport.(*http.Transport); ok {
		tr.ResponseHeaderTimeout = defaultResponseHeaderTimeout
		tr.MaxIdleConnsPerHost = defaultMaxIdleConnsPerHost
	}

	return c
}

func copyClient(c *http.Client) *http.Client {
	origClientVal := reflect.ValueOf(c).Elem()

	c = &http.Client{}
	copyStructFields(reflect.ValueOf(c).Elem(), origClientVal)

	// Returned c.Transport will be nil if DefaultClient doesn't have a
	// transport set and the http.DefaultTransport is not a http.Transport.
	if c.Transport == nil {
		if tr, ok := http.DefaultTransport.(*http.Transport); ok {
			c.Transport = copyTransport(tr)
		}
	} else if tr, ok := c.Transport.(*http.Transport); ok {
		c.Transport = copyTransport(tr)
	}

	return c
}

func copyTransport(tr *http.Transport) *http.Transport {
	origTransportVal := reflect.ValueOf(tr).Elem()

	tr = &http.Transport{}
	copyStructFields(reflect.ValueOf(tr).Elem(), origTransportVal)

	return tr
}

func copyStructFields(dst, src reflect.Value) {
	srcType := src.Type()

	for i := 0; i < srcType.NumField(); i++ {
		ft := srcType.Field(i)
		if len(ft.PkgPath) != 0 {
			// Unexported fields include package path. Exported fields do not.
			// https://godoc.org/reflect#StructField
			continue
		}

		dst.Field(i).Set(src.Field(i))
	}
}
