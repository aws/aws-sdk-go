package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/endpoints"
)

// A Session provides a central location to create service clients from and
// store configurations and request handlers for those services.
//
// Sessions are safe to create service clients concurrently, but it is not safe
// to mutate the Session concurrently.
//
// The Session satisfies the service client's client.ClientConfigProvider.
type Session struct {
	Config   *aws.Config
	Handlers request.Handlers
}

// New creates a new instance of the handlers merging in the provided configs
// on top of the SDK's default configurations. Once the Session is created it
// can be mutated to modify Config or Handlers. The Session is safe to be read
// concurrently, but it should not be written to concurrently.
//
// If the AWS_SDK_LOAD_CONFIG environment variable is set to a truthy value
// the shared config (~/.aws/config) will also be loaded. Values from the shared
// credentials file will have taken over those in the shared config.
//
// Example:
//     // Create a Session with the default config and request handlers.
//     sess := session.New()
//
//     // Create a Session with a custom region
//     sess := session.New(&aws.Config{Region: aws.String("us-east-1")})
//
//     // Create a Session, and add additional handlers for all service
//     // clients created with the Session to inherit. Adds logging handler.
//     sess := session.New()
//     sess.Handlers.Send.PushFront(func(r *request.Request) {
//          // Log every request made and its payload
//          logger.Println("Request: %s/%s, Payload: %s", r.ClientInfo.ServiceName, r.Operation, r.Params)
//     })
//
//     // Create a S3 client instance from a Session
//     sess := session.New()
//     svc := s3.New(sess)
func New(cfgs ...*aws.Config) *Session {
	// Load initial config from environment
	envCfg := loadEnvConfig()

	// Load user's shared config, and passed in config
	return newSession(envCfg, cfgs...)
}

// NewFromProfile creates a new Session loading configuration from the SDKs
// defaults, and credentials from the shared credentials file.
//
// Uses the profile to configure which profile to load the shared configuration
// from. This will override the profile specififed in the environment.
//
// Same as New, but specifies the profile that will be loaded from the shared
// configuration files. This function also uses the AWS_SDK_LOAD_CONFIG the
// same as the New function.
func NewFromProfile(profile string, cfgs ...*aws.Config) *Session {
	// Load initial config from environment
	envCfg := loadEnvConfig()
	envCfg.Profile = profile

	// Load user's shared config, and passed in config
	return newSession(envCfg, cfgs...)
}

// NewFromSharedConfig creates a new Session loading configuration from the
// shared configuration, as a base which the passed in configurations will
// be merged on top off.
//
// Same as New, but loads the configuration as if the AWS_SDK_LOAD_CONFIG
// environment variable is set.
func NewFromSharedConfig(cfgs ...*aws.Config) *Session {
	// Load initial config from environment
	envCfg := loadSharedEnvConfig()

	// Load user's shared config, and passed in config
	return newSession(envCfg, cfgs...)
}

// NewFromSharedConfigProfile creates a new Session loading configuration from
// the shared configuration, as a base which the passed in configurations will
// be merged on top off.
//
// Uses the profile to configure which profile to load the shared configuration
// from. This will override the profile specififed in the environment.
//
// Same as New, but loads the configuration as if the AWS_SDK_LOAD_CONFIG
// environment variable is set.
func NewFromSharedConfigProfile(profile string, cfgs ...*aws.Config) *Session {
	// Load initial config from environment
	envCfg := loadSharedEnvConfig()
	envCfg.Profile = profile

	// Load user's shared config, and passed in config
	return newSession(envCfg, cfgs...)
}

