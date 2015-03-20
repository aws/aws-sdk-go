package aws

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/vaughan0/go-ini"
)

var (
	// ErrDefaultRegionNotFound is returned when the AWS Default Region
	// can't be found in the process's environment.
	ErrDefaultRegionNotFound = fmt.Errorf("AWS_DEFAULT_REGION not found in environment")
)

// DetectRegion returns the region based on the available information.
//
// If the AWS_DEFAULT_REGION environment variable is available, it returns that region.
//
// If a profile configuration file is available in the default location and has
// a default profile configured with the region set, it returns that region.
//
// Otherwise, it returns an error
func DetectRegion() (string, error) {
	region, err := EnvRegion()
	if err == nil {
		return region, nil
	}

	region, err = ProfileRegion("", "")
	if err == nil {
		return region, nil
	}

	return "", fmt.Errorf("%s and %s", ErrDefaultRegionNotFound.Error(), err.Error())
}

// EnvRegion returns a string of the Deafult AWS Region from the process's
// environment, or an error if it is not found.
func EnvRegion() (string, error) {
	region := os.Getenv("AWS_DEFAULT_REGION")

	if region == "" {
		return "", ErrDefaultRegionNotFound
	}

	return region, nil
}

// ProfileRegion returns a string of the AWS Region from the profile
// configuration file, or an error if it is not found.
func ProfileRegion(filename, profile string) (string, error) {
	if filename == "" {
		u, err := user.Current()
		if err != nil {
			return "", err
		}

		filename = path.Join(u.HomeDir, ".aws", "config")
	}

	config, err := ini.LoadFile(filename)
	if err != nil {
		return "", err
	}

	if profile == "" {
		profile = "default"
	}

	section := config.Section(profile)
	region, ok := section["region"]
	if !ok {
		return "", fmt.Errorf("profile %s in %s did not contain region", profile, filename)
	}

	return region, nil
}
