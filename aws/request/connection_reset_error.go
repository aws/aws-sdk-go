package request

import (
	"strings"
)

func isErrConnectionReset(err error) bool {
	if strings.Contains(err.Error(), "connection reset") ||
		strings.Contains(err.Error(), "broken pipe") {
		return true
	}

	return false
}
