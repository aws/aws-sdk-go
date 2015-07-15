package aws

import (
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/awslog"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// DefaultChainCredentials is a Credentials which will find the first available
// credentials Value from the list of Providers.
//
// This should be used in the default case. Once the type of credentials are
// known switching to the specific Credentials will be more efficient.
var DefaultChainCredentials = credentials.NewChainCredentials(
	[]credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
		&credentials.EC2RoleProvider{ExpiryWindow: 5 * time.Minute},
	})

// The default number of retries for a service. The value of -1 indicates that
// the service specific retry default will be used.
const DefaultRetries = -1

// A LogLevel defines the level logging should be performed at. Used to instruct
// the SDK which statements should be logged.
type LogLevel uint

// Matches returns true if the v LogLevel is enabled by this LogLevel. Should be
// used with logging sub levels.
func (l LogLevel) Matches(v LogLevel) bool {
	return l&v == v
}

// AtLeast returns true if this LogLevel is at least high enough to satisfies v.
func (l LogLevel) AtLeast(v LogLevel) bool {
	return l >= v
}

const (
	// LogOff states that no logging should be performed by the SDK. This is the
	// default state of the SDK, and should be use to disable all logging.
	LogOff LogLevel = iota * 0x1000

	// LogDebug state that debug output should be logged by the SDK. This should
	// be used to inspect request made and responses received.
	LogDebug
)

// Debug Logging Sub Levels
const (
	// LogDebugWithSigning states that the SDK should log request signing and
	// presigning events. This should be used to log the signing details of
	// requests for debugging. Will also enable LogDebug.
	LogDebugWithSigning LogLevel = LogDebug | (1 << iota)

	// LogDebugWithHTTPBody states the SDK should log HTTP request and response
	// HTTP bodys in addition to the headers and path. This should be used to
	// see the body content of requests and responses made while using the SDK
	// Will also enable LogDebug.
	LogDebugWithHTTPBody
)

// DefaultConfig is the default all service configuration will be based off of.
// By default, all clients use this structure for initialization options unless
// a custom configuration object is passed in.
//
// You may modify this global structure to change all default configuration
// in the SDK. Note that configuration options are copied by value, so any
// modifications must happen before constructing a client.
var DefaultConfig = &Config{
	Credentials:             DefaultChainCredentials,
	Endpoint:                "",
	Region:                  os.Getenv("AWS_REGION"),
	DisableSSL:              false,
	HTTPClient:              http.DefaultClient,
	LogLevel:                LogOff,
	Logger:                  awslog.NewDefaultLogger(),
	MaxRetries:              DefaultRetries,
	DisableParamValidation:  false,
	DisableComputeChecksums: false,
	S3ForcePathStyle:        false,
}

// A Config provides service configuration for service clients. By default,
// all clients will use the {DefaultConfig} structure.
type Config struct {
	// The credentials object to use when signing requests. Defaults to
	// {DefaultChainCredentials}.
	Credentials *credentials.Credentials

	// An optional endpoint URL (hostname only or fully qualified URI)
	// that overrides the default generated endpoint for a client. Set this
	// to `""` to use the default generated endpoint.
	//
	// @note You must still provide a `Region` value when specifying an
	//   endpoint for a client.
	Endpoint string

	// The region to send requests to. This parameter is required and must
	// be configured globally or on a per-client basis unless otherwise
	// noted. A full list of regions is found in the "Regions and Endpoints"
	// document.
	//
	// @see http://docs.aws.amazon.com/general/latest/gr/rande.html
	//   AWS Regions and Endpoints
	Region string

	// Set this to `true` to disable SSL when sending requests. Defaults
	// to `false`.
	DisableSSL bool

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`.
	HTTPClient *http.Client

	// An integer value representing the logging level. The default log level
	// is zero (LogOff), which represents no logging. To enable logging set
	// to a LogLevel Value.
	LogLevel LogLevel

	// The logger writer interface to write logging messages to. Defaults to
	// standard out.
	Logger awslog.Logger

	// The maximum number of times that a request will be retried for failures.
	// Defaults to -1, which defers the max retry setting to the service specific
	// configuration.
	MaxRetries int

	// Disables semantic parameter validation, which validates input for missing
	// required fields and/or other semantic request input errors.
	DisableParamValidation bool

	// Disables the computation of request and response checksums, e.g.,
	// CRC32 checksums in Amazon DynamoDB.
	DisableComputeChecksums bool

	// Set this to `true` to force the request to use path-style addressing,
	// i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will
	// use virtual hosted bucket addressing when possible
	// (`http://BUCKET.s3.amazonaws.com/KEY`).
	//
	// @note This configuration option is specific to the Amazon S3 service.
	// @see http://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html
	//   Amazon S3: Virtual Hosting of Buckets
	S3ForcePathStyle bool
}

