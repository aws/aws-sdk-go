package util

import "go/format"

func GoFmt(buf string) string {
	formatted, err := format.Source([]byte(buf))
	if err != nil {
		panic(err)
	}
	return string(formatted)
}
