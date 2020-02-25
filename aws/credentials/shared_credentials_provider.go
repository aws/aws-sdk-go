package credentials

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/internal/ini"
	"github.com/aws/aws-sdk-go/internal/shareddefaults"
)

// SharedCredsProviderName provides a name of SharedCreds provider
const SharedCredsProviderName = "SharedCredentialsProvider"

var (
	// ErrSharedCredentialsHomeNotFound is emitted when the user directory cannot be found.
	ErrSharedCredentialsHomeNotFound = awserr.New("UserHomeNotFound", "user home directory not found.", nil)
)

// A SharedCredentialsProvider retrieves credentials from the current user's home
// directory, and keeps track if those credentials are expired.
//
// Profile ini file example: $HOME/.aws/credentials
type SharedCredentialsProvider struct {
	// Path to the shared credentials file.
	//
	// If empty will look for "AWS_SHARED_CREDENTIALS_FILE" env variable. If the
	// env value is empty will default to current user's home directory.
	// Linux/OSX: "$HOME/.aws/credentials"
	// Windows:   "%USERPROFILE%\.aws\credentials"
	Filename string

	// AWS Profile to extract credentials from the shared credentials file. If empty
	// will default to environment variable "AWS_PROFILE" or "default" if
	// environment variable is also not set.
	Profile string

	// retrieved states if the credentials have been successfully retrieved.
	retrieved bool

	// expiration is the session expiration time, if available.
	expiration *time.Time
}

// NewSharedCredentials returns a pointer to a new Credentials object
// wrapping the Profile file provider.
func NewSharedCredentials(filename, profile string) *Credentials {
	return NewCredentials(&SharedCredentialsProvider{
		Filename: filename,
		Profile:  profile,
	})
}

// Retrieve reads and extracts the shared credentials from the current
// users home directory.
func (p *SharedCredentialsProvider) Retrieve() (Value, error) {
	p.retrieved = false

	filename, err := p.filename()
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, err
	}

	creds, expiration, err := loadProfile(filename, p.profile())
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, err
	}

	p.expiration = expiration

	p.retrieved = true
	return creds, nil
}

// IsExpired returns if the shared credentials have expired.
func (p *SharedCredentialsProvider) IsExpired() bool {
	// Use the session expiration time, if available
	if p.expiration != nil {
		return p.expiration.Before(time.Now())
	}

	// Otherwise, mark the creds as expired if they
	// haven't been successfully retrieved
	return !p.retrieved
}

// loadProfiles loads from the file pointed to by shared credentials filename for profile.
// The credentials retrieved from the profile will be returned or error. Error will be
// returned if it fails to read from the file, or the data is invalid.
func loadProfile(filename, profile string) (Value, *time.Time, error) {
	config, err := ini.OpenFile(filename)
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, nil, awserr.New("SharedCredsLoad", "failed to load shared credentials file", err)
	}

	iniProfile, ok := config.GetSection(profile)
	if !ok {
		return Value{ProviderName: SharedCredsProviderName}, nil, awserr.New("SharedCredsLoad", "failed to get profile", nil)
	}

	id := iniProfile.String("aws_access_key_id")
	if len(id) == 0 {
		return Value{ProviderName: SharedCredsProviderName}, nil, awserr.New("SharedCredsAccessKey",
			fmt.Sprintf("shared credentials %s in %s did not contain aws_access_key_id", profile, filename),
			nil)
	}

	secret := iniProfile.String("aws_secret_access_key")
	if len(secret) == 0 {
		return Value{ProviderName: SharedCredsProviderName}, nil, awserr.New("SharedCredsSecret",
			fmt.Sprintf("shared credentials %s in %s did not contain aws_secret_access_key", profile, filename),
			nil)
	}

	var sessionExpiration *time.Time
	sessionExpirationStr := iniProfile.String("aws_session_expiration")
	if len(sessionExpirationStr) != 0 {
		parsedTime, err := time.Parse(time.RFC3339, sessionExpirationStr)
		if err == nil {
			sessionExpiration = &parsedTime
		}
	}


	// Default to empty string if not found
	token := iniProfile.String("aws_session_token")

	return Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
		SessionToken:    token,
		ProviderName:    SharedCredsProviderName,
	}, sessionExpiration, nil
}

// filename returns the filename to use to read AWS shared credentials.
//
// Will return an error if the user's home directory path cannot be found.
func (p *SharedCredentialsProvider) filename() (string, error) {
	if len(p.Filename) != 0 {
		return p.Filename, nil
	}

	if p.Filename = os.Getenv("AWS_SHARED_CREDENTIALS_FILE"); len(p.Filename) != 0 {
		return p.Filename, nil
	}

	if home := shareddefaults.UserHomeDir(); len(home) == 0 {
		// Backwards compatibility of home directly not found error being returned.
		// This error is too verbose, failure when opening the file would of been
		// a better error to return.
		return "", ErrSharedCredentialsHomeNotFound
	}

	p.Filename = shareddefaults.SharedCredentialsFilename()

	return p.Filename, nil
}

// profile returns the AWS shared credentials profile.  If empty will read
// environment variable "AWS_PROFILE". If that is not set profile will
// return "default".
func (p *SharedCredentialsProvider) profile() string {
	if p.Profile == "" {
		p.Profile = os.Getenv("AWS_PROFILE")
	}
	if p.Profile == "" {
		p.Profile = "default"
	}

	return p.Profile
}
