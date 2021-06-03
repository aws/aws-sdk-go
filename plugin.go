package main

import (
	"github.com/aws/aws-sdk-go/implementation"
	blink_sdk "github.com/blinkops/blink-sdk"
	"github.com/blinkops/blink-sdk/plugin/config"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

func main() {

	log.SetLevel(log.DebugLevel)

	// Get the current directory.
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Error("Failed getting current directory: ", err)
		panic(err)
	}

	log.Info("Current directory is: ", currentDirectory)

	// Initialize the configuration.
	err = os.Setenv(config.ConfigurationPathEnvVar, path.Join(currentDirectory, "config.yaml"))
	if err != nil {
		log.Error("Failed to set configuration env variable: ", err)
		panic(err)
	}

	plugin, err := implementation.NewAWSPlugin(currentDirectory)
	if err != nil {
		log.Error("Failed to create plugin implementation: ", err)
		panic(err)
	}

	err = blink_sdk.Start(plugin)
	if err != nil {
		log.Fatal("Error during server startup: ", err)
	}
}
