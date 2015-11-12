// Package template provides types and functions that enable programmatic generation of CloudFormation templates
package template

import "encoding/json"

// Template defines a CloudFormation template.
//
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

// Parameter defines a CloudFormation template parameter
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html
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

// Mapping defines a CloudFormation template mapping
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html
type Mapping map[string]map[string]string

// Condition defines a CloudFormation template condition
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/conditions-section-structure.html
type Condition map[string]interface{}

// Resource defines a CloudFormation template resource
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resources-section-structure.html
type Resource struct {
	Type       string
	Properties map[string]interface{} `json:",omitempty"`
	Metadata   map[string]interface{} `json:",omitempty"`

	DependsOn      interface{} `json:",omitempty"`
	CreationPolicy interface{} `json:",omitempty"`
	UpdatePolicy   interface{} `json:",omitempty"`
	DeletionPolicy interface{} `json:",omitempty"`
}

// Output defines a CloudFormation template output
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html
type Output struct {
	Description string `json:",omitempty"`
	Value       interface{}
}

// String returns a JSON representation of the Template suitable for use
// in CloudFormation requests such as CreateStack and UpdateStack
func (t *Template) String() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
