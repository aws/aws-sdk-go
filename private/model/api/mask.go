package api

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strings"
)

const maskFilePath = "./private/model/api/mask.yaml"

type (
	Mask struct {
		Actions map[string]*MaskedAction `yaml:"actions,omitempty"`
	}
	MaskedAction struct {
		Alias       string                            `yaml:"alias,omitempty"`
		DisplayName string                            `yaml:"display_name"`
		Description string                            `yaml:"description,omitempty"`
		Parameters  map[string]*MaskedActionParameter `yaml:"parameters,omitempty"`
	}
	MaskedActionParameter struct {
		ConstValue  string   `yaml:"const_value,omitempty"`
		Alias       string   `yaml:"alias,omitempty"`
		Required    bool     `yaml:"required,omitempty"`
		Type        string   `yaml:"type,omitempty"` // password/date - 2017-07-21/date_time - 2017-07-21T17:32:28Z/date_epoch - 1631399887
		Index       int64    `yaml:"index,omitempty"`
		IsMulti     bool     `yaml:"is_multi,omitempty"` // is this a multi-select field
		Default     string   `yaml:"default,omitempty"`  // override parameter default value
		Description string   `yaml:"description,omitempty"`
		Placeholder *string  `yaml:"placeholder,omitempty"` // type pointer to differentiate between an empty string and nil
		Options     []string `yaml:"options"`               // optional: the option list in case of dropdown\checkbox
	}
)

// ParseMask parses  a mask file and returns a new mask object.
func ParseMask() (mask Mask, err error) {
	var rawMaskData []byte
	rawMaskData, err = ioutil.ReadFile(maskFilePath)
	if err != nil {
		return
	}

	rawMaskData = []byte(strings.ReplaceAll(string(rawMaskData), "\\", ""))
	if err = yaml.Unmarshal(rawMaskData, &mask); err != nil {
		return
	}

	return
}

func getMaskedActionParameter(maskedAction *MaskedAction, paramName string) (param *MaskedActionParameter, ok bool) {
	if maskedAction == nil {
		return nil, false
	}

	param, ok = maskedAction.Parameters[paramName]

	return
}
