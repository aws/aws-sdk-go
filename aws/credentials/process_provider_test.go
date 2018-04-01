package credentials

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessProvider(t *testing.T) {
	os.Clearenv()

	p := ProcessProvider{Filename: "example.ini", Profile: "process", executionFunc: executeCredentialProcess}
	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "accessKey", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "secret", creds.SecretAccessKey, "Expect secret access key to match")
	assert.Equal(t, "tokenProcess", creds.SessionToken, "Expect session token to match")
}

func fakeExectuteCredsExpired(process string) ([]byte, error) {
	return []byte(`{"Version": 1, "AccessKeyId": "accessKey", "SecretAccessKey": "secret", "SessionToken": "tokenDefault", "Expiration": "2000-01-01T00:00:00-00:00"}`), nil
}

func TestProcessProviderIsExpired(t *testing.T) {
	os.Clearenv()

	p := ProcessProvider{Filename: "example.ini", Profile: "process", executionFunc: fakeExectuteCredsExpired}

	assert.True(t, p.IsExpired(), "Expect creds to be expired before retrieve")
}

func TestProcessProviderWithAWS_CONFIG_FILE(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_CONFIG_FILE", "example.ini")
	os.Setenv("AWS_DEFAULT_PROFILE", "process")
	p := ProcessProvider{Filename: "", Profile: "", executionFunc: executeCredentialProcess}
	creds, err := p.Retrieve()

	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "accessKey", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "secret", creds.SecretAccessKey, "Expect secret access key to match")
	assert.Equal(t, "tokenProcess", creds.SessionToken, "Expect session token to match")
}

func TestProcessProviderWithAWS_CONFIG_FILEAbsPath(t *testing.T) {
	os.Clearenv()
	wd, err := os.Getwd()
	assert.NoError(t, err)
	os.Setenv("AWS_CONFIG_FILE", filepath.Join(wd, "example.ini"))
	p := ProcessProvider{executionFunc: executeCredentialProcess}
	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "accessKey", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "secret", creds.SecretAccessKey, "Expect secret access key to match")
	assert.Equal(t, "tokenDefault", creds.SessionToken, "Expect session token to match")
}

func fakeExectuteCredsSuccess(process string) ([]byte, error) {
	return []byte(`{"Version": 1, "AccessKeyId": "accessKey", "SecretAccessKey": "secret", "SessionToken": "tokenFake", "Expiration": "2000-01-01T00:00:00-00:00"}`), nil
}

func TestProcessProviderWithAWS_PROFILE(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_PROFILE", "process")

	p := ProcessProvider{Filename: "example.ini", Profile: "", executionFunc: fakeExectuteCredsSuccess}
	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "accessKey", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "secret", creds.SecretAccessKey, "Expect secret access key to match")
	assert.Equal(t, "tokenFake", creds.SessionToken, "Expect token to match")
}

func fakeExectuteCredsFailMalformed(process string) ([]byte, error) {
	return []byte(`{"Version": 1, "AccessKeyId": "accessKey", "SecretAccessKey": "secret", "SessionToken": "tokenDefault", "Expiration": `), nil
}

func TestProcessProviderMalformed(t *testing.T) {
	os.Clearenv()
	os.Setenv("AWS_PROFILE", "process")

	p := ProcessProvider{Filename: "example.ini", Profile: "", executionFunc: fakeExectuteCredsFailMalformed}
	_, err := p.Retrieve()
	assert.NotNil(t, err, "Expect an error")
}

func fakeExectuteCredsNoToken(process string) ([]byte, error) {
	return []byte(`{"Version": 1, "AccessKeyId": "accessKey", "SecretAccessKey": "secret"}`), nil
}

func TestProcessProviderNoToken(t *testing.T) {
	os.Clearenv()

	p := ProcessProvider{Filename: "example.ini", Profile: "process", executionFunc: fakeExectuteCredsNoToken}
	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")
	assert.Empty(t, creds.SessionToken, "Expect no token")
}

func fakeExectuteCredsFailVersion(process string) ([]byte, error) {
	return []byte(`{"Version": 2, "AccessKeyId": "accessKey", "SecretAccessKey": "secret", "SessionToken": "tokenDefault"}`), nil
}

func TestProcessProviderWrongVersion(t *testing.T) {
	os.Clearenv()
	p := ProcessProvider{Filename: "example.ini", Profile: "process", executionFunc: fakeExectuteCredsFailVersion}
	_, err := p.Retrieve()
	assert.NotNil(t, err, "Expect an error")
}

func fakeExectuteCredsFailExpiration(process string) ([]byte, error) {
	return []byte(`{"Version": 1, "AccessKeyId": "accessKey", "SecretAccessKey": "secret", "SessionToken": "tokenDefault", "Expiration": "20222"}`), nil
}
func TestProcessProviderBadExpiry(t *testing.T) {
	os.Clearenv()
	p := ProcessProvider{Filename: "example.ini", Profile: "process", executionFunc: fakeExectuteCredsFailExpiration}
	_, err := p.Retrieve()
	assert.NotNil(t, err, "Expect an error")
}

func BenchmarkProcessProvider(b *testing.B) {
	os.Clearenv()

	p := ProcessProvider{Filename: "example.ini", Profile: "process", executionFunc: executeCredentialProcess}
	_, err := p.Retrieve()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Retrieve()
		if err != nil {
			b.Fatal(err)
		}
	}
}