// Copy will return a shallow copy of the Config object.
func (c Config) Copy() Config {
	dst := Config{}
	dst.Credentials = c.Credentials
	dst.Endpoint = c.Endpoint
	dst.Region = c.Region
	dst.DisableSSL = c.DisableSSL
	dst.HTTPClient = c.HTTPClient
	dst.LogLevel = c.LogLevel
	dst.Logger = c.Logger
	dst.MaxRetries = c.MaxRetries
	dst.DisableParamValidation = c.DisableParamValidation
	dst.DisableComputeChecksums = c.DisableComputeChecksums
	dst.S3ForcePathStyle = c.S3ForcePathStyle

	return dst
}

// Merge merges the newcfg attribute values into this Config. Each attribute
// will be merged into this config if the newcfg attribute's value is non-zero.
// Due to this, newcfg attributes with zero values cannot be merged in. For
// example bool attributes cannot be cleared using Merge, and must be explicitly
// set on the Config structure.
func (c Config) Merge(newcfg *Config) *Config {
	if newcfg == nil {
		return &c
	}

	cfg := Config{}

	if newcfg.Credentials != nil {
		cfg.Credentials = newcfg.Credentials
	} else {
		cfg.Credentials = c.Credentials
	}

	if newcfg.Endpoint != "" {
		cfg.Endpoint = newcfg.Endpoint
	} else {
		cfg.Endpoint = c.Endpoint
	}

	if newcfg.Region != "" {
		cfg.Region = newcfg.Region
	} else {
		cfg.Region = c.Region
	}

	if newcfg.DisableSSL {
		cfg.DisableSSL = newcfg.DisableSSL
	} else {
		cfg.DisableSSL = c.DisableSSL
	}

	if newcfg.HTTPClient != nil {
		cfg.HTTPClient = newcfg.HTTPClient
	} else {
		cfg.HTTPClient = c.HTTPClient
	}

	if newcfg.LogLevel != 0 {
		cfg.LogLevel = newcfg.LogLevel
	} else {
		cfg.LogLevel = c.LogLevel
	}

	if newcfg.Logger != nil {
		cfg.Logger = newcfg.Logger
	} else {
		cfg.Logger = c.Logger
	}

	if newcfg.MaxRetries != DefaultRetries {
		cfg.MaxRetries = newcfg.MaxRetries
	} else {
		cfg.MaxRetries = c.MaxRetries
	}

	if newcfg.DisableParamValidation {
		cfg.DisableParamValidation = newcfg.DisableParamValidation
	} else {
		cfg.DisableParamValidation = c.DisableParamValidation
	}

	if newcfg.DisableComputeChecksums {
		cfg.DisableComputeChecksums = newcfg.DisableComputeChecksums
	} else {
		cfg.DisableComputeChecksums = c.DisableComputeChecksums
	}

	if newcfg.S3ForcePathStyle {
		cfg.S3ForcePathStyle = newcfg.S3ForcePathStyle
	} else {
		cfg.S3ForcePathStyle = c.S3ForcePathStyle
	}

	return &cfg
}
