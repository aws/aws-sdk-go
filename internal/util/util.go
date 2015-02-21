package util

import (
	"go/format"
	"regexp"
)

func GoFmt(buf string) string {
	formatted, err := format.Source([]byte(buf))
	if err != nil {
		panic(err)
	}
	return string(formatted)
}

var reTrim = regexp.MustCompile("\\s")

func Trim(s string) string {
	return reTrim.ReplaceAllString(s, "")
}
