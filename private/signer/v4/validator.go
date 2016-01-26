package v4

import (
	"net/http"
	"strings"
)

// Validator houses a set of rule needed for validation of a
// string value
type Validator struct {
	rules []Rule
}

// Rule interface allows for more flexible rules and just simply
// checks whether or not a value adheres to that rule
type Rule interface {
	IsValid(value string) bool
}

// Validate will iterate through all rules and see if any rules
// apply to the value
func (v Validator) Validate(value string) bool {
	for _, rule := range v.rules {
		if rule.IsValid(value) {
			return true
		}
	}
	return false
}

// AddRule adds a new rule to the validator
func (v *Validator) AddRule(r Rule) {
	v.rules = append(v.rules, r)
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
