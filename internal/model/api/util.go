package api

import "go/format"

func gofmt(buf string) string {
	formatted, err := format.Source([]byte(buf))
	if err != nil {
		panic(err)
	}
	return string(formatted)
}
