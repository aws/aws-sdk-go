package s3crypto

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Finish SHA256 and add MD5 tests
// From Go stdlib encoding/sha256 test cases
func TestSHA256(t *testing.T) {
	sha := newSHA256Writer(nil)
	expected, _ := hex.DecodeString("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	b := sha.GetValue()
	assert.Equal(t, expected, b)
}
