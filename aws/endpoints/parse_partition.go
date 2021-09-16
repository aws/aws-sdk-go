package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type cfgPartitions struct {
	Partitions partitions `json:"partitions"`
}

func configPartitions() partitions {
	configFile := os.Getenv("ENDPOINT_CONFIG_FILE")
	if configFile == "" {
		return partitions{}
	}

	jsonFile, err := os.Open(configFile)
	if err != nil {
		fmt.Errorf("Error while opening file", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cparts cfgPartitions

	json.Unmarshal(byteValue, &cparts)

	return cparts.Partitions
}