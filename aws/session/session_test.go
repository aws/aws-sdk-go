package session

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
)

func TestNewDefaultSession(t *testing.T) {
	env := stashEnv()
	defer popEnv(env)

	s := New(&aws.Config{Region: aws.String("region")})

	assert.Equal(t, "region", *s.Config.Region)
	assert.Equal(t, http.DefaultClient, s.Config.HTTPClient)
	assert.NotNil(t, s.Config.Logger)
	assert.Equal(t, aws.LogOff, *s.Config.LogLevel)
}

func TestNew_WithCustomCreds(t *testing.T) {
	env := stashEnv()
	defer popEnv(env)

	customCreds := credentials.NewStaticCredentials("AKID", "SECRET", "TOKEN")
	s := New(&aws.Config{Credentials: customCreds})

	assert.Equal(t, customCreds, s.Config.Credentials)
}

func TestSessionCopy(t *testing.T) {
	env := stashEnv()
	defer popEnv(env)

	os.Setenv("AWS_REGION", "orig_region")

	s := Session{
		Config:   defaults.Config(),
		Handlers: defaults.Handlers(),
	}

	newSess := s.Copy(&aws.Config{Region: aws.String("new_region")})

	assert.Equal(t, "orig_region", *s.Config.Region)
	assert.Equal(t, "new_region", *newSess.Config.Region)
}

func TestSessionClientConfig(t *testing.T) {
	s := New(&aws.Config{Region: aws.String("orig_region")})

	cfg := s.ClientConfig("s3", &aws.Config{Region: aws.String("us-west-2")})

	assert.Equal(t, "https://s3-us-west-2.amazonaws.com", cfg.Endpoint)
	assert.Empty(t, cfg.SigningRegion)
	assert.Equal(t, "us-west-2", *cfg.Config.Region)
}

func TestNewFromProfile(t *testing.T) {
	env := stashEnv()
	defer popEnv(env)

	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "other_profile")

	s := NewFromProfile("full_profile")

	assert.Empty(t, *s.Config.Region)

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "full_profile_akid", creds.AccessKeyID)
	assert.Equal(t, "full_profile_secret", creds.SecretAccessKey)
	assert.Empty(t, creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "SharedConfigCredentials")
}

func TestNewFromProfile_WithSharedConfig(t *testing.T) {
	env := stashEnv()
	defer popEnv(env)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "other_profile")

	s := NewFromProfile("full_profile")

	assert.Equal(t, "full_profile_region", *s.Config.Region)

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "full_profile_akid", creds.AccessKeyID)
	assert.Equal(t, "full_profile_secret", creds.SecretAccessKey)
	assert.Empty(t, creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "SharedConfigCredentials")
}

func TestNewFromSharedConfig(t *testing.T) {
	env := stashEnv()
	defer popEnv(env)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "0")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "full_profile")

	s := NewFromSharedConfig()

	assert.Equal(t, "full_profile_region", *s.Config.Region)

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "full_profile_akid", creds.AccessKeyID)
	assert.Equal(t, "full_profile_secret", creds.SecretAccessKey)
	assert.Empty(t, creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "SharedConfigCredentials")
}

func TestNewFromSharedConfigProfile(t *testing.T) {
	cases := []struct {
		InEnvs    map[string]string
		InProfile string
		OutRegion string
		OutCreds  credentials.Value
	}{
		{
			InEnvs: map[string]string{
				"AWS_SDK_LOAD_CONFIG":         "0",
				"AWS_SHARED_CREDENTIALS_FILE": testConfigFilename,
				"AWS_PROFILE":                 "other_profile",
			},
			InProfile: "full_profile",
			OutRegion: "full_profile_region",
			OutCreds: credentials.Value{
				AccessKeyID:     "full_profile_akid",
				SecretAccessKey: "full_profile_secret",
				ProviderName:    "SharedConfigCredentials",
			},
		},
		{
			InEnvs: map[string]string{
				"AWS_SDK_LOAD_CONFIG":         "0",
				"AWS_SHARED_CREDENTIALS_FILE": testConfigFilename,
				"AWS_REGION":                  "env_region",
				"AWS_ACCESS_KEY":              "env_akid",
				"AWS_SECRET_ACCESS_KEY":       "env_secret",
				"AWS_PROFILE":                 "other_profile",
			},
			InProfile: "full_profile",
			OutRegion: "env_region",
			OutCreds: credentials.Value{
				AccessKeyID:     "env_akid",
				SecretAccessKey: "env_secret",
				ProviderName:    "EnvConfigCredentials",
			},
		},
		{
			InEnvs: map[string]string{
				"AWS_SDK_LOAD_CONFIG":         "0",
				"AWS_SHARED_CREDENTIALS_FILE": testConfigFilename,
				"AWS_SHARED_CONFIG_FILE":      testConfigOtherFilename,
				"AWS_PROFILE":                 "shared_profile",
			},
			InProfile: "config_file_load_order",
			OutRegion: "shared_config_region",
			OutCreds: credentials.Value{
				AccessKeyID:     "shared_config_akid",
				SecretAccessKey: "shared_config_secret",
				ProviderName:    "SharedConfigCredentials",
			},
		},
	}

	for _, c := range cases {
		env := stashEnv()
		defer popEnv(env)

		for k, v := range c.InEnvs {
			os.Setenv(k, v)
		}

		s := NewFromSharedConfigProfile(c.InProfile)

		creds, err := s.Config.Credentials.Get()
		assert.NoError(t, err)
		assert.Equal(t, c.OutRegion, *s.Config.Region)
		assert.Equal(t, c.OutCreds.AccessKeyID, creds.AccessKeyID)
		assert.Equal(t, c.OutCreds.SecretAccessKey, creds.SecretAccessKey)
		assert.Equal(t, c.OutCreds.SessionToken, creds.SessionToken)
		assert.Contains(t, creds.ProviderName, c.OutCreds.ProviderName)
	}
}
