package awsutil

import (
	"unicode"
	"unicode/utf8"
)

// StringHasPrefixFold tests whether the string s begins with prefix, interpreted as UTF-8 strings,
// under Unicode case-folding.
//
// This implementation is a modified version of the Go Standard Library's strings.EqualFold
// Copyright 2009 The Go Authors. All rights reserved.
func StringHasPrefixFold(s, prefix string) bool {
	for s != "" && prefix != "" {
		// Extract first rune from each string.
		var sr, tr rune
		if s[0] < utf8.RuneSelf {
			sr, s = rune(s[0]), s[1:]
		} else {
			r, size := utf8.DecodeRuneInString(s)
			sr, s = r, s[size:]
		}
		if prefix[0] < utf8.RuneSelf {
			tr, prefix = rune(prefix[0]), prefix[1:]
		} else {
			r, size := utf8.DecodeRuneInString(prefix)
			tr, prefix = r, prefix[size:]
		}

		// If they match, keep going; if not, return false.

		// Easy case.
		if tr == sr {
			continue
		}

		// Make sr < tr to simplify what follows.
		if tr < sr {
			tr, sr = sr, tr
		}
		// Fast check for ASCII.
		if tr < utf8.RuneSelf {
			// ASCII only, sr/tr must be upper/lower case
			if 'A' <= sr && sr <= 'Z' && tr == sr+'a'-'A' {
				continue
			}
			return false
		}

		// General case. SimpleFold(x) returns the next equivalent rune > x
		// or wraps around to smaller values.
		r := unicode.SimpleFold(sr)
		for r != sr && r < tr {
			r = unicode.SimpleFold(r)
		}
		if r == tr {
			continue
		}
		return false
	}

	if len(s) == 0 && len(prefix) > 0 {
		return false
	}

	return true
}
