package template

import "encoding/json"

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
type Template struct {
	AWSTemplateFormatVersion string                 `json:",omitempty"`
	Description              string                 `json:",omitempty"`
	Metadata                 map[string]interface{} `json:",omitempty"`
	Parameters               map[string]Parameter   `json:",omitempty"`
	Mappings                 map[string]Mapping     `json:",omitempty"`
	Conditions               map[string]Condition   `json:",omitempty"`
	Resources                map[string]Resource    `json:",omitempty"`
	Outputs                  map[string]Output      `json:",omitempty"`
}

type Parameter struct {
	Type                  string
	Default               string   `json:",omitempty"`
	NoEcho                bool     `json:",omitempty,string"`
	AllowedValues         []string `json:",omitempty"`
	AllowedPattern        string   `json:",omitempty"`
	MaxLength             int      `json:",omitempty,string"`
	MinLength             int      `json:",omitempty,string"`
	Description           string   `json:",omitempty"`
	ConstraintDescription string   `json:",omitempty"`
}

type Mapping map[string]map[string]string

type Condition map[string]interface{}

type Resource struct {
	Type       string
	Properties map[string]interface{}
	Metadata   map[string]interface{} `json:",omitempty"`
}

type Output struct {
	Description string `json:",omitempty"`
	Value       interface{}
}

func (t *Template) String() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
