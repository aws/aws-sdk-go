package v4

import (
	"net/http"
	"strings"
)

// filter is an interface that will check to see if a value
// needs to be filtered
type filter interface {
	Allow(value string) bool
}

// whiltelistFilter takes a whitelist as a field and uses that to
// check for allowance
type whitelistFilter struct {
	whitelist map[string]bool
}

// blacklistFilter takes a whitelist as a field and uses that to
// check for allowance
type blacklistFilter struct {
	blacklist map[string]struct{}
}

// Allow check for allowance of a key
func (a whitelistFilter) Allow(value string) bool {
	if _, ok := a.whitelist[value]; ok {
		return true
	}
	for k, v := range a.whitelist {
		if v {
			if strings.HasPrefix(http.CanonicalHeaderKey(value), k) {
				return true
			}
		}
	}
	return false
}

// Allow check for allowance of a key
func (a blacklistFilter) Allow(value string) bool {
	_, ok := a.blacklist[value]
	return !ok
}
