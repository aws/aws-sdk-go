package credentials

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharedCredentialsProviderFromHomeDirectory(t *testing.T) {
	os.Clearenv()
	wd, err := os.Getwd()
	assert.NoError(t, err)
	os.Setenv("HOME", filepath.Join(wd, "testhomedir"))
	os.Setenv("USERPROFILE", filepath.Join(wd, "whatwouldbetesthomedironwindows"))
	p := SharedCredentialsProvider{}
	creds, err := p.Retrieve()

	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "homeDirAccessKey", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "homeDirSecret", creds.SecretAccessKey, "Expect secret access key to match")
	assert.Equal(t, "homeDirToken", creds.SessionToken, "Expect session token to match")
}
