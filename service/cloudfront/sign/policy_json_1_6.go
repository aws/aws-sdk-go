// +build !go1.7

package sign

import (
	"bytes"
	"encoding/json"
	"log"
)

func encodePolicyJSON(p *Policy) ([]byte, error) {
	buffer := &bytes.Buffer{}
	src, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	start := 0
	for i, c := range src {
		// Convert \u0026 back to &
		if c == '\\' && i+5 < len(src) && src[i+1] == 'u' && src[i+2] == '0' && src[i+3] == '0' && src[i+4] == '2' && src[i+5] == '6' {
			if start < i {
				buffer.Write(src[start:i])
			}
			buffer.WriteString("&")
			start = i + 6
		}
		if start < i {
			buffer.Write(src[start:i])
			start = start + 1
		}
		log.Printf("%+q", c)
	}
	if start < len(src) {
		buffer.Write(src[start:])
	}
	return buffer.Bytes(), err
}
