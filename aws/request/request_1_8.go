// +build go1.8

package request

import "io"

// Is a nil instructing Go HTTP client to not include
// and body in the HTTP request.
var noBodyReader io.ReadCloser
