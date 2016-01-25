package v4

// filter is an interface that will check to see if a value
// needs to be filtered
type filter interface {
	Allow(value string) bool
}

// whiltelistFilter takes a whitelist as a field and uses that to
// check for allowance
type whitelistFilter struct {
	whitelist map[string]struct{}
}

// blacklistFilter takes a whitelist as a field and uses that to
// check for allowance
type blacklistFilter struct {
	blacklist map[string]struct{}
}

// Allow check for allowance of a key
func (a whitelistFilter) Allow(value string) bool {
	_, ok := a.whitelist[value]
	return ok
}

// Allow check for allowance of a key
func (a blacklistFilter) Allow(value string) bool {
	_, ok := a.blacklist[value]
	return !ok
}
