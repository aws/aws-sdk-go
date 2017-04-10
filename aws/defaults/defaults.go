// Package defaults is a collection of helpers to retrieve the SDK's default
// configuration and handlers.
//
// Generally this package shouldn't be used directly, but session.Session
// instead. This package is useful when you need to reset the defaults
// of a session or service client to the SDK defaults before setting
// additional parameters.
package defaults

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/credentials/endpointcreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/request"
)

// A Defaults provides a collection of default values for SDK clients.
type Defaults struct {
	Config   *aws.Config
	Handlers request.Handlers
}

// Get returns the SDK's default values with Config and handlers pre-configured.
func Get() Defaults {
	cfg := Config()
	handlers := Handlers()
	cfg.Credentials = CredChain(cfg, handlers)

	return Defaults{
		Config:   cfg,
		Handlers: handlers,
	}
}

// Config returns the default configuration without credentials.
// To retrieve a config with credentials also included use
// `defaults.Get().Config` instead.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the configuration of an
// existing service client or session.
func Config() *aws.Config {
	return aws.NewConfig().
		WithCredentials(credentials.AnonymousCredentials).
		WithRegion(os.Getenv("AWS_REGION")).
		WithHTTPClient(http.DefaultClient).
		WithMaxRetries(aws.UseServiceDefaultRetries).
		WithLogger(aws.NewDefaultLogger()).
		WithLogLevel(aws.LogOff).
		WithEndpointResolver(endpoints.DefaultResolver())
}

// Handlers returns the default request handlers.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the request handlers of an
// existing service client or session.
func Handlers() request.Handlers {
	var handlers request.Handlers

	handlers.Validate.PushBackNamed(corehandlers.ValidateEndpointHandler)
	handlers.Validate.AfterEachFn = request.HandlerListStopOnError
	handlers.Build.PushBackNamed(corehandlers.SDKVersionUserAgentHandler)
	handlers.Build.AfterEachFn = request.HandlerListStopOnError
	handlers.Sign.PushBackNamed(corehandlers.BuildContentLengthHandler)
	handlers.Send.PushBackNamed(corehandlers.ValidateReqSigHandler)
	handlers.Send.PushBackNamed(corehandlers.SendHandler)
	handlers.AfterRetry.PushBackNamed(corehandlers.AfterRetryHandler)
	handlers.ValidateResponse.PushBackNamed(corehandlers.ValidateResponseHandler)

	return handlers
}

// CredChain returns the default credential chain.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the credentials of an
// existing service client or session's Config.
func CredChain(cfg *aws.Config, handlers request.Handlers) *credentials.Credentials {
	return credentials.NewCredentials(&credentials.ChainProvider{
		VerboseErrors: aws.BoolValue(cfg.CredentialsChainVerboseErrors),
		Providers: []credentials.Provider{
			&credentials.EnvProvider{},
			&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
			RemoteCredProvider(*cfg, handlers),
		},
	})
}

const (
	httpProviderEnvVar     = "AWS_CONTAINER_CREDENTIALS_FULL_URI"
	ecsCredsProviderEnvVar = "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI"
)

// RemoteCredProvider returns a credenitials provider for the default remote
// endpoints such as EC2 or ECS Roles.
func RemoteCredProvider(cfg aws.Config, handlers request.Handlers) credentials.Provider {
	if u := os.Getenv(httpProviderEnvVar); len(u) > 0 {
		parsed, err := url.Parse(u)
		// Need to log error because this API did not expose an error
		if err != nil {
			log(cfg.Logger,
				"Ignoring,", httpProviderEnvVar, "failed to parse url", err)
		} else if host := aws.URLHostname(parsed); !(host == "localhost" || host == "127.0.0.1") {
			log(cfg.Logger,
				"Ignoring,", httpProviderEnvVar, "specified URL with invalid hostname",
				host, ", only localhost and 127.0.0.1 are valid.")
		} else {
			return httpCredProvider(cfg, handlers, u)
		}
	}

	if uri := os.Getenv(ecsCredsProviderEnvVar); len(uri) > 0 {
		u := fmt.Sprintf("http://169.254.170.2%s", uri)
		return httpCredProvider(cfg, handlers, u)
	}

	return ec2RoleProvider(cfg, handlers)
}

func log(logger aws.Logger, msg ...interface{}) {
	if logger == nil {
		return
	}
	logger.Log(msg...)
}

func httpCredProvider(cfg aws.Config, handlers request.Handlers, u string) credentials.Provider {
	return endpointcreds.NewProviderClient(cfg, handlers, u,
		func(p *endpointcreds.Provider) {
			p.ExpiryWindow = 5 * time.Minute
		},
	)
}

func ec2RoleProvider(cfg aws.Config, handlers request.Handlers) credentials.Provider {
	resolver := cfg.EndpointResolver
	if resolver == nil {
		resolver = endpoints.DefaultResolver()
	}

	e, _ := resolver.EndpointFor(endpoints.Ec2metadataServiceID, "")
	return &ec2rolecreds.EC2RoleProvider{
		Client:       ec2metadata.NewClient(cfg, handlers, e.URL, e.SigningRegion),
		ExpiryWindow: 5 * time.Minute,
	}
}
