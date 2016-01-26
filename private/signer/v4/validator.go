package v4

import (
	"net/http"
	"strings"
)

// validator houses a set of rule needed for validation of a
// string value
type validator struct {
	rules []rule
}

// rule interface allows for more flexible rules and just simply
// checks whether or not a value adheres to that rule
type rule interface {
	IsValid(value string) bool
}

// validate will iterate through all rules and see if any rules
// apply to the value
func (v validator) Validate(value string) bool {
	for _, rule := range v.rules {
		if rule.IsValid(value) {
			return true
		}
	}
	return false
}

type whitelist map[string]struct{}
type blacklist map[string]struct{}
type patterns []string

// IsValid for whitelist checks if the value is within the whitelist
func (wlist whitelist) IsValid(value string) bool {
	_, ok := wlist[value]
	return ok
}

// IsValid for whitelist checks if the value is within the whitelist
func (blist blacklist) IsValid(value string) bool {
	_, ok := blist[value]
	return !ok
}

// IsValid for patterns checks each pattern and returns if a match has
// been found
func (p patterns) IsValid(value string) bool {
	for _, pattern := range p {
		if strings.HasPrefix(http.CanonicalHeaderKey(value), pattern) {
			return true
		}
	}
	return false
}
