package session

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/stretchr/testify/assert"
)

var (
	testConfigFilename      = filepath.Join("testdata", "shared_config")
	testConfigOtherFilename = filepath.Join("testdata", "shared_config_other")
)

func TestLoadSharedConfigFromFile(t *testing.T) {
	cases := []struct {
		Filename    string
		Profile     string
		Expected    sharedConfig
		ErrContains string
	}{
		{
			Filename: testConfigFilename, Profile: "default",
			Expected: sharedConfig{Region: "default_region"},
		},
		{
			Filename: testConfigFilename, Profile: "alt_profile_name",
			Expected: sharedConfig{Region: "alt_profile_name_region"},
		},
		{
			Filename: testConfigFilename, Profile: "short_profile_name_first",
			Expected: sharedConfig{Region: "short_profile_name_first_short"},
		},
		{
			Filename: testConfigFilename, Profile: "partial_creds",
			Expected: sharedConfig{},
		},
		{
			Filename: testConfigFilename, Profile: "complete_creds",
			Expected: sharedConfig{
				Creds: credentials.Value{
					AccessKeyID:     "complete_creds_akid",
					SecretAccessKey: "complete_creds_secret",
					ProviderName:    fmt.Sprintf("SharedConfigCredentials: %s", testConfigFilename),
				},
			},
		},
		{
			Filename: testConfigFilename, Profile: "complete_creds_with_token",
			Expected: sharedConfig{
				Creds: credentials.Value{
					AccessKeyID:     "complete_creds_with_token_akid",
					SecretAccessKey: "complete_creds_with_token_secret",
					SessionToken:    "complete_creds_with_token_token",
					ProviderName:    fmt.Sprintf("SharedConfigCredentials: %s", testConfigFilename),
				},
			},
		},
		{
			Filename: testConfigFilename, Profile: "full_profile",
			Expected: sharedConfig{
				Creds: credentials.Value{
					AccessKeyID:     "full_profile_akid",
					SecretAccessKey: "full_profile_secret",
					ProviderName:    fmt.Sprintf("SharedConfigCredentials: %s", testConfigFilename),
				},
				Region: "full_profile_region",
			},
		},
	}

	for _, c := range cases {
		cfg := sharedConfig{}

		err := cfg.setFromFile(c.Profile, c.Filename)
		if len(c.ErrContains) > 0 {
			assert.Error(t, err)
			assert.Contains(t, err, c.ErrContains)
			continue
		}

		assert.NoError(t, err)
	}
}

func TestLoadSharedConfig(t *testing.T) {
	cases := []struct {
		Filenames   []string
		Profile     string
		Expected    sharedConfig
		ErrContains string
	}{
		{
			Filenames: []string{"file_not_exists"},
			Profile:   "default",
		},
		{
			Filenames: []string{testConfigFilename},
			Expected: sharedConfig{
				Region: "default_region",
			},
		},
		{
			Filenames: []string{testConfigOtherFilename, testConfigFilename},
			Profile:   "config_file_load_order",
			Expected: sharedConfig{
				Region: "shared_config_region",
				Creds: credentials.Value{
					AccessKeyID:     "shared_config_akid",
					SecretAccessKey: "shared_config_secret",
					ProviderName:    "SharedConfigCredentials",
				},
			},
		},
		{
			Filenames: []string{testConfigFilename, testConfigOtherFilename},
			Profile:   "config_file_load_order",
			Expected: sharedConfig{
				Region: "shared_config_other_region",
				Creds: credentials.Value{
					AccessKeyID:     "shared_config_other_akid",
					SecretAccessKey: "shared_config_other_secret",
					ProviderName:    "SharedConfigCredentials",
				},
			},
		},
	}

	for _, c := range cases {
		cfg, err := loadSharedConfig(c.Profile, c.Filenames...)
		if len(c.ErrContains) > 0 {
			assert.Error(t, err)
			assert.Contains(t, err, c.ErrContains)
			continue
		}

		assert.NoError(t, err)

		assert.Contains(t, cfg.Creds.ProviderName, c.Expected.Creds.ProviderName)
		cfg.Creds.ProviderName = c.Expected.Creds.ProviderName

		assert.Equal(t, c.Expected, cfg)
	}
}
