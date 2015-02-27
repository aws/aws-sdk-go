package util

import (
	"fmt"
	"go/format"
	"regexp"
)

func GoFmt(buf string) string {
	formatted, err := format.Source([]byte(buf))
	if err != nil {
		panic(fmt.Errorf("%s\nOriginal code:\n%s", err.Error(), buf))
	}
	return string(formatted)
}

var reTrim = regexp.MustCompile("\\s")

func Trim(s string) string {
	return reTrim.ReplaceAllString(s, "")
}
