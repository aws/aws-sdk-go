package credentials

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/internal/shareddefaults"
	"github.com/go-ini/ini"
)

// ProcessProviderName provides a name of Process provider
const ProcessProviderName = "ProcessProvider"

var (
	// ErrProcessProviderHomeNotFound is emitted when the user directory cannot be found.
	ErrProcessProviderHomeNotFound = awserr.New("UserHomeNotFound", "user home directory not found.", nil)
)

// A ProcessProvider retrieves credentials from a process specified as
// 'credential_process' in a profile, and keeps track if those credentials are
// expired.
//
// Profile ini file example: $HOME/.aws/config
type ProcessProvider struct {
	Expiry

	// Path to the shared config file.
	//
	// If empty will look for "AWS_CONFIG_FILE" env variable. If the
	// env value is empty will default to current user's home directory.
	// Linux/OSX: "$HOME/.aws/config"
	// Windows:   "%USERPROFILE%\.aws\config"
	Filename string

	// AWS Profile to retrieve the 'credential_process' from. If empty
	// will default to environment variable "AWS_PROFILE" or "default" if
	// environment variable is also not set.
	Profile string

	// ExpiryWindow will allow the credentials to trigger refreshing prior to
	// the credentials actually expiring. This is beneficial so race conditions
	// with expiring credentials do not cause request to fail unexpectedly
	// due to ExpiredTokenException exceptions.
	//
	// So a ExpiryWindow of 10s would cause calls to IsExpired() to return true
	// 10 seconds before the credentials are actually expired.
	//
	// If ExpiryWindow is 0 or less it will be ignored.
	ExpiryWindow time.Duration

	// A function to perform the credential_process execution. Can be replaced
	// for testing
	executionFunc func(process string) ([]byte, error)
}

// NewProcessProvider returns a pointer to a new Credentials object
// wrapping the Profile file provider.
func NewProcessProvider(filename, profile string) *Credentials {
	return NewCredentials(&ProcessProvider{
		Filename:      filename,
		Profile:       profile,
		executionFunc: executeCredentialProcess,
	})
}

// Retrieve executes the 'credential_process' and returns the credentials
func (p *ProcessProvider) Retrieve() (Value, error) {
	filename, err := p.filename()
	if err != nil {
		return Value{ProviderName: ProcessProviderName}, err
	}

	creds, err := p.loadProfile(filename, p.profile())
	if err != nil {
		return Value{ProviderName: ProcessProviderName}, err
	}

	return creds, nil
}

// loadProfiles executes the 'credential_process' from the profile of the given
// file, and returns the credentials retrieved. An error will be returned if it
// fails to read from the file, or the data is invalid.
func (p *ProcessProvider) loadProfile(filename, profile string) (Value, error) {
	config, err := ini.Load(filename)
	if err != nil {
		return Value{ProviderName: ProcessProviderName}, awserr.New("ProcessProviderLoad", "failed to load shared config file", err)
	}

	var iniProfile *ini.Section
	if len(profile) == 0 || profile == "default" {
		iniProfile, err = config.GetSection("default")
	} else {
		iniProfile, err = config.GetSection("profile " + profile)
	}
	if err != nil {
		return Value{ProviderName: ProcessProviderName}, awserr.New("ProcessProviderLoad", "failed to get profile", err)
	}

	process, err := iniProfile.GetKey("credential_process")
	if err != nil {
		return Value{ProviderName: ProcessProviderName}, awserr.New("ProcessProviderLoad", "No credential_process specified", err)
	}
	return p.credentialProcess(process.String())
}

type credentialProcessResponse struct {
	Version         int
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string
	SessionToken    string
	Expiration      string
}

func executeCredentialProcess(process string) ([]byte, error) {
	processArgs := strings.Split(process, " ")
	var command string
	var cmdArgs []string

	if len(processArgs) == 1 {
		command = processArgs[0]
	} else if len(processArgs) > 1 {
		command = processArgs[0]
		cmdArgs = processArgs[1:]
	}

	// TODO: Check for executability?
	_, err := os.Stat(command)
	if os.IsNotExist(err) {
		return nil, awserr.New("ProcessProviderLoad", "credential_process "+command+" not found", err)
	} else if err != nil {
		return nil, awserr.New("ProcessProviderLoad", "failed to open credential_process "+command, err)
	}

	// Execute command
	cmd := exec.Command(command, cmdArgs...)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
	if err != nil {
		return nil, awserr.New("ProcessProviderLoad", "Error executing credential_process", err)
	}
	return out, nil
}

func (p *ProcessProvider) credentialProcess(process string) (Value, error) {
	out, err := p.executionFunc(process)
	if err != nil {
		return Value{ProviderName: ProcessProviderName}, err
	}

	// Serialize and validate response
	resp := &credentialProcessResponse{}
	err = json.Unmarshal(out, resp)
	if err != nil {
		return Value{ProviderName: ProcessProviderName}, awserr.New("ProcessProviderLoad", "Error parsing credential_process output: "+string(out), err)
	}
	if resp.Version != 1 {
		return Value{ProviderName: ProcessProviderName}, awserr.New("ProcessProviderLoad", " Version in credential_process output is not 1", err)
	}
	if len(resp.AccessKeyID) == 0 {
		return Value{ProviderName: ProcessProviderName}, awserr.New("ProcessProviderLoad", " Missing required AccessKeyId in credential_process output", err)
	}
	if len(resp.SecretAccessKey) == 0 {
		return Value{ProviderName: ProcessProviderName}, awserr.New("ProcessProviderLoad", " Missing required SecretAccessKey in credential_process output", err)
	}

	// Handle expiration
	if len(resp.Expiration) > 0 {
		expiry, err := time.Parse(time.RFC3339, resp.Expiration)
		if err != nil {
			return Value{ProviderName: ProcessProviderName}, awserr.New("ProcessProviderLoad", "Error parsing expiration of credential_process output: "+resp.Expiration, err)
		}
		p.SetExpiration(expiry, p.ExpiryWindow)
	}

	return Value{
		ProviderName:    ProcessProviderName,
		AccessKeyID:     resp.AccessKeyID,
		SecretAccessKey: resp.SecretAccessKey,
		SessionToken:    resp.SessionToken,
	}, nil

}

// filename returns the filename to use to read AWS shared credentials.
//
// Will return an error if the user's home directory path cannot be found.
func (p *ProcessProvider) filename() (string, error) {
	if len(p.Filename) != 0 {
		return p.Filename, nil
	}

	if p.Filename = os.Getenv("AWS_CONFIG_FILE"); len(p.Filename) != 0 {
		return p.Filename, nil
	}

	if home := shareddefaults.UserHomeDir(); len(home) == 0 {
		// Backwards compatibility of home directly not found error being returned.
		// This error is too verbose, failure when opening the file would of been
		// a better error to return.
		return "", ErrProcessProviderHomeNotFound
	}

	p.Filename = shareddefaults.SharedConfigFilename()

	return p.Filename, nil
}

// profile returns the AWS shared config profile. If empty will read
// environment variable "AWS_PROFILE". If that is not set profile will
// return "default".
func (p *ProcessProvider) profile() string {
	if p.Profile == "" {
		p.Profile = os.Getenv("AWS_PROFILE")
	}
	if p.Profile == "" {
		p.Profile = os.Getenv("AWS_DEFAULT_PROFILE")
	}
	if p.Profile == "" {
		p.Profile = "default"
	}

	return p.Profile
}
