package aws

import (
	"io"
	"net/http"
	"os"
	"time"

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
	LogHTTPBody:             false,
	LogLevel:                0,
	Logger:                  os.Stdout,
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

	// Set this to `true` to also log the body of the HTTP requests made by the
	// client.
	//
	// @note `LogLevel` must be set to a non-zero value in order to activate
	//   body logging.
	LogHTTPBody bool

	// An integer value representing the logging level. The default log level
	// is zero (0), which represents no logging. Set to a non-zero value to
	// perform logging.
	LogLevel uint

	// The logger writer interface to write logging messages to. Defaults to
	// standard out.
	Logger io.Writer

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
	dst.LogHTTPBody = c.LogHTTPBody
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

	if newcfg.LogHTTPBody {
		cfg.LogHTTPBody = newcfg.LogHTTPBody
	} else {
		cfg.LogHTTPBody = c.LogHTTPBody
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
