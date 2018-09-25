// +build !go1.7

package sign

import (
	"encoding/json"
)

func encodePolicyJSON(p *Policy) ([]byte, error) {
	return json.Marshal(p)
}
