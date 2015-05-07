package s3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostedZoneIDForRegion(t *testing.T) {
	assert.Equal(t, "Z3AQBSTGFYJSTF", HostedZoneIDForRegion("us-east-1"))
	assert.Equal(t, "Z1WCIGYICN2BYD", HostedZoneIDForRegion("ap-southeast-2"))

	// Empty string should default to us-east-1
	assert.Equal(t, "Z3AQBSTGFYJSTF", HostedZoneIDForRegion(""))

	// Bad input should be empty string
	assert.Equal(t, "", HostedZoneIDForRegion("not-a-region"))
}