func newSession(envCfg envConfig, cfgs ...*aws.Config) *Session {
	cfg := defaults.Config()
	handlers := defaults.Handlers()

	// Load user shared config
	err := loadConfig(envCfg, cfg)

	// Get a merged version of the user provided config to determine if
	// credentials were.
	userCfg := aws.Config{}
	userCfg.MergeIn(cfgs...)

	// Merge in user provided configuration
	cfg.MergeIn(&userCfg)

	// Need to wait until after the user, shared, and default configs are merged,
	// so any log options are set to report the error to.
	if err != nil && cfg.Logger != nil && cfg.LogLevel.Matches(aws.LogDebug) {
		cfg.Logger.Log("DEBUG: failed to load shared config, error", err)
	}

	// Set default credential chain if none found in env/shared config, and
	// not set by the user.
	if cfg.Credentials == credentials.AnonymousCredentials && userCfg.Credentials == nil {
		// Use default credentials chain if none in env/shared config
		cfg.Credentials = defaults.CredChain(cfg, handlers)
	}

	s := &Session{
		Config:   cfg,
		Handlers: handlers,
	}

	initHandlers(s)

	return s
}

func loadConfig(envCfg envConfig, cfg *aws.Config) error {
	// Order config files will be loaded in with later files overwriting
	// previous config file values.
	cfgFiles := []string{envCfg.SharedConfigFile, envCfg.SharedCredentialsFile}
	if !envCfg.EnableSharedConfig {
		// The shared config file (~/.aws/config) is only loaded if instructed
		// to load via the envConfig.EnableSharedConfig (AWS_SDK_LOAD_CONFIG).
		cfgFiles = cfgFiles[1:]
	}

	// Load additional config from file(s)
	sharedCfg, err := loadSharedConfig(envCfg.Profile, cfgFiles...)
	if err != nil {
		return err
	}

	// Region
	if len(envCfg.Region) > 0 {
		cfg.WithRegion(envCfg.Region)
	} else if envCfg.EnableSharedConfig && len(sharedCfg.Region) > 0 {
		cfg.WithRegion(sharedCfg.Region)
	}

	// Configure credentials if known
	if len(envCfg.Creds.AccessKeyID) > 0 {
		cfg.Credentials = credentials.NewCredentials(
			&credentials.StaticProvider{Value: envCfg.Creds},
		)
	} else if len(sharedCfg.Creds.AccessKeyID) > 0 {
		cfg.Credentials = credentials.NewCredentials(
			&credentials.StaticProvider{Value: sharedCfg.Creds},
		)
	}

	return nil
}

func initHandlers(s *Session) {
	// Add the Validate parameter handler if it is not disabled.
	s.Handlers.Validate.Remove(corehandlers.ValidateParametersHandler)
	if !aws.BoolValue(s.Config.DisableParamValidation) {
		s.Handlers.Validate.PushBackNamed(corehandlers.ValidateParametersHandler)
	}
}

// Copy creates and returns a copy of the current Session, coping the config
// and handlers. If any additional configs are provided they will be merged
// on top of the Session's copied config.
//
// Example:
//     // Create a copy of the current Session, configured for the us-west-2 region.
//     sess.Copy(&aws.Config{Region: aws.String("us-west-2")})
func (s *Session) Copy(cfgs ...*aws.Config) *Session {
	newSession := &Session{
		Config:   s.Config.Copy(cfgs...),
		Handlers: s.Handlers.Copy(),
	}

	initHandlers(newSession)

	return newSession
}

// ClientConfig satisfies the client.ConfigProvider interface and is used to
// configure the service client instances. Passing the Session to the service
// client's constructor (New) will use this method to configure the client.
//
// Example:
//     sess := session.New()
//     s3.New(sess)
func (s *Session) ClientConfig(serviceName string, cfgs ...*aws.Config) client.Config {
	s = s.Copy(cfgs...)
	endpoint, signingRegion := endpoints.NormalizeEndpoint(
		aws.StringValue(s.Config.Endpoint), serviceName,
		aws.StringValue(s.Config.Region), aws.BoolValue(s.Config.DisableSSL))

	return client.Config{
		Config:        s.Config,
		Handlers:      s.Handlers,
		Endpoint:      endpoint,
		SigningRegion: signingRegion,
	}
}
