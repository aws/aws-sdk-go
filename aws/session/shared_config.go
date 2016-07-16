package session

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-ini/ini"
)

const (
	accessKeyIDKey  = `aws_access_key_id`
	secretAccessKey = `aws_secret_access_key`
	sessionTokenKey = `aws_session_token`

	regionKey = `region`

	// DefaultSharedConfigProfile is the default profile to be used when
	// loading configuration from the shared configuration files if another
	// profile name is not provided.
	DefaultSharedConfigProfile = `default`
)

// sharedConfig represents the configuration fields of the shared configuration
// files.
type sharedConfig struct {
	// Credentials values from the shared configuration file. Both aws_access_key_id
	// and aws_secret_access_key must be provided together in the same file
	// to be considered valid. The values will be ignored if not a complete group.
	// aws_session_token is an optional field that can be provided if both of the
	// other two fields are also provided.
	//
	//	aws_access_key_id
	//	aws_secret_access_key
	//	aws_session_token
	Creds credentials.Value

	// Region is the region the SDK should use for looking up AWS service endpoints
	// and signing requests.
	//
	//	region
	Region string
}

// loadSharedConfig retrieves the shared configuration from the list of files
// using the profile provided. The order the files are listed will determine
// precedence. Values in subsequent files will overwrite values defined in
// earlier files.
//
// For example, given two files A and B. Both define credentials. If the order
// of the files are A then B, B's credential values will be used instead of A's.
//
// See sharedConfig.setFromFile for information how the shared config files
// will be loaded.
func loadSharedConfig(profile string, configFiles ...string) (sharedConfig, error) {
	if len(profile) == 0 {
		profile = DefaultSharedConfigProfile
	}

	cfg := sharedConfig{}
	for _, filename := range configFiles {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			// Ignore config files that don't exist.
			continue
		}

		if err := cfg.setFromFile(profile, filename); err != nil {
			return sharedConfig{}, err
		}
	}

	return cfg, nil
}

// setFromFile loads the shared configuration from the file using
// the profile provided. A sharedConfig pointer type value is used so that
// multiple shared config file loadings can be chained.
//
// Only loads complete logically grouped values, and will not set fields in cfg
// for incomplete grouped values in the config. Such as credentials. For example
// if a config file only includes aws_access_key_id but no aws_secret_access_key
// the aws_access_key_id will be ignored.
func (cfg *sharedConfig) setFromFile(profile, filename string) error {
	file, err := ini.Load(filename)
	if err != nil {
		return awserr.New("LoadSharedConfigError",
			"failed to load shared config file.", err)
	}

	section, err := file.GetSection(profile)
	if err != nil {
		// Fallback to to alternate profile name: profile %s
		section, err = file.GetSection(fmt.Sprintf("profile %s", profile))
		if err != nil {
			return awserr.New("LoadSharedConfigError",
				fmt.Sprintf("failed to get profile %s.", profile), err)
		}
	}

	// Credentials
	creds := credentials.Value{
		AccessKeyID:     section.Key(accessKeyIDKey).String(),
		SecretAccessKey: section.Key(secretAccessKey).String(),
		SessionToken:    section.Key(sessionTokenKey).String(),
		ProviderName:    fmt.Sprintf("SharedConfigCredentials: %s", filename),
	}

	// Require logical grouping of credentials
	if len(creds.AccessKeyID) > 0 && len(creds.SecretAccessKey) > 0 {
		cfg.Creds = creds
	}

	// Region
	if v := section.Key(regionKey).String(); len(v) > 0 {
		cfg.Region = v
	}

	return nil
}
