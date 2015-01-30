package aws

import (
	"net/http"
	"os"
)

var DefaultConfig = &Config{
	Credentials: DefaultCreds(),
	Endpoint:    "",
	Region:      os.Getenv("AWS_REGION"),
	DisableSSL:  false,
	ManualSend:  false,
	HTTPClient:  http.DefaultClient,
	LogLevel:    0,
}

type Config struct {
	Credentials CredentialsProvider
	Endpoint    string
	Region      string
	DisableSSL  bool
	ManualSend  bool
	HTTPClient  *http.Client
	LogLevel    uint
}

func (c Config) Merge(newcfg *Config) *Config {
	cfg := Config{}

	if newcfg != nil && newcfg.Credentials != nil {
		cfg.Credentials = newcfg.Credentials
	} else {
		cfg.Credentials = c.Credentials
	}

	if newcfg != nil && newcfg.Region != "" {
		cfg.Region = newcfg.Region
	} else {
		cfg.Region = c.Region
	}

	if newcfg != nil && newcfg.DisableSSL {
		cfg.DisableSSL = newcfg.DisableSSL
	} else {
		cfg.DisableSSL = c.DisableSSL
	}

	if newcfg != nil && newcfg.ManualSend {
		cfg.ManualSend = newcfg.ManualSend
	} else {
		cfg.ManualSend = c.ManualSend
	}

	if newcfg != nil && newcfg.HTTPClient != nil {
		cfg.HTTPClient = newcfg.HTTPClient
	} else {
		cfg.HTTPClient = c.HTTPClient
	}

	if newcfg != nil && newcfg.LogLevel != 0 {
		cfg.LogLevel = newcfg.LogLevel
	} else {
		cfg.LogLevel = c.LogLevel
	}

	return &cfg
}
