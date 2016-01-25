package v4

// allower is an interface that will check to see if a value
// is allowed
type allower interface {
	allow(value string) bool
}

// whitelistAllower takes a whitelist as a field and uses that to
// check for allowance
type whitelistAllower struct {
	whitelist map[string]struct{}
}

// blacklistAllower takes a whitelist as a field and uses that to
// check for allowance
type blacklistAllower struct {
	blacklist map[string]struct{}
}

// allow check for allowance of a key
func (allower whitelistAllower) allow(value string) bool {
	_, ok := allower.whitelist[value]
	return ok
}

// allow check for allowance of a key
func (allower blacklistAllower) allow(value string) bool {
	_, ok := allower.blacklist[value]
	return !ok
}
